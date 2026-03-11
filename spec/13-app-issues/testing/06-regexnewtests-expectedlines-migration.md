# Issue: Migrate regexnewtests from ExpectedLines to MapGherkins

## Status: ✅ Completed

## Summary

Test cases in `tests/integratedtests/regexnewtests/` used opaque `ExpectedLines: []string{"true", "false", ...}` 
where each line's meaning was unknowable without reading the test runner. These have been migrated to 
`MapGherkins` (or updated to use `Expected: args.Map{}`) with semantic keys, and a centralized `params.go` 
file eliminates all raw string key usage.

## Problem

```go
// ❌ Old — opaque, unreadable
ExpectedLines: []string{
    "hello",   // pattern? input? output?
    "true",    // isDefined? isApplicable? isMatch?
    "true",    // ???
    "true",    // ???
    "false",   // ???
},
```

## Solution

Created `params.go` with a typed struct holding all reusable key constants, then migrated all test cases 
to use `MapGherkins` with `Input` for arrange data and `Expected` for assertion data:

```go
// ✅ Completed — self-documenting with params keys
Input: args.Map{
    params.pattern:      "hello",
    params.compareInput: "hello world",
},
Expected: args.Map{
    params.isDefined:    true,
    params.isApplicable: true,
    params.isMatch:      true,
    params.isFailedMatch: false,
},
```

## Migrated Files

### params.go (new)

Centralized key constants file with 30 keys covering input keys (`pattern`, `compareInput`, `isLock`, 
`customizer`, etc.) and expected keys (`isDefined`, `isApplicable`, `isMatch`, `regexNotNil`, `hasError`, etc.).

### Testcase files (5 migrated from ExpectedLines → MapGherkins)

| File | Test Case Slices | Cases |
|------|-----------------|-------|
| `LazyRegex_testcases.go` | `lazyRegexNewTestCases` (5), `lazyRegexLockTestCases` (2), 5 named cases | 12 |
| `LazyRegex_Methods_testcases.go` | `lazyRegexCompileTestCases` (3), `lazyRegexHasErrorTestCases` (2), `lazyRegexMatchBytesTestCases` (2), `lazyRegexMatchErrorTestCases` (3), `lazyRegexStringTestCases` (2) | 12 |
| `Create_testcases.go` | `createTestCases` (4), `createIsMatchLockTestCases` (4), `createIsMatchFailedTestCases` (3), 4 named cases | 15 |
| `CreateMust_testcases.go` | `createMustTestCases` (3), `createMustLockIfTestCases` (2), `createLockIfTestCases` (4), `createApplicableLockTestCases` (3), `newMustLockTestCases` (3), `matchUsingFuncErrorLockTestCases` (3), `matchUsingCustomizeErrorFuncLockTestCases` (4) | 22 |
| `IsMatchLock_testcases.go` | `isMatchLockTestCases` (5), `isMatchFailedTestCases` (3), `isMatchLockLazyIsMatchTestCases` (2), `isMatchLockCompileTestCases` (1), `isMatchLockIsFailedMatchTestCases` (2), `isMatchLockPatternStringTestCases` (1), `isMatchLockMatchErrorTestCases` (2) | 16 |

### Testcase files (2 updated to use params.go keys)

| File | Updates |
|------|---------|
| `LazyRegex_EdgeCases_testcases.go` | Raw string keys → `params.` constants |
| `LazyRegex_ExtendedMethods_testcases.go` | Raw string keys → `params.` constants |

### Test runners (8 updated)

| File | Changes |
|------|---------|
| `LazyRegex_New_test.go` | `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap`, `fmt.Sprintf` removed |
| `LazyRegex_PatternMatch_test.go` | `ShouldBeEqualUsingExpectedFirst` → `ShouldBeEqualMapFirst` |
| `LazyRegex_Methods_test.go` | `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap` |
| `Create_test.go` | `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap` |
| `CreateMust_test.go` | `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap`, branching eliminated |
| `IsMatchLock_test.go` | `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap` |
| `LazyRegex_EdgeCases_test.go` | Raw string keys → `params.` constants |
| `LazyRegex_ExtendedMethods_test.go` | Raw string keys → `params.` constants |

### Unchanged (no migration needed)

| File | Reason |
|------|--------|
| `LazyRegex_NilReceiver_testcases.go` | Uses `CaseNilSafe` with `results.ResultAny` — different pattern |
| `LazyRegex_Concurrency_test.go` | Channel-based goroutine tests — not table-driven |

## Key Improvements

1. **No more `fmt.Sprintf`**: All stringification removed; native types (`bool`, `string`) used throughout.
2. **No branching in runners**: `CustomizeErrorFuncLock` runner now uses consistent map shape instead of conditional `append`.
3. **Zero raw string keys**: Every `args.Map` key uses `params.` constants from centralized `params.go`.
4. **Self-documenting**: Every test case is readable without consulting the test runner.

## Total: 77 test cases across 7 testcase files + 8 test runner files + 1 params.go
