# Issue: LineDiff actual/expected Label Alignment Insufficient

## Summary

In `LineDiff` mismatch output, the `actual` and `expected` labels had inconsistent
visual padding, making them harder to scan in terminal output.

## Symptom

```
  Line   0 [MISMATCH]:
          actual   : `containsName : false`
          expected : `hasError : false`
```

Although the colons were technically aligned (both at column 19), the visual difference
between `actual   ` (6 chars + 3 spaces) and `expected ` (8 chars + 1 space) made
the labels appear misaligned in some terminal environments.

## Root Cause

In `errcore/LineDiff.go`, the format strings used minimal padding to align the colon:

```go
"          actual   : `%s`\n"+
"          expected : `%s`\n",
```

Both labels totaled 9 characters before the colon, but the uneven trailing space
distribution (`actual` had 3 trailing spaces, `expected` had 1) created a visual
imbalance, especially in terminals using proportional or variable-width rendering.

## Fix

Adjusted leading spaces so both colons align at the same column:

```go
"              actual : `%s`\n"+
"            expected : `%s`\n",
```

- `actual`: 14 leading spaces + 6 chars + 1 space = colon at column 21
- `expected`: 12 leading spaces + 8 chars + 1 space = colon at column 21

Both colons now sit at column 21 from the line start.

## Affected Files

- `errcore/LineDiff.go` — `LineDiffToString` format strings (lines 100-102)

## Spec Reference

`spec/testing-guidelines/06-diagnostics-output-standards.md` → Multi-line comparison alignment
