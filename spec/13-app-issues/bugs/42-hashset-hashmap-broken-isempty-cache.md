# Hashset/Hashmap IsEmpty Caching Logic is Broken

## Status: ✅ RESOLVED

## Fix Applied

Removed all broken caching fields (`hasMapUpdated`, `isEmptySet`, `length`) from both `coregeneric.Hashset` and `coregeneric.Hashmap`.

- `IsEmpty()` now simply returns `it == nil || len(it.items) == 0` — O(1) in Go
- Constructors no longer set misleading `length: capacity`
- All mutation methods no longer set `hasMapUpdated` flag

**Note:** The `corestr.Hashset` and `corestr.Hashmap` string-specific types still use this pattern and should be cleaned up separately in a future pass.
