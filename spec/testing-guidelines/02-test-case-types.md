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
type AnyGherkins    = GenericGherkins[any, any]
type StringGherkins = GenericGherkins[string, string]
```

### When to Use

- Complex scenarios with structured Given/When/Then
- Regex/validation tests where `IsMatching` semantics apply
- Tests requiring both typed expected values and string-line assertions

### Example

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
