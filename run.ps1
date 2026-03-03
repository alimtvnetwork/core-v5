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
    [string[]]$Args
)

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

function Invoke-FetchLatest {
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) { Write-Success "Dependencies up to date" }
    else { Write-Fail "go mod tidy failed" }
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
        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 -coverprofile=coverage.out ./... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if (Test-Path coverage.out) {
            go tool cover -func=coverage.out
            Write-Success "Coverage report generated: tests/coverage.out"
            Write-Host "  Run 'go tool cover -html=tests/coverage.out' to view in browser" -ForegroundColor Yellow
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

    $port = if ($Args -and $Args[0]) { $Args[0] } else { "8080" }
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
    Write-Host "    TC  | -tc  | test-cover    Run tests with coverage"
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
    { $_ -in "tp", "-tp", "test-pkg" }        { Invoke-PackageTests $Args[0] }
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
