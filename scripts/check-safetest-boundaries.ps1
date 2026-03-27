#!/usr/bin/env pwsh

$ErrorActionPreference = "Stop"

$dir = "tests/integratedtests/corestrtests"
$issues = 0

if (-not (Test-Path -LiteralPath $dir -PathType Container)) {
    Write-Host "⚠ Directory $dir not found, skipping check."
    exit 0
}

$repoRoot = (Get-Location).Path
$files = Get-ChildItem -LiteralPath $dir -Filter "*_test.go" -File | Sort-Object Name

foreach ($file in $files) {
    $lines = Get-Content -LiteralPath $file.FullName
    $rel = $file.FullName.Replace($repoRoot, "").TrimStart([char]'\', [char]'/') -replace '\\', '/'

    # Check 1: malformed safeTest boundary (missing inner `}` before `\t})`)
    for ($i = 1; $i -lt $lines.Count; $i++) {
        $curr = $lines[$i]
        if ($curr -eq "`t})" -or $curr -match "^\t\)\}$") {
            $prev = $lines[$i - 1]
            if (
                $prev -match '^\t\t\t' -and
                $prev -notmatch '^\t\t\t\s*\}\s*$' -and
                $prev -notmatch '^\t\t\t\s*//'
            ) {
                Write-Host "  ${rel}:$($i + 1): missing closing } before %})"
                $issues = 1
            }
        }
    }

    # Check 2: closure arg `}` missing `)` (avoid false-positives on if/for/switch/select blocks)
    $controlRe = '^\t\t(?:if|for|switch|select|else)\b'
    $funcOpenRe = '^\t\t[^\t](?!if\b|for\b|switch\b|select\b|else\b).*(?:\bfunc\(|\(func\().*\{\s*$'
    $blockOpenRe = '^\t\t[^\t].*\{\s*$'

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $line = $lines[$i].TrimEnd()
        if ($line -ne "`t`t}") { continue }

        $depth = 0
        $start = [Math]::Max($i - 40, 0)

        for ($j = $i - 1; $j -ge $start; $j--) {
            $l = $lines[$j].TrimEnd()
            if ([string]::IsNullOrWhiteSpace($l)) { continue }

            if ($l -eq "`t`t})") {
                $depth++
                continue
            }

            if ($l.StartsWith("func Test_") -or $l.StartsWith("`tsafeTest(")) {
                break
            }

            if ($l -eq "`t`t}") {
                if ($depth -eq 0) { break }
                continue
            }

            if ($l -match $blockOpenRe) {
                if ($l -match $controlRe) { break }

                if ($l -match $funcOpenRe) {
                    if ($depth -gt 0) {
                        $depth--
                    }
                    else {
                        Write-Host "  ${rel}:$($i + 1): closure } missing )"
                        $issues = 1
                    }
                }
                break
            }
        }
    }
}

# ── Check 3: empty if blocks (comment-only body with no actual statement) ──
foreach ($file in $files) {
    $lines = Get-Content -LiteralPath $file.FullName
    $rel = $file.FullName.Replace($repoRoot, "").TrimStart('\\', '/') -replace '\\', '/'

    for ($i = 0; $i -lt $lines.Count; $i++) {
        if ($lines[$i] -notmatch '^(\t+)if\b.*\{\s*$') { continue }
        $indent = $Matches[1]
        $close = "$indent}"
        $hasStmt = $false
        for ($j = $i + 1; $j -lt [Math]::Min($i + 20, $lines.Count); $j++) {
            $body = $lines[$j]
            $stripped = $body.Trim()
            if ($body.TrimEnd() -eq $close -or $body.TrimEnd() -eq "$close)") { break }
            if ($stripped -eq '' -or $stripped.StartsWith('//')) { continue }
            $hasStmt = $true
            break
        }
        if (-not $hasStmt) {
            Write-Host "  ${rel}:$($i + 1): empty if block (no statements)"
            $issues = 1
        }
    }
}

if ($issues -ne 0) {
    Write-Host ""
    Write-Host "✗ Malformed safeTest boundaries or empty if blocks detected."
    exit 1
}

Write-Host "✓ All safeTest boundaries and if blocks are clean."
exit 0
