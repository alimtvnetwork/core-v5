# CaseNilSafe Migration Tracker

## Status Legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Migrated to CaseNilSafe |
| ⬜ | Not yet migrated |

---

## Migrated (✅)

| # | Package | File | Style Before | Cases |
|---|---------|------|-------------|-------|
| 1 | `corestrtests` | `Hashset_NilReceiver_testcases.go` | Inline `t.Error` | 5 |
| 2 | `regexnewtests` | `LazyRegex_NilReceiver_testcases.go` | Inline `t.Error` | 10 |
| 3 | `coreinstructiontests` | `StringCompare_NilReceiver_testcases.go` | CaseV1 string-dispatch | 5 |
| 4 | `coregenerictests` | `LinkedList_NilReceiver_testcases.go` | CaseV1 | 3 |
| 5 | `namevaluetests` | `Collection_NilReceiver_testcases.go` | CaseV1 | — |
| 6 | `coreoncetests` | `BytesErrorOnce_NilReceiver_testcases.go` | Custom `IsNilReceiver` wrapper | — |
| 7 | `corepayloadtests` | `TypedCollection_NilReceiver_testcases.go` | CaseV1 / GenericGherkins | 3 |
| 8 | `coreapitests` | `TypedConversions_NilReceiver_testcases.go` | CaseV1 string-dispatch | 4 |
| 9 | `casenilsafetests` | `CaseNilSafe_test.go` | N/A (self-test) | 12 |

---

## Remaining — Inline `t.Error` Style (Priority A)

These use raw `t.Error`/`t.Errorf` with manual nil-receiver setup. Highest migration value.

| # | Package | File | Methods to Migrate | Est. Cases |
|---|---------|------|--------------------|-----------|
| 1 | `reflectmodeltests` | `FieldProcessor_test.go` | `IsFieldType`, `IsFieldKind` | 2 |
| 2 | `reflectmodeltests` | `MethodProcessor_test.go` | `HasValidFunc`, `IsInvalid`, `Func`, `IsPublicMethod`, `GetType`, `Invoke` | 6 |
| 3 | `reflectmodeltests` | `ReflectValueKind_test.go` | `IsInvalid` + others | ~2 |
| 4 | `coredatatests` | `BytesError_test.go` | `HasError`, `IsEmptyError`, `IsEmpty`, `HandleError` + others | ~5 |
| 5 | `corevalidatortests` | `SliceValidators_test.go` | `IsMatch` | 1 |
| 6 | `corevalidatortests` | `SliceValidatorUnit_test.go` | `IsValid`, `ActualLinesLength`, `AllVerifyError`, `VerifyFirstError` | 4 |
| 7 | `corevalidatortests` | `SliceValidatorExtra_test.go` | `AllVerifyErrorExceptLast`, `AllVerifyErrorQuick`, `AllVerifyErrorTestCase`, `ActualLinesString`, `ExpectingLinesString`, `IsUsedAlready`, `VerifySimpleError` | 7 |
| 8 | `corevalidatortests` | `TextValidator_test.go` | `IsMatchMany`, `VerifyDetailError` | 2 |
| 9 | `corevalidatortests` | `TextValidators_test.go` | `Length`, `VerifyErrorMany` | 2 |
| 10 | `corevalidatortests` | `BaseLinesValidators_test.go` | `LinesValidatorsLength`, `IsEmptyLinesValidators`, `HasLinesValidators` | 3 |
| 11 | `corevalidatortests` | `LineValidator_test.go` | `IsMatchMany` | 1 |

**Subtotal: ~35 cases across 11 files**

---

## Remaining — CaseV1 with Nil Receiver (Priority B)

These use `CaseV1` with `(*Type)(nil)` in `ArrangeInput`. Well-structured but verbose.

| # | Package | File | Methods / Sections | Est. Cases |
|---|---------|------|--------------------|-----------|
| 12 | `coreapitests` | `PageRequest_testcases.go` | `IsPageSizeEmpty`, `IsPageIndexEmpty`, `HasPageSize`, `HasPageIndex`, `Clone` (nil) | 5 |
| 13 | `coredynamictests` | `Dynamic_testcases.go` + `Dynamic_test.go` | `ClonePtr`, `Bytes`, `ValueNullErr`, `ValueString`, `IntDefault` (nil receiver) | 5 |
| 14 | `coredynamictests` | `MapAnyItemsEdge_testcases.go` + `_test.go` | `IsEqualRaw`, `ClonePtr`, `Length`, `HasKey` (nil receiver) | 4 |
| 15 | `coredynamictests` | `AnyCollectionNewCreator_testcases.go` | `From` nil, `Clone` nil | 2 |
| 16 | `coregenerictests` | `Hashmap_testcases.go` + `Hashmap_test.go` | `IsEmpty`, `Length`, `HasItems` (nil receiver) | 3 |
| 17 | `coregenerictests` | `PairTripleExtended_test.go` | `Pair.Clear`, `Triple.Clear` (nil receiver) | 2 |
| 18 | `coregenerictests` | `Hashset_test.go` | `IsEmpty`, `Length`, `HasItems` (duplicate nil tests, generic type) | 5 |
| 19 | `corestrtests` | `BugfixRegression_testcases.go` + `_test.go` | `Hashmap.Clear` nil receiver | 1 |
| 20 | `coreinstructiontests` | `IdentifiersWithGlobals_testcases.go` | `Length` nil receiver | 1 |
| 21 | `coreinstructiontests` | `FromTo_test.go` | `ClonePtr` nil receiver | 1 |
| 22 | `trydotests` | `WrappedErr_testcases.go` | nil receiver state, string, exception type | 3 |
| 23 | `errcoretests` | `ErrorChain_testcases.go` | `ConcatMessageWithErr` nil | 1 |
| 24 | `coretaskinfotests` | `InfoCreate_testcases.go` | `SafeName`, `SafeDescription`, `SafeUrl`, `SafeHintUrl`, `SafeErrorUrl`, `SafeExampleUrl` | 6 |
| 25 | `coreapitests` | `TypedConversions_testcases.go` | `Clone` nil, `RequestInTo` nil (remaining CaseV1 sections) | 2 |

**Subtotal: ~41 cases across 14 files**

---

## Summary

| Category | Files | Est. Cases | Status |
|----------|-------|-----------|--------|
| ✅ Migrated | 9 | ~42 | Done |
| ⬜ Priority A (inline `t.Error`) | 11 | ~35 | Not started |
| ⬜ Priority B (CaseV1 nil) | 14 | ~41 | Not started |
| **Total** | **34** | **~118** | **~36% done** |

---

## Recommended Migration Order

1. **Priority A** — `corevalidatortests/` (5 files, 17 cases) — highest density of raw `t.Error`
2. **Priority A** — `reflectmodeltests/` (3 files, 10 cases) — inline style
3. **Priority A** — `coredatatests/BytesError_test.go` (5 cases)
4. **Priority B** — `coredynamictests/` (3 files, 11 cases) — CaseV1 with manual setup
5. **Priority B** — `coregenerictests/` (3 files, 10 cases) — requires generic literal wrappers
6. **Priority B** — remaining scattered files (5 files, ~14 cases)

## Notes

- Generic types (`Hashset[T]`, `Hashmap[K,V]`, `Pair[A,B]`, `Triple[A,B,C]`) require the **function literal wrapper** pattern documented in the design doc §7.
- Some CaseV1 nil tests also test non-nil behavior in the same variable (e.g., `PageRequest` has nil + valid cases in one slice). Migration should extract only the nil cases into `CaseNilSafe`, leaving the rest in CaseV1.
- `corevalidatortests` has the most files (5) and would benefit most from consolidation into per-type `_NilReceiver_testcases.go` files.
