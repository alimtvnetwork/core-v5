# Issue: Map Diagnostic Output Uses Double-Quoted String Format Instead of Go Literal

## Summary

When `ShouldBeEqualMap` fails, the diagnostic output shows map entries wrapped
in double quotes with commas (`"key : value",`) via `StringLinesToQuoteLinesToSingle`.
This format is not copy-pasteable into `_testcases.go` and makes entries harder
to read as distinct items.

## Symptom

```
============================>
0 ) Actual Received:
    AsActionReturnsErrorFunc returns nil on success
============================>
"containsName : false",
"hasError : false",
============================>
```

Entries are shown in opaque `"key : value",` format instead of per-item
Go literal format, and header separators (`============================>`)
were removed during initial fix.

## Root Cause

`ShouldBeEqualMap` delegated to the generic `ShouldBe` pipeline, which uses
`SliceValidator` → `SliceValidatorMessages` → `StringLinesToQuoteLinesToSingle`
for formatting. This wraps each compiled line in double quotes with commas,
which is designed for generic string slice comparison but is unsuitable for
map diagnostics where copy-pasteability matters.

During the initial fix, the separator headers (`============================>`)
were removed entirely, breaking the visual structure that distinguishes the
Actual and Expected blocks.

## Fix

1. Created `errcore/MapMismatchError.go` — formats map mismatches with
   tab-indented Go literal lines, each entry on its own line, wrapped in
   separator headers (`============================>`):
   ```
   ============================>
   Actual Received (2 entries):
       Case Title
   ============================>
   	"containsName": false,
   	"hasError":      false,
   ============================>

   ============================>
   Expected Input (1 entries):
       Case Title
   ============================>
   	"hasError": false,
   ============================>
   ```

2. Modified `CaseV1MapAssertions.go` — `ShouldBeEqualMap` now handles the full
   assertion directly instead of delegating to generic `ShouldBe`. On mismatch:
   - Prints `LineDiff` for detailed line-by-line comparison
   - Builds error message using `MapMismatchError` with Go literal format
   - Asserts via `convey.So(validationErr, should.BeNil)`

## Key Rules

- Do NOT use indexed numbering (`0:`, `1:`, etc.) in diagnostic output lines.
- Use tab indentation for each entry line.
- Each entry must be on its own line in Go literal format (`"key": value,`).
- Separator headers (`============================>`) MUST wrap each block.
- Title MUST appear under the section label, indented with 4 spaces.

## Affected Files

- `errcore/MapMismatchError.go` — map-specific diagnostic formatter
- `coretests/coretestcases/CaseV1MapAssertions.go` — primary fix location

## Spec Reference

`spec/testing-guidelines/06-diagnostics-output-standards.md` → Map Expected Output
