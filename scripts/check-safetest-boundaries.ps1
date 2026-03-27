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
    $funcOpenRe = '^\t\t[^\t].*\(func\(.*\{\s*$'
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

    # Check 3: empty if blocks (comment-only body with no actual statement)
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

    # Check 4: func assignment closure must end with `}` (not `})`)
    $assignOpenRe = '^\s*(?:var\s+\w+\s*=\s*func\s*\(|\w+\s*(?::=|=)\s*func\s*\()'
    for ($i = 0; $i -lt $lines.Count; $i++) {
        if ($lines[$i] -notmatch $assignOpenRe) { continue }

        $depth = 0
        $started = $false
        $end = [Math]::Min($i + 300, $lines.Count)

        for ($j = $i; $j -lt $end; $j++) {
            $l = $lines[$j]
            $openCount = ([regex]::Matches($l, '\{')).Count
            $closeCount = ([regex]::Matches($l, '\}')).Count
            if ($openCount -gt 0) { $started = $true }
            $depth += $openCount
            $depth -= $closeCount

            if ($started -and $depth -eq 0) {
                if ($lines[$j].TrimEnd() -match '\}\)\s*$') {
                    Write-Host "  ${rel}:$($j + 1): func assignment closes with }) (should close with } only)"
                    $issues = 1
                }
                break
            }
        }
    }

    # Check 5: placeholder `...` lines are forbidden
    for ($i = 0; $i -lt $lines.Count; $i++) {
        if ($lines[$i].Trim() -eq '...') {
            Write-Host "  ${rel}:$($i + 1): placeholder line ... is not allowed"
            $issues = 1
        }
    }
}

if ($issues -ne 0) {
    Write-Host ""
    Write-Host "✗ Malformed safeTest boundaries, empty if blocks, func assignment closures, or placeholder lines detected."
    exit 1
}

Write-Host "✓ All safeTest boundaries, if blocks, assignment closures, and placeholder lines are clean."
exit 0
