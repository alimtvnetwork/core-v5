# Testing Guidelines — How to Write Tests for `core`

> This document is the authoritative guide for writing unit and integration tests in the `core` repository.
> Follow this exactly so that tests are consistent, readable, and maintainable across all packages.

## Table of Contents

1. [Philosophy](#philosophy)
2. [Directory Structure](#directory-structure)
3. [File Naming Conventions](#file-naming-conventions)
4. [Function Naming Conventions](#function-naming-conventions)
5. [Core Testing Framework](#core-testing-framework)
6. [Test Case Structure (CaseV1)](#test-case-structure-casev1)
7. [Complete Example: Simple Package Test](#complete-example-simple-package-test)
8. [Complete Example: With Custom Test Wrapper](#complete-example-with-custom-test-wrapper)
9. [Assertion Methods Reference](#assertion-methods-reference)
10. [Type Verification](#type-verification)
11. [Error Testing Patterns](#error-testing-patterns)
12. [Anti-Patterns](#anti-patterns)
13. [Checklist for New Tests](#checklist-for-new-tests)

---

## Philosophy

- **AAA (Arrange-Act-Assert)**: Every test follows this structure with explicit `// Arrange`, `// Act`, `// Assert` comments.
- **Table-driven**: All test inputs are defined as `[]coretestcases.CaseV1` slices in separate `_testcases.go` files.
- **Separation of data and logic**: Test data lives in `*_testcases.go`, test execution logic lives in `*_test.go`.
- **String-line comparison**: Results are converted to `[]string` lines and compared line-by-line for readable diffs.
- **Index-based tracking**: Every test loop passes `caseIndex` for precise failure identification.

---

## Directory Structure

All tests live under `tests/integratedtests/` in per-package subdirectories:

```
tests/
└── integratedtests/
    ├── {package}tests/              ← one directory per package under test
    │   ├── {Function}_test.go       ← test execution logic
    │   ├── {Function}_testcases.go  ← test case data ([]CaseV1)
    │   ├── testWrapper.go           ← optional: custom test wrapper struct
    │   └── helpers.go               ← optional: shared test helpers
    ├── GetAssert_Quick_test.go      ← root-level tests for coretests itself
    └── GetAssert_testcases.go
```

### Naming the directory

- Package `chmodhelper` → directory `chmodhelpertests`
- Package `coredata` → directory `coredatatests`
- Package `simplewrap` → directory `simplewraptests`
- Package `conditional` → directory `conditionaltests`

**Rule**: `{packagename}tests` — always lowercase, always suffixed with `tests`.

---

## File Naming Conventions

| File Type | Pattern | Example |
|-----------|---------|---------|
| Test logic | `{StructOrFunc}_test.go` | `DirFilesWithContent_test.go` |
| Test cases | `{StructOrFunc}_testcases.go` | `DirFilesWithContent_testcases.go` |
| Test wrapper | `{StructOrFunc}TestWrapper.go` | `EnumImplDynamicMapTestWrapper.go` |
| Shared helpers | `helpers.go` or descriptive name | `pathInstructionsV3.go` |

**Critical**: Test case data is NEVER inline in `_test.go` files. Always in separate `_testcases.go` files.

---

## Function Naming Conventions

```go
// Pattern: Test_{FunctionOrStructName}_{Behavior}_Verification
func Test_GetAssert_Quick_Verification(t *testing.T) { ... }

// Pattern: Test_{StructName}_{MethodName}_Verification
func Test_FuncWrap_Creation_Verification(t *testing.T) { ... }

// Pattern: Test_{Action}_{Subject}
func Test_DirFilesWithContent_CreateRead(t *testing.T) { ... }
```

**Rules**:
- Always start with `Test_` (Go convention)
- Use PascalCase for struct/function names
- End with `_Verification` or a descriptive suffix
- Underscores separate logical parts: `Test_{What}_{How}`

---

## Core Testing Framework

### Key Imports

```go
import (
    "testing"

    "gitlab.com/auk-go/core/coretests"              // GetAssert, BaseTestCase
    "gitlab.com/auk-go/core/coretests/args"          // Map, One, ThreeFunc, Holder
    "gitlab.com/auk-go/core/coretests/coretestcases" // CaseV1
    "gitlab.com/auk-go/core/coredata/corestr"        // SimpleSlice for building output lines
)
```

### Framework Components

| Component | Purpose | Usage |
|-----------|---------|-------|
| `coretestcases.CaseV1` | Test case struct with Title, ArrangeInput, ExpectedInput | Define test data |
| `coretests.GetAssert` | Assertion helper singleton | `GetAssert.ToStrings()`, `GetAssert.Quick()` |
| `args.Map` | `map[string]any` with typed getters | `input.When()`, `input.Actual()`, `input.Expect()` |
| `args.One` | Single-value holder with First, Second, Third | Simple struct inputs |
| `args.ThreeFunc` | Function holder with 3 params + WorkFunc | Dynamic function invocation tests |
| `coretests.BaseTestCase` | Base struct for custom test wrappers | Embed in custom wrappers |
| `coretests.VerifyTypeOf` | Type verification for Arrange/Actual/Expected | Catches type mismatches |

---

## Test Case Structure (CaseV1)

```go
type CaseV1 struct {
    Title         string              // Human-readable test description
    ArrangeInput  any                 // Input data (args.Map, args.One, custom struct, etc.)
    ActualInput   any                 // Set during Act phase via SetActual()
    ExpectedInput any                 // Expected output (usually []string)
    VerifyTypeOf  *coretests.VerifyTypeOf  // Optional: type validation
    IsEnable      issetter.Value      // Optional: set to issetter.False to skip
    HasError      bool                // Optional: flag if error is expected
    HasPanic      bool                // Optional: flag if panic is expected
}
```

### Using `args.Map` for arrange input

```go
ArrangeInput: args.Map{
    "when":   "given a valid name",
    "actual": "hello",
    "expect": "HELLO",
},
```

Access in test:
```go
input := testCase.ArrangeInput.(args.Map)
when := input.When()          // "given a valid name"
actual := input.Actual()      // "hello" (as any)
expected := input.Expect()    // "HELLO" (as any)
counter := input.GetAsIntDefault("counter", 0)
```

### Using `args.One` for struct inputs

```go
ArrangeInput: []args.One{
    {First: someInstruction},
},
```

### Using `args.ThreeFunc` for function testing

```go
ArrangeInput: args.ThreeFunc{
    First:    "param1",
    Second:   "param2",
    Third:    "param3",
    WorkFunc: myTargetFunction,
},
```

---

## Complete Example: Simple Package Test

### Step 1: Create test case file — `tests/integratedtests/conditionaltests/If_testcases.go`

```go
package conditionaltests

import (
    "gitlab.com/auk-go/core/coretests/args"
    "gitlab.com/auk-go/core/coretests/coretestcases"
)

var ifTestCases = []coretestcases.CaseV1{
    {
        Title: "If[string] - true condition returns trueValue",
        ArrangeInput: args.Map{
            "when":      "condition is true",
            "condition": true,
            "trueVal":   "yes",
            "falseVal":  "no",
        },
        ExpectedInput: []string{
            "yes",
        },
    },
    {
        Title: "If[string] - false condition returns falseValue",
        ArrangeInput: args.Map{
            "when":      "condition is false",
            "condition": false,
            "trueVal":   "yes",
            "falseVal":  "no",
        },
        ExpectedInput: []string{
            "no",
        },
    },
}
```

### Step 2: Create test file — `tests/integratedtests/conditionaltests/If_test.go`

```go
package conditionaltests

import (
    "testing"

    "gitlab.com/auk-go/core/conditional"
    "gitlab.com/auk-go/core/coretests"
    "gitlab.com/auk-go/core/coretests/args"
)

func Test_If_Verification(t *testing.T) {
    for caseIndex, testCase := range ifTestCases {
        // Arrange
        input := testCase.ArrangeInput.(args.Map)
        condition := input["condition"].(bool)
        trueVal := input["trueVal"].(string)
        falseVal := input["falseVal"].(string)
        toStringsConv := coretests.GetAssert.ToStrings

        // Act
        result := conditional.If[string](condition, trueVal, falseVal)

        // Assert
        actualLines := toStringsConv(result)
        testCase.ShouldBeEqual(t, caseIndex, actualLines...)
    }
}
```

---

## Complete Example: With Custom Test Wrapper

When the `ArrangeInput` needs type-safe casting, create a wrapper:

### `tests/integratedtests/mypkgtests/MyTestWrapper.go`

```go
package mypkgtests

import (
    "gitlab.com/auk-go/core/coretests"
    "gitlab.com/auk-go/core/errcore"
)

type MyTestWrapper struct {
    coretests.BaseTestCase
}

func (it MyTestWrapper) ArrangeAsMyInput() MyInput {
    casted, isSuccess := it.ArrangeInput.(MyInput)
    if !isSuccess {
        errcore.HandleErrMessage("casting failed to MyInput")
    }
    return casted
}
```

### Usage in test

```go
func Test_MyFeature_Verification(t *testing.T) {
    for caseIndex, testCase := range myTestCases {
        // Arrange
        input := testCase.ArrangeAsMyInput()

        // Act
        result := mypackage.DoSomething(input.Value)

        // Assert
        testCase.ShouldBeEqual(t, caseIndex, result...)
    }
}
```

---

## Assertion Methods Reference

`CaseV1` provides these assertion methods:

| Method | Comparison | Use When |
|--------|-----------|----------|
| `ShouldBeEqual(t, idx, lines...)` | Exact string equality | Default — most tests |
| `ShouldBeTrimEqual(t, idx, lines...)` | Trim + equal | Whitespace-insensitive |
| `ShouldBeSortedEqual(t, idx, lines...)` | Sort + trim + equal | Order doesn't matter |
| `ShouldContains(t, idx, lines...)` | Contains check | Partial matching |
| `ShouldStartsWith(t, idx, lines...)` | Starts-with check | Prefix matching |
| `ShouldEndsWith(t, idx, lines...)` | Ends-with check | Suffix matching |
| `ShouldBeNotEqual(t, idx, lines...)` | Not equal | Negative assertions |
| `ShouldBeRegex(t, idx, lines...)` | Regex match per line | Pattern matching |
| `ShouldBeTrimRegex(t, idx, lines...)` | Trim + regex | Pattern with whitespace |
| `ShouldHaveNoError(t, title, idx, err)` | Error is nil | Error absence check |
| `AssertDirectly(t, title, msg, idx, actual, assertion, expected)` | Any convey assertion | Custom assertions |

---

## Type Verification

Optional but recommended for catching type drift:

```go
import "reflect"

VerifyTypeOf: &coretests.VerifyTypeOf{
    ArrangeInput:  reflect.TypeOf(args.Map{}),
    ActualInput:   reflect.TypeOf([]string{}),
    ExpectedInput: reflect.TypeOf([]string{}),
},
```

If set, CaseV1 automatically validates that ArrangeInput, ActualInput, and ExpectedInput match the declared types.

---

## Error Testing Patterns

### Expecting an error in output lines

```go
{
    Title: "nil function returns error",
    ArrangeInput: args.ThreeFunc{
        WorkFunc: nil,
    },
    ExpectedInput: []string{
        "error : ",
        "  func-wrap is invalid:",
        "      given type: <nil>",
    },
    HasError: true,
},
```

### Asserting error is nil

```go
testCase.ShouldHaveNoError(t, "deserialization", caseIndex, err)
```

---

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Inline test data in `_test.go` | Separate `_testcases.go` file |
| Skip `// Arrange` / `// Act` / `// Assert` comments | Always label AAA sections |
| Omit `caseIndex` from assertions | Always pass `caseIndex` |
| Use `t.Fatal()` directly | Use `ShouldBeEqual` or `GetAssert` |
| Name test dir same as package | Suffix with `tests` (e.g., `chmodhelpertests`) |
| Put test wrapper in `_test.go` | Separate `TestWrapper.go` file |
| Hardcode expected strings inline | Define in `_testcases.go` var block |

---

## Checklist for New Tests

When adding tests for a new package:

- [ ] Create directory: `tests/integratedtests/{package}tests/`
- [ ] Create test cases file: `{Feature}_testcases.go` with `var {feature}TestCases = []coretestcases.CaseV1{...}`
- [ ] Create test file: `{Feature}_test.go` with `func Test_{Feature}_Verification(t *testing.T)`
- [ ] Follow AAA pattern with explicit comments
- [ ] Pass `caseIndex` to all assertion calls
- [ ] Use `coretests.GetAssert.ToStrings()` to convert results to comparable lines
- [ ] Add `VerifyTypeOf` if the test has distinct input/output types
- [ ] Create custom `TestWrapper` if ArrangeInput needs type-safe casting
- [ ] Run `make run-tests` to verify

---

## Related Docs

- [Testing Patterns Overview](/spec/01-app/13-testing-patterns.md)
- [coretests Folder Spec](/spec/01-app/folders/07-coretests.md)
- [Missing Unit Tests Issue](/spec/13-app-issues/testing/01-missing-unit-tests.md)
