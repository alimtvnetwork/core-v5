# Plan — Future Work Roadmap

## Last Updated: 2026-03-16T09:50:00+08:00

---

## Status Overview

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 (Foundation) | ✅ Done | `interface{}` → `any`, Go 1.24, bug fixes |
| Phase 2 (Generics — Collections) | ✅ Done | Collection[T], Hashset[T], Hashmap[K,V], SimpleSlice[T], LinkedList[T] |
| Phase 3 (Generics — Payload/Dynamic) | ✅ Done | TypedPayloadWrapper[T], TypedDynamic[T], generic deserialize helpers |
| Phase 4 (Test Coverage Expansion) | ✅ Done (P0) | `conditional/`, `errcore/`, `converters/` expanded |
| Phase 5 (File Splitting) | ✅ Done | PayloadWrapper, Attributes, Info, Dynamic, BaseTestCase |
| Phase 6 (Value Receiver Migration) | ✅ Done | issetter, coreversion, corepayload; remaining audited — no convertible methods |
| Phase 7 (Expert Code Review Fixes) | ✅ Done | 16 findings across 4 sub-phases |
| Phase 8 (Deep Quality Sweep) | ✅ Done | ~190 inline negation refactors, bug fixes, regression tests |
| Error Modernization | ✅ Done | errors.Join, errors.Is/As, fmt.Errorf with %w |
| Go Modernization (Phases 1-7) | ✅ Done | All 7 phases complete including slog, legacy removal |
| Test Title Audit (Batches 1-5) | ✅ Done | ~375+ titles renamed |
| Package READMEs | ✅ Done | All core packages documented |

---

## Phase A: Coverage Stabilization (CURRENT PRIORITY)

### A.0 — Compile Baseline Refresh ⚠️ USER ACTION REQUIRED
- **Objective**: Run `./run.ps1 PC` to validate all existing test files compile, especially Batch 4
- **Dependencies**: None (user action)
- **Expected outputs**: Blocked package list; Batch 4 file compilation status
- **Acceptance criteria**: `./run.ps1 PC` runs; results shared and documented
- **Suggestion ref**: S-001, S-002

### A.1 — Phase 1 Quick Wins (6 packages, ~195 test cases)
- **Objective**: Push 6 near-100% packages to full branch coverage
- **Dependencies**: A.0 (compile baseline must be clean)
- **Expected outputs**: 6-8 new coverage test files
- **Acceptance criteria**: `./run.ps1 TC` shows 100% for each
- **Packages**: `coreonce` (95.7%), `keymk` (95.6%), `corerange` (94.3%), `enumimpl` (95.9%), `corevalidator` (91.2%), `stringslice` (90.6%)
- **Suggestion ref**: S-003

### A.2 — Phase 2 Moderate Effort (5 packages, ~215 test cases)
- **Objective**: Push 5 medium-gap packages to 100%
- **Dependencies**: A.1
- **Expected outputs**: 5-10 new coverage test files
- **Acceptance criteria**: `./run.ps1 TC` shows 100% for each
- **Packages**: `errcore` (90.2%), `reflectmodel` (72.6%), `reflectinternal` (80.4%), `corejson` (45%) ⚠️, `corepayload` (56%) ⚠️
- **Suggestion ref**: S-004

### A.3 — Phase 3 Heavy Lift (4 packages, ~365+ test cases)
- **Objective**: Full coverage for largest uncovered packages
- **Dependencies**: A.2
- **Expected outputs**: 10-20+ new coverage test files
- **Acceptance criteria**: `./run.ps1 TC` shows 100% for each
- **Packages**: `codestack` (0%), `corecmp` (10.8%), `corestr` (3.3%) ⚠️, `coredynamic` (0.9%) ⚠️
- **Suggestion ref**: S-005

---

## Phase B: Code Cleanup

### B.1 — Codegen Removal ⏭️ SKIPPED / DEFERRED
- **Objective**: Remove deprecated `codegen/` entirely
- **Dependencies**: User runs external audit (`grep` across auk-go repos)
- **Expected outputs**: Deleted `codegen/`, `cmd/main/unitTestGenerator.go`, `tests/integratedtests/codegentests/`; updated `go.mod`
- **Acceptance criteria**: All exit criteria in `spec/01-app/10-codegen-deprecation-plan.md` met
- **Suggestion ref**: S-006
- **Note**: User chose to skip this task (2026-03-16). Revisit when coverage stabilization is complete.

### B.2 — Spec Reconciliation ✅ DONE
- **Objective**: Remove stale entries from spec files
- **Dependencies**: None
- **Expected outputs**: Updated spec files with accurate status markers
- **Acceptance criteria**: No spec references completed work as pending
- **Suggestion ref**: S-007
- **Completed**: 2026-03-16 — Fixed 9 spec files

---

## Phase C: Future Architecture (Low Priority)

### C.1 — Generic Interfaces in `coreinterface/`
- **Objective**: Evaluate `ValueGetter[T]` generic interfaces
- **Dependencies**: None
- **Expected outputs**: Architecture decision doc
- **Acceptance criteria**: Decision documented with rationale
- **Spec reference**: `spec/01-app/15-code-review-report.md`

### C.2 — `iter` Package Adoption (Go 1.23+)
- **Objective**: Use `iter.Seq` for collection iteration patterns
- **Dependencies**: None
- **Expected outputs**: Prototype in `coregeneric/`
- **Acceptance criteria**: Working iterator pattern with tests
- **Spec reference**: `spec/01-app/11-go-modernization.md`

### C.3 — CI Pipeline
- **Objective**: Add `golangci-lint`, test coverage, and security scanning
- **Dependencies**: None
- **Expected outputs**: CI config file, lint config
- **Acceptance criteria**: CI runs on push, blocks on failures
- **Suggestion ref**: S-008

### C.4 — Module Splitting
- **Objective**: Evaluate splitting monorepo into focused Go modules
- **Dependencies**: All coverage work complete
- **Expected outputs**: Architecture decision doc
- **Acceptance criteria**: Decision documented with migration path
- **Spec reference**: `spec/01-app/15-code-review-report.md`

---

## Next Task Selection

Pick one to implement next:

| # | Task | Effort | Risk | Prerequisite |
|---|------|--------|------|-------------|
| 1 | **A.0 — Run `./run.ps1 PC`** | User action | None | — |
| 2 | **A.1 — Coverage Phase 1: `coreonce`** | Small (30 tests) | Low | A.0 results |
| 3 | **A.1 — Coverage Phase 1: `keymk`** | Small (20 tests) | Low | A.0 results |
| 4 | **B.1 — Codegen removal** | Medium | Medium | User runs grep audit |
| 5 | ~~**B.2 — Spec reconciliation**~~ | ✅ Done | — | — |
| 6 | **C.3 — CI pipeline** | Medium | Low | — |

**Recommended**: Start with **A.0** (user runs `./run.ps1 PC`), then pick from **A.1** packages one at a time.

---

## Process Rules (Mandatory for Any AI)

1. **Read source before every test edit.** Never infer APIs from naming patterns.
2. **One package at a time.** Fix → compile verify → move on.
3. **Do not trust coverage percentages while blockers exist.** Fix blockers first.
4. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
5. **Do not bulk-create coverage suites.** Especially for `errcore`, `corejson`, `corepayload`, `coredynamic`, `corestr`.
6. **Honor naming standards.** Coverage tests: `Test_Cov[N]_{Method}_{Context}`. Titles: `"{Function} returns {Result} -- {Input Context}"`.
7. **Honor project behavior standards.** Vacuous truth (`All*` on empty = true, `Any*` on empty = false), nil-handling, byte-slice clone.
