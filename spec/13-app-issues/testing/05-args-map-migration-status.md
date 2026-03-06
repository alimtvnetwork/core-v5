# `args.Map` ExpectedInput Migration Status

> **Last updated:** 2026-03-06

## Summary

| Category | Count | % of Total |
|----------|-------|------------|
| ✅ Migrated to `args.Map` | **59 files** | 42.8% |
| 🔶 Using `args.Two`–`args.Six` (typed tuples) | **12 files** | 8.7% |
| 🔴 Using `[]string` | **33 files** | 23.9% |
| 🟡 Using plain `string` / other | **~34 files** | 24.6% |
| **Total testcase files** | **~138** | — |

> Note: Some files use multiple patterns (e.g., `[]string` for some cases, `args.Map` for others).

---

## ✅ Fully Migrated to `args.Map` (59 files)

| Package | File | Notes |
|---------|------|-------|
| `bytetypetests` | `Variant_testcases.go` | Reference implementation |
| `chmodhelpertests` | `PartialRwxVerify_testcases.go` | Simple boolean expectations |
| `coreapitests` | `PageRequest_testcases.go` | Clone fields + independence |
| `coreapitests` | `TypedRequest_testcases.go` | New/Invalid/Clone/ToTypedResponse/Message |
| `coreapitests` | `TypedApiTypes_testcases.go` | RequestIn/Response/ResponseResult |
| `coreapitests` | `TypedConversions_testcases.go` | SimpleGenericRequest/Conversions |
| `coredynamictests` | `MapAnyItems_testcases.go` | 4 named test cases |
| `coredynamictests` | `MapAnyItemsEdge_testcases.go` | 19 named test cases |
| `coredynamictests` | `Dynamic_testcases.go` | 16 cases: constructors, clone, bytes, loop, items |
| `coredynamictests` | `AnyCollectionLock_testcases.go` | Lock/unlock assertions |
| `coredynamictests` | `CollectionLock_testcases.go` | Lock/unlock assertions |
| `coredynamictests` | `CollectionSort_testcases.go` | Sort result assertions |
| `coredynamictests` | `CollectionDistinct_testcases.go` | Distinct count + items |
| `coredynamictests` | `CollectionGroupBy_testcases.go` | GroupBy result assertions |
| `coredynamictests` | `LeftRight_testcases.go` | Left/Right pair assertions |
| `coredynamictests` | `CollectionMap_testcases.go` | Map transformation assertions |
| `coredynamictests` | `CollectionNewCreator_testcases.go` | Creator result assertions |
| `coredynamictests` | `AnyCollectionNewCreator_testcases.go` | Creator result assertions |
| `coregenerictests` | `Collection_testcases.go` | Collection operations |
| `coregenerictests` | `CollectionEdgeCases_testcases.go` | Edge cases with semantic keys |
| `coregenerictests` | `duplicateEdgeCases_testcases.go` | Distinct/duplicate operations |
| `coregenerictests` | `orderedfuncs_testcases.go` | Clamp/MinMax operations |
| `coregenerictests` | `PointerSliceSorter_testcases.go` | Mixed: args.Map + ordered []string |
| `coreinstructiontests` | `Identifier_testcases.go` | Identifiers, Specs, BaseTags |
| `coreinstructiontests` | `IdentifiersWithGlobals_testcases.go` | WithGlobals CRUD operations |
| `coreinstructiontests` | `FromTo_testcases.go` | Already migrated |
| `coreinstructiontests` | `StringCompare_testcases.go` | Already migrated |
| `coreinstructiontests` | `StringSearch_testcases.go` | Already migrated |
| `corejsontests` | `New_NewPtr_testcases.go` | 6 cases: New/NewPtr constructors |
| `corejsontests` | `Result_Unmarshal_testcases.go` | 4 cases: Unmarshal valid/nil/invalid/error |
| `coremathtests` | `MinMaxInt_testcases.go` | MinMax operations |
| `coreoncetests` | `IntegerOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `StringOnce_testcases.go` | 5 test case arrays |
| `coreoncetests` | `BoolOnce_testcases.go` | 3 test case arrays |
| `coreoncetests` | `BytesOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `BytesErrorOnce_testcases.go` | 13 test case arrays |
| `coreoncetests` | `ErrorOnce_testcases.go` | 6 test case arrays |
| `coreoncetests` | `ByteOnce_testcases.go` | 5 test case arrays |
| `coreoncetests` | `IntegersOnce_testcases.go` | 7 test case arrays |
| `coreoncetests` | `StringsOnce_testcases.go` | 7 test case arrays |
| `coreoncetests` | `MapStringStringOnce_testcases.go` | 7 test case arrays |
| `coreoncetests` | `AnyOnce_testcases.go` | 5 test case arrays |
| `coreoncetests` | `AnyErrorOnce_testcases.go` | 6 test case arrays |
| `corepayloadtests` | `TypedCollection_testcases.go` | Collection operations |
| `corepayloadtests` | `TypedCollectionFlatMap_testcases.go` | FlatMap operations |
| `corepayloadtests` | `TypedCollectionGroupBy_testcases.go` | GroupBy operations |
| `corepayloadtests` | `TypedCollectionPartition_testcases.go` | Partition operations |
| `corepayloadtests` | `PayloadWrapper_testcases.go` | Complex struct assertions |
| `corepayloadtests` | `PagingInfo_testcases.go` | Paging metadata |
| `corepayloadtests` | `TypedWrapper_testcases.go` | Wrapper round-trips |
| `corepayloadtests` | `Attributes_testcases.go` | Attribute key-value |
| `coreversiontests` | `ComparisonExtended_testcases.go` | Version comparison |
| `coreversiontests` | `Comparison_testcases.go` | Version comparison |
| `coreversiontests` | `Parse_testcases.go` | Version parsing |
| `coreversiontests` | `String_testcases.go` | Version string output |
| `coresorttests` | `Sort_testcases.go` | Plain string expectations |
| `namevaluetests` | `Collection_testcases.go` | Collection CRUD operations |
| `namevaluetests` | `Instance_testcases.go` | Instance formatting/dispose |
| `pagingutiltests` | `Paging_testcases.go` | Paging calculations |
| `typesconvtests` | `TypesConv_testcases.go` | All Bool/Byte/Float/Int/String conversions |

---

## 🔶 Using Typed Tuples `args.Two`–`args.Six` (12 files)

These use positional typed tuples — better than `[]string` but lack semantic keys. **Migration to `args.Map` recommended.**

| Package | File | Tuple Types | Cases |
|---------|------|-------------|-------|
| `coreappendtests` | `Append_testcases.go` | `args.Three` | 2 |
| `corestrtests` | `BugfixRegression_testcases.go` | `args.Two`–`args.Five` | ~20 |
| `errcoretests` | `MergeErrors_testcases.go` | `args.Two`/`args.Three` | ~10 |
| `iserrortests` | `iserror_testcases.go` | `args.Two`/`args.Three` | ~10 |
| `keymktests` | `KeyLegend_testcases.go` | `args.Three` | 1 |
| `reqtypetests` | `Request_testcases.go` | `args.Three`–`args.Five` | ~20 |
| `stringcompareastests` | `Glob_testcases.go` | `args.Two` | ~13 |
| `stringslicetests` | `CloneIf_testcases.go` | `args.Two`–`args.Five` | 6 |
| `versionindexestests` | `Index_testcases.go` | `args.Two` | ~6 |

> **Note:** Files previously listed here that have been migrated: `coreapitests/*` (4 files), `typesconvtests/*` (1 file), `converterstests/*`, `coregenerictests/*`, `corestrtests/Collection`, `corestrtests/Hashmap`, `corestrtests/Hashset`, `corestrtests/ValidValue`, `regexnewtests/*`, `simplewraptests/*`, `keymktests/Key_testcases.go` — these were either migrated or confirmed to not use typed tuples.

---

## 🔴 Using `[]string` Expectations (33 files)

### Batch A — Migratable to `args.Map` (18 files)

| Package | File | Fields | Difficulty |
|---------|------|--------|------------|
| `conditionaltests` | `If_testcases.go` | 1–2 values | 🟢 Easy |
| `conditionaltests` | `ValueOrZeroNilVal_testcases.go` | 1–2 values | 🟢 Easy |
| `converterstests` | `StringTo_testcases.go` | 2 values (result + hasError) | 🟢 Easy |
| `converterstests` | `AnyItemConverter_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `CastedResult_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `CollectionSearch_testcases.go` | 2–3 values | 🟢 Easy |
| `corestrtests` | `LeftMiddleRightFromSplit_testcases.go` | 4 values × 7 cases | 🟡 Medium |
| `corestrtests` | `SimpleSlice_testcases.go` | 1–3 values | 🟢 Easy |
| `coretaskinfotests` | `InfoCreate_testcases.go` | 2–4 values | 🟢 Easy |
| `coreutilstests` | `StringUtil_testcases.go` | 1–2 values | 🟢 Easy |
| `defaulterrtests` | `DefaultErr_testcases.go` | 2 booleans × 11 cases | 🟢 Easy |
| `defaultcapacitytests` | `DefaultCapacity_testcases.go` | 1 value | 🟢 Easy |
| `enumimpltests` | `BasicEnum_testcases.go` | 2–3 values | 🟢 Easy |
| `isanytests` | `IsAny_testcases.go` | 2–3 booleans | 🟢 Easy |
| `isanytests` | `DeepEqual_testcases.go` | 2 booleans | 🟢 Easy |
| `issettertests` | `Value_testcases.go` | 2–8 booleans | 🟡 Medium |
| `ostypetests` | `OsType_testcases.go` | 1–2 values | 🟢 Easy |
| `codetestcasestests` | `GenericGherkins_testcases.go` | 1–2 values | 🟢 Easy |

### Batch B — Keep as `[]string` (15 files)

Variable-length output, multi-line error messages, formatted type inspection. Not suitable for `args.Map`.

| Package | File | Reason |
|---------|------|--------|
| `chmodhelpertests` | `ApplyOnPath_testcases.go` | Multi-line file operation results |
| `chmodhelpertests` | `LinuxApplyRecursiveOnPath_testcases.go` | OS-dependent results |
| `chmodhelpertests` | `SimpleFileWriter_CreateDir_testcases.go` | Multi-line formatted output |
| `chmodhelpertests` | `DirFilesWithContent_testcases.go` | Multi-line formatted output |
| `chmodhelpertests` | `VerifyPartialRwxLocations_testcases.go` | Multi-line comparison output |
| `coredatatests` | `FuncWrap_testcases.go` | Multi-line error messages |
| `coreflecttests` | `FuncWrap_testcases.go` | Multi-line error messages |
| `coredynamictests` | `CollectionGetPagesSize_testcases.go` | Variable-length paging output |
| `coredynamictests` | `CollectionGetPagesSize_Others_testcases.go` | Variable-length paging output |
| `isanytests` | `ExtendedTypedNil_testcases.go` | Formatted type inspection output |
| `integratedtests` | `GetAssert_testcases.go` | Multi-line assertion output |
| `integratedtests` | `GetAssert_ToStrings_testcases.go` | Formatted conversion output |
| `integratedtests` | `GetAssert_SimpleTestCasesWrapper_testcases.go` | Multi-line wrapper output |
| `corejsontests` | `Deserializer_Apply_testcases.go` | JSON comparison output |
| `corepayloadtests` | `TypedCollectionPagingEdge_testcases.go` | Variable paging |

---

## 🟡 Using Plain `string` / Other Expectations (~34 files)

Single-value expectations stored as bare strings or other simple types. **Low priority** — already simple and readable.

---

## Migration Progress

```
Migrated █████████░░░░░░░░░░░░  59/138 (42.8%)
Tuples   ██░░░░░░░░░░░░░░░░░░  12/138 ( 8.7%)
[]string █████░░░░░░░░░░░░░░░░  33/138 (23.9%)
Other    █████░░░░░░░░░░░░░░░░  34/138 (24.6%)
```

### Changelog

| Date | Change |
|------|--------|
| 2026-03-06 | Full audit: +8 migrated (coreapitests×4, typesconvtests×1, coremathtests×1, LeftRight TypeStatus fix×1, recount×1). Tuples reduced 27→12 (many were already migrated or didn't exist). Updated []string batch lists. Total 59/138 (42.8%) |
| 2026-03-06 | Fixed coredynamictests tracking: moved 9 previously migrated files from tuples/[]string to migrated — total 51 |
| 2026-03-06 | +2 migrated: `corejsontests` (New_NewPtr, Result_Unmarshal) — total 42 |
| 2026-03-06 | +1 migrated: `Dynamic_testcases.go` (16 args.Two/Three/Four → args.Map) — total 40 |
| 2026-03-06 | +12 migrated: `coregenerictests` (5), `coreinstructiontests` (2), `namevaluetests` (2), `coresorttests` (1), `coreuniquetests` (1), `PointerSliceSorter` (1) — total 39 |
| 2026-03-06 | Initial audit — 19 migrated, 52 `[]string` |

---

## Migration Priority

### Priority 1 — `[]string` Quick Wins (🟢 Easy, 16 files)
All single/dual-boolean or 1–3 value expectations. Estimated: ~1 hour total.

### Priority 2 — `[]string` Medium (🟡 Medium, 2 files)
Multi-field structs or 4+ positional values. Estimated: ~30 min total.

### Priority 3 — Typed Tuples → `args.Map` (12 files)
`args.Two`–`args.Five` → `args.Map` with semantic keys. Higher effort but significant diagnostic improvement.

### Keep As-Is (15 files)
Variable-length output, multi-line error messages, formatted type inspection. Not suitable for `args.Map`.

---

## Related Docs

- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md) — `args.Map` mandate and patterns
- [Testing Patterns](/spec/01-app/13-testing-patterns.md) — AAA pattern and `CaseV1` usage
- [Edge-Case Coverage Audit](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md)
- [GoConvey Migration Plan](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)
