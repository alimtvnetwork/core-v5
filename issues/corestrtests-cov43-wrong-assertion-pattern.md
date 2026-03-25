# Issue: Coverage43 Tests Using Wrong Assertion Pattern

## Status: FIXED

## Affected Tests
All ~100 tests in `tests/integratedtests/corestrtests/Coverage43_Iteration39_test.go`

## Root Cause
Every test in the Cov43 file used `CaseV1.ShouldBeEqual(t, 0)` without passing 
actual values via the variadic `...string` parameter.

The `ShouldBeEqual` method signature is:
```go
func (it CaseV1) ShouldBeEqual(t *testing.T, caseIndex int, actualElements ...string)
```

It compares `ExpectedInput` (from the struct) with the variadic `actualElements` parameter.
It does NOT use the `ActualInput` struct field.

The tests were setting `ActualInput` in the struct but not passing it to `ShouldBeEqual`,
resulting in comparisons against an empty/nil actual value.

## Error Pattern
```
ActualLines, ExpectedLines any is nil and other is not. - Expect [""] != ["0"] Actual
```

## Solution
Converted all tests to use `ShouldBeEqualMapFirst(t, args.Map{...})` pattern which
properly passes actual values through the `args.Map` parameter:

**Before (broken):**
```go
tc := coretestcases.CaseV1{Title: "X", ExpectedInput: 0, ActualInput: hm.Length()}
tc.ShouldBeEqual(t, 0)
```

**After (fixed):**
```go
tc := coretestcases.CaseV1{Title: "X", ExpectedInput: args.Map{"len": 0}}
tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
```

## Files Modified
- `tests/integratedtests/corestrtests/Coverage43_Iteration39_test.go` — full rewrite
