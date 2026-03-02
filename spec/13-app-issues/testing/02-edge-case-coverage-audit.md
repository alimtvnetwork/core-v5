# Test Coverage Audit: corejson, corepayload, coreinstruction

## Date: 2026-03-02

## Summary

| Package | Test Files | Test Cases | Source Files | Coverage Rating |
|---------|-----------|------------|-------------|----------------|
| `corejson` | 1 | 1 | 41 | 🔴 CRITICAL — nearly untested |
| `corepayload` | 6 | ~55 | ~20 | 🟡 PARTIAL — core types tested, supporting types not |
| `coreinstruction` | 1 | ~30 | 45 | 🟡 PARTIAL — Identifier/Specification covered, many types not |

---

## corejson — 🔴 CRITICAL GAPS

### What IS tested (1 test)
- `Deserialize.FromTo` — single positive case

### What is NOT tested (priority order)

| Function/Area | Risk | Notes |
|---|---|---|
| `Result.Unmarshal` | HIGH | Used everywhere; nil receiver, nil target, invalid bytes untested |
| `Result.IsEmpty` / `Result.HasError` | HIGH | Guard clause logic |
| `Result.IsEqual` | MEDIUM | Equality comparison |
| `Result.PrettyJsonStringOrErrString` | MEDIUM | Nil receiver path exists in cmd/main smoke tests |
| `New` / `NewPtr` | HIGH | Core constructors; marshal failure path untested |
| `Serialize.Apply` | HIGH | Used in PayloadWrapper tests but never directly tested |
| `BytesDeepClone` / `BytesCloneIf` | MEDIUM | Used in Attributes deep clone |
| `MapResults` / `ResultCollection` / `ResultsPtrCollection` | MEDIUM | Collection types |
| `anyTo` / `castingAny` | LOW | Internal helpers |
| `deserializeFromBytesTo` / `deserializeFromResultTo` | HIGH | Core deserialization logic |

### Recommended minimum (15 new test cases)
1. `New` — valid struct, nil input, unmarshalable type (channel)
2. `NewPtr` — same 3 cases
3. `Result.Unmarshal` — valid, nil receiver, nil target, invalid bytes
4. `Result.IsEmpty` — empty bytes, nil, valid
5. `Result.IsEqual` — equal, different bytes, nil vs non-nil

---

## corepayload — 🟡 PARTIAL GAPS

### What IS tested (~55 cases across 6 files)
- `PayloadWrapper` — Create, DeserializeRoundtrip, Clone, DeserializeToMany (4 tests, 4 cases)
- `TypedPayloadCollection` — Creation, Add, Filter, Map, Reduce, Group, Partition, AllData, ElementAccess, Any/All (10 tests, ~15 cases)
- `TypedPayloadCollection` paging — GetPagesSize, GetSinglePageCollection, GetPagedCollection, GetPagedCollectionWithInfo + edge cases (8 tests, ~15 cases)
- `TypedPayloadCollection` FlatMap — wrapper-level, data-level, empty, nil output, nil wrapper, deserialization failure, nil receiver (7 tests, ~7 cases)
- `TypedPayloadWrapper` — Deserialization, RoundTrip, Clone, SetData, Nil/Invalid, DeserializeToMany (6 tests, ~10 cases)

### What is NOT tested

| Type/Area | Risk | Notes |
|---|---|---|
| `Attributes.IsEqual` | HIGH | Complex equality with 6 sub-comparisons; just fixed `IsSafeValid` bug here |
| `Attributes.Clone` / `ClonePtr` / `deepClonePtr` | HIGH | Deep clone with error paths |
| `Attributes` getters/setters | MEDIUM | ~20 accessor methods |
| `AuthInfo` | MEDIUM | Just fixed missing `Identifier` in Clone |
| `PagingInfo.IsEqual` / `ClonePtr` | MEDIUM | Used in Attributes.IsEqual |
| `User` / `SessionInfo` | LOW | Simple structs |
| `PayloadWrapper.IsEmpty` / `IsEmptyPayloads` | MEDIUM | Guard clause methods |

### Recommended minimum (10 new test cases)
1. `Attributes.IsEqual` — both nil, one nil, equal, different error, different paging, different KV
2. `Attributes.Clone(true)` — deep clone independence verification
3. `Attributes.IsSafeValid` — valid, invalid, nil (regression for Bug #3)
4. `AuthInfo.Clone` — all fields including Identifier (regression for missing field bug)

---

## coreinstruction — 🟡 PARTIAL GAPS

### What IS tested (~30 cases)
- `BaseIdentifier` — 4 cases (positive, special chars, empty, whitespace)
- `Identifiers` — Length (3), GetById (6), IndexOf (5), Clone (2), Add (2) = 18 cases
- `Specification.Clone` — 2 table cases + nil safety + deep copy verification = 4 tests
- `BaseTags` — 4 cases (all match, partial, empty-empty, empty-nonempty)

### What is NOT tested

| Type/Area | Risk | Notes |
|---|---|---|
| `IdentifiersWithGlobals` | HIGH | Listed in testing roadmap Phase 2; GetById, Length, Clone, Add |
| `FromTo` / `BaseFromTo` | MEDIUM | Used in Attributes; has ClonePtr |
| `SourceDestination` | LOW | Simple struct |
| `StringCompare` / `StringSearch` | MEDIUM | Comparison logic |
| `NameList` / `NameListCollection` | MEDIUM | Collection with potential edge cases |
| `NameRequests` / `NameRequestsCollection` | LOW | Request types |
| `Rename` | LOW | Simple struct |
| `FlatSpecification` | MEDIUM | Flattening logic |
| `RequestSpecification` | LOW | Composition type |

### Recommended minimum (8 new test cases)
1. `IdentifiersWithGlobals` — GetById found/not-found, Length, Clone, Add (per roadmap Phase 2)
2. `FromTo.ClonePtr` — positive, nil receiver
3. `StringCompare` — equal, not equal, case sensitivity

---

## Priority Order for Implementation

1. **corejson** `New`/`NewPtr`/`Result.Unmarshal` — highest risk, nearly zero coverage
2. **corepayload** `Attributes.IsEqual`/`Clone`/`IsSafeValid` — regression prevention for 3 recent bugs
3. **coreinstruction** `IdentifiersWithGlobals` — per testing roadmap Phase 2
4. **corejson** remaining Result methods
5. **corepayload** AuthInfo.Clone regression test
