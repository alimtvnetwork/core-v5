# Plan — Future Work Roadmap

## Last Updated: 2026-03-16T04:58:00+08:00

---

## Status Overview

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1-3 | ✅ Done | Repo scan, per-folder specs, README upgrades |
| Phase 4 | ✅ Done | Special module docs and conventions |
| Phase 5 | ✅ Done | Codegen deprecation plan + Go modernization plan |
| Phase 6 | ✅ Done | Code review report + improvement backlog |
| Phase 7 | ✅ Done | Package-level README docs |
| Phase 8 | ✅ Done | Code review fixes (return types, typos, README accuracy) |
| Foundation | ✅ Done | `interface{}` → `any`, Go 1.24, bug fixes |
| Generics | ✅ Done | Collection[T], TypedPayloadWrapper[T], TypedDynamic[T] |
| File Splitting | ✅ Done | All large files split (PayloadWrapper, Attributes, Info, Dynamic, BaseTestCase) |
| Deep Quality Sweep | ✅ Done | ~190 inline negation refactors, low-priority bug fixes, regression tests |
| Expert Code Review | ✅ Done | 16 findings across 4 sub-phases, all fixed |
| Error Modernization | ✅ Done | errors.Join, errors.Is/As, fmt.Errorf with %w |

---

## Phase A: Coverage Stabilization (CURRENT PRIORITY)

> **Prerequisite**: User must run `./run.ps1 PC` and share results before AI proceeds.

### A.1 — Compile Baseline Refresh
- **Objective**: Get the real blocked-package list from `./run.ps1 PC`
- **Dependencies**: None (user action)
- **Expected outputs**: Blocked package list documented in workflow memory
- **Acceptance criteria**: `./run.ps1 PC` runs; blocked list captured

### A.2 — Audit High-Risk Coverage Files (6 files)
- **Objective**: Verify each unverified coverage file compiles
- **Dependencies**: A.1
- **Expected outputs**: Fixed test files that compile individually
- **Acceptance criteria**: `go build ./tests/integratedtests/<pkg>/...` passes for each
- **Files**: `errcoretests/Coverage9_test.go`, `simplewraptests/Coverage7_test.go`, `issettertests/Coverage7_test.go`, `isanytests/Coverage9_test.go`, `converterstests/Coverage4_test.go`, `stringslicetests/Coverage7_test.go`

### A.3 — Remaining 12 Packages to 100%
- **Objective**: Push all remaining packages to 100% branch coverage
- **Dependencies**: A.1, A.2
- **Expected outputs**: New coverage test files per package
- **Acceptance criteria**: `./run.ps1 TC` shows 100% for each; no blocked packages
- **Packages**: `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`, `corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`

---

## Phase B: Code Cleanup & Modernization

### B.1 — Codegen Removal
- **Objective**: Remove deprecated codegen package entirely
- **Dependencies**: External consumer audit (user must run grep across auk-go repos)
- **Expected outputs**: Deleted `codegen/`, `cmd/main/unitTestGenerator.go`, `tests/integratedtests/codegentests/`; updated `go.mod`
- **Acceptance criteria**: All exit criteria in `spec/01-app/10-codegen-deprecation-plan.md` met
- **Spec reference**: `spec/01-app/10-codegen-deprecation-plan.md`

### B.2 — Value Receiver Migration (Phase 6 Completion)
- **Objective**: Migrate remaining read-only methods to value receivers
- **Dependencies**: None (can be done incrementally)
- **Expected outputs**: Updated method signatures; verified interface satisfaction
- **Acceptance criteria**: All read-only methods use value receivers; all tests pass
- **Spec reference**: `spec/01-app/20-improvement-plan.md` Phase 6

### B.3 — Test Title Audit (Remaining 17 Packages)
- **Objective**: Rename 1400+ test titles to standard format
- **Dependencies**: None
- **Expected outputs**: Renamed test functions
- **Acceptance criteria**: All titles follow `"{Function} returns {Result} -- {Input Context}"`

---

## Phase C: Documentation Completion

### C.1 — Remaining Package READMEs
- **Objective**: Create README.md for undocumented packages
- **Dependencies**: None
- **Expected outputs**: README.md files with folder trees, API docs, examples
- **Acceptance criteria**: Each README has verified method signatures and usage examples
- **Packages**: `coregeneric`, `corestr`, `coreonce`, `corerange`, `stringslice`

### C.2 — Spec Reconciliation
- **Objective**: Remove stale/contradictory entries from specs
- **Dependencies**: None
- **Expected outputs**: Updated spec files with accurate status markers
- **Acceptance criteria**: No spec file references completed work as pending

---

## Phase D: Future Architecture (Low Priority)

### D.1 — Generic Interfaces in `coreinterface/`
- **Objective**: Evaluate `ValueGetter[T]` generic interfaces
- **Spec reference**: `spec/01-app/15-code-review-report.md`

### D.2 — `iter` Package Adoption (Go 1.23+)
- **Objective**: Use `iter.Seq` for collection iteration patterns
- **Spec reference**: `spec/01-app/11-go-modernization.md`

### D.3 — CI Pipeline
- **Objective**: Add `golangci-lint`, test coverage, and security scanning
- **Spec reference**: `spec/01-app/15-code-review-report.md`

### D.4 — Module Splitting
- **Objective**: Evaluate splitting monorepo into focused Go modules
- **Spec reference**: `spec/01-app/15-code-review-report.md`

---

## Next Task Selection

Pick one of these to implement next:

| # | Task | Effort | Risk | Prerequisite |
|---|------|--------|------|-------------|
| 1 | **A.1 — Run `./run.ps1 PC`** | User action | None | — |
| 2 | **A.2 — Audit 6 high-risk coverage files** | Medium | High (API mismatches) | A.1 results |
| 3 | **B.1 — Remove codegen** | Medium | Medium (external audit needed) | User runs grep |
| 4 | **B.3 — Test title audit** | Large but mechanical | Low | — |
| 5 | **C.1 — Package READMEs** | Medium | Low | — |
| 6 | **B.2 — Value receiver migration** | Small per package | Medium | — |

**Recommended**: Start with **A.1** (user runs compile baseline), then **C.1** or **B.1** while waiting for results.
