# Issue: chmodhelper tests fail due to stack traces in error messages

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Tests

- `Test_VerifyRwxChmodUsingRwxInstructions_Unix`
- `Test_VerifyRwxPartialChmodLocations_Unix`

## Root Cause

The `errcore` error types (`PathsMissingOrHavingIssuesType`, `NotSupportedType`)
now include stack trace information in their `.Error()` output. The actual error
messages contain lines like:

```
"\"[/tmp/core/test-cases-3s /tmp/core/test-cases-3x]\" Path Ref(s) access having issues! missing or other { }"
"- /Users/.../chmodhelper/GetExistsFilteredPathFileInfoMap.go:40"
"- /Users/.../chmodhelper/RwxInstructionExecutor.go:135"
"- ErrorRefOnly"
"Stack-Trace:"
```

But the test expectations only contain the error message text (1 line), not the
stack trace lines (6+ lines). The `assertNonWhiteSortedEqual` function compares
line counts first, causing `"ActualLines, ExpectedLines Length is not equal"`.

## Solution

Added `isStackTraceLine()` helper to `nonWhiteSortedEqual.go` that filters out:
- Lines starting with `"- /"` (file paths)
- Lines equal to `"Stack-Trace:"`
- Lines starting with `"- ErrorRefOnly"` or `"- getVerifyRwxInternalError"`

The `nonWhiteSortedLines()` function now skips these lines before normalization
and comparison, so tests compare only the semantic error content.

## Affected Files

- `tests/integratedtests/chmodhelpertests/nonWhiteSortedEqual.go` — test infrastructure fix
