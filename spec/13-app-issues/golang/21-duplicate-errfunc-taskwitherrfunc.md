# Issue: Duplicate ErrFunc / TaskWithErrFunc Types

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`errcore/funcs.go`

## Description

`ErrFunc` and `TaskWithErrFunc` were both defined as `func() error` — identical duplicate types with no semantic distinction.

## Fix

Converted `TaskWithErrFunc` to a type alias (`= ErrFunc`) with a deprecation comment, preserving backward compatibility while eliminating the redundancy.
