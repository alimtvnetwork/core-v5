#!/usr/bin/env pwsh
$ErrorActionPreference = "Stop"

$shScript = Join-Path $PSScriptRoot "check-safetest-boundaries.sh"
if (-not (Test-Path -LiteralPath $shScript -PathType Leaf)) {
    Write-Host "✗ Boundary checker shell script not found: $shScript"
    exit 1
}

$bashCmd = Get-Command bash -ErrorAction SilentlyContinue
if (-not $bashCmd) {
    Write-Host "✗ bash is required to run boundary checks on this machine."
    exit 1
}

& bash $shScript
exit $LASTEXITCODE
