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
Go literal format.

## Root Cause

`ShouldBeEqualMap` delegated to the generic `ShouldBe` pipeline, which uses
`SliceValidator` → `SliceValidatorMessages` → `StringLinesToQuoteLinesToSingle`
for formatting. This wraps each compiled line in double quotes with commas,
which is designed for generic string slice comparison but is unsuitable for
map diagnostics where copy-pasteability matters.

## Fix

1. Created `errcore/MapMismatchError.go` — formats map mismatches with indexed
   Go literal lines, showing each entry on its own numbered line:
   ```
   Actual Received (2 entries):
     0: "containsName": false,
     1: "hasError":      false,

   Expected Input (1 entries):
     0: "hasError": false,
   ```

2. Modified `CaseV1MapAssertions.go` — `ShouldBeEqualMap` now handles the full
   assertion directly instead of delegating to generic `ShouldBe`. On mismatch:
   - Prints `LineDiff` for detailed line-by-line comparison
   - Builds error message using `MapMismatchError` with Go literal format
   - Asserts via `convey.So(validationErr, should.BeNil)`

## Affected Files

- `errcore/MapMismatchError.go` — new, map-specific diagnostic formatter
- `coretests/coretestcases/CaseV1MapAssertions.go` — primary fix location

## Spec Reference

`spec/testing-guidelines/06-diagnostics-output-standards.md` → Map Expected Output
