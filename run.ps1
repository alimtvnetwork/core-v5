#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Project runner script with shorthands for common operations.

.DESCRIPTION
    Usage: ./run.ps1 <command> [options]

    Commands (uppercase shorthands OR hyphen-lowercase):
        T   | -t   | test          Run all tests (verbose)
        TP  | -tp  | test-pkg      Run tests for a specific package: ./run.ps1 TP regexnewtests
        TC  | -tc  | test-cover    Run tests with coverage (parallel by default)
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

    Mode options (for TC/TCP):
        --sync        Run precompile + tests sequentially (default: parallel)
        --no-open     Skip auto-opening HTML coverage report

.EXAMPLE
    ./run.ps1 T
    ./run.ps1 -t
    ./run.ps1 TP regexnewtests
    ./run.ps1 -tp regexnewtests
    ./run.ps1 TC --sync
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
    $match = [regex]::Match($rawLine, '(?i)run\.ps1\s+(-?\w[\w-]*)\s*(.*)')
    if ($match.Success) {
        $Command = $match.Groups[1].Value
        # Capture remaining args that PowerShell swallowed
        $trailing = $match.Groups[2].Value.Trim()
        if ($trailing -and (-not $ExtraArgs -or $ExtraArgs.Count -eq 0)) {
            $ExtraArgs = @($trailing -split '\s+')
        }
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

    # Clean data folder before running tests
    $dataDir = Join-Path $PSScriptRoot "data"
    if (Test-Path $dataDir) {
        Remove-Item -Recurse -Force $dataDir
        Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow
    }

    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    $partialDir = Join-Path $coverDir "partial"
    New-Item -ItemType Directory -Path $partialDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage.out"
    $coverHtml    = Join-Path $coverDir "coverage.html"
    $coverSummary = Join-Path $coverDir "coverage-summary.txt"

    # Build coverpkg list: all source packages EXCLUDING tests/
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    # Get all test packages to run individually
    $allTestPkgs = go list ./tests/integratedtests/... 2>&1 | ForEach-Object { $_.ToString() }

    # ── Pre-coverage compile check ──────────────────────────────────
    # Compile each test package individually (go test -c) to detect
    # build failures BEFORE running coverage. Packages that fail to
    # compile are excluded from the coverage run so they don't produce
    # misleading 0% profiles or cascade failures.

    # Determine sync vs parallel mode from ExtraArgs (--sync flag)
    $isSyncMode = $false
    if ($ExtraArgs) {
        foreach ($ea in $ExtraArgs) {
            if ($ea -eq "--sync") { $isSyncMode = $true }
        }
    }

    $modeLabel = if ($isSyncMode) { "sync" } else { "parallel" }
    Write-Host ""
    Write-Header "Pre-coverage compile check ($($allTestPkgs.Count) packages, $modeLabel mode)"

    $blockedPkgs = [System.Collections.Generic.List[string]]::new()
    $blockedErrors = [System.Collections.Generic.Dictionary[string, string]]::new()
    $testPkgs = [System.Collections.Generic.List[string]]::new()
    $compileTemp = Join-Path $coverDir "compile-check"
    New-Item -ItemType Directory -Path $compileTemp -Force | Out-Null
    [int]$jobCounter = 0

    if ($isSyncMode) {
        # ── Sequential compile check ──
        foreach ($testPkg in $allTestPkgs) {
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $compileOut = & go test -c -o (Join-Path $compileTemp "test.exe") "-coverpkg=$covPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $compileExit = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($compileExit -eq 0) {
                Write-Host "  ✓ $shortName" -ForegroundColor Green
                $testPkgs.Add($testPkg)
            } else {
                Write-Host "  ✗ $shortName [build failed]" -ForegroundColor Red
                $blockedPkgs.Add($shortName)
                $errLines = ($compileOut | Where-Object { $_ -match '\.go:\d+:' }) -join "`n"
                $blockedErrors[$shortName] = $errLines
            }
        }
    } else {
        # ── Parallel compile check (ForEach-Object -Parallel, runspace-based) ──
        $throttle = [Math]::Min($allTestPkgs.Count, [Environment]::ProcessorCount)
        Write-Host "  Launching $($allTestPkgs.Count) compile checks ($throttle parallel)..." -ForegroundColor Gray

        $compileResults = $allTestPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:covPkgList
            $tempDir = $using:compileTemp
            $idx = [System.Threading.Interlocked]::Increment([ref]$using:jobCounter)
            $outFile = Join-Path $tempDir "compile-$idx.exe"
            $ErrorActionPreference = "Continue"
            # Capture output to array first, THEN read $LASTEXITCODE before any pipe resets it
            $rawOut = & go test -c -o $outFile "-coverpkg=$covPkgs" "$pkg" 2>&1
            $ec = $LASTEXITCODE
            $out = @($rawOut | ForEach-Object { $_.ToString() })
            [pscustomobject]@{
                Pkg      = $pkg
                ExitCode = $ec
                Output   = $out
            }
        }

        foreach ($result in $compileResults) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            if ($result.ExitCode -eq 0) {
                Write-Host "  ✓ $shortName" -ForegroundColor Green
                $testPkgs.Add($result.Pkg)
            } else {
                Write-Host "  ✗ $shortName [build failed]" -ForegroundColor Red
                $blockedPkgs.Add($shortName)
                # Capture specific .go errors first; fall back to ALL output if none matched
                $goLines = @($result.Output | Where-Object { $_ -match '\.go:\d+:' })
                if ($goLines.Count -gt 0) {
                    $blockedErrors[$shortName] = $goLines -join "`n"
                } else {
                    $blockedErrors[$shortName] = ($result.Output -join "`n")
                }
            }
        }
    }

    # Clean up compile artifacts
    if (Test-Path $compileTemp) { Remove-Item -Recurse -Force $compileTemp }

    # Print blocked summary
    if ($blockedPkgs.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ BLOCKED PACKAGES ($($blockedPkgs.Count) failed to compile)" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($bp in ($blockedPkgs | Sort-Object)) {
            Write-Host "  │   ✗ $bp" -ForegroundColor Red
        }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ These packages will be SKIPPED in coverage." -ForegroundColor Yellow
        Write-Host "  │ Fix their build errors to include them." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host ""

        # Write blocked details to file for AI/human review
        $blockedFile = Join-Path $coverDir "blocked-packages.txt"
        $sortedBlocked = $blockedPkgs | Sort-Object
        $blockedContent = @(
            "# Blocked Packages — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')",
            "# Count: $($blockedPkgs.Count)",
            "",
            "# ── CLI Summary ──",
            "# ┌─────────────────────────────────────────────────",
            "# │ BLOCKED PACKAGES ($($blockedPkgs.Count) failed to compile)"
        )
        foreach ($bp in $sortedBlocked) {
            $blockedContent += "# │   ✗ $bp"
        }
        $blockedContent += @(
            "# │",
            "# │ These packages will be SKIPPED in coverage.",
            "# │ Fix their build errors to include them.",
            "# └─────────────────────────────────────────────────",
            "",
            "# ── Package List ──"
        )
        foreach ($bp in $sortedBlocked) {
            $blockedContent += "# - $bp"
        }
        $blockedContent += ""

        # Write to project root as well
        $rootBlockedFile = Join-Path $PSScriptRoot "blocked-packages.txt"

        foreach ($bp in $sortedBlocked) {
            $blockedContent += "## $bp"
            if ($blockedErrors.ContainsKey($bp)) {
                $blockedContent += $blockedErrors[$bp]
            }
            $blockedContent += ""
        }
        $fileContent = $blockedContent -join "`n"
        Set-Content -Path $blockedFile -Value $fileContent -Encoding UTF8
        Set-Content -Path $rootBlockedFile -Value $fileContent -Encoding UTF8
        Write-Host "  Blocked details → $blockedFile" -ForegroundColor Gray
    } else {
        Write-Host ""
        Write-Success "All $($testPkgs.Count) packages compiled successfully"
    }

    if ($testPkgs.Count -eq 0) {
        Write-Fail "No packages compiled — aborting coverage run"
        return
    }

    # ── Coverage run (only compilable packages) ─────────────────────
    $allOutput = [System.Collections.Generic.List[string]]::new()
    $pkgCoverMap = [ordered]@{}
    $overallExit = 0
    $pkgIndex = 0

    Write-Host ""
    Write-Host "  Running $($testPkgs.Count) test packages with individual coverage profiles ($modeLabel)..." -ForegroundColor Yellow
    Write-Host ""

    if ($isSyncMode) {
        # ── Sequential coverage run ──
        foreach ($testPkg in $testPkgs) {
            $pkgIndex++
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }
            $srcTarget = $shortName -replace 'tests$', '' -replace 'tests/', '/'
            if (-not $srcTarget) { $srcTarget = $shortName }

            $partialProfile = Join-Path $partialDir "cover-$pkgIndex.out"

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $output = & go test -v -count=1 "-coverprofile=$partialProfile" "-coverpkg=$covPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $pkgExit = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($pkgExit -ne 0) { $overallExit = $pkgExit }

            $statusIcon = if ($pkgExit -eq 0) { "✓" } else { "✗" }
            $statusColor = if ($pkgExit -eq 0) { "Green" } else { "Red" }

            $partialPct = ""
            if (Test-Path $partialProfile) {
                $srcMatchPattern = $srcTarget -replace '/', '/'
                $pStmts = 0; $pCovered = 0
                $pTotalLines = 0; $pMatchedLines = 0
                foreach ($pLine in (Get-Content $partialProfile)) {
                    if ($pLine -match "^mode:") { continue }
                    $pTotalLines++
                    if ($pLine -notmatch "/$srcMatchPattern/") { continue }
                    $pMatchedLines++
                    if ($pLine -match "\s+(\d+)\s+(\d+)\s*$") {
                        $pStmts += [int]$Matches[1]
                        if ([int]$Matches[2] -gt 0) { $pCovered += [int]$Matches[1] }
                    }
                }
                if ($pStmts -gt 0) {
                    $partialPct = " — $([math]::Round(($pCovered / $pStmts) * 100, 1))%"
                }
                Write-Host "    [debug] filter=/$srcMatchPattern/ matched=$pMatchedLines/$pTotalLines stmts=$pCovered/$pStmts" -ForegroundColor DarkGray
            }

            Write-Host "  [$pkgIndex/$($testPkgs.Count)] $statusIcon $srcTarget$partialPct" -ForegroundColor $statusColor

            if ($output) { foreach ($line in $output) { $allOutput.Add([string]$line) } }
        }
    } else {
        # ── Parallel coverage run (ForEach-Object -Parallel) ──
        $throttle = [Math]::Min($testPkgs.Count, [Environment]::ProcessorCount)
        Write-Host "  Launching $($testPkgs.Count) test packages ($throttle parallel)..." -ForegroundColor Gray
        [int]$coverCounter = 0

        $coverResults = $testPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:covPkgList
            $pDir = $using:partialDir
            $idx = [System.Threading.Interlocked]::Increment([ref]$using:coverCounter)
            $profile = Join-Path $pDir "cover-$idx.out"
            $ErrorActionPreference = "Continue"
            $out = & go test -v -count=1 "-coverprofile=$profile" "-coverpkg=$covPkgs" "$pkg" 2>&1 | ForEach-Object { $_.ToString() }
            [pscustomobject]@{
                Pkg      = $pkg
                Profile  = $profile
                Index    = $idx
                ExitCode = $LASTEXITCODE
                Output   = $out
            }
        }

        # Collect results in order
        foreach ($result in ($coverResults | Sort-Object Index)) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }
            $srcTarget = $shortName -replace 'tests$', '' -replace 'tests/', '/'
            if (-not $srcTarget) { $srcTarget = $shortName }

            if ($result.ExitCode -ne 0) { $overallExit = $result.ExitCode }

            $statusIcon = if ($result.ExitCode -eq 0) { "✓" } else { "✗" }
            $statusColor = if ($result.ExitCode -eq 0) { "Green" } else { "Red" }

            $partialPct = ""
            if (Test-Path $result.Profile) {
                $srcMatchPattern = $srcTarget -replace '/', '/'
                $pStmts = 0; $pCovered = 0
                $pTotalLines = 0; $pMatchedLines = 0
                foreach ($pLine in (Get-Content $result.Profile)) {
                    if ($pLine -match "^mode:") { continue }
                    $pTotalLines++
                    if ($pLine -notmatch "/$srcMatchPattern/") { continue }
                    $pMatchedLines++
                    if ($pLine -match "\s+(\d+)\s+(\d+)\s*$") {
                        $pStmts += [int]$Matches[1]
                        if ([int]$Matches[2] -gt 0) { $pCovered += [int]$Matches[1] }
                    }
                }
                if ($pStmts -gt 0) {
                    $partialPct = " — $([math]::Round(($pCovered / $pStmts) * 100, 1))%"
                }
                Write-Host "    [debug] filter=/$srcMatchPattern/ matched=$pMatchedLines/$pTotalLines stmts=$pCovered/$pStmts" -ForegroundColor DarkGray
            }

            Write-Host "  [$($result.Index)/$($testPkgs.Count)] $statusIcon $srcTarget$partialPct" -ForegroundColor $statusColor

            if ($result.Output) { foreach ($line in $result.Output) { $allOutput.Add([string]$line) } }
        }
        $pkgIndex = $testPkgs.Count
    }

    # Print to console
    $allOutput | ForEach-Object { Write-Host $_ }
    Write-TestLogs $allOutput.ToArray()

    # Merge all partial profiles into one, using MAX count per unique line.
    # This is critical because -coverpkg instruments ALL source packages in every test run,
    # so each line appears N times (once per test package). Without dedup, the last occurrence
    # (usually count=0 from a package that didn't exercise this code) overwrites the covered entry.
    Write-Host ""
    Write-Host "  Merging $pkgIndex coverage profiles (max-count dedup)..." -ForegroundColor Yellow

    $partialFiles = Get-ChildItem -Path $partialDir -Filter "cover-*.out" | Sort-Object Name
    $coverMap = [System.Collections.Generic.Dictionary[string, int]]::new()

    foreach ($pf in $partialFiles) {
        $lines = Get-Content $pf.FullName
        foreach ($line in $lines) {
            if (-not $line -or $line -match "^mode:") { continue }
            # Coverage line format: "pkg/file.go:startLine.startCol,endLine.endCol numStatements count"
            # Require full format with colon before line numbers to reject malformed lines
            if ($line -match "^(\S+\.go:\d+\.\d+,\d+\.\d+\s+\d+)\s+(\d+)\s*$") {
                $key = $Matches[1]
                $count = [int]$Matches[2]
                if ($coverMap.ContainsKey($key)) {
                    if ($count -gt $coverMap[$key]) {
                        $coverMap[$key] = $count
                    }
                } else {
                    $coverMap[$key] = $count
                }
            }
        }
    }

    $mergedLines = [System.Collections.Generic.List[string]]::new()
    $mergedLines.Add("mode: set")
    foreach ($entry in $coverMap.GetEnumerator()) {
        $mergedLines.Add("$($entry.Key) $($entry.Value)")
    }

    Set-Content -Path $coverProfile -Value ($mergedLines -join "`n") -Encoding UTF8
    Write-Success "Merged profile: $coverProfile ($($coverMap.Count) unique coverage lines)"

    # Keep partial profiles for per-package inspection
    Write-Success "Partial profiles kept in: $partialDir"

    if (Test-Path $coverProfile) {
        # Generate func-level summary
        Write-Host "  [debug] coverProfile = $coverProfile" -ForegroundColor DarkGray
        Write-Host "  [debug] coverHtml    = $coverHtml" -ForegroundColor DarkGray
        Write-Host "  [debug] file exists  = $(Test-Path $coverProfile)" -ForegroundColor DarkGray

        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }

        # Generate HTML report — use explicit argument list to avoid variable interpolation issues
        $htmlArgs = @("-html=$coverProfile", "-o=$coverHtml")
        Write-Host "  [debug] go tool cover args: $($htmlArgs -join ' ')" -ForegroundColor DarkGray
        $htmlErr = & go tool cover $htmlArgs 2>&1
        $htmlExitCode = $LASTEXITCODE

        if ($htmlExitCode -ne 0 -or -not (Test-Path $coverHtml)) {
            Write-Host "  ⚠ Failed to generate HTML report via 'go tool cover -html' (exit: $htmlExitCode)" -ForegroundColor Red
            if ($htmlErr) { Write-Host "  Error: $htmlErr" -ForegroundColor Red }
            Write-Host "  [debug] Attempted command: go tool cover -html=`"$coverProfile`" -o=`"$coverHtml`"" -ForegroundColor DarkGray
            # Fallback: generate a basic HTML from the func output
            $fallbackHtml = @"
<!DOCTYPE html><html><head><meta charset="utf-8"><title>Coverage Report</title>
<style>body{font-family:monospace;padding:20px;background:#1e1e2e;color:#cdd6f4}
pre{white-space:pre-wrap}</style></head><body>
<h1>Coverage Report</h1><pre>$($funcOutput -join "`n")</pre></body></html>
"@
            Set-Content -Path $coverHtml -Value $fallbackHtml -Encoding UTF8
            Write-Host "  Generated fallback HTML report" -ForegroundColor Yellow
        }

        # Build AI-friendly coverage text for the copy button
        $aiTextLines = [System.Collections.Generic.List[string]]::new()
        $aiTextLines.Add("# Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
        $aiTextLines.Add("")

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

        # Per-SOURCE-package coverage from merged profile
        $srcPkgCovMap = [ordered]@{}
        foreach ($line in $funcOutput) {
            if ($line -match "^(\S+):" -and $line -notmatch "^total:") {
                $srcPkg = $Matches[1]
                # Extract short name from full import path
                $shortSrc = $srcPkg -replace '.*alimtvnetwork/core/?', ''
                if (-not $shortSrc) { $shortSrc = "(root)" }
                if (-not $srcPkgCovMap.Contains($shortSrc)) {
                    $srcPkgCovMap[$shortSrc] = @{ Stmts = 0; Covered = 0 }
                }
            }
        }
        # Parse coverage.out lines to compute per-source-package %
        $srcPkgStmts = [ordered]@{}
        foreach ($covLine2 in $mergedLines) {
            if ($covLine2 -match "^mode:") { continue }
            # Format: pkg/file.go:startLine.col,endLine.col numStmts count
            if ($covLine2 -match "^(\S+?):(\d+)\.(\d+),(\d+)\.(\d+)\s+(\d+)\s+(\d+)") {
                $filePath2 = $Matches[1]
                $stmts = [int]$Matches[6]
                $count = [int]$Matches[7]
                # Extract package from file path
                $shortSrc2 = $filePath2 -replace '.*alimtvnetwork/core/?', ''
                $shortSrc2 = $shortSrc2 -replace '/[^/]+$', ''  # remove filename
                if (-not $shortSrc2) { $shortSrc2 = "(root)" }
                if (-not $srcPkgStmts.Contains($shortSrc2)) {
                    $srcPkgStmts[$shortSrc2] = @{ Stmts = 0; Covered = 0 }
                }
                $srcPkgStmts[$shortSrc2].Stmts += $stmts
                if ($count -gt 0) { $srcPkgStmts[$shortSrc2].Covered += $stmts }
            }
        }
        if ($srcPkgStmts.Count -gt 0) {
            $summaryLines.Add("## Per-Package Coverage (Source)")
            $sortedSrcPkgs = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal }
            } | Sort-Object Pct -Descending
            foreach ($entry in $sortedSrcPkgs) {
                $summaryLines.Add("  $($entry.Pct)%`t$($entry.Name)")
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

        # Build AI-friendly text for copy button
        $aiTextLines.Add("## Goal: Improve test coverage for the packages listed below.")
        $aiTextLines.Add("Please write tests for uncovered functions, following the project's AAA pattern.")
        $aiTextLines.Add("")
        if ($totalLine) {
            $aiTextLines.Add("## Total Coverage")
            $aiTextLines.Add($totalLine)
            $aiTextLines.Add("")
        }
        if ($srcPkgStmts.Count -gt 0) {
            $aiTextLines.Add("## Per-Source-Package Coverage")
            $computedSrcPkgs = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal3 = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal3; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
            } | Sort-Object Pct
            foreach ($e in $computedSrcPkgs) {
                $aiTextLines.Add("  $($e.Pct)%  $($e.Name)  ($($e.Covered)/$($e.Stmts) stmts)")
            }
            $aiTextLines.Add("")
        }
        if ($lowCovFuncs.Count -gt 0) {
            $aiTextLines.Add("## Uncovered/Low-Coverage Functions (< 50%)")
            $aiTextLines.Add("Count: $($lowCovFuncs.Count)")
            $aiTextLines.Add("")
            foreach ($f in $lowCovFuncs) { $aiTextLines.Add($f.TrimStart()) }
            $aiTextLines.Add("")
        }
        $aiTextLines.Add("## Instructions")
        $aiTextLines.Add("- Tests go in tests/integratedtests/{pkg}tests/")
        $aiTextLines.Add("- Use CaseV1 table-driven pattern with AAA comments")
        $aiTextLines.Add("- Focus on the lowest coverage packages first")

        $aiTextEscaped = ($aiTextLines -join "`n") -replace '\\', '\\\\' -replace "'", "\\\'" -replace "`n", '\n' -replace "`r", '' -replace '"', '\"'

        # Inject "Copy for AI" button into the Go HTML report
        if (Test-Path $coverHtml) {
            $htmlContent = Get-Content -Path $coverHtml -Raw

            $buttonHtml = @'
<div id="ai-copy-panel" style="position:fixed;top:12px;right:12px;z-index:9999;font-family:system-ui,sans-serif;">
<button onclick="copyForAI()" style="
  background:linear-gradient(135deg,#6366f1,#8b5cf6);color:#fff;border:none;
  padding:10px 20px;border-radius:8px;font-size:14px;font-weight:600;
  cursor:pointer;box-shadow:0 4px 12px rgba(99,102,241,0.4);
  display:flex;align-items:center;gap:6px;transition:all 0.2s;
" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'">
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="9" y="9" width="13" height="13" rx="2"/><path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/></svg>
  Copy for AI
</button>
<span id="ai-copy-status" style="display:none;color:#22c55e;font-size:13px;margin-top:4px;text-align:center;">Copied!</span>
</div>
<script>
var __aiCoverageText =
'@
            # Insert the escaped text between the two halves
            $scriptEnd = @'
';
function copyForAI(){
  try {
    var ta = document.createElement("textarea");
    ta.value = __aiCoverageText;
    ta.style.position = "fixed";
    ta.style.left = "-9999px";
    document.body.appendChild(ta);
    ta.select();
    document.execCommand("copy");
    document.body.removeChild(ta);
    var s = document.getElementById("ai-copy-status");
    s.style.display = "block";
    setTimeout(function(){ s.style.display = "none"; }, 2000);
  } catch(e) {
    alert("Copy failed: " + e.message);
  }
}
</script>
'@
            $injectedHtml = $buttonHtml + $aiTextEscaped + $scriptEnd
            $htmlContent = $htmlContent -replace '</body>', ($injectedHtml + "`n</body>")
            Set-Content -Path $coverHtml -Value $htmlContent -Encoding UTF8
            Write-Host "  ✓ Injected 'Copy for AI' button into HTML report" -ForegroundColor Green
        }

        # Print per-source-package coverage to console
        if ($srcPkgStmts.Count -gt 0) {
            Write-Host ""
            Write-Host "  === Per-Source-Package Coverage ===" -ForegroundColor Cyan
            Write-Host ""
            $sortedSrcPkgs2 = $srcPkgStmts.GetEnumerator() | ForEach-Object {
                $pctVal2 = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
                [pscustomobject]@{ Name = $_.Key; Pct = $pctVal2 }
            } | Sort-Object Pct -Descending
            foreach ($entry2 in $sortedSrcPkgs2) {
                $color = if ($entry2.Pct -ge 50) { "Green" } elseif ($entry2.Pct -ge 20) { "Yellow" } else { "Red" }
                Write-Host "  $($entry2.Pct)%`t$($entry2.Name)" -ForegroundColor $color
            }
        }

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

        # Auto-open HTML report in browser
        $openHtml = $true
        if ($ExtraArgs -and $ExtraArgs[0] -eq "--no-open") { $openHtml = $false }
        if ($openHtml -and (Test-Path $coverHtml)) {
            Write-Host ""
            Write-Host "  Opening HTML coverage report in browser..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
    Open-FailingTestsIfAny
}

function Invoke-PackageTestCoverage {
    param([string]$pkg)

    if (-not $pkg) {
        Write-Fail "Usage: ./run.ps1 TCP <package-name>"
        Write-Host "  Example: ./run.ps1 TCP regexnewtests" -ForegroundColor Gray
        return
    }

    Write-Header "Running coverage for package: $pkg"
    Invoke-FetchLatest

    # Clean data folder before running tests
    $dataDir = Join-Path $PSScriptRoot "data"
    if (Test-Path $dataDir) {
        Remove-Item -Recurse -Force $dataDir
        Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow
    }

    # Build check from tests/ dir
    Push-Location tests
    try { if (-not (Invoke-BuildCheck "./integratedtests/$pkg/...")) { return } }
    finally { Pop-Location }

    $coverDir = Join-Path $PSScriptRoot "data" "coverage"
    New-Item -ItemType Directory -Path $coverDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage-$pkg.out"
    $coverHtml    = Join-Path $coverDir "coverage-$pkg.html"
    $coverSummary = Join-Path $coverDir "coverage-$pkg-summary.txt"

    # Build coverpkg list: all source packages EXCLUDING tests/
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    # Run from project ROOT so -coverpkg can instrument all source packages
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 "-coverprofile=$coverProfile" "-coverpkg=$covPkgList" "./tests/integratedtests/$pkg/..." 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    $output | ForEach-Object { Write-Host $_ }
    Write-TestLogs $output

    if (Test-Path $coverProfile) {
        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }
        $htmlArgs = @("-html=$coverProfile", "-o=$coverHtml")
        & go tool cover $htmlArgs 2>&1 | Out-Null

        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $summaryLines = [System.Collections.Generic.List[string]]::new()
        $summaryLines.Add("# Coverage Summary ($pkg) — $timestamp")
        $summaryLines.Add("")

        $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
        if ($totalLine) {
            $summaryLines.Add("## Total Coverage")
            $summaryLines.Add("  $totalLine")
            $summaryLines.Add("")
        }

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

        $summaryLines.Add("## Reports")
        $summaryLines.Add("  Profile:  $coverProfile")
        $summaryLines.Add("  HTML:     $coverHtml")
        $summaryLines.Add("  Summary:  $coverSummary")

        Set-Content -Path $coverSummary -Value ($summaryLines -join "`n") -Encoding UTF8

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

        $openHtml = $true
        if ($ExtraArgs -and $ExtraArgs[-1] -eq "--no-open") { $openHtml = $false }
        if ($openHtml -and (Test-Path $coverHtml)) {
            Write-Host ""
            Write-Host "  Opening HTML coverage report..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
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
    Write-Host "  Project Runner — ./run.ps1 <command> [options]" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  Testing:" -ForegroundColor Yellow
    Write-Host "    T   | -t   | test          Run all tests (verbose)"
    Write-Host "    TP  | -tp  | test-pkg      Run tests for a specific package"
    Write-Host "    TC  | -tc  | test-cover    Run tests with coverage (HTML + summary)"
    Write-Host "    TCP | -tcp | test-cover-pkg Run coverage for a specific package"
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
    Write-Host "  Mode Options (for TC/TCP):" -ForegroundColor Yellow
    Write-Host "    --sync      Run precompile + tests sequentially (default: parallel)"
    Write-Host "    --no-open   Skip auto-opening HTML coverage report"
    Write-Host ""
    Write-Host "  Examples:" -ForegroundColor Gray
    Write-Host "    ./run.ps1 T"
    Write-Host "    ./run.ps1 -t"
    Write-Host "    ./run.ps1 TP regexnewtests"
    Write-Host "    ./run.ps1 -tp regexnewtests"
    Write-Host "    ./run.ps1 TCP regexnewtests  (package coverage)"
    Write-Host "    ./run.ps1 TC                 (parallel by default)"
    Write-Host "    ./run.ps1 TC --sync          (sequential mode)"
    Write-Host "    ./run.ps1 TC --sync --no-open"
    Write-Host "    ./run.ps1 -gc"
    Write-Host "    ./run.ps1 -gc 9090          (custom port)"
    Write-Host ""
}

# -- Dispatch --
switch ($Command.ToLower()) {
    { $_ -in "t", "-t", "test" }              { Invoke-AllTests }
    { $_ -in "tp", "-tp", "test-pkg" }        { Invoke-PackageTests $ExtraArgs[0] }
    { $_ -in "tc", "-tc", "test-cover" }      { Invoke-TestCoverage }
    { $_ -in "tcp", "-tcp", "test-cover-pkg" } { Invoke-PackageTestCoverage $ExtraArgs[0] }
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
