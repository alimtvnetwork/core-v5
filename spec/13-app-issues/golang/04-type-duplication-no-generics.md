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

- [ ] Upgrade Go version first (prerequisite)
- [ ] Add generic `conditional.If[T]`
- [ ] Add generic `EmptySlicePtr[T]`
- [ ] Deprecate per-type functions
- [ ] Remove in next major version

## Done Checklist

- [ ] Generic versions created
- [ ] Old functions deprecated
- [ ] Tests pass
