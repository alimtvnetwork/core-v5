# CaseNilSafe Migration Tracker

## Status Legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Migrated to CaseNilSafe with `results.ResultAny` |

---

## Architecture (Corrected)

All migrated test cases now use the **corrected architecture**:

| Aspect | Old (incorrect) | New (correct) |
|--------|----------------|---------------|
| **Expected type** | `args.Map` (untyped) | `results.ResultAny` (typed struct) |
| **Assertion owner** | Methods on `CaseNilSafe` | Methods on `results.Result[T]` (`ShouldMatchResult`) |
| **CaseNilSafe role** | Logic + data | Data-only + thin `ShouldBeSafe` convenience that delegates to `Result` |
| **Error sentinel** | `"hasError": true` in map | `results.ExpectAnyError` error value |
| **Field comparison** | Manual map key matching | Auto-derived from `Expected` fields, with optional `CompareFields` override |

See `spec/01-app/designs/CaseNilSafe-design.md` for full architecture.

---

## Migrated (✅)

### Priority A — Inline `t.Error` (Complete)

| # | Package | File | Style Before | Cases |
|---|---------|------|-------------|-------|
| 1 | `corestrtests` | `Hashset_NilReceiver_testcases.go` | Inline `t.Error` | 5 |
| 2 | `regexnewtests` | `LazyRegex_NilReceiver_testcases.go` | Inline `t.Error` | 10 |
| 3 | `reflectmodeltests` | `FieldProcessor_NilReceiver_testcases.go` | Inline `t.Error` | 2 |
| 4 | `reflectmodeltests` | `MethodProcessor_NilReceiver_testcases.go` | Inline `t.Error` | 10 |
| 5 | `reflectmodeltests` | `ReflectValueKind_NilReceiver_testcases.go` | Inline `t.Error` | 8 |
| 6 | `coredatatests` | `BytesError_NilReceiver_testcases.go` | Inline `t.Error` | 6 |
| 7 | `corevalidatortests` | `SliceValidator_NilReceiver_testcases.go` | Inline `t.Error` | 11 |
| 8 | `corevalidatortests` | `SliceValidators_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 9 | `corevalidatortests` | `TextValidator_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 10 | `corevalidatortests` | `TextValidators_NilReceiver_testcases.go` | Inline `t.Error` | 3 |
| 11 | `corevalidatortests` | `BaseLinesValidators_NilReceiver_testcases.go` | Inline `t.Error` | 5 |
| 12 | `corevalidatortests` | `LineValidator_NilReceiver_testcases.go` | Inline `t.Error` | 1 |
| 13 | `coregenerictests` | `Hashset_NilReceiver_testcases.go` | Inline `t.Error` | 3 |

### Priority A — CaseV1 / Custom Wrapper (Complete)

| # | Package | File | Style Before | Cases |
|---|---------|------|-------------|-------|
| 14 | `coreinstructiontests` | `StringCompare_NilReceiver_testcases.go` | CaseV1 string-dispatch | 5 |
| 15 | `coregenerictests` | `LinkedList_NilReceiver_testcases.go` | CaseV1 | 3 |
| 16 | `namevaluetests` | `Collection_NilReceiver_testcases.go` | CaseV1 | 1 |
| 17 | `coreoncetests` | `BytesErrorOnce_NilReceiver_testcases.go` | Custom `IsNilReceiver` wrapper | 4 |
| 18 | `corepayloadtests` | `TypedCollection_NilReceiver_testcases.go` | CaseV1 / GenericGherkins | 3 |
| 19 | `coreapitests` | `TypedConversions_NilReceiver_testcases.go` | CaseV1 string-dispatch | 4 |

### Self-Test

| # | Package | File | Cases |
|---|---------|------|-------|
| 20 | `casenilsafetests` | `CaseNilSafe_test.go` | 12 |

### Priority B — CaseV1 with Nil Receiver (Complete)

| # | Package | File | Style Before | Cases |
|---|---------|------|-------------|-------|
| 21 | `coredynamictests` | `Dynamic_NilReceiver_testcases.go` | CaseV1 standalone vars | 3 |
| 22 | `coredynamictests` | `LeftRight_NilReceiver_testcases.go` | CaseV1 mixed slices | 5 |
| 23 | `coredynamictests` | `CastedResult_NilReceiver_testcases.go` | CaseV1 mixed slices | 6 |
| 24 | `coredynamictests` | `MapAnyItemsEdge_NilReceiver_testcases.go` | CaseV1 standalone vars | 3 |
| 25 | `coregenerictests` | `Hashmap_NilReceiver_testcases.go` | CaseV1 standalone vars | 3 |
| 26 | `coregenerictests` | `PairTriple_NilReceiver_testcases.go` | Inline `t.Error` | 2 |
| 27 | `coreapitests` | `PageRequest_NilReceiver_testcases.go` | CaseV1 mixed slices | 5 |
| 28 | `trydotests` | `WrappedErr_NilReceiver_testcases.go` | CaseV1 mixed slices | 6 |
| 29 | `errcoretests` | `ErrorChain_NilReceiver_testcases.go` | CaseV1 standalone var | 1 |
| 30 | `coretaskinfotests` | `InfoCreate_NilReceiver_testcases.go` | CaseV1 standalone vars | 12 |
| 31 | `coreinstructiontests` | `IdentifiersWithGlobals_NilReceiver_testcases.go` | CaseV1 standalone var | 1 |
| 32 | `coreinstructiontests` | `FromTo_NilReceiver_testcases.go` | CaseV1 standalone vars | 6 |
| 33 | `corestrtests` | `BugfixRegression_NilReceiver_testcases.go` | CaseV1 standalone var | 1 |

---

## Summary

| Category | Files | Est. Cases | Status |
|----------|-------|-----------|--------|
| ✅ Priority A (inline `t.Error`) | 13 | ~70 | **Done** |
| ✅ Priority A (CaseV1/custom) | 6 | ~20 | **Done** |
| ✅ Self-test | 1 | 12 | **Done** |
| ✅ Priority B (CaseV1 nil) | 13 | ~52 | **Done** |
| **Total** | **33** | **~154** | **100% ✅** |

---

## Notes

- **Architecture corrected**: All migrated files use `results.ResultAny` for `Expected` (not `args.Map`). Assertion logic lives on `Result[T].ShouldMatchResult`, with `CaseNilSafe.ShouldBeSafe` as a thin convenience.
- Generic types (`Hashset[T]`, `Hashmap[K,V]`, `Pair[A,B]`, `Triple[A,B,C]`, `TypedPayloadCollection[T]`, `TypedSimpleGenericRequest[T]`) use the **function literal wrapper** pattern documented in the design doc §7.
- Original CaseV1 nil cases remain in their source files alongside non-nil cases for backward compatibility. The CaseNilSafe duplicates provide structured, standardized nil-safety coverage.
- `Expected` uses `results.ResultAny` with `CompareFields` for subset assertion.
- `ExpectAnyError` sentinel is used for methods expected to return non-nil errors.
- `ConcatMessageWithErr` is a package function (not a method), wrapped via `func(_ *struct{})` pattern.
