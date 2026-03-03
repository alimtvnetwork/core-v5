# Test Coverage Audit: corejson, corepayload, coreinstruction

## Date: 2026-03-03 (updated)

## Summary

| Package | Test Files | Test Cases | Source Files | Coverage Rating |
|---------|-----------|------------|-------------|----------------|
| `corejson` | 3 | ~18 | 41 | 🟡 PARTIAL — core constructors and Result methods covered |
| `corepayload` | 7 | ~85 | ~20 | 🟢 GOOD — core types, paging, attributes, PagingInfo fully tested |
| `coreinstruction` | 3 | ~55 | 45 | 🟡 PARTIAL — StringCompare, StringSearch, Identifier, Specification covered |
| `coreapi` | 1 | 15 | ~5 | 🟢 GOOD — PageRequest fully tested |

---

## corejson — 🟡 PARTIAL (previously 🔴 CRITICAL)

### What IS tested (~18 cases across 3 files)
- ✅ `New` — valid struct, nil input, unmarshalable type (channel)
- ✅ `NewPtr` — valid struct, nil input, unmarshalable type (channel)
- ✅ `Result.Unmarshal` — valid, nil receiver, invalid bytes, error result
- ✅ `Result.IsEmpty` — empty bytes, nil receiver, valid bytes
- ✅ `Result.IsEqual` — equal, different content, both nil, one nil, same pointer

### What is NOT tested (remaining gaps)

| Function/Area | Risk | Notes |
|---|---|---|
| `Result.PrettyJsonStringOrErrString` | MEDIUM | Nil receiver path exists in cmd/main smoke tests |
| `Serialize.Apply` | MEDIUM | Used in PayloadWrapper tests but never directly tested |
| `BytesDeepClone` / `BytesCloneIf` | MEDIUM | Used in Attributes deep clone |
| `MapResults` / `ResultCollection` / `ResultsPtrCollection` | MEDIUM | Collection types |
| `anyTo` / `castingAny` | LOW | Internal helpers |
| `deserializeFromBytesTo` / `deserializeFromResultTo` | LOW | Core logic covered indirectly via Unmarshal tests |

### Recommended next (8 new test cases)
1. `Serialize.Apply` — valid struct, nil, unmarshalable type
2. `BytesDeepClone` — valid bytes, empty, nil
3. `Result.PrettyJsonStringOrErrString` — valid, error result

---

## corepayload — 🟢 GOOD (previously 🟡 PARTIAL)

### What IS tested (~85 cases across 7 files)
- ✅ `PayloadWrapper` — Create, DeserializeRoundtrip, Clone, DeserializeToMany (4 tests, 4 cases)
- ✅ `TypedPayloadCollection` — Creation, Add, Filter, Map, Reduce, Group, Partition, AllData, ElementAccess, Any/All (~15 cases)
- ✅ `TypedPayloadCollection` paging — GetPagesSize, GetSinglePageCollection, GetPagedCollection, GetPagedCollectionWithInfo + edge cases (~15 cases)
- ✅ `TypedPayloadCollection` FlatMap — wrapper-level, data-level, empty, nil output, nil wrapper, deserialization failure, nil receiver (~7 cases)
- ✅ `TypedPayloadWrapper` — Deserialization, RoundTrip, Clone, SetData, Nil/Invalid, DeserializeToMany (~10 cases)
- ✅ `Attributes.IsEqual` — 6 cases (pointer identity, content equality, nil handling)
- ✅ `Attributes.Clone` — 3 cases (deep clone independence, nil safety)
- ✅ `Attributes.IsSafeValid` — 3 cases (regression for negation logic bug)
- ✅ `AuthInfo.Clone` — 4 cases (deep copy of nested fields including Identifier)
- ✅ `PagingInfo.IsEqual` — 8 cases (both nil, left nil, right nil, equal, each field differing)
- ✅ `PagingInfo.ClonePtr` — 3 cases (nil, field copy, independence)
- ✅ `PagingInfo.IsEmpty` — 3 cases (nil, zero, non-zero)
- ✅ `PagingInfo.HasTotalPages` / `HasCurrentPageIndex` / `HasPerPageItems` / `HasTotalItems` — 8 cases (nil + positive)
- ✅ `PagingInfo.IsInvalidTotalPages` / `IsInvalidCurrentPageIndex` / `IsInvalidPerPageItems` / `IsInvalidTotalItems` — 10 cases (nil, zero/negative, positive)
- ✅ `PagingInfo.Clone` (value) — 2 cases (field copy, independence)

### What is NOT tested (low priority)

| Type/Area | Risk | Notes |
|---|---|---|
| `Attributes` getters/setters | LOW | ~20 accessor methods, straightforward |
| `User` / `SessionInfo` | LOW | Simple structs |
| `PayloadWrapper.IsEmpty` / `IsEmptyPayloads` | LOW | Guard clause methods |

---

## coreinstruction — 🟡 PARTIAL (improved)

### What IS tested (~55 cases across 3 files)
- ✅ `BaseIdentifier` — 4 cases (positive, special chars, empty, whitespace)
- ✅ `Identifiers` — Length (3), GetById (6), IndexOf (5), Clone (2), Add (2) = 18 cases
- ✅ `Specification.Clone` — 2 table cases + nil safety + deep copy verification = 4 tests
- ✅ `BaseTags` — 4 cases (all match, partial, empty-empty, empty-nonempty)
- ✅ `StringCompare.IsMatch` — Equal (match, no-match, case-sensitive), Contains, StartsWith (match, no-match, ignore-case), EndsWith (match, no-match, ignore-case), Regex (match, no-match) = 12 cases
- ✅ `StringCompare.IsMatchFailed` — Equal (match, mismatch), Contains (match, mismatch), Regex (match, no-match), nil receiver = 7 cases
- ✅ `StringCompare.VerifyError` — Equal match/mismatch, Contains mismatch, Regex match/no-match, invalid regex, nil = 7 cases
- ✅ `StringCompare` nil receiver — IsMatch vacuous truth = 1 case
- ✅ `StringSearch.IsMatch` — Equal (match, no-match), Contains (match, no-match), nil receiver = 5 cases
- ✅ `StringSearch.IsMatchFailed` — match, no-match, nil = 3 cases
- ✅ `StringSearch.IsAllMatch` — all pass, one fails, empty contents = 3 cases
- ✅ `StringSearch.IsAnyMatchFailed` — all match, one fails = 2 cases
- ✅ `StringSearch.IsEmpty` / `IsExist` / `Has` — nil + non-nil = 2 cases
- ✅ `StringSearch.VerifyError` — match, no-match, nil = 3 cases

### What is NOT tested

| Type/Area | Risk | Notes |
|---|---|---|
| `IdentifiersWithGlobals` | MEDIUM | Listed in testing roadmap Phase 2; GetById, Length, Clone, Add |
| `FromTo` / `BaseFromTo` | MEDIUM | Used in Attributes; has ClonePtr |
| `StringCompare` / `StringSearch` constructors | LOW | Covered indirectly |
| `NameList` / `NameListCollection` | LOW | Collection with potential edge cases |
| `FlatSpecification` | LOW | Flattening logic |

### Recommended next (6 new test cases)
1. `IdentifiersWithGlobals` — GetById found/not-found, Length, Clone, Add (per roadmap Phase 2)
2. `FromTo.ClonePtr` — positive, nil receiver

---

## coreapi — 🟢 GOOD (new)

### What IS tested (15 cases in 1 file)
- ✅ `PageRequest.IsPageSizeEmpty` — nil, zero, negative, positive = 4 cases
- ✅ `PageRequest.IsPageIndexEmpty` — nil, zero, positive = 3 cases
- ✅ `PageRequest.HasPageSize` — nil, positive = 2 cases
- ✅ `PageRequest.HasPageIndex` — nil, positive = 2 cases
- ✅ `PageRequest.Clone` — nil, field copy, independence = 3 cases

---

## Priority Order for Remaining Implementation

1. **coreinstruction** `IdentifiersWithGlobals` — per testing roadmap Phase 2
2. **coreinstruction** `FromTo.ClonePtr` — regression prevention
3. **corejson** `Serialize.Apply` / `BytesDeepClone` — direct coverage for shared utilities
4. **corejson** `Result.PrettyJsonStringOrErrString` — error path coverage
