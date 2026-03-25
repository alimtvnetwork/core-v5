# Issue: Specification.Clone() Missing Nil Guard and Shallow-Copying Tags

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## File

`coreinstruction/Specification.go`

## Description

1. `Clone()` had no nil receiver guard — calling it on a nil `*Specification` would panic.
2. The `Tags` slice was assigned directly (`Tags: r.Tags`), meaning the clone shared the same backing array. Mutating tags on the clone would corrupt the original.

## Fix

Added `if r == nil { return nil }` guard and deep-copied `Tags` using `make` + `copy`.
