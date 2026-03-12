# Build Failure: `args.Map.ShouldBeEqual` Undefined — 36 Test Files Affected

## Status: ✅ RESOLVED

## Issue Summary

36 test files (4197 call sites) use `expected.ShouldBeEqual(t, caseIndex, "title", actual)` on `args.Map`, but this method did not exist on the `Map` type. This caused build failures in every affected test package, blocking coverage measurement.

## Error

```
expected.ShouldBeEqual undefined (type args.Map has no field or method ShouldBeEqual)
```

## Root Cause

The `ShouldBeEqual` method was implemented on `CaseV1` and `SimpleTestCase` but never added to `args.Map`. Test files were written assuming the method existed on `Map` directly.

## Fix

Added `coretests/args/MapShouldBeEqual.go` implementing:

```go
func (it Map) ShouldBeEqual(t *testing.T, caseIndex int, title string, actual Map)
```

This method replicates the same assertion logic as `CaseV1.ShouldBeEqualMap`:
1. Compiles both maps to sorted `"key : value"` lines via `CompileToStrings()`
2. Detects mismatches via `errcore.HasAnyMismatchOnLines`
3. Prints diff and Go literal diagnostics via `errcore.PrintDiffOnMismatch` and `errcore.MapMismatchError`
4. Asserts via goconvey `So(validationErr, should.BeNil)`

## Why Not Rewrite 36 Files

Adding the method to `args.Map` is the correct fix because:
- 4197 call sites across 36 files all use the same `expected.ShouldBeEqual(t, 0, "title", actual)` pattern
- The pattern is ergonomic and self-documenting
- No import cycle exists: `args` can safely import `errcore` and `goconvey`
- Rewriting all 36 files to use `CaseV1{...}.ShouldBeEqualMapFirst(t, actual)` would be a massive, error-prone change

## Affected Files (36)

All `Coverage*.go` and `Coverage2*.go` test files across `tests/integratedtests/` that use the `args.Map` assertion pattern.

## What Not to Repeat

- When designing assertion APIs, implement the method on the type that callers will use directly
- Before committing test files en masse, verify that called methods actually compile
