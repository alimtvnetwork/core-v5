# Issue: Migrate regexnewtests from ExpectedLines to MapGherkins

## Status: Open

## Summary

Test cases in `tests/integratedtests/regexnewtests/` use opaque `ExpectedLines: []string{"true", "false", ...}` 
where each line's meaning is unknowable without reading the test runner. These should be migrated to 
`MapGherkins` (or updated to use `Expected: args.Map{}`) with semantic keys.

## Problem

```go
// ❌ Current — opaque, unreadable
ExpectedLines: []string{
    "hello",   // pattern? input? output?
    "true",    // isDefined? isApplicable? isMatch?
    "true",    // ???
    "true",    // ???
    "false",   // ???
},
```

## Solution

Use `MapGherkins` with `Input` for arrange data and `Expected` for assertion data:

```go
// ✅ Target — self-documenting
Input: args.Map{
    "pattern":      "hello",
    "compareInput": "hello world",
},
Expected: args.Map{
    "isDefined":    true,
    "isApplicable": true,
    "isMatch":      true,
    "isFailedMatch": false,
},
```

## Affected Files

### regexnewtests (5 files)

| File | Test Case Slices | Estimated Cases |
|------|-----------------|-----------------|
| `LazyRegex_testcases.go` | `lazyRegexNewTestCases` (5), `lazyRegexLockTestCases` (2), 5 named cases | 12 |
| `LazyRegex_Methods_testcases.go` | `lazyRegexCompileTestCases` (3), `lazyRegexHasErrorTestCases` (2), `lazyRegexMatchBytesTestCases` (2), `lazyRegexMatchErrorTestCases` (3), `lazyRegexStringTestCases` (2) | 12 |
| `Create_testcases.go` | `createTestCases`, `isMatchLockTestCases`, `isMatchFailedTestCases`, 4 named cases | ~15 |
| `CreateMust_testcases.go` | `createMustTestCases`, `createMustLockTestCases`, `hasErrorTestCases`, `isApplicableTestCases`, `createAndMatchTestCases`, `matchErrorTestCases`, `matchUsingFuncErrorTestCases` | ~25 |
| `IsMatchLock_testcases.go` | `isMatchLockTestCases`, `isMatchFailedTestCases`, `lazyIsMatchTestCases`, `lazyIsMatchBuiltinTestCases`, `lazyIsFailedMatchTestCases`, `patternStringTestCases`, `lazyMatchErrorTestCases` | ~15 |

### Corresponding test runners (5 files)

| File | Updates Needed |
|------|---------------|
| `LazyRegex_New_test.go` | Replace `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap`, extract from `tc.Input` |
| `LazyRegex_Methods_test.go` | Same pattern |
| `LazyRegex_PatternMatch_test.go` | Same pattern |
| `Create_test.go` | Same pattern |
| `CreateMust_test.go` | Same pattern |
| `IsMatchLock_test.go` | Same pattern |

## Migration Steps Per File

1. Change test case type from `StringBoolGherkins` / `StringGherkins` → `MapGherkins`
2. Move `Input` (pattern) and `ExtraArgs["compareInput"]` into `Input: args.Map{}`
3. Replace `ExpectedLines: []string{...}` with `Expected: args.Map{...}` using semantic keys
4. Update test runner to extract from `tc.Input.GetAsString("pattern")` etc.
5. Replace `ShouldBeEqualArgs` / `ShouldBeEqualUsingExpected` → `ShouldBeEqualMap`

## Total: ~79 test cases across 5 testcase files + 6 test runner files
