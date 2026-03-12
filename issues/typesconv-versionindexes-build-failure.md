# Build Failure: `args.Map.ShouldBeEqual` Undefined in typesconvtests and versionindexestests

## Status: ✅ RESOLVED

## Issue Summary

Two test packages failed to build due to calling `ShouldBeEqual` on `args.Map`, which does not have that method:

- `tests/integratedtests/typesconvtests/Coverage2_test.go` — 22 occurrences
- `tests/integratedtests/versionindexestests/Coverage_test.go` — 6 occurrences

## Error

```
expected.ShouldBeEqual undefined (type args.Map has no field or method ShouldBeEqual)
```

## Root Cause

Tests were written using an incorrect API: `expected.ShouldBeEqual(t, 0, "title", actual)` called directly on `args.Map`. The `ShouldBeEqual` method belongs to `CaseV1` and `SimpleTestCase`, not `args.Map`.

## Fix

Replaced all `expected.ShouldBeEqual(t, 0, "title", actual)` calls with:

```go
coretestcases.CaseV1{Title: "title", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
```

This uses the correct `CaseV1.ShouldBeEqualMapFirst` method which accepts an `args.Map` for comparison.

## Impact

These build failures triggered the **build failure cascade** (see `spec/05-failing-tests/12-zero-coverage-build-failure-cascade.md`), potentially blocking coverage measurement for all test packages.

## What Not to Repeat

- Always verify that assertion methods exist on the type being used before committing tests.
- `args.Map` is a data container — assertion methods live on `CaseV1`, `GenericGherkins`, or `SimpleTestCase`.
