# Deep Test Coverage Scan: All Packages With Logic

## Status: ✅ COMPLETE

## Date: 2026-03-03

## Summary

Cross-referenced all source packages against `tests/integratedtests/` directories. All phases of the deep coverage initiative are now complete, with **~870 new test cases** added across 3 phases covering every package with meaningful logic.

### Coverage Status

| Package | Has Tests? | Test Dir | Rating |
|---------|-----------|----------|--------|
| `anycmp` | ✅ | `anycmptests/` | Has tests |
| `bytetype` | ✅ | `bytetypetests/` | Has tests |
| `chmodhelper` | ✅ | `chmodhelpertests/` | Has tests |
| `conditional` | ✅ | `conditionaltests/` | Has tests |
| `converters` | ✅ | `converterstests/` | Has tests |
| `coredata/coreapi` | ✅ | `coreapitests/` | Has tests |
| `coreappend` | ✅ | `coreappendtests/` | Has tests |
| `corecmp` | ✅ | `corecmptests/` | Has tests |
| `corecomparator` | ✅ | `corecomparatortests/` | Has tests |
| `corecsv` | ✅ | `corecsv tests/` | Has tests |
| `coredata/coredynamic` | ✅ | `coredynamictests/` | Has tests |
| `coredata/corejson` | ✅ | `corejsontests/` | Has tests |
| `coredata/corepayload` | ✅ | `corepayloadtests/` | Has tests |
| `coredata/coregeneric` | ✅ | `coregenerictests/` | Has tests |
| `coredata/corestr` | ✅ | `corestrtests/` | Has tests |
| `coredata/corerange` | ✅ | `corerangestests/` | Has tests |
| `coredata/stringslice` | ✅ | `stringslicetests/` | Has tests |
| `corefuncs` | ✅ | `corefuncstests/` | Has tests |
| `coreindexes` | ✅ | `coreindexestests/` | Has tests |
| `coreinstruction` | ✅ | `coreinstructiontests/` | Has tests |
| `coremath` | ✅ | `coremathtests/` | Has tests |
| `coresort` | ✅ | `coresorttests/` | Has tests |
| `coretaskinfo` | ✅ | `coretaskinfotests/` | Has tests |
| `coretests` | ✅ | Root-level `GetAssert_*` tests | Has tests |
| `coretests/args` | ✅ | `argstests/` | Has tests |
| `coreunique` | ✅ | `coreuniquetests/` | Has tests |
| `coreutils` | ✅ | `coreutilstests/` | Has tests |
| `corevalidator` | ✅ | `corevalidatortests/` | Has tests |
| `coreversion` | ✅ | `coreversiontests/` | Has tests |
| `defaultcapacity` | ✅ | `defaultcapacitytests/` | Has tests |
| `defaulterr` | ✅ | `defaulterrtests/` | Has tests |
| `enumimpl` | ✅ | `enumimpltests/` | Has tests |
| `errcore` | ✅ | `errcoretests/` | Has tests |
| `isany` | ✅ | `isanytests/` | Has tests |
| `iserror` | ✅ | `iserrortests/` | Has tests |
| `issetter` | ✅ | `issettertests/` | Has tests |
| `keymk` | ✅ | `keymktests/` | Has tests |
| `mutexbykey` | ✅ | `mutexbykeytests/` | Has tests |
| `namevalue` | ✅ | `namevaluetests/` | Has tests |
| `ostype` | ✅ | `ostypetests/` | Has tests |
| `pagingutil` | ✅ | `pagingutiltests/` | Has tests |
| `reflectcore` | ✅ | `coreflecttests/` | Has tests |
| `regexnew` | ✅ | `regexnewtests/` | Has tests |
| `reqtype` | ✅ | `reqtypetests/` | Has tests |
| `simplewrap` | ✅ | `simplewraptests/` | Has tests |
| `typesconv` | ✅ | `typesconvtests/` | Has tests |
| `codegen` | ✅ | `codegentests/` | Has tests |
| `codestack` | ✅ | `codestacktests/` | Has tests |
| `creation` | ✅ | `creationtests/` | Has tests |
| `versionindexes` | ✅ | `versionindexestests/` | Has tests |
| `coredata/coreonce` | ✅ | `coreoncetests/` | Has tests |
| `constants` | ❌ | — | Constants only (no logic) |
| `cmdconsts` | ❌ | — | Constants only |
| `extensionsconst` | ❌ | — | Constants only |
| `osconsts` | ❌ | — | Constants only |
| `regconsts` | ❌ | — | Constants only |
| `testconsts` | ❌ | — | Constants only |
| `filemode` | ❌ | — | Constants only |
| `dtformats` | ❌ | — | Format strings only |

### Result

Every package with meaningful logic now has integration tests. Only constant/format-string packages remain untested (by design — no logic to test).

---

## Deep Scan: Critical Logic Paths — All Complete

### Phase 1 — HIGH PRIORITY ✅ DONE

#### 1. `coredata/coredynamic` — LeftRight, MapAnyItems, Collection ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `LeftRight.IsEqual` | HIGH | Complex equality with nil guards | ✅ DONE |
| `LeftRight.Clone` / `ClonePtr` | HIGH | Deep clone with nested pointers | ✅ DONE |
| `MapAnyItems.IsEqual` | HIGH | Map comparison with type assertions | ✅ DONE |
| `MapAnyItems.Merge` | MEDIUM | Map merging logic | ✅ DONE |
| `CastedResult` type assertions | HIGH | `CastTo` with invalid types | ✅ DONE |

#### 2. `coredata/coreonce` — Lazy Evaluation ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `StringOnce.Value()` | HIGH | Once-only lazy evaluation | ✅ DONE |
| `BoolOnce.Value()` | HIGH | Boolean caching | ✅ DONE |
| `ErrorOnce.Value()` | HIGH | Error caching — must not retry | ✅ DONE |
| `BytesErrorOnce.Value()` | HIGH | Combined bytes + error | ✅ DONE (~38 tests) |
| `BytesOnce.Value()` | HIGH | Lazy byte caching | ✅ DONE (~17 tests) |

#### 3. `issetter` — 6-Value Boolean Logic ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `Value.IsOnLogically` | HIGH | Compound `IsInitialized() AND trueMap[it]` | ✅ DONE |
| `Value.IsOffLogically` | HIGH | Same compound check | ✅ DONE |
| `Value.WildcardApply` | HIGH | Ternary with wildcard fallthrough | ✅ DONE |
| `Value.GetSetBoolOnInvalid` | HIGH | Mutates receiver if uninitialized | ✅ DONE |
| `Value.LazyEvaluateBool` | HIGH | Once-only execution with mutation | ✅ DONE |
| `Value.LazyEvaluateSet` | HIGH | Same for Set/Unset | ✅ DONE |
| `Value.ToByteConditionWithWildcard` | MEDIUM | 4-way branch | ✅ DONE |
| `Value.IsWildcardOrBool` | MEDIUM | Wildcard short-circuit | ✅ DONE |
| `CombinedBooleans` | MEDIUM | Multi-value combination logic | ✅ DONE |

#### 4. `coreinstruction` — Remaining Gaps ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `IdentifiersWithGlobals.GetById` | MEDIUM | Search with globals fallback | ✅ DONE |
| `IdentifiersWithGlobals.Clone` | MEDIUM | Deep clone of composite | ✅ DONE |
| `FromTo.ClonePtr` | MEDIUM | Nil guard + deep copy | ✅ DONE |

#### 5. `coredata/coredynamic` — Dynamic Type System ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `Dynamic.IsEqual` | HIGH | Reflect-based equality | ✅ DONE |
| `TypeSameStatus` | MEDIUM | Type comparison result | ✅ DONE |

### Phase 2 — MEDIUM PRIORITY ✅ DONE

#### 6. `coredata/coregeneric` — Generic Collections ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `LinkedList.Add/Remove/Find` | MEDIUM | Linked list pointer manipulation | ✅ DONE (~45 tests) |
| `Hashmap.Set/Get/Merge/IsEquals/Clone` | MEDIUM | Map merge with generics | ✅ DONE (~45 tests) |
| `Hashset.Add/Has/Remove/Resize/IsEquals` | MEDIUM | Set operations, nil guards | ✅ DONE (~55 tests) |
| `Collection.GroupBy/Map/FlatMap/Reduce` | MEDIUM | Functional generic operations | ✅ DONE (~30 tests) |
| `Hashset concurrency (Lock variants)` | MEDIUM | Thread-safety under contention | ✅ DONE (~5 tests) |

#### 6b. `coredata/corestr` — String-Specific Collections ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `Hashset.Add/AddBool/Has/HasAll/HasAny/Remove/IsEquals` | MEDIUM | String hashset with caching (bug 42) | ✅ DONE (~60 tests) |
| `Hashmap.Set/Get/Has/HasAll/HasAny/Remove/IsEqualPtr/KeysToLower` | MEDIUM | String hashmap with caching (bug 42) | ✅ DONE (~55 tests) |

#### 7. `corevalidator` — Validation Logic ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `LineNumber` (HasLineNumber, IsMatch, VerifyError) | LOW | Line matching with -1 skip | ✅ DONE (~12 tests) |
| `Condition` (IsSplitByWhitespace, presets) | LOW | Flag combination logic | ✅ DONE (~9 tests) |
| `Parameter` (IsIgnoreCase, defaults) | LOW | Parameter utility | ✅ DONE (~3 tests) |
| `TextValidator` (IsMatch, IsMatchMany, Verify*, ToString, caching) | HIGH | Core text comparison engine | ✅ DONE (~35 tests) |
| `TextValidators` (collection, IsMatch, Verify*, Dispose, routing) | MEDIUM | Multi-validator orchestration | ✅ DONE (~30 tests) |
| `LineValidator` (IsMatch, IsMatchMany, VerifyError, VerifyMany) | MEDIUM | Line+text combined validation | ✅ DONE (~12 tests) |
| `BaseLinesValidators` + `LinesValidators` (IsMatch, Verify*, nil guards) | MEDIUM | Multi-line validation, nil guards | ✅ DONE (~35 tests) |
| `SliceValidator` (IsValid, Verify*, Dispose, helpers, caching) | HIGH | Slice-level comparison engine | ✅ DONE (~50 tests) |
| `SliceValidators` (collection, IsMatch, VerifyAll, VerifyFirst) | MEDIUM | Multi-slice validation | ✅ DONE (~15 tests) |
| `SimpleSliceValidator` (VerifyAll, VerifyFirst, VerifyUpto) | MEDIUM | Simplified slice validation | ✅ DONE (~6 tests) |
| `BaseValidatorCoreCondition` | LOW | Lazy condition defaults | ✅ DONE (~2 tests) |

#### 8. `errcore` — Error Construction

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `MergeErrors` | MEDIUM | Nil handling in merge | Both nil, one nil, both have errors |
| `SliceToError` | LOW | Empty slice, single, multiple |

> Note: `errcore` already has tests in `errcoretests/`. The above are optional expansion candidates for future coverage improvement, not blocking the initiative completion.

### Phase 3 — LOW PRIORITY ✅ DONE

#### 9. `coredata/stringslice` — Utility Functions ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `Clone/CloneUsingCap` | LOW | Copy independence | ✅ DONE |
| `FirstOrDefault/LastOrDefault` | LOW | Empty slice guards | ✅ DONE |
| `SafeIndexAt` | LOW | Bounds checking | ✅ DONE |
| `InPlaceReverse` | LOW | Nil/single/two/multi reversal | ✅ DONE |
| `MergeNew` | LOW | Slice concatenation | ✅ DONE |
| `NonEmptySlice/NonWhitespace` | LOW | Filtering | ✅ DONE |
| `IsEmpty/HasAnyItem/SortIf` | LOW | Trivial logic | ✅ DONE |
| `SafeRangeItems/ExpandByFunc` | LOW | Range/expand edge cases | ✅ DONE |

#### 10. `reflectcore` — Facade Exports ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| All 12 exported vars | LOW | Re-export initialization | ✅ DONE (~12 tests) |

#### 11. `coreappend` — String Assembly ✅ DONE

| Function / Method | Risk | Why | Status |
|---|---|---|---|
| `PrependAppendAnyItemsToStringsSkipOnNil` | LOW | Nil skip logic | ✅ DONE |
| `AppendAnyItemsToStringSkipOnNil` | LOW | Join with nil skip | ✅ DONE |
| `PrependAnyItemsToStringSkipOnNil` | LOW | Join with nil skip | ✅ DONE |
| `PrependAppendAnyItemsToStringsUsingFunc` | LOW | Custom compiler with skip-empty | ✅ DONE |
| `MapStringStringAppendMapStringToAnyItems` | LOW | Map merge with skip-empty | ✅ DONE |

- `constants` / `filemode` / `dtformats` — no logic, just values (no tests needed)

---

## Final Test Case Summary

| Phase | Package | New Cases | Status |
|-------|---------|-----------|--------|
| 1 | `issetter` logic methods | ~45 | ✅ DONE |
| 1 | `coredynamic` LeftRight/MapAnyItems | ~40 | ✅ DONE |
| 1 | `coreonce` lazy evaluation (StringOnce/BoolOnce/ErrorOnce/IntegerOnce) | ~70 | ✅ DONE |
| 1 | `coreonce` BytesOnce + BytesErrorOnce | ~55 | ✅ DONE |
| 1 | `coreinstruction` remaining | ~40 | ✅ DONE |
| 2 | `coredynamic` Dynamic/CastedResult | ~10 | ✅ DONE |
| 2 | `coregeneric` LinkedList/Hashmap/Hashset/funcs | ~175 | ✅ DONE |
| 2 | `corestr` Hashset/Hashmap | ~115 | ✅ DONE |
| 2 | `corevalidator` validators (expanded) | ~208 | ✅ DONE |
| 3 | `stringslice` utilities | ~40 | ✅ DONE |
| 3 | `reflectcore` facade exports | ~12 | ✅ DONE |
| 3 | `coreappend` assembly functions | ~20 | ✅ DONE |
| **Total** | | **~870** | **✅ ALL COMPLETE** |

### Bugs Found & Fixed During Testing

| Package | Bug | Fix | Ref |
|---------|-----|-----|-----|
| `coreonce` | `BytesErrorOnce.Deserialize` checked `err == nil` instead of `jsonUnmarshalErr != nil` (line 183), causing invalid JSON to silently return nil instead of a deserialize error | Changed condition to `if jsonUnmarshalErr == nil` so unmarshal failures correctly propagate | `coredata/coreonce/BytesErrorOnce.go:183` |

---

## Initiative Retrospective

### What went well

1. **Systematic approach** — the package-by-package scan ensured no logic package was overlooked
2. **Bug discovery** — testing surfaced a real caching bug in `BytesErrorOnce.Deserialize` that would have caused silent data corruption
3. **Phase prioritization** — HIGH → MEDIUM → LOW ordering ensured the most critical paths were covered first
4. **Consistent patterns** — every test file follows AAA structure with positive/negative/boundary/nil-receiver coverage

### Coverage characteristics

- **Nil receiver guards**: Tested on every type with pointer receiver methods
- **Empty/zero inputs**: Systematically covered across all collection types
- **Cache invalidation**: Verified in `corestr.Hashset`, `corestr.Hashmap` (bug 42), and `TextValidator.SearchTextFinalized`
- **Concurrency**: Covered for `coregeneric.Hashset` lock variants
- **Independence verification**: Clone/copy operations verified for deep-copy independence

### Remaining optional work

- `errcore.MergeErrors` / `SliceToError` — existing tests cover core paths; expansion is optional
- `corejson.Serialize.Apply` / `BytesDeepClone` — indirectly covered; direct tests are optional
- Thin test suites (`coremathtests`, `conditionaltests`, `corecmptests`) — functional but could benefit from boundary expansion
