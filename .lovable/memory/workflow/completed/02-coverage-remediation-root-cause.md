# Coverage Remediation Root Cause Memory

## Status: ✅ Finalized Postmortem (2026-03-16)

## Why the coverage work kept failing
The repeated failures were caused by **agent-side process mistakes**, not just isolated package bugs.

### Root cause
1. Coverage tests were sometimes written against **assumed** APIs instead of fully verified source.
2. Multiple packages were edited in bulk before a compile-first verification loop was completed.
3. The required sequence **identify root cause → define solution → apply solution** was not followed strictly enough.
4. Progress was sometimes described from file creation rather than from verified `PC` / `TC` results.
5. Because integrated test packages drive source coverage, compile blockers created a **build-failure cascade** that made coverage appear worse or unstable.

## Strict guidelines for the next agent
1. **List first, then fix one-by-one.** Regenerate blocked packages before new coverage work.
2. **Read source before every test edit.** Do not infer signatures from naming patterns.
3. **Use a package gate.** Fix one package, run compile verification, then move on.
4. **Do not trust coverage percentages while blockers exist.** Blocked packages must be reported first.
5. **Do not report success from edits alone.** Only `./run.ps1 PC` and `./run.ps1 TC` are evidence.
6. **Do not bulk-create coverage suites for unfamiliar packages.** `errcore`, `issetter`, `corejson`, `corepayload`, `coredynamic`, and converters/helpers are especially high-risk.
7. **Honor project behavior standards.** Examples already recorded in memory include vacuous truth (`All*` on empty is true, `Any*` on empty is false), nil-handling conventions, and byte-slice clone behavior.
8. **Honor naming/reporting standards.** Coverage tests use `Test_Cov[N]_{Method}_{Context}` and every user-facing coverage response must show **BLOCKED PACKAGES** before the sub-100% package list.

## Things already done before this postmortem
- Added deterministic parallel-sync reporting and blocked-package reporting.
- Fixed major diagnostic-formatting regressions.
- Completed several large coverage push batches.
- Built the coverage prompt generator system.
- Documented/fixed prior API mismatch cases such as `enumimpltests/Coverage7_test.go`, `corejsontests/Coverage4_test.go`, and `corestrtests/Coverage8_test.go`.
- Completed earlier failing-test cleanup and title-audit batches.

## What should NOT be assumed after this postmortem
- Newly written coverage files are **not** automatically valid.
- A lower coverage report does **not** always mean the source lost coverage; it may mean integrated test packages are blocked.
- “Similar package = similar API” is false in this repository.

## Current plan items
1. Run `./run.ps1 PC` and regenerate the current blocked-package truth.
2. Audit the latest high-risk coverage files one package at a time.
3. Fix compile mismatches package-by-package.
4. Run `./run.ps1 TC` only after blockers are reduced/cleared.
5. Update active workflow/suggestions with the real blocked list and real coverage numbers.
6. Resume the 100% coverage push only after the baseline is stable again.

## Memory references
- `issues/coverage-test-api-mismatch-cascade.md`
- `issues/coverage-regression-parallel-sync-build-cascade.md`
- `issues/repeated-coverage-remediation-failure-root-cause.md`
