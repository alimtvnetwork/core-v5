# Issue: LineIdentifier.IsNewLineRequest() Missing Nil Guard

## Status: 🔴 OPEN

## Phase: 8.1 — Critical Bugs

## File

`coreinstruction/LineIdentifier.go`

## Description

`IsNewLineRequest()`, `IsDeleteLineRequest()`, and `IsModifyLineRequest()` access `it.LineModifyAs` without first checking `it == nil`. Other methods on the same receiver (e.g., `HasLineNumber`) include nil guards. Calling any of these three methods on a nil `*LineIdentifier` will panic.

## Fix

Add `it == nil` guard returning `false` at the start of each method.
