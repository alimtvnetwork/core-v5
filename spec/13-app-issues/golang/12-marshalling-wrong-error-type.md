# Issue: MarshallingFailedDueToNilOrEmpty Uses Wrong Error Type

## Status: ✅ RESOLVED

## Phase: 7.4 — Minor Cleanup

## File

`defaulterr/defaulterr.go`

## Description

`MarshallingFailedDueToNilOrEmpty` was constructed with `UnMarshallingFailedType` instead of `MarshallingFailedType`, causing marshalling errors to be misclassified as unmarshalling errors.

## Fix

Changed to `MarshallingFailedType`.
