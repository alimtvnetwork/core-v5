# Issue: Missing Parentheses in IsDeleteLineRequest()

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`coreinstruction/LineIdentifier.go`

## Description

`IsDeleteLineRequest()` used `HasLineNumber() && IsDelete() || IsDrop()` without parentheses, so `IsDrop()` alone could make the expression true regardless of `HasLineNumber()`.

## Fix

Added parentheses: `HasLineNumber() && (IsDelete() || IsDrop())`.
