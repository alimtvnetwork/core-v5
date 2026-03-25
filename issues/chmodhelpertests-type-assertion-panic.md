# Issue: chmodhelpertests — Type Assertion Panic in `assertNonWhiteSortedEqual`

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

`Test_VerifyRwxPartialChmodLocations_Unix` — panics with:
```
panic: interface conversion: interface {} is string, not []string
```

## Stack Trace

```
nonWhiteSortedEqual.go:41 — testCase.ExpectedInput.([]string)
VerifyPartialRwxLocations_test.go:33 — assertNonWhiteSortedEqual(t, testCase, caseIndex, err)
```

## Root Cause

`nonWhiteSortedEqual.go:41` performs a hard type assertion:
```go
expectedLines := testCase.ExpectedInput.([]string)
```

But test case index 1 in `VerifyPartialRwxLocations_testcases.go` has:
```go
ExpectedInput: "",  // string, not []string
```

This is the "no error expected" case where the expected output is an empty string. The type assertion `.([]string)` panics because `ExpectedInput` is a `string`, not `[]string`.

The other 2 test cases correctly use `[]string{...}` for `ExpectedInput`, so they would pass — but Go runs test cases sequentially in the loop, and the panic at index 1 crashes the entire test function.

## Fix

In `nonWhiteSortedEqual.go:41`, replace the hard assertion with a type switch that handles both `string` and `[]string`:

```go
var expectedStr string
switch v := testCase.ExpectedInput.(type) {
case []string:
    expectedStr = strings.Join(v, "\n")
case string:
    expectedStr = v
default:
    expectedStr = fmt.Sprintf("%v", testCase.ExpectedInput)
}
```

## Affected Files

- `tests/integratedtests/chmodhelpertests/nonWhiteSortedEqual.go` — **fix location**
- `tests/integratedtests/chmodhelpertests/VerifyPartialRwxLocations_testcases.go` — test data (correct, no change needed)

## Learning

Always use type switches or comma-ok assertions (`v, ok := x.(T)`) when `ExpectedInput` is `interface{}` — different test cases may legitimately use different concrete types.
