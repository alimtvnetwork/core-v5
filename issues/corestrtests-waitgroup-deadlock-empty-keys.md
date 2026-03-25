# Issue: corestrtests — WaitGroup Deadlock in `AddOrUpdateStringsPtrWgLock` with Empty Keys

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

`Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock_Empty` — times out after 10 minutes with:
```
panic: test timed out after 10m0s
goroutine 3911 [sync.WaitGroup.Wait, 9 minutes]
```

## Stack Trace

```
Coverage36_Hashmap_Hashset_Full_Coverage_test.go:164 — wg.Wait()
```

## Root Cause

In `coredata/corestr/Hashmap.go:157-184`, the `AddOrUpdateStringsPtrWgLock` method:

```go
func (it *Hashmap) AddOrUpdateStringsPtrWgLock(
    wg *sync.WaitGroup,
    keys, values []string,
) *Hashmap {
    // ...
    if len(keys) == 0 {
        return it  // ← EARLY RETURN WITHOUT wg.Done()
    }
    // ...
    wg.Done()  // ← Only called when keys are non-empty
    return it
}
```

When called with empty slices, the method returns early at line 172 **without calling `wg.Done()`**. The test does:

```go
wg.Add(1)
h.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
wg.Wait()  // ← DEADLOCK: wg counter is still 1
```

The `wg.Wait()` blocks forever because the counter was incremented but never decremented.

## Fix

Add `wg.Done()` before the early return for empty keys:

```go
if len(keys) == 0 {
    wg.Done()
    return it
}
```

This follows the convention documented in memory: callers of `AddOrUpdateStringsPtrWgLock` call `wg.Add(1)` before invocation, so the method **must** always call `wg.Done()` regardless of the code path.

## Affected Files

- `coredata/corestr/Hashmap.go:171-173` — **fix location**
- `tests/integratedtests/corestrtests/Coverage36_Hashmap_Hashset_Full_Coverage_test.go:159-165` — test (correct, no change needed)

## Learning

Any method that accepts a `*sync.WaitGroup` and is expected to call `wg.Done()` must do so on **every** code path, including early returns. This is a classic Go concurrency bug pattern.
