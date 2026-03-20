# Coverage 100% Push — Segmented Iteration Plan

## Status: 🔧 Active
## Last Updated: 2026-03-20
## Source Data: PC run `blocked-packages-10.json` + existing coverage data

---

## Rules

1. **Internal packages are EXCLUDED** — never write coverage tests for `internal/*`
2. **Large packages (>1000 uncovered stmts)** are split into ~200-line segments
3. **Small packages (<200 uncovered)** are grouped 2-4 per iteration
4. **Each iteration** = one "next" command from the user
5. **Follow AAA pattern**, `CaseV1` / `args.Map` / `ShouldBeEqual` conventions per `spec/testing-guidelines/`
6. **Read source files before writing tests** — never assume APIs
7. **Test file naming**: `Coverage{N}_test.go` in `tests/integratedtests/{pkg}tests/`
8. **Function naming**: `Test_Cov{N}_{Method}_{Context}`
9. **ShouldBeEqual** is on `CaseV1` / `SimpleTestCase`, NOT on `coretests.GetAssert`

---

## Blocked Packages (Must Fix First)

9 blocked packages from PC run. Root cause: `coretests.GetAssert.ShouldBeEqual` undefined.

| Package | File(s) | Error | Fix |
|---------|---------|-------|-----|
| `coreoncetests` | `Coverage12_Iteration1_test.go` | `GetAssert.ShouldBeEqual` undefined (6 calls) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `corepropertytests` | `Coverage2_Iteration2_test.go` | `GetAssert.ShouldBeEqual` undefined (2 calls) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `corerangetests` | `Coverage7_Iteration2_test.go` | `GetAssert.ShouldBeEqual` undefined (1 call) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `corestrtests` | `Coverage33_LinkedList_LinkedColl_test.go` | `LinkedList` API mismatch (Collection, SimpleSlice, LoopLock, etc.) | Read source, fix all API calls |
| `coretaskinfotests` | `Coverage3_Iteration2_test.go` | `GetAssert.ShouldBeEqual` undefined (1 call) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `coreversiontests` | `Coverage5_Iteration2_test.go` | `GetAssert.ShouldBeEqual` undefined (2 calls) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `resultstests` | `Coverage3_Iteration1_test.go` | `GetAssert.ShouldBeEqual` + `InvokeWithPanicRecovery` wrong signature | Rewrite both |
| `simplewraptests` | `Coverage8_Iteration2_test.go` | `GetAssert.ShouldBeEqual` undefined (2 calls) | Rewrite to use `CaseV1.ShouldBeEqual` |
| `stringslicetests` | `Coverage12_Iteration1_test.go` | `GetAssert.ShouldBeEqual` undefined (4 calls) | Rewrite to use `CaseV1.ShouldBeEqual` |

---

## Package Summary (Non-Internal, Below 100%)

| Package | Coverage | Uncovered Stmts | Segments |
|---------|----------|-----------------|----------|
| `namevalue` | 0.0% | 188 | 1 |
| `reflectcore/reflectmodel` | 0.8% | 251 | 2 |
| `coredata/corestr` | 3.6% | 5553 | **28** |
| `coredata/coredynamic` | 3.7% | 2191 | **11** |
| `regexnew` | 87.0% | 29 | 1 |
| `chmodhelper` | 89.6% | 170 | 1 |
| `coretests` | 90.5% | 35 | 1 |
| `coretests/args` | 90.9% | 156 | 1 |
| `coredata/corepayload` | 92.9% | 117 | 1 |
| `errcore` | 93.6% | 53 | 1 |
| `coredata/coregeneric` | 94.7% | 57 | 1 |
| `coredata/corejson` | 95.0% | 106 | 1 |
| `corecmp` | 95.1% | 9 | 1 |
| `codestack` | 95.2% | 24 | 1 |
| `corevalidator` | 95.4% | 33 | 1 |
| `coretests/coretestcases` | 95.9% | 11 | 1 |
| `coreinstruction` | 95.9% | 16 | 1 |
| `codegen/coreproperty` | 96.2% | 2 | 1 |
| `coreimpl/enumimpl` | 96.3% | 55 | 1 |
| `coretests/results` | 96.6% | 5 | 1 |
| `iserror` | 97.4% | 1 | 1 |
| `coreutils/stringutil` | 98.0% | 9 | 1 |
| `simplewrap` | 98.1% | 2 | 1 |
| `keymk` | 98.5% | 6 | 1 |
| `coremath` | 98.5% | 1 | 1 |
| `enums/versionindexes` | 98.6% | 1 | 1 |
| `reqtype` | 99.1% | 2 | 1 |
| `coreversion` | 99.2% | 3 | 1 |
| `coredata/stringslice` | 99.2% | 4 | 1 |
| `coretaskinfo` | 99.2% | 2 | 1 |
| `coredata/coreonce` | 99.3% | 5 | 1 |
| `isany` | 99.4% | 1 | 1 |
| `ostype` | 99.4% | 1 | 1 |
| `issetter` | 99.6% | 1 | 1 |
| `coredata/corerange` | 99.7% | 2 | 1 |
| `converters` | 99.8% | 1 | 1 |

**Total uncovered: ~9,179 stmts across 36 packages**

---

## Iteration Plan

### Iteration 0 — Fix 9 Blocked Packages (PREREQUISITE)

Fix all 9 compile-blocked packages listed above. No new coverage — just make them build.
- 8 packages: replace `coretests.GetAssert.ShouldBeEqual` with correct `CaseV1.ShouldBeEqual` pattern
- 1 package (`corestrtests`): fix `LinkedList` API calls in `Coverage33` file
- 1 package (`resultstests`): fix `InvokeWithPanicRecovery` signature

### Phase 1 — Quick Wins (≤10 uncovered stmts each) — Iterations 1-3 ✅ DONE

**Iteration 1** ✅: `corecmp` (9), `coretests/results` (5), `coredata/coreonce` (5), `coredata/stringslice` (4) → 23 stmts
**Iteration 2** ✅: `coreversion` (3), `coretaskinfo` (2), `codegen/coreproperty` (2), `simplewrap` (2), `coredata/corerange` (2), `reqtype` (2) → 14 stmts
**Iteration 3** ✅: `keymk` (6), `converters` (1), `ostype` (1), `enums/versionindexes` (1) → 13 stmts

### Phase 2 — Small-Medium Packages — Iterations 4-9 ✅ DONE

**Iteration 4** ✅: `coreinstruction` (16), `coretests/coretestcases` (11) → 27 stmts
**Iteration 5** ✅: `codestack` (24), `regexnew` (29) → 53 stmts
**Iteration 6** ✅: `corevalidator` (33), `coretests` (35) → 68 stmts
**Iteration 7** ✅: `errcore` (53), `coreimpl/enumimpl` (55) → 108 stmts
**Iteration 8** ✅: `coredata/coregeneric` (57), `coredata/corejson` (106) → 163 stmts
**Iteration 9** ✅: `coretests/args` (156) → 156 stmts

### Phase 3 — Medium Packages — Iterations 10-14

**Iteration 10** ✅: `chmodhelper` (170)
→ 170 stmts — Coverage15_Iteration10_test.go (SFRW read/write/json/clone/expire, RwxInstructionExecutor, RwxInstructionExecutors, AttrVariant, FilteredPathFileInfoMap, RecursivePathsApply, fileBytesWriter, anyItemWriter, fileReader, newSimpleFileReaderWriterCreator)

**Iteration 11** ✅: `coredata/corepayload` (117), `coreutils/stringutil` (9)
→ 126 stmts — Coverage16_Iteration11_test.go (newPayloadWrapperCreator: UsingBytesCreateInstruction, TypeStringer variants, CreateUsingTypeStringer, NameIdCategoryStringer, RecordsTypeStringer, RecordTypeStringer, ManyRecords, CastOrDeserializeFrom valid, DeserializeToMany/Collection, string payload branch; PayloadsCollection: Dynamic accessors, SkipDynamic/TakeDynamic/LimitDynamic, SkipCollection/TakeCollection/LimitCollection, IsEqualItems, GetPagedCollection, GetSinglePageCollection, DeserializeMust/ToMany/UsingJsonResult; structs: PayloadCreateInstructionTypeStringer, PayloadTypeExpander, BytesCreateInstructionStringer; emptyCreator; stringutil: Coverage7_Iteration11_test.go)

**Iteration 12** ✅: `namevalue` (188)
→ 188 stmts — Coverage7_Iteration12_test.go (exercised all type aliases: StringInt, StringMapAny, StringMapString through every Collection method; Instance type-specific String/JsonString/Dispose/IsNull; AppendsIf/PrependsIf per type; IsEqualByString, Error, LazyString, FuncIf, AppendPrependIf, CollectionUsing per type)

**Iteration 13** ✅: `reflectcore/reflectmodel` — Segment A (first ~200 of 251)
→ ~200 stmts — Coverage10_Iteration13_test.go (Invoke with ptr/slice/map/chan/func/interface nil+non-nil returns to exercise ReflectValueToAnyValue + IsNull branches; InvokeFirstAndError/InvokeError with various return types; cached GetOutArgsTypes/GetInArgsTypes/GetInArgsTypesNames; ValidateMethodArgs multi-type-mismatch; VerifyInArgs/VerifyOutArgs; InArgsVerifyRv/OutArgsVerifyRv match+mismatch; ReflectValueKind with Int/Bool/Slice/Map/Struct/Ptr kinds; FieldProcessor Int+Bool fields; ReflectValue struct)

**Iteration 14** ✅: `reflectcore/reflectmodel` — Segment B (remaining ~51)
→ ~51 stmts — Coverage11_Iteration14_test.go (nil receiver paths for MethodProcessor/ReflectValueKind/FieldProcessor; IsEqual nil+same-pointer+same-signature+diff-signature; InvalidReflectValueKindModel constructor; InvokeFirstAndError single-return error; InvokeError non-error panic; ValidateMethodArgs too-few/too-many/empty; GetFirstResponseOfInvoke success; InvokeResultOfIndex second result; RVK invalid-not-nil PointerRv)

### Phase 4 — coredata/coredynamic (2191 uncovered) — Iterations 15-25

Each segment = ~200 stmts, aligned to logical file/function boundaries.

| Iteration | Segment | Stmts |
|-----------|---------|-------|
| 15 | `coredynamic` Segment A | ~200 |
| 16 | `coredynamic` Segment B | ~200 |
| 17 | `coredynamic` Segment C | ~200 |
| 18 | `coredynamic` Segment D | ~200 |
| 19 | `coredynamic` Segment E | ~200 |
| 20 | `coredynamic` Segment F | ~200 |
| 21 | `coredynamic` Segment G | ~200 |
| 22 | `coredynamic` Segment H | ~200 |
| 23 | `coredynamic` Segment I | ~200 |
| 24 | `coredynamic` Segment J | ~200 |
| 25 | `coredynamic` Segment K | ~191 |

### Phase 5 — coredata/corestr (5553 uncovered) — Iterations 26-53

Each segment = ~200 stmts, aligned to logical file/function boundaries.

| Iteration | Segment | Stmts |
|-----------|---------|-------|
| 26 | `corestr` Segment A | ~200 |
| 27 | `corestr` Segment B | ~200 |
| 28 | `corestr` Segment C | ~200 |
| 29 | `corestr` Segment D | ~200 |
| 30 | `corestr` Segment E | ~200 |
| 31 | `corestr` Segment F | ~200 |
| 32 | `corestr` Segment G | ~200 |
| 33 | `corestr` Segment H | ~200 |
| 34 | `corestr` Segment I | ~200 |
| 35 | `corestr` Segment J | ~200 |
| 36 | `corestr` Segment K | ~200 |
| 37 | `corestr` Segment L | ~200 |
| 38 | `corestr` Segment M | ~200 |
| 39 | `corestr` Segment N | ~200 |
| 40 | `corestr` Segment O | ~200 |
| 41 | `corestr` Segment P | ~200 |
| 42 | `corestr` Segment Q | ~200 |
| 43 | `corestr` Segment R | ~200 |
| 44 | `corestr` Segment S | ~200 |
| 45 | `corestr` Segment T | ~200 |
| 46 | `corestr` Segment U | ~200 |
| 47 | `corestr` Segment V | ~200 |
| 48 | `corestr` Segment W | ~200 |
| 49 | `corestr` Segment X | ~200 |
| 50 | `corestr` Segment Y | ~200 |
| 51 | `corestr` Segment Z | ~200 |
| 52 | `corestr` Segment AA | ~200 |
| 53 | `corestr` Segment AB | ~153 |

---

## Total Summary

| Phase | Iterations | Stmts | Status |
|-------|-----------|-------|--------|
| Blocker Fix | 0 | 0 (fix only) | 🔲 Next |
| Phase 1 (Quick Wins) | 1-3 | 50 | ✅ Done |
| Phase 2 (Small-Medium) | 4-9 | 575 | ✅ Done |
| Phase 3 (Medium) | 10-14 | 735 | 🔲 Pending |
| Phase 4 (coredynamic) | 15-25 | 2,191 | 🔲 Pending |
| Phase 5 (corestr) | 26-53 | 5,553 | 🔲 Pending |
| **Total** | **54 iterations** | **~9,104** | |

---

## Execution Protocol (per iteration)

1. **Read** the uncovered source files relevant to the segment
2. **Identify** uncovered branches/paths from coverage data
3. **Write tests** following:
   - AAA format (Arrange / Act / Assert)
   - `CaseV1` + `args.Map` for inputs/expectations
   - `ShouldBeEqual` / `ShouldBeEqualMap` for assertions
   - File: `Coverage{N}_{SegmentLabel}_test.go` in `tests/integratedtests/{pkg}tests/`
   - Function: `Test_Cov{N}_{Method}_{Context}`
4. **Segment boundaries** align to logical file/function boundaries (not arbitrary line numbers)
5. **For Phases 4-5**: each iteration will map specific source files at execution time

## Process Rules (Mandatory)

1. **Read source before every test edit.** Never infer APIs from naming patterns.
2. **One package at a time.** Fix → verify → move on.
3. **Do not trust coverage percentages while blockers exist.** Fix blockers first.
4. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
5. **Do not bulk-create coverage suites.** Especially for large packages.
6. **Honor project behavior standards.** Vacuous truth, nil-handling, byte-slice clone.
7. **Follow spec/testing-guidelines/ for all conventions.**
