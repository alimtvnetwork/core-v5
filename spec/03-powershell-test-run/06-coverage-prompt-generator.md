# Coverage Prompt Generator

## Overview

After `./run.ps1 TC` completes, the runner automatically generates AI-friendly prompt files listing all functions below 100% coverage with their specific uncovered line ranges.

## Scripts

```
scripts/coverage/
├── Generate-CoveragePrompts.ps1   # Main generator (called by run.ps1)
├── Get-UncoveredLines.ps1         # Standalone: uncovered lines for one file
└── Get-FunctionCoverage.ps1       # Standalone: filter functions by threshold
```

## Output

```
data/prompts/
├── coverage-prompt-1.txt          # Functions 1-500 (sorted by coverage ascending)
├── coverage-prompt-2.txt          # Functions 501-1000
├── ...
└── prompts-summary.json           # Metadata (counts, batch info)
```

## Prompt File Format

```text
# Coverage Improvement Prompt — Batch 1/3
# Generated: 2026-03-15 12:00:00
# Functions: 500 (of 1200 total below 100%)

Please improve the code coverage to 100% for these functions.
Each function lists its current coverage and the specific uncovered lines.
Write tests in tests/integratedtests/{pkg}tests/ using the AAA pattern with args.Map + ShouldBeEqual.

─────────────────────────────────────────────────────────────

## NewError
   File:     errcore/ErrorNew.go
   Package:  errcore
   Coverage: 66.7%
   Uncovered lines: L15-L17, L22

## SplitLeftRight
   File:     internal/strutilinternal/all-left-right-splits.go
   Package:  internal/strutilinternal
   Coverage: 40.0%
   Uncovered lines: L8-L12, L18
```

## Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `CoverProfile` | (required) | Path to merged coverage.out |
| `FuncOutput` | (required) | Lines from `go tool cover -func` |
| `OutputDir` | `data/prompts` | Where to write prompt files |
| `BatchSize` | 500 | Functions per file |
| `ProjectRoot` | (auto-detect) | Project root for path resolution |

## Standalone Usage

```powershell
# Get uncovered lines for a specific file
./scripts/coverage/Get-UncoveredLines.ps1 `
  -CoverProfile data/coverage/coverage.out `
  -SourceFile "github.com/alimtvnetwork/core/errcore/ErrorNew.go"

# Get all functions below a threshold
$funcLines = go tool cover -func=data/coverage/coverage.out
./scripts/coverage/Get-FunctionCoverage.ps1 -FuncOutput $funcLines -Threshold 80
```

## Integration Point

Called automatically at end of `Invoke-TestCoverage` in `run.ps1` (line ~1062):

```powershell
$promptScript = Join-Path $PSScriptRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
if (Test-Path $promptScript) {
    & $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput `
      -OutputDir $promptsDir -BatchSize 500 -ProjectRoot $PSScriptRoot
}
```

## Related Docs

- [PowerShell Test Runner Overview](01-overview.md)
- [Parallel Threading Strategy](05-parallel-threading.md)
