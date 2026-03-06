# `args.Map` ExpectedInput Migration Status

> **Last updated:** 2026-03-06

## Summary

| Category | Count | % of Total |
|----------|-------|------------|
| ✅ Migrated to `args.Map` | **19 files** | 15.8% |
| 🔶 Using `args.Two`–`args.Six` (typed tuples) | **36 files** | 30.0% |
| 🟡 Using plain `string` | **53 files** | 44.2% |
| 🔴 Using `[]string` | **52 files** | 43.3% |
| **Total testcase files** | **~120** | — |

> Note: Some files use multiple patterns (e.g., `[]string` for some cases, `string` for others), so percentages overlap.

---

## ✅ Fully Migrated to `args.Map` (19 files)

| Package | File | Notes |
|---------|------|-------|
| `bytetypetests` | `Variant_testcases.go` | Reference implementation |
| `chmodhelpertests` | `PartialRwxVerify_testcases.go` | Simple boolean expectations |
| `coreoncetests` | `IntegerOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `StringOnce_testcases.go` | 5 test case arrays |
| `coreoncetests` | `BoolOnce_testcases.go` | 3 test case arrays |
| `coreoncetests` | `BytesOnce_testcases.go` | 4 test case arrays |
| `coreoncetests` | `BytesErrorOnce_testcases.go` | 13 test case arrays |
| `coreoncetests` | `ErrorOnce_testcases.go` | 6 test case arrays |
| `corepayloadtests` | `TypedCollection_testcases.go` | Collection operations |
| `corepayloadtests` | `TypedCollectionFlatMap_testcases.go` | FlatMap operations |
| `corepayloadtests` | `PayloadWrapper_testcases.go` | Complex struct assertions |
| `corepayloadtests` | `PagingInfo_testcases.go` | Paging metadata |
| `coreversiontests` | `ComparisonExtended_testcases.go` | Version comparison |
| `coreversiontests` | `Comparison_testcases.go` | Version comparison |
| `coreversiontests` | `Parse_testcases.go` | Version parsing |
| `coreversiontests` | `String_testcases.go` | Version string output |
| `pagingutiltests` | `Paging_testcases.go` | Paging calculations |
| `corepayloadtests` | `TypedCollectionGroupBy_testcases.go` | GroupBy operations |
| `corepayloadtests` | `TypedCollectionPartition_testcases.go` | Partition operations |

---

## 🔶 Using Typed Tuples `args.Two`–`args.Six` (36 files)

These use positional typed tuples — better than `[]string` but lack semantic keys. **Migration to `args.Map` is recommended** for improved diagnostics.

| Package | File |
|---------|------|
| `coreapitests` | `PageRequest_testcases.go` |
| `coreapitests` | `SearchRequest_testcases.go` |
| `coreapitests` | `RequestAttribute_testcases.go` |
| `coreapitests` | `ResponseAttribute_testcases.go` |
| `converterstests` | `StringsTo_testcases.go` |
| `converterstests` | `StringTo_testcases.go` |
| `coredynamictests` | `MapAnyItems_testcases.go` |
| `coregenerictests` | `CollectionBranch_testcases.go` |
| `coregenerictests` | `CollectionSerialization_testcases.go` |
| `coregenerictests` | `HashsetBranch_testcases.go` |
| `corestrtests` | `Collection_testcases.go` |
| `corestrtests` | `Hashmap_testcases.go` |
| `corestrtests` | `Hashset_testcases.go` |
| `corestrtests` | `ValidValue_testcases.go` |
| `keymktests` | `KeyLegend_testcases.go` |
| `namevaluetests` | `Collection_testcases.go` |
| `regexnewtests` | `LazyRegex_testcases.go` |
| `simplewraptests` | `Wrapper_testcases.go` |
| `stringcompareastests` | `Glob_testcases.go` |
| `typesconvtests` | `Bool_testcases.go` |
| `typesconvtests` | `Byte_testcases.go` |
| `typesconvtests` | `Float_testcases.go` |
| `typesconvtests` | `Integer_testcases.go` |
| `typesconvtests` | `String_testcases.go` |
| *(and ~12 more across various packages)* | |

---

## 🔴 Using `[]string` Expectations (52 files)

These use raw `[]string` slices with positional semantics. **Highest priority for migration** — failure messages are opaque.

### Migratable to `args.Map` (simple boolean/value expectations)

| Package | File | Fields | Difficulty |
|---------|------|--------|------------|
| `defaulterrtests` | `DefaultErr_testcases.go` | 2 booleans | 🟢 Easy |
| `issettertests` | `Value_testcases.go` | 2–8 booleans | 🟡 Medium (many fields) |
| `coregenerictests` | `HashsetBranch_testcases.go` | 1–4 values | 🟢 Easy |
| `coregenerictests` | `Hashmap_testcases.go` | 1–3 values | 🟢 Easy |
| `coregenerictests` | `LinkedList_testcases.go` | 1–3 values | 🟢 Easy |
| `coregenerictests` | `SimpleSlice_testcases.go` | 1–4 values | 🟢 Easy |
| `corestrtests` | `SimpleSlice_testcases.go` | 1–3 values | 🟢 Easy |
| `enumimpltests` | `BasicEnum_testcases.go` | 2–3 values | 🟢 Easy |
| `reqtypetests` | `Type_testcases.go` | 1–2 values | 🟢 Easy |
| `ostypetests` | `Type_testcases.go` | 1–2 values | 🟢 Easy |
| `anycmptests` | `QuickCompare_testcases.go` | 1 value | 🟢 Easy |

### Keep as `[]string` (variable-length or formatted error messages)

| Package | File | Reason |
|---------|------|--------|
| `chmodhelpertests` | `VerifyRwxChmod*_testcases.go` | Multi-line error messages |
| `chmodhelpertests` | `VerifyPartialRwxLocations_testcases.go` | Variable error output |
| `chmodhelpertests` | `DirFilesWithContent_testcases.go` | Dynamic file listings |
| `chmodhelpertests` | `SimpleFileWriter_testcases.go` | File operation results |
| `chmodhelpertests` | `RwxCompileValue_testcases.go` | Formatted rwx strings |
| `errcoretests` | `MergeErrors_testcases.go` | Variable error merge output |
| `corejsontests` | `Serialize_testcases.go` | JSON string output |
| `stringslicetests` | `*_testcases.go` (many) | Variable-length slice results |
| `corevalidatortests` | `*_testcases.go` | Multi-line validation output |
| `corestrtests` | `*_testcases.go` (several) | Collection operation results |
| `coreappendtests` | `Append_testcases.go` | Dynamic append results |

---

## 🟡 Using Plain `string` Expectations (53 files)

Single-value expectations stored as bare strings. These are **low priority** — they're already simple and readable.

| Pattern | Example | Count |
|---------|---------|-------|
| Boolean string | `ExpectedInput: "true"` | ~30 files |
| Numeric string | `ExpectedInput: "42"` | ~15 files |
| Value string | `ExpectedInput: "hello"` | ~8 files |

---

## Migration Priority

### Wave 1 — Quick Wins (🟢 Easy, high diagnostic value)
1. `defaulterrtests/DefaultErr_testcases.go` — 11 cases, 2 boolean fields each
2. `coregenerictests/HashsetBranch_testcases.go` — many single-var cases
3. `coregenerictests/Hashmap_testcases.go` — many single-var cases
4. `issettertests/Value_testcases.go` — 8-field boolean matrix → very high diagnostic value
5. `enumimpltests/BasicEnum_testcases.go` — simple enum checks

### Wave 2 — Typed Tuples → `args.Map`
6. `typesconvtests/*.go` — 5 files, consistent Ptr/PtrToSimple pattern
7. `stringcompareastests/Glob_testcases.go` — `args.Two` → `args.Map{isMatch, isInverse}`
8. `namevaluetests/Collection_testcases.go` — `args.Two`/`args.Three` → semantic keys
9. `coreapitests/*.go` — `args.Two` → `args.Map` for request/response fields
10. `converterstests/*.go` — `args.Five` → `args.Map` with named fields

### Wave 3 — Skip / Keep As-Is
- `stringslicetests/` — 80+ pure functions with variable-length output → keep `[]string`
- `chmodhelpertests/` error-message tests → keep `[]string`
- Plain `string` single-value tests → no migration needed

---

## Related Docs

- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md) — `args.Map` mandate and patterns
- [Testing Patterns](/spec/01-app/13-testing-patterns.md) — AAA pattern and `CaseV1` usage
- [Edge-Case Coverage Audit](/spec/13-app-issues/testing/02-edge-case-coverage-audit.md)
- [GoConvey Migration Plan](/spec/13-app-issues/testing/04-goconvey-migration-plan.md)
