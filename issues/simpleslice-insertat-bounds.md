# Issue: SimpleSlice.InsertAt Slice Bounds Panic (FIXED)

## Summary
`InsertAt()` panicked with `slice bounds out of range [:4] with capacity 3` when inserting at the end of a 3-element slice.

## Root Cause
```go
// BEFORE (broken):
s = append(s[:index+1], s[index:]...)
```

For a 3-element slice inserting at index 2: `s[:3]` works, but `s[:index+1]` where `index == length-1` requires `index+1 == length` which is fine. However, for appending at `index == length` (valid per the bounds check), `s[:index+1]` = `s[:length+1]` exceeds capacity.

## Fix
```go
s = append(s, "")            // grow by one
copy(s[index+1:], s[index:]) // shift right
s[index] = item              // insert
```

This is the standard Go insert-at-index pattern that avoids bounds issues.

## Affected Files
- `coredata/corestr/SimpleSlice.go` — line 198-201

## Learning
- The `append(s[:i+1], s[i:]...)` pattern only works when `i < len(s)`, not for `i == len(s)`
- Use the grow-copy-assign pattern for robust insertions
