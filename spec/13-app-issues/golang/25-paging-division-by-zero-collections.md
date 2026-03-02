# Issue: PayloadsCollection and TypedPayloadCollection GetPagesSize Division by Zero

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## Files

- `coredata/corepayload/PayloadsCollectionPaging.go`
- `coredata/corepayload/typed_collection_paging.go`

## Description

Both `GetPagesSize()` methods divided by `eachPageSize` without checking for zero, causing a panic when called with `eachPageSize == 0`. This is the same class of bug previously fixed in `pagingutil.GetPagesSize()`.

## Fix

Added an early return of `0` when `eachPageSize <= 0` to both methods.
