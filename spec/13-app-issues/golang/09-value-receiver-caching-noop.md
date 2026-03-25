# Issue: Value Receiver Makes Caching a No-Op

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`coreinstruction/BaseTypeDotFilter.go`

## Description

`GetDotSplitTypes()` used a value receiver, so the cached `splitDotFilters` field was set on a copy and discarded after the call. The cache was recomputed every time.

## Fix

Changed to pointer receiver `*BaseTypeDotFilter`.
