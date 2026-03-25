# Current Failing Tests Analysis

## Date: 2026-03-25

## Overview

Two test runs were analyzed (failing-tests-43 with 5 failures, failing-tests-44 with 25 failures). The bugs fall into **5 categories** across `chmodhelpertests` and `corestrtests`.

---

## Bug #1: `dirCreator.ByChecking` always returns error even on success

**Failing tests:** `Test_Cov8_NotDirError_ExistsButNotDir`  
**File:** `chmodhelper/dirCreator.go`, lines 86-133

### Root Cause

`ByChecking()` calls `newError.chmodApplyFailed()` unconditionally at line 128 — even when `chmodErr` is `nil`. It should only return an error when `chmodErr != nil`.

Also at line 94-100: when the dir already exists and IS a directory, it calls `os.Chmod` but then passes the result through `newError.chmodApplyFailed()` which wraps it as an error even on success.

### Fix

```go
// Line 93-100: When dir exists, apply chmod and return directly
if isExists && isDir {
    return os.Chmod(dirPath, applyChmod)
}

// Lines 123-132: After MkdirAll, only return error if chmod actually failed
chmodErr := os.Chmod(dirPath, applyChmod)
if chmodErr != nil {
    return newError.chmodApplyFailed(applyChmod, dirPath, chmodErr)
}
return nil
```

### Info from codebase

The `dirCreator.go` file is already in context (see current code). The bug is clearly visible at lines 94-100 and 128-132.

---

## Bug #2: `chmodVerifier` tests expect errors on non-existent paths but get success

**Failing tests (11):**
- `Test_Cov4_ChmodVerifier_IsEqual_ValidPath`
- `Test_Cov4_ChmodVerifier_IsEqualRwxFull`
- `Test_Cov4_ChmodVerifier_Path`
- `Test_Cov4_ChmodVerifier_PathIf_True`
- `Test_Cov4_ChmodVerifier_PathsUsingFileMode`
- `Test_Cov4_ChmodVerifier_PathsUsingFileModeContinueOnError`
- `Test_Cov4_ChmodVerifier_PathsUsingFileModeImmediateReturn`
- `Test_Cov4_FwChmodVerifier_IsEqualFile`
- `Test_Cov5_RwxWrapper_VerifyPaths`
- `Test_Cov13_ChmodVerifier_PathsUsingRwxFull_ContinueOnError` (nil pointer dereference)

**File:** `chmodhelper/chmodVerifier.go`

### Root Cause

Tests pass non-existent file paths expecting `IsEqual → false` and `Path/PathIf → error`. But the verifier is returning `true`/no-error, meaning it's not properly checking path existence before comparing permissions.

The nil pointer dereference at line 412 in `PathsUsingRwxFull` suggests a nil `PathExistStat` result is being used without a nil check.

### Fix

Need to read `chmodVerifier.go` to see the exact issue. The verifier methods (`IsEqual`, `Path`, `PathIf`, `PathsUsingFileMode`, `PathsUsingRwxFull`) need to:
1. Check if path exists before comparing permissions
2. Return `false`/error for non-existent paths
3. Add nil guard before dereferencing `PathExistStat` result

### Info needed

I need to see `chmodhelper/chmodVerifier.go` and the test file `Coverage4_test.go` to understand the exact test inputs and expected behavior.

---

## Bug #3: `chmodApplier.RwxPartial` returns nil for nil condition

**Failing test:** `Test_Cov13_ChmodApplier_RwxPartial_Error`  
**File:** `chmodhelper/chmodApplier.go`

### Root Cause

Test expects an error when a nil condition is passed, but `RwxPartial` returns `nil` instead.

### Fix

Need to read `chmodApplier.go` to see the `RwxPartial` method — it likely needs a nil guard that returns an error.

---

## Bug #4: `CreateDirFilesWithRwxPermission` / `CreateDirWithFiles` path handling

**Failing tests:**
- `Test_Cov13_CreateDirFilesWithRwxPermission_FileModeErr` — "expected error"
- `Test_Cov13_CreateDirFilesWithRwxPermissions_Error` — "expected error"  
- `Test_Cov13_CreateDirFilesWithRwxPermissionsMust_Panic` — "expected panic"
- `Test_LinuxApplyRecursiveOnPath_Unix` — panic: `mkdir /temp: read-only file system`

### Root Cause

Two issues:
1. `CreateDirFilesWithRwxPermission` is succeeding when it should fail (invalid file mode test)
2. `Test_LinuxApplyRecursiveOnPath_Unix` uses `/temp` (root-level) instead of `os.TempDir()` (`/tmp`), which is read-only on macOS

### Fix

- For the error tests: Need to read `CreateDirWithFiles.go` and `CreateDirFilesWithRwxPermission.go` to find missing validation
- For `Test_LinuxApplyRecursiveOnPath_Unix`: The test should use `os.TempDir()` or `t.TempDir()` instead of hardcoded `/temp`

---

## Bug #5: `corestr` — JSON unmarshal failures and type mismatches

**Failing tests:**
- `Test_C25_LinkedCollections_ParseInjectUsingJson` / `ParseInjectUsingJsonMust` — unmarshal fails for `LinkedCollections`
- `Test_C27_KeyValueCollection_ParseInjectUsingJson` / `JsonParseSelfInject` / `Deserialize` — unmarshal fails for `KeyValueCollection`
- `Test_C28_97/100/106_KeyValueCollection_*` — same unmarshal failures
- `Test_C28_149_SimpleStringOnce_Int16_OutOfRange` — "expected 0"
- `Test_C29_44_SimpleSlice_SkipDynamic` — type assertion panic: `interface {} is corestr.SimpleSlice, not []string`
- `Test_Cov21_Collection_AddNonEmptyStrings` — "expected 3, got 2"
- `Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single` — "expected a"

### Root Cause

Multiple issues:

1. **KeyValueCollection/LinkedCollections UnmarshalJSON broken** (`KeyValueCollection.go:494`, `LinkedCollections.go:1478`): The `UnmarshalJSON` method fails even with valid JSON payloads like `{"KeyValuePairs":[{"Key":"k","Value":"v"}]}`. Likely a field name mismatch or the unmarshal target struct doesn't match the JSON model.

2. **SimpleSlice.SkipDynamic** returns `corestr.SimpleSlice` but the test casts to `[]string`: The `SkipDynamic` method needs to return `[]string` or the test needs to handle the `SimpleSlice` type.

3. **Collection.AddNonEmptyStrings** adds 2 items instead of 3: Empty string filtering is too aggressive (possibly trimming strings it shouldn't, or the method signature changed).

4. **LeftRight.LeftRightTrimmedUsingSlice** returns wrong value for single-element slice.

5. **SimpleStringOnce.Int16** out-of-range handling returns wrong default.

### Fix

Need to read:
- `coredata/corestr/KeyValueCollection.go` (around line 494) — check `UnmarshalJSON`
- `coredata/corestr/LinkedCollections.go` (around line 1478) — check `UnmarshalJSON`  
- `coredata/corestr/SimpleSlice.go` — check `SkipDynamic` return type
- `coredata/corestr/Collection.go` — check `AddNonEmptyStrings`

---

## Summary Table

| # | Category | Tests | File(s) to Fix | Info Available? |
|---|----------|-------|-----------------|-----------------|
| 1 | `ByChecking` unconditional error | 1 | `dirCreator.go` | ✅ Full code in context |
| 2 | chmodVerifier missing path-exists check | 11 | `chmodVerifier.go` | ❌ Need to read file |
| 3 | `RwxPartial` nil guard | 1 | `chmodApplier.go` | ❌ Need to read file |
| 4 | CreateDirFiles error handling + `/temp` path | 4 | `CreateDirWithFiles.go`, test file | ❌ Need to read files |
| 5 | corestr JSON unmarshal + type issues | 8 | `KeyValueCollection.go`, `LinkedCollections.go`, etc. | ❌ Need to read files |
| **Total** | | **25** | | |
