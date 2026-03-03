# coretests — Test Utilities & Assertions

## Overview

The `coretests` package provides shared testing utilities, assertion helpers, type conversion tools, and test case structures used across all integration tests in the repository.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `GetAssert` | `getAssert` | Central assertion message builder |

## Sub-packages

| Package | Description |
|---------|-------------|
| `args/` | Typed argument maps for test case inputs (`args.Map`) |
| `coretestcases/` | Standardized test case structures (`CaseV1`) |

## Key Capabilities

### Assertion Messages (`GetAssert`)

- **`Quick(when, actual, expected, counter)`** — Standard indexed test message
- **`ToStrings(anyItem)`** — Convert any value to `[]string` for comparison
- **`ToQuoteLines(spaceCount, lines)`** — Format lines as double-quoted with indentation
- **`AnyToDoubleQuoteLines(spaceCount, anyItem)`** — Convert any value to quoted lines
- **`ToString(anyItem)`** — Convert any value to a single string
- **`SortedMessage(isPrint, message, joiner)`** — Sort and join message words
- **`SortedArrayNoPrint(message)`** — Split message into sorted word array
- **`ErrorToLinesWithSpacesDefault(err)`** — Format error as indented lines

### Test Structures

- **`SimpleTestCase`** — Basic test case with input/expected
- **`SimpleTestCaseWrapper`** — Wraps test cases with helper methods
- **`BaseTestCase`** / **`BaseTestCaseWrapper`** — Extended base structures
- **`SimpleGherkins`** — Gherkin-style (Given/When/Then) test representation
- **`DraftType`** / **`AnyToDraftType`** — Draft type conversion for test assertions

### Comparison & Validation

- **`Compare`** — Structural comparison helpers
- **`CompareInstruction`** — Comparison with instructions
- **`isCompare`** — Internal comparison logic
- **`VerifyTypeOf`** — Reflection-based type verification

### Utility Functions

- **`LogOnFail(isPass, expected, actual)`** — Print diff on test failure
- **`ToStringValues(anyItem)`** — Format value using `%v`
- **`ToStringNameValues(anyItem)`** — Format value using `%#v`
- **`SkipOnUnix` / `SkipOnWindows`** — Platform-conditional test skipping

## Test Pattern

Tests follow the AAA pattern with `coretestcases.CaseV1`:

```go
func Test_Example(t *testing.T) {
    cases := []coretestcases.CaseV1{
        {
            Input: args.Map{"key": "value"},
            Expected: "value",
        },
    }

    for i, tc := range cases {
        actual := doSomething(tc.Input)
        assert.Equal(t, tc.Expected, actual,
            coretests.GetAssert.Quick("description", actual, tc.Expected, i))
    }
}
```

## Future: GenericGherkins

A typed generic version of `SimpleGherkins` is proposed to replace untyped `args.Map` patterns:

```go
type GenericGherkins[TInput, TExpect any] struct {
    Title, Feature, Given, When, Then string
    Input      TInput
    Expected   TExpect
    Actual     TExpect
    IsMatching bool
    ExtraArgs  args.Map
}

type AnyGherkins = GenericGherkins[any, any]
```

**Key methods:** `IsFailedToMatch()`, `ShouldBeEqual(t, idx, act, exp)`, `CompareWith(other)`, `String()`, `FullString()`.

See [Testing Guidelines — GenericGherkins Proposal](/spec/01-app/16-testing-guidelines.md#future-genericgherkins-proposal) for full design.

## Contributors

## Issues for Future Reference
