# Unit Coverage Fix — Workflow Spec

## Status: 🔄 Active

## Instructions (Verbatim from User)

Fix blocked packages and improve code coverage to 100%. List packages that have less than 100% code coverage. Work on two packages at a time and list the remaining ones that need to be done next. If the .out file, and related JSON files not given then ask for it.

You do not need to instruct me to run -TC. I will handle that.

Every time I say "next", write tests for the remaining packages.

## Execution Model

1. Process **exactly two packages per iteration**
2. Maintain a rolling list of remaining packages
3. Continue until all packages reach 100% coverage
4. On "next" command: write tests for the next two packages

## Priority Order

1. **Blocked packages first** — fix compilation errors
2. **Failing tests** — fix assertion mismatches
3. **Coverage gaps** — write new tests for uncovered paths

## Test Writing Standards

- All tests in `tests/integratedtests/{pkg}tests/`
- AAA pattern (Arrange, Act, Assert)
- Naming: `Test_Cov{N}_{Method}_{Context}`
- Title format: `"{Function} returns {Result} -- {Input Context}"`
- Use `args.Map` and `CaseV1.ShouldBeEqual` patterns
- No flaky tests, deterministic outcomes
- Cover: normal paths, edge cases, error handling, boundary conditions

## Reporting Format (per iteration)

1. **Completed**: package name, previous coverage, target 100%
2. **Remaining**: all packages still below 100% with coverage %
3. **Notes**: blockers, fixes applied, ambiguities

## Rules

- Do NOT skip edge cases or error paths
- Do NOT modify production logic unless required to fix blockers
- Preserve existing working behavior
- Read source before writing tests — never infer APIs
- Internal packages excluded from 100% mandate

## Data Requirements

- `coverage-*.out` file (coverage profile)
- `per-package-coverage-*.json` (coverage summary)
- `blocked-packages-*.txt` (compilation failures)
- `failing-tests-*.txt` (test failures)

If not provided, ask for them.
