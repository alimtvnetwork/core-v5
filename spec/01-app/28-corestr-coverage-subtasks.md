# corestr Coverage Sub-Task Breakdown

## Status: 🔄 Active

## Summary

Package: `coredata/corestr`
Total uncovered statements: ~3,917 (across 54 files)
Target: 100% coverage
Segment size: ~200 statements per task

---

## Sub-Tasks (20 segments)

### Segment 1: Collection.go — Part 1 (~200 stmts)
- **File**: `Collection.go` (527 uncovered total)
- **Scope**: First ~200 uncovered statements (basic methods, Add/Remove, Filter, Map)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Collection_Part1_test.go`

### Segment 2: Collection.go — Part 2 (~200 stmts)
- **File**: `Collection.go`
- **Scope**: Next ~200 uncovered statements (Sort, Join, EqualFold, Lock variants)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Collection_Part2_test.go`

### Segment 3: Collection.go — Part 3 (~127 stmts)
- **File**: `Collection.go`
- **Scope**: Remaining ~127 uncovered statements (JSON, String, Summary, remaining methods)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Collection_Part3_test.go`

### Segment 4: LinkedCollections.go (~200 stmts)
- **File**: `LinkedCollections.go` (359 uncovered)
- **Scope**: First ~200 statements (Add, Remove, Loop, Chain methods)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_LinkedCollections_Part1_test.go`

### Segment 5: LinkedCollections.go — Part 2 (~159 stmts)
- **File**: `LinkedCollections.go`
- **Scope**: Remaining ~159 statements (JSON, Equals, Append, remaining)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_LinkedCollections_Part2_test.go`

### Segment 6: Hashset.go (~200 stmts)
- **File**: `Hashset.go` (324 uncovered)
- **Scope**: First ~200 statements (Add, Remove, Contains, Lock variants)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Hashset_Part1_test.go`

### Segment 7: Hashset.go — Part 2 (~124 stmts)
- **File**: `Hashset.go`
- **Scope**: Remaining ~124 statements (JSON, Diff, Merge, remaining)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Hashset_Part2_test.go`

### Segment 8: SimpleSlice.go (~200 stmts)
- **File**: `SimpleSlice.go` (308 uncovered)
- **Scope**: First ~200 statements (Add, Filter, Skip, Take, Sort)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_SimpleSlice_Part1_test.go`

### Segment 9: SimpleSlice.go — Part 2 (~108 stmts)
- **File**: `SimpleSlice.go`
- **Scope**: Remaining ~108 statements (JSON, remaining methods)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_SimpleSlice_Part2_test.go`

### Segment 10: Hashmap.go (~200 stmts)
- **File**: `Hashmap.go` (288 uncovered)
- **Scope**: First ~200 statements (Get, Set, Delete, Keys, Values)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Hashmap_Part1_test.go`

### Segment 11: Hashmap.go — Part 2 (~88 stmts) + HashmapDiff.go (28 stmts)
- **Files**: `Hashmap.go` remaining + `HashmapDiff.go` (28 uncovered)
- **Scope**: Remaining Hashmap + full HashmapDiff coverage
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Hashmap_Part2_test.go`

### Segment 12: LinkedList.go (~200 stmts)
- **File**: `LinkedList.go` (255 uncovered)
- **Scope**: First ~200 statements (Push, Pop, Remove, Loop)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_LinkedList_Part1_test.go`

### Segment 13: LinkedList.go — Part 2 (~55 stmts) + LinkedListNode.go (67 stmts)
- **Files**: `LinkedList.go` remaining + `LinkedListNode.go` (67 uncovered)
- **Scope**: Remaining LinkedList + LinkedListNode methods
- **Test file**: `tests/integratedtests/corestrtests/Coverage_LinkedList_Part2_test.go`

### Segment 14: CharCollectionMap.go (~200 stmts)
- **File**: `CharCollectionMap.go` (247 uncovered)
- **Scope**: First ~200 statements (Get, Add, Lock variants)
- **Test file**: `tests/integratedtests/corestrtests/Coverage_CharCollectionMap_Part1_test.go`

### Segment 15: CharCollectionMap.go — Part 2 (~47 stmts) + CharHashsetMap.go — Part 1 (~153 stmts)
- **Files**: `CharCollectionMap.go` remaining + `CharHashsetMap.go` first portion
- **Scope**: ~200 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_CharMaps_Part1_test.go`

### Segment 16: CharHashsetMap.go — Part 2 (~90 stmts) + SimpleStringOnce.go (~110 stmts)
- **Files**: `CharHashsetMap.go` remaining (90) + `SimpleStringOnce.go` first portion (110 of 174)
- **Scope**: ~200 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_CharHash_SimpleOnce_test.go`

### Segment 17: SimpleStringOnce.go — Part 2 (~64 stmts) + KeyValueCollection.go (~136 stmts)
- **Files**: `SimpleStringOnce.go` remaining (64) + `KeyValueCollection.go` first portion (136 of 144)
- **Scope**: ~200 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Once_KeyValue_test.go`

### Segment 18: KeyValueCollection.go remaining (8) + HashsetsCollection.go (102) + ValidValue.go (100)
- **Files**: `KeyValueCollection.go` remaining + `HashsetsCollection.go` + `ValidValue.go`
- **Scope**: ~210 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_HashsetsColl_Valid_test.go`

### Segment 19: ValidValues.go (90) + CollectionsOfCollection.go (70) + LinkedCollectionNode.go (66)
- **Files**: Three mid-size files
- **Scope**: ~226 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_ValidVals_CollOfColl_test.go`

### Segment 20: Remaining small files (~213 stmts total)
- **Files**: `KeyValuePair.go` (52), `LeftRight.go` (50), `NonChainedLinkedListNodes.go` (37), `NonChainedLinkedCollectionNodes.go` (37), `LeftMiddleRight.go` (32), `KeyAnyValuePair.go` (36)
- **Scope**: ~244 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_SmallTypes_Part1_test.go`

### Segment 21: Remaining creator + utility files (~193 stmts total)
- **Files**: `newSimpleSliceCreator.go` (36), `newHashsetCreator.go` (26), `newHashmapCreator.go` (24), `newCollectionCreator.go` (18), `emptyCreator.go` (16), `newLinkedListCreator.go` (14), `newCharCollectionMapCreator.go` (14), `newKeyValuesCreator.go` (13), `utils.go` (13), `newLinkedListCollectionsCreator.go` (11), `newHashsetsCollectionCreator.go` (11), `newCollectionsOfCollectionCreator.go` (9), `newCharHashsetMapCreator.go` (9), `newSimpleStringOnceCreator.go` (7), `TextWithLineNumber.go` (12), `reflectInterfaceVal.go` (7), `AnyToString.go` (5), `AllIndividualsLengthOfSimpleSlices.go` (5), `AllIndividualStringsOfStringsLength.go` (5), `LeftRightFromSplit.go` (4), `LeftMiddleRightFromSplit.go` (4), `ValueStatus.go` (3), `CloneSlice.go` (3), `CloneSliceIf.go` (2), + 6 DataModel files (2 each = 12)
- **Scope**: ~193 statements combined
- **Test file**: `tests/integratedtests/corestrtests/Coverage_Creators_Utils_test.go`

---

## Execution Plan

| Iteration | Segments | Est. Stmts | Files |
|-----------|----------|------------|-------|
| 1 | Seg 1 + Seg 2 | ~400 | Collection.go (Part 1+2) |
| 2 | Seg 3 + Seg 4 | ~327 | Collection.go (Part 3) + LinkedCollections.go (Part 1) |
| 3 | Seg 5 + Seg 6 | ~359 | LinkedCollections.go (Part 2) + Hashset.go (Part 1) |
| 4 | Seg 7 + Seg 8 | ~324 | Hashset.go (Part 2) + SimpleSlice.go (Part 1) |
| 5 | Seg 9 + Seg 10 | ~308 | SimpleSlice.go (Part 2) + Hashmap.go (Part 1) |
| 6 | Seg 11 + Seg 12 | ~316 | Hashmap.go (Part 2) + LinkedList.go (Part 1) |
| 7 | Seg 13 + Seg 14 | ~322 | LinkedList.go (Part 2) + CharCollectionMap.go (Part 1) |
| 8 | Seg 15 + Seg 16 | ~400 | CharMaps + CharHashset + SimpleStringOnce |
| 9 | Seg 17 + Seg 18 | ~410 | SimpleStringOnce + KeyValue + HashsetsColl + ValidValue |
| 10 | Seg 19 + Seg 20 | ~470 | ValidValues + CollOfColl + SmallTypes |
| 11 | Seg 21 | ~193 | Creators + Utils + DataModels |

**Total: 11 iterations × 2 segments each (except last) = ~3,917 statements**

---

## Rules

- Follow AAA pattern and naming conventions from spec/01-app/27-unit-coverage-fix.md
- Read source files before writing tests — never infer APIs
- Reference memory/tech/api-conventions-corestr for correct method signatures
- Each test file is self-contained and independently compilable
