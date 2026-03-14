# Issue: Coverage Regression from Build-Failure Cascade + Non-Deterministic Parallel Feedback

## Summary
Coverage appeared to "get worse" after recent changes because multiple integrated test packages failed to compile, causing a broad measurement gap; parallel output ordering also made the failure pattern harder to read and trust.

## Root Cause

### 1) Build-Failure Cascade (Primary)
Coverage depends on integrated test packages compiling. When blocked packages increase, major source packages lose active coverage contribution, which manifests as very low/0% package totals.

### 2) API Signature Drift in Coverage Tests
Several coverage test files used mismatched calls versus production APIs (field-vs-method access, changed signatures, value-vs-pointer expectations, callback signatures), which created compile blockers.

### 3) Parallel Output Perception Gap
Even when execution is correct, asynchronous completion order can produce unstable display order if not normalized, creating the impression that reported numbers are random or degraded.

## Evidence
- `coverage-summary-13.txt` shows broad low/0% package coverage after blocked test-package set increased.
- `blocked-packages-12.txt` shows 12 compile-blocked integrated test packages and concrete signature mismatches.

## Solution

### A) Deterministic Parallel-Sync Reporting
Implement explicit plan-array + async execution + ordered sync rendering:
- pre-sort package list
- attach stable index to each package
- collect async results
- display/write reports in index order

### B) Compile-Blocker Cleanup
Fix test code to align with verified production signatures (no assumed APIs), focusing on blocked packages first.

### C) Parallel Artifact Safety
Use package-derived sanitized filenames for compile/test artifacts to avoid collisions.

## Iteration Details
- **Iteration 1:** Added pre-coverage compile-check and blocked-package reporting.
- **Iteration 2:** Added deterministic sorting for compile and coverage result display.
- **Iteration 3 (current):** Added formal parallel-sync spec and continued API-alignment fixes in remaining blocked test packages.

## Learnings
1. Coverage quality metrics are only meaningful when compile blockers are near-zero.
2. Assumed test signatures are a high-cost failure mode.
3. Parallel execution must be followed by deterministic sync to preserve operator trust.

## Prevention (What Not To Repeat)
1. **Never write coverage tests against assumed signatures**; verify from source first.
2. **Never stream worker output directly** in parallel mode; always sync-sort before display.
3. **Never use collision-prone temp names** (`compile-1`, `cover-1`) for parallel jobs.
4. **Always treat blocked-packages count as a release-quality gate** before interpreting coverage percentages.

## Permanent Guardrails
- Enforce plan-indexed parallel-sync pattern for all multi-package runners.
- Keep blocked-packages report deterministic and machine-parseable.
- Require source-signature verification before committing new Coverage*.go files.