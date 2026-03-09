# CaseNilSafe Migration Tracker

## Status Legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Migrated to CaseNilSafe with `results.ResultAny` |
| ⬜ | Not yet migrated |

---

## Architecture (Corrected)

All migrated test cases now use the **corrected architecture**:

| Aspect | Old (incorrect) | New (correct) |
|--------|----------------|---------------|
| **Expected type** | `args.Map` (untyped) | `results.ResultAny` (typed struct) |
| **Assertion owner** | Methods on `CaseNilSafe` | Methods on `results.Result[T]` (`ShouldMatchResult`) |
| **CaseNilSafe role** | Logic + data | Data-only + thin `ShouldBeSafe` convenience that delegates to `Result` |
| **Error sentinel** | `"hasError": true` in map | `results.ExpectAnyError` error value |
| **Field comparison** | Manual map key matching | Auto-derived from `Expected` fields, with optional `CompareFields` override |

See `spec/01-app/designs/CaseNilSafe-design.md` for full architecture.

---

## Migrated (✅)

| # | Package | File | Style Before | Cases |
|---|---------|------|-------------|-------|
| 1 | `corestrtests` | `Hashset_NilReceiver_testcases.go` | Inline `t.Error` | 5 |
| 2 | `regexnewtests` | `LazyRegex_NilReceiver_testcases.go` | Inline `t.Error` | 10 |
| 3 | `coreinstructiontests` | `StringCompare_NilReceiver_testcases.go` | CaseV1 string-dispatch | 5 |
| 4 | `coregenerictests` | `LinkedList_NilReceiver_testcases.go` | CaseV1 | 3 |
| 5 | `namevaluetests` | `Collection_NilReceiver_testcases.go` | CaseV1 | 1 |
| 6 | `coreoncetests` | `BytesErrorOnce_NilReceiver_testcases.go` | Custom `IsNilReceiver` wrapper | 4 |
| 7 | `corepayloadtests` | `TypedCollection_NilReceiver_testcases.go` | CaseV1 / GenericGherkins | 3 |
| 8 | `coreapitests` | `TypedConversions_NilReceiver_testcases.go` | CaseV1 string-dispatch | 4 |
| 9 | `casenilsafetests` | `CaseNilSafe_test.go` | N/A (self-test) | 12 |
| 10 | `reflectmodeltests` | `FieldProcessor_NilReceiver_testcases.go` | Inline `t.Error` | 2 |
| 11 | `reflectmodeltests` | `MethodProcessor_NilReceiver_testcases.go` | Inline `t.Error` | 10 |
| 12 | `reflectmodeltests` | `ReflectValueKind_NilReceiver_testcases.go` | Inline `t.Error` | 8 |
| 13 | `coredatatests` | `BytesError_NilReceiver_testcases.go` | Inline `t.Error` | 6 |
| 14 | `corevalidatortests` | `SliceValidator_NilReceiver_testcases.go` | Inline `t.Error` | 11 |
| 15 | `corevalidatortests` | `SliceValidators_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 16 | `corevalidatortests` | `TextValidator_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 17 | `corevalidatortests` | `TextValidators_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 18 | `corevalidatortests` | `BaseLinesValidators_NilReceiver_testcases.go` | Inline `t.Error` | 5 |
| 19 | `corevalidatortests` | `LineValidator_NilReceiver_testcases.go` | Inline `t.Error` | 1 |

**Subtotal: 19 files, ~99 cases**

---

## Remaining — CaseV1 with Nil Receiver (Priority B)

These use `CaseV1` with `(*Type)(nil)` in `ArrangeInput`. Well-structured but verbose.

| # | Package | File | Methods / Sections | Est. Cases |
|---|---------|------|--------------------|-----------|
| 1 | `coreapitests` | `PageRequest_testcases.go` | `IsPageSizeEmpty`, `IsPageIndexEmpty`, `HasPageSize`, `HasPageIndex`, `Clone` (nil) | 5 |
| 2 | `coredynamictests` | `Dynamic_testcases.go` + `Dynamic_test.go` | `ClonePtr`, `Bytes`, `ValueNullErr`, `ValueString`, `IntDefault` (nil receiver) | 5 |
| 3 | `coredynamictests` | `MapAnyItemsEdge_testcases.go` + `_test.go` | `IsEqualRaw`, `ClonePtr`, `Length`, `HasKey` (nil receiver) | 4 |
| 4 | `coredynamictests` | `AnyCollectionNewCreator_testcases.go` | `From` nil, `Clone` nil | 2 |
| 5 | `coregenerictests` | `Hashmap_testcases.go` + `Hashmap_test.go` | `IsEmpty`, `Length`, `HasItems` (nil receiver) | 3 |
| 6 | `coregenerictests` | `PairTripleExtended_test.go` | `Pair.Clear`, `Triple.Clear` (nil receiver) | 2 |
| 7 | `coregenerictests` | `Hashset_test.go` | `IsEmpty`, `Length`, `HasItems` (duplicate nil tests, generic type) | 5 |
| 8 | `corestrtests` | `BugfixRegression_testcases.go` + `_test.go` | `Hashmap.Clear` nil receiver | 1 |
| 9 | `coreinstructiontests` | `IdentifiersWithGlobals_testcases.go` | `Length` nil receiver | 1 |
| 10 | `coreinstructiontests` | `FromTo_test.go` | `ClonePtr` nil receiver | 1 |
| 11 | `trydotests` | `WrappedErr_testcases.go` | nil receiver state, string, exception type | 3 |
| 12 | `errcoretests` | `ErrorChain_testcases.go` | `ConcatMessageWithErr` nil | 1 |
| 13 | `coretaskinfotests` | `InfoCreate_testcases.go` | `SafeName`, `SafeDescription`, `SafeUrl`, `SafeHintUrl`, `SafeErrorUrl`, `SafeExampleUrl` | 6 |
| 14 | `coreapitests` | `TypedConversions_testcases.go` | `Clone` nil, `RequestInTo` nil (remaining CaseV1 sections) | 2 |

**Subtotal: ~41 cases across 14 files**

---

## Summary

| Category | Files | Est. Cases | Status |
|----------|-------|-----------|--------|
| ✅ Migrated (ResultAny) | 19 | ~99 | Done |
| ⬜ Priority A (inline `t.Error`) | 0 | 0 | **All done** ✅ |
| ⬜ Priority B (CaseV1 nil) | 14 | ~41 | Not started |
| **Total** | **33** | **~140** | **~71% done** |

---

## Recommended Migration Order

1. ~~**Priority A** — `reflectmodeltests/` (3 files, 20 cases) — inline style~~ ✅ Done
2. ~~**Priority A** — `coredatatests/BytesError_test.go` (6 cases)~~ ✅ Done
3. ~~**Priority A** — `corevalidatortests/` (7 files, 26 cases)~~ ✅ Done
4. **Priority B** — `coredynamictests/` (3 files, 11 cases) — CaseV1 with manual setup
5. **Priority B** — `coregenerictests/` (3 files, 10 cases) — requires generic literal wrappers
6. **Priority B** — remaining scattered files (8 files, ~20 cases)

## Notes

- **Architecture corrected**: All migrated files now use `results.ResultAny` for `Expected` (not `args.Map`). Assertion logic lives on `Result[T].ShouldMatchResult`, with `CaseNilSafe.ShouldBeSafe` as a thin convenience.
- Generic types (`Hashset[T]`, `Hashmap[K,V]`, `Pair[A,B]`, `Triple[A,B,C]`) require the **function literal wrapper** pattern documented in the design doc §7.
- Some CaseV1 nil tests also test non-nil behavior in the same variable (e.g., `PageRequest` has nil + valid cases in one slice). Migration should extract only the nil cases into `CaseNilSafe`, leaving the rest in CaseV1.
- All Priority A (inline `t.Error`) migrations are now complete. Only Priority B (CaseV1) remains.
- `Expected` uses `results.ResultAny` with `CompareFields` for subset assertion.
- `ExpectAnyError` sentinel is used for methods expected to return non-nil errors.
- `BytesErrorOnce` and `Collection` testcases were retroactively corrected from `args.Map` → `results.ResultAny`.
