# Testing Guidelines — Comprehensive Reference

> **This document is the single source of truth** for writing integration tests in the `auk-go/core` repository.
> Any AI agent, contributor, or reviewer must follow these patterns precisely.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [File Organization](#file-organization)
- [Test Case Structures](#test-case-structures)
- [AAA Pattern (Arrange-Act-Assert)](#aaa-pattern-arrange-act-assert)
- [Assertion Methods](#assertion-methods)
- [Input Management (args.Map)](#input-management-argsmap)
- [Named Test Case Variables](#named-test-case-variables)
- [ExpectedInput Formats](#expectedinput-formats)
- [SliceValidator and VerifyAll](#slicevalidator-and-verifyall)
- [GoConvey Integration](#goconvey-integration)
- [errcore.AssertDiffOnMismatch](#errcoreassertdiffonmismatch)
- [Comparison Modes](#comparison-modes)
- [Conditions (Trim, Sort)](#conditions-trim-sort)
- [Type Verification](#type-verification)
- [Error Testing Patterns](#error-testing-patterns)
- [Panic Testing](#panic-testing)
- [Concurrency Testing](#concurrency-testing)
- [Custom Test Wrappers](#custom-test-wrappers)
- [Anti-Patterns (Banned)](#anti-patterns-banned)
- [Complete Examples](#complete-examples)
- [Checklist for New Tests](#checklist-for-new-tests)
- [Future: GenericGherkins Proposal](#future-genericgherkins-proposal)
- [Future: CaseV2 Proposal](#future-casev2-proposal)
- [Related Docs](#related-docs)

---

## Architecture Overview

```
coretests/                          # Test utility package
├── coretestcases/
│   ├── CaseV1.go                   # Primary test case structure
│   ├── GenericGherkins.go          # Generic typed test case (Gherkin pattern)
│   ├── GenericGherkinsAliases.go   # AnyGherkins, StringGherkins, StringBoolGherkins
│   ├── GenericGherkinsAssertions.go # ShouldBeEqual, ShouldBeEqualUsingExpected
│   ├── GenericGherkinsCompare.go   # CompareWith — structural diffing
│   ├── GenericGherkinsFormatting.go # String, ToString, FullString
│   ├── GenericGherkinsGetters.go   # IsFailedToMatch, HasExtraArgs, GetExtra
│   └── GenericGherkinsTypedWrapper.go # TypedTestCaseWrapper implementation
├── args/
│   ├── Map.go                      # Typed input map for test arrangements
│   ├── FuncWrap.go                 # Function reflection wrapper
│   ├── Holder.go                   # Multi-parameter holder
│   ├── OneFunc.go .. SixFunc.go    # Typed function argument structs
│   └── ...
├── BaseTestCase.go                 # Root struct (Title, ArrangeInput, ExpectedInput, etc.)
├── BaseTestCaseAssertions.go       # ShouldBe, ShouldBeExplicit (goconvey)
├── BaseTestCaseGetters.go          # Input(), Expected(), FormTitle()
├── BaseTestCaseValidation.go       # TypeValidationError()
├── SimpleTestCase.go               # Lightweight test case (no type verification)
├── SimpleTestCaseWrapper.go        # Legacy all-any test wrapper interface
├── TypedTestCaseWrapper.go         # Generic-first typed interface
├── TypedTestCaseWrapperContractsBinder.go # Generic contracts binder
├── ShouldAsserter.go               # Assertion interfaces
├── getAssert.go                    # GetAssert singleton — formatting & conversion
└── vars.go                         # Package-level exports

errcore/
└── AssertDiffOnMismatch.go         # Diff-based assertion (non-goconvey)

tests/integratedtests/              # All test implementations
├── {pkg}tests/
│   ├── SomeFeature_test.go         # Test logic (linear, no branching)
│   └── SomeFeature_testcases.go    # Test data (named variables)
```

---

## File Organization

### Rule: Separate test data from test logic

| File Pattern | Contains | Example |
|---|---|---|
| `*_testcases.go` | Named `CaseV1` variables, custom wrapper structs | `Hashmap_testcases.go` |
| `*_test.go` | `Test_*` functions with AAA structure | `Hashmap_test.go` |

### Rule: One test function per scenario

**NEVER** use switch/case, if/else, or index-based dispatching inside test functions. Each logical branch gets its own `Test_Type_Method_Scenario` function.

```go
// ✅ CORRECT — one function per scenario
func Test_Hashmap_GetFound(t *testing.T) { ... }
func Test_Hashmap_GetNotFound(t *testing.T) { ... }

// ❌ BANNED — dispatching inside a test
func Test_Hashmap_Get(t *testing.T) {
    for _, tc := range cases {
        if tc.ShouldFind {
            // ...
        } else {
            // ...
        }
    }
}
```

### Rule: Package naming

Test packages use the `{pkg}tests` suffix:

```
coredata/coregeneric/  →  tests/integratedtests/coregenerictests/
coredata/corestr/      →  tests/integratedtests/corestrtests/
errcore/               →  tests/integratedtests/errcoretests/
chmodhelper/           →  tests/integratedtests/chmodhelpertests/
```

### Rule: File naming

| File Type | Pattern | Example |
|---|---|---|
| Test logic | `{StructOrFunc}_test.go` | `PairFromSplit_test.go` |
| Test cases | `{StructOrFunc}_testcases.go` | `PairFromSplit_testcases.go` |
| Test wrapper | `{StructOrFunc}TestWrapper.go` | `RwxInstructionTestWrapper.go` |
| Shared helpers | descriptive name | `pathInstructionsV3.go` |

---

## Test Case Structures

### CaseV1 (Primary — use this)

`CaseV1` is a type definition of `BaseTestCase` and is the standard test case structure:

```go
type CaseV1 coretests.BaseTestCase
```

#### BaseTestCase fields

```go
type BaseTestCase struct {
    Title           string         // Test case header / description
    ArrangeInput    any            // Preparation input (often args.Map)
    ActualInput     any            // Dynamically set after Act phase via SetActual
    ExpectedInput   any            // Expected result (string or []string)
    Additional      any            // Extra context data
    CustomFormat    string         // Custom output format
    VerifyTypeOf    *VerifyTypeOf  // Optional type verification
    Parameters      *args.Holder   // Extra parameters for complex tests
    IsEnable        issetter.Value // Set to False to disable a test case
    HasError        bool           // Whether error is expected
    HasPanic        bool           // Whether panic is expected
    IsValidateError bool           // Whether to validate error content
}
```

### SimpleTestCase (Lightweight alternative)

For simpler scenarios without type verification:

```go
type SimpleTestCase struct {
    Title         string
    ArrangeInput  any
    ActualInput   any
    ExpectedInput any
    Params        args.Map
}
```

### Test Case Wrapper Interfaces

#### TypedTestCaseWrapper[TInput, TExpect] (Generic-first — use this for new code)

```go
type TypedTestCaseWrapper[TInput, TExpect any] interface {
    CaseTitle() string
    TypedInput() TInput
    TypedExpected() TExpect
    TypedActual() TExpect
    SetTypedActual(actual TExpect)
}
```

`GenericGherkins[TInput, TExpect]` implements this interface.

#### SimpleTestCaseWrapper (Legacy — backward compatible)

```go
type SimpleTestCaseWrapper interface {
    CaseTitle() string
    Input() any
    Expected() any
    Actual() any
    SetActual(actual any)
}
```

`CaseV1`, `BaseTestCase`, and `SimpleTestCase` implement this interface.

### Key Imports

```go
import (
    "testing"

    "gitlab.com/auk-go/core/coretests"              // GetAssert, BaseTestCase, TypedTestCaseWrapper
    "gitlab.com/auk-go/core/coretests/args"          // Map, FuncWrap, Holder
    "gitlab.com/auk-go/core/coretests/coretestcases" // CaseV1, GenericGherkins
    "gitlab.com/auk-go/core/errcore"                 // AssertDiffOnMismatch
)
```

---

## AAA Pattern (Arrange-Act-Assert)

**Every test function MUST have three clearly labeled sections:**

```go
func Test_SomeType_SomeMethod_Scenario(t *testing.T) {
    tc := someNamedTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    value, _ := input.GetAsString("key")

    // Act
    result := pkg.SomeMethod(value)

    // Assert
    tc.ShouldBeEqual(t, 0, result)
}
```

### Rules

1. **Always write `// Arrange`, `// Act`, `// Assert` comments** — never skip them.
2. **Arrange** sets up inputs from the test case.
3. **Act** invokes exactly one function/method under test.
4. **Assert** uses framework assertions — **never** raw `t.Error()` / `t.Fatal()`.

---

## Assertion Methods

### Two Assertion Paradigms

This project supports two assertion styles. Both are correct — choose based on context:

| Style | When to Use | Method |
|---|---|---|
| **GoConvey (CaseV1)** | Most tests — structured BDD-style output | `tc.ShouldBeEqual(t, idx, actuals...)` |
| **Diff-based (errcore)** | Custom wrappers, non-CaseV1 structs | `errcore.AssertDiffOnMismatch(t, idx, title, act, exp)` |

### CaseV1 Assertion Methods (GoConvey-based)

These are the primary assertion methods — they use GoConvey's `convey.Convey` and `convey.So` under the hood:

| Method | Comparison | Use When |
|---|---|---|
| `ShouldBeEqual(t, idx, actuals...)` | Exact string equality | **Default choice** for most tests |
| `ShouldBeTrimEqual(t, idx, actuals...)` | Trim whitespace, then equal | Output may have leading/trailing spaces |
| `ShouldBeSortedEqual(t, idx, actuals...)` | Sort + trim, then equal | Order doesn't matter (e.g., map keys) |
| `ShouldContains(t, idx, actuals...)` | Substring contains | Partial matching |
| `ShouldStartsWith(t, idx, actuals...)` | Prefix match | Check beginning of output |
| `ShouldEndsWith(t, idx, actuals...)` | Suffix match | Check ending of output |
| `ShouldBeNotEqual(t, idx, actuals...)` | Not equal | Negative assertions |
| `ShouldBeRegex(t, idx, actuals...)` | Regex match per line | Pattern-based validation |
| `ShouldBeTrimRegex(t, idx, actuals...)` | Trim + regex | Trimmed pattern matching |
| `ShouldHaveNoError(t, title, idx, err)` | Error is nil | Verify no error occurred |
| `AssertDirectly(t, title, msg, idx, actual, assertion, expected)` | Any convey assertion | Custom assertions |

#### ShouldBeEqual — primary usage

```go
func Test_Pair_FromSplit_Valid(t *testing.T) {
    tc := pairFromSplitValidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("sep")

    // Act
    pair := coregeneric.PairFromSplit(str, sep)

    // Assert
    tc.ShouldBeEqual(t, 0,
        pair.Left,
        pair.Right,
        fmt.Sprintf("%v", pair.IsValid),
    )
}
```

#### ShouldBeSortedEqual — order-independent

```go
func Test_Hashmap_Keys_ReturnAll(t *testing.T) {
    tc := hashmapKeysTestCase

    // Arrange
    hm := coregeneric.HashmapFrom(map[string]int{"b": 2, "a": 1})

    // Act
    keys := hm.Keys()

    // Assert — map key order is non-deterministic
    tc.ShouldBeSortedEqual(t, 0, keys...)
}
```

#### ShouldContains — partial matching

```go
func Test_Error_ContainsContext(t *testing.T) {
    tc := errorContextTestCase

    // Act
    err := someFunc()

    // Assert
    tc.ShouldContains(t, 0, err.Error())
}
```

#### ShouldHaveNoError — nil error check via GoConvey

```go
tc.ShouldHaveNoError(t, "serialization step", 0, err)
```

#### AssertDirectly — custom convey.Assertion

```go
tc.AssertDirectly(
    t,
    "custom context",
    "value should be positive",
    caseIndex,
    actualValue,
    convey.ShouldBeGreaterThan,
    0,
)
```

### How ShouldBeEqual works internally

Understanding the chain helps when debugging:

```
ShouldBeEqual(t, idx, actuals...)
  → ShouldBe(t, idx, stringcompareas.Equal, actuals...)
    → ShouldBeUsingCondition(t, idx, Equal, DefaultDisabledCondition, actuals...)
      → VerifyAllCondition(idx, Equal, condition, actuals)
        → SliceValidator{}.AllVerifyError(&Parameter{})
      → convey.Convey(title, t, func() {
            convey.So(validationError, should.BeNil)
        })
```

Each actual element is compared 1:1 against its corresponding expected line.

---

## errcore.AssertDiffOnMismatch

For tests using custom wrappers or when not using GoConvey assertion chains, use the diff-based assertion:

```go
func AssertDiffOnMismatch(
    t *testing.T,
    caseIndex int,
    title string,
    actLines []string,
    expectedLines []string,     // can be any (normalizes string → []string)
    contextLines ...string,     // optional diagnostic lines printed on failure
)
```

### Basic usage

```go
func Test_LinkedList_Empty(t *testing.T) {
    tc := linkedListEmptyTestCase
    ll := coregeneric.EmptyLinkedList[int]()

    // Act
    actLines := []string{
        fmt.Sprintf("%v", ll.IsEmpty()),
        fmt.Sprintf("%v", ll.Length()),
        fmt.Sprintf("%v", ll.HasItems()),
    }

    // Assert
    errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}
```

### With context lines (extra diagnostics on failure)

```go
errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
    fmt.Sprintf("  InitValue: %d", tc.InitValue),
    fmt.Sprintf("  CompareValue: %d", tc.CompareValue),
)
```

### Error-specific variant

```go
errcore.AssertErrorDiffOnMismatch(t, caseIndex, tc.Title, err, expectedLines)
```

---

## Input Management (args.Map)

`args.Map` is a `map[string]any` with typed accessors. Use it as `ArrangeInput` for structured test inputs.

### Creating test cases with args.Map

```go
var myTestCase = coretestcases.CaseV1{
    Title: "Positive: valid email is accepted",
    ArrangeInput: args.Map{
        "input":  "user@example.com",
        "strict": true,
    },
    ExpectedInput: "true",
}
```

### Accessing values — ALWAYS handle errors

```go
// Arrange
input := tc.ArrangeInput.(args.Map)

// String — returns (string, bool)
email, valid := input.GetAsString("input")
errcore.HandleErrMessage("input", !valid)  // panics if key missing/invalid

// Int — returns (int, bool)
count, valid := input.GetAsInt("count")

// String with default (empty string if missing)
name := input.GetAsStringDefault("name")

// Int with default
limit := input.GetAsIntDefault("limit", 10)

// Direct access (no type safety, returns nil if missing)
raw := input.GetDirectLower("key")

// String slice
parts, valid := input.GetAsStrings("parts")

// Check existence before access
if input.HasDefined("optional") { ... }
if input.IsKeyMissing("key") { ... }
```

### Special keys in args.Map

| Key | Accessor | Purpose |
|---|---|---|
| `"when"` | `input.When()` | Scenario description |
| `"actual"` | `input.Actual()` | The value under test |
| `"expect"` | `input.Expect()` | Expected outcome |
| `"func"` | `input.WorkFunc()` | Function to invoke via reflection |
| `"first"` / `"f1"` / `"p1"` | `input.FirstItem()` | Positional arg |
| `"second"` / `"f2"` | `input.SecondItem()` | Positional arg |

### Function invocation via args.Map

```go
ArrangeInput: args.Map{
    "func": myTargetFunction,
    "p1":   "arg1",
    "p2":   42,
},

// In test:
input := tc.ArrangeInput.(args.Map)
results, err := input.InvokeWithValidArgs()
```

---

## Named Test Case Variables

### Rule: NEVER use indexed slice access

```go
// ❌ BANNED — fragile, order-dependent
var testCases = []coretestcases.CaseV1{...}
tc := testCases[0]  // breaks when cases are reordered

// ✅ CORRECT — resilient, self-documenting
var hashmapGetFoundTestCase = coretestcases.CaseV1{...}
tc := hashmapGetFoundTestCase
```

### Naming convention

```
{lowerCamelType}{Method}{Scenario}TestCase

Examples:
  hashmapGetFoundTestCase
  linkedListAddFrontEmptyTestCase
  pairFromSplitValidTestCase
  hashsetIsEqualsBothNilTestCase
  bytesErrorOnceExecuteTestCase
```

### When loops are still appropriate

Loops over **named slices** are fine when all cases share **identical test logic** (no branching):

```go
// ✅ OK — homogeneous cases, identical Arrange/Act/Assert paths
var createTestCases = []coretestcases.CaseV1{
    {Title: "Valid pattern", ArrangeInput: args.Map{"pattern": `\d+`}, ExpectedInput: []string{"true", "false"}},
    {Title: "Invalid pattern", ArrangeInput: args.Map{"pattern": "[bad"}, ExpectedInput: []string{"false", "true"}},
}

func Test_Create_Verification(t *testing.T) {
    for caseIndex, tc := range createTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        pattern, _ := input.GetAsString("pattern")

        // Act
        regex, err := regexnew.New.DefaultLock(pattern)

        // Assert
        tc.ShouldBeEqual(t, caseIndex,
            fmt.Sprintf("%v", regex != nil),
            fmt.Sprintf("%v", err != nil),
        )
    }
}
```

---

## ExpectedInput Formats

### Single string (preferred for single values)

```go
var myTestCase = coretestcases.CaseV1{
    Title:         "Length returns 3",
    ExpectedInput: "3",
}
```

### String slice (for multi-line positional assertions)

```go
var myTestCase = coretestcases.CaseV1{
    Title:         "Returns left, right, valid",
    ExpectedInput: []string{"hello", "world", "true"},
}
```

### args.Map (preferred for multi-value assertions)

Use `args.Map` as `ExpectedInput` when the test verifies multiple properties.
This produces **self-documenting failure output** where every value has a descriptive key.

```go
var myTestCase = coretestcases.CaseV1{
    Title: "New creates Variant with correct value",
    ArrangeInput: args.Map{
        "when":  "given byte value 5",
        "input": 5,
    },
    ExpectedInput: args.Map{
        "value":     5,      // raw int — no fmt.Sprintf needed
        "isZero":    false,   // raw bool
        "isInvalid": false,
        "isValid":   true,
    },
}
```

In the test body, construct actual results as `args.Map` with the **same keys** and raw values:

```go
func Test_Variant(t *testing.T) {
    for caseIndex, tc := range testCases {
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

**How it works internally:**
- `ShouldBeEqualMap` calls `CompileToStrings()` on both actual and expected maps
- `CompileToStrings()` sorts keys and formats each entry as `"key : %v"` (e.g., `"isZero : false"`)
- Comparison is done as sorted string lines, ensuring deterministic order

**Failure output example:**
```
isValid : true       ← expected
isValid : false      ← actual     [MISMATCH]
```

**When to use which format:**

| Format | Use When |
|--------|----------|
| `string` | Single value assertion |
| `[]string` | Multi-value positional (collection items, ordered output) |
| `args.Map` | Multi-property assertions (object state, method results) |

### How ExpectedLines() works internally

`CaseV1.ExpectedLines()` normalizes `ExpectedInput` to `[]string` via `convertinternal.AnyTo.Strings()`:

- `string` → `[]string{s}`
- `[]string` → as-is
- `int`, `bool`, `byte`, etc. → converted via `strconv`
- `map[string]any` → sorted `"key : value"` lines
- `args.Map` → handled via `CompileToStrings()` in `ShouldBeEqualMap`
- Other types → PrettyJSON fallback

### Rule: Match actual output to expected line-by-line

Each actual element corresponds to one expected line:

```go
// Test case
ExpectedInput: []string{"alice", "30", "true"},

// Test assertion — each arg maps positionally
tc.ShouldBeEqual(t, 0,
    user.Name,                           // → "alice"
    fmt.Sprintf("%d", user.Age),         // → "30"
    fmt.Sprintf("%v", user.IsActive),    // → "true"
)
```

---

## SliceValidator and VerifyAll

For advanced validation scenarios, `CaseV1` exposes a `SliceValidator` builder:

```go
// Create a validator with custom comparison mode
validator := tc.SliceValidator(stringcompareas.Contains, actualLines)

// Or with conditions (trim, sort)
validator := tc.SliceValidatorCondition(
    stringcompareas.Equal,
    corevalidator.DefaultTrimCoreCondition,
    actualLines,
)

// Verify and get error
err := tc.VerifyAllSliceValidator(caseIndex, validator)
```

### VerifyAll / VerifyError (return error instead of asserting)

```go
// Returns error instead of calling convey.So
err := tc.VerifyAllEqual(caseIndex, actuals...)

// Combines value verification + type verification
err := tc.VerifyError(caseIndex, stringcompareas.Equal, actuals...)

// First-line-only verification (for partial checks)
err := tc.VerifyFirst(caseIndex, stringcompareas.Equal, actuals)
```

---

## GoConvey Integration

All `ShouldBe*` methods use GoConvey's `convey.Convey` and `convey.So` internally. You should **never** call `convey.Convey` directly in test functions — let the assertion methods handle it.

### Direct GoConvey (only for custom assertions)

Use `AssertDirectly` if you need a custom `convey.Assertion`:

```go
tc.AssertDirectly(
    t,
    "additional context",
    "assertion message",
    caseIndex,
    actualValue,
    convey.ShouldBeGreaterThan,
    expectedMinimum,
)
```

---

## Comparison Modes

The `stringcompareas.Variant` enum controls how actual vs expected lines are compared:

| Variant | Behavior | Method Shortcut |
|---|---|---|
| `Equal` | Exact string match | `ShouldBeEqual` |
| `Contains` | Actual contains expected | `ShouldContains` |
| `StartsWith` | Actual starts with expected | `ShouldStartsWith` |
| `EndsWith` | Actual ends with expected | `ShouldEndsWith` |
| `NotEqual` | Actual ≠ expected | `ShouldBeNotEqual` |
| `Regex` | Expected is regex pattern | `ShouldBeRegex` |

---

## Conditions (Trim, Sort)

Conditions preprocess actual/expected lines before comparison:

| Condition | Effect | Method Shortcut |
|---|---|---|
| `DefaultDisabledCoreCondition` | No preprocessing (default) | `ShouldBeEqual` |
| `DefaultTrimCoreCondition` | `strings.TrimSpace` on each line | `ShouldBeTrimEqual` |
| `DefaultSortTrimCoreCondition` | Sort + trim both sides | `ShouldBeSortedEqual` |

Custom condition via `ShouldBeUsingCondition`:

```go
tc.ShouldBeUsingCondition(
    t, caseIndex,
    stringcompareas.Equal,
    corevalidator.DefaultTrimCoreCondition,
    actuals...,
)
```

---

## Type Verification

Optional automatic type checking via `VerifyTypeOf`:

```go
var myTestCase = coretestcases.CaseV1{
    Title:         "Returns string result",
    ExpectedInput: "hello",
    VerifyTypeOf: &coretests.VerifyTypeOf{
        ArrangeInput:  reflect.TypeOf(args.Map{}),
        ActualInput:   reflect.TypeOf(""),
        ExpectedInput: reflect.TypeOf(""),
    },
}
```

This automatically asserts that `ArrangeInput`, `ActualInput`, and `ExpectedInput` match the specified `reflect.Type` values after the test runs. Mismatches produce a clear error like `"Arrange Type Mismatch"`.

---

## Error Testing Patterns

### Testing error existence

```go
var errorExpectedTestCase = coretestcases.CaseV1{
    Title:         "Invalid input returns error",
    ArrangeInput:  args.Map{"input": ""},
    ExpectedInput: []string{"true"},
}

func Test_Validate_InvalidInput(t *testing.T) {
    tc := errorExpectedTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")

    // Act
    _, err := validate(str)

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err != nil))
}
```

### Testing error message content

```go
var errorMessageTestCase = coretestcases.CaseV1{
    Title:         "Error contains field name",
    ExpectedInput: []string{"true", "true"},
}

func Test_Validate_ErrorContainsFieldName(t *testing.T) {
    tc := errorMessageTestCase

    // Act
    _, err := validate("")

    // Assert
    tc.ShouldBeEqual(t, 0,
        fmt.Sprintf("%v", err != nil),
        fmt.Sprintf("%v", strings.Contains(err.Error(), "required")),
    )
}
```

### Using errcore.AssertErrorDiffOnMismatch

Splits an error into lines and compares:

```go
errcore.AssertErrorDiffOnMismatch(t, caseIndex, tc.Title, err, expectedLines,
    fmt.Sprintf("  Input: %s", inputValue),
)
```

### ShouldHaveNoError (GoConvey nil check)

```go
tc.ShouldHaveNoError(t, "deserialization step", 0, err)
```

---

## Panic Testing

Use a helper function to catch panics:

```go
// Helper — define once per test package
func callPanics(fn func()) (panicked bool) {
    defer func() {
        if r := recover(); r != nil {
            panicked = true
        }
    }()
    fn()
    return false
}
```

### Usage in test

```go
var mustPanicTestCase = coretestcases.CaseV1{
    Title:         "MustHaveSafeItems panics when empty",
    ExpectedInput: "true",
}

func Test_MustHaveSafeItems_PanicsWhenEmpty(t *testing.T) {
    tc := mustPanicTestCase

    // Arrange
    once := createEmptyOnce()

    // Act
    panicked := callPanics(func() { once.MustHaveSafeItems() })

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", panicked))
}
```

### Panic with value capture

```go
func callPanicsWithValue(fn func()) (panicked bool, value any) {
    defer func() {
        if r := recover(); r != nil {
            panicked = true
            value = r
        }
    }()
    fn()
    return false, nil
}
```

---

## Concurrency Testing

For thread-safe method tests (`*Lock` methods), use `sync.WaitGroup`:

```go
func Test_Hashset_AddLock_ConcurrentSafety(t *testing.T) {
    const goroutines = 500
    hs := coregeneric.NewHashset[int](goroutines)

    wg := sync.WaitGroup{}
    wg.Add(goroutines)

    for i := 0; i < goroutines; i++ {
        go func(idx int) {
            hs.AddLock(idx)
            wg.Done()
        }(i)
    }

    wg.Wait()

    got := hs.Length()
    if got != goroutines {
        t.Errorf("AddLock concurrent: expected %d items, got %d", goroutines, got)
    }
}
```

**Note:** Concurrency tests **may** use `t.Errorf` directly since they validate thread safety (no panics, correct final state) rather than functional output lines.

### Mixed read/write concurrency

```go
func Test_Hashset_ContainsLock_ConcurrentReadsWrites(t *testing.T) {
    const writers = 200
    const readers = 200
    hs := coregeneric.NewHashset[int](writers)

    wg := sync.WaitGroup{}
    wg.Add(writers + readers)

    // Concurrent writers
    for i := 0; i < writers; i++ {
        go func(idx int) {
            hs.AddLock(idx)
            wg.Done()
        }(i)
    }

    // Concurrent readers — must not panic
    for i := 0; i < readers; i++ {
        go func(idx int) {
            _ = hs.ContainsLock(idx)
            wg.Done()
        }(i)
    }

    wg.Wait()

    got := hs.Length()
    if got != writers {
        t.Errorf("Expected %d, got %d", writers, got)
    }
}
```

---

## Custom Test Wrappers

For domain-specific test data, wrap `CaseV1` in a custom struct:

### Defining the wrapper

```go
// bytesErrorOnce_testcases.go
type bytesErrorOnceTestCase struct {
    Case          coretestcases.CaseV1
    InitBytes     []byte
    InitErr       error
    IsNilReceiver bool
}
```

### Named test case with wrapper

```go
var bytesErrorOnceExecuteTestCase = bytesErrorOnceTestCase{
    Case: coretestcases.CaseV1{
        Title:         "Execute returns same as Value",
        ExpectedInput: []string{"true"},
    },
    InitBytes: []byte("exec"),
}
```

### Using the wrapper in a test

```go
func Test_BytesErrorOnce_Execute(t *testing.T) {
    tc := bytesErrorOnceExecuteTestCase

    // Arrange
    once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
        return tc.InitBytes, nil
    })

    // Act
    v1, _ := once.Execute()
    v2, _ := once.Value()

    // Assert
    actLines := []string{fmt.Sprintf("%v", string(v1) == string(v2))}
    expectedLines := tc.Case.ExpectedInput.([]string)
    errcore.AssertDiffOnMismatch(t, 0, tc.Case.Title, actLines, expectedLines)
}
```

### Looping over wrapper slices (homogeneous logic only)

```go
var bytesErrorOnceCoreTestCases = []bytesErrorOnceTestCase{
    {Case: coretestcases.CaseV1{Title: "abc", ExpectedInput: []string{...}}, InitBytes: []byte("abc")},
    {Case: coretestcases.CaseV1{Title: "nil", ExpectedInput: []string{...}}, InitBytes: nil},
}

func Test_BytesErrorOnce_Core(t *testing.T) {
    for caseIndex, tc := range bytesErrorOnceCoreTestCases {
        // Arrange
        once := createOnce(tc.InitBytes, tc.InitErr)

        // Act
        val, err := once.Value()
        actLines := buildActLines(val, err, once)

        // Assert
        expectedLines := tc.Case.ExpectedInput.([]string)
        errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
    }
}
```

---

## Anti-Patterns (Banned)

### ❌ Never use raw `t.Error` / `t.Fatal` / `t.Errorf` for functional tests

```go
// ❌ BANNED
if result != expected {
    t.Errorf("got %v, want %v", result, expected)
}

// ✅ USE INSTEAD
tc.ShouldBeEqual(t, 0, result)
// or
errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLines)
```

### ❌ Never inline test data in test functions

```go
// ❌ BANNED
func Test_Something(t *testing.T) {
    expected := "hello"
    actual := doThing()
    if actual != expected { t.Error("nope") }
}

// ✅ USE INSTEAD — data in _testcases.go file
var somethingTestCase = coretestcases.CaseV1{
    Title:         "Returns hello",
    ExpectedInput: "hello",
}
```

### ❌ Never use indexed slice access for test cases

```go
// ❌ BANNED
tc := testCases[0]

// ✅ USE INSTEAD
tc := mySpecificTestCase
```

### ❌ Never ignore input getter errors

```go
// ❌ BANNED — silently fails, test passes incorrectly
str, _ := input.GetAsString("key")

// ✅ CORRECT — panics immediately on missing input
str, valid := input.GetAsString("key")
errcore.HandleErrMessage("key", !valid)
```

### ❌ Never branch inside test functions

```go
// ❌ BANNED
if tc.HasError {
    // different assertion path
} else {
    // another path
}

// ✅ CORRECT — separate test functions
func Test_Method_WithError(t *testing.T) { ... }
func Test_Method_WithoutError(t *testing.T) { ... }
```

### ❌ Never skip AAA comments

```go
// ❌ BANNED
func Test_X(t *testing.T) {
    tc := xTestCase
    result := doX()
    tc.ShouldBeEqual(t, 0, result)
}

// ✅ CORRECT — always label sections
func Test_X(t *testing.T) {
    tc := xTestCase

    // Arrange
    // (no setup needed)

    // Act
    result := doX()

    // Assert
    tc.ShouldBeEqual(t, 0, result)
}
```

### ❌ Never call convey.Convey directly

```go
// ❌ BANNED — framework handles convey wrapping
convey.Convey("my test", t, func() {
    convey.So(actual, convey.ShouldEqual, expected)
})

// ✅ USE INSTEAD
tc.ShouldBeEqual(t, 0, actual)
```

---

## Complete Examples

### Example 1: Named test cases with ShouldBeEqual

```go
// PairFromSplit_testcases.go
package coregenerictests

import (
    "gitlab.com/auk-go/core/coretests/args"
    "gitlab.com/auk-go/core/coretests/coretestcases"
)

var pairFromSplitValidTestCase = coretestcases.CaseV1{
    Title: "PairFromSplit with dot separator",
    ArrangeInput: args.Map{
        "input": "key.value",
        "sep":   ".",
    },
    ExpectedInput: []string{"key", "value", "true", ""},
}

var pairFromSplitMissingSepTestCase = coretestcases.CaseV1{
    Title: "PairFromSplit without separator returns invalid",
    ArrangeInput: args.Map{
        "input": "noseparator",
        "sep":   ".",
    },
    ExpectedInput: []string{"noseparator", "", "false", "separator not found"},
}
```

```go
// PairFromSplit_test.go
package coregenerictests

import (
    "fmt"
    "testing"

    "gitlab.com/auk-go/core/coredata/coregeneric"
    "gitlab.com/auk-go/core/coretests/args"
    "gitlab.com/auk-go/core/errcore"
)

func Test_PairFromSplit_Valid(t *testing.T) {
    tc := pairFromSplitValidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("sep")

    // Act
    pair := coregeneric.PairFromSplit(str, sep)

    // Assert
    tc.ShouldBeEqual(t, 0,
        pair.Left,
        pair.Right,
        fmt.Sprintf("%v", pair.IsValid),
        pair.Message,
    )
}

func Test_PairFromSplit_MissingSep(t *testing.T) {
    tc := pairFromSplitMissingSepTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("sep")

    // Act
    pair := coregeneric.PairFromSplit(str, sep)

    // Assert
    tc.ShouldBeEqual(t, 0,
        pair.Left,
        pair.Right,
        fmt.Sprintf("%v", pair.IsValid),
        pair.Message,
    )
}
```

### Example 2: Using errcore.AssertDiffOnMismatch with custom wrapper

```go
// IntegerOnce_testcases.go
type integerOnceTestCase struct {
    Case         coretestcases.CaseV1
    InitValue    int
    CompareValue int
}

var integerOnceBasicTestCase = integerOnceTestCase{
    Case: coretestcases.CaseV1{
        Title:         "IntegerOnce returns initialized value",
        ExpectedInput: []string{"42", "true"},
    },
    InitValue: 42,
}
```

```go
// IntegerOnce_test.go
func Test_IntegerOnce_Basic(t *testing.T) {
    tc := integerOnceBasicTestCase

    // Arrange
    once := coreonce.NewIntegerOncePtr(func() int { return tc.InitValue })

    // Act
    val := once.Value()

    // Assert
    actLines := []string{
        fmt.Sprintf("%d", val),
        fmt.Sprintf("%v", once.IsInitialized()),
    }
    expectedLines := tc.Case.ExpectedInput.([]string)
    errcore.AssertDiffOnMismatch(t, 0, tc.Case.Title, actLines, expectedLines,
        fmt.Sprintf("  InitValue: %d", tc.InitValue),
    )
}
```

### Example 3: Single-value ExpectedInput (string shorthand)

```go
// _testcases.go
var linkedListCollectionTestCase = coretestcases.CaseV1{
    Title:         "Collection converts",
    ExpectedInput: "2",                 // single string, NOT []string{"2"}
}

// _test.go
func Test_LinkedList_Collection(t *testing.T) {
    tc := linkedListCollectionTestCase

    // Arrange
    ll := coregeneric.LinkedListFrom([]int{1, 2})

    // Act
    col := ll.Collection()

    // Assert — ShouldBeEqual normalizes "2" → []string{"2"}
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", col.Length()))
}
```

---

## Checklist for New Tests

When adding tests for a new package:

- [ ] Create directory: `tests/integratedtests/{package}tests/`
- [ ] Create test cases file: `{Feature}_testcases.go` with named variables
- [ ] Create test file: `{Feature}_test.go` with individual `Test_` functions
- [ ] Follow AAA pattern with explicit `// Arrange`, `// Act`, `// Assert` comments
- [ ] Pass `caseIndex` to all assertion calls (use `0` for single-case tests)
- [ ] Use named test case variables (never indexed slice access)
- [ ] Use `string` for single-value expectations, `[]string` for multi-value
- [ ] Use `ShouldBeEqual` or `errcore.AssertDiffOnMismatch` — never raw `t.Error`
- [ ] Handle `args.Map` getter errors (never ignore the `bool` return)
- [ ] One test function per scenario — no branching in test bodies
- [ ] Run `make run-tests` to verify

---

## Future: GenericGherkins Proposal

The existing `SimpleGherkins` struct uses all-string fields (`Feature`, `Given`, `When`, `Then`, `Expect`, `Actual`), which lacks type safety and forces callers to convert everything to strings. Many test cases (e.g., regex tests) extract typed values like `pattern`, `input`, `isMatch` from `args.Map` — losing compile-time checks and IDE discoverability.

### Problem

```go
// Current — untyped args.Map, easy to misspell keys, no compile-time safety
ArrangeInput: args.Map{
    "when":    "given an invalid regex pattern",
    "pattern": "[invalid",
    "input":   "anything",
    "isMatch": false,  // bool stored as any
},
```

### Solution: GenericGherkins[TInput, TExpect]

A generic struct that provides typed fields for common test case properties while keeping an optional `ExtraArgs` map for dynamic overflow:

```go
// GenericGherkins — typed test case representation
//
// Use this when you want compile-time type safety for Input/Expected fields
// instead of extracting values from args.Map at runtime.
//
// TInput  — type of the test input (e.g., string for regex pattern)
// TExpect — type of the expected result (e.g., bool for IsMatch)
type GenericGherkins[TInput, TExpect any] struct {
    Title      string       // Test case header / scenario name
    Feature    string       // Feature being tested
    Given      string       // Precondition
    When       string       // Scenario description
    Then       string       // Expected outcome description
    Input      TInput       // Typed input value
    Expected   TExpect      // Typed expected value
    Actual     TExpect      // Typed actual value (set after Act phase)
    IsMatching bool         // Whether a match is expected (for validation tests)
    ExtraArgs  args.Map     // Optional dynamic key-value pairs for overflow
}
```

### Type Aliases

```go
// AnyGherkins — all-any version for maximum flexibility
type AnyGherkins = GenericGherkins[any, any]

// SimpleGherkins — backward-compatible all-string version (existing struct unchanged)
// Note: SimpleGherkins remains as the legacy struct for backward compatibility.
// New tests should prefer GenericGherkins with concrete types.
```

### Key Methods

```go
// TypedTestCaseWrapper implementation — compile-time safe access
func (it *GenericGherkins[TInput, TExpect]) CaseTitle() string
func (it *GenericGherkins[TInput, TExpect]) TypedInput() TInput
func (it *GenericGherkins[TInput, TExpect]) TypedExpected() TExpect
func (it *GenericGherkins[TInput, TExpect]) TypedActual() TExpect
func (it *GenericGherkins[TInput, TExpect]) SetTypedActual(actual TExpect)
func (it *GenericGherkins[TInput, TExpect]) AsTypedTestCaseWrapper() coretests.TypedTestCaseWrapper[TInput, TExpect]

// IsFailedToMatch — inverse of IsMatching.
// Use when validating that a mismatch is expected.
func (it *GenericGherkins[TInput, TExpect]) IsFailedToMatch() bool

// ShouldBeEqual — assert actLines match expectedLines using the struct's Title.
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqual(
    t *testing.T, caseIndex int, actLines []string, expectedLines []string,
)

// CompareWith — structural comparison against another GenericGherkins.
// Returns whether they are equal and a diff string on mismatch.
func (it *GenericGherkins[TInput, TExpect]) CompareWith(
    other *GenericGherkins[TInput, TExpect],
) (isEqual bool, diff string)

// String / ToString — formatted printing with Gherkins layout.
func (it *GenericGherkins[TInput, TExpect]) String() string
func (it *GenericGherkins[TInput, TExpect]) ToString(testIndex int) string

// FullString — verbose representation including all fields for debugging.
func (it *GenericGherkins[TInput, TExpect]) FullString() string
```

### Interface Hierarchy

```
TypedTestCaseWrapper[TInput, TExpect]    (generic-first — new code)
├── GenericGherkins[TInput, TExpect]      implements via GenericGherkinsTypedWrapper.go
│   ├── AnyGherkins     = GenericGherkins[any, any]
│   ├── StringGherkins  = GenericGherkins[string, string]
│   └── StringBoolGherkins = GenericGherkins[string, bool]
│
SimpleTestCaseWrapper                     (legacy all-any — backward compatible)
├── CaseV1                                implements via CaseV1.go
├── BaseTestCase                          implements via BaseTestCase.go
└── SimpleTestCase                        implements via SimpleTestCase.go
```

### Usage Example — Regex Matching

```go
// _testcases.go — typed, self-documenting, searchable
var lazyRegexInvalidPatternTestCase = GenericGherkins[string, bool]{
    Title:      "Invalid regex pattern has error",
    When:       "given an invalid regex pattern",
    Input:      "[invalid",
    Expected:   false,
    IsMatching: false,
}

// _test.go
func Test_LazyRegex_InvalidPattern(t *testing.T) {
    tc := lazyRegexInvalidPatternTestCase

    // Arrange
    regex := regexnew.New.Lazy(tc.TypedInput())

    // Act
    result := regex.IsMatch("anything")
    tc.SetTypedActual(result)

    // Assert
    tc.ShouldBeEqual(t, 0,
        fmt.Sprintf("%v", tc.TypedActual()),
        fmt.Sprintf("%v", tc.TypedExpected()),
    )
}
```

### Migration Path

| From | To |
|---|---|
| `SimpleGherkins` (all string) | Stays as-is — backward compatible |
| `args.Map` with typed keys | `GenericGherkins[ConcreteInput, ConcreteExpect]` |
| `args.Map` with mixed types | `AnyGherkins` (all-any) with `ExtraArgs` for overflow |
| `SimpleTestCaseWrapper` consumers | `TypedTestCaseWrapper[TInput, TExpect]` for new code |

`SimpleTestCaseWrapper` and `SimpleGherkins` remain unchanged. `TypedTestCaseWrapper` and `GenericGherkins` are implemented in `coretests/` alongside the legacy types.

---

## Future: CaseV2 Proposal

The current `CaseV1` serves well for string-line-based comparison but has known limitations:

### CaseV1 Limitations

1. **String-only comparison** — all actual values must be converted to `string` via `fmt.Sprintf`.
2. **No structured diff** — comparison is line-by-line strings, not typed values.
3. **No built-in error expectation** — `HasError` field exists but isn't enforced automatically.
4. **Manual ActualInput management** — `SetActual` must be called explicitly.
5. **No subtest integration** — doesn't use `t.Run()` for Go's native subtest naming.

### CaseV2 Ideas

A potential `CaseV2` could address these:

| Feature | CaseV1 | CaseV2 (Proposed) |
|---|---|---|
| Expected type | `any` (string or []string) | `ExpectedResult[T]` — typed generics |
| Error handling | Manual `HasError` bool | `ExpectedError *ErrorExpectation` — auto-assert |
| Comparison | String lines only | `reflect.DeepEqual` + typed diff |
| Panic assertion | External helper required | `ExpectsPanic bool` — auto-recover |
| Timeout | Not supported | `Timeout time.Duration` — deadline enforcement |
| Subtest naming | Manual `caseIndex` | Auto `t.Run(tc.Title)` subtests |
| Cleanup | Not supported | `Cleanup func()` — auto `t.Cleanup` |

### Proposed CaseV2 structure

```go
type CaseV2[T any] struct {
    Title          string
    ArrangeInput   any
    ExpectedResult T                    // Typed expected value — no fmt.Sprintf needed
    ExpectedError  *ErrorExpectation    // nil = no error expected
    ExpectsPanic   bool
    Timeout        time.Duration
    Cleanup        func()
    Tags           []string             // for filtering: "unit", "integration", "slow"
}

type ErrorExpectation struct {
    ShouldExist  bool
    Contains     string          // substring match on error message
    MatchesRegex string          // pattern match
    IsType       reflect.Type    // typed error assertion (e.g., *os.PathError)
}
```

### CaseV2 usage would look like

```go
var v2TestCase = coretestcases.CaseV2[int]{
    Title:          "Length after 3 adds",
    ExpectedResult: 3,
}

func Test_LinkedList_Length_V2(t *testing.T) {
    tc := v2TestCase

    // Arrange
    ll := coregeneric.EmptyLinkedList[int]()
    ll.Adds(1, 2, 3)

    // Act
    result := ll.Length()

    // Assert — typed comparison, no fmt.Sprintf needed
    tc.AssertEqual(t, result)
}
```

### CaseV2 with error expectation

```go
var v2ErrorCase = coretestcases.CaseV2[*User]{
    Title: "Invalid ID returns not-found error",
    ExpectedError: &coretestcases.ErrorExpectation{
        ShouldExist:  true,
        Contains:     "not found",
    },
}

func Test_FindUser_NotFound_V2(t *testing.T) {
    tc := v2ErrorCase

    // Act
    result, err := findUser("invalid-id")

    // Assert — auto-checks error existence + message content
    tc.AssertResult(t, result, err)
}
```

This is a **proposal only** — CaseV1 remains the standard until CaseV2 is implemented and validated.

---

## Related Docs

- [Folder Spec: coretests](/spec/01-app/folders/07-coretests.md)
- [Testing Patterns (legacy summary)](/spec/01-app/13-testing-patterns.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Edge Case Coverage Audit](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md)
- [Deep Coverage Scan](/spec/13-app-issues/testing/03-deep-coverage-scan.md)
- [GoConvey Migration Plan](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)
