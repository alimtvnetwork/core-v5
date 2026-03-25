# Issue: ValidValue.ValueByte Should Return constants.Zero on Error

## Status: FIXED (Reverted)

## Affected Tests
- Test_Cov11_ValidValue_Conversions (badByte: expected 0, got 55)
- Test_I27_ValidValue_ValueByte (err: expected 0, got 7; neg: expected 0, got 9)
- Test_S01_ValidValue_ValueByte_Invalid
- Test_S01_ValidValue_ValueByte_Negative
- Test_Seg8_VV_ValueByte_Negative

## Root Cause
A previous fix incorrectly changed `ValueByte` to return `defVal` on parse error 
or negative values, when the original (correct) behavior was to return `constants.Zero`.

The `defVal` parameter is present in the signature for API consistency with 
`ValueInt(defaultInteger int)`, but the error-path behavior is intentionally different:
- `ValueByte` returns `0` on error (consistent with `ValueDefByte`)
- `defVal` is unused in error paths by design

## Solution
Reverted `ValueByte` to return `constants.Zero` on error/negative, matching the 
original production behavior that all tests expect.

## Files Modified
- `coredata/corestr/ValidValue.go` — reverted error return to `constants.Zero`
