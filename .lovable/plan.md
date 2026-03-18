

## Fix: Strip `warning: no packages being tested` Lines Everywhere

### Problem
The warning filter exists in only 2 places (inside `Write-TestLogs` at line 106 and the TC console output at line 740), but there are **5 other code paths** that print raw test output to console without filtering and feed unfiltered output to `Write-TestLogs`:

| Line | Command | Issue |
|------|---------|-------|
| 265  | `Invoke-GoTestAndLog` | Unfiltered console print |
| 357  | `Invoke-AllTests` (TA) | Unfiltered console print |
| 389  | `Invoke-PackageTests` (TP) | Unfiltered console print |
| 1231 | `Invoke-PackageCoverage` (PC) | Unfiltered console print |
| 1323 | `Invoke-IntegratedTests` (TI) | Unfiltered console print |

### Solution
Create a shared filter function and apply it at all 7 output points:

1. **Add a reusable filter function** near the top of `run.ps1` (after the helper functions):
   ```powershell
   function Filter-TestWarnings([string[]]$lines) {
       return $lines | Where-Object {
           $_ -notmatch '^\s*warning: no packages being tested depend on matches for pattern'
       }
   }
   ```

2. **Replace all 5 unfiltered console output locations** with filtered versions:
   - Line 265: `$output | ForEach-Object { Write-Host $_ }` → `Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }`
   - Line 357: same pattern
   - Line 389: same pattern
   - Line 1231: same pattern
   - Line 1323: same pattern

3. **Update `Write-TestLogs` (line 106-108)** to use the shared function:
   ```powershell
   $filteredOutput = Filter-TestWarnings $rawOutput
   ```

4. **Update TC console output (line 740)** to use the shared function:
   ```powershell
   Filter-TestWarnings $allOutput | ForEach-Object { Write-Host $_ }
   ```

This ensures warnings are stripped from both console output AND log files across every command (TA, TP, TI, TC, PC).

