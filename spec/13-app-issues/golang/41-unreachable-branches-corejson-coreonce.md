# Issue: Unreachable Code Branches in corecmp, corejson, and coreonce

## Status: ✅ RESOLVED

## Phase: 9.2 — Code Quality / Coverage

## Packages

- `corecmp`
- `coredata/coreonce`
- `coredata/corejson`

---

## corecmp

### Dead `return corecomparator.NotEqual` in 9 comparison functions

**Files:** `Byte.go`, `Integer.go`, `Integer8.go`, `Integer16.go`, `Integer32.go`, `Integer64.go`, `Time.go`, `VersionSliceByte.go`, `VersionSliceInteger.go` — all at line 14 or 41.

Each function has an exhaustive `if equal / else if less / else if greater` chain. The final `return corecomparator.NotEqual` is unreachable because the three conditions are logically exhaustive for comparable types.

**9 uncovered statements total. Cannot be covered by any test input.**

**Fix option:** Remove the dead return or convert `else if greater` to bare `else`.

---

## coredata/coreonce

### 1. `StringOnce.SplitLeftRight` — dead `len(items) > 2` branch

**File:** `StringOnce.go`, lines 108–110

```go
items := strings.SplitN(it.Value(), splitter, constants.Two) // max 2 elements

if len(items) == 2 { return items[0], items[1] }
if len(items) > 2 { return items[0], items[len(items)-1] } // UNREACHABLE
```

**Reason:** `strings.SplitN(..., 2)` returns at most 2 elements. The `> 2` branch can never execute.

**Fix options:**
- A) Remove the dead branch (it's protective but misleading).
- B) Change `SplitN(..., 2)` to `Split(...)` if the intent was to handle multi-split and return first/last.

---

## coredata/corejson

### 2. `BytesCollection.Clone` — loop body unreachable due to wrong length check

**File:** `BytesCollection.go`, lines 760–773

```go
func (it BytesCollection) Clone(isDeepCloneEach bool) BytesCollection {
    newResults := NewBytesCollection.UsingCap(it.Length())

    if newResults.Length() == 0 { // BUG: checks NEW (always empty) instead of SOURCE
        return *newResults
    }

    for _, item := range it.Items { // UNREACHABLE
        newResults.Add(BytesCloneIf(isDeepCloneEach, item))
    }
    return *newResults
}
```

**Reason:** `NewBytesCollection.UsingCap(n)` creates an empty collection with capacity `n` but length 0. The check `newResults.Length() == 0` is always true, so the method always returns an empty clone.

**Fix:** Change `newResults.Length() == 0` to `it.Length() == 0`.

### 3. `BytesCollection.ClonePtr` — same bug as Clone

**File:** `BytesCollection.go`, lines 775–792

Same pattern: `newResults.Length() == 0` checks the new empty collection instead of `it.Length() == 0`.

**Fix:** Change `newResults.Length() == 0` to `it.Length() == 0`.

---

## Summary Table

| # | Package | File | Branch | Root Cause | Fixable? |
|---|---------|------|--------|------------|----------|
| 1 | corecmp | 9 files (Byte, Integer*, Time, VersionSlice*) | `return NotEqual` after exhaustive if/else | All 3 comparison outcomes already covered | Yes — use bare `else` |
| 2 | coreonce | `StringOnce.go:108` | `len(items) > 2` | `SplitN(..., 2)` caps at 2 | Yes — remove branch or change to `Split` |
| 3 | corejson | `BytesCollection.go:764` | `Clone` loop body | Checks new collection length (always 0) instead of source | Yes — check `it.Length()` |
| 4 | corejson | `BytesCollection.go:783` | `ClonePtr` loop body | Same as #3 | Yes — check `it.Length()` |

## Notes

- The `AnyOnce.Deserialize` and `AnyErrorOnce.Deserialize` bugs (`if err == nil` instead of `if unmarshallErr == nil`) were **already fixed** in a prior commit and are no longer unreachable.
- All other branches in these two packages are believed to be reachable and covered by existing tests (Iterations 1–20).
