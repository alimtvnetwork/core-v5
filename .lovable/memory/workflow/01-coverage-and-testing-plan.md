# Coverage & Testing Master Plan

## Status: In Progress — Coverage Plan Generated, Ready for Phase 1
## Last Updated: 2026-03-16T15:00:00+08:00

## Critical Root Cause Checkpoint
Coverage work has been repeatedly invalidated by assumed APIs, broad unverified coverage-file generation, and skipping the compile-first gate. Do **not** treat newly written coverage files as successful until `./run.ps1 PC` and then `./run.ps1 TC` confirm the result.

See finalized postmortem memory: `.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`
See issue record: `issues/repeated-coverage-remediation-failure-root-cause.md`

---

## TC Run Results (2026-03-16)

- **68 packages**, **1210 files**, **755 at 100%**, **455 below 100%**
- **21 packages at 100%**
- **0 blocked packages** (all compile)
- Detailed file-level plan: `.lovable/memory/workflow/03-coverage-file-level-plan.md`

---

## Completed Tasks

### 1. ✅ 100% Coverage Push — Batch 1 (11 packages)
**Completed: 2026-03-14** — Created 11 coverage test files for packages 75-97%.

### 2. ✅ 100% Coverage Push — Batch 2 (6 packages)
**Completed: 2026-03-14** — Created 6 coverage test files for packages 0-57%.

### 3. ✅ 100% Coverage Push — Batch 3 (7 packages)
**Completed: 2026-03-15** — Created 7 coverage test files for generic/utility packages.

### 4. ✅ Coverage Prompt Generator System
**Completed: 2026-03-15** — PowerShell-based system for auto-generating AI-friendly prompt files.

### 5. ✅ Compile Baseline Refresh
**Completed: 2026-03-16** — Ran `./run.ps1 PC`. Only 1 blocked package (corestrtests). Fixed Coverage10_test.go.

### 6. ✅ Audit 6 High-Risk Coverage Files
**Completed: 2026-03-16** — All 6 files audited. One assertion fix in converterstests.

### 7. ✅ TC Coverage Run & Plan Generation
**Completed: 2026-03-16** — Full coverage results analyzed. 3-tier plan created with ~975 test cases across 20 packages.

---

## Active Tasks (Ordered by Priority)

### 8. 🔲 Phase 1 — Quick Wins (6 packages, ~195 test cases)
Packages near 100% that need targeted branch coverage:
1. `coreonce` (95.7%) — 7 files, ~30 tests
2. `keymk` (95.6%) — 5 files, ~20 tests
3. `corerange` (94.3%) — 11 files, ~30 tests
4. `enumimpl` (95.9%) — 16 files, ~40 tests
5. `corevalidator` (91.2%) — 10 files, ~35 tests
6. `stringslice` (90.6%) — 24 files, ~50 tests

### 9. 🔲 Phase 2 — Moderate Effort (5 packages, ~215 test cases)
7. `errcore` (90.2%) — 15 files, ~20 tests
8. `reflectmodel` (72.6%) — 3 files, ~15 tests
9. `reflectinternal` (80.4%) — 11 files, ~40 tests
10. `corejson` (45.0%) — 18 files, ~60 tests ⚠️ HIGH RISK
11. `corepayload` (56.4%) — 23 files, ~80 tests ⚠️ HIGH RISK

### 10. 🔲 Phase 3 — Heavy Lift (4 packages, ~365+ test cases)
12. `codestack` (0.0%) — 11 files, ~25 tests
13. `corecmp` (10.8%) — 22 files, ~40 tests
14. `corestr` (3.3%) — 52 files, ~150+ tests ⚠️ VERY HIGH RISK
15. `coredynamic` (0.9%) — 57 files, ~150+ tests ⚠️ VERY HIGH RISK

### 11. 🔲 Test Title Audit — Remaining 17 Packages
Not blocked. Can proceed independently.

### 12. 🔲 Diagnostic Output Regression Tests
Create snapshot tests for diagnostic output formatting.

---

## Process Rules (From Postmortem)

1. **List first, then fix one-by-one.** Regenerate blocked packages before new work.
2. **Read source before every test edit.** Never infer signatures from naming patterns.
3. **Use a package gate.** Fix one package → compile verify → move on.
4. **Do not trust coverage percentages while blockers exist.**
5. **Do not report success from edits alone.** Only `./run.ps1 PC` / `TC` are evidence.
6. **Do not bulk-create coverage suites for unfamiliar packages.**
7. **Honor project behavior standards.** Vacuous truth, nil-handling, byte-slice clone.
8. **Honor naming standards.** `Test_Cov[N]_{Method}_{Context}` format.
