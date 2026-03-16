# Reliability & Failure-Chance Report

## Date: 2026-03-16
## Scope: Full spec set for `github.com/alimtvnetwork/core`

---

## 1. Success Probability Estimates

### By Module Complexity Tier

| Tier | Modules | Success Probability | Assumptions |
|------|---------|:-------------------:|-------------|
| **Simple** (mechanical, well-scoped) | `interface{}` → `any` migration ✅, Go version update ✅, README rewrites ✅, deprecation notices ✅ | **95%** | Specs are precise, acceptance criteria are binary (done/not done), no cross-package side effects. Already completed — success confirmed. |
| **Medium** (multi-file, API-aware) | File splitting (Phase 5) ✅, value receiver migration (Phase 6), codegen removal (Track B), per-package READMEs, test title audit | **75-80%** | Requires reading source to verify method signatures. Failure occurs when an AI assumes API shape from naming patterns. Specs exist but some lack explicit function signatures. |
| **Complex / Agentic** (coverage push, generics adoption) | 100% coverage push (Batches 1-3 + remaining), generic Collection/Payload/Dynamic, branch coverage strategy | **50-60%** | Coverage work has a **documented history of repeated failures** (see postmortem). Root cause: assumed APIs, bulk generation without compile gates, build-failure cascades. Generics work requires understanding 40+ packages and their interplay. |
| **End-to-End** (full pipeline: write tests → compile → run → verify coverage) | Coverage verification via `./run.ps1 PC` → `./run.ps1 TC` pipeline | **40-50%** | Requires PowerShell runtime, Go toolchain, and the full repo. AI cannot run these commands in Lovable's sandbox. Verification must happen on the user's machine. |

### Key Assumption Behind All Estimates

- The AI **cannot run Go compilation or tests**. All Go work is write-only — the user must verify.
- The AI **must read source before writing tests** — never infer from naming conventions.
- The spec set is extensive (25+ spec files, 18+ failing-test docs, 8+ testing guideline docs) but **some specs are stale** (reference completed work as pending).

---

## 2. Failure Map

### 2.1 Where Failures Are Likely

| Module / Workflow | Failure Likelihood | Why | Symptoms |
|---|:---:|---|---|
| **Coverage test generation** (remaining 12 packages) | **HIGH** | Documented root cause: APIs assumed from naming, bulk generation without compile check. `errcore`, `issetter`, `corejson`, `corepayload`, `coredynamic` are explicitly flagged as high-risk. | Tests don't compile. Coverage reports show blocked packages. Build-failure cascade across integrated test packages. |
| **Codegen removal** (Track B) | **MEDIUM** | Exit criteria checklist is clear but requires external audit (`grep` across auk-go repos). AI cannot verify external consumers. | Broken imports in unknown downstream repos. Missing rollback tag. |
| **Value receiver migration** (Phase 6) | **MEDIUM** | Requires verifying interface satisfaction after each change. Some methods need nil-safety guards that prevent value receiver use. | Interface compliance failures at compile time. Nil pointer panics at runtime. |
| **Test title audit** (1400+ titles, 40+ packages) | **LOW-MEDIUM** | Mechanical but massive scope. Risk: inconsistent application, missing edge cases in naming convention. | Titles don't follow `"{Function} returns {Result} -- {Input Context}"` format. Partial completion. |
| **Documentation** (remaining READMEs) | **LOW** | Specs clearly list which packages need READMEs. Risk: API drift if README written without reading current source. | Inaccurate usage examples. Missing methods in docs. |

### 2.2 Cross-File Inconsistency Issues

| Issue | Location | Impact |
|---|---|---|
| `plan.md` backlog says `interface{} → any` is pending but it's ✅ complete | `plan.md` line 151 vs `20-improvement-plan.md` Phase 1 | AI may re-do completed work |
| `plan.md` says "Split BaseTestCase.go" is 🟡 Medium priority but it's ✅ done | `plan.md` line 155 vs `20-improvement-plan.md` Phase 5 | Wasted effort |
| Workflow plan lists tasks 10-13 as pending but suggestions tracker marks 7-9 as completed | `.lovable/memory/workflow/01-*` vs `.lovable/memory/suggestions/01-*` | Contradictory status |
| `15-code-review-report.md` still shows "update go.mod to 1.22+" as pending | `15-code-review-report.md` line 57 | Stale recommendation |
| Coverage batches 1-3 marked ✅ in workflow but suggestions say "Blocked by compile verification" | Workflow completed section vs suggestions #3 | Unclear what's actually verified |

### 2.3 How Failures Would Manifest

1. **Silent compilation failures**: Test files are created but never compiled. User runs `./run.ps1 PC` and gets a wall of errors.
2. **Coverage regression**: Blocked packages cause coverage numbers to drop, creating false alarms.
3. **Duplicate work**: Stale spec files cause AI to re-implement already-completed phases.
4. **API mismatch**: Tests call methods with wrong signatures, wrong parameter counts, or wrong return types.

---

## 3. Corrective Actions (Prioritized)

| # | Fix | Where | Expected Reliability Gain |
|---|-----|-------|:---:|
| 1 | **Reconcile plan.md with improvement-plan.md** — Mark completed items as ✅, remove from backlog | `plan.md`, `20-improvement-plan.md` | +15% — Prevents duplicate work |
| 2 | **Add explicit API signatures to coverage specs** — For each of the 12 remaining packages, list exact method signatures in the spec | `spec/05-failing-tests/`, new per-package coverage specs | +20% — Prevents assumed-API failures |
| 3 | **Mark stale recommendations as DONE in code review report** | `spec/01-app/15-code-review-report.md` | +5% — Reduces confusion |
| 4 | **Establish a "read source first" gate in all coverage specs** — Add a mandatory checklist item: "I have read the source file and confirmed the method signature" | `.lovable/memory/workflow/01-*` | +10% — Process fix for root cause |
| 5 | **Resolve contradictory statuses** — Suggestions #1-3 say "blocked/pending" but completed suggestions include overlapping work | `.lovable/memory/suggestions/01-*` | +5% — Status clarity |
| 6 | **Add compile-verification acceptance criteria to every coverage task** — "Task is not done until `./run.ps1 PC` passes" | All coverage-related specs | +10% — Prevents premature success claims |

---

## 4. Readiness Decision

### Verdict: **CONDITIONALLY READY** ⚠️

The spec set is **architecturally sound** and impressively comprehensive. The project has:
- ✅ Clear folder structure and repo overview
- ✅ Detailed per-package specs with acceptance criteria
- ✅ A well-documented postmortem explaining past failures
- ✅ Testing guidelines, branch coverage strategy, and naming conventions
- ✅ A clear improvement plan with phased milestones

**However, before starting implementation:**

1. **MUST FIX**: Reconcile `plan.md` with `20-improvement-plan.md` to eliminate stale/contradictory status (this report does this — see updated `plan.md`).
2. **MUST FIX**: Update suggestions tracker to reflect true current state.
3. **SHOULD FIX**: Add method signatures to remaining coverage package specs.
4. **SHOULD FIX**: Establish that coverage work requires user-side verification (`./run.ps1 PC`) — AI cannot confirm success alone.

### Estimated Overall Success Rate for Remaining Work

| If handed to another AI as-is | After corrective fixes |
|:---:|:---:|
| **55-60%** | **75-80%** |

The 20% gap is almost entirely due to stale specs causing duplicate/wrong work and the coverage generation root cause not being enforced by the spec structure itself.

---

## Related Documents

- [Improvement Plan](../../spec/01-app/20-improvement-plan.md)
- [Coverage Remediation Root Cause](../workflow/completed/02-coverage-remediation-root-cause.md)
- [Code Review Report](../../spec/01-app/15-code-review-report.md)
- [Branch Coverage Strategy](../../spec/01-app/23-branch-coverage-strategy.md)
- [Suggestions Tracker](../suggestions/01-suggestions-tracker.md)
