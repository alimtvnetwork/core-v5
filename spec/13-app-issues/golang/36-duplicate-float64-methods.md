# Issue: stringTo.Float64Default and Float64Conditional Are Identical

## Status: 🟡 OPEN

## Phase: 8.2 — Code Quality

## File

`converters/stringTo.go`

## Description

`Float64Default()` (line 161) and `Float64Conditional()` (line 173) have identical signatures and implementations. This is likely a copy-paste artifact. One should be removed or differentiated.

## Recommendation

Remove the duplicate or differentiate their behavior/signature.
