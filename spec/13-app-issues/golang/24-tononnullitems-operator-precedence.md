# Issue: anyItemConverter.ToNonNullItems() Operator Precedence Bug

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## File

`converters/anyItemConverter.go`

## Description

The condition `isSkipOnNil && anyItem == nil || reflectinternal.Is.Null(anyItem)` evaluated as `(isSkipOnNil && anyItem == nil) || reflectinternal.Is.Null(anyItem)` due to Go operator precedence. This meant the reflection null check always ran regardless of the `isSkipOnNil` flag.

## Fix

Added explicit parentheses: `isSkipOnNil && (anyItem == nil || reflectinternal.Is.Null(anyItem))`.
