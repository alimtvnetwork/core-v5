# Test Structures Reference — Complete API Guide

> **Purpose**: Detailed API reference for every test structure, argument type, and result type used in the testing framework. Use this as a lookup when writing test cases.

## Table of Contents

- [CaseV1 — Primary Test Case](#casev1--primary-test-case)
- [CaseNilSafe — Nil Receiver Testing](#casenilsafe--nil-receiver-testing)
- [GenericGherkins — Typed Test Case](#genericgherkins--typed-test-case)
- [args.Map — Typed Input Map](#argsmap--typed-input-map)
- [args.One through args.Six — Positional Arguments](#argsone-through-argssix--positional-arguments)
- [args.Dynamic — Variable-Length Arguments](#argsdynamic--variable-length-arguments)
- [args.Holder — Multi-Parameter Holder](#argsholder--multi-parameter-holder)
- [args.FuncWrap — Function Reflection](#argsfuncwrap--function-reflection)
- [results.Result — Invocation Result](#resultsresult--invocation-result)
- [results.ResultAny — Untyped Result Alias](#resultsresultany--untyped-result-alias)
- [results.ExpectAnyError — Error Sentinel](#resultsexpectanyerror--error-sentinel)
- [errcore.AssertDiffOnMismatch — Diff Assertion](#errcoreassertdiffonmismatch--diff-assertion)
- [Custom Test Wrappers](#custom-test-wrappers)
- [When to Use What](#when-to-use-what)
- [Related Docs](#related-docs)

---

## CaseV1 — Primary Test Case

**Package**: `gitlab.com/auk-go/core/coretests/coretestcases`

**Definition**: `type CaseV1 coretests.BaseTestCase`

### Fields

| Field | Type | Purpose | Required |
|-------|------|---------|----------|
| `Title` | `string` | Test case header / scenario name | Yes |
| `ArrangeInput` | `any` | Input data (usually `args.Map`) | Yes (for parameterized tests) |
| `ActualInput` | `any` | Set after Act phase via `SetActual` | Auto |
| `ExpectedInput` | `any` | Expected result (`string`, `[]string`, or `args.Map`) | Yes |
| `Additional` | `any` | Extra context data | Optional |
| `CustomFormat` | `string` | Custom output format | Optional |
| `VerifyTypeOf` | `*VerifyTypeOf` | Type verification config | Optional |
| `Parameters` | `*args.Holder` | Extra parameters | Optional |
| `IsEnable` | `issetter.Value` | Set to False to disable | Optional |
| `HasError` | `bool` | Whether error is expected | Optional |
| `HasPanic` | `bool` | Whether panic is expected | Optional |
| `IsValidateError` | `bool` | Whether to validate error content | Optional |

### Key Methods

```go
// Assertion methods (GoConvey-based)
tc.ShouldBeEqual(t, caseIndex, actuals...)           // exact match
tc.ShouldBeEqualFirst(t, actuals...)                  // same, caseIndex=0
tc.ShouldBeEqualMap(t, caseIndex, actualMap)          // sorted map comparison
tc.ShouldBeEqualMapFirst(t, actualMap)                // same, caseIndex=0
tc.ShouldBeTrimEqual(t, caseIndex, actuals...)        // trim whitespace
tc.ShouldBeSortedEqual(t, caseIndex, actuals...)      // sort + trim
tc.ShouldContains(t, caseIndex, actuals...)           // substring match
tc.ShouldStartsWith(t, caseIndex, actuals...)         // prefix match
tc.ShouldEndsWith(t, caseIndex, actuals...)           // suffix match
tc.ShouldBeNotEqual(t, caseIndex, actuals...)         // not equal
tc.ShouldBeRegex(t, caseIndex, actuals...)            // regex match
tc.ShouldHaveNoError(t, title, caseIndex, err)        // nil error check

// Data access
tc.Input() any                                         // ArrangeInput
tc.Expected() any                                      // ExpectedInput
tc.ExpectedLines() []string                            // normalized to []string
tc.ExpectedAsMap() args.Map                            // type-asserted to args.Map
tc.CaseTitle() string                                  // Title
```

### Usage Pattern

```go
// _testcases.go
var myTestCase = coretestcases.CaseV1{
    Title: "Positive: valid input returns expected output",
    ArrangeInput: args.Map{
        "input":  "hello",
        "option": true,
    },
    ExpectedInput: args.Map{
        "result":  "HELLO",
        "isValid": true,
    },
}

// _test.go
func Test_MyFunc_ValidInput(t *testing.T) {
    tc := myTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    opt, _ := input.GetAsBool("option")

    // Act
    result := MyFunc(str, opt)
    actual := args.Map{
        "result":  result.Value,
        "isValid": result.IsValid,
    }

    // Assert
    tc.ShouldBeEqualMapFirst(t, actual)
}
```

---

## CaseNilSafe — Nil Receiver Testing

**Package**: `gitlab.com/auk-go/core/coretests/coretestcases`

### Fields

| Field | Type | Purpose | Required |
|-------|------|---------|----------|
| `Title` | `string` | Scenario name | Optional (falls back to method name) |
| `Func` | `any` | Method expression: `(*Type).Method` | Yes |
| `Args` | `[]any` | Extra arguments after receiver | Optional |
| `Expected` | `results.ResultAny` | Expected invocation result | Yes |
| `CompareFields` | `[]string` | Override auto-derived field comparison | Optional |

### Key Methods

```go
tc.MethodName() string                    // reflected method name
tc.CaseTitle() string                     // Title or MethodName fallback
tc.Invoke(receiver) results.ResultAny     // invoke with specific receiver
tc.InvokeNil() results.ResultAny          // invoke with nil receiver
tc.ShouldBeSafe(t, caseIndex)             // invoke nil + assert (primary usage)
tc.ShouldBeSafeFirst(t)                   // same, caseIndex=0
```

### Func Field Patterns

```go
// 1. Non-generic type — direct method expression
Func: (*MyStruct).IsValid,

// 2. Generic type — function literal wrapper (compile-time safe)
Func: func(c *Collection[string]) bool {
    return c.IsEmpty()
},

// 3. Method returning interface/pointer — check nil via wrapper
Func: func(w *MyStruct) bool {
    return w.GetThing() == nil
},

// 4. Package function (not a method) — struct{} receiver wrapper
Func: func(_ *struct{}) bool {
    return mypackage.SomeFunc(nil) == nil
},

// 5. Method with arguments
Func: (*MyStruct).Process,
Args: []any{"input", 42},
```

### Expected Field Patterns

```go
// Boolean return
Expected: results.ResultAny{Value: "true", Panicked: false}
Expected: results.ResultAny{Value: "false", Panicked: false}

// String return
Expected: results.ResultAny{Value: "hello", Panicked: false}
Expected: results.ResultAny{Value: "", Panicked: false}

// Int return
Expected: results.ResultAny{Value: "0", Panicked: false}
Expected: results.ResultAny{Value: "42", Panicked: false}

// Error return
Expected: results.ResultAny{Panicked: false, Error: results.ExpectAnyError}

// Panic expected
Expected: results.ResultAny{Panicked: true}

// Void method
Expected: results.ResultAny{Panicked: false}
CompareFields: []string{"panicked", "returnCount"}

// Multi-return with count
Expected: results.ResultAny{
    Value:       "0",
    Panicked:    false,
    Error:       results.ExpectAnyError,
    ReturnCount: 2,
}
```

### Complete File Example

```go
// Feature_NilReceiver_testcases.go
package mypkgtests

import (
    "gitlab.com/auk-go/core/coretests/coretestcases"
    "gitlab.com/auk-go/core/coretests/results"
    "gitlab.com/auk-go/core/mypkg"
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
        Title: "Length on nil returns 0",
        Func:  (*mypkg.MyStruct).Length,
        Expected: results.ResultAny{
            Value:    "0",
            Panicked: false,
        },
    },
}

// NilReceiver_test.go
package mypkgtests

import "testing"

func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

---

## GenericGherkins — Typed Test Case

**Package**: `gitlab.com/auk-go/core/coretests/coretestcases`

### Definition

```go
type GenericGherkins[TInput, TExpect any] struct {
    Title      string
    Feature    string
    Given      string
    When       string
    Then       string
    Input      TInput
    Expected   TExpect
    Actual     TExpect
    IsMatching bool
    ExtraArgs  args.Map
}
```

### Aliases

```go
type AnyGherkins         = GenericGherkins[any, any]
type StringGherkins      = GenericGherkins[string, string]
type StringBoolGherkins  = GenericGherkins[string, bool]
```

### Key Methods

```go
tc.CaseTitle() string
tc.TypedInput() TInput
tc.TypedExpected() TExpect
tc.TypedActual() TExpect
tc.SetTypedActual(actual TExpect)
tc.IsFailedToMatch() bool
tc.HasExtraArgs() bool
tc.GetExtra(key string) any
tc.GetExtraAsBool(key string) bool

// Assertions
tc.ShouldBeEqual(t, caseIndex, actLines, expectedLines)
tc.ShouldBeEqualFirst(t, actLines, expectedLines)
tc.ShouldBeEqualArgs(t, caseIndex, actLines...)
tc.ShouldBeEqualArgsFirst(t, actLines...)
tc.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
tc.ShouldBeEqualUsingExpectedFirst(t, actLines)
tc.ShouldMatchExpected(t, caseIndex, result)
tc.ShouldMatchExpectedFirst(t, result)
```

### When to Use

- New typed tests where `TInput` and `TExpect` are known at compile time
- Tests requiring Gherkin-style descriptions (Feature, Given, When, Then)
- When you want `TypedTestCaseWrapper` interface satisfaction

### Usage Example

```go
// _testcases.go
var regexMatchTestCase = coretestcases.StringBoolGherkins{
    Title:      "Valid pattern matches input",
    When:       "given a digit pattern",
    Input:      `\d+`,
    Expected:   true,
    IsMatching: true,
    ExtraArgs: args.Map{
        "testInput": "42",
    },
}

// _test.go
func Test_Regex_Match(t *testing.T) {
    tc := regexMatchTestCase

    // Arrange
    pattern := tc.TypedInput()
    testInput := tc.ExtraArgs["testInput"].(string)

    // Act
    regex := regexnew.New.Lazy(pattern)
    result := regex.IsMatch(testInput)
    tc.SetTypedActual(result)

    // Assert
    tc.ShouldMatchExpectedFirst(t, result)
}
```

---

## args.Map — Typed Input Map

**Package**: `gitlab.com/auk-go/core/coretests/args`

**Definition**: `type Map map[string]any`

### Typed Accessors

| Method | Returns | Purpose |
|--------|---------|---------|
| `GetAsString(key)` | `(string, bool)` | Get string value |
| `GetAsStringDefault(key)` | `string` | Get string, empty if missing |
| `GetAsInt(key)` | `(int, bool)` | Get int value |
| `GetAsIntDefault(key, def)` | `int` | Get int with default |
| `GetAsBool(key)` | `(bool, bool)` | Get bool value |
| `GetAsBoolDefault(key)` | `bool` | Get bool, false if missing |
| `GetAsStrings(key)` | `([]string, bool)` | Get string slice |
| `GetDirectLower(key)` | `any` | Raw access, nil if missing |
| `HasDefined(key)` | `bool` | Key exists check |
| `IsKeyMissing(key)` | `bool` | Key absent check |

### Special Key Accessors

| Method | Looks Up Keys |
|--------|--------------|
| `When()` | `"when"` |
| `Actual()` | `"actual"` |
| `Expect()` | `"expect"` |
| `WorkFunc()` | `"func"`, `"work.func"` |
| `FirstItem()` | `"first"`, `"f1"`, `"p1"` |
| `SecondItem()` | `"second"`, `"f2"` |

### Assertion Methods

| Method | Purpose |
|--------|---------|
| `CompileToStrings()` | Sorted `"key : value"` lines for comparison |

### Usage in ArrangeInput

```go
// Simple input
ArrangeInput: args.Map{
    "input": "hello",
},

// Multiple inputs
ArrangeInput: args.Map{
    "firstName": "Alice",
    "lastName":  "Smith",
    "age":       30,
},

// With function reference
ArrangeInput: args.Map{
    "func":  myTargetFunction,
    "p1":    "arg1",
    "p2":    42,
},
```

### Usage in ExpectedInput

```go
// Multi-property assertion (preferred for 2+ values)
ExpectedInput: args.Map{
    "isValid": true,
    "length":  3,
    "name":    "hello",
},

// In test body — build actual map with SAME keys
actual := args.Map{
    "isValid": result.IsValid(),
    "length":  result.Length(),
    "name":    result.Name(),
}
tc.ShouldBeEqualMap(t, caseIndex, actual)
```

### Error Handling — Always Validate

```go
// ✅ CORRECT — validate getter returns
str, valid := input.GetAsString("key")
errcore.HandleErrMessage("key", !valid)

// ❌ BANNED — silently ignores missing key
str, _ := input.GetAsString("key")
```

---

## args.One through args.Six — Positional Arguments

**Package**: `gitlab.com/auk-go/core/coretests/args`

### Overview

Typed positional argument holders for up to 6 values. **Used in `ArrangeInput` only** — never in `ExpectedInput`.

### Types

```go
type One[T any]                     struct { First T }
type Two[T1, T2 any]               struct { First T1; Second T2 }
type Three[T1, T2, T3 any]         struct { First T1; Second T2; Third T3 }
type Four[T1, T2, T3, T4 any]     struct { First T1; Second T2; Third T3; Fourth T4 }
type Five[T1, T2, T3, T4, T5 any] struct { ... }
type Six[T1, T2, T3, T4, T5, T6 any] struct { ... }
```

### Aliases (any-based, backward compatible)

```go
type OneAny     = One[any]
type TwoAny     = Two[any, any]
type ThreeAny   = Three[any, any, any]
type FourAny    = Four[any, any, any, any]
// etc.
```

### Usage in ArrangeInput

```go
// Typed positional input
ArrangeInput: args.Two[string, int]{
    First:  "hello",
    Second: 42,
},

// In test body
input := tc.ArrangeInput.(args.Two[string, int])
str := input.First
num := input.Second
```

### When to Use

- When inputs are simple positional values with obvious meaning
- When you need compile-time type safety for inputs
- **Never for ExpectedInput** — use `args.Map` or `string`/`[]string` instead

### Functional Variants

```go
type OneFunc[T any]   struct { First T; Func func(T) any }
type TwoFunc[T1, T2 any] struct { ... }
// etc.

type OneFuncAny   = OneFunc[any]
type TwoFuncAny   = TwoFunc[any, any]
// etc.
```

These embed a function alongside the arguments, primarily used for function-wrapping test patterns.

---

## args.Dynamic — Variable-Length Arguments

**Package**: `gitlab.com/auk-go/core/coretests/args`

```go
type Dynamic[T any] struct {
    Items []T
}

type DynamicAny = Dynamic[any]
```

### Usage

For test cases that need variable-length inputs:

```go
ArrangeInput: args.Dynamic[string]{
    Items: []string{"a", "b", "c", "d"},
},
```

---

## args.Holder — Multi-Parameter Holder

**Package**: `gitlab.com/auk-go/core/coretests/args`

Used as `Parameters` field on `BaseTestCase` for complex test configurations:

```go
var tc = coretestcases.CaseV1{
    Title: "Complex test",
    Parameters: &args.Holder{
        Items: args.Map{
            "config":  myConfig,
            "timeout": 5,
        },
    },
}
```

---

## args.FuncWrap — Function Reflection

**Package**: `gitlab.com/auk-go/core/coretests/args`

```go
type FuncWrap[T any] struct {
    Func T
    Name string
}

type FuncWrapAny = FuncWrap[any]
```

Wraps a function with its name for reflection-based invocation in test frameworks.

---

## results.Result — Invocation Result

**Package**: `gitlab.com/auk-go/core/coretests/results`

### Fields

| Field | Type | Purpose |
|-------|------|---------|
| `Value` | `T` | Primary return value |
| `Error` | `error` | Error from function or nil |
| `Panicked` | `bool` | True if recovered from panic |
| `PanicValue` | `any` | Raw `recover()` value |
| `AllResults` | `[]any` | All return values (multi-return) |
| `ReturnCount` | `int` | Number of return values |

### Methods

```go
result.IsSafe() bool               // !Panicked && Error == nil
result.HasError() bool             // Error != nil
result.HasPanicked() bool          // Panicked
result.IsResult(expected) bool     // %v equality
result.IsResultTypeOf(exp) bool    // reflect.Type assignability
result.IsError(msg) bool           // error message match
result.ValueString() string        // fmt.Sprintf("%v", Value)
result.ResultAt(index) any         // AllResults[index]
result.ToMap() args.Map            // full map for assertion
result.ToMapCompact() args.Map     // minimal map (value + panicked only)
result.String() string             // human-readable summary

// Primary assertion
result.ShouldMatchResult(t, caseIndex, title, expected, compareFields...)
```

### ToMap Output

```go
result.ToMap() → args.Map{
    "value":       fmt.Sprintf("%v", result.Value),
    "panicked":    result.Panicked,
    "isSafe":      result.IsSafe(),
    "hasError":    result.HasError(),
    "returnCount": result.ReturnCount,
}
```

---

## results.ResultAny — Untyped Result Alias

```go
type ResultAny = Result[any]
```

Used as the `Expected` field in `CaseNilSafe`. Other aliases:

```go
type ResultBool   = Result[bool]
type ResultString = Result[string]
type ResultInt    = Result[int]
type ResultError  = Result[error]
```

### Field Auto-Derivation in ShouldMatchResult

When `CompareFields` is empty, compared fields are auto-derived:

| Condition | Field Compared |
|-----------|---------------|
| Always | `"panicked"` |
| `Expected.Value != nil` | `"value"` |
| `Expected.Error != nil` | `"hasError"` |
| `Expected.ReturnCount != 0` | `"returnCount"` |

### Override with CompareFields

```go
CaseNilSafe{
    Func: (*MyStruct).Reset,
    Expected: results.ResultAny{
        Panicked: false,
    },
    CompareFields: []string{"panicked", "returnCount"},
}
```

---

## results.ExpectAnyError — Error Sentinel

```go
var ExpectAnyError = fmt.Errorf("expect-any-error")
```

Use in `Expected.Error` to assert that the method returns **any** non-nil error (without checking the specific message):

```go
Expected: results.ResultAny{
    Panicked: false,
    Error:    results.ExpectAnyError,
}
```

This causes `ShouldMatchResult` to compare `hasError: true` in the filtered map.

---

## errcore.AssertDiffOnMismatch — Diff Assertion

**Package**: `gitlab.com/auk-go/core/errcore`

### Signature

```go
func AssertDiffOnMismatch(
    t *testing.T,
    caseIndex int,
    title string,
    actLines []string,
    expectedLines any,        // normalizes string → []string
    contextLines ...string,   // optional diagnostics on failure
)
```

### When to Use

- Custom test wrappers that don't use CaseV1
- When you need context lines in failure output
- When not using GoConvey assertion chains

### Usage

```go
// Basic
errcore.AssertDiffOnMismatch(
    t,
    caseIndex,
    tc.Title,
    actLines,
    expectedLines,
)

// With context
errcore.AssertDiffOnMismatch(
    t,
    caseIndex,
    tc.Title,
    actLines,
    expectedLines,
    fmt.Sprintf("  InitValue: %d", tc.InitValue),
    fmt.Sprintf("  Config: %s", tc.Config),
)
```

### Error-Specific Variant

```go
errcore.AssertErrorDiffOnMismatch(
    t,
    caseIndex,
    tc.Title,
    err,
    expectedLines,
)
```

---

## Custom Test Wrappers

### When to Create

- The test needs extra fields beyond `CaseV1` (e.g., `InitBytes`, `CompareValue`)
- Complex arrange logic requires structured data

### Pattern

```go
// _testcases.go
type myCustomTestCase struct {
    Case       coretestcases.CaseV1
    InitValue  int
    ExtraConfig string
}

var myCustomTest = myCustomTestCase{
    Case: coretestcases.CaseV1{
        Title:         "Custom test with extra config",
        ExpectedInput: []string{"42", "true"},
    },
    InitValue:   42,
    ExtraConfig: "verbose",
}

// _test.go
func Test_Custom(t *testing.T) {
    tc := myCustomTest

    // Arrange
    obj := NewThing(tc.InitValue, tc.ExtraConfig)

    // Act
    result := obj.Process()
    actLines := []string{
        fmt.Sprintf("%d", result.Value),
        fmt.Sprintf("%v", result.IsValid),
    }

    // Assert
    expectedLines := tc.Case.ExpectedInput.([]string)
    errcore.AssertDiffOnMismatch(
        t,
        0,
        tc.Case.Title,
        actLines,
        expectedLines,
        fmt.Sprintf("  InitValue: %d", tc.InitValue),
    )
}
```

---

## When to Use What

### Quick Decision Guide

| Scenario | Structure | Assertion |
|----------|-----------|-----------|
| Standard functional test | `CaseV1` | `ShouldBeEqual` |
| Multi-property assertion | `CaseV1` + `args.Map` | `ShouldBeEqualMap` |
| Nil receiver test | `CaseNilSafe` | `ShouldBeSafe` |
| Typed input/expected | `GenericGherkins[T1,T2]` | `ShouldMatchExpected` |
| Custom wrapper needed | Custom struct + `CaseV1` | `errcore.AssertDiffOnMismatch` |
| Order-independent | `CaseV1` | `ShouldBeSortedEqual` |
| Substring match | `CaseV1` | `ShouldContains` |
| Error existence | `CaseV1` + `args.Map` | `ShouldBeEqualMap` |
| Panic test | `CaseV1` + helper | `ShouldBeEqual` |
| Nil panic test | `CaseNilSafe` | `ShouldBeSafe` (Panicked: true) |

### ExpectedInput Format Decision

```
How many values to assert?
├── 1 value → string
├── 2+ values, ordered/positional → []string
├── 2+ values, semantic keys → args.Map (preferred)
└── Variable-length collection → []string
```

### ArrangeInput Format Decision

```
What kind of input?
├── Named parameters → args.Map (preferred)
├── Typed positional → args.One through args.Six
├── Function + args → args.Map with "func" key
└── Variable-length → args.Dynamic[T]
```

---

## Related Docs

- [Testing Guidelines](./16-testing-guidelines.md) — full testing workflow and examples
- [Branch Coverage Strategy](./23-branch-coverage-strategy.md) — positive/negative path coverage
- [Test Quality Guide](./24-test-quality-guide.md) — good vs bad test comparison
- [CaseNilSafe Design](./designs/CaseNilSafe-design.md) — architecture and edge cases
- [Coding Guidelines](./17-coding-guidelines.md) — formatting and naming
