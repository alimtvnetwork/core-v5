# Test Title Audit — Tracking Document

## Convention

All test case `Title` fields must follow the self-documenting format:

```
"{Function} returns {Result} -- {Input Context}"
```

This ensures diagnostic output immediately identifies the function under test and the scenario that failed.

## Audit Status

| Package | File | Titles Renamed | Status |
|---------|------|:--------------:|--------|
| corestrtests | `LeftRightFromSplit_testcases.go` | 14 | ✅ Done |
| corestrtests | `BugfixRegression_testcases.go` | 30 | ✅ Done |
| coreflecttests | `FuncWrap_testcases.go` | 7 | ✅ Done |
| corefuncstests | `corefuncs_testcases.go` | 8 | ✅ Done |
| corevalidatortests | `Condition_testcases.go` | 9 | ✅ Done |
| corevalidatortests | `TextValidator_testcases.go` | 22 | ✅ Done |
| corevalidatortests | `TextValidators_testcases.go` | 14 | ✅ Done |
| corevalidatortests | `HeaderSliceValidators_testcases.go` | 16 | ✅ Done |
| corevalidatortests | `SliceValidatorUnit_testcases.go` | 21 | ✅ Done |
| corevalidatortests | `RangeSegmentsValidator_testcases.go` | 14 | ✅ Done |
| corevalidatortests | `BaseValidatorCoreCondition_testcases.go` | 2 | ✅ Done |
| **Total** | **11 files** | **~157** | |

### Skipped (Already Compliant or Excluded)

| Package | File | Reason |
|---------|------|--------|
| corevalidatortests | `LineNumber_testcases.go` | Already compliant |
| corevalidatortests | `LineValidator_testcases.go` | Already compliant |
| corevalidatortests | `Parameter_testcases.go` | Already compliant |
| corevalidatortests | `SimpleSliceValidator_testcases.go` | Already compliant |
| corevalidatortests | `corevalidator_testCases.go` | Snapshot-based — titles are header assertions |

---

## Renamed Titles by File

### corestrtests/LeftRightFromSplit_testcases.go (14 titles)

| # | New Title |
|---|-----------|
| 1 | `LeftRightFromSplit returns valid split -- 'key=value' normal input` |
| 2 | `LeftRightFromSplit returns invalid -- no separator found` |
| 3 | `LeftRightFromSplit returns empty invalid -- empty input string` |
| 4 | `LeftRightFromSplit returns empty left -- separator at start` |
| 5 | `LeftRightFromSplit returns empty right -- separator at end` |
| 6 | `LeftRightFromSplit returns first-left and last-right -- multiple separators` |
| 7 | `LeftRightFromSplitTrimmed returns trimmed parts -- whitespace around both` |
| 8 | `LeftRightFromSplitTrimmed returns trimmed invalid -- no separator found` |
| 9 | `LeftRightFromSplitTrimmed returns empty parts -- whitespace-only values` |
| 10 | `LeftRightFromSplitFull returns remainder in right -- 'a:b:c:d' multi-separator` |
| 11 | `LeftRightFromSplitFull returns same as normal -- single separator` |
| 12 | `LeftRightFromSplitFull returns invalid -- no separator found` |
| 13 | `LeftRightFromSplitFullTrimmed returns trimmed remainder -- ' a : b : c : d ' with spaces` |
| 14 | `LeftRightFromSplitFullTrimmed returns trimmed invalid -- no separator found` |

### corestrtests/BugfixRegression_testcases.go (30 titles)

| # | New Title |
|---|-----------|
| 1 | `AddNonEmpty returns length 1 -- non-empty string added` |
| 2 | `AddNonEmpty returns length 0 -- empty string skipped` |
| 3 | `AddNonEmpty returns length 3 -- chained three items` |
| 4 | `InsertAt returns shifted items -- middle index insertion` |
| 5 | `InsertAt returns prepended item -- index 0` |
| 6 | `InsertAt returns appended item -- end index` |
| 7 | `InsertAt returns unchanged slice -- negative index` |
| 8 | `InsertAt returns unchanged slice -- out-of-bounds index` |
| 9 | `RemoveAt returns true -- valid middle index` |
| 10 | `RemoveAt returns true -- index 0` |
| 11 | `RemoveAt returns true -- last index` |
| 12 | `RemoveAt returns false -- negative index` |
| 13 | `RemoveAt returns false -- out-of-bounds index` |
| 14 | `RemoveAt returns false -- empty collection` |
| 15 | `IsEqualPtr returns true -- same keys same values` |
| 16 | `IsEqualPtr returns false -- same keys different values` |
| 17 | `IsEqualPtr returns false -- different keys` |
| 18 | `IsEqualPtr returns true -- both empty` |
| 19 | `IsEqualPtr returns false -- nil vs non-nil` |
| 20 | `Hashset returns isEmpty true length 0 -- fresh instance` |
| 21 | `Hashset returns isEmpty false length 2 -- after Add` |
| 22 | `Hashmap returns isEmpty true length 0 -- fresh instance` |
| 23 | `Hashmap returns isEmpty false length 2 -- after Set` |
| 24 | `Skip returns empty -- count beyond length` |
| 25 | `Take returns all items -- count beyond length` |
| 26 | `Skip returns all items -- count 0` |
| 27 | `Take returns empty -- count 0` |
| 28 | `SimpleSlice.HasIndex returns false -- negative index` |
| 29 | `Collection.HasIndex returns false -- negative index` |
| 30 | `Clear returns nil -- nil Hashmap receiver` |

(Plus: `Clear returns empty hashmap`, `Clear returns chainable instance`, `AddBool returns false existed`, `AddBool returns true existed`, `AddBool returns length 3`, `AddOrUpdateCollection returns length 0/2`)

### coreflecttests/FuncWrap_testcases.go (7 titles)

| # | New Title |
|---|-----------|
| 1 | `FuncWrap returns correct output -- someFunctionV1 with valid params` |
| 2 | `FuncWrap returns args count mismatch error -- someFunctionV1 with nil third param` |
| 3 | `FuncWrap returns type mismatch error -- someFunctionV1 with int instead of string at arg 2` |
| 4 | `FuncWrap returns invalid error -- nil work func given` |
| 5 | `FuncWrap returns invalid error -- int given as work func` |
| 6 | `FuncWrap returns string and error output -- someFunctionV2 with valid params` |
| 7 | `FuncWrap returns int, string, error output -- someFunctionV3 with valid params` |

### corefuncstests/corefuncs_testcases.go (8 titles)

| # | New Title |
|---|-----------|
| 1 | `GetFuncName returns short name -- named function input` |
| 2 | `ActionReturnsErrorFuncWrapper.Exec returns nil -- successful action` |
| 3 | `ActionReturnsErrorFuncWrapper.Exec returns error -- failing action` |
| 4 | `IsSuccessFuncWrapper.Exec returns true -- action returns true` |
| 5 | `IsSuccessFuncWrapper.Exec returns false -- action returns false` |
| 6 | `InOutErrFuncWrapperOf.Exec returns output 5 -- string 'hello' input` |
| 7 | `InOutErrFuncWrapperOf.Exec returns error -- empty string input` |
| 8 | `New.ActionErr returns named wrapper -- 'my-action' factory` |

### corevalidatortests/Condition_testcases.go (9 titles)

| # | New Title |
|---|-----------|
| 1 | `IsSplitByWhitespace returns false -- all flags false` |
| 2 | `IsSplitByWhitespace returns true -- IsUniqueWordOnly enabled` |
| 3 | `IsSplitByWhitespace returns true -- IsNonEmptyWhitespace enabled` |
| 4 | `IsSplitByWhitespace returns true -- IsSortStringsBySpace enabled` |
| 5 | `IsSplitByWhitespace returns false -- IsTrimCompare only` |
| 6 | `DefaultDisabled returns isSplit false -- preset disabled` |
| 7 | `DefaultTrim returns isSplit false, isTrimCompare true -- preset trim` |
| 8 | `DefaultSortTrim returns isSplit true -- preset sort-trim` |
| 9 | `DefaultUniqueWords returns isSplit true, isUniqueWordOnly true -- preset unique-words` |

### corevalidatortests/TextValidator_testcases.go (22 titles)

| # | New Title |
|---|-----------|
| 1 | `IsMatch returns true -- exact equal text` |
| 2 | `IsMatch returns false -- different text` |
| 3 | `IsMatch returns true -- case-insensitive match` |
| 4 | `IsMatch returns false -- case-sensitive mismatch` |
| 5 | `IsMatch returns true -- trimmed search matches content` |
| 6 | `IsMatch returns true -- trim applied to both search and content` |
| 7 | `IsMatch returns true -- contains substring found` |
| 8 | `IsMatch returns false -- contains substring not found` |
| 9 | `IsMatch returns true -- NotEqual with different text` |
| 10 | `IsMatch returns false -- NotEqual with same text` |
| 11 | `IsMatch returns true -- unique+sorted reordered words` |
| 12 | `IsMatch returns true -- both search and content empty` |
| 13 | `IsMatch returns false -- empty search vs non-empty content` |
| 14 | `IsMatchMany returns true -- all lines identical` |
| 15 | `IsMatchMany returns false -- one line mismatched` |
| 16 | `IsMatchMany returns true -- empty contents with skip` |
| 17 | `VerifyDetailError returns nil -- matching text` |
| 18 | `VerifyDetailError returns error -- mismatched text` |
| 19 | `VerifyMany returns first error -- firstOnly mode` |
| 20 | `VerifyMany returns all errors -- collect mode` |
| 21 | `VerifyFirstError returns nil -- empty contents with skip` |
| 22 | `SearchTextFinalized returns cached trimmed value -- 'hello' with whitespace` |

### corevalidatortests/TextValidators_testcases.go (14 titles)

| # | New Title |
|---|-----------|
| 1 | `TextValidators returns isEmpty true length 0 -- new instance` |
| 2 | `TextValidators.Add returns length 2 -- two items added` |
| 3 | `TextValidators.Adds returns length 2 -- variadic two items` |
| 4 | `TextValidators.Adds returns length 0 -- no items given` |
| 5 | `TextValidators.AddSimple returns length 1 -- one item added` |
| 6 | `TextValidators.HasIndex returns true for 0, false for 1 -- single item` |
| 7 | `TextValidators.LastIndex returns 1 -- two items` |
| 8 | `TextValidators.IsMatch returns true -- empty validators` |
| 9 | `TextValidators.IsMatch returns true -- all validators pass` |
| 10 | `TextValidators.IsMatch returns false -- one validator fails` |
| 11 | `TextValidators.IsMatchMany returns true -- empty validators` |
| 12 | `TextValidators.VerifyFirstError returns nil -- all match` |
| 13 | `TextValidators.VerifyFirstError returns error -- one mismatch` |
| 14 | `TextValidators.Dispose returns nil Items -- after dispose` |

### corevalidatortests/HeaderSliceValidators_testcases.go (16 titles)

| # | New Title |
|---|-----------|
| 1 | `HeaderSliceValidators.Length returns 0 -- nil input` |
| 2 | `HeaderSliceValidators.Length returns 0 -- empty slice` |
| 3 | `HeaderSliceValidators.Length returns 1 -- single item` |
| 4 | `HeaderSliceValidators.Length returns 2 -- two items` |
| 5 | `HeaderSliceValidators.IsEmpty returns true -- nil input` |
| 6 | `HeaderSliceValidators.IsEmpty returns true -- empty slice` |
| 7 | `HeaderSliceValidators.IsEmpty returns false -- non-empty slice` |
| 8 | `HeaderSliceValidators.IsMatch returns true -- nil input` |
| 9 | `HeaderSliceValidators.IsMatch returns true -- empty slice` |
| 10 | `HeaderSliceValidators.IsMatch returns true -- all matching` |
| 11 | `HeaderSliceValidators.IsMatch returns false -- one mismatch` |
| 12 | `HeaderSliceValidators.VerifyAll returns nil -- empty slice` |
| 13 | `HeaderSliceValidators.VerifyAll returns nil -- all matching` |
| 14 | `HeaderSliceValidators.VerifyAll returns error -- mismatch found` |
| 15 | `HeaderSliceValidators.VerifyFirst returns nil -- empty slice` |
| 16 | `HeaderSliceValidators.VerifyFirst returns error -- mismatch found` |

### corevalidatortests/SliceValidatorUnit_testcases.go (21 titles)

| # | New Title |
|---|-----------|
| 1 | `SliceValidator.IsValid returns true -- exact match` |
| 2 | `SliceValidator.IsValid returns false -- content mismatch` |
| 3 | `SliceValidator.IsValid returns false -- length mismatch` |
| 4 | `SliceValidator.IsValid returns true -- both nil` |
| 5 | `SliceValidator.IsValid returns false -- one nil` |
| 6 | `SliceValidator.IsValid returns true -- both empty` |
| 7 | `SliceValidator.IsValid returns true -- trimmed match` |
| 8 | `SliceValidator.IsValid returns true -- contains substrings` |
| 9 | `SliceValidator.ActualLinesLength returns 2 -- two actual lines` |
| 10 | `SliceValidator.ExpectingLinesLength returns 3 -- three expected lines` |
| 11 | `SliceValidator.IsUsedAlready returns false -- fresh instance` |
| 12 | `SliceValidator.IsUsedAlready returns true -- after ComparingValidators` |
| 13 | `SliceValidator.MethodName returns 'IsContains' -- Contains compare mode` |
| 14 | `SliceValidator.SetActual returns length 1 -- one line set` |
| 15 | `SliceValidator.SetActualVsExpected returns both set -- one actual one expected` |
| 16 | `SliceValidator.IsValidOtherLines returns true -- matching lines` |
| 17 | `SliceValidator.IsValidOtherLines returns false -- mismatching lines` |
| 18 | `SliceValidator.AllVerifyError returns nil -- matching lines` |
| 19 | `SliceValidator.AllVerifyError returns error -- mismatched lines` |
| 20 | `SliceValidator.AllVerifyError returns nil -- skip when actual empty` |
| 21 | `SliceValidator.Dispose returns nil lines -- after dispose` |

### corevalidatortests/RangeSegmentsValidator_testcases.go (14 titles)

| # | New Title |
|---|-----------|
| 1 | `LengthOfVerifierSegments returns 0 -- no segments` |
| 2 | `LengthOfVerifierSegments returns 1 -- one segment` |
| 3 | `LengthOfVerifierSegments returns 2 -- two segments` |
| 4 | `Validators returns HeaderSliceValidators -- one segment input` |
| 5 | `VerifyAll returns nil -- matching segment` |
| 6 | `VerifyAll returns error -- mismatching segment` |
| 7 | `VerifySimple returns nil -- matching segment range 1-3` |
| 8 | `VerifySimple returns error -- mismatched segment range 0-2` |
| 9 | `VerifyFirst returns nil -- matching segment range 0-2` |
| 10 | `VerifyFirst returns error -- mismatched segment range 0-2` |
| 11 | `VerifyUpto returns nil -- matching segment within length` |
| 12 | `VerifyUpto returns error -- mismatched segment range 0-2` |
| 13 | `SetActual returns self and match true -- matching segment propagated` |
| 14 | `SetActualOnAll returns match true -- all segments matching` |

### corevalidatortests/BaseValidatorCoreCondition_testcases.go (2 titles)

| # | New Title |
|---|-----------|
| 1 | `ValidatorCoreConditionDefault returns all-false condition -- nil preset` |
| 2 | `ValidatorCoreConditionDefault returns existing condition -- non-nil preset` |

---

## Pending Packages (Not Yet Audited)

The following test packages have `_testcases.go` files that have **not been audited** for title compliance:

| Package | Est. Files | Notes |
|---------|:----------:|-------|
| coreoncetests | 3+ | BytesErrorOnce, ErrorOnce, etc. |
| corejsontests | 5+ | Result, JsonParsing, etc. |
| corepayloadtests | 2+ | Attributes, etc. |
| coreversiontests | 3+ | Comparison, ComparisonExtended |
| ostypetests | 1 | OsType |
| reqtypetests | 1 | Request |
| issettertests | 1 | Value |
| stringslicetests | 2+ | CloneIf, AnyItemsCloneIf |
| trydotests | 3+ | WrappedErr, NilReceiver |
| codefuncstests | 1 | LegacyWrappers |
| codestacktests | 2+ | |
| coreappendtests | 1+ | |
| corecomparatortests | 2+ | |
| corecmptests | 1+ | |
| converterstests | 2+ | |
| enumimpltests | 2+ | |
| regexnewtests | 1+ | |
| coretaskinfotests | 1 | Already compliant (verified) |

## Related Docs

- [Test Case Naming Convention](/.lovable/memory/testing/test-case-naming-convention)
- [Test Audit Plan](/.lovable/memory/workflow/01-test-audit-plan.md)
- [Testing Patterns](/spec/01-app/13-testing-patterns.md)
