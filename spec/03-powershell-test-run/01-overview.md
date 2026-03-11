# PowerShell Test Runner (`run.ps1`)

## Overview

`run.ps1` is the primary task runner for the project. It provides short, memorable commands for running tests, building, formatting, and more.

## Quick Reference

```powershell
./run.ps1 -t              # Run all tests
./run.ps1 -tc             # Run tests with coverage (HTML + summary)
./run.ps1 -h              # Show help
```

## All Commands

| Short | Flag | Long | Description |
|-------|------|------|-------------|
| `T` | `-t` | `test` | Run all tests (verbose, with log output) |
| `TP` | `-tp` | `test-pkg` | Run tests for a specific package |
| `TC` | `-tc` | `test-cover` | Run tests with coverage report |
| `TI` | `-ti` | `test-int` | Run integrated tests only |
| `TF` | `-tf` | `test-fail` | Show last failing tests log |
| `GC` | `-gc` | `goconvey` | Launch GoConvey browser test runner |
| `R` | `-r` | `run` | Run the main application |
| `B` | `-b` | `build` | Build the binary |
| `BR` | `-br` | `build-run` | Build then run |
| `F` | `-f` | `fmt` | Format all Go files |
| `L` | `-l` | `lint` | Run `go vet` |
| `V` | `-v` | `vet` | Run `go vet` |
| `TY` | `-ty` | `tidy` | Run `go mod tidy` |
| `C` | `-c` | `clean` | Clean build artifacts + coverage |
| `H` | `-h` | `help` | Show help |

## Usage Examples

```powershell
# Run all tests
./run.ps1 T
./run.ps1 -t
./run.ps1 test

# Run a specific package
./run.ps1 TP regexnewtests
./run.ps1 -tp corestrtests

# Run tests with coverage (auto-opens HTML report)
./run.ps1 TC
./run.ps1 -tc
./run.ps1 -tc --no-open    # skip auto-open

# Show last failing tests
./run.ps1 TF

# Launch GoConvey on custom port
./run.ps1 GC 9090

# Show help
./run.ps1 -h
./run.ps1 help
./run.ps1              # defaults to help
```

## Test Execution Pipeline

When you run `./run.ps1 -t`, the following happens in order:

```
1. git pull               (Invoke-FetchLatest)
2. go mod tidy            (dependency sync)
3. go build ./...         (Invoke-BuildCheck — fails fast if compilation errors)
4. go test -v -count=1    (run all tests, no caching)
5. Write-TestLogs         (parse output → passing/failing logs)
6. Open-FailingTestsIfAny (auto-open failing-tests.txt if failures exist)
```

### Build Check Gate

Before running tests, the script compiles the test packages. If compilation fails:
- Tests are **skipped entirely**
- Build errors are written to `data/test-logs/failing-tests.txt`
- The failing log is auto-opened

This prevents confusing test output when the code doesn't compile.

## Test Output & Logs

All test runs produce structured log files in `data/test-logs/`:

| File | Content |
|------|---------|
| `passing-tests.txt` | List of passing test names with count and timestamp |
| `failing-tests.txt` | Summary of failed tests + full diagnostic details |
| `raw-output.txt` | Complete unprocessed `go test` output |

### Failing Tests Log Format

```
# Failing Tests — 2026-03-11 10:30:00
# Count: 3

# ── Summary ──
  - TestFoo/Case_1
  - TestBar/Case_3
  - TestBaz/Case_0

# ── Details ──
FAIL: TestFoo/Case_1
  expected: "hello"
  actual:   "world"

FAIL: TestBar/Case_3
  ...
```

The summary section lists all failed test names sorted alphabetically, followed by detailed diagnostic output for each failure.

## Coverage Reports (`-tc`)

Running `./run.ps1 -tc` produces:

| File | Description |
|------|-------------|
| `data/coverage/coverage.out` | Raw Go coverage profile |
| `data/coverage/coverage.html` | Visual HTML report (auto-opens in browser) |
| `data/coverage/coverage-summary.txt` | Text summary with per-package and low-coverage highlights |

### Coverage Summary Contents

1. **Total Coverage** — aggregate percentage
2. **Per-Package Coverage** — breakdown by test package
3. **Low Coverage Functions (< 50%)** — functions needing attention
4. **Report file paths**

Pass `--no-open` to skip auto-opening the HTML report:
```powershell
./run.ps1 -tc --no-open
```

## Command Dispatch

All three forms are equivalent and case-insensitive:

```powershell
./run.ps1 T       # uppercase shorthand
./run.ps1 -t      # hyphen-lowercase flag
./run.ps1 test    # long name
```

The dispatch table in `run.ps1` normalizes all forms via `$Command.ToLower()` and matches against a set of aliases.

## Cleanup

```powershell
./run.ps1 -c      # removes: build/, tests/coverage.out, data/coverage/
```

## Related Docs

- [Repo Overview](/spec/01-app/00-repo-overview.md)
- [CMD Entrypoints](/spec/01-app/12-cmd-entrypoints.md)
- [Testing Patterns](/spec/01-app/13-testing-patterns.md)
