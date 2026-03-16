# Suggestions Tracker

## Last Updated: 2026-03-16T09:50:00+08:00

## Convention

- **Location**: `.lovable/memory/suggestions/` — this file for active tracking, `completed/` for archives.
- **File naming**: Single tracker file (`01-suggestions-tracker.md`). Individual completed suggestions in `completed/NN-slug.md`.
- **Statuses**: `open` → `inProgress` → `done`
- **Completion handling**: When done, update status here and optionally create detail file in `completed/`.

---

## Active Suggestions

### S-001: Run Compile Baseline Before More Coverage Work
- **suggestionId**: S-001
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (all integrated test packages)
- **description**: Run `./run.ps1 PC` to regenerate the real blocked-package baseline before any further coverage expansion.
- **rationale**: Coverage work has repeatedly failed because tests were written against assumed APIs. Postmortem mandates this as step 1.
- **proposed change**: User runs `./run.ps1 PC`, shares output. AI uses output to identify which packages are truly blocked.
- **acceptance criteria**: `./run.ps1 PC` runs cleanly; blocked package list documented.
- **status**: open
- **completion notes**: —

### S-002: Verify Batch 4 Coverage Files (6 files)
- **suggestionId**: S-002
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (test packages)
- **description**: Compile-verify Batch 4 files: `coreindexestests/Coverage2_test.go`, `coremathtests/Coverage3_test.go`, `corecsvtests/Coverage3_test.go`, `intuniquetests/Coverage_test.go`, `stringutiltests/Coverage5_test.go`, `conditionaltests/Coverage8_test.go`
- **rationale**: Written 2026-03-16 but never compiled. Session log: `.lovable/memory/workflow/02-coverage-batch4-session-log.md`
- **proposed change**: User runs `./run.ps1 PC`, AI fixes any API mismatches.
- **acceptance criteria**: All 6 files compile via `./run.ps1 PC`.
- **status**: open (depends on S-001)
- **completion notes**: —

### S-003: Coverage Push — Phase 1 Quick Wins (6 packages)
- **suggestionId**: S-003
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Push 6 near-100% packages to full coverage: `coreonce` (95.7%), `keymk` (95.6%), `corerange` (94.3%), `enumimpl` (95.9%), `corevalidator` (91.2%), `stringslice` (90.6%).
- **rationale**: Highest ROI — small test additions for significant coverage gains.
- **proposed change**: One package at a time: read source → write tests → compile verify → run coverage.
- **acceptance criteria**: Each package reaches 100% via `./run.ps1 TC`.
- **status**: open (depends on S-001)
- **completion notes**: —

### S-004: Coverage Push — Phase 2 Moderate Effort (5 packages)
- **suggestionId**: S-004
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: `errcore` (90.2%), `reflectmodel` (72.6%), `reflectinternal` (80.4%), `corejson` (45%) ⚠️, `corepayload` (56%) ⚠️
- **rationale**: Medium effort. `corejson` and `corepayload` are HIGH RISK — require method signature inventory first.
- **proposed change**: One package at a time with mandatory source reading.
- **acceptance criteria**: Each reaches 100% via `./run.ps1 TC`.
- **status**: open (depends on S-003)
- **completion notes**: —

### S-005: Coverage Push — Phase 3 Heavy Lift (4 packages)
- **suggestionId**: S-005
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: `codestack` (0%), `corecmp` (10.8%), `corestr` (3.3%) ⚠️ VERY HIGH RISK, `coredynamic` (0.9%) ⚠️ VERY HIGH RISK
- **rationale**: Largest uncovered packages. 5-8 sessions each for `corestr` and `coredynamic`.
- **proposed change**: Method signature inventory → one file at a time → compile gate → coverage verify.
- **acceptance criteria**: Each reaches 100%.
- **status**: open (depends on S-004)
- **completion notes**: —

### S-006: Codegen Removal (Track B)
- **suggestionId**: S-006
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Complete codegen removal per `spec/01-app/10-codegen-deprecation-plan.md`.
- **rationale**: Deprecated package adds maintenance burden.
- **proposed change**: External audit → remove files → `go mod tidy` → update docs.
- **acceptance criteria**: All exit criteria in `10-codegen-deprecation-plan.md` met.
- **status**: open (prerequisite: user runs external audit `grep`)
- **completion notes**: —

### S-007: Spec Reconciliation
- **suggestionId**: S-007
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (spec files)
- **description**: Remove stale/contradictory entries from spec files. `15-code-review-report.md` still shows completed items as recommendations.
- **rationale**: Stale specs cause AI to re-implement completed work.
- **proposed change**: Audit each spec file, mark completed items, remove outdated recommendations.
- **acceptance criteria**: No spec file references completed work as pending.
- **status**: open
- **completion notes**: —

### S-008: CI Pipeline Setup
- **suggestionId**: S-008
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Add `golangci-lint`, test coverage reporting, and security scanning to CI.
- **rationale**: Currently no automated quality gates. Manual verification is error-prone.
- **proposed change**: Create CI config (GitHub Actions or GitLab CI) with lint, test, coverage steps.
- **acceptance criteria**: CI runs on push/PR, blocks on lint errors or test failures.
- **status**: open (low priority — Phase D)
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

> Detail files in `completed/` subfolder.
