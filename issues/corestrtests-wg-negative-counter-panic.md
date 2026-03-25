# Issue: corestrtests — Negative WaitGroup Counter Panic in Test_Cov18

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

`Test_Cov18_Hashmap_StringsPtrWgLock` — panics with:
```
panic: sync: negative WaitGroup counter
```

## Root Cause

`Coverage18_HashmapHashset_test.go:344` passes a fresh `&sync.WaitGroup{}` (counter=0) with nil keys:
```go
h.AddOrUpdateStringsPtrWgLock(&sync.WaitGroup{}, nil, nil)
```

The `AddOrUpdateStringsPtrWgLock` method contract requires callers to call `wg.Add(1)` before invocation, because the method always calls `wg.Done()`. When the WaitGroup counter is 0, calling `wg.Done()` panics with "negative WaitGroup counter".

This bug was exposed after fixing the deadlock in `Hashmap.go` (see `issues/corestrtests-waitgroup-deadlock-empty-keys.md`) — the added `wg.Done()` on the empty-keys early return path now executes on this fresh WaitGroup.

## Fix

Add `wg.Add(1)` before the call:
```go
wg2 := &sync.WaitGroup{}
wg2.Add(1)
h.AddOrUpdateStringsPtrWgLock(wg2, nil, nil)
```

## Affected Files

- `tests/integratedtests/corestrtests/Coverage18_HashmapHashset_test.go:344` — **fix location**

## Learning

Any method that calls `wg.Done()` requires its caller to have called `wg.Add(1)` first. When testing edge cases (nil/empty input), ensure the WaitGroup contract is still honored.
