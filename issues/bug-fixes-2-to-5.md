# Bug Fixes #2–#5: Complete Analysis & Code Changes

## Summary of All Bugs

| # | Package | Failing Tests | Root Cause |
|---|---------|--------------|------------|
| 2 | chmodhelper | `PathsUsingRwxFull_ContinueOnError` (nil panic) | Missing `return nil` causes fall-through; also missing nil guard in PathsUsingRwxFull |
| 3 | chmodhelper | 11 Cov4 verifier tests + Cov5 + Cov8 | Verifier tests expect mismatch but get match (path comparison issue on macOS) |
| 4 | chmodhelper | `RwxPartial_Error`, `CreateDirFiles*`, `LinuxApplyRecursiveOnPath_Unix` | Missing nil guard in `RwxPartial`; `/temp` hardcoded (read-only on macOS) |
| 5 | corestr | `KeyValueCollection` unmarshal (6 tests), `LinkedCollections` ParseInject (2 tests), `SimpleSlice_SkipDynamic`, `AddNonEmptyStrings`, `LeftRightTrimmedUsingSlice_Single`, `SimpleStringOnce_Int16_OutOfRange` | UnmarshalJSON doesn't handle struct-wrapped JSON; SimpleSlice type assertion |

---

## Bug #2: `chmodVerifier.PathsUsingRwxFull` nil pointer crash

**File:** `chmodhelper/chmodVerifier.go`  
**Test:** `Test_Cov13_ChmodVerifier_PathsUsingRwxFull_ContinueOnError`  
**Error:** `panic: runtime error: invalid memory address or nil pointer dereference` at line 412

### Root Cause
When `isContinueOnError=false`, the loop at lines 396–404 iterates all locations. If no error is found, execution **falls through** to the `isContinueOnError=true` code block (line 406+), re-iterating unnecessarily. Missing `return nil` after line 404.

### Fix
**Replace lines 386–417** with:

```go
func (it chmodVerifier) PathsUsingRwxFull(
	isContinueOnError bool,
	expectedHyphenedRwx string,
	locations ...string,
) error {
	if locations == nil || len(locations) == 0 {
		return errcore.CannotBeNilOrEmptyType.
			Error(constants.EmptyString, nil)
	}

	if !isContinueOnError {
		for _, location := range locations {
			err := it.RwxFull(location, expectedHyphenedRwx)

			if err != nil {
				return err
			}
		}

		return nil // <-- FIX: was missing, caused fall-through
	}

	slice := corestr.New.Collection.Cap(constants.Zero)

	for _, location := range locations {
		err := it.RwxFull(location, expectedHyphenedRwx)

		if err != nil {
			slice.Add(err.Error())
		}
	}

	return errcore.SliceErrorDefault(slice.List())
}
```

**Key change:** Added `return nil` after the `!isContinueOnError` loop (after the closing `}` of line 403).

---

## Bug #3: chmodVerifier Cov4 tests — verifier returns `true` when it should return `false`

**File:** `chmodhelper/chmodVerifier.go`  
**Tests:** `Test_Cov4_ChmodVerifier_IsEqual_ValidPath`, `Test_Cov4_ChmodVerifier_IsEqualRwxFull`, `Test_Cov4_ChmodVerifier_Path`, `Test_Cov4_ChmodVerifier_PathIf_True`, `Test_Cov4_ChmodVerifier_PathsUsingFileMode`, `Test_Cov4_ChmodVerifier_PathsUsingFileModeContinueOnError`, `Test_Cov4_ChmodVerifier_PathsUsingFileModeImmediateReturn`, `Test_Cov4_FwChmodVerifier_IsEqualFile`, `Test_Cov5_RwxWrapper_VerifyPaths`, `Test_Cov8_NotDirError_ExistsButNotDir`

### Root Cause
These tests create a file/dir with one permission and then verify against a **different** permission, expecting a mismatch error. But the verifier returns "match" (`true` / `nil`).

The issue is in the `RwxFull` comparison at lines 273-276:
```go
existingFileMode := fileInfo.Mode().String()[1:]
if existingFileMode == expectedHyphenedRwx[1:] {
    return nil
}
```

On **macOS**, `os.FileMode.String()` can return a 10-char string like `"-rwxr-xr-x"`, but it can also include special bits (setuid/setgid/sticky) making it longer. More importantly, `fileInfo.Mode()` on macOS includes the file type bits. For a **directory**, `Mode().String()` returns `"drwxr-xr-x"` (starts with `d`), and for a file `"-rw-r--r--"` (starts with `-`).

The comparison at line 274 strips the first character from both `existingFileMode` (the full mode string) and `expectedHyphenedRwx` (which starts with `-`). But if the existing mode string is `"-rw-r--r--"` and expected is `"-rwx------"`, these should NOT match. 

**The real issue is likely that the test creates a file with `os.FileMode(0644)` but verifies against `0644` as well** — or the file is being created with a mode that the umask adjusts to match. 

### How to diagnose
Since I don't have the test file (`Coverage4_test.go`), please share:
1. The test setup code that creates the file
2. What file mode is set during creation
3. What expected mode is used for verification

**If the test creates a temp file and expects specific permissions, macOS umask (typically `022`) will mask the file mode.** For example, creating with `0777` results in `0755` due to umask. If the test then verifies against `0755`, it matches when the test expects mismatch.

### Possible Fix (if the issue is test environment)
The tests may need to explicitly `os.Chmod(path, desiredMode)` AFTER creation instead of relying on `os.Create/MkdirAll` modes. Share the test file and I'll give the exact fix.

### Possible Fix (if the issue is in RwxFull comparison)
If the comparison logic is genuinely wrong, add permission-bits-only comparison:

```go
// In RwxFull, replace lines 273-276 with:
existingPerm := fileInfo.Mode().Perm()
expectedPerm := fileMode // need to convert expectedHyphenedRwx back to FileMode
// or compare the 9-char rwx strings after normalizing
existingRwx := fileInfo.Mode().Perm().String()[1:] // strip type char
expectedRwx := expectedHyphenedRwx[1:]             // strip type char
if existingRwx == expectedRwx {
    return nil
}
```

---

## Bug #4: chmodApplier.RwxPartial nil guard + path issues

### Bug 4a: `Test_Cov13_ChmodApplier_RwxPartial_Error`
**File:** `chmodhelper/chmodApplier.go`  
**Error:** `expected error for nil condition`

The test passes `nil` condition to `RwxPartial` and expects an error. Looking at the code (lines 257-276):

```go
func (it chmodApplier) RwxPartial(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
		rwxPartial,
		condition)
	...
```

**Root Cause:** When `condition` is nil AND `locations` is not empty, it passes nil condition to `RwxPartialToInstructionExecutor` which may not guard against nil. But the test passes nil condition WITH locations, expecting an error return.

**Fix:** Add nil guard before the `RwxPartialToInstructionExecutor` call:

```go
func (it chmodApplier) RwxPartial(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
		rwxPartial,
		condition)

	if err != nil {
		return err
	}

	return rwxInstructionExecutor.
		ApplyOnPathsPtr(&locations)
}
```

### Bug 4b: `Test_Cov13_CreateDirFilesWithRwxPermission_FileModeErr` & `CreateDirFilesWithRwxPermissions_Error` & `CreateDirFilesWithRwxPermissionsMust_Panic`
**File:** `chmodhelper/CreateDirWithFiles.go` or `CreateDirFilesWithRwxPermission.go`  
**Error:** `expected error` / `expected panic`

These tests pass invalid parameters to the CreateDirFiles functions and expect errors/panics. The functions likely need validation guards. Without seeing these files, the fix would be:

```go
// In CreateDirFilesWithRwxPermission (or CreateDirWithFiles):
// Add validation at the top:
if len(locations) == 0 {
    return errcore.CannotBeNilOrEmptyType.ErrorNoRefs("locations")
}
```

**Please share** `CreateDirWithFiles.go` and `CreateDirFilesWithRwxPermission.go` for exact fix.

### Bug 4c: `Test_LinuxApplyRecursiveOnPath_Unix`
**File:** `chmodhelpertests/LinuxApplyRecursiveOnPath_test.go`  
**Error:** `mkdir /temp: read-only file system`

**Root Cause:** The test uses hardcoded `/temp` (root-level, read-only on macOS) instead of `os.TempDir()` which returns `/tmp` or `/var/folders/...`.

**Fix** (in the test file):
```go
// Replace hardcoded "/temp" path:
// BEFORE:
basePath := "/temp/core/test-cases"

// AFTER:
basePath := filepath.Join(os.TempDir(), "core", "test-cases")
```

---

## Bug #5: corestr JSON unmarshal + type assertion issues

### Bug 5a: `KeyValueCollection.UnmarshalJSON` fails for struct-wrapped JSON (6 tests)
**File:** `coredata/corestr/KeyValueCollection.go`  
**Tests:** `Test_C27_KeyValueCollection_ParseInjectUsingJson`, `Test_C27_KeyValueCollection_JsonParseSelfInject`, `Test_C27_KeyValueCollection_Deserialize`, `Test_C28_97_*`, `Test_C28_100_*`, `Test_C28_106_*`

**Root Cause:** `MarshalJSON()` serializes as `[{"Key":"k","Value":"v"}]` (array), but `corejson.New(&it)` (used in `Json()` and `JsonPtr()`) serializes the struct directly as `{"KeyValuePairs":[{"Key":"k","Value":"v"}]}` (struct wrapper). When `ParseInjectUsingJson` tries to unmarshal, it calls `UnmarshalJSON` which only handles the array format.

**Fix — Replace lines 467-481:**

```go
func (it *KeyValueCollection) UnmarshalJSON(data []byte) error {
	// Try compact array format first: [{"Key":"k","Value":"v"}]
	var dataModelItems []KeyValuePair
	err := corejson.Deserialize.UsingBytes(
		data,
		&dataModelItems,
	)

	if err == nil {
		if len(dataModelItems) > 0 {
			it.KeyValuePairs = dataModelItems
		} else {
			it.KeyValuePairs = []KeyValuePair{}
		}
		return nil
	}

	// Try struct-wrapped format: {"KeyValuePairs":[{"Key":"k","Value":"v"}]}
	type kvAlias KeyValueCollection
	var wrapper kvAlias
	err2 := json.Unmarshal(data, &wrapper)

	if err2 == nil {
		if len(wrapper.KeyValuePairs) > 0 {
			it.KeyValuePairs = wrapper.KeyValuePairs
		} else {
			it.KeyValuePairs = []KeyValuePair{}
		}
		return nil
	}

	return err
}
```

**Key change:** Added fallback to unmarshal as struct wrapper using a type alias (to avoid infinite recursion on `UnmarshalJSON`).

### Bug 5b: `LinkedCollections.ParseInjectUsingJson` fails (2 tests)
**File:** `coredata/corestr/LinkedCollections.go`  
**Tests:** `Test_C25_LinkedCollections_ParseInjectUsingJson`, `Test_C25_LinkedCollections_ParseInjectUsingJsonMust`

**Root Cause:** Similar to KeyValueCollection — `corejson.New(&it)` produces a JSON format that `UnmarshalJSON` doesn't handle. But LinkedCollections has no exported fields, so `corejson.New` should fall back to `MarshalJSON` which produces `[]string`.

The issue is likely in `corejson.Result.Unmarshal` — it may not be calling `json.Unmarshal` properly on the stored bytes. The `Unmarshal` method at `Result.go:535` (from the stack trace for KeyValueCollection) may have a bug.

**Without seeing `corejson/Result.go`**, the safest fix is to make `ParseInjectUsingJson` handle both formats:

```go
func (it *LinkedCollections) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*LinkedCollections, error) {
	// Get raw bytes and try direct json.Unmarshal
	rawBytes := jsonResult.RawMust()
	err := json.Unmarshal(rawBytes, it)

	if err != nil {
		// Fallback: try standard Unmarshal
		err = jsonResult.Unmarshal(it)
		if err != nil {
			return nil, err
		}
	}

	return it, nil
}
```

**Alternative (if corejson.Result doesn't have RawMust):** Share `corejson/Result.go` so I can fix the root cause.

### Bug 5c: `SimpleSlice.SkipDynamic` type assertion panic
**File:** `coredata/corestr/SimpleSlice.go`  
**Test:** `Test_C29_44_SimpleSlice_SkipDynamic`  
**Error:** `interface conversion: interface {} is corestr.SimpleSlice, not []string`

**Root Cause:** A method returns `SimpleSlice` as `interface{}` but the caller asserts it as `[]string`. `SimpleSlice` is defined as `type SimpleSlice []string`, so it IS a `[]string` underneath but Go's type system treats them as distinct named types.

**Fix:** In `SimpleSlice.go`, find the method that returns `any`/`interface{}` (likely `JsonModelAny()` or `SkipDynamic()`) and return `[]string(it)` instead of `it`:

```go
// Find the method returning SimpleSlice as interface{} and change:
// BEFORE:
func (it SimpleSlice) SomeMethod() any {
    return it  // returns SimpleSlice
}

// AFTER:
func (it SimpleSlice) SomeMethod() any {
    return []string(it)  // returns []string
}
```

**Please share `SimpleSlice.go`** so I can identify the exact method.

### Bug 5d: `Test_Cov21_Collection_AddNonEmptyStrings` — expected 3, got 2
**File:** `coredata/corestr/Collection.go`  
**Test:** `Test_Cov21_Collection_AddNonEmptyStrings`

**Root Cause:** `AddNonEmptyStrings` or `AddNonEmptyStringsSlice` is filtering out a valid string. Per memory note: *"`AddNonEmptyStringsSlice` explicitly filters out empty strings."* The test adds 3 non-empty strings but only 2 are added.

This is likely a bug in the empty-string check logic — it may be trimming strings before checking, or using `len(s) <= 1` instead of `len(s) == 0`.

**Please share** the `AddNonEmptyStrings` / `AddNonEmptyStringsSlice` method from `Collection.go`.

### Bug 5e: `Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single` — expected `a`
Per memory note: *"`LeftRightTrimmedUsingSlice` trims single items (length=1 case) as well as the 2-item case."* The fix should ensure single-element slices are trimmed.

**Please share** the `LeftRightTrimmedUsingSlice` method.

### Bug 5f: `Test_C28_149_SimpleStringOnce_Int16_OutOfRange` — expected 0
This test expects 0 for an out-of-range Int16 conversion but gets a non-zero value. Without seeing `SimpleStringOnce`, the fix would be adding range validation.

**Please share** `SimpleStringOnce.go`.

---

## Files I Need to Complete All Fixes

To fix bugs #3, #4b, #5c, #5d, #5e, #5f, I need these files:

1. `chmodhelpertests/Coverage4_test.go` (lines around 441, 492, 503, 514) — to understand test setup for bug #3
2. `chmodhelper/CreateDirWithFiles.go` — for bug #4b
3. `chmodhelper/CreateDirFilesWithRwxPermission.go` — for bug #4b
4. `chmodhelpertests/LinuxApplyRecursiveOnPath_test.go` — for bug #4c
5. `coredata/corestr/SimpleSlice.go` — for bug #5c
6. `coredata/corestr/Collection.go` (AddNonEmptyStrings method) — for bug #5d
7. `internal/strutilinternal/LeftRightTrimmed.go` or wherever `LeftRightTrimmedUsingSlice` lives — for bug #5e
8. `coredata/coreonce/SimpleStringOnce.go` — for bug #5f
9. `coredata/corejson/Result.go` (around line 535) — root cause of unmarshal issues

---

## Fixes You Can Apply Right Now (No Additional Files Needed)

### Fix 1: `chmodVerifier.go` — Add `return nil` (Bug #2)
**Line 404**, after the closing `}` of `if !isContinueOnError`, add `return nil`:

```diff
 	if !isContinueOnError {
 		for _, location := range locations {
 			err := it.RwxFull(location, expectedHyphenedRwx)
 
 			if err != nil {
 				return err
 			}
 		}
+
+		return nil
 	}
```

### Fix 2: `chmodApplier.go` — Add nil guard for condition (Bug #4a)
**Line 262**, add nil check before `RwxPartialToInstructionExecutor`:

```diff
 func (it chmodApplier) RwxPartial(
 	rwxPartial string,
 	condition *chmodins.Condition,
 	locations ...string,
 ) error {
 	if len(locations) == 0 {
 		return nil
 	}
 
+	if condition == nil {
+		return errcore.CannotBeNilOrEmptyType.
+			ErrorNoRefs("condition")
+	}
+
 	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
```

### Fix 3: `KeyValueCollection.go` — Handle struct-wrapped JSON (Bug #5a)
**Replace lines 467–481** with the dual-format UnmarshalJSON shown above in Bug #5a section.
