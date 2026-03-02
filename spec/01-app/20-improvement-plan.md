# Improvement Plan — Phased Roadmap

> A prioritized, phase-by-phase plan for modernizing the `core` framework.

## Phase 1: Foundation ✅ COMPLETE

### 1.1 Complete `interface{}` → `any` Migration
- **Status**: ~80% done (corepayload, coredynamic creators migrated)
- **Remaining**: `coreinterface/` (569 matches in 27 files — large interfaces, defer to dedicated session)
- **Effort**: ~1-2 sessions remaining
- **Risk**: Low (alias, no behavioral change)

### 1.2 Fix Known Bugs ✅
- [x] `PayloadWrapper.IsIdentifier` compares `it.Name` instead of `it.Identifier` — **FIXED**
- [x] Remaining `convertinteranl` / `refeflectcore` typo references — **FIXED** (standardized to `convertinternal`/`reflectcore`)
- **Effort**: Done

### 1.3 Complete Go Version Update ✅
- [x] Updated `go.mod` to Go 1.24
- [x] Updated `makefile` GoVersion
- **Effort**: Done

---

## Phase 2: Generics — Core Collections ✅ COMPLETE

### 2.1 `Collection[T]` — Created in `coredata/coregeneric/`
- Thread-safe generic collection with Map, Filter, Reduce, GroupBy, Distinct
- `Hashset[T]`, `Hashmap[K,V]`, `SimpleSlice[T]`, `LinkedList[T]` all created
- Generic Typed Creator pattern with 16 primitive types

### 2.2 Pre-Built Type Aliases ✅
- `CollectionTypes.go` with String, Int, Byte, Bool, Float64, etc.

### 2.3 Backward Compatibility ✅
- `DynamicCollection` and `AnyCollection` remain with deprecation notices

---

## Phase 3: Generics — Payload & Dynamic Types ✅ COMPLETE

### 3.1 Generic Deserialize Helpers ✅
- `DeserializePayloadTo[T]`, `DeserializePayloadToSlice[T]`, `DeserializePayloadToMust[T]`
- `DeserializeAttributesPayloadTo[T]`, `DeserializeAttributesPayloadToMust[T]`
- All in `coredata/corepayload/generic_helpers.go`

### 3.2 TypedPayloadWrapper[T] ✅
- Full API: 40+ methods including GetAs*, Value*, JSON, Clone, Setters
- Factory functions: `TypedPayloadWrapperFrom`, `TypedPayloadWrapperRecord`, `TypedPayloadWrapperAll`, etc.
- Deserialization: `TypedPayloadWrapperDeserialize[T]`, `TypedPayloadWrapperDeserializeToMany[T]`

### 3.3 TypedDynamic[T], TypedSimpleRequest[T], TypedSimpleResult[T] ✅
- Full parity with legacy Dynamic/SimpleRequest/SimpleResult
- GetAs* (String/Int/Int64/Float64/Float32/Bool/Bytes/Strings)
- Value* convenience methods with safe defaults
- JSON operations, Clone, conversion to legacy types

### 3.4 `interface{}` → `any` Migration in corepayload ✅
- `newPayloadWrapperCreator.go` — all `interface{}` → `any`
- `newAttributesCreator.go` — all `interface{}` → `any`
- `PayloadCreateInstructionTypeStringer.go` — `interface{}` → `any`

---

## Phase 4: Test Coverage Expansion ✅ IN PROGRESS

### Priority Order (by risk/usage):

| Priority | Package | Status |
|----------|---------|--------|
| P0 | `conditional/` | ✅ 8 test functions, 30+ test cases |
| P0 | `errcore/` | ✅ Expanded: 5 test functions, 10+ test cases, panic tests |
| P0 | `converters/` | ✅ Expanded: 9 test functions, 35+ test cases (added Bool, UnsignedInteger) |
| P1 | `coretaskinfo/` | Tests exist |
| P1 | `coredata/corepayload/` | Coverage via demo cmd |
| P1 | `regexnew/` | Tests exist |
| P2 | `coremath/` | Tests exist |
| P2 | `coresort/` | Tests exist |
| P2 | `coreutils/` | Tests exist |
| P3 | `mutexbykey/` | Tests exist |
| P3 | `namevalue/` | Tests exist |
| P3 | `pagingutil/` | Tests exist |
| P3 | `typesconv/` | Tests exist |
| P3 | `coreappend/` | Tests exist |
| P3 | `coreunique/` | Tests exist |

---

## Phase 5: Refactoring Large Files — IN PROGRESS

| File | Lines | Action | Status |
|------|-------|--------|--------|
| `PayloadWrapper.go` | 842→817 | Split: extracted `PayloadWrapperGetters.go`, `PayloadWrapperJson.go` | ✅ Started |
| `Attributes.go` | 768 | Split planned: `AttributesGetters.go`, `AttributesSetters.go` | Pending |
| `Info.go` | 646 | Split planned | Pending |
| `DynamicCollection.go` | 636 | Deprecated, replaced by `Collection[T]` | ✅ Done |
| `AnyCollection.go` | 707 | Deprecated, replaced by `Collection[any]` | ✅ Done |
| `Dynamic.go` | 674 | Split planned: `DynamicGetters.go`, `DynamicReflect.go` | Pending |
| `BaseTestCase.go` | 435 | Split planned | Pending |

---

## Phase 6: Value Receiver Migration — PLANNED

Migrate read-only methods from pointer to value receivers, package by package:
- Start with small packages (`coreversion/`, `issetter/`)
- Graduate to larger packages (`coretaskinfo/`, `corepayload/`)
- Always verify interface satisfaction after changes
- Initial migration done for `PayloadWrapperJson.go` (value receivers)

**Effort**: Ongoing, 1-2 files per session alongside other work

---

## Phase 7: Expert Code Review Fixes ✅ COMPLETE

> Identified via comprehensive code review — 16 findings across 4 sub-phases.

### 7.1 Critical Bugs ✅

- [x] `coreinstruction/BaseRequestIds.go` — Inverted nil guard in `RequestIdsLength()` caused panic on nil receiver — **FIXED**
- [x] `ostype/Variation.go` — `IsOpenBsd()` compared against `NetBsd` instead of `OpenBsd` — **FIXED**
- [x] `reqtype/RangesInBetween.go` — Off-by-one: `i < endVal` excluded last element, changed to `i <= endVal` — **FIXED**
- [x] `coreindexes/HasIndexPlusRemoveIndex.go` — `RemoveIndex` mutated local slice copy; changed to `*[]int` — **FIXED**

### 7.2 Code Quality & Style ✅

- [x] `coreinstruction/BaseTypeDotFilter.go` — Value receiver on `GetDotSplitTypes()` made caching a no-op; changed to pointer receiver — **FIXED**
- [x] `coreinstruction/SourceDestination.go`, `FromTo.go`, `Rename.go` — Removed unreachable `IsNull()` checks on value receivers — **FIXED**
- [x] `coreinstruction/SourceDestination.go`, `FromTo.go`, `Rename.go` — Fixed `form` → `from` typo in `SetFromName` — **FIXED**
- [x] `coreinstruction/LineIdentifier.go` — Missing parentheses in `IsDeleteLineRequest()` caused wrong operator precedence — **FIXED**
- [x] `coreappend/PrependAppendAnyItemsToStringsUsingFunc.go` — Collapsed redundant `if/else-if` into single condition — **FIXED**
- [x] `reqtype/start.go`, `end.go` — Removed unnecessary parentheses — **FIXED**
- [x] `coreinstruction/BaseIdentifier.go`, `BaseUsername.go`, `BaseDisplay.go` — Removed trailing blank lines — **FIXED**

### 7.3 Design Improvements ✅

- [x] `coreinstruction/BaseContinueOnError.go` — Consolidated into `BaseIsContinueOnError.go`, deleted redundant type — **FIXED**
- [x] `reqtype/start.go`, `end.go` — Changed return type from `any` to `*Request` for type safety — **FIXED**
- [x] `coreinstruction/BaseRequestIds.go` — `NewRequestIds` now returns `[]IdentifierWithIsGlobal` instead of `*[]IdentifierWithIsGlobal` — **FIXED**
- [x] `coreinstruction/ById.go` vs `BaseIdentifier.go` — Kept both: different JSON tags (`omitempty` vs not) are intentional — **NO CHANGE (by design)**

### 7.4 Minor Cleanup ✅

- [x] `defaulterr/defaulterr.go` — `MarshallingFailedDueToNilOrEmpty` used wrong error type (`UnMarshallingFailedType`); changed to `MarshallingFailedType` — **FIXED**

---

## Phase 8: Deep Quality Sweep ✅ COMPLETE

> Systematic codebase-wide sweep covering inline negation refactoring, defensive fixes, and regression test coverage.

### 8.1 Inline Negation Refactoring ✅

Refactored **~190 inline negations** across **~45 files** to use named boolean variables:

| Pass | Scope | Files | Fixes |
|------|-------|-------|-------|
| 1 | `coredata/` core types | 22 | ~65 |
| 2 | `coredata/` remaining | 15 | ~30 |
| 3 | `internal/`, `coreutils/`, other packages | 30+ | ~95 |

**Pattern**: `if !condition {` → `isNegated := !condition` + `if isNegated {`

**Packages covered**: `coredata/corestr/`, `coredata/coredynamic/`, `coredata/coregeneric/`, `coredata/corejson/`, `coredata/corepayload/`, `coredata/coreonce/`, `coredata/stringslice/`, `internal/reflectinternal/`, `internal/strutilinternal/`, `internal/mapdiffinternal/`, `coreutils/stringutil/`, `codestack/`, `simplewrap/`, `chmodhelper/`, `namevalue/`, `errcore/`, `conditional/`, `issetter/`, `coreimpl/enumimpl/`, `coretests/`

**Remaining exceptions**: `coretests/args/Map.go` — `!ok` guard clauses in type-assertion getters are standard Go idiom and intentionally preserved.

### 8.2 Low-Priority Bug Fixes ✅

- [x] `coredata/corestr/Hashmap.Clear()` — Added nil receiver guard + nil check on `cachedList` before slicing — **FIXED**
- [x] `coredata/corestr/Hashset.AddBool()` — Added `it.hasMapUpdated = true` when new item is added to invalidate cached data — **FIXED**
- [x] `coredata/corestr/Hashmap.AddOrUpdateCollection()` — Added length mismatch guard returning early if `keys` and `values` have different lengths — **FIXED**

### 8.3 Regression Tests ✅

Added **10 regression test cases** for the three low-priority fixes:

| Fix | Test Cases | Covers |
|-----|-----------|--------|
| `Hashmap.Clear` nil safety | 3 | nil receiver, populated clear, chainability |
| `Hashset.AddBool` cache invalidation | 3 | new item cache rebuild, existing item no-op, multiple additions |
| `AddOrUpdateCollection` length mismatch | 4 | mismatched lengths, equal lengths, nil keys, empty keys |

### 8.4 PairFromSplit / TripleFromSplit Tests ✅

Added **30 test cases** across 4 new files covering all split constructors:

| Function | Cases |
|----------|-------|
| `PairFromSplit` | 7 (standard, no sep, empty, multi-sep, sep at start/end, multi-char) |
| `PairFromSplitTrimmed` | 2 |
| `PairFromSplitFull` | 3 |
| `PairFromSplitFullTrimmed` | 2 |
| `PairFromSlice` | 4 (two-element, single, empty, three-element) |
| `TripleFromSplit` | 6 (standard, no sep, two parts, four parts, empty, multi-char) |
| `TripleFromSplitTrimmed` | 1 |
| `TripleFromSplitN` | 3 |
| `TripleFromSplitNTrimmed` | 1 |
| `TripleFromSlice` | 4 (three-element, empty, single, two-element) |

---

## Summary Timeline

| Phase | Focus | Sessions | Status |
|-------|-------|----------|--------|
| 1 | Foundation (any, bugs, Go version) | 3-4 | ✅ Complete |
| 2 | Generic Collection[T] | 3-4 | ✅ Complete |
| 3 | Generic Payload/Dynamic helpers | 2 | ✅ Complete |
| 4 | Test coverage | 8-10 | ✅ P0 Complete |
| 5 | File splitting | 2-3 | 🔄 Started |
| 6 | Value receivers | Ongoing | 🔄 Started |
| 7 | Expert code review fixes | 1 | ✅ Complete |
| 8 | Deep quality sweep | 1 | ✅ Complete |

## Remaining Work

### `interface{}` → `any` Migration

- `coreinterface/` — 569 matches in 27 files (largest remaining surface)
- `coredata/coredynamic/Dynamic.go` — ~20 matches
- Various other packages with scattered usage

### File Splitting

- `Attributes.go` (768 lines)
- `Dynamic.go` (674 lines)
- `Info.go` (646 lines)
- `BaseTestCase.go` (435 lines)
