# Hashset/Hashmap IsEmpty Caching Logic is Broken

## Issue Summary

`Hashset.IsEmpty()` and `Hashmap.IsEmpty()` use a `hasMapUpdated` flag + `isEmptySet` cache for optimization. However, `hasMapUpdated` is never reset to `false` after recalculation, so every `IsEmpty()` call after any mutation recalculates `len(items)` — making the cache useless.

Additionally, `NewHashset`/`NewHashmap` set `length` to `capacity` (not 0), which is misleading even though `Length()` uses `len(items)` at runtime.

## Fix Description

Either:
1. Remove the caching entirely — just use `return it == nil || len(it.items) == 0`
2. Or properly implement: reset `hasMapUpdated = false` after recalculation, and set `length: 0` in constructors

Option 1 is preferred — `len()` on a map is O(1) in Go, so caching provides no benefit.

## Done Checklist

- [ ] Simplify `IsEmpty()` to remove broken caching
- [ ] Fix `length` initialization in `NewHashset`/`NewHashmap`
- [ ] Remove `hasMapUpdated`, `isEmptySet`, `length` fields if unused elsewhere
