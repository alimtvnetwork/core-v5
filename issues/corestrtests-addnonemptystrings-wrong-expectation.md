# Issue: corestrtests — Test_Seg3_Collection_AddNonEmptyStrings wrong expected length

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

`Test_Seg3_Collection_AddNonEmptyStrings` — assertion failure: expected length 3, got 2.

## Root Cause

The test at `Coverage37_Seg3_CollectionEnd_test.go:501` expects `len: 3` after calling:
```go
c.AddNonEmptyStrings("a", "", "b")
```

But `AddNonEmptyStrings` delegates to `AddNonEmptyStringsSlice` (Collection.go:1791),
which explicitly filters out empty strings:
```go
for _, addingItem := range slice {
    if addingItem == "" {
        continue
    }
    items = append(items, addingItem)
}
```

So `"a", "", "b"` → only `"a"` and `"b"` are added → length is **2**, not 3.

The comment on line 501 was also misleading:
```go
expected := args.Map{"len": 3} // AddNonEmptyStrings delegates to AddNonEmptyStringsSlice which adds all
```

"which adds all" is wrong — it adds all **non-empty** strings.

## Fix

Change expected length from `3` to `2` and fix the comment.

## Affected Files

- `tests/integratedtests/corestrtests/Coverage37_Seg3_CollectionEnd_test.go:501` — **fix location**
- `coredata/corestr/Collection.go:1815-1835` — production code (correct, no change needed)
