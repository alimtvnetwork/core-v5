# Issue: coreonce Types Not Thread-Safe

## Status: ✅ RESOLVED

## Phase: 9.2 — Code Quality

## Files

- `coredata/coreonce/StringOnce.go`
- `coredata/coreonce/IntegerOnce.go`
- `coredata/coreonce/BoolOnce.go`

## Description

`StringOnce`, `IntegerOnce`, and `BoolOnce` all use a manual `isInitialized` flag to implement lazy initialization, but without any synchronization. Concurrent calls to `Value()` can race — multiple goroutines may execute `initializerFunc()` simultaneously, and the write to `isInitialized`/`innerData` is not atomic.

## Fix

Replaced the manual `isInitialized` flag with `sync.Once` to guarantee the initializer runs exactly once, even under concurrent access.
