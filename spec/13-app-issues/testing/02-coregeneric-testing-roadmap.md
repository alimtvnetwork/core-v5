# Testing Roadmap — Comprehensive Coverage Plan

## Status: 🟡 IN PROGRESS (Phase 2 complete)

## Summary

This document outlines the phased plan for achieving full test coverage across all `coregeneric` types, with emphasis on the new `Pair` and `Triple` types and existing data structures.

---

## Phase 1 — ✅ Completed: Pair & Triple Foundation

### 1.1 Pair[L, R] Tests

- **Added:** `NewPair` — 2 cases (valid with values, valid with empty strings)
- **Added:** `InvalidPair` / `InvalidPairNoMessage` — 2 cases
- **Added:** `Clone` — 2 cases (independence verification, nil safety)
- **Added:** `IsEqual` — 4 cases (equal, unequal left, nil vs non-nil, both nil)
- **Added:** `Values()` — 1 case
- **Added:** `Clear/Dispose` — 1 case (zero values after clear)
- **Added:** `New.Pair` Creator — 4 shortcut tests (StringString, StringInt, Any, InvalidStringString)

### 1.2 Triple[A, B, C] Tests

- **Added:** `NewTriple` — 1 case (valid with values)
- **Added:** `InvalidTriple` / `InvalidTripleNoMessage` — 2 cases
- **Added:** `Clone` — 2 cases (independence, nil safety)
- **Added:** `Values()` — 1 case
- **Added:** `Clear/Dispose` — 1 case
- **Added:** `New.Triple` Creator — 3 shortcut tests (StringStringString, Any, InvalidStringStringString)

---

## Phase 2 — ✅ Completed: Regression Tests for Fixed Functions

### 2.1 converters Package (`tests/integratedtests/converterstests/`)

- **Added:** `IntegersWithDefaults` — 5 cases (all valid, mixed valid/invalid, empty, all invalid, negatives)
- **Added:** `BytesWithDefaults` — 5 cases (valid bytes, >255 overflow, negative, non-numeric, empty)
- **Added:** `CloneIf` — 3 cases (clone true, clone false, empty input)
- **Added:** `PtrOfPtrToPtrStrings` — 4 cases (valid, nil entries, nil outer, nil inner)
- **Added:** `PtrOfPtrToMapStringBool` — 4 cases (valid, nil entries skipped, nil input, empty)
- **Added:** `ToNonNullItems` — 3 cases (nil with skip, valid slice, nil without skip)

### 2.2 stringslice Package (`tests/integratedtests/stringslicetests/`)

- **Added:** `CloneIf` — 4 cases (clone with extra cap, return original, nil + false, nil + true)
- **Added:** `AnyItemsCloneIf` — 2 cases (clone true, return original)

### 2.3 PayloadsCollection Paging Edge Cases (`tests/integratedtests/corepayloadtests/`)

- **Added:** `GetPagedCollection` — 4 edge cases (single item, page size 1, exact count, oversized page)
- **Added:** `GetSinglePageCollection` — 2 edge cases (middle page, page size 1)
- **Added:** `GetPagedCollectionWithInfo` — 2 edge cases (single item metadata, exact division metadata)
- **Added:** `GetPagesSize` — 2 edge cases (page size 1, single item)
- **Added:** `PagingWithInfo_Empty` — 1 case (empty collection with info)

---

## Phase 2b — ✅ Completed: Pair & Triple Extended Coverage

### Tests Added (`PairTripleExtended_testcases.go` / `PairTripleExtended_test.go`)

- **Pair.IsEqual extended** — 6 cases: same values diff validity, different right, both invalid zero, Pair[int,int] same/diff, Pair[string,int] mixed
- **Pair.HasMessage** — 4 cases: no message, with message, whitespace-only, nil receiver
- **Pair.IsInvalid** — 3 cases: valid, invalid, nil receiver
- **Pair.String** — 4 cases: valid string pair, invalid zero, nil, Pair[string,int]
- **Pair.NewPairWithMessage** — 2 cases: valid with msg, invalid with error
- **Pair.Dispose** — 1 case: resets same as Clear
- **Pair nil Clear** — 1 case: no panic on nil receiver
- **Triple.IsEqual extended** — 5 cases: same, diff validity, diff middle, both nil, nil vs non-nil
- **Triple.HasMessage** — 3 cases: no msg, with msg, nil
- **Triple.IsInvalid** — 3 cases: valid, invalid, nil
- **Triple.String** — 3 cases: valid, invalid zero, nil
- **Triple.NewTripleWithMessage** — 2 cases: valid, invalid
- **Triple.Dispose** — 1 case: resets same as Clear
- **Triple nil Clear** — 1 case: no panic on nil receiver
- **New.Pair all shortcuts** — 7 creators: StringInt64, StringFloat64, StringBool, StringAny, IntInt, IntString, InvalidAny
- **New.Triple all shortcuts** — 3 creators: StringIntString, StringAnyAny, InvalidAny

**Total: 49 new test cases across 18 test functions**

---

## Phase 3 — 🔲 Collection[T] Full Branch Coverage

| Function | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `Add/Adds/AddSlice` | ~3 | 8+ | AddIf false, AddIfMany, AddFunc, AddCollection nil |
| `RemoveAt` | 0 | 4+ | valid index, negative index, out-of-bounds, single item |
| `First/Last/FirstOrDefault/LastOrDefault` | ~2 | 6+ | empty collection panic, SafeAt OOB, SafeAt valid |
| `Skip/Take` | 0 | 4+ | skip all, take more than length, zero, normal |
| `ForEach/ForEachBreak` | 0 | 3+ | full iteration, early break, empty |
| `Filter/CountFunc` | ~1 | 4+ | no match, all match, empty, predicate |
| `Clone` | ~1 | 3+ | empty, populated, nil |
| `SortFunc/Reverse` | 0 | 4+ | sort ascending/descending, reverse, single element |
| `ConcatNew` | 0 | 3+ | empty + items, items + empty, both populated |
| `Lock variants` | ~1 | 4+ | AddLock, AddsLock, LengthLock, IsEmptyLock |

---

## Phase 4 — 🔲 Hashset[T] Full Branch Coverage

| Function | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `Add/AddSlice/Adds` | ~1 | 4+ | duplicate add, empty slice, AddIf false |
| `Remove/RemoveAll` | ~1 | 4+ | remove existing, remove missing, remove from empty |
| `Has/HasAll/HasAny` | ~1 | 6+ | has existing, has missing, HasAll partial, HasAny none |
| `Clone` | 0 | 2+ | populated, empty |
| `ToSlice/ToSortedSlice` | 0 | 3+ | populated, empty, sorted order |
| `Intersect/Difference/Union` | 0 | 6+ | overlap, disjoint, empty sets |
| `Lock variants` | 0 | 4+ | AddLock, HasLock, RemoveLock, LengthLock |

---

## Phase 5 — 🔲 Hashmap[K,V] Full Branch Coverage

| Function | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `Set/Get/GetOrDefault` | ~1 | 5+ | existing key, missing key, overwrite, GetOrDefault missing |
| `Delete/Has` | 0 | 4+ | delete existing, delete missing, Has existing/missing |
| `Keys/Values/Items` | 0 | 3+ | populated, empty, single item |
| `Clone` | 0 | 2+ | populated, empty |
| `ForEach` | 0 | 2+ | populated, empty |
| `Lock variants` | 0 | 4+ | SetLock, GetLock, DeleteLock, LengthLock |

---

## Phase 6 — 🔲 SimpleSlice[T] & LinkedList[T]

| Type | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `SimpleSlice` Add/Adds | ~1 | 4+ | AddIf false, AddsIf, AddFunc |
| `SimpleSlice` First/Last/OrDefault | 0 | 4+ | empty panic, populated, default on empty |
| `SimpleSlice` Skip/Take | 0 | 4+ | boundaries, normal |
| `SimpleSlice` Filter/Clone | 0 | 4+ | no match, clone independence |
| `SimpleSlice` InsertAt | 0 | 3+ | start, middle, end |
| `LinkedList` Add/AddFront | ~2 | 4+ | empty list add, single item |
| `LinkedList` Remove/RemoveAt | 0 | 4+ | head, tail, middle, out-of-bounds |
| `LinkedList` Find/Contains | 0 | 4+ | found, not found, empty list |
| `LinkedList` Reverse/Clone | 0 | 4+ | populated, empty, independence |

---

## Phase 7 — 🔲 Generic Functions (funcs.go, comparablefuncs.go, orderedfuncs.go)

| Function | Current Cases | Target Cases | Key Gaps |
|---|---|---|---|
| `MapCollection` | ~1 | 3+ | empty, identity map, type-changing map |
| `FlatMapCollection` | 0 | 3+ | normal, empty inner, empty outer |
| `Reduce` | 0 | 3+ | sum, concat, empty collection |
| `GroupBy` | 0 | 3+ | multiple groups, single group, empty |
| `ContainsFunc` | 0 | 3+ | found, not found, empty |
| `ContainsAll` | ~1 | 4+ | all present, partial, empty search, empty collection |
| `Distinct` | ~1 | 3+ | duplicates, unique, empty |
| `Sort/Min/Max/Sum/Clamp` | ~1 | 10+ | ascending, descending, single, empty, boundary values |

---

## Testing Standards

1. **AAA Pattern:** Arrange → Act → Assert in every test
2. **Table-Driven:** Use `coretestcases.CaseV1` with `args.Map`
3. **File Separation:** `_testcases.go` for data, `_test.go` for logic
4. **Error Handling:** Never ignore `args.Map.GetAs*` errors
5. **No Branching in Tests:** Each scenario = one test case row
6. **Coverage Targets per function:**
   - ✅ Positive (happy path)
   - ❌ Negative (invalid input, not found)
   - 🔲 Boundary (zero, nil, empty, max values)
   - 🔲 Guard clauses (nil receiver, empty collection)

---

## Execution Order

1. **Phase 2** — Complete Pair/Triple edge cases
2. **Phase 3** — Collection[T] full coverage (highest-complexity type)
3. **Phase 4** — Hashset[T] (set operations are bug-prone)
4. **Phase 5** — Hashmap[K,V]
5. **Phase 6** — SimpleSlice[T] & LinkedList[T]
6. **Phase 7** — Generic standalone functions
