#!/usr/bin/env pwsh

param(
    [string]$SinglePackage = ""
)

$ErrorActionPreference = "Stop"

function Get-BraceDelta([string]$line) {
    $openCount = ([regex]::Matches($line, '\{')).Count
    $closeCount = ([regex]::Matches($line, '\}')).Count
    return $openCount - $closeCount
}

function Add-Issue {
    param(
        [System.Collections.Generic.List[object]]$issues,
        [System.Collections.Generic.HashSet[string]]$issueKeys,
        [string]$packageName,
        [string]$relFile,
        [int]$line,
        [string]$category,
        [string]$message,
        [string]$snippet
    )

    $key = "$relFile|$line|$category|$message"
    if ($issueKeys.Contains($key)) { return }
    $issueKeys.Add($key) | Out-Null

    $issues.Add([pscustomobject]@{
            package  = $packageName
            file     = $relFile
            line     = $line
            category = $category
            message  = $message
            snippet  = $snippet.Trim()
        })
}

function Get-CoreTestCasesAliases([string]$rawFileContent) {
    $aliases = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    $importMatches = [regex]::Matches(
        $rawFileContent,
        '(?m)^\s*(?:(?<alias>[A-Za-z_]\w*)\s+)?"github\.com/alimtvnetwork/core/coretests/coretestcases"'
    )

    foreach ($m in $importMatches) {
        $alias = $m.Groups['alias'].Value
        if ([string]::IsNullOrWhiteSpace($alias)) {
            $aliases.Add("coretestcases") | Out-Null
        }
        else {
            $aliases.Add($alias) | Out-Null
        }
    }

    return $aliases.ToArray()
}

function Scan-FileForRegressions {
    param(
        [string]$repoRoot,
        [string]$packageName,
        [System.IO.FileInfo]$file,
        [System.Collections.Generic.List[object]]$issues,
        [System.Collections.Generic.HashSet[string]]$issueKeys
    )

    $raw = Get-Content -Path $file.FullName -Raw
    $lines = Get-Content -Path $file.FullName
    $relFile = $file.FullName.Replace($repoRoot, '').TrimStart('\\', '/') -replace '\\', '/'

    # Rule 1: Invalid Result initializer field (corejson.Result{Err: ...})
    for ($i = 0; $i -lt $lines.Count; $i++) {
        if ($lines[$i] -match '\bcorejson\.Result\s*\{[^}]*\bErr\s*:') {
            Add-Issue $issues $issueKeys $packageName $relFile ($i + 1) "corejson-result-err" "Use corejson.Result.Error instead of Err" $lines[$i]
        }
    }

    # Rule 2: Invalid Result field access (x.Err) for known corejson.Result vars
    $resultVarNames = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $line = $lines[$i]

        if ($line -match '\bvar\s+([A-Za-z_]\w*)\s+\*?corejson\.Result\b') {
            $resultVarNames.Add($Matches[1]) | Out-Null
        }

        if ($line -match '\b([A-Za-z_]\w*)\s*:?=\s*&?corejson\.Result\b') {
            $resultVarNames.Add($Matches[1]) | Out-Null
        }

        if ($line -match '\b([A-Za-z_]\w*)\s*:=\s*.+\.Json(?:Ptr)?\(\)\s*$') {
            $resultVarNames.Add($Matches[1]) | Out-Null
        }
    }

    foreach ($name in $resultVarNames) {
        $escapedName = [regex]::Escape($name)
        for ($i = 0; $i -lt $lines.Count; $i++) {
            if ($lines[$i] -match "\b$escapedName\.Err\b") {
                Add-Issue $issues $issueKeys $packageName $relFile ($i + 1) "corejson-result-err" "Use .Error instead of .Err on corejson.Result" $lines[$i]
            }
        }
    }

    # Rule 3: Legacy CaseV1 fields inside coretestcases.CaseV1 literals
    $aliases = Get-CoreTestCasesAliases $raw
    if ($aliases.Count -eq 0) { return }

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $line = $lines[$i]

        $isCaseV1Start = $false
        foreach ($alias in $aliases) {
            $escapedAlias = [regex]::Escape($alias)
            if ($line -match "\b$escapedAlias\.CaseV1\s*\{") {
                $isCaseV1Start = $true
                break
            }
        }

        if (-not $isCaseV1Start) { continue }

        if ($line -match '\b(Name|Input|Expected|Actual)\s*:') {
            $legacyFieldInline = $Matches[1]
            Add-Issue $issues $issueKeys $packageName $relFile ($i + 1) "legacy-casev1-field" "Legacy coretestcases.CaseV1 field '$legacyFieldInline' found (use Title/ArrangeInput/ActualInput/ExpectedInput)" $line
        }

        $braceDepth = Get-BraceDelta $line
        if ($braceDepth -le 0) { continue }

        $j = $i + 1
        while ($j -lt $lines.Count -and $braceDepth -gt 0) {
            $currentLine = $lines[$j]

            if ($currentLine -match '\b(Name|Input|Expected|Actual)\s*:') {
                $legacyField = $Matches[1]
                Add-Issue $issues $issueKeys $packageName $relFile ($j + 1) "legacy-casev1-field" "Legacy coretestcases.CaseV1 field '$legacyField' found (use Title/ArrangeInput/ActualInput/ExpectedInput)" $currentLine
            }

            $braceDepth += Get-BraceDelta $currentLine

            $j++
        }
    }
}

$repoRoot = Split-Path $PSScriptRoot -Parent
$integratedRoot = Join-Path $repoRoot "tests" "integratedtests"

if (-not (Test-Path $integratedRoot)) {
    Write-Host "  ✗ Missing integrated tests directory: $integratedRoot" -ForegroundColor Red
    exit 1
}

if ([string]::IsNullOrWhiteSpace($SinglePackage)) {
    $packageDirs = @(Get-ChildItem -Path $integratedRoot -Directory)
}
else {
    $pkgPath = Join-Path $integratedRoot $SinglePackage
    if (-not (Test-Path $pkgPath)) {
        Write-Host "  ✗ Package not found for regression scan: $SinglePackage" -ForegroundColor Red
        exit 1
    }
    $packageDirs = @([System.IO.DirectoryInfo]$pkgPath)
}

$issues = [System.Collections.Generic.List[object]]::new()
$issueKeys = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
$scannedFiles = 0

foreach ($pkg in $packageDirs) {
    $coverageFiles = @(Get-ChildItem -Path $pkg.FullName -Filter "Coverage*.go" -File -ErrorAction SilentlyContinue)
    foreach ($file in $coverageFiles) {
        $scannedFiles++
        Scan-FileForRegressions $repoRoot $pkg.Name $file $issues $issueKeys
    }
}

if ($issues.Count -eq 0) {
    Write-Host "  ✓ Regression guard passed ($scannedFiles Coverage* file(s) scanned)" -ForegroundColor Green
    exit 0
}

$groupedByPackage = $issues | Group-Object package | Sort-Object Name

Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
Write-Host "  │ ✗ REGRESSION GUARD FOUND $($issues.Count) ISSUE(S)" -ForegroundColor Red
Write-Host "  │" -ForegroundColor Red
foreach ($g in $groupedByPackage) {
    Write-Host "  │   ✗ $($g.Name) ($($g.Count) issue(s))" -ForegroundColor Red
}
Write-Host "  │" -ForegroundColor Red
Write-Host "  │ Fix these before running pre-commit compile check." -ForegroundColor Yellow
Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
Write-Host ""

$issues |
Sort-Object package, file, line |
ForEach-Object {
    Write-Host "  $($_.file):$($_.line) [$($_.category)] $($_.message)" -ForegroundColor Yellow
}

exit 1