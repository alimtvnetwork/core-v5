# 02 — Test Case Types

## Overview

| Type | Use When | Input | Expected |
|------|----------|-------|----------|
| `CaseV1` | Testing package functions or methods with explicit Act step | `ArrangeInput` | `ExpectedInput` |
| `CaseNilSafe` | Testing nil-receiver safety of pointer receiver methods | `Func` (method ref) | `Expected` (ResultAny) |
| `GenericGherkins[TInput, TExpect]` | BDD-style scenarios with typed input/output | `Input` | `Expected` / `ExtraArgs` |

---

## CaseV1

The primary workhorse. Use for any test where you explicitly control Arrange → Act → Assert.

### Structure

```go
type CaseV1 struct {
    Title         string         // Test case name / scenario description
    ArrangeInput  any            // Input data (args.Map, args.One, etc.)
    ActualInput   any            // Set dynamically after Act phase
    ExpectedInput any            // Expected output (args.Map, string, []string, bool, etc.)
    VerifyTypeOf  *VerifyTypeOf  // Optional: type verification
    Parameters    *args.HolderAny // Optional: extra parameters
}
```

### Key Fields

| Field | Type | Purpose |
|-------|------|---------|
| `Title` | `string` | Displayed in test output on failure |
| `ArrangeInput` | `any` | Holds input data. Usually `args.Map` or positional types |
| `ExpectedInput` | `any` | Holds expected output. Must be `args.Map` for map assertions, `string`/`[]string` for line assertions |

### ExpectedInput Auto-Normalization

`ExpectedInput` is normalized to `[]string` via `ExpectedLines()`. Supported types:

| Type | Conversion |
|------|-----------|
| `string` | `[]string{s}` |
| `[]string` | as-is |
| `int` | `[]string{strconv.Itoa(v)}` |
| `bool` | `[]string{"true"}` or `[]string{"false"}` |
| `args.Map` | sorted `"key : value"` lines |
| other | `PrettyJSON` fallback |

### Assertion Methods

| Method | Use When |
|--------|----------|
| `ShouldBeEqual(t, caseIndex, actual...)` | Loop-based, exact string match |
| `ShouldBeEqualFirst(t, actual...)` | Single test case (caseIndex=0) |
| `ShouldBeEqualMap(t, caseIndex, actualMap)` | Map-based comparison |
| `ShouldBeEqualMapFirst(t, actualMap)` | Single test case, map comparison |
| `ShouldContains(t, caseIndex, actual...)` | Substring match |
| `ShouldStartsWith(t, caseIndex, actual...)` | Prefix match |
| `ShouldEndsWith(t, caseIndex, actual...)` | Suffix match |
| `ShouldBeNotEqual(t, caseIndex, actual...)` | Inverse match |
| `ShouldBeTrimEqual(t, caseIndex, actual...)` | Trimmed comparison |
| `ShouldBeSortedEqual(t, caseIndex, actual...)` | Sorted + trimmed comparison |
| `ShouldBeRegex(t, caseIndex, actual...)` | Regex match |

### Complete Example

**`_testcases.go`:**
```go
package mypkgtests

import (
    "myproject/coretests/args"
    "myproject/coretests/coretestcases"
)

// =============================================================================
// MyFunc positive path
// =============================================================================

var myFuncPositiveTestCases = []coretestcases.CaseV1{
    {
        Title: "MyFunc returns sum of two integers",
        ArrangeInput: args.Map{
            "a": 3,
            "b": 5,
        },
        ExpectedInput: args.Map{
            "result":  8,
            "isValid": true,
        },
    },
    {
        Title: "MyFunc handles zero values",
        ArrangeInput: args.Map{
            "a": 0,
            "b": 0,
        },
        ExpectedInput: args.Map{
            "result":  0,
            "isValid": true,
        },
    },
}

// =============================================================================
// MyFunc negative path — nil input
// =============================================================================

var myFuncNilInputTestCase = coretestcases.CaseV1{
    Title: "MyFunc with nil returns error",
    ArrangeInput: args.Map{
        "input": nil,
    },
    ExpectedInput: args.Map{
        "hasError": true,
        "result":   0,
    },
}
```

**`_test.go`:**
```go
package mypkgtests

import (
    "testing"

    "myproject/coretests/args"
    "myproject/mypkg"
)

// ==========================================
// MyFunc — positive path
// ==========================================

func Test_MyFunc_Positive_Verification(t *testing.T) {
    for caseIndex, tc := range myFuncPositiveTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        a, _ := input.GetAsInt("a")
        b, _ := input.GetAsInt("b")

        // Act
        result, err := mypkg.MyFunc(a, b)
        actual := args.Map{
            "result":  result,
            "isValid": err == nil,
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}

// ==========================================
// MyFunc — negative path (nil input)
// ==========================================

func Test_MyFunc_NilInput(t *testing.T) {
    tc := myFuncNilInputTestCase

    // Arrange
    // (nil input — no setup needed)

    // Act
    result, err := mypkg.MyFunc(0, 0)
    actual := args.Map{
        "hasError": err != nil,
        "result":   result,
    }

    // Assert
    tc.ShouldBeEqualMapFirst(t, actual)
}
```

---

## CaseNilSafe

Designed exclusively for testing nil-receiver safety of **pointer receiver methods**.

### Structure

```go
type CaseNilSafe struct {
    Title         string          // Scenario name
    Func          any             // Direct method reference: (*Type).Method
    Args          []any           // Optional arguments for the method call
    Expected      results.ResultAny  // Expected outcome
    CompareFields []string        // Override auto-derived field comparison
}
```

### When to Use

✅ Use for: pointer receiver methods that must not panic on nil  
❌ Do NOT use for: package-level functions (use CaseV1 instead)

### How Func Works

The `Func` field accepts a **method expression** — a direct reference to a method:

```go
// Zero-arg method — use method expression directly:
Func: (*MyStruct).IsValid

// Method with arguments — wrap in a function literal:
Func: func(m *MyStruct) bool {
    return m.HasKey("someKey")
}

// Void method — wrap to suppress no-return:
Func: func(m *MyStruct) {
    m.SetName("x")
}
```

### Expected Fields (auto-derived)

| Field | Auto-compared when... | Meaning |
|-------|----------------------|---------|
| `Panicked` | always | Whether a panic occurred |
| `Value` | `Expected.Value != nil` | The stringified return value |
| `Error` | `Expected.Error != nil` | Whether an error was returned |
| `ReturnCount` | `Expected.ReturnCount != 0` | Number of return values |

### CompareFields Override

When auto-derivation isn't sufficient, explicitly specify which fields to compare:

```go
{
    Title: "SetName on nil does not panic",
    Func: func(m *MyStruct) {
        m.SetName("x")
    },
    Expected: results.ResultAny{
        Panicked: false,
    },
    // Void method has no "value" — only compare panicked + returnCount
    CompareFields: []string{"panicked", "returnCount"},
}
```

### Complete Example

**`_NilReceiver_testcases.go`:**
```go
package mypkgtests

import (
    "myproject/coretests/coretestcases"
    "myproject/coretests/results"
    "myproject/mypkg"
)

var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*mypkg.MyStruct).IsValid,
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "Name on nil returns empty",
        Func:  (*mypkg.MyStruct).Name,
        Expected: results.ResultAny{
            Value:    "",
            Panicked: false,
        },
    },
    {
        Title: "HasKey on nil returns false",
        Func: func(m *mypkg.MyStruct) bool {
            return m.HasKey("anything")
        },
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "ClonePtr on nil returns nil",
        Func: func(m *mypkg.MyStruct) bool {
            return m.ClonePtr() == nil
        },
        Expected: results.ResultAny{
            Value:    "true",
            Panicked: false,
        },
    },
    {
        Title: "Clear on nil does not panic",
        Func:  (*mypkg.MyStruct).Clear,
        Expected: results.ResultAny{
            Panicked: false,
        },
        CompareFields: []string{"panicked", "returnCount"},
    },
}
```

**`NilReceiver_test.go`:**
```go
package mypkgtests

import "testing"

func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

### Pattern Abuse Warning

**Never** use `CaseNilSafe` for package-level functions. If `ConcatMessageWithErr` is `func(string, error) error` (not a method), use `CaseV1`:

```go
// ❌ BAD — pattern abuse
var badTestCase = coretestcases.CaseNilSafe{
    Func: func(_ *struct{}) bool {
        return errcore.ConcatMessageWithErr("msg", nil) == nil
    },
}

// ✅ GOOD — use CaseV1
var goodTestCase = coretestcases.CaseV1{
    Title: "ConcatMessageWithErr nil error returns nil",
    ArrangeInput: args.Map{"message": "should not appear"},
    ExpectedInput: args.Map{"isNil": true},
}
```

---

## GenericGherkins[TInput, TExpect]

BDD-style test case with typed fields for input and expectations.

### Structure

```go
type GenericGherkins[TInput, TExpect any] struct {
    Title         string
    Feature       string
    Given         string
    When          string
    Then          string
    Input         TInput
    Expected      TExpect
    Actual        TExpect      // Set after Act
    IsMatching    bool
    ExpectedLines []string
    ExtraArgs     args.Map     // Overflow key-value pairs
}
```

### Common Aliases

```go
type AnyGherkins      = GenericGherkins[any, any]
type StringGherkins   = GenericGherkins[string, string]
type StringBoolGherkins = GenericGherkins[string, bool]
type MapGherkins      = GenericGherkins[args.Map, args.Map]
```

### When to Use Which Alias

| Alias | Input | Expected | Use When |
|-------|-------|----------|----------|
| `StringGherkins` | `string` | `string` | Single string input → single string result |
| `StringBoolGherkins` | `string` | `bool` | String input → boolean result (e.g., IsMatch) |
| `MapGherkins` | `args.Map` | `args.Map` | Multi-field input → multi-field result |
| `AnyGherkins` | `any` | `any` | Heterogeneous or unknown types |

### Field Responsibility

| Field | Purpose | Use When |
|-------|---------|----------|
| `Input` | Primary typed input data | Always — holds the main test input |
| `Expected` | Primary typed expected result | Always — holds what the test asserts against |
| `ExtraArgs` | Overflow key-value pairs | Only when `Input` cannot hold all arrange data |
| `ExpectedLines` | Legacy raw string assertions | **Deprecated for new tests** — use `Expected` with `args.Map` instead |
| `IsMatching` | Boolean match flag | Validation/matching tests (e.g., regex, search) |

### MapGherkins — Preferred for Multi-Field Tests

When a test has multiple inputs (e.g., pattern + compareInput) and multiple expected
results (e.g., isDefined, isApplicable, isMatch), use `MapGherkins`:

- **Input** (`args.Map`): holds all arrange data with semantic keys
- **Expected** (`args.Map`): holds all assertion data with semantic keys
- **ExtraArgs**: only used if needed beyond `Input`

This replaces the opaque `ExpectedLines: []string{"true", "false", "true"}` pattern
where each line's meaning is unknowable without reading the test runner.

#### Example — MapGherkins Test Case

```go
// params.go
package regexnewtests

var params = struct {
    pattern      string
    compareInput string
    isDefined    string
    isApplicable string
    isMatch      string
    isFailedMatch string
}{
    pattern:      "pattern",
    compareInput: "compareInput",
    isDefined:    "isDefined",
    isApplicable: "isApplicable",
    isMatch:      "isMatch",
    isFailedMatch: "isFailedMatch",
}
```

```go
// _testcases.go
var lazyRegexTestCases = []coretestcases.MapGherkins{
    {
        Title: "New.Lazy matches word pattern",
        When:  "given a simple word pattern",
        Input: args.Map{
            params.pattern:      "hello",
            params.compareInput: "hello world",
        },
        Expected: args.Map{
            params.isDefined:    true,
            params.isApplicable: true,
            params.isMatch:      true,
            params.isFailedMatch: false,
        },
    },
}
```

```go
// _test.go
func Test_LazyRegex_New_Verification(t *testing.T) {
    for caseIndex, tc := range lazyRegexTestCases {
        // Arrange
        pattern, _ := tc.Input.GetAsString(params.pattern)
        compareInput, _ := tc.Input.GetAsString(params.compareInput)

        // Act
        lazyRegex := regexnew.New.Lazy(pattern)
        actual := args.Map{
            params.isDefined:    lazyRegex.IsDefined(),
            params.isApplicable: lazyRegex.IsApplicable(),
            params.isMatch:      lazyRegex.IsMatch(compareInput),
            params.isFailedMatch: lazyRegex.IsFailedMatch(compareInput),
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

#### Why MapGherkins Over ExpectedLines

❌ **Bad — opaque ExpectedLines:**
```go
{
    Title:      "New.Lazy with simple word pattern",
    Input:      "hello",
    ExtraArgs:  map[string]any{"compareInput": "hello world"},
    ExpectedLines: []string{
        "hello",   // what is this?
        "true",    // isDefined? isApplicable? isMatch?
        "true",    // ???
        "true",    // ???
        "false",   // ???
    },
}
```

✅ **Good — self-documenting MapGherkins:**
```go
{
    Title: "New.Lazy matches word pattern",
    Input: args.Map{
        "pattern":      "hello",
        "compareInput": "hello world",
    },
    Expected: args.Map{
        "isDefined":    true,
        "isApplicable": true,
        "isMatch":      true,
        "isFailedMatch": false,
    },
}
```

### Assertion Methods — GenericGherkins

| Method | Use When |
|--------|----------|
| `ShouldBeEqual(t, caseIndex, actLines, expLines)` | Raw string line comparison |
| `ShouldBeEqualFirst(t, actLines, expLines)` | Single case, string lines |
| `ShouldBeEqualArgs(t, caseIndex, lines...)` | Variadic string args vs ExpectedLines |
| `ShouldBeEqualArgsFirst(t, lines...)` | Single case, variadic args |
| `ShouldBeEqualUsingExpected(t, caseIndex, actLines)` | Act lines vs struct's ExpectedLines |
| `ShouldBeEqualUsingExpectedFirst(t, actLines)` | Single case, vs ExpectedLines |
| `ShouldBeEqualMap(t, caseIndex, actual)` | Map-based comparison (MapGherkins) |
| `ShouldBeEqualMapFirst(t, actual)` | Single case, map comparison |

### Legacy Example — StringGherkins (still valid for simple cases)

```go
var regexMatchTestCases = []coretestcases.StringGherkins{
    {
        Title:      "Email pattern matches valid email",
        Given:      "a compiled email regex",
        When:       "matching against user@example.com",
        Input:      "user@example.com",
        IsMatching: true,
        ExtraArgs: args.Map{
            "pattern": `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
        },
    },
}
```
