# Issue: corestrtests Build Failure — 76 Errors Blocking Coverage (3.9%)

## Date: 2026-03-25

## Summary

`corestrtests` failed to compile with 76 build errors across 10 test files, causing the entire package to be **BLOCKED** from coverage. This produced a phantom coverage drop to **3.9%** for `coredata/corestr`.

## Impact

Because Go runs all tests in a package as a single binary, **one compile error blocks the entire package**. All 85+ coverage test files are rendered useless until every error is fixed.

---

## Root Causes (6 Categories)

### 1. `Hashset.Strings` Signature Mismatch (~40 errors)
**Files:** `Coverage83_S23_*.go`, `Coverage84_S24_*.go`

`newHashsetCreator.Strings` takes a **single `[]string` slice**, not variadic strings.

**Wrong:** `corestr.New.Hashset.Strings("a", "b", "c")`
**Right (slice):** `corestr.New.Hashset.Strings([]string{"a", "b", "c"})`
**Right (variadic):** `corestr.New.Hashset.StringsSpreadItems("a", "b", "c")`

This is the **dominant error** — roughly 40 of the 76 errors.

### 2. Unexported Field Access: `SimpleStringOnce.isInitialize` (6 errors)
**Files:** `Coverage81_S21_*.go`, `Coverage82_S22_*.go`

`SimpleStringOnce.isInitialize` is an **unexported** field. Tests cannot access it directly.

**Wrong:** `sso.IsInitialize`
**Right:** Use the public API (e.g., `sso.IsInit()` or `sso.HasValue()` — check actual exported methods)

### 3. Non-Existent Factory Methods (5 errors)
**Files:** `Coverage74_S14_*.go`, `Coverage84_S24_*.go`

| Wrong Call | Correct Alternative |
|---|---|
| `corestr.New.Hashmap.StringsMap(m)` | `corestr.New.Hashmap.UsingMap(m)` |
| `corestr.New.Hashmap.StringsOfPairs("a","1","b","2")` | `corestr.New.Hashmap.KeyValuesStrings(keys, vals)` or build manually |
| `corestr.New.CharCollectionMap.Default()` | `corestr.New.CharCollectionMap.Empty()` |
| `corestr.New.CharHashsetMap.Default()` | `corestr.New.CharHashsetMap.Cap(0, 0)` |
| `corestr.New.SimpleSlice.Strings("a", "b")` (variadic) | `corestr.New.SimpleSlice.Strings([]string{"a", "b"})` (takes `[]string`) |

### 4. Pointer Method on Value Receiver: `corejson.Result.HasError` (3 errors)
**Files:** `Coverage_LinkedCollections_S13_*.go`, `Coverage_LinkedList_S12_*.go`, `Coverage_SimpleSlice_S11b_*.go`

`HasError()` has a pointer receiver (`*Result`), but the test calls it on a value `corejson.Result`.

**Wrong:** `jsonResult.HasError()` (where `jsonResult` is `corejson.Result`)
**Right:** `(&jsonResult).HasError()` or store as `*corejson.Result`

### 5. Type Redeclaration Conflict: `testErr` (2 errors)
**File:** `Coverage_Hashset_S10a_test.go` vs `Coverage40_Seg5_Hashset_test.go`

Both files declare a `testErr` type in the same package. Go treats all `_test.go` files in a directory as one compilation unit.

**Also:** `strings.NewReader("err")` is used where an `error` interface is expected — `*strings.Reader` does not implement `error`.

**Fix:** Rename the type in S10a (e.g., `testErrS10a`) and use a proper error value.

### 6. Type Mismatch & Incorrect Error Calling (2 errors)
**Files:** `Coverage85_S25_*.go`, `Coverage74_S14_*.go`

| Error | Cause |
|---|---|
| `LeftRightExpectingLengthMessager == ""` | `LeftRightExpectingLengthMessager` is `*errcore.ExpectingRecord`, not `string` |
| `jsonResult.Error(...)` called as function | `.Error` is a field/variable of type `error`, not a callable function |

---

## Fix Plan (by file)

| # | File | Errors | Fix |
|---|---|---|---|
| 1 | `Coverage83_S23_*.go` | ~40 | Change all `Hashset.Strings("a")` → `Hashset.Strings([]string{"a"})` |
| 2 | `Coverage82_S22_*.go` | 5 | Remove `sso.IsInitialize` → use exported method |
| 3 | `Coverage84_S24_*.go` | 9 | Fix factory names + Hashset.Strings signature |
| 4 | `Coverage81_S21_*.go` | 1 | Remove `sso.IsInitialize` → use exported method |
| 5 | `Coverage74_S14_*.go` | 4 | Fix `StringsMap` → real factory; fix `.Error` call |
| 6 | `Coverage_Hashset_S10a_*.go` | 3 | Rename `testErr` type; fix `strings.NewReader` |
| 7 | `Coverage85_S25_*.go` | 1 | Fix type comparison for `ExpectingRecord` |
| 8 | `Coverage_*_S11b/S12/S13_*.go` | 3 | Use pointer receiver for `HasError()` |

---

## Lessons Reinforced

1. **`newHashsetCreator.Strings` takes `[]string`, not variadic** — the most common mistake across all segments.
2. **Unexported fields cannot be tested directly** — always verify export status before accessing struct fields.
3. **All `_test.go` files share one compilation scope** — type names must be globally unique within the test package.
4. **Pointer vs value receivers matter** — `HasError()` on `*corejson.Result` cannot be called on a stack `corejson.Result`.
5. **Read the source before writing tests** — factory methods were invented (`StringsMap`, `StringsOfPairs`, `Default`) that don't exist.
