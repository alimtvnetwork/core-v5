# Suggestions Tracker

## Last Updated: 2026-03-16T04:58:00+08:00

---

## Active Suggestions

### S-001: Run Compile Baseline Before More Coverage Work
- **suggestionId**: S-001
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (all integrated test packages)
- **description**: Run `./run.ps1 PC` to regenerate the real blocked-package baseline before any further coverage expansion.
- **rationale**: Coverage work has repeatedly failed because tests were written against assumed APIs. The postmortem (`.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`) mandates this as step 1.
- **proposed change**: User runs `./run.ps1 PC`, shares output. AI uses output to identify which packages are truly blocked.
- **acceptance criteria**: `./run.ps1 PC` runs cleanly; blocked package list is captured and documented.
- **status**: open
- **completion notes**: —

### S-002: Audit High-Risk Coverage Files One-By-One
- **suggestionId**: S-002
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (test packages)
- **description**: Compile-check these files individually instead of in bulk: `errcoretests/Coverage9_test.go`, `simplewraptests/Coverage7_test.go`, `issettertests/Coverage7_test.go`, `isanytests/Coverage9_test.go`, `converterstests/Coverage4_test.go`, `stringslicetests/Coverage7_test.go`
- **rationale**: These files were generated in bulk and never compile-verified. Each may have API mismatches.
- **proposed change**: Read source for each package → verify test file signatures → fix mismatches → compile one at a time.
- **acceptance criteria**: Each file compiles individually via `go build ./tests/integratedtests/<pkg>/...`
- **status**: open
- **completion notes**: —

### S-003: Complete Remaining Package Coverage (12 packages)
- **suggestionId**: S-003
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Push remaining packages to 100% coverage: `keymk`, `corerange`, `coreonce`, `enumimpl`, `stringslice`, `corevalidator`, `corepayload`, `reflectinternal`, `corejson`, `corestr`, `coredynamic`, `reflectmodel`
- **rationale**: These packages are below 100% and have no verified coverage files yet.
- **proposed change**: One package at a time: read source → write tests → compile verify → run coverage.
- **acceptance criteria**: Each package reaches 100% as reported by `./run.ps1 TC`.
- **status**: open (blocked by S-001 and S-002)
- **completion notes**: —

### S-004: Test Title Audit (Remaining 17 Packages)
- **suggestionId**: S-004
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core (test packages)
- **description**: 1400+ test titles across 40+ packages need renaming to `"{Function} returns {Result} -- {Input Context}"`. Top violators: coregenerictests (347), coredynamictests (174), coreutilstests (87).
- **rationale**: Consistent naming aids debugging and makes test output scannable.
- **proposed change**: Batch rename test titles package-by-package.
- **acceptance criteria**: All test titles follow the convention. `grep -c` confirms zero violations.
- **status**: open
- **completion notes**: —

### S-005: Codegen Removal (Track B)
- **suggestionId**: S-005
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Complete codegen removal per `spec/01-app/10-codegen-deprecation-plan.md` Track B.
- **rationale**: Codegen is deprecated, adds maintenance burden, and has no confirmed external consumers.
- **proposed change**: Remove `cmd/main/unitTestGenerator.go`, `tests/integratedtests/codegentests/`, `codegen/` tree. Run `go mod tidy`. Update docs.
- **acceptance criteria**: All exit criteria in `10-codegen-deprecation-plan.md` are met. Module compiles without codegen.
- **status**: open
- **completion notes**: —

### S-006: Value Receiver Migration (Phase 6 Completion)
- **suggestionId**: S-006
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Continue migrating read-only methods from pointer to value receivers. `issetter/` and `coreversion/` are done. Remaining: `coretaskinfo/`, `corepayload/`, others.
- **rationale**: Value receivers communicate immutability, enable compiler optimizations, and clarify API intent.
- **proposed change**: Package-by-package, verify interface satisfaction after each change.
- **acceptance criteria**: All read-only methods use value receivers. All interfaces still satisfied. All tests pass.
- **status**: open
- **completion notes**: —

### S-007: Remaining Package READMEs
- **suggestionId**: S-007
- **createdAt**: 2026-03-16
- **source**: Lovable
- **affectedProject**: core
- **description**: Create README.md for: `coregeneric`, `corestr`, `coreonce`, `corerange`, `stringslice`.
- **rationale**: These are core packages without documentation.
- **proposed change**: Read source, document API with examples, verify method names against source.
- **acceptance criteria**: Each README has folder tree, usage examples, and verified method signatures.
- **status**: open
- **completion notes**: —

---

## Completed Suggestions (Archive)

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
| 12 | Coverage Push Batch 4 (6 packages) | 2026-03-16 (pending PC verification) |

> Completed suggestion detail files are in `completed/` subfolder.
> Batch 4 session log: `.lovable/memory/workflow/02-coverage-batch4-session-log.md`
