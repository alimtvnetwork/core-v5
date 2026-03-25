# Issue: Hashset.Transpile mutates map during iteration

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Tests

- `Test_Seg5_HS_Transpile` (expected len 1, got 4)
- `Test_Seg5_HS_WrapDoubleQuote` (expected len 1, got 2)
- `Test_Seg5_HS_WrapSingleQuote` (expected len 1, got 5)
- `Test_Seg5_HS_WrapDoubleQuoteIfMissing` (expected len 1, got 2)
- `Test_Seg5_HS_WrapSingleQuoteIfMissing` (expected len 1, got 2)

## Root Cause

`Hashset.Transpile()` in `coredata/corestr/Hashset.go` was modifying the
`items` map **in-place during iteration**:

```go
for k, v := range it.items {
    it.items[fmtFunc(k)] = v  // adds new key but never deletes old key
}
```

This caused:
1. Original key `"a"` remains in the map
2. Transformed key `"\"a\""` is added
3. In some cases, the transformed key gets re-iterated and transformed again
4. Result: map grows instead of staying the same size

All `Wrap*` methods delegate to `Transpile`, so they all exhibited the same bug.

## Solution

Build a **new map** from the transformed keys, then replace `it.items`:

```go
newItems := make(map[string]bool, len(it.items))
for k, v := range it.items {
    newItems[fmtFunc(k)] = v
}
it.items = newItems
```

## Affected Files

- `coredata/corestr/Hashset.go` (line ~1453) — production fix
- `tests/integratedtests/corestrtests/Coverage40_Seg5_Hashset_test.go` — tests now pass
