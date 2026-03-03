# GoConvey → CaseV1 Migration Plan

> Prioritized migration roadmap for all remaining inline goconvey test files.
> Ordered by complexity (ascending), grouped into waves for batch execution.

## Status Summary

| Status | Count |
|--------|-------|
| ✅ Already migrated | ~30+ files across packages |
| 🔲 Remaining goconvey files | **27 files** |

---

## Wave 1 — Trivial (≤50 lines, simple boolean/struct assertions)

Straightforward 1:1 mapping to CaseV1. No dependencies, no filesystem side effects.

| # | File | Lines | Package | Notes |
|---|------|-------|---------|-------|
| 1 | `creationtests/enumimpltests/DynamicMap_test.go` | 37 | enumimpltests | Simple constant checks |
| 2 | `corejsontests/Deserializer_Apply_test.go` | 40 | corejsontests | Single test, nil safety |
| 3 | `chmodhelpertests/ApplyOnPath_test.go` | 44 | chmodhelpertests | Filesystem helper, single test |
| 4 | `chmodhelpertests/LinuxApplyRecursiveOnPath_test.go` | 42 | chmodhelpertests | Linux-only, single test |
| 5 | `corejsontests/Result_IsEmpty_test.go` | 52 | corejsontests | Boolean checks on Result |
| 6 | `chmodhelpertests/VerifyPartialRwxLocations_test.go` | 56 | chmodhelpertests | Struct field verification |
| 7 | `chmodhelpertests/VerifyRwxChmodUsingRwxInstructions_test.go` | 59 | chmodhelpertests | RWX instruction checks |

**Effort:** ~1 session per 3-4 files. **Total: 7 files.**

---

## Wave 2 — Simple (50–110 lines, uniform assertion patterns)

Repetitive boolean/value checks with clear grouping. Easy to table-drive.

| # | File | Lines | Package | Notes |
|---|------|-------|---------|-------|
| 8 | `corevalidatortests/TestValidators_test.go` | 68 | corevalidatortests | Validator construction checks |
| 9 | `corejsontests/Result_IsEqual_test.go` | 85 | corejsontests | Equality comparisons |
| 10 | `simplewraptests/DoubleQuoteWrapElements_test.go` | 89 | simplewraptests | String wrapping output |
| 11 | `enumimpltests/Enum_test.go` | 97 | enumimpltests | Enum value/string checks |
| 12 | `corejsontests/Result_Unmarshal_test.go` | 100 | corejsontests | Unmarshal result checks |
| 13 | `corerangestests/range_test.go` | 101 | corerangestests | Range boundary checks |
| 14 | `corerangestests/within_range_test.go` | 101 | corerangestests | WithinRange boolean checks |
| 15 | `creationtests/enumimpltests/DynamicMapCreation_test.go` | 101 | enumimpltests | Dynamic map factory checks |
| 16 | `corejsontests/New_NewPtr_test.go` | 106 | corejsontests | Constructor nil/valid checks |
| 17 | `coredynamictests/ReflectSetFromTo_InvalidCases_test.go` | 98 | coredynamictests | Custom wrapper + goconvey |

**Effort:** ~1 session per 2-3 files. **Total: 10 files.**

---

## Wave 3 — Medium (110–260 lines, mixed assertion types)

Larger files with multiple method groups. May require multiple CaseV1 slices per file.

| # | File | Lines | Package | Notes |
|---|------|-------|---------|-------|
| 18 | `corerangestests/start_end_test.go` | 127 | corerangestests | StartEnd struct methods |
| 19 | `simplewraptests/TitleCurlyMeta_test.go` | 143 | simplewraptests | JSON + string wrapping |
| 20 | `coreapitests/PageRequest_test.go` | 207 | coreapitests | API pagination logic |
| 21 | `coreinstructiontests/FromTo_test.go` | 247 | coreinstructiontests | Instruction FromTo methods |
| 22 | `coregenerictests/funcs_test.go` | 250 | coregenerictests | Generic utility functions |

**Effort:** ~1 session per 1-2 files. **Total: 5 files.**

---

## Wave 4 — Complex (300+ lines, deep method coverage)

Large files with many distinct method groups, closures, generics, or concurrency patterns.

| # | File | Lines | Package | Notes |
|---|------|-------|---------|-------|
| 23 | `coreinstructiontests/IdentifiersWithGlobals_test.go` | 313 | coreinstructiontests | Complex identifier resolution |
| 24 | `coregenerictests/Hashmap_test.go` | 375 | coregenerictests | Full hashmap API coverage |
| 25 | `coregenerictests/LinkedList_test.go` | 386 | coregenerictests | Linked list operations |
| 26 | `issettertests/ValueLogic_test.go` | 443 | issettertests | ~30 boolean logic permutations |
| 27 | `coregenerictests/Hashset_test.go`* | 572 | coregenerictests | *Verify — may already be migrated* |

**Effort:** ~1 session per file. **Total: 5 files.**

---

## Dependency & Ordering Notes

### No cross-file dependencies
All goconvey files are self-contained within their packages. Migration order is driven purely by complexity, not by import chains.

### Shared test infrastructure already exists
- `coretestcases.CaseV1` — available in all packages
- `args.Map` — available for ArrangeInput encoding
- `errcore.PrintLineDiff` — available for diff diagnostics
- Custom wrapper structs (like `castedResultTestCase`) — create per-file as needed

### Special cases
| File | Consideration |
|------|---------------|
| `ReflectSetFromTo_InvalidCases_test.go` (#17) | Uses `FromToTestWrapper` + `corevalidator.TextValidator`. Keep wrapper, replace goconvey assertions. |
| `chmodhelpertests/ApplyOnPath_test.go` (#3) | Filesystem side effects. Keep `assertSingleChmod` helper. |
| `chmodhelpertests/LinuxApplyRecursiveOnPath_test.go` (#4) | Linux-only build tag may apply. |
| `issettertests/ValueLogic_test.go` (#26) | 30+ tiny tests → collapse into 3-4 CaseV1 slices by method. |
| `Hashset_test.go` (#27) | 572 lines but may already be migrated (no goconvey import found). Verify before starting. |

---

## Recommended Batch Schedule

| Session | Files | Wave | Est. Effort |
|---------|-------|------|-------------|
| 1 | #1–#4 | Wave 1 | Light |
| 2 | #5–#7 | Wave 1 | Light |
| 3 | #8–#11 | Wave 2 | Light-Medium |
| 4 | #12–#14 | Wave 2 | Light-Medium |
| 5 | #15–#17 | Wave 2 | Medium |
| 6 | #18–#19 | Wave 3 | Medium |
| 7 | #20–#21 | Wave 3 | Medium |
| 8 | #22 | Wave 3 | Medium |
| 9 | #23 | Wave 4 | Heavy |
| 10 | #24 | Wave 4 | Heavy |
| 11 | #25 | Wave 4 | Heavy |
| 12 | #26 | Wave 4 | Heavy |

---

## Completion Criteria

- [ ] All 27 files migrated (or #27 confirmed already done → 26)
- [ ] Zero remaining `goconvey/convey` imports under `tests/`
- [ ] Each migrated file has corresponding `_testcases.go`
- [ ] All tests compile and pass
- [ ] `goconvey` dependency removable from `go.mod`
