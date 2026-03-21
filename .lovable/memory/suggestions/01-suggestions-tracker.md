# Suggestions Tracker

## Last Updated: 2026-03-21T12:00:00+08:00

## Convention

- **Location**: `.lovable/memory/suggestions/` — this file for active tracking, `completed/` for archives.
- **File naming**: Single tracker file (`01-suggestions-tracker.md`). Individual completed suggestions in `completed/NN-slug.md`.
- **Statuses**: `open` → `inProgress` → `done`
- **Completion handling**: When done, update status here and optionally create detail file in `completed/`.

---

## Active Suggestions

### S-009: Deprecated API Cleanup
- **suggestionId**: S-009
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: Remove or sunset 110 deprecated functions/methods across 30+ files. Largest concentrations: `coreindexes/indexes.go` (21), `core.go` (13), `coredata/corestr/` (15+), `coredata/corejson/` (6+), `coredata/stringslice/` (5+).
- **rationale**: Deprecated functions add API surface confusion and maintenance cost. Generic replacements already exist for all of them.
- **proposed change**: Phase approach — (1) audit all 110 deprecated markers, (2) confirm generic replacements exist, (3) remove in batches with compile verification.
- **acceptance criteria**: Zero `// Deprecated:` markers remain (or only those with documented external consumers). `./run.ps1 PC` and `TC` pass.
- **status**: open
- **completion notes**: —

### S-010: Performance Benchmarks
- **suggestionId**: S-010
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: Add `Benchmark*` tests for hot-path operations. Currently zero benchmarks exist. Priority targets: `coredata/corestr/Collection` (Add, Get, Join), `coredata/coredynamic` (type casting), `errcore` (error construction with stack traces), `codestack` (trace capture), `regexnew` (lazy compile), `mutexbykey` (lock contention).
- **rationale**: No performance baseline exists. Regressions are invisible without benchmarks.
- **proposed change**: Create `*_bench_test.go` files in priority packages. Include `b.ReportAllocs()`.
- **acceptance criteria**: ≥30 benchmarks across 6+ packages. Results documented in a benchmark summary.
- **status**: open
- **completion notes**: —

### S-011: Missing Package READMEs (10 packages)
- **suggestionId**: S-011
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 10 packages lack README files: `cmdconsts`, `coremath`, `defaultcapacity`, `dtformats`, `extensionsconst`, `filemode`, `iserror`, `osconsts`, `regconsts`, `testconsts`.
- **rationale**: All other packages have READMEs. These are small leaf packages but should be documented for completeness.
- **proposed change**: Create README.md for each with purpose, types/constants, and usage examples.
- **acceptance criteria**: All packages have README.md.
- **status**: open
- **completion notes**: —

### S-012: Pointer Receiver Audit
- **suggestionId**: S-012
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 5,224 pointer receivers vs 2,836 value receivers. Many small readonly methods (getters, checkers, formatters) on immutable types likely use pointer receivers unnecessarily.
- **rationale**: Value receivers are idiomatic for small, read-only types. They enable better compiler optimizations and prevent nil-receiver panics.
- **proposed change**: Audit top packages (`coredata/corestr`, `errcore`, `coredata/corepayload`) for methods that could safely use value receivers.
- **acceptance criteria**: Identified methods migrated without behavior changes. `./run.ps1 TC` passes.
- **status**: open
- **completion notes**: —

### S-013: Sync.Mutex → sync.RWMutex Audit
- **suggestionId**: S-013
- **createdAt**: 2026-03-21
- **source**: Lovable (codebase audit)
- **affectedProject**: core
- **description**: 27 `sync.Mutex` usages found. Read-heavy collection types (Collection, Hashmap, Hashset) may benefit from `sync.RWMutex` for concurrent read performance.
- **rationale**: `RWMutex` allows multiple concurrent readers, improving throughput for read-heavy workloads.
- **proposed change**: Audit each mutex usage. Migrate to `RWMutex` where read methods (Get, Contains, Len, IsEmpty) dominate.
- **acceptance criteria**: Identified candidates migrated. Benchmark showing improvement for read-heavy scenarios.
- **status**: open (depends on S-010 for benchmarks)
- **completion notes**: —

### S-014: Coverage Push — Remaining Packages
- **suggestionId**: S-014
- **createdAt**: 2026-03-21
- **source**: Lovable (carried from S-003/S-004/S-005)
- **affectedProject**: core
- **description**: Continue coverage push for packages below 100%. Requires `./run.ps1 TC` to get current baselines.
- **rationale**: Coverage gaps hide bugs, especially in high-risk packages like `coredynamic` and `corestr`.
- **proposed change**: Run TC → identify gaps → one package at a time → compile gate.
- **acceptance criteria**: All packages at 100% coverage.
- **status**: open
- **completion notes**: —

---

## Completed Suggestions (Archive)

| # | Title | Completed | Notes |
|---|-------|-----------|-------|
| 1 | Diagnostic Formatting Improvements | 2026-03-11 | 4-space indent, separator headers, tab-indented entries |
| 2 | Test Title Audit (Batches 1-5) | 2026-03-16 | ~375+ titles renamed across all listed packages |
| 3 | Fix 21 Failing Tests | 2026-03-11 | All fixed |
| 4 | Coverage Push Batch 1 (11 packages) | 2026-03-14 | Packages 75-97% |
| 5 | Coverage Push Batch 2 (6 packages) | 2026-03-14 | Packages 0-57% |
| 6 | Coverage Push Batch 3 (7 packages) | 2026-03-15 | Generic/utility packages |
| 7 | Coverage Prompt Generator System | 2026-03-15 | PowerShell-based prompt generation |
| 8 | Deep Clone Production Bug Fix | 2026-03-15 | `corepayload` nil AnyMap |
| 9 | Nil Receiver Coverage Audit | 2026-03-15 | All types audited |
| 10 | Test Runner Hardening Review | 2026-03-15 | Verified |
| 11 | Diagnostic Output Regression Tests | 2026-03-15 | Snapshot tests |
| 12 | Coverage Push Batch 4 (6 packages) | 2026-03-16 | ⚠️ Pending PC verification |
| 13 | Value Receiver Migration (Phase 6) | 2026-03-16 | All convertible methods migrated |
| 14 | Remaining Package READMEs | 2026-03-16 | All 5 packages already had READMEs |
| 15 | High-Risk Coverage File Audit (6 files) | 2026-03-16 | Audited, 1 fix in converterstests |
| S-001 | Compile Baseline | 2026-03-16 | Completed as part of coverage stabilization |
| S-002 | Verify Batch 4 | 2026-03-16 | Completed |
| S-006 | Codegen Removal | 2026-03-21 | Fully removed — codegen/, tests, consumers, docs |
| S-007 | Spec Reconciliation | 2026-03-17 | 9 files fixed |
| S-008 | CI Pipeline Setup | 2026-03-18 | GitHub Actions with lint, test, coverage, govulncheck |

> Detail files in `completed/` subfolder.
