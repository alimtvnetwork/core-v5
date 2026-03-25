# Coverage Batch Testing — API Mismatch Risk

## Date: 2026-03-15

## Context

When creating 24 Coverage test files across 3 batches (Batches 1-3), all tests were written by reading source code and inferring API signatures. These files have NOT been compiled or run yet.

## Root Cause

Coverage test files are created based on source code reading, but:
1. Source APIs may have changed between the time of reading and test creation
2. Generic type constraints may not be correctly inferred from source
3. Method signatures with optional/variadic parameters may be misunderstood
4. Some methods may exist on pointer receivers vs value receivers

## Solution

1. Run `./run.ps1 PC` (Pre-Commit API Checker) to compile-check all Coverage files
2. Run `./run.ps1 TC` to execute full coverage and identify blocked packages
3. Fix any API mismatches found in the 24 new files
4. The `blocked-packages.json` will identify which files fail

## Iteration

This is a known risk pattern documented in `knowledge://memory/issues/coverage-test-api-mismatch-cascade`. Each batch of coverage tests should be validated with `./run.ps1 PC` before committing.

## Learning

- Always run pre-commit check after creating Coverage test files
- The Pre-Commit API Checker was specifically built to catch these issues early
- Batch creation of test files amplifies the risk — 24 files means 24 potential failure points

## What Not To Repeat

- Don't create large batches of coverage tests without running compile checks
- Don't assume source code reading is sufficient — always validate against the compiler
- Don't skip the `./run.ps1 PC` step before committing Coverage files
