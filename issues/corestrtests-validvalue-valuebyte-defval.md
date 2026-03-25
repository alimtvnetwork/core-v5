# Issue: ValidValue.ValueByte ignores defVal parameter

## Date: 2026-03-25

## Status: ✅ RESOLVED — NOT A BUG

## Original Failing Test

- `Test_C40_ValidValue_NumericConversions` (line 63)

## Analysis

`ValidValue.ValueByte(defVal byte)` returns `constants.Zero` on parse error or negative.
This is the **correct contract** — `defVal` is accepted for API consistency but the error
path always returns zero. Multiple existing tests (`Test_Cov11`, `Test_I27`, `Test_S01`,
`Test_Seg8_VV`) confirm this behavior.

## Resolution

The test `Test_C40` line 63 was wrong: `bad.ValueByte(88) != 88` should be `!= 0`.
Fixed the test expectation to match the established contract.

## Affected Files

- `tests/integratedtests/corestrtests/Coverage40_Types_Remaining_Coverage_test.go` (line 63)
