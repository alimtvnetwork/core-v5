# `args.Map` ExpectedInput Migration Status

> **Last updated:** 2026-03-06

## Summary

| Category | Count | % of Total |
|----------|-------|------------|
| ✅ Migrated to `args.Map` | **27 files** | 20.1% |
| 🔶 Using `args.Two`–`args.Six` (typed tuples) | **36 files** | 26.9% |
| 🔴 Using `[]string` | **50 files** | 37.3% |
| 🟡 Using plain `string` / other | **21 files** | 15.7% |
| **Total testcase files** | **~134** | — |

> Note: Some files use multiple patterns (e.g., `[]string` for some cases, `args.Map` for others).

---

## ✅ Fully Migrated to `args.Map` (27 files)

| Package | File | Notes |
|---------|------|-------|
| `bytetypetests` | `Variant_testcases.go` | Reference implementation |
| `chmodhelpertests` | `PartialRwxVerify_testcases.go` | Simple boolean expectations |
| `coredynamictests` | `MapAnyItems_testcases.go` | 4 named test cases |
| `coredynamictests` | `MapAnyItemsEdge_testcases.go` | 19 named test cases |
| `coreoncetests` | `IntegerOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `StringOnce_testcases.go` | 5 test case arrays |
| `coreoncetests` | `BoolOnce_testcases.go` | 3 test case arrays |
| `coreoncetests` | `BytesOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `BytesErrorOnce_testcases.go` | 13 test case arrays |
| `coreoncetests` | `ErrorOnce_testcases.go` | 6 test case arrays |
| `coreoncetests` | `ByteOnce_testcases.go` | 5 test case arrays (new) |
| `coreoncetests` | `IntegersOnce_testcases.go` | 7 test case arrays (new) |
| `coreoncetests` | `StringsOnce_testcases.go` | 7 test case arrays (new) |
| `coreoncetests` | `MapStringStringOnce_testcases.go` | 7 test case arrays (new) |
| `coreoncetests` | `AnyOnce_testcases.go` | 5 test case arrays (new) |
| `coreoncetests` | `AnyErrorOnce_testcases.go` | 6 test case arrays (new) |
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
| `pagingutiltests` | `Paging_testcases.go` | Paging calculations |

---

## 🔶 Using Typed Tuples `args.Two`–`args.Six` (36 files)

These use positional typed tuples — better than `[]string` but lack semantic keys. **Migration to `args.Map` recommended.**

| Package | File | Tuple Type |
|---------|------|-----------|
| `coreapitests` | `PageRequest_testcases.go` | `args.Two` |
| `coreapitests` | `SearchRequest_testcases.go` | `args.Two` |
| `coreapitests` | `RequestAttribute_testcases.go` | `args.Two` |
| `coreapitests` | `ResponseAttribute_testcases.go` | `args.Two` |
| `coreapitests` | `TypedRequest_testcases.go` | `args.Two` |
| `converterstests` | `StringsTo_testcases.go` | `args.Five` |
| `converterstests` | `StringTo_testcases.go` | `args.Five` |
| `coredynamictests` | `CollectionGroupBy_testcases.go` | `args.Three` |
| `coredynamictests` | `LeftRight_testcases.go` | `args.Two` |
| `coregenerictests` | `CollectionBranch_testcases.go` | `args.Two` |
| `coregenerictests` | `CollectionSerialization_testcases.go` | `args.Two` |
| `coregenerictests` | `HashsetBranch_testcases.go` | `args.Two` |
| `coregenerictests` | `Hashmap_testcases.go` | `args.Two` |
| `coregenerictests` | `TripleFromSplit_testcases.go` | `args.Three` |
| `coregenerictests` | `comparablefuncs_testcases.go` | `args.Two` |
| `corejsontests` | `New_NewPtr_testcases.go` | `args.Two` |
| `corejsontests` | `Result_Unmarshal_testcases.go` | `args.Two`/`args.Three` |
| `coreinstructiontests` | `StringSearch_testcases.go` | `args.Two` |
| `coreinstructiontests` | `FromTo_testcases.go` | `args.Two` |
| `corestrtests` | `Collection_testcases.go` | `args.Two` |
| `corestrtests` | `Hashmap_testcases.go` | `args.Two` |
| `corestrtests` | `Hashset_testcases.go` | `args.Two` |
| `corestrtests` | `ValidValue_testcases.go` | `args.Two` |
| `iserrortests` | `iserror_testcases.go` | `args.Two`/`args.Three` |
| `keymktests` | `KeyLegend_testcases.go` | `args.Two` |
| `keymktests` | `Key_testcases.go` | `args.Two` |
| `namevaluetests` | `Collection_testcases.go` | `args.Two`/`args.Three` |
| `namevaluetests` | `Instance_testcases.go` | `args.Two` |
| `regexnewtests` | `LazyRegex_testcases.go` | `args.Two` |
| `regexnewtests` | `LazyRegex_Methods_testcases.go` | `args.Two` |
| `simplewraptests` | `Wrapper_testcases.go` | `args.Two` |
| `stringcompareastests` | `Glob_testcases.go` | `args.Two` |
| `typesconvtests` | `Bool_testcases.go` | `args.Two` |
| `typesconvtests` | `Byte_testcases.go` | `args.Two` |
| `typesconvtests` | `Float_testcases.go` | `args.Two` |
| `typesconvtests` | `Integer_testcases.go` | `args.Two` |
| `typesconvtests` | `String_testcases.go` | `args.Two` |
| `versionindexestests` | `Index_testcases.go` | `args.Two` |

---

## 🔴 Using `[]string` Expectations (50 files)

These use raw `[]string` slices with positional semantics. **Highest priority for migration.**

### Wave 1 — Quick Wins (simple boolean/value expectations)

| Package | File | Fields | Difficulty |
|---------|------|--------|------------|
| `defaulterrtests` | `DefaultErr_testcases.go` | 2 booleans × 11 cases | 🟢 Easy |
| `issettertests` | `Value_testcases.go` | 2–8 booleans | 🟡 Medium |
| `coregenerictests` | `CollectionEdgeCases_testcases.go` | 1–4 values × 25 cases | 🟡 Medium |
| `coregenerictests` | `funcs_testcases.go` | 1–3 values × 10 cases | 🟢 Easy |
| `enumimpltests` | `BasicEnum_testcases.go` | 2–3 values | 🟢 Easy |
| `reqtypetests` | `Type_testcases.go` | 1–2 values | 🟢 Easy |
| `ostypetests` | `OsType_testcases.go` | 1–2 values | 🟢 Easy |
| `anycmptests` | `QuickCompare_testcases.go` | 1 value | 🟢 Easy |
| `corecmptests` | `Time_testcases.go` | 1–2 values | 🟢 Easy |
| `coretaskinfotests` | `InfoCreate_testcases.go` | 2–4 values | 🟢 Easy |
| `converterstests` | `AnyItemConverter_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `AnyCollectionNewCreator_testcases.go` | 2–4 values × 12 cases | 🟡 Medium |
| `coredynamictests` | `CollectionNewCreator_testcases.go` | 2–4 values | 🟡 Medium |
| `coredynamictests` | `AnyCollectionLock_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `CollectionLock_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `CollectionMap_testcases.go` | 2–4 values × 7 cases | 🟡 Medium |
| `coredynamictests` | `CastedResult_testcases.go` | 1–3 values | 🟢 Easy |
| `coredynamictests` | `Dynamic_testcases.go` | 2–5 values | 🟡 Medium |
| `coredynamictests` | `CollectionSort_testcases.go` | 2–3 values | 🟢 Easy |
| `coredynamictests` | `CollectionSearch_testcases.go` | 2–3 values | 🟢 Easy |
| `coredynamictests` | `CollectionDistinct_testcases.go` | 1–2 values | 🟢 Easy |
| `corestrtests` | `LeftMiddleRightFromSplit_testcases.go` | 4 values × 7 cases | 🟡 Medium |
| `corestrtests` | `BugfixRegression_testcases.go` | 1–2 values | 🟢 Easy |
| `corestrtests` | `SimpleSlice_testcases.go` | 1–3 values | 🟢 Easy |
| `coresorttests` | `Sort_testcases.go` | 1 value | 🟢 Easy |
| `coreinstructiontests` | `Identifier_testcases.go` | 3–5 values × 10 cases | 🟡 Medium |
| `coreinstructiontests` | `IdentifiersWithGlobals_testcases.go` | 3–5 values | 🟡 Medium |
| `codefuncstests` | `GetFuncName_testcases.go` | 1 value | 🟢 Easy |
| `coreutilstests` | `StringUtil_testcases.go` | 1–2 values | 🟢 Easy |
| `coretestcasestests` | `GenericGherkins_testcases.go` | 1–2 values | 🟢 Easy |
| `defaultcapacitytests` | `DefaultCapacity_testcases.go` | 1 value | 🟢 Easy |

### Wave 3 — Keep as `[]string` (variable-length / formatted output)

| Package | File | Reason |
|---------|------|--------|
| `chmodhelpertests` | `ApplyOnPath_testcases.go` | Multi-line file operation results |
| `chmodhelpertests` | `LinuxApplyRecursiveOnPath_testcases.go` | OS-dependent results |
| `coredatatests` | `FuncWrap_testcases.go` | Multi-line error messages |
| `coreflecttests` | `FuncWrap_testcases.go` | Multi-line error messages |
| `coredynamictests` | `CollectionGetPagesSize_testcases.go` | Variable-length paging output |
| `coredynamictests` | `CollectionGetPagesSize_Others_testcases.go` | Variable-length paging output |
| `coredynamictests` | `CollectionGroupBy_testcases.go` | Empty `[]string{}` cases |
| `isanytests` | `ExtendedTypedNil_testcases.go` | Formatted type inspection output |
| `isanytests` | `testCases.go` | Multi-line type verification |
| `integratedtests` | `GetAssert_testcases.go` | Multi-line assertion output |
| `integratedtests` | `GetAssert_ToStrings_testcases.go` | Formatted conversion output |
| `integratedtests` | `GetAssert_SimpleTestCasesWrapper_testcases.go` | Multi-line wrapper output |
| `corejsontests` | `Deserializer_Apply_testcases.go` | JSON comparison output |
| `corejsontests` | `Result_IsEqual_testcases.go` | Equality comparison |
| `corepayloadtests` | `TypedCollectionPagingEdge_testcases.go` | Variable paging |

---

## 🟡 Using Plain `string` / Other Expectations (~21 files)

Single-value expectations stored as bare strings or other simple types. **Low priority** — already simple and readable.

---

## Migration Progress

```
Migrated ████████░░░░░░░░░░░░ 27/134 (20.1%)
Tuples   ██████████░░░░░░░░░░ 36/134 (26.9%)
[]string ████████████████░░░░ 50/134 (37.3%)
Other    ████░░░░░░░░░░░░░░░░ 21/134 (15.7%)
```

### Changelog

| Date | Change |
|------|--------|
| 2026-03-06 | +8 migrated (`MapAnyItems*`, 6 new `coreoncetests`) — total 27 |
| 2026-03-06 | Split Wave 1 into 31 actionable files with difficulty ratings |
| 2026-03-06 | Added Wave 3 "keep as-is" with 15 files + justifications |
| 2026-03-06 | Initial audit — 19 migrated, 52 `[]string` |

---

## Migration Priority

### Wave 1 — Quick Wins (🟢 Easy, 20 files)
All single/dual-boolean or 1–3 value expectations. Estimated: ~1 hour total.

### Wave 2 — Typed Tuples → `args.Map` (36 files)
`args.Two`–`args.Six` → `args.Map` with semantic keys. Higher effort but significant diagnostic improvement.

### Wave 3 — Keep As-Is (15 files)
Variable-length output, multi-line error messages, formatted type inspection. Not suitable for `args.Map`.

---

## Related Docs

- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md) — `args.Map` mandate and patterns
- [Testing Patterns](/spec/01-app/13-testing-patterns.md) — AAA pattern and `CaseV1` usage
- [Edge-Case Coverage Audit](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md)
- [GoConvey Migration Plan](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)
