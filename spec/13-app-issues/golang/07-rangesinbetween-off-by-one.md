# Issue: RangesInBetween Off-By-One Error

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## File

`reqtype/RangesInBetween.go`

## Description

The loop condition `i < endVal` excluded the last element from the generated range, contradicting the documented inclusive `[start, end]` semantics used by `IsInBetween` and `GetInBetweenStatus`.

## Fix

Changed loop condition to `i <= endVal` to include the end value.
