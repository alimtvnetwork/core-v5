# coretestcases — Typed Test Case Structures

## Overview

The `coretestcases` package provides typed test case structures following the Gherkin (Given/When/Then) pattern and the AAA (Arrange-Act-Assert) pattern. It includes assertion helpers, structural comparison, and formatting utilities for consistent, self-documenting test output.

## Architecture

```
coretestcases/
├── GenericGherkins.go              # Core generic struct definition
├── GenericGherkinsAliases.go       # Type aliases (AnyGherkins, StringGherkins, MapGherkins, etc.)
├── GenericGherkinsAssertions.go    # String-line-based assertion methods (ShouldBeEqual, ShouldBeEqualArgs)
├── GenericGherkinsMapAssertions.go # Map-based assertion methods (ShouldBeEqualMap)
├── GenericGherkinsTypedAssertions.go # Typed value assertion (ShouldMatchExpected)
├── GenericGherkinsTypedWrapper.go  # TypedTestCaseWrapper interface implementation
├── GenericGherkinsGetters.go       # Field accessors (IsFailedToMatch, GetExtra, GetExtraAsBool)
├── GenericGherkinsFormatting.go    # String/FullString/ToString formatting
├── GenericGherkinsCompare.go       # Structural comparison (CompareWith)
├── CaseV1.go                       # Legacy CaseV1 (BaseTestCase alias)
├── CaseV1FirstAssertions.go        # CaseV1 convenience assertions
├── CaseV1MapAssertions.go          # CaseV1 map-based assertions
├── CaseNilSafe.go                  # Nil-receiver safety test case
├── CaseNilSafeAssertHelper.go      # Nil-safe assertion helpers
└── consts.go                       # Package constants
```

## Core Types

### GenericGherkins[TInput, TExpect]

The primary typed test case structure. Provides compile-time type safety for `Input` and `Expected` fields.

| Field          | Type       | Description                                      |
|----------------|------------|--------------------------------------------------|
| `Title`        | `string`   | Test case header / scenario name                 |
| `Feature`      | `string`   | Feature being tested                             |
| `Given`        | `string`   | Precondition                                     |
| `When`         | `string`   | Scenario description (fallback for Title)        |
| `Then`         | `string`   | Expected outcome description                     |
| `Input`        | `TInput`   | Typed input value                                |
| `Expected`     | `TExpect`  | Typed expected result                            |
| `Actual`       | `TExpect`  | Typed actual result (set after Act phase)        |
| `IsMatching`   | `bool`     | Whether a match is expected                      |
| `ExpectedLines`| `[]string` | Legacy string output lines (prefer `Expected`)   |
| `ExtraArgs`    | `args.Map` | Optional dynamic key-value overflow              |

### Type Aliases

| Alias              | Definition                              | Use Case                                     |
|--------------------|-----------------------------------------|----------------------------------------------|
| `AnyGherkins`      | `GenericGherkins[any, any]`             | Heterogeneous / unknown types                |
| `StringGherkins`   | `GenericGherkins[string, string]`       | Text-based validation tests                  |
| `StringBoolGherkins`| `GenericGherkins[string, bool]`        | Match/validation tests (regex, search)       |
| `MapGherkins`      | `GenericGherkins[args.Map, args.Map]`   | Multi-field semantic assertions (preferred)  |

### CaseV1

Legacy test case structure aliased from `coretests.BaseTestCase`. Provides `ArrangeInput`, `ActualInput`, `ExpectedInput` fields with map-based assertion support via `ShouldBeEqualMap`.

### CaseNilSafe

Specialized test case for nil-receiver safety testing. Uses direct method references (`Func` field) for compile-time safety and `results.ResultAny` for structured assertions.

## Assertion Methods

### String-Line Assertions (GenericGherkins)

| Method                          | Description                                          |
|---------------------------------|------------------------------------------------------|
| `ShouldBeEqual(t, idx, act, exp)` | Compare actual lines against expected lines        |
| `ShouldBeEqualFirst(t, act, exp)` | ShouldBeEqual with caseIndex=0                     |
| `ShouldBeEqualArgs(t, idx, ...)`  | Each arg treated as one line, compared to ExpectedLines |
| `ShouldBeEqualArgsFirst(t, ...)`  | ShouldBeEqualArgs with caseIndex=0                 |
| `ShouldBeEqualUsingExpected(t, idx, act)` | Compare act against struct's own ExpectedLines |

### Map Assertions (GenericGherkins — MapGherkins)

| Method                            | Description                                        |
|-----------------------------------|----------------------------------------------------|
| `ShouldBeEqualMap(t, idx, actual)` | Compare actual args.Map against Expected args.Map |
| `ShouldBeEqualMapFirst(t, actual)` | ShouldBeEqualMap with caseIndex=0                 |

Map assertions compile both maps to sorted `"key : value"` lines using `CompileToStrings()`, then compare line-by-line. On mismatch, a copy-pasteable Go literal block is printed.

### Typed Assertions (GenericGherkins)

| Method                              | Description                                     |
|-------------------------------------|-------------------------------------------------|
| `ShouldMatchExpected(t, idx, result)` | Compare result against Expected using `%v`    |
| `ShouldMatchExpectedFirst(t, result)` | ShouldMatchExpected with caseIndex=0          |

## Preferred Pattern: MapGherkins

The `MapGherkins` pattern replaces opaque `ExpectedLines` with self-documenting semantic keys:

```go
var testCases = []coretestcases.MapGherkins{
    {
        Title: "Lazy regex matches word pattern",
        When:  "given a simple word pattern",
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
    },
}

// Test runner
for i, tc := range testCases {
    input := tc.Input
    pattern := input["pattern"].(string)
    compareInput := input["compareInput"].(string)

    actual := args.Map{
        "isDefined":    regex.IsDefined(),
        "isApplicable": regex.IsApplicable(),
        "isMatch":      regex.IsMatch(compareInput),
        "isFailedMatch": regex.IsFailedMatch(compareInput),
    }

    tc.ShouldBeEqualMap(t, i, actual)
}
```

### Why MapGherkins over ExpectedLines

| Aspect         | `ExpectedLines`          | `MapGherkins`                    |
|----------------|--------------------------|----------------------------------|
| Readability    | Positional, opaque       | Semantic keys, self-documenting  |
| Diagnostics    | Index-based diff         | Key-labeled diff with Go literal |
| Maintenance    | Order-dependent          | Order-independent                |
| Type safety    | All strings              | Native types (bool, int, string) |

## Getters

| Method                    | Returns           | Description                        |
|---------------------------|-------------------|------------------------------------|
| `CaseTitle()`             | `string`          | Title or When fallback             |
| `IsFailedToMatch()`       | `bool`            | Inverse of IsMatching              |
| `HasExtraArgs()`          | `bool`            | Whether ExtraArgs is non-empty     |
| `GetExtra(key)`           | `any`             | Value from ExtraArgs               |
| `GetExtraAsString(key)`   | `(string, bool)`  | Typed string accessor              |
| `GetExtraAsBool(key)`     | `(bool, bool)`    | Typed bool accessor                |

## Formatting

| Method                              | Description                                    |
|-------------------------------------|------------------------------------------------|
| `String()`                          | Gherkins-formatted string (index 0)            |
| `ToString(testIndex)`               | Gherkins-formatted string with index           |
| `GetWithExpectation(testIndex)`      | Includes Actual and Expected values            |
| `GetMessageConditional(isExp, idx)` | Conditional expectation inclusion              |
| `FullString()`                      | Verbose multi-line debug output of all fields  |

## Comparison

`CompareWith(other)` performs structural field-by-field comparison, returning `(isEqual bool, diff string)`. Handles nil receivers and nil arguments gracefully.

## Related Docs

- [coretests README](../README.md)
- [Testing Guidelines — Test Case Types](../../spec/testing-guidelines/02-test-case-types.md)
- [Testing Guidelines — Assertion Patterns](../../spec/testing-guidelines/05-assertion-patterns.md)
- [Migration Issue — ExpectedLines to MapGherkins](../../spec/13-app-issues/testing/06-regexnewtests-expectedlines-migration.md)
