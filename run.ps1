#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Project runner script with shorthands for common operations.

.DESCRIPTION
    Usage: ./run.ps1 <command>

    Commands:
        T     | test          Run all tests (verbose)
        TP    | test-pkg      Run tests for a specific package: ./run.ps1 TP regexnewtests
        TC    | test-cover    Run tests with coverage report
        TI    | test-int      Run integrated tests only
        R     | run           Run the main application
        B     | build         Build the binary
        BR    | build-run     Build then run
        F     | fmt           Format all Go files
        L     | lint          Run go vet on all packages
        V     | vet           Run go vet
        TY    | tidy          Run go mod tidy
        C     | clean         Clean build artifacts
        H     | help          Show this help

.EXAMPLE
    ./run.ps1 T
    ./run.ps1 test
    ./run.ps1 TP regexnewtests
    ./run.ps1 TC
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

# -- Commands --

function Invoke-AllTests {
    Write-Header "Running all tests"
    Push-Location tests
    try {
        go test -v -count=1 ./...
        if ($LASTEXITCODE -eq 0) { Write-Success "All tests passed" }
        else { Write-Fail "Some tests failed (exit code: $LASTEXITCODE)" }
    }
    finally { Pop-Location }
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
    Push-Location tests
    try {
        go test -v -count=1 "./integratedtests/$pkg/..."
        if ($LASTEXITCODE -eq 0) { Write-Success "Package tests passed" }
        else { Write-Fail "Package tests failed (exit code: $LASTEXITCODE)" }
    }
    finally { Pop-Location }
}

function Invoke-TestCoverage {
    Write-Header "Running tests with coverage"
    Push-Location tests
    try {
        go test -v -count=1 -coverprofile=coverage.out ./...
        if (Test-Path coverage.out) {
            go tool cover -func=coverage.out
            Write-Success "Coverage report generated: tests/coverage.out"
            Write-Host "  Run 'go tool cover -html=tests/coverage.out' to view in browser" -ForegroundColor Yellow
        }
    }
    finally { Pop-Location }
}

function Invoke-IntegratedTests {
    Write-Header "Running integrated tests only"
    Push-Location tests
    try {
        go test -v -count=1 ./integratedtests/...
        if ($LASTEXITCODE -eq 0) { Write-Success "Integrated tests passed" }
        else { Write-Fail "Integrated tests failed (exit code: $LASTEXITCODE)" }
    }
    finally { Pop-Location }
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

function Invoke-Clean {
    Write-Header "Cleaning build artifacts"
    if (Test-Path build) { Remove-Item -Recurse -Force build }
    if (Test-Path tests/coverage.out) { Remove-Item tests/coverage.out }
    Write-Success "Clean complete"
}

function Show-Help {
    Write-Host ""
    Write-Host "  Project Runner — ./run.ps1 <command>" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  Testing:" -ForegroundColor Yellow
    Write-Host "    T   | test          Run all tests (verbose)"
    Write-Host "    TP  | test-pkg      Run tests for a specific package"
    Write-Host "    TC  | test-cover    Run tests with coverage"
    Write-Host "    TI  | test-int      Run integrated tests only"
    Write-Host ""
    Write-Host "  Build & Run:" -ForegroundColor Yellow
    Write-Host "    R   | run           Run the main application"
    Write-Host "    B   | build         Build the binary"
    Write-Host "    BR  | build-run     Build then run"
    Write-Host ""
    Write-Host "  Code Quality:" -ForegroundColor Yellow
    Write-Host "    F   | fmt           Format all Go files"
    Write-Host "    L   | lint          Run go vet"
    Write-Host "    V   | vet           Run go vet"
    Write-Host "    TY  | tidy          Run go mod tidy"
    Write-Host ""
    Write-Host "  Other:" -ForegroundColor Yellow
    Write-Host "    C   | clean         Clean build artifacts"
    Write-Host "    H   | help          Show this help"
    Write-Host ""
    Write-Host "  Examples:" -ForegroundColor Gray
    Write-Host "    ./run.ps1 T"
    Write-Host "    ./run.ps1 TP regexnewtests"
    Write-Host "    ./run.ps1 TC"
    Write-Host ""
}

# -- Dispatch --
switch ($Command.ToLower()) {
    { $_ -in "t", "test" }        { Invoke-AllTests }
    { $_ -in "tp", "test-pkg" }   { Invoke-PackageTests $Args[0] }
    { $_ -in "tc", "test-cover" } { Invoke-TestCoverage }
    { $_ -in "ti", "test-int" }   { Invoke-IntegratedTests }
    { $_ -in "r", "run" }         { Invoke-RunMain }
    { $_ -in "b", "build" }       { Invoke-Build }
    { $_ -in "br", "build-run" }  { Invoke-BuildRun }
    { $_ -in "f", "fmt" }         { Invoke-Format }
    { $_ -in "l", "lint", "v", "vet" } { Invoke-Vet }
    { $_ -in "ty", "tidy" }       { Invoke-Tidy }
    { $_ -in "c", "clean" }       { Invoke-Clean }
    { $_ -in "h", "help", "" }    { Show-Help }
    default {
        Write-Fail "Unknown command: '$Command'"
        Show-Help
    }
}
