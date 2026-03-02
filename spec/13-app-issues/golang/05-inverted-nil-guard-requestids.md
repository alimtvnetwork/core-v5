# Issue: Inverted Nil Guard in BaseRequestIds.RequestIdsLength()

## Status: ✅ RESOLVED

## Phase: 7.1 — Critical Bugs

## File

`coreinstruction/BaseRequestIds.go`

## Description

The nil guard in `RequestIdsLength()` was inverted — it checked `b != nil` instead of `b == nil`, causing a nil pointer panic when the receiver was nil.

## Fix

Changed condition to `if b == nil || b.RequestIds == nil` to correctly short-circuit and return `constants.Zero`.
