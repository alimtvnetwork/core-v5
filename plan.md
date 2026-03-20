# Plan — Future Work Roadmap

## Last Updated: 2026-03-20T12:00:00+08:00

---

## Status Overview

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 (Foundation) | ✅ Done | `interface{}` → `any`, Go 1.24, bug fixes |
| Phase 2 (Generics — Collections) | ✅ Done | Collection[T], Hashset[T], Hashmap[K,V], SimpleSlice[T], LinkedList[T] |
| Phase 3 (Generics — Payload/Dynamic) | ✅ Done | TypedPayloadWrapper[T], TypedDynamic[T], generic deserialize helpers |
| Phase 4 (Test Coverage Expansion) | ✅ Done | `conditional/`, `errcore/`, `converters/` expanded |
| Phase 5 (File Splitting) | ✅ Done | PayloadWrapper, Attributes, Info, Dynamic, BaseTestCase |
| Phase 6 (Value Receiver Migration) | ✅ Done | issetter, coreversion, corepayload; remaining audited |
| Phase 7 (Expert Code Review Fixes) | ✅ Done | 16 findings across 4 sub-phases |
| Phase 8 (Deep Quality Sweep) | ✅ Done | ~190 inline negation refactors, bug fixes, regression tests |
| Error Modernization | ✅ Done | errors.Join, errors.Is/As, fmt.Errorf with %w |
| Go Modernization (Phases 1-7) | ✅ Done | All 7 phases complete including slog, legacy removal |
| Test Title Audit (Batches 1-5) | ✅ Done | ~375+ titles renamed |
| Package READMEs | ✅ Done | All core packages documented |
| Phase A (Coverage Stabilization) | ✅ Done | 20-iteration plan executed; all non-internal packages covered |
| Phase B.2 (Spec Reconciliation) | ✅ Done | 9 spec files cleaned |

---

## Phase B: Code Cleanup

### B.1 — Codegen Removal ⏭️ SKIPPED / DEFERRED
- **Objective**: Remove deprecated `codegen/` entirely
- **Dependencies**: User runs external audit (`grep` across auk-go repos)
- **Expected outputs**: Deleted `codegen/`, `cmd/main/unitTestGenerator.go`, `tests/integratedtests/codegentests/`; updated `go.mod`
- **Acceptance criteria**: All exit criteria in `spec/01-app/10-codegen-deprecation-plan.md` met
- **Note**: User chose to skip (2026-03-16). Revisit when convenient.

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

### C.4 — Module Splitting
- **Objective**: Evaluate splitting monorepo into focused Go modules
- **Dependencies**: All coverage work complete ✅
- **Expected outputs**: Architecture decision doc
- **Acceptance criteria**: Decision documented with migration path
- **Spec reference**: `spec/01-app/15-code-review-report.md`

---

## Phase D: Tooling & Runner Improvements

### D.1 — Test Title Audit — Remaining Packages
- **Objective**: Audit remaining 17 packages for test title consistency
- **Dependencies**: None
- **Acceptance criteria**: All test titles follow `"{Function} returns {Result} -- {Input Context}"` format

### D.2 — Diagnostic Output Regression Tests
- **Objective**: Create snapshot tests for diagnostic output formatting
- **Dependencies**: None
- **Acceptance criteria**: Snapshot tests pass for all formatter outputs

---

## Next Task Selection

| # | Task | Effort | Risk |
|---|------|--------|------|
| 1 | **B.1 — Codegen removal** | Medium | Low (deferred — needs user audit) |
| 2 | **C.3 — CI pipeline** | Medium | Low |
| 3 | **C.1 — Generic interfaces** | Medium | Low |
| 4 | **C.2 — iter adoption** | Small | Low |
| 5 | **C.4 — Module splitting** | Large | Medium |
| 6 | **D.1 — Test title audit** | Small | Low |
| 7 | **D.2 — Diagnostic snapshots** | Small | Low |

**Recommended**: **C.3** (CI pipeline) or **D.1** (test title audit) as the next high-value, low-risk tasks.

---

## Process Rules (Mandatory for Any AI)

1. **Read source before every test edit.** Never infer APIs from naming patterns.
2. **One package at a time.** Fix → compile verify → move on.
3. **Do not trust coverage percentages while blockers exist.** Fix blockers first.
4. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
5. **Do not bulk-create coverage suites.** Especially for `errcore`, `corejson`, `corepayload`, `coredynamic`, `corestr`.
6. **Honor naming standards.** Coverage tests: `Test_Cov[N]_{Method}_{Context}`. Titles: `"{Function} returns {Result} -- {Input Context}"`.
7. **Honor project behavior standards.** Vacuous truth (`All*` on empty = true, `Any*` on empty = false), nil-handling, byte-slice clone.
