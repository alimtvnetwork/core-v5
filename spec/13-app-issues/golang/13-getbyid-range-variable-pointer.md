# Issue: GetById Returns Pointer to Range-Loop Copy

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## Files

- `coreinstruction/Identifiers.go`
- `coreinstruction/IdentifiersWithGlobals.go`

## Description

`GetById()` in both `Identifiers` and `IdentifiersWithGlobals` used a `for _, v := range` loop and returned `&v`. This returns a pointer to the loop variable (a copy), not to the actual slice element. The returned pointer would be invalidated or point to the wrong element on the next iteration.

## Fix

Changed to index-based iteration (`for i := range`) and return `&slice[i]` to point directly at the backing array element.
