# Suggestions Tracker

## Last Updated: 2026-03-16

---

## Active Suggestions

### 1. Regenerate Blocked Package Truth Before More Coverage Work
**Priority:** Critical
**Status:** In Progress
Run `./run.ps1 PC` first. The current root cause is process failure: assumed APIs + unverified bulk coverage generation. No additional coverage progress should be trusted until the blocked-package baseline is refreshed.

### 2. Audit Latest High-Risk Coverage Files Package-By-Package
**Priority:** High
**Status:** Pending
Audit and compile-check these files one-by-one instead of in bulk:
- `errcoretests/Coverage9_test.go`
- `simplewraptests/Coverage7_test.go`
- `issettertests/Coverage7_test.go`
- `isanytests/Coverage9_test.go`
- `converterstests/Coverage4_test.go`
- `stringslicetests/Coverage7_test.go`

### 3. Finish Remaining Package Coverage Only After Compile Baseline Is Stable
**Priority:** High
**Status:** Blocked by compile verification
Packages: `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`,
`corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`
Previously documented API-mismatch fixes exist for: `enumimpltests/Coverage7_test.go`, `corejsontests/Coverage4_test.go`, `corestrtests/Coverage8_test.go`.

### 4. Test Title Audit (Remaining Packages)
**Priority:** Medium
**Status:** Pending — Scoped
1400+ titles across 40+ packages need renaming to `"{Function} returns {Result} -- {Input Context}"`.
Top violators by count: coregenerictests (347), coredynamictests (174), coreutilstests (87), chmodhelpertests (79), coreinstructiontests (47), coremathtests (44), coresorttests (41), issettertests (40).

### 5. Coverage Prompt Generator Validation
**Priority:** Medium
**Status:** Pending
Run `./run.ps1 TC` and inspect `data/prompts/` to verify prompt files are generated correctly after blockers are cleared.

---

## Completed Suggestions (Moved to `completed/`)

| # | Title | Completed |
|---|-------|-----------|
| 1 | Diagnostic Formatting Improvements | 2026-03-11 |
| 2 | Test Title Audit (Batches 1-4) | 2026-03-11 |
| 3 | Fix 21 Failing Tests | 2026-03-11 |
| 4 | Coverage Push Batch 1 (11 packages) | 2026-03-14 |
| 5 | Coverage Push Batch 2 (6 packages) | 2026-03-14 |
| 6 | Coverage Push Batch 3 (7 packages) | 2026-03-15 |
| 7 | Coverage Prompt Generator System | 2026-03-15 |
| 8 | Deep Clone Production Bug Fix | 2026-03-15 |
| 9 | Nil Receiver Coverage Audit | 2026-03-15 |
| 10 | Test Runner Hardening Review | 2026-03-15 |
| 11 | Diagnostic Output Regression Tests | 2026-03-15 |
