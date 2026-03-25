# Issue: run.ps1 PC — Cannot index into a null array

## Date: 2026-03-25

## Status: ✅ FIXED

## Symptom

Running:
```powershell
./run.ps1 PC
```
failed with:
```text
run.ps1: Cannot index into a null array.
```

## Root Cause

In the command dispatch block, `PC` (and a few other commands) directly accessed `$ExtraArgs[0]` even when no extra argument was provided.

```powershell
{ $_ -in "pc", "-pc", "pre-commit" } { Invoke-PreCommitCheck $ExtraArgs[0] }
```

When `$ExtraArgs` is `$null`, indexing `[0]` throws `Cannot index into a null array`.

## Fix

Use a safe first-argument variable and pass that instead:

```powershell
$firstExtraArg = if ($ExtraArgs -and $ExtraArgs.Count -gt 0) { $ExtraArgs[0] } else { $null }
```

Then dispatch with `$firstExtraArg` for `TP`, `TCP`, and `PC`.

## Affected File

- `run.ps1` — dispatch section

## Learning

For optional CLI arguments in PowerShell, never index directly into arrays without a null/length check.
