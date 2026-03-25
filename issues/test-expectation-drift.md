# Issue: Test Expectation Drift Across Multiple Packages (FIXED)

## Summary
16 test cases across 9 files had expected values that no longer matched production behavior. These were pre-existing mismatches surfaced during the test title audit.

## Categories of Drift

### 1. Enum/Type Name Changes
- `ostype.Group.Name()` now returns full names with "Group" suffix (e.g., "WindowsGroup" not "Windows")
- `ostype.GetVariant("unknown")` returns "Unknown" (capitalized) not "unknown"
- `reqtype.Create.Name()` returns "CreateUsingAliasMap" not "Create"

### 2. Behavioral Changes
- `issetter.IsOutOfRange(5)` now returns `true` (was expected `false`)
- `reqtype.Drop.IsCrudOnly()` now returns `true`
- `reqtype.CreateOrUpdate` now returns `isCrudOnly: true, isReadOrEdit: true`
- `coreversion.IsExpectedVersion(v4, v4.0, LeftGreater)` returns `true`
- `BytesErrorOnce` with `[]byte{}`: `isDefined` is `true`, `isEmpty` is `false`
- `BytesErrorOnce` lifecycle: error causes all panic-guarded methods to trigger
- `ErrorOnce.ConcatNewString` wraps nil-error result in quotes
- `corepayload.Attributes.Clone(deep=true)` returns error

### 3. Test Logic Bugs
- `codefuncstests`: unconditional `containsName: false` in else branches
- `corejsontests`: value vs pointer type assertion
- `stringslicetests`: `isIndependentCopy` computed incorrectly

## Learning
- Test expectations should be verified against production behavior after any refactoring
- Named type method returns (`.Name()`, `.String()`) should be checked against actual enum definitions
- Boolean logic in test assertions should be kept simple and direct

## What Not To Repeat
- Don't add map keys unconditionally in else branches (creates extra entries)
- Don't assume Go value types match pointer interface assertions
- Don't conflate "returned same reference" with "is independent copy" in test logic
