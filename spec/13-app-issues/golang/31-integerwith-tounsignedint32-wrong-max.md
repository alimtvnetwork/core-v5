# Issue: integerWithin.ToUnsignedInt32() Wrong Max Constant

## Status: ✅ RESOLVED

## Phase: 8.1 — Critical Bugs

## File

`coremath/integerWithin.go`

## Description

`ToUnsignedInt32()` uses `math.MaxInt32` instead of `math.MaxUint32`. This mirrors the same class of bug previously fixed in `integerOutOfRange`. The function will reject valid values in the range `[2147483648, 4294967295]`.

## Fix

Change `math.MaxInt32` to `math.MaxUint32`.
