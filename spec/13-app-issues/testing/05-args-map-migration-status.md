# `args.Map` ExpectedInput Migration Status

> **Last updated:** 2026-03-06

## Summary

| Category | Count | % of Total |
|----------|-------|------------|
| ✅ Migrated to `args.Map` | **88 files** | 63.8% |
| 🔶 Using `args.Two`–`args.Six` (typed tuples) | **0 files** | 0% |
| 🔴 Using `[]string` | **15 files** | 10.9% |
| 🟡 Using plain `string` / other | **~35 files** | 25.4% |
| **Total testcase files** | **~138** | — |

> Note: Some files use multiple patterns (e.g., `[]string` for some cases, `args.Map` for others).

---

## ✅ Fully Migrated to `args.Map` (88 files)

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
| `corestrtests` | `BugfixRegression_testcases.go` | ~20 cases: Collection/Hashmap/Hashset/ValidValue |
| `coreversiontests` | `String_testcases.go` | Version string output |
| `corestrtests` | `LeftRightFromSplit_testcases.go` | 14 cases: left/right/isValid |
| `corestrtests` | `LeftMiddleRightFromSplit_testcases.go` | 14 cases: left/middle/right/isValid |
| `coresorttests` | `Sort_testcases.go` | Plain string expectations |
| `namevaluetests` | `Collection_testcases.go` | Collection CRUD operations |
| `namevaluetests` | `Instance_testcases.go` | Instance formatting/dispose |
| `pagingutiltests` | `Paging_testcases.go` | Paging calculations |
| `typesconvtests` | `TypesConv_testcases.go` | All Bool/Byte/Float/Int/String conversions |
| `reqtypetests` | `Request_testcases.go` | ~20 cases: identity/logical/HTTP methods |
| `iserrortests` | `iserror_testcases.go` | ~10 cases: Empty/Equal/EqualString |
| `errcoretests` | `MergeErrors_testcases.go` | ~10 cases: SliceToError/MergeErrors/ToStrings |
| `stringcompareastests` | `Glob_testcases.go` | 13 cases: glob match/inverse |
| `stringslicetests` | `CloneIf_testcases.go` | 6 cases: CloneIf/AnyItemsCloneIf |
| `versionindexestests` | `Index_testcases.go` | ~7 cases: JSON roundtrip/Name/Inject |
| `coreappendtests` | `Append_testcases.go` | 2 cases: PrependAppend |
| `keymktests` | `KeyLegend_testcases.go` | 1 case: GroupIntRange |
| `conditionaltests` | `If_testcases.go` | ifSlice: length+first → args.Map |
| `conditionaltests` | `ValueOrZeroNilVal_testcases.go` | ptrOrZero/nilValPtr: isNotNil+value → args.Map |
| `converterstests` | `StringTo_testcases.go` | value+hasError/isSuccess → args.Map |
| `converterstests` | `AnyItemConverter_testcases.go` | count+items → args.Map |
| `isanytests` | `IsAny_testcases.go` | isDefined/isNull, definedBoth/nullBoth → args.Map |
| `isanytests` | `DeepEqual_testcases.go` | isDeepEqual/isNotDeepEqual + named keys → args.Map |
| `coreutilstests` | `StringUtil_testcases.go` | splitLeftRight: left+right → args.Map |
| `ostypetests` | `OsType_testcases.go` | variationGroup/Identity: args.Three/Four → args.Map |
| `defaultcapacitytests` | `DefaultCapacity_testcases.go` | custom structs → args.Map ArrangeInput |
| `defaulterrtests` | `DefaultErr_testcases.go` | isNotNil+hasMessage → args.Map |
| `enumimpltests` | `enumTestCases.go` | min+max → args.Map |
| `issettertests` | `Value_testcases.go` | hasError/name, 8 boolean logic keys, conversions → args.Map |
| `coretaskinfotests` | `InfoCreate_testcases.go` | 20 test vars: name/desc/url/isSecure/isPlainText/hasExamples/field checks → args.Map |
| `coretestcasestests` | `ExpectedLines_testcases.go` | 8 scattered expected vars → table-driven with lineCount/line0..N args.Map |
| `coretestcasestests` | `GenericGherkins_testcases.go` | ExpectedLines []string → ExtraArgs with expectedDiff/expectedResult/line0..N keys |
| `isanytests` | `ExtendedTypedNil_testcases.go` | 4 legacy BaseTestCase vars → CaseV1 with result0..N/type0..N/types0..N args.Map |

---

## 🔶 Using Typed Tuples `args.Two`–`args.Six` (0 files)

All typed tuple files have been migrated to `args.Map`. ✅

---

## 🔴 Using `[]string` Expectations (20 files)

### Batch A — Migratable to `args.Map` (0 files)

All migratable `[]string` files have been migrated. ✅

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
Migrated ██████████████░░░░░░  88/138 (63.8%)
Tuples   ░░░░░░░░░░░░░░░░░░░░   0/138 ( 0.0%)
[]string ██░░░░░░░░░░░░░░░░░░░  15/138 (10.9%)
Other    █████░░░░░░░░░░░░░░░░  34/138 (24.6%)
```

### Changelog

| Date | Change |
|------|--------|
| 2026-03-06 | +2 migrated: `codestacktests/FileWithLine_testcases.go` (args.Three → args.Map with filePath/lineNumber/isValid) + `corecomparatortests/CompareExtended_testcases.go` (args.Three → args.Map with marshaledJson/unmarshaledName/unmarshaledValue). Total 88/138 (63.8%). All args.Three cleared. |
| 2026-03-06 | +1 migrated: `isanytests/ExtendedTypedNil_testcases.go` (4 legacy BaseTestCase vars → CaseV1 with result/type indexed args.Map). []string 16→15. Total 86/138 (62.3%). Batch A fully cleared. |
| 2026-03-06 | +2 migrated: `coretestcasestests/ExpectedLines_testcases.go` (8 scattered vars → table-driven args.Map) + `GenericGherkins_testcases.go` (ExpectedLines → ExtraArgs with semantic keys). []string 18→16. Total 85/138 (61.6%) |
| 2026-03-06 | +1 migrated: `coretaskinfotests/InfoCreate_testcases.go` (20 test vars, all []string → args.Map with semantic keys). []string 19→18. Total 83/138 (60.1%) |
| 2026-03-06 | Confirmed `issettertests/Value_testcases.go` already migrated — moved from Batch A to migrated. []string 20→19. Total 82/138 (59.4%) |
| 2026-03-06 | +11 migrated: easy []string batch (conditionaltests×2, converterstests×2, isanytests×2, coreutilstests×1, ostypetests×1, defaultcapacitytests×1, defaulterrtests×1, enumimpltests×1). []string 31→20. Total 81/138 (58.7%) |
| 2026-03-06 | +8 migrated: all remaining typed tuples (reqtypetests, iserrortests, errcoretests, stringcompareastests, stringslicetests, versionindexestests, coreappendtests, keymktests/KeyLegend). Tuples 11→0. Total 70/138 (50.7%) |
| 2026-03-06 | +2 migrated: `corestrtests/LeftRightFromSplit` (14 cases) + `LeftMiddleRightFromSplit` (14 cases) from []string → args.Map. Total 62/138 (44.9%) |
| 2026-03-06 | +1 migrated: `corestrtests/BugfixRegression_testcases.go` (~20 cases: args.Two–Five → args.Map). Tuples 12→11. Total 60/138 (43.5%) |
| 2026-03-06 | Full audit: +8 migrated (coreapitests×4, typesconvtests×1, coremathtests×1, LeftRight TypeStatus fix×1, recount×1). Tuples reduced 27→12 (many were already migrated or didn't exist). Updated []string batch lists. Total 59/138 (42.8%) |
| 2026-03-06 | Fixed coredynamictests tracking: moved 9 previously migrated files from tuples/[]string to migrated — total 51 |
| 2026-03-06 | +2 migrated: `corejsontests` (New_NewPtr, Result_Unmarshal) — total 42 |
| 2026-03-06 | +1 migrated: `Dynamic_testcases.go` (16 args.Two/Three/Four → args.Map) — total 40 |
| 2026-03-06 | +12 migrated: `coregenerictests` (5), `coreinstructiontests` (2), `namevaluetests` (2), `coresorttests` (1), `coreuniquetests` (1), `PointerSliceSorter` (1) — total 39 |
| 2026-03-06 | Initial audit — 19 migrated, 52 `[]string` |

---

## Migration Priority

### Priority 1 — `[]string` Medium (🟡 Medium, 0 files)
Batch A fully cleared. All remaining `[]string` files are in Batch B (kept as `[]string` for ordered sequences).
ExtendedTypedNil. Requires semantic key mapping for legacy BaseTestCase pattern.

### Priority 2 — Typed Tuples → `args.Map` (0 files) ✅ COMPLETE
All typed tuple files have been migrated.

### Priority 3 — Easy `[]string` (0 files) ✅ COMPLETE
All easy []string files have been migrated.

### Keep As-Is (15 files)
Variable-length output, multi-line error messages, formatted type inspection. Not suitable for `args.Map`.

---

## Related Docs

- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md) — `args.Map` mandate and patterns
- [Testing Patterns](/spec/01-app/13-testing-patterns.md) — AAA pattern and `CaseV1` usage
- [Edge-Case Coverage Audit](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md)
- [GoConvey Migration Plan](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)
