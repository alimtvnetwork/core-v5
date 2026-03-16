# Reliability & Failure-Chance Report

## Date: 2026-03-16 (Refreshed)
## Scope: Full spec set for `github.com/alimtvnetwork/core`

---

## 1. Success Probability Estimates

### By Module Complexity Tier

| Tier | Modules | Success Probability | Assumptions |
|------|---------|:-------------------:|-------------|
| **Simple** (mechanical, well-scoped) | `interface{}` → `any` ✅, Go 1.24 update ✅, README rewrites ✅, deprecation notices ✅, slog adoption ✅, error modernization ✅ | **95%** | Already completed — confirmed success. |
| **Medium** (multi-file, API-aware) | File splitting ✅, value receiver migration ✅ (S-006 done), codegen removal (Track B open), per-package READMEs ✅ (S-007 done), test title audit ✅ (S-004 done) | **80%** | Requires reading source to verify method signatures. Codegen removal needs external audit. |
| **Complex / Agentic** (coverage push, reflection-heavy packages) | 100% coverage push for 20 remaining packages (Tiers 1-3), especially `corestr` (3.3%), `coredynamic` (0.9%), `corejson` (45%), `corepayload` (56%) | **45-55%** | Coverage work has a **documented root cause of repeated failure**: assumed APIs, bulk generation, build cascades. Reflection-heavy packages resist test generation without deep source reading. |
| **End-to-End** (full verification pipeline) | Write tests → `./run.ps1 PC` → fix mismatches → `./run.ps1 TC` → confirm % | **35-45%** | AI cannot run Go/PowerShell in sandbox. Every cycle requires user-side verification. Latency between write and verify amplifies error accumulation. |

### Key Global Assumptions

1. AI **cannot compile or run Go tests** — all Go work is write-only until user verifies.
2. AI **must read source before every test edit** — naming-pattern inference is the #1 root cause of failure.
3. Spec set is extensive (25+ spec files, 40+ bug docs, 8 testing guidelines) but **some cross-references are stale**.

---

## 2. Failure Map

### 2.1 Where Failures Are Likely

| Module / Workflow | Likelihood | Why | Symptoms |
|---|:---:|---|---|
| **Coverage: `coredynamic`** (57 files, 0.9% avg) | **VERY HIGH** | 53 files at 0%. Reflection-heavy, complex generics, Dynamic typing. Largest uncovered package. | Massive API mismatch. Tests won't compile. Build cascade blocks other packages. |
| **Coverage: `corestr`** (52 files, 3.3% avg) | **VERY HIGH** | 42 files at 0%. Collection/Hashmap/Hashset/LinkedList. Many data structure methods. | Wrong method signatures, missing type fixtures. |
| **Coverage: `corejson`** (18 files, 45% avg) | **HIGH** | Serialization/deserialization with generics and reflection. | Type assertion failures, wrong generic parameters. |
| **Coverage: `corepayload`** (23 files, 56% avg) | **HIGH** | Typed generics, complex collection methods, JSON interop. | Wrong factory function signatures, paging logic errors. |
| **Coverage: Tier 1 quick wins** (6 packages, 90-96%) | **MEDIUM** | Targeted branch coverage. Risk: missing edge cases in nil/boundary logic. | Tests compile but don't cover intended branches. |
| **Codegen removal** (Track B) | **MEDIUM** | External consumer audit required. AI cannot verify. | Broken imports in unknown downstream repos. |
| **Batch 4 verification** (6 unverified files) | **MEDIUM** | Written but never compiled. `coreindexes`, `coremath`, `corecsv`, `intunique`, `stringutil`, `conditional`. | API mismatches discovered when user runs `./run.ps1 PC`. |

### 2.2 Cross-File Inconsistency Issues (Resolved vs Remaining)

| Issue | Status | Action |
|---|---|---|
| `plan.md` showed completed items as pending | **FIXED** in this update | plan.md rewritten with accurate statuses |
| `20-improvement-plan.md` Phase 6 says "In Progress" but S-006 is done | **FIXED** in this update | Noted in suggestions tracker |
| `15-code-review-report.md` recommends "update go.mod to 1.22+" (already at 1.24) | **STALE** | Low priority — report is historical |
| Coverage batches 1-3 ✅ in workflow vs suggestions "blocked by compile" | **CLARIFIED** | Batches 1-3 were written; batch 4 pending verification |

### 2.3 How Failures Manifest

1. **Silent compilation failures** — Test files created, user gets wall of errors from `./run.ps1 PC`.
2. **Coverage regression** — Blocked test packages make coverage numbers drop across the board.
3. **API mismatch cascade** — One wrong method signature blocks the entire integrated test package.
4. **Duplicate work** — Stale specs cause AI to re-implement completed phases.

---

## 3. Corrective Actions (Prioritized)

| # | Fix | Where | Reliability Gain |
|---|-----|-------|:---:|
| 1 | **Verify Batch 4 compilation** — Run `./run.ps1 PC` to validate 6 unverified coverage files | User action → report results | +15% — Establishes real baseline |
| 2 | **One-package-at-a-time gate for Tier 3** — Never bulk-generate coverage for `corestr`, `coredynamic`, `corejson` | Process rule in workflow memory | +15% — Prevents cascade failures |
| 3 | **Add method signature snapshots** — Before writing tests for any HIGH RISK package, create a method-signature inventory file | New spec per package | +10% — Prevents assumed-API errors |
| 4 | **Reconcile plan.md** (done in this update) | `plan.md` | +5% — Prevents duplicate work |
| 5 | **Mark stale code review recommendations** | `spec/01-app/15-code-review-report.md` | +2% — Reduces confusion |
| 6 | **External codegen audit** — User must run `grep` across auk-go repos before removal | User action | +3% — Prerequisite for Track B |

---

## 4. Readiness Decision

### Verdict: **CONDITIONALLY READY** ⚠️

**Strengths:**
- ✅ Comprehensive spec set (25+ architecture docs, 40+ bug audits, 8 testing guidelines)
- ✅ Well-documented postmortem explaining root cause of repeated failures
- ✅ Clear improvement plan with 8 completed phases
- ✅ Testing framework with AAA pattern, naming conventions, and branch coverage strategy
- ✅ Suggestions tracker with structured completion handling

**Before starting implementation:**
1. **MUST**: Run `./run.ps1 PC` to validate Batch 4 files (S-001)
2. **MUST**: Follow one-package-at-a-time gate for all coverage work
3. **SHOULD**: Create method signature inventory for HIGH RISK packages before writing tests
4. **SHOULD**: Accept that coverage for `coredynamic` and `corestr` will take 5-8 sessions each

### Overall Success Rate

| Scenario | Estimate |
|:---|:---:|
| Handed to another AI as-is (before corrective fixes) | **50-55%** |
| After corrective fixes + process enforcement | **70-75%** |
| With user-side verification loop (./run.ps1 PC/TC after each batch) | **80-85%** |

The 30% gap between as-is and best-case is: stale specs (~5%), assumed APIs (~10%), no compile verification (~10%), cascade from bulk generation (~5%).

---

## Related Documents

- [Improvement Plan](../../spec/01-app/20-improvement-plan.md)
- [Coverage Remediation Root Cause](../workflow/completed/02-coverage-remediation-root-cause.md)
- [Coverage File-Level Plan](../workflow/03-coverage-file-level-plan.md)
- [Branch Coverage Strategy](../../spec/01-app/23-branch-coverage-strategy.md)
- [Suggestions Tracker](../suggestions/01-suggestions-tracker.md)
