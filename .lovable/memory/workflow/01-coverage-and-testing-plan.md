# Coverage & Testing Master Plan

## Status: In Progress — Audit Complete, Awaiting TC Run
## Last Updated: 2026-03-16T14:30:00+08:00

## Critical Root Cause Checkpoint
Coverage work has been repeatedly invalidated by assumed APIs, broad unverified coverage-file generation, and skipping the compile-first gate. Do **not** treat newly written coverage files as successful until `./run.ps1 PC` and then `./run.ps1 TC` confirm the result.

See finalized postmortem memory: `.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`
See issue record: `issues/repeated-coverage-remediation-failure-root-cause.md`

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
**Completed: 2026-03-16** — Ran `./run.ps1 PC`. Only 1 blocked package (corestrtests). Fixed Coverage10_test.go: embedded struct literal mismatches, wrong ValueStatus fields, non-existent Hashmap.SortedKeys(). Issue recorded.

### 6. ✅ Audit 6 High-Risk Coverage Files
**Completed: 2026-03-16** — All 6 files audited against source. All APIs verified to exist with correct signatures. One assertion fix: `converterstests/Coverage4_test.go` had wrong expected len for IntegersSkipAndDefaultValue (4→3). Files audited:
- `errcoretests/Coverage9_test.go` ✅ — 755 lines, all APIs verified
- `simplewraptests/Coverage7_test.go` ✅ — 264 lines, all APIs verified
- `issettertests/Coverage7_test.go` ✅ — 371 lines, all APIs verified
- `isanytests/Coverage9_test.go` ✅ — 353 lines, all APIs verified
- `converterstests/Coverage4_test.go` ✅ — Fixed len assertion bug
- `stringslicetests/Coverage7_test.go` ✅ — 304 lines, all APIs verified

### 7. 🔲 Remaining 12 Packages to 100%
**Blocked by**: Tasks 5, 6
Packages: `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`, `corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`

### 8. 🔲 Test Title Audit — Remaining 17 Packages
Not blocked. Can proceed independently.

### 9. 🔲 Diagnostic Output Regression Tests
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
