# Massive Type Duplication Without Generics

## Status: ✅ RESOLVED

## Issue Summary

Multiple packages (`conditional/`, `coremath/`, `core.go`, `isany/`, `issetter/`) had near-identical functions duplicated for every primitive type. This was a direct consequence of Go 1.17 lacking generics.

## Root Cause Analysis

Pre-generics Go required separate implementations per type.

## Fix Description

After upgrading to Go 1.24, generic versions were introduced. See [Go Modernization Plan](/spec/01-app/11-go-modernization.md).

## Prevention and Non-Regression

- After generics adoption, add lint rules to prevent new per-type duplicates.

## Done Checklist

- [x] Upgrade Go version first (prerequisite) — Go 1.24.0
- [x] Add generic `conditional.If[T]` — implemented in `generic.go`
- [x] Add typed convenience wrappers (`IfInt`, `IfBool`, etc.) — `typed_wrappers.go`
- [x] Add generic `EmptySlicePtr[T]` — already implemented in `generic.go`
- [x] Deprecate per-type functions — all legacy files have deprecation comments
- [x] Remove deprecated files — Done (Phase 5: 27 files deleted, ~99 functions removed)
- [x] Tests pass
