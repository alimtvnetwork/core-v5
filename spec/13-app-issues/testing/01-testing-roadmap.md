# Testing Roadmap — Comprehensive Coverage Plan

## Status: ✅ COMPLETE

## Summary

This document outlined the prioritized plan for achieving full integration test coverage across all packages. **All phases are now complete**, with ~870 new test cases added through the deep coverage scan initiative.

---

## Phase 1 — ✅ Completed: Fix Broken Tests & Expand Critical Coverage

### 1.1 Paging Tests (`pagingutiltests`)

- **Fixed:** `GetPagingInfo` test expected wrong `EndingLength` after bug fix
- **Added:** 10 `GetPagesSize` cases (positive, zero items, zero/negative page size)
- **Added:** 9 `GetPagingInfo` cases (multi-page, last page clamping, not-pageable, exact fit, zero length)

### 1.2 Core Instruction Tests (`coreinstructiontests`)

- **Added:** `Identifiers.Length()` — 3 cases (positive, empty, single)
- **Added:** `Identifiers.GetById()` — 6 cases (found first/middle/last, not found, empty search, empty collection)
- **Added:** `Identifiers.IndexOf()` — 5 cases (found, first, missing, empty search, empty collection)
- **Added:** `Identifiers.Clone()` — 2 cases (positive, empty)
- **Added:** `Identifiers.Add()` — 2 cases (positive, skip empty)
- **Added:** `Specification.Clone()` — 2 cases (all fields, empty tags) + nil safety + deep-copy verification
- **Added:** `BaseTags` — 4 cases (all match, partial, empty-empty, empty-nonempty)

---

## Phase 2 — ✅ Completed: Recently Fixed Functions & Core Packages

### 2.1 `converters/stringsTo` Tests ✅

- `IntegersWithDefaults`, `CloneIf`, `PtrOfPtrToPtrStrings`, `BytesWithDefaults` — covered

### 2.2 `converters/anyItemConverter` Tests ✅

- `ToNonNullItems` — covered

### 2.3 `PayloadsCollection` Paging Tests ✅

- `GetPagesSize`, `GetSinglePageCollection`, `GetPagedCollection` — covered

### 2.4 `IdentifiersWithGlobals` Tests ✅

- `GetById`, `Length/IsEmpty`, `Clone`, `Add/Adds` — covered

### 2.5 Deep Coverage Scan Additions ✅

- `issetter` — 45 tests (6-value boolean logic)
- `coredynamic` — 50 tests (LeftRight, MapAnyItems, Dynamic, CastedResult)
- `coreonce` — 125 tests (StringOnce, BoolOnce, ErrorOnce, BytesOnce, BytesErrorOnce)
- `coreinstruction` — 40 tests (IdentifiersWithGlobals, FromTo)
- `coregeneric` — 175 tests (LinkedList, Hashmap, Hashset, functional ops)
- `corestr` — 115 tests (Hashset, Hashmap, bug 42 caching verification)
- `corevalidator` — 208 tests (Parameter, TextValidator, LineValidator, SliceValidator, SliceValidators, SimpleSliceValidator, LinesValidators, BaseLinesValidators)

---

## Phase 3 — ✅ Completed: Missing Test Packages

| Package | Status | Tests Added |
|---|---|---|
| `coredata/corestr` | ✅ DONE | ~115 (Hashset, Hashmap) |
| `coredata/stringslice` | ✅ DONE | ~40 (Clone, First/Last, SafeIndex, NonEmpty, Reverse, etc.) |
| `coredata/coreonce` | ✅ DONE | ~125 (all Once types) |
| `reflectcore` | ✅ DONE | ~12 (facade export verification) |
| `coreappend` | ✅ DONE | ~20 (string assembly with nil skipping) |

---

## Phase 4 — ✅ Completed: Expanded Coverage via Deep Scan

The deep coverage scan initiative (documented in `03-deep-coverage-scan.md`) systematically expanded test coverage across all remaining packages, surpassing the original Phase 4 targets.

---

## Testing Standards (per project guidelines)

1. **AAA Pattern:** Arrange → Act → Assert in every test
2. **Table-Driven:** Use `coretestcases.CaseV1` with `args.Map` for complex scenarios
3. **File Separation:** `_testcases.go` for data, `_test.go` for logic
4. **Error Handling:** Never ignore `args.Map.GetAs*` errors — use `errcore.HandleErrMessage`
5. **No Branching in Tests:** Each scenario = one test case row
6. **Coverage Targets per function:**
   - ✅ Positive (happy path)
   - ✅ Negative (invalid input, not found)
   - ✅ Boundary (zero, nil, empty, max values)
   - ✅ Guard clauses (nil receiver, division by zero)

---

## Final Results

| Metric | Value |
|--------|-------|
| **New test cases added** | ~870 |
| **Packages with logic covered** | 100% |
| **Bugs found during testing** | 1 (`BytesErrorOnce.Deserialize` silent failure) |
| **Phases completed** | 4/4 |

## Related Docs

- [Edge Case Coverage Audit](./02-edge-case-coverage-audit.md)
- [Deep Coverage Scan](./03-deep-coverage-scan.md)
- [Testing Patterns](../01-app/13-testing-patterns.md)
