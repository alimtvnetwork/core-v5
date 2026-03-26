# Dead Code Analysis: Unreachable Coverage Lines

## Status: DOCUMENTED

These lines are logically unreachable and cannot be covered by any test.

## 1. `coreversion/hasDeductUsingNilNess.go:20-22`

```go
if left == nil || right == nil {
    return corecomparator.NotEqual, true
}
```

**Reason**: Prior branches exhaustively handle all nil combinations:
- L8: `left == nil && right == nil`
- L12: `left != nil && right == nil`
- L16: `left == nil && right != nil`

After these, both `left` and `right` are guaranteed non-nil. The `||` check is dead.

## 2. `coredata/corerange/MinMaxByte.go:46-48`

```go
if diff < 0 {
    return diff
}
```

**Reason**: `byte` is an unsigned type in Go (`uint8`). The expression `diff < 0` is always false for unsigned types.

## 3. `coredata/corerange/within.go:89`

```go
return 0, isInRange
```

**Reason**: `StringRangeUint32` calls `StringRangeInteger(true, 0, MaxInt32, input)`. With `isUsageMinMaxBoundary=true`, `RangeInteger` clamps the result to `[0, MaxInt32]`. Therefore `finalInt <= math.MaxInt32` is always true, making the else branch unreachable.

## 4. `coretaskinfo/InfoJson.go:25-27`

```go
func (it Info) JsonString() string {
    if it.IsNull() {  // IsNull checks it == nil, but value receiver is never nil
        return ""
    }
```

**Reason**: `JsonString` uses a value receiver (`Info`), so `it` is a copy on the stack. `IsNull()` is a pointer method checking `it == nil`, but `&it` of a value receiver is never nil.

## 5. `coredata/stringslice/MergeSlicesOfSlices.go:13-15`

```go
if sliceLength == constants.Zero {
    return []string{}
}
```

**Reason**: Redundant check. Line 7 already returns `[]string{}` when `len(slicesOfSlice) == 0`, and `sliceLength` is assigned from the same `len()` call.

## 6. `coredata/stringslice/RegexTrimmedSplitNonEmptyAll.go:17-19`

```go
if len(items) == 0 {
    return []string{}
}
```

**Reason**: `regexp.Split(content, -1)` always returns at least one element (even for empty string → `[""]`). The `len(items) == 0` check is unreachable.

## Recommendation

These are defensive guards. They could be removed for cleanliness, but keeping them is harmless. Coverage for these packages is effectively 100% for all reachable code.
