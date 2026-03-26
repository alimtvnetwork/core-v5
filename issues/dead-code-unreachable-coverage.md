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

## Recommendation

These are defensive guards. They could be removed for cleanliness, but keeping them is harmless. Coverage for these packages is effectively 100% for all reachable code.
