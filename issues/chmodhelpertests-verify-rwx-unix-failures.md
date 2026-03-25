# Issue: chmodhelpertests — VerifyRwx Unix test failures

## Date: 2026-03-25

## Status: 🔍 ANALYSIS

## Failing Tests

- `Test_VerifyRwxChmodUsingRwxInstructions_Unix`
- `Test_VerifyRwxPartialChmodLocations_Unix`

## Test Structure

Both tests:
1. Create directories with `CreateDirFilesWithRwxPermissionsMust(true, pathInstructionsV2)`
2. Iterate over test cases that verify chmod expectations against `SimpleLocations`
3. Assert using `assertNonWhiteSortedEqual` (normalizes whitespace + sorts lines)

### Test data

- `pathInstructionsV2` creates: `/tmp/core/test-cases`, `/tmp/core/test-cases-2`, `/tmp/core/test-cases-3`
- `SimpleLocations` includes non-existent paths: `/tmp/core/test-cases-3s`, `/tmp/core/test-cases-3x`
- Default permissions applied: `rwxr-xr--` (0754)

### Previous fix

The `assertNonWhiteSortedEqual` function had a type assertion panic (`.([]string)` on a `string`).
This was fixed with a type switch in `nonWhiteSortedEqual.go` (see `issues/chmodhelpertests-type-assertion-panic.md`).

## Possible Root Causes

1. **Error message format mismatch**: The production code may produce error messages with slightly different
   formatting than what test cases expect. After `nonWhiteSortedLines` normalization (sort tokens per line,
   sort lines), a single extra/missing word would cause a mismatch.

2. **Filesystem state**: If `/tmp/core/test-cases-*` directories persist from prior runs with different
   permissions, `isRemoveAllDirBeforeCreate=true` should clean them, but edge cases may exist.

3. **macOS SIP restrictions**: macOS may apply additional restrictions on `/tmp` that affect
   permission verification, especially for `0754` vs actual applied mode.

## Investigation Steps

1. Run the failing tests in isolation and capture actual vs expected output:
   ```bash
   go test -v -run "Test_VerifyRwxChmodUsingRwxInstructions_Unix" ./tests/integratedtests/chmodhelpertests/
   ```

2. Compare normalized actual error with normalized expected error.

3. Check if `SimpleLocations` paths match what `pathInstructionsV2` creates.

## Affected Files

- `tests/integratedtests/chmodhelpertests/VerifyRwxChmodUsingRwxInstructions_test.go`
- `tests/integratedtests/chmodhelpertests/VerifyPartialRwxLocations_test.go`
- `tests/integratedtests/chmodhelpertests/nonWhiteSortedEqual.go`
- `tests/integratedtests/chmodhelpertests/VerifyRwxChmodUsingRwxInstructions_testcases.go`
- `tests/integratedtests/chmodhelpertests/VerifyPartialRwxLocations_testcases.go`

## Next Action

Need the **actual error output** from `failing-tests.txt` for these two tests to determine
the exact mismatch.
