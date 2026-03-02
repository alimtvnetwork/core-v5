# Issue: PointerStrings/PointerIntegers Less() Nil Element Dereference

## Status: ✅ RESOLVED

## Phase: 8.1 — Critical Bugs

## Files

- `coredata/PointerStrings.go`
- `coredata/PointerStringsDsc.go`
- `coredata/PointerIntegers.go`
- `coredata/PointerIntegersDsc.go`

## Description

All `Less()` methods dereference pointer elements (`*p[i]`) without nil checks. If any element in the slice is nil, calling `sort.Sort()` will panic with a nil pointer dereference.

## Fix

Add nil checks in `Less()`. Treat nil as less-than any non-nil value (or greater-than for Dsc variants) to ensure stable, panic-free sorting.
