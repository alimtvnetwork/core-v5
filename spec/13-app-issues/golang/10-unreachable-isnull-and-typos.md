# Issue: Unreachable IsNull() Checks and SetFromName Typo

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## Files

- `coreinstruction/SourceDestination.go`
- `coreinstruction/FromTo.go`
- `coreinstruction/Rename.go`

## Description

1. `String()` methods (value receivers) contained `if it.IsNull()` guards that are unreachable — a value receiver can never be nil.
2. `SetFromName` parameter was named `form` instead of `from` — a typo.

## Fix

Removed unreachable nil checks and renamed parameter to `from`.
