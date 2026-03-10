# Issue: Map Diagnostic Output Format

## Summary

When `ShouldBeEqualMap` fails, the diagnostic output must use the project's
standard header separators (`============================>`) and show each entry
line-by-line with aligned `actual` / `expected` labels — not as two separate
blocks of entries.

## Symptom (before fix)

```
Actual Received (2 entries):
	"containsName": false,
	"hasError":     false,

Expected Input (1 entries):
	"hasError": false,
```

Entries were shown as two separate blocks with trailing commas, no header
separators, and no per-line actual/expected alignment.

## Root Cause

`MapMismatchError` was formatting actual and expected as two separate blocks
with indexed Go literal lines and trailing commas. This:
1. Lost the standard `============================>` header separators
2. Did not show per-line comparison with aligned actual/expected labels
3. Added unnecessary trailing commas and index numbers

## Fix

Rewrote `errcore/MapMismatchError.go` to:
1. Use standard `============================>` header separators
2. Show each line with aligned `actual   :` / `expected :` labels
3. No trailing commas — uses `CompileToStrings()` format (`key : value`)
4. Mark missing lines as `<missing>`

`CaseV1MapAssertions.go` passes `CompileToStrings()` lines (not GoLiteralLines)
to `MapMismatchError` for the diagnostic, and separately prints a `slog.Warn`
with Go literal format for copy-paste convenience.

## Correct Output Format

```
============================>
0 ) Map Mismatch:
    title here
============================>
    Actual lines: 2, Expected lines: 1
============================>
	actual   : containsName : false
	expected : hasError : false
============================>
	actual   : hasError : false
	expected : <missing>
============================>
```

## Key Rules

- Always use `============================>` header separators.
- Show per-line comparison with aligned `actual   :` / `expected :` labels.
- No trailing commas on entry values.
- No indexed numbering (`0:`, `1:`, etc.) before entries.
- `actual   :` and `expected :` colons must start at the same column.

## Affected Files

- `errcore/MapMismatchError.go` — map-specific diagnostic formatter
- `coretests/coretestcases/CaseV1MapAssertions.go` — assertion entry point

## Spec Reference

`spec/testing-guidelines/06-diagnostics-output-standards.md` → Map Expected Output
