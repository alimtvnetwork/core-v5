# Coverage Prompt Generator — Integration with run.ps1

## Date: 2026-03-15

## Context

After each coverage run (`./run.ps1 TC`), the runner already produced coverage summaries, JSON exports, and HTML reports. However, the output for low-coverage functions only listed function names and percentages — it did not include the **specific uncovered lines**, making it hard for AI agents to write targeted tests.

## Root Cause

The `go tool cover -func` output only gives per-function coverage percentages, not line-level detail. The line-level detail exists in `coverage.out` but was not being cross-referenced with function boundaries to produce actionable prompts.

## Solution

Created a modular PowerShell script system under `scripts/coverage/`:

1. **`Generate-CoveragePrompts.ps1`** — Parses both `go tool cover -func` output AND `coverage.out` to:
   - Identify all functions below 100% coverage
   - Extract their specific uncovered line ranges (where count=0 in coverage.out)
   - Match line ranges to functions using start-line boundaries
   - Batch functions into groups of 500
   - Write AI-friendly prompt files to `data/prompts/`

2. **`Get-UncoveredLines.ps1`** — Standalone utility for debugging single files
3. **`Get-FunctionCoverage.ps1`** — Standalone utility for threshold filtering

Integrated into `run.ps1` at the end of `Invoke-TestCoverage`, after the coverage summary is written.

## Iteration

- First implementation: monolithic approach considered
- Final implementation: 3 small focused scripts in `scripts/coverage/` as requested by user
- Integration point: single call at end of `Invoke-TestCoverage` in `run.ps1`

## Learning

- Coverage.out provides line-level granularity but requires parsing the `file.go:startLine.col,endLine.col stmts count` format
- Function boundaries from `go tool cover -func` can be approximated by sorting functions per file by start line
- Batching at 500 functions keeps prompt files at a manageable size for AI context windows

## What Not To Repeat

- Don't put all logic in `run.ps1` — keep utility scripts modular in dedicated subdirectories
- Don't rely solely on `go tool cover -func` for AI prompts — always include specific uncovered lines
- Don't generate unbounded prompt files — always batch to prevent context window overflow
