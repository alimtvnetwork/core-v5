#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Project runner script with shorthands for common operations.

.DESCRIPTION
    Usage: ./run.ps1 <command>

    Commands (uppercase shorthands OR hyphen-lowercase):
        T   | -t   | test          Run all tests (verbose)
        TP  | -tp  | test-pkg      Run tests for a specific package: ./run.ps1 TP regexnewtests
        TC  | -tc  | test-cover    Run tests with coverage
        TCP | -tcp | test-cover-pkg Run coverage for a specific package: ./run.ps1 TCP regexnewtests
        TI  | -ti  | test-int      Run integrated tests only
        TF  | -tf  | test-fail     Show last failing tests log
        GC  | -gc  | goconvey      Launch GoConvey (browser test runner)
        R   | -r   | run           Run the main application
        B   | -b   | build         Build the binary
        BR  | -br  | build-run     Build then run
        F   | -f   | fmt           Format all Go files
        L   | -l   | lint          Run go vet on all packages
        V   | -v   | vet           Run go vet
        TY  | -ty  | tidy          Run go mod tidy
        C   | -c   | clean         Clean build artifacts
        H   | -h   | help          Show this help

.EXAMPLE
    ./run.ps1 T
    ./run.ps1 -t
    ./run.ps1 TP regexnewtests
    ./run.ps1 -tp regexnewtests
    ./run.ps1 -gc
#>

param(
    [Parameter(Position = 0)]
    [string]$Command = "help",

    [Parameter(Position = 1, ValueFromRemainingArguments)]
    [string[]]$ExtraArgs
)

# Normalize: if $Command was swallowed by PowerShell as a switch
# (e.g. -gc parsed away), $Command will be "help" — detect via $PSBoundParameters.
if (-not $PSBoundParameters.ContainsKey('Command')) {
    # Check $MyInvocation.Line for the actual argument
    $rawLine = $MyInvocation.Line
    $match = [regex]::Match($rawLine, '(?i)run\.ps1\s+(-?\w[\w-]*)')
    if ($match.Success) {
        $Command = $match.Groups[1].Value
    }
}

$ErrorActionPreference = "Stop"

# -- Colors --
function Write-Header([string]$msg) {
    Write-Host "`n=== $msg ===" -ForegroundColor Cyan
}

function Write-Success([string]$msg) {
    Write-Host "  ✓ $msg" -ForegroundColor Green
}

function Write-Fail([string]$msg) {
    Write-Host "  ✗ $msg" -ForegroundColor Red
}

# -- Test Log Directory --
$TestLogDir = Join-Path $PSScriptRoot "data" "test-logs"

function Ensure-TestLogDir {
    if (-not (Test-Path $TestLogDir)) {
        New-Item -ItemType Directory -Path $TestLogDir -Force | Out-Null
    }
}

function Write-TestLogs([string[]]$rawOutput) {
    Ensure-TestLogDir

    $passingFile = Join-Path $TestLogDir "passing-tests.txt"
    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    $rawFile     = Join-Path $TestLogDir "raw-output.txt"

    # Clear previous log files before writing new results
    @($passingFile, $failingFile, $rawFile) | ForEach-Object {
        if (Test-Path $_) { Remove-Item $_ -Force }
    }

    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $passing = [System.Collections.Generic.List[string]]::new()
    $failing = [System.Collections.Generic.List[string]]::new()

    # Save raw output for debugging
    Set-Content -Path $rawFile -Value ($rawOutput -join "`n") -Encoding UTF8

    # Two-pass approach:
    # Pass 1: Identify which tests passed and which failed
    $failedNames = [System.Collections.Generic.HashSet[string]]::new()
    $passedNames = [System.Collections.Generic.HashSet[string]]::new()

    foreach ($line in $rawOutput) {

        if ($line -match "^\s*--- FAIL:\s+(.+?)\s+\(") {
            $failedNames.Add($Matches[1].Trim()) | Out-Null
        }
        elseif ($line -match "^\s*--- PASS:\s+(.+?)\s+\(") {
            $passedNames.Add($Matches[1].Trim()) | Out-Null
        }
    }

    # Pass 2: Collect diagnostic details for failed tests
    $currentTest = ""
    $currentBlock = [System.Collections.Generic.List[string]]::new()

    foreach ($line in $rawOutput) {

        if ($line -match "^=== RUN\s+(.+)$") {
            # Flush previous block if it was a failed test
            if ($currentTest -and $failedNames.Contains($currentTest)) {
                $failing.Add("FAIL: $currentTest")

                foreach ($detail in $currentBlock) {
                    $failing.Add("  $detail")
                }

                $failing.Add("")
            }

            $currentTest = $Matches[1].Trim()
            $currentBlock.Clear()
        }
        elseif ($line -match "^\s*--- PASS:\s+(.+?)\s+\(") {
            # Passing test — flush and reset
            $currentTest = ""
            $currentBlock.Clear()
        }
        elseif ($line -match "^\s*--- FAIL:\s+(.+?)\s+\(") {
            # Capture the --- FAIL line itself as part of diagnostics
            if ($currentTest) {
                $currentBlock.Add($line)
            }
        }
        else {
            if ($currentTest) {
                # Keep all diagnostic lines: t.Errorf output, diff lines,
                # assertion details (expected vs actual), file:line references
                $currentBlock.Add($line)
            }
        }
    }

    # Flush last block
    if ($currentTest -and $failedNames.Contains($currentTest)) {
        $failing.Add("FAIL: $currentTest")

        foreach ($detail in $currentBlock) {
            $failing.Add("  $detail")
        }

        $failing.Add("")
    }

    # Collect passing test names
    foreach ($name in $passedNames) {
        $passing.Add($name)
    }

    # Write passing tests
    $passingContent = @("# Passing Tests — $timestamp", "# Count: $($passing.Count)", "")
    $passingContent += $passing
    Set-Content -Path $passingFile -Value ($passingContent -join "`n") -Encoding UTF8

    # Write failing tests
    $failCount = $failedNames.Count
    $failingContent = @("# Failing Tests — $timestamp", "# Count: $failCount", "")

    # Summary section: list failed test names first
    if ($failCount -gt 0) {
        $failingContent += "# ── Summary ──"
        $sortedFailed = $failedNames | Sort-Object
        foreach ($name in $sortedFailed) {
            $failingContent += "  - $name"
        }
        $failingContent += @("", "# ── Details ──", "")
    }
    $failingContent += $failing

    # Also capture compilation errors (no === RUN lines at all)
    $hasAnyRun = $rawOutput | Where-Object { $_ -match "^=== RUN" } | Select-Object -First 1

    if (-not $hasAnyRun) {
        $compileErrors = $rawOutput | Where-Object {
            $_ -match "\.go:\d+:" -or $_ -match "^#\s+" -or $_ -match "FAIL\s+"
        }

        if ($compileErrors) {
            $failingContent += @("", "# Compilation Errors:", "")
            $failingContent += $compileErrors
            $failCount = $failCount + 1
        }
    }

    Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8

    $passCount = $passing.Count

    Write-Host ""
    if ($passCount -gt 0) { Write-Success "$passCount passing test(s) → $passingFile" }
    if ($failCount -gt 0) { Write-Fail "$failCount failing test(s) → $failingFile" }
    elseif ($failCount -eq 0) { Write-Success "No failing tests" }
    Write-Host "  Raw output → $rawFile" -ForegroundColor Gray
}

function Invoke-GoTestAndLog([string]$testArgs) {
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 $testArgs 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    # Print to console
    $output | ForEach-Object { Write-Host $_ }

    # Write logs
    Write-TestLogs $output

    return $exitCode
}

# -- Commands --

function Invoke-GitPull {
    Write-Header "Pulling latest from remote"
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    git pull 2>&1 | ForEach-Object { Write-Host "  $_" -ForegroundColor Gray }
    if ($LASTEXITCODE -eq 0) { Write-Success "Git pull complete" }
    else { Write-Fail "git pull failed (continuing anyway)" }
    $ErrorActionPreference = $prevPref
}

function Invoke-FetchLatest {
    Invoke-GitPull
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) { Write-Success "Dependencies up to date" }
    else { Write-Fail "go mod tidy failed" }
}

function Invoke-BuildCheck([string]$buildPath) {
    Write-Header "Build check: $buildPath"
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go build $buildPath 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    if ($exitCode -ne 0) {
        Write-Fail "Build failed — skipping tests"

        Ensure-TestLogDir
        $failingFile = Join-Path $TestLogDir "failing-tests.txt"
        $rawFile     = Join-Path $TestLogDir "raw-output.txt"
        $timestamp   = Get-Date -Format "yyyy-MM-dd HH:mm:ss"

        $failingContent = @(
            "# Failing Tests — $timestamp",
            "# Count: 0",
            "",
            "# Build Failed — tests were NOT run",
            "",
            "# ── Build Errors ──",
            ""
        )
        $failingContent += $output

        Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8
        Set-Content -Path $rawFile -Value ($output -join "`n") -Encoding UTF8

        $output | ForEach-Object { Write-Host "  $_" -ForegroundColor Red }
        Open-FailingTestsIfAny
        return $false
    }

    Write-Success "Build OK"
    return $true
}

function Open-FailingTestsIfAny {
    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    if ((Test-Path $failingFile)) {
        $content = Get-Content $failingFile -Raw
        if ($content -and $content -notmatch '# Count: 0') {
            Write-Host ""
            Write-Host "  Opening failing tests log..." -ForegroundColor Yellow
            Start-Process $failingFile
        }
    }
}

function Invoke-AllTests {
    Write-Header "Running all tests"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 ./... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "All tests passed" }
        else { Write-Fail "Some tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-PackageTests([string]$pkg) {
    if (-not $pkg) {
        Write-Fail "Package name required. Usage: ./run.ps1 TP <package>"
        Write-Host "  Available packages:" -ForegroundColor Yellow
        Get-ChildItem -Path tests/integratedtests -Directory | ForEach-Object {
            Write-Host "    - $($_.Name)" -ForegroundColor Gray
        }
        return
    }

    Write-Header "Running tests for package: $pkg"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./integratedtests/$pkg/...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 "./integratedtests/$pkg/..." 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "Package tests passed" }
        else { Write-Fail "Package tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-TestCoverage {
    Write-Header "Running tests with coverage"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./...")) { return }

        $coverDir = Join-Path $PSScriptRoot "data" "coverage"
        if (-not (Test-Path $coverDir)) {
            New-Item -ItemType Directory -Path $coverDir -Force | Out-Null
        }

        $coverProfile = Join-Path $coverDir "coverage.out"
        $coverHtml    = Join-Path $coverDir "coverage.html"
        $coverSummary = Join-Path $coverDir "coverage-summary.txt"

        # Discover test packages, excluding testwrappers (test data only, 0% noise)
        $allPkgs = & go list ./... 2>&1 | ForEach-Object { $_.ToString() }
        $testPkgs = $allPkgs | Where-Object { $_ -notmatch "testwrappers" }

        # Use -coverpkg to measure coverage of the MAIN source packages
        # (the parent module), not just the test helper packages
        $coverpkg = "github.com/alimtvnetwork/core/..."

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 -coverprofile=$coverProfile -coverpkg=$coverpkg @testPkgs 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if (Test-Path $coverProfile) {
            # Generate func-level summary
            $funcOutput = & go tool cover -func=$coverProfile 2>&1 | ForEach-Object { $_.ToString() }

            # Generate HTML report
            & go tool cover -html=$coverProfile -o $coverHtml 2>&1 | Out-Null

            # Build summary report
            $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
            $summaryLines = [System.Collections.Generic.List[string]]::new()
            $summaryLines.Add("# Coverage Summary — $timestamp")
            $summaryLines.Add("")

            # Extract total line
            $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
            if ($totalLine) {
                $summaryLines.Add("## Total Coverage")
                $summaryLines.Add("  $totalLine")
                $summaryLines.Add("")
            }

            # Extract per-package coverage from test output
            $pkgCoverLines = $output | Where-Object { $_ -match "coverage:" } | Sort-Object
            if ($pkgCoverLines) {
                $summaryLines.Add("## Per-Package Coverage")
                foreach ($line in $pkgCoverLines) {
                    $summaryLines.Add("  $line")
                }
                $summaryLines.Add("")
            }

            # Extract low-coverage functions (< 50%)
            $lowCovFuncs = [System.Collections.Generic.List[string]]::new()
            foreach ($line in $funcOutput) {
                if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
                    $pct = [double]$Matches[1]
                    if ($pct -lt 50.0) {
                        $lowCovFuncs.Add("  $line")
                    }
                }
            }

            if ($lowCovFuncs.Count -gt 0) {
                $summaryLines.Add("## Low Coverage Functions (< 50%)")
                $summaryLines.Add("  Count: $($lowCovFuncs.Count)")
                $summaryLines.Add("")
                foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }
                $summaryLines.Add("")
            }

            # File paths
            $summaryLines.Add("## Reports")
            $summaryLines.Add("  Profile:  $coverProfile")
            $summaryLines.Add("  HTML:     $coverHtml")
            $summaryLines.Add("  Summary:  $coverSummary")

            Set-Content -Path $coverSummary -Value ($summaryLines -join "`n") -Encoding UTF8

            # Print summary to console
            Write-Host ""
            if ($totalLine) {
                Write-Host "  $totalLine" -ForegroundColor Cyan
            }
            Write-Host ""
            Write-Success "Coverage profile:  $coverProfile"
            Write-Success "HTML report:       $coverHtml"
            Write-Success "Summary:           $coverSummary"

            if ($lowCovFuncs.Count -gt 0) {
                Write-Host ""
                Write-Host "  ⚠ $($lowCovFuncs.Count) function(s) below 50% coverage" -ForegroundColor Yellow
            }

            # Auto-open HTML report
            $openHtml = $true
            if ($ExtraArgs -and $ExtraArgs[0] -eq "--no-open") { $openHtml = $false }
            if ($openHtml -and (Test-Path $coverHtml)) {
                Write-Host ""
                Write-Host "  Opening HTML coverage report..." -ForegroundColor Yellow
                Start-Process $coverHtml
            }
        }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-IntegratedTests {
    Write-Header "Running integrated tests only"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./integratedtests/...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 ./integratedtests/... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "Integrated tests passed" }
        else { Write-Fail "Integrated tests failed (exit code: $exitCode)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-RunMain {
    Write-Header "Running main application"
    go run ./cmd/main/*.go
}

function Invoke-Build {
    Write-Header "Building binary"
    $buildDir = "build"
    if (-not (Test-Path $buildDir)) { New-Item -ItemType Directory -Path $buildDir | Out-Null }
    go build -o "$buildDir/cli" ./cmd/main/
    if ($LASTEXITCODE -eq 0) { Write-Success "Build complete: $buildDir/cli" }
    else { Write-Fail "Build failed" }
}

function Invoke-BuildRun {
    Invoke-Build
    if ($LASTEXITCODE -eq 0) {
        Write-Header "Running built binary"
        & ./build/cli
    }
}

function Invoke-Format {
    Write-Header "Formatting Go files"
    gofmt -w -s .
    Write-Success "Formatting complete"
}

function Invoke-Vet {
    Write-Header "Running go vet"
    go vet ./...
    if ($LASTEXITCODE -eq 0) { Write-Success "No issues found" }
    else { Write-Fail "Issues found" }
}

function Invoke-Tidy {
    Write-Header "Running go mod tidy"
    go mod tidy
    Write-Success "Tidy complete"
}

function Invoke-GoConvey {
    Write-Header "Launching GoConvey"

    # Check if goconvey is installed
    $gcPath = Get-Command goconvey -ErrorAction SilentlyContinue
    if (-not $gcPath) {
        Write-Host "  GoConvey not found. Installing..." -ForegroundColor Yellow
        go install github.com/smartystreets/goconvey@latest
        if ($LASTEXITCODE -ne 0) {
            Write-Fail "Failed to install GoConvey"
            return
        }
        Write-Success "GoConvey installed"
    }

    $port = if ($ExtraArgs -and $ExtraArgs[0]) { $ExtraArgs[0] } else { "8080" }
    Write-Host "  Starting GoConvey on http://localhost:$port" -ForegroundColor Yellow
    Write-Host "  Press Ctrl+C to stop" -ForegroundColor Gray

    Push-Location tests
    try {
        goconvey -port $port
    }
    finally { Pop-Location }
}

function Invoke-Clean {
    Write-Header "Cleaning build artifacts"
    if (Test-Path build) { Remove-Item -Recurse -Force build }
    if (Test-Path tests/coverage.out) { Remove-Item tests/coverage.out }
    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    if (Test-Path $coverDir) { Remove-Item -Recurse -Force $coverDir; Write-Success "Removed coverage reports" }
    Write-Success "Clean complete"
}

function Invoke-ShowFailLog {
    $failingFile = Join-Path $TestLogDir "failing-tests.txt"
    if (-not (Test-Path $failingFile)) {
        Write-Header "No failing tests log found"
        Write-Host "  Run tests first: ./run.ps1 T" -ForegroundColor Yellow
        return
    }

    Write-Header "Last Failing Tests"
    $content = Get-Content $failingFile -Raw
    if ($content -match '# Count: 0') {
        Write-Success "No failing tests in last run"
    }
    else {
        Write-Host $content
    }
    Write-Host ""
    Write-Host "  Log file: $failingFile" -ForegroundColor Gray
}

function Show-Help {
    Write-Host ""
    Write-Host "  Project Runner — ./run.ps1 <command>" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  Testing:" -ForegroundColor Yellow
    Write-Host "    T   | -t   | test          Run all tests (verbose)"
    Write-Host "    TP  | -tp  | test-pkg      Run tests for a specific package"
    Write-Host "    TC  | -tc  | test-cover    Run tests with coverage (HTML + summary)"
    Write-Host "    TI  | -ti  | test-int      Run integrated tests only"
    Write-Host "    TF  | -tf  | test-fail     Show last failing tests log"
    Write-Host "    GC  | -gc  | goconvey      Launch GoConvey (browser test runner)"
    Write-Host ""
    Write-Host "  Build & Run:" -ForegroundColor Yellow
    Write-Host "    R   | -r   | run           Run the main application"
    Write-Host "    B   | -b   | build         Build the binary"
    Write-Host "    BR  | -br  | build-run     Build then run"
    Write-Host ""
    Write-Host "  Code Quality:" -ForegroundColor Yellow
    Write-Host "    F   | -f   | fmt           Format all Go files"
    Write-Host "    L   | -l   | lint          Run go vet"
    Write-Host "    V   | -v   | vet           Run go vet"
    Write-Host "    TY  | -ty  | tidy          Run go mod tidy"
    Write-Host ""
    Write-Host "  Other:" -ForegroundColor Yellow
    Write-Host "    C   | -c   | clean         Clean build artifacts"
    Write-Host "    H   | -h   | help          Show this help"
    Write-Host ""
    Write-Host "  Examples:" -ForegroundColor Gray
    Write-Host "    ./run.ps1 T"
    Write-Host "    ./run.ps1 -t"
    Write-Host "    ./run.ps1 TP regexnewtests"
    Write-Host "    ./run.ps1 -tp regexnewtests"
    Write-Host "    ./run.ps1 -gc"
    Write-Host "    ./run.ps1 -gc 9090          (custom port)"
    Write-Host ""
}

# -- Dispatch --
switch ($Command.ToLower()) {
    { $_ -in "t", "-t", "test" }              { Invoke-AllTests }
    { $_ -in "tp", "-tp", "test-pkg" }        { Invoke-PackageTests $ExtraArgs[0] }
    { $_ -in "tc", "-tc", "test-cover" }      { Invoke-TestCoverage }
    { $_ -in "ti", "-ti", "test-int" }        { Invoke-IntegratedTests }
    { $_ -in "tf", "-tf", "test-fail" }       { Invoke-ShowFailLog }
    { $_ -in "gc", "-gc", "goconvey" }        { Invoke-GoConvey }
    { $_ -in "r", "-r", "run" }               { Invoke-RunMain }
    { $_ -in "b", "-b", "build" }             { Invoke-Build }
    { $_ -in "br", "-br", "build-run" }       { Invoke-BuildRun }
    { $_ -in "f", "-f", "fmt" }               { Invoke-Format }
    { $_ -in "l", "-l", "lint", "v", "-v", "vet" } { Invoke-Vet }
    { $_ -in "ty", "-ty", "tidy" }            { Invoke-Tidy }
    { $_ -in "c", "-c", "clean" }             { Invoke-Clean }
    { $_ -in "h", "-h", "help", "" }          { Show-Help }
    default {
        Write-Fail "Unknown command: '$Command'"
        Show-Help
    }
}
