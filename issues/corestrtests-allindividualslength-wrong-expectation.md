# Issue: AllIndividualsLength tests expect character count instead of item count

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Tests

- `Test_I8_AllIndividualStringsOfStringsLength` (expected 6, got 3)
- `Test_I8_AllIndividualsLengthOfSimpleSlices` (expected 6, got 3)

## Root Cause

Both tests used items with varying string lengths (`"a"`, `"bb"`, `"ccc"`) and
expected the sum of **character lengths** (1+2+3=6). However, both production
functions count **items**, not characters:

- `AllIndividualStringsOfStringsLength`: `length += len(stringsItems)` — counts slice length
- `AllIndividualsLengthOfSimpleSlices`: `length += simpleSlice.Length()` — counts items

This is confirmed by existing tests in `Coverage01_Helpers_test.go` and
`Coverage10_Full_test.go` which correctly expect item counts (e.g., 3 for
`{"a","b"}, {"c"}`).

## Solution

Fix test expectations from 6 to 3.

## Affected Files

- `tests/integratedtests/corestrtests/Coverage41_Iteration8_test.go` (lines 1217, 1224) — test fix
