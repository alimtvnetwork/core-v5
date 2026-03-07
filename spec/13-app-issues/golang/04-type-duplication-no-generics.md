# Massive Type Duplication Without Generics

## Issue Summary

Multiple packages (`conditional/`, `coremath/`, `core.go`, `isany/`, `issetter/`) have near-identical functions duplicated for every primitive type. This is a direct consequence of Go 1.17 lacking generics.

## Root Cause Analysis

Pre-generics Go required separate implementations per type.

## Fix Description

After upgrading to Go 1.22+, introduce generic versions. See [Go Modernization Plan](/spec/01-app/11-go-modernization.md).

## Prevention and Non-Regression

- After generics adoption, add lint rules to prevent new per-type duplicates.

## TODO and Follow-Ups

- [x] Upgrade Go version first (prerequisite)
- [x] Add generic `conditional.If[T]` — implemented in `generic.go`
- [x] Add typed convenience wrappers (`IfInt`, `IfBool`, etc.) — `typed_wrappers.go`
- [x] Add generic `EmptySlicePtr[T]` — already implemented in `generic.go`
- [x] Deprecate per-type functions — all legacy files have deprecation comments
- [ ] Remove deprecated files in next major version

## Done Checklist

- [x] Generic versions created (`If[T]`, `IfFunc[T]`, `IfTrueFunc[T]`, `IfSlice[T]`, `NilDef[T]`, `ValueOrZero[T]`, etc.)
- [x] Typed convenience wrappers created (`IfInt`, `IfString`, `IfSliceBool`, etc.)
- [x] Old functions deprecated with migration comments
- [ ] Tests pass (run locally to verify)
