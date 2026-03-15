# Suggestions Tracker

## Last Updated: 2026-03-15

---

## Active Suggestions

### 1. Finish Remaining Package Coverage (12 packages)
**Priority:** High
**Status:** Pending
Packages still below 100% that need Coverage test files:
- `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`
- `corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`
See workflow plan task #5 for full details.

### 2. Verify All 24 New Coverage Tests Compile
**Priority:** High
**Status:** Pending
Run `./run.ps1 TC` to confirm all newly created Coverage test files from Batches 1-3 compile without errors and produce coverage improvements. Fix any API mismatches found by the Pre-Commit checker.

### 3. Continue Test Title Audit (Remaining Packages)
**Priority:** Medium
**Status:** Pending
17 packages remain unaudited for `"{Function} returns {Result} -- {Input Context}"` convention. See workflow plan task #7.

### 4. Nil Receiver Coverage Audit
**Priority:** Medium
**Status:** Pending
Systematically migrate nil receiver test cases from CaseV1 to CaseNilSafe pattern across all packages.

### 5. Deep Clone Bug Investigation
**Priority:** High
**Status:** Pending
`corepayload.Attributes.Clone(deep=true)` returns error. Investigate root cause in production code.

### 6. Test Runner Hardening
**Priority:** Low
**Status:** Pending
Review all test runners for unconditional map key insertion, value vs pointer type assertions, incorrect independence/equality check logic.

### 7. Diagnostic Output Regression Tests
**Priority:** Low
**Status:** Pending
Create snapshot tests for diagnostic output formatting to prevent regressions.

### 8. Coverage Prompt Generator Validation
**Priority:** Medium
**Status:** Pending — New
Validate the new `scripts/coverage/Generate-CoveragePrompts.ps1` system produces correct prompt files by running `./run.ps1 TC` and inspecting `data/prompts/`. Verify function-to-line-range matching accuracy.

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
