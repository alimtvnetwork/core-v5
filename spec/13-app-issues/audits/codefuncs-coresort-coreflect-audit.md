# Source-Code-Driven Audit: codefuncstests, coresorttests, coreflecttests

**Date**: 2026-03-09  
**Scope**: Negative path and boundary condition coverage gaps  
**Method**: Source code → branch identification → test inventory comparison

---

## 1. codefuncstests

### Source inventory (corefuncs/ — 15 files)

| File | Functions/Methods | Branches |
|------|-------------------|----------|
| `GetFuncName.go` | `GetFuncName(any) string` | nil input |
| `GetFuncFullName.go` | `GetFuncFullName(any) string` | nil input (comment warns "may panic") |
| `GetFunc.go` | `GetFunc(any) *runtime.Func` | nil input |
| `IsSuccessFuncWrapper` | `.Exec()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()` | success path, failure path (`!isSuccess`) |
| `ActionReturnsErrorFuncWrapper` | `.Exec()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()` | nil error, non-nil error |
| `NamedActionFuncWrapper` | `.Exec()`, `.Next()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()` | nil next pointer |
| `ResultDelegatingFuncWrapper` | `.Exec()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()` | nil error, non-nil error |
| `InOutErrFuncWrapper` | `.Exec()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()` | nil error, non-nil error |
| `InOutErrFuncWrapperOf[T]` | same + `.ToLegacy()` | nil error, non-nil error, type assertion panic |
| `InOutFuncWrapperOf[T]` | `.Exec()`, `.AsActionFunc()`, `.AsActionReturnsErrorFunc()`, `.ToLegacy()` | type assertion panic |
| `ResultDelegatingFuncWrapperOf[T]` | same + `.ToLegacy()` | nil error, non-nil error, type assertion |
| `InActionReturnsErrFuncWrapperOf[T]` | same + `.ToLegacy()` | nil error, non-nil error, type assertion |
| `SerializeOutputFuncWrapperOf[T]` | `.Exec()`, `.AsActionReturnsErrorFunc()` | nil error, non-nil error |

### Current test coverage

| Test | Cases | What's tested |
|------|-------|---------------|
| `GetFuncName_Verification` | 1 | Positive: named function returns non-empty |

### Gaps (Priority order)

| # | Gap | Severity | Category |
|---|-----|----------|----------|
| 1 | **`GetFuncName(nil)`** — no nil/non-func input test | High | Negative |
| 2 | **`GetFuncFullName`** — zero tests (comment warns panic risk) | High | Missing |
| 3 | **`GetFunc`** — zero tests | Medium | Missing |
| 4 | **All 6 legacy FuncWrapper types** — zero tests | High | Missing |
| 5 | **All 5 generic FuncWrapperOf types** — zero tests | High | Missing |
| 6 | **FuncWrapper error branches** — `AsActionReturnsErrorFunc` has `if err != nil` / `if !isSuccess` branches, untested | High | Negative |
| 7 | **`ToLegacy()` type assertion panics** — generic wrappers cast `input.(TIn)`, untested with wrong type | Medium | Boundary |
| 8 | **`NamedActionFuncWrapper.Next(nil)`** — nil pointer dereference risk | Medium | Negative |

**Total missing**: ~35+ test cases across 11 wrapper types + 3 standalone functions.

---

## 2. coresorttests

### Source inventory

| Package | Function | Signature |
|---------|----------|-----------|
| `intsort` | `Quick` | `(*[]int) → *[]int` |
| `intsort` | `QuickPtr` | `(*[]*int) → *[]*int` |
| `intsort` | `QuickDsc` | `(*[]int) → *[]int` |
| `intsort` | `QuickDscPtr` | `(*[]*int) → *[]*int` |
| `strsort` | `Quick` | `(*[]string) → *[]string` |
| `strsort` | `QuickPtr` | `(*[]*string) → *[]*string` |
| `strsort` | `QuickDsc` | `(*[]string) → *[]string` |
| `strsort` | `QuickDscPtr` | `(*[]*string) → *[]*string` |

### Current test coverage

| Test | Cases | What's tested |
|------|-------|---------------|
| `IntSort_Quick_Verification` | 2 | Positive: unsorted → sorted, already sorted → unchanged |
| `StrSort_Quick_Verification` | 1 | Positive: unsorted → sorted |

### Gaps

| # | Gap | Severity | Category |
|---|-----|----------|----------|
| 1 | **`intsort.QuickDsc`** — zero tests | High | Missing |
| 2 | **`intsort.QuickPtr`** — zero tests | High | Missing |
| 3 | **`intsort.QuickDscPtr`** — zero tests | High | Missing |
| 4 | **`strsort.QuickDsc`** — zero tests | High | Missing |
| 5 | **`strsort.QuickPtr`** — zero tests | High | Missing |
| 6 | **`strsort.QuickDscPtr`** — zero tests | High | Missing |
| 7 | **Empty slice** — `Quick(&[]int{})` boundary case | Medium | Boundary |
| 8 | **Single element** — `Quick(&[]int{42})` | Low | Boundary |
| 9 | **Negative numbers** — `Quick(&[]int{-3, 0, 2, -1})` | Medium | Boundary |
| 10 | **All duplicates** — `Quick(&[]int{5, 5, 5})` | Low | Boundary |
| 11 | **Reverse-sorted input** — worst-case for naive quicksort | Medium | Boundary |
| 12 | **Already-sorted `strsort.Quick`** — missing (int has it) | Low | Positive |

**Total missing**: 6 untested functions + ~6 boundary cases = ~18 test cases.

---

## 3. coreflecttests

### Source inventory

`reflectcore/vars.go` exports 14 facade variables. `FuncWrap` is tested via `args.ThreeFuncAny`.

### Current test coverage

| Test | Cases | What's tested |
|------|-------|---------------|
| `FuncWrap_Creation_Verification` | 8 | Positive (3 func variants), arg count mismatch, type mismatch, nil func, non-func input |
| `ReflectcoreVars_test.go` | 12 | Compile-time existence for 12 vars |

### Gaps

| # | Gap | Severity | Category |
|---|-----|----------|----------|
| 1 | **Duplicate test case** — lines 116-128 are identical to 103-115 (`someFunctionV2` test repeated verbatim) | High | Bug |
| 2 | **Raw `t.Error`** — `TypeName_NotNil` (line 43) and `TypeNames_NotNil` (line 48) use `t.Error` instead of project assertion framework | Medium | Standards |
| 3 | **Missing vars** — `TypeNamesString`, `TypeNamesReferenceString`, `ReflectGetterUsingReflectValue` not in existence checks | Medium | Missing |
| 4 | **Zero-arg function** — FuncWrap never tested with a zero-arg function | Low | Boundary |
| 5 | **Function with >3 args** — FuncWrap never tested with 4+ arg function | Low | Boundary |
| 6 | **Inconsistent test patterns** — existence checks use `_ = x` (compile-only) vs `if x == nil` (runtime), should be uniform | Low | Standards |

**Total missing**: 1 duplicate to fix, 2 raw `t.Error` violations, 3 missing vars, ~3 boundary cases.

---

## Summary

| Package | Existing Cases | Missing Cases | Priority |
|---------|---------------|---------------|----------|
| `codefuncstests` | 1 | ~35+ | 🔴 Critical |
| `coresorttests` | 3 | ~18 | 🟡 High |
| `coreflecttests` | 20 | ~9 (incl. 1 duplicate bug) | 🟡 High |
| **Total** | **24** | **~62** | |

### Recommended action order

1. **Fix coreflecttests duplicate** (line 116-128) — immediate, it inflates pass count
2. **Fix coreflecttests raw `t.Error`** — standards violation
3. **Add coresort missing functions** (QuickDsc, QuickPtr, QuickDscPtr × 2 packages) — high impact, low effort
4. **Add coresort boundary cases** (empty, single, negatives, duplicates, reverse)
5. **Add corefuncs FuncWrapper tests** — largest gap, highest effort
6. **Add corefuncs GetFuncFullName/GetFunc** — with nil input negative cases
