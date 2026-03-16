# Issue: Repeated Coverage Remediation Failure from Assumed APIs + Missing Compile Gate

## Summary
Coverage work kept failing because coverage test files were written against partially-read or assumed APIs, then expanded across multiple packages before a compile-first verification loop was completed. The failure was primarily **agent process failure**: hallucinated signatures, broad bulk edits, and reporting progress before `./run.ps1 PC` / `./run.ps1 TC` confirmed reality.

## Root Cause

### 1) Assumed API Shapes Instead of Verified Source
The agent repeatedly inferred constructors, method names, parameters, return types, and enum/value behavior from naming patterns or earlier packages instead of fully verifying each target package from source.

Examples of repeated failure patterns:
- treating fields as methods
- inventing helper methods that do not exist
- using wrong parameter counts
- asserting pointer behavior for value types
- assuming common behavior across unrelated packages

### 2) Violation of the Required 3-Step Methodology
The project requires:
1. identify the root cause
2. define the solution
3. apply the solution

The repeated failure happened because the agent often skipped from “coverage is low” directly to “write more tests” without first proving the exact uncovered branch behavior and current package API.

### 3) Bulk Multi-Package Test Generation Without a Package Gate
Coverage files were created across several packages in one pass, which amplified errors:
- one mistaken assumption created multiple compile failures
- blocked integrated packages caused a build-failure cascade
- coverage reports then looked worse or misleadingly low

### 4) Missing Compile-First Feedback Loop
The agent did not consistently enforce this sequence:
1. write or repair **one package**
2. run `./run.ps1 PC`
3. fix compile issues for that package
4. move to the next package
5. run `./run.ps1 TC` only after blockers are under control

Without that gate, invalid test files accumulated faster than they were verified.

### 5) Reporting Progress Before Objective Verification
Some status reporting described coverage work as if it were complete when only files had been written. In this repo, written files are **not** evidence; only `PC` and `TC` results are evidence.

## Solution

### A) Freeze New Coverage Expansion Until the Baseline Is Verified
Do not create more `Coverage*.go` files until the current blocked-package state is re-measured.

### B) Regenerate the Blocked Package List First
Use the runner flow to regenerate the compile-truth baseline, then fix packages one-by-one in blocked order.

### C) Read Full Relevant Source Before Editing Tests
Before writing or fixing a coverage file, verify:
- constructor signatures
- exported methods and fields
- enum/value semantics
- pointer vs value returns
- callback signatures
- nil/empty branch behavior

### D) Enforce Package-at-a-Time Remediation
For each blocked package:
1. inspect failing test file
2. inspect production source
3. inspect behavior standards / existing memory
4. patch only that package
5. rerun compile verification
6. only then continue

### E) Treat Coverage as a Post-Compile Metric
Blocked packages must be reported before coverage percentages. Coverage numbers are not trustworthy while compile blockers remain.

## Iteration Details
- Prior improvements already completed in this repository include deterministic parallel-sync reporting, blocked-package reporting, diagnostic formatting fixes, large title-audit work, and earlier API-alignment fixes.
- Despite those improvements, a later coverage expansion pass repeated the same failure mode by writing additional coverage suites from incomplete API understanding.
- This means the main problem is not the existence of tooling; it is failure to follow the tooling/process strictly.

## Learnings
1. The fastest path is not bulk generation; it is verified package-by-package remediation.
2. Coverage files are high-risk because even small API mistakes block whole integrated packages.
3. “Looks right” is not enough in this repository; signatures and behavior must be proven from source.
4. `PC` is the gate; `TC` is the outcome. Skipping the gate causes false progress.
5. Memory must explicitly record agent failure modes, not just code defects.

## What Not To Repeat
1. **Never write coverage tests from naming intuition or cross-package pattern matching.**
2. **Never bulk-create many coverage files and assume they compile.**
3. **Never claim coverage improvement before `./run.ps1 TC` verifies it.**
4. **Never interpret low coverage without first checking blocked packages.**
5. **Never skip full source verification for unfamiliar packages like `errcore`, `issetter`, `corejson`, `corepayload`, `coredynamic`, or converter/helper packages.**
6. **Never continue adding tests while the compile baseline is unknown.**

## Strict Guardrails For The Next Agent
- Start every coverage session by reading the latest blocked-package report or generating it.
- Fix one package at a time.
- Read the source before every test edit.
- Use project behavioral memories (vacuous truth, nil handling, byte slice behavior, etc.) when asserting behavior.
- Keep test names in `Test_Cov[N]_{Method}_{Context}` format.
- Use `args.Map` assertions where appropriate, but do not let assertion style distract from API verification.
- Report **BLOCKED PACKAGES** before listing sub-100% coverage packages.
- If a package API is unclear, stop and read more source instead of guessing.
