# Coverage Prompts from PowerShell Coverage Run

## Date: 2026-03-15

## Goal
Generate AI-ready prompt files at the end of `./run.ps1 TC` so every function below 100% coverage is listed with uncovered lines.

## Requirements

1. **Trigger point**: Execute after coverage profile + `go tool cover -func` output are available in `Invoke-TestCoverage`.
2. **Input parameters** (script):
   - `CoverProfile` (required): merged `coverage.out`
   - `FuncOutput` (required): lines from `go tool cover -func`
   - `OutputDir` (optional): default `data/prompts`
   - `BatchSize` (optional): default `500`
   - `ProjectRoot` (optional): for module-relative paths
3. **Selection rule**: Include **only** functions with coverage `< 100%`.
4. **Uncovered lines**: Parse `coverage.out` blocks with `count=0`, map ranges to each function by file + line boundaries.
5. **Batching**: Write one prompt file per 500 functions:
   - `coverage-prompt-1.txt`, `coverage-prompt-2.txt`, ...
6. **Prompt header**: Each file must start with:
   - `Please make the code coverage to a hundred percent for these functions.`
7. **Entry format** (per function):
   - Method/function name
   - File path
   - Package path
   - Coverage percentage
   - Uncovered line ranges (`L10-L12, L18`, etc.)
8. **Metadata output**: Write `prompts-summary.json` with total functions, batch size, batch count, and generated files.

## Acceptance Criteria

- Running `./run.ps1 TC` creates `data/prompts/` if missing.
- Prompt files are generated in 500-function chunks.
- All listed functions have `< 100%` coverage.
- Every listed function includes uncovered lines (or explicit fallback when none found).
- Header sentence appears at top of every prompt file exactly as required.
