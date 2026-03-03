# Deep Test Coverage Scan: All Packages With Logic

## Date: 2026-03-03

## Summary

Cross-referenced all source packages against `tests/integratedtests/` directories.

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
| `regexnew` | ✅ | `regexnewtests/` | Has tests |
| `reqtype` | ✅ | `reqtypetests/` | Has tests |
| `simplewrap` | ✅ | `simplewraptests/` | Has tests |
| `typesconv` | ✅ | `typesconvtests/` | Has tests |
| `codegen` | ✅ | `codegentests/` | Has tests |
| `codestack` | ✅ | `codestacktests/` | Has tests |
| `creation` | ✅ | `creationtests/` | Has tests |
| `versionindexes` | ✅ | `versionindexestests/` | Has tests |
| `coredata/coreonce` | ✅ | `coreoncetests/` | Has tests (Phase 1 + BytesOnce/BytesErrorOnce) |
| `reflectcore` | ❌ | — | **NO TESTS** |
| `constants` | ❌ | — | Constants only (no logic) |
| `cmdconsts` | ❌ | — | Constants only |
| `extensionsconst` | ❌ | — | Constants only |
| `osconsts` | ❌ | — | Constants only |
| `regconsts` | ❌ | — | Constants only |
| `testconsts` | ❌ | — | Constants only |
| `filemode` | ❌ | — | Constants only |
| `dtformats` | ❌ | — | Format strings only |

### Packages with NO tests that HAVE logic

| Package | Risk | Files with Logic | Priority |
|---------|------|-----------------|----------|
| `coredata/coreonce` | MEDIUM | 12 `*Once.go` files — lazy evaluation with sync patterns | Phase 1 |
| `reflectcore` | LOW | `vars.go` only — thin wrapper | Phase 2 |

---

## Deep Scan: Critical Logic Paths Needing Edge Case Tests

### Phase 1 — HIGH PRIORITY (branching logic, nil guards, edge cases)

#### 1. `coredata/coredynamic` — LeftRight, MapAnyItems, Collection

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `LeftRight.IsEqual` | HIGH | Complex equality with nil guards on both sides | Both nil, left nil, right nil, equal, different left, different right |
| `LeftRight.Clone` / `ClonePtr` | HIGH | Deep clone with nested pointers | Nil receiver, field copy, independence |
| `MapAnyItems.IsEqual` | HIGH | Map comparison with type assertions | Both nil, different lengths, same keys diff values, nested maps |
| `MapAnyItems.Merge` | MEDIUM | Map merging logic with overwrite behavior | Empty maps, overlapping keys, nil receiver |
| `Collection.Filter` / `Map` | MEDIUM | Generic collection operations | Empty collection, all match, none match, nil predicate |
| `CastedResult` type assertions | HIGH | `CastTo` with invalid types | Valid cast, wrong type, nil input |

#### 2. `coredata/coreonce` — Lazy Evaluation

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `StringOnce.Value()` | HIGH | Once-only lazy evaluation; concurrent safety | First call, second call returns cached, concurrent calls |
| `BoolOnce.Value()` | HIGH | Same pattern | True result, false result, nil func |
| `ErrorOnce.Value()` | HIGH | Error caching — must not retry on error | Error result cached, nil error cached |
| `BytesErrorOnce.Value()` | ✅ DONE | Combined bytes + error | Caching, Deserialize edge cases, HasIssuesOrEmpty, nil guards, state queries (~38 tests) |
| `BytesOnce.Value()` | ✅ DONE | Lazy byte caching | Caching, nil initializer, Length, IsEmpty, JSON, String (~17 tests) |

#### 3. `issetter` — 6-Value Boolean Logic

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `Value.IsOnLogically` | HIGH | Combines `IsInitialized()` AND `trueMap[it]` | Each of 6 values: Uninitialized, True, False, Unset, Set, Wildcard |
| `Value.IsOffLogically` | HIGH | Same compound check | All 6 values |
| `Value.WildcardApply` | HIGH | Ternary with wildcard fallthrough | Wildcard+true, Wildcard+false, True+any, False+any, Uninitialized+any |
| `Value.GetSetBoolOnInvalid` | HIGH | Mutates receiver if uninitialized | Already set, uninitialized+true, uninitialized+false |
| `Value.LazyEvaluateBool` | HIGH | Once-only execution with mutation | Already defined, uninitialized triggers func |
| `Value.LazyEvaluateSet` | HIGH | Same for Set/Unset | Already set, uninitialized triggers func |
| `Value.ToByteConditionWithWildcard` | MEDIUM | 4-way branch | True, False, Wildcard, Uninitialized |
| `Value.IsWildcardOrBool` | MEDIUM | Wildcard short-circuit | Wildcard returns true, True+true, False+false |
| `CombinedBooleans` | MEDIUM | Multi-value combination logic | All combinations |

#### 4. `coreinstruction` — Remaining Gaps

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `IdentifiersWithGlobals.GetById` | MEDIUM | Search with globals fallback | Found in main, found in globals, not found |
| `IdentifiersWithGlobals.Clone` | MEDIUM | Deep clone of composite | Nil, populated, independence |
| `FromTo.ClonePtr` | MEDIUM | Nil guard + deep copy | Nil receiver, valid copy |

#### 5. `coredata/coredynamic` — Dynamic Type System

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `Dynamic.IsEqual` | HIGH | Reflect-based equality | Same type same value, same type diff value, different types, nil |
| `TypeSameStatus` | MEDIUM | Type comparison result | Same, different, nil inputs |
| `SafeZeroSet` / `ZeroSet` | HIGH | Reflect-based zero value assignment | Valid target, nil target, non-settable |

### Phase 2 — MEDIUM PRIORITY

#### 6. `coredata/coregeneric` — Generic Collections

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `LinkedList.Add/Remove/Find` | MEDIUM | Linked list pointer manipulation | Empty list ops, single element, head/tail removal |
| `Hashmap.Merge` | MEDIUM | Map merge with generics | Empty, overlapping keys |
| `Hashset.Intersect/Union/Diff` | MEDIUM | Set operations | Empty sets, disjoint, overlapping |
| `Collection.GroupBy` | MEDIUM | Grouping with key function | Empty, single group, multiple groups |

#### 7. `corevalidator` — Validation Logic

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `LinesValidators` | MEDIUM | Multi-line validation with error aggregation | All pass, one fails, empty lines |
| Range validators | MEDIUM | Boundary checks | At boundary, below, above |

#### 8. `errcore` — Error Construction

| Function / Method | Risk | Why | Test Cases Needed |
|---|---|---|---|
| `MergeErrors` | MEDIUM | Nil handling in merge | Both nil, one nil, both have errors |
| `SliceToError` | LOW | Empty slice, single, multiple |

### Phase 3 — LOW PRIORITY (simple logic, well-tested indirectly)

- `coredata/stringslice` — many utility functions, most are straightforward
- `reflectcore` — thin wrapper, low risk
- `coreappend` — simple append operations
- `constants` / `filemode` / `dtformats` — no logic, just values

---

## Implementation Order

1. ~~**`issetter` logic methods** — 6-value boolean has the most complex branching~~ ✅ DONE (45 tests)
2. ~~**`coredata/coredynamic` LeftRight + MapAnyItems** — equality and clone with nil guards~~ ✅ DONE (40 tests)
3. ~~**`coredata/coreonce`** — lazy evaluation correctness (was only package with zero tests)~~ ✅ DONE (70 + 55 = 125 tests)
4. ~~**`coreinstruction` IdentifiersWithGlobals + FromTo** — per existing roadmap~~ ✅ DONE (40 tests)
5. **`coredata/coredynamic` Dynamic type system** — reflect-based operations
6. **`coredata/coregeneric` LinkedList** — pointer manipulation edge cases

## Estimated Test Cases

| Phase | Package | New Cases | Status |
|-------|---------|-----------|--------|
| 1 | `issetter` logic methods | ~45 | ✅ DONE |
| 1 | `coredynamic` LeftRight/MapAnyItems | ~40 | ✅ DONE |
| 1 | `coreonce` lazy evaluation (StringOnce/BoolOnce/ErrorOnce/IntegerOnce) | ~70 | ✅ DONE |
| 1 | `coreonce` BytesOnce + BytesErrorOnce | ~55 | ✅ DONE |
| 1 | `coreinstruction` remaining | ~40 | ✅ DONE |
| 2 | `coredynamic` Dynamic/CastedResult | ~10 | TODO |
| 2 | `coregeneric` LinkedList/Hashmap | ~12 | TODO |
| 2 | `corevalidator` validators | ~8 | TODO |
| 3 | Remaining low-priority | ~10 | TODO |
| **Total** | | **~290** | **Phase 1 complete** |
