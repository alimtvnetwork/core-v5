# Testing Roadmap — Comprehensive Coverage Plan

## Status: 🟡 IN PROGRESS

## Summary

This document outlines the prioritized plan for achieving full integration test coverage across all packages, with emphasis on critical functions, branch coverage, and positive/negative/boundary cases.

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

## Phase 2 — 🔲 High Priority: Recently Fixed Functions

These functions had bugs fixed and need test coverage to prevent regression.

### 2.1 `converters/stringsTo` Tests
| Function | Cases Needed | Coverage Targets |
|---|---|---|
| `IntegersWithDefaults` | 5+ | valid input, invalid input with defaults, mixed valid/invalid, empty input, all-invalid |
| `CloneIf(true)` | 4+ | clone=true returns copy, clone=false returns original, empty slice, verify independence |
| `PtrOfPtrToPtrStrings` | 4+ | valid ptrs, nil element in slice, nil outer ptr, empty slice |
| `BytesWithDefaults` | 4+ | valid bytes, out-of-range (>255), negative, non-numeric |

### 2.2 `converters/anyItemConverter` Tests
| Function | Cases Needed | Coverage Targets |
|---|---|---|
| `ToNonNullItems` | 4+ | skipOnNil=true with nil, skipOnNil=false with nil, valid input, reflect-nil interface |

### 2.3 `PayloadsCollection` Paging Tests
| Function | Cases Needed | Coverage Targets |
|---|---|---|
| `GetPagesSize` | 3+ | positive, zero page size, zero length |
| `GetSinglePageCollection` | 4+ | page 1, last page (partial), length < page size, exact fit |
| `GetPagedCollection` | 3+ | multi-page split, single page, concurrent correctness |

### 2.4 `IdentifiersWithGlobals` Tests
| Function | Cases Needed | Coverage Targets |
|---|---|---|
| `GetById` | 4+ | found, not found, empty search, empty collection |
| `Length/IsEmpty` | 3+ | positive, empty, nil receiver |
| `Clone` | 2+ | positive, empty |
| `Add/Adds` | 3+ | single add, multi-add, empty string skipped |

---

## Phase 3 — 🔲 Medium Priority: Missing Test Packages

These packages have **no** integrated test directory at all:

| Package | Complexity | Priority | Estimated Cases |
|---|---|---|---|
| `coredata/corestr` | HIGH | ⭐⭐⭐ | 30+ (Hashset, Collection, SimpleSlice, LinkedList, KeyValuePair) |
| `coredata/stringslice` | HIGH | ⭐⭐⭐ | 25+ (Clone, First/Last, SafeIndex, NonEmpty, Process) |
| `coredata/coreonce` | MEDIUM | ⭐⭐ | 15+ (StringOnce, IntegerOnce, BoolOnce — verify once-semantics) |
| `osconsts` | LOW | ⭐ | 5+ (architecture detection, platform flags) |
| `internal/strutilinternal` | MEDIUM | ⭐⭐ | 8+ (Clone, SliceToMapConverter, IsNullOrEmpty) |

---

## Phase 4 — 🔲 Lower Priority: Expand Existing Thin Tests

| Test Suite | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `coremathtests` | ~4 | 15+ | `integerOutOfRange` all methods, `integer64OutOfRange`, boundary values |
| `errcoretests` | ~3 | 10+ | `RawErrorType` methods (.Error, .Fmt, .MergeError), `SliceToError`, `ManyErrorToSingle` |
| `conditionaltests` | unknown | 10+ | typed conditionals, nil defaults, function execution |
| `corecmptests` | unknown | 10+ | integer/string comparison, pointer comparisons, ordering |

---

## Testing Standards (per project guidelines)

1. **AAA Pattern:** Arrange → Act → Assert in every test
2. **Table-Driven:** Use `coretestcases.CaseV1` with `args.Map`
3. **File Separation:** `_testcases.go` for data, `_test.go` for logic
4. **Error Handling:** Never ignore `args.Map.GetAs*` errors — use `errcore.HandleErrMessage`
5. **No Branching in Tests:** Each scenario = one test case row
6. **Coverage Targets per function:**
   - ✅ Positive (happy path)
   - ❌ Negative (invalid input, not found)
   - 🔲 Boundary (zero, nil, empty, max values)
   - 🔲 Guard clauses (nil receiver, division by zero)

---

## Execution Order

1. **Phase 2** — Cover all recently-fixed functions (regression prevention)
2. **Phase 3** — Create test suites for `corestr` and `stringslice` (highest-complexity untested packages)
3. **Phase 4** — Expand thin test suites to full branch coverage
