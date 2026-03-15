# Completed: Coverage Prompt Generator System

## Completed: 2026-03-15

### What Was Done

Built a PowerShell-based system to auto-generate AI-friendly prompt files after coverage runs.

### Architecture
```
scripts/coverage/
├── Generate-CoveragePrompts.ps1   # Main: parse coverage.out → batched prompt files
├── Get-UncoveredLines.ps1         # Utility: uncovered lines for a specific file
└── Get-FunctionCoverage.ps1       # Utility: filter functions by coverage threshold
```

### How It Works
1. Parses `go tool cover -func` output to find all functions <100% coverage
2. Parses `coverage.out` to extract specific uncovered line ranges (count=0)
3. Matches uncovered ranges to functions by file + line boundaries
4. Writes batched prompt files (500 functions/file) to `data/prompts/`
5. Each prompt includes: function name, file, package, coverage %, uncovered lines

### Integration
Called automatically at end of `./run.ps1 TC` via:
```powershell
$promptScript = Join-Path $PSScriptRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
& $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput ...
```

### Output Format
```
data/prompts/
├── coverage-prompt-1.txt    # Functions 1-500
├── coverage-prompt-2.txt    # Functions 501-1000
└── prompts-summary.json     # Metadata
```
