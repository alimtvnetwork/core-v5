# Coverage Recovery Spec (Ordered Execution)

## Goal
Bring **chmodhelper**, **coredata/corejson**, and **coredata/corestr** to **100% coverage** while keeping tests stable and unblocked.

## Inputs Used
- `user-uploads://failing-tests-11.txt` (0 runtime failures)
- `user-uploads://blocked-packages-9.txt` (compile blockers in `corejsontests`, `corestrtests`)
- `user-uploads://coverage-summary-11.txt`
- `user-uploads://per-package-coverage-15.json`

## Why previous attempts failed
1. New coverage suites introduced compile blockers (duplicate test names, stale method signatures, missing imports).
2. Blocked test packages were skipped by coverage runner, so coverage stayed very low despite added tests.
3. Some tests targeted outdated APIs (e.g., `AnyTo.UsingSerializerFunc`, old `Collection` helper signatures).

## Ordered Tasks
1. **Unblock packages first**
   - Fix `corejsontests` compile errors:
     - replace stale API usage
     - remove/relabel duplicate test function names
     - add missing imports
   - Fix `corestrtests` compile errors:
     - update tests to current method signatures
     - remove duplicate test declarations by unique test naming
2. **Re-validate baseline**
   - Confirm blocked packages are now compilable/runnable.
3. **Coverage-gap targeting**
   - Use machine-readable coverage profile (`.out`) to map uncovered lines → methods.
   - Generate JSON with only `method` + `uncovered_lines` for targeted test generation.
4. **Targeted test additions**
   - Add positive, negative, branch, and edge-path tests for remaining gaps.
5. **Final validation**
   - Re-run compile and coverage.
   - Verify 100% on requested packages.

## Coverage artifact preference
- **Best for automation and precision:** `.out` coverage profile
- **Best for visual inspection:** HTML report
- For programmatic JSON extraction, `.out` is required/preferred.

## PowerShell feasibility confirmation
Yes, it is possible to generate a JSON report from PowerShell containing only:
- method names
- uncovered lines per method
