# Testing Patterns

## Dominant Testing Style

The repository uses a **table-driven test pattern** with an AAA (Arrange-Act-Assert) structure and the **goconvey** assertion library.

### Framework & Libraries

| Tool | Purpose |
|------|---------|
| `testing` (stdlib) | Test runner |
| `github.com/smartystreets/goconvey` | BDD-style assertions |
| `github.com/smarty/assertions` | Assertion functions |
| `coretests.GetAssert` | Custom assertion wrapper |
| `coretests/args.Map` | Typed test input map |
| `coretests/coretestcases.CaseV1` | Test case structure with ArrangeInput + expected |

### Test File Organization

- Integration tests live in `tests/integratedtests/`.
- Per-package test directories: `tests/integratedtests/{pkg}tests/`.
- Test case data files: `*_testcases.go` (separate from test logic `*_test.go`).

## Template Test Structure

### Standard (positional string assertions)

```go
package sometests

import (
    "testing"
    "github.com/alimtvnetwork/core/coretests"
    "github.com/alimtvnetwork/core/coretests/args"
)

// Test cases defined in a separate _testcases.go file
var myTestCases = []coretestcases.CaseV1{
    {
        ArrangeInput: args.Map{
            "when":   "given valid input",
            "actual": "hello",
            "expect": "HELLO",
        },
        ExpectedLines: []string{"HELLO"},
    },
}

func Test_MyFunction(t *testing.T) {
    for caseIndex, testCase := range myTestCases {
        // Arrange
        input := testCase.ArrangeInput.(args.Map)

        // Act
        result := MyFunction(input.Actual())

        // Assert
        testCase.ShouldBeEqual(t, caseIndex, result)
    }
}
```

### Map-Based (self-documenting multi-property assertions)

Use `args.Map` as `ExpectedInput` when asserting multiple properties. This eliminates
magic indices and produces labeled failure output (e.g., `"isZero : false"` instead of `"false"`).

```go
// _testcases.go — raw typed values, no fmt.Sprintf
var variantTestCases = []coretestcases.CaseV1{
    {
        Title: "New creates Variant with correct value",
        ArrangeInput: args.Map{
            "when":  "given byte value 5",
            "input": 5,
        },
        ExpectedInput: args.Map{
            "value":     5,
            "isZero":    false,
            "isInvalid": false,
            "isValid":   true,
        },
    },
}

// _test.go — pass raw values, CompileToStrings handles conversion
func Test_Variant(t *testing.T) {
    for caseIndex, tc := range variantTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        inputVal, _ := input.GetAsInt("input")

        // Act
        v := bytetype.New(byte(inputVal))
        actual := args.Map{
            "value":     v.ValueInt(),
            "isZero":    v.IsZero(),
            "isInvalid": v.IsInvalid(),
            "isValid":   v.IsValid(),
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

**Key methods:**
- `args.Map.CompileToStrings()` — sorted `"key : value"` lines using `%v` format
- `CaseV1.ShouldBeEqualMap(t, idx, actual)` — compiles both maps and compares
- `CaseV1.ShouldBeEqualMapFirst(t, actual)` — convenience for non-looping tests (hardcodes `caseIndex=0`)
- `CaseV1.ShouldBeEqualFirst(t, values...)` — convenience for non-looping positional tests
- `CaseV1.ExpectedAsMap()` — type-asserts `ExpectedInput` to `args.Map`

> **Migration status:** 39 of ~138 testcase files (28.3%) now use `args.Map`. See [migration tracker](../13-app-issues/testing/05-args-map-migration-status.md).

## Nil-Receiver Safety Pattern (CaseNilSafe)

The `CaseNilSafe` structure provides a **compile-time-safe, table-driven** pattern for testing nil-receiver safety on pointer receiver methods.

### Why CaseNilSafe

- **Compile-time safety**: Uses direct method references (`(*Type).Method`) — renaming a method causes a build error instead of a silent test failure.
- **Unified pattern**: Replaces 4 different nil-test styles (inline `t.Error`, CaseV1, custom wrappers, GenericGherkins) with one consistent structure.
- **Panic recovery**: Built-in `recover()` via `results.InvokeWithPanicRecovery` — tests can assert both safe and panicking methods.

### Structure

```go
type CaseNilSafe struct {
    Title    string     // scenario name
    Func     any        // direct method ref: (*Type).Method
    Args     []any      // optional input arguments
    Expected args.Map   // expected outcome map
}
```

### Test Case File (`_NilReceiver_testcases.go`)

```go
var myStructNilReceiverTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyStruct).IsValid,
        Expected: args.Map{
            "value":    "false",
            "panicked": false,
        },
    },
    {
        Title: "Parse on nil with args returns zero",
        Func:  (*MyStruct).Parse,
        Args:  []any{"hello"},
        Expected: args.Map{
            "value":       "0",
            "panicked":    false,
            "hasError":    false,
            "returnCount": 2,
        },
    },
    {
        Title: "UnsafeMethod on nil panics",
        Func:  (*MyStruct).UnsafeMethod,
        Expected: args.Map{
            "panicked": true,
        },
    },
}
```

### Test Logic File (`_NilReceiver_test.go`)

```go
func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilReceiverTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

### Key Methods

| Method | Purpose |
|--------|---------|
| `MethodName()` | Extracts method name via reflection |
| `CaseTitle()` | Returns Title, falling back to MethodName |
| `InvokeNil()` | Calls method with nil receiver + panic recovery |
| `Invoke(receiver)` | Calls with a specific receiver |
| `ShouldBeSafe(t, idx)` | One-liner: invoke nil → assert ToMap vs Expected |
| `ShouldBeEqualMap(t, idx, actual)` | Manual assertion with custom actual map |

### Expected Map Keys

| Key | Type | Description |
|-----|------|-------------|
| `"value"` | `string` | Primary return value via `%v` |
| `"panicked"` | `bool` | Whether invocation panicked |
| `"isSafe"` | `bool` | `!panicked && !hasError` |
| `"hasError"` | `bool` | Whether an error was returned |
| `"returnCount"` | `int` | Number of return values |

Expected can be a **subset** — only keys present in Expected are compared.

### Migration Guide

**Priority 1 — Inline `if` + `t.Error`** (violates "no raw `t.Error`" standard):
1. Create `_NilReceiver_testcases.go` with `CaseNilSafe` entries
2. Replace individual `Test_*_NilReceiver_*` functions with one loop
3. Delete old inline tests

**Priority 2 — CaseV1 nil sections**:
1. Extract nil-specific CaseV1 entries into `CaseNilSafe` entries
2. Use method references instead of `ArrangeInput` maps

**Priority 3 — Custom wrappers**:
1. Replace custom `IsNilReceiver` fields with `CaseNilSafe.InvokeNil()`

### Related Docs

- [CaseNilSafe Design Document](./designs/CaseNilSafe-design.md)
- [results package](./folders/07-coretests.md)

## Best Patterns Observed

1. **Separation of test data and test logic** — `_testcases.go` files keep data separate.
2. **Consistent AAA structure** — every test follows Arrange-Act-Assert.
3. **Index-based case tracking** — `caseIndex` enables precise failure identification.
4. **Formatted output** — `GetAssert.Quick` provides readable failure messages.

## Anti-Patterns to Avoid

1. **Inline test data in test functions** — always use separate testcases files.
2. **Skipping the Arrange comment** — always label AAA sections.
3. **Ignoring caseIndex** — always pass it for debugging.
4. **Direct `t.Fatal` without context** — use `ShouldBeEqual` or `GetAssert` for rich output.

## Coverage Expectations

No formal coverage requirements are documented. Recommended minimum: critical packages (`chmodhelper`, `errcore`, `coredata/corestr`, `converters`) should have ≥80% coverage.

## Related Docs

- [coretests folder spec](./folders/07-coretests.md)
- [Repo Overview](./00-repo-overview.md)
- [CaseNilSafe Design](./designs/CaseNilSafe-design.md)
