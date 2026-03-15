# Suggestions Tracker

## Last Updated: 2026-03-15

---

## Active Suggestions

### 1. Finish Remaining Package Coverage (12 packages)
**Priority:** High
**Status:** In Progress — Coverage test files created, 3 have API mismatches needing fixes
Packages: `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`,
`corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`
Files created but need `./run.ps1 PC` to verify compilation.
API mismatches flagged in: `enumimpltests/Coverage7_test.go`, `corejsontests/Coverage4_test.go`, `corestrtests/Coverage8_test.go`.

### 2. Verify All New Coverage Tests Compile
**Priority:** High
**Status:** Pending
Run `./run.ps1 PC` to confirm all newly created Coverage test files (Batches 1-4, 36 files total) compile without errors.

### 3. Test Title Audit (Remaining Packages)
**Priority:** Medium
**Status:** Pending — Scoped
1400+ titles across 40+ packages need renaming to `"{Function} returns {Result} -- {Input Context}"`.
Top violators by count: coregenerictests (347), coredynamictests (174), coreutilstests (87), chmodhelpertests (79), coreinstructiontests (47), coremathtests (44), coresorttests (41), issettertests (40).
Recommend incremental batches of 5-6 packages per session.

### 4. Coverage Prompt Generator Validation
**Priority:** Medium
**Status:** Pending
Run `./run.ps1 TC` and inspect `data/prompts/` to verify prompt files are generated correctly.

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
