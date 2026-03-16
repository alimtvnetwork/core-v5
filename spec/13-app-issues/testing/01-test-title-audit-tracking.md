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
| coreoncetests | `MapStringStringOnce_testcases.go` | 8 | ✅ Done |
| coreoncetests | `AnyErrorOnce_testcases.go` | 12 | ✅ Done |
| coreoncetests | `BoolOnce_testcases.go` | 8 | ✅ Done |
| coreoncetests | `AnyOnce_testcases.go` | 10 | ✅ Done |
| coreoncetests | `ByteOnce_testcases.go` | 6 | ✅ Done |
| coreoncetests | `BytesErrorOnce_testcases.go` | 34 | ✅ Done |
| coreoncetests | `BytesOnce_testcases.go` | 12 | ✅ Done |
| coreoncetests | `ErrorOnce_testcases.go` | 12 | ✅ Done |
| coreoncetests | `IntegerOnce_testcases.go` | 7 | ✅ Done |
| coreoncetests | `IntegersOnce_testcases.go` | 9 | ✅ Done |
| coreoncetests | `StringOnce_testcases.go` | 13 | ✅ Done |
| coreoncetests | `StringsOnce_testcases.go` | 10 | ✅ Done |
| corejsontests | `Result_IsEqual_testcases.go` | 6 | ✅ Done |
| corejsontests | `Result_Unmarshal_testcases.go` | 4 | ✅ Done |
| coreversiontests | `VersionCompare_testcases.go` | 8 | ✅ Done |
| ostypetests | `OsType_testcases.go` | 6 | ✅ Done |
| reqtypetests | `Request_testcases.go` | 4 | ✅ Done |
| issettertests | `Value_testcases.go` | 6 | ✅ Done |
| stringslicetests | `CloneIf_testcases.go` | 4 | ✅ Done |
| codefuncstests | `GenericWrappers_testcases.go` | 6 | ✅ Done |
| codefuncstests | `GetFuncName_testcases.go` | 3 | ✅ Done |
| codefuncstests | `LegacyWrappers_testcases.go` | 4 | ✅ Done |
| codestacktests | `FileWithLine_testcases.go` | 2 | ✅ Done |
| coreappendtests | `Append_testcases.go` | 2 | ✅ Done |
| enumimpltests | `enumTestCases.go` | 6 | ✅ Done |
| **Total** | **36 files** | **~352** | |

### Skipped (Already Compliant or Excluded)

| Package | File | Reason |
|---------|------|--------|
| corevalidatortests | `LineNumber_testcases.go` | Already compliant |
| corevalidatortests | `LineValidator_testcases.go` | Already compliant |
| corevalidatortests | `Parameter_testcases.go` | Already compliant |
| corevalidatortests | `SimpleSliceValidator_testcases.go` | Already compliant |
| corevalidatortests | `corevalidator_testCases.go` | Snapshot-based — titles are header assertions |
| trydotests | `WrappedErr_testcases.go` | Already compliant |
| trydotests | `WrappedErr_NilReceiver_testcases.go` | Already compliant |
| coreoncetests | `BytesErrorOnce_NilReceiver_testcases.go` | Already compliant |
| coretaskinfotests | All files | Already compliant (verified) |

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

### coreoncetests/MapStringStringOnce_testcases.go (8 titles)

| # | New Title |
|---|-----------|
| 1 | `MapStringStringOnce returns length 2 and hasAnyItem true -- {a:1,b:2} input` |
| 2 | `MapStringStringOnce returns length 0 and isEmpty true -- nil input` |
| 3 | `MapStringStringOnce returns length 0 and isEmpty true -- empty map input` |
| 4 | `MapStringStringOnce.Value returns cached result -- initializer runs once` |
| 5 | `MapStringStringOnce.MarshalJSON returns no error -- {x:y} input` |
| 6 | `MapStringStringOnce.Serialize returns no error -- {s:v} input` |
| 7 | `NewMapStringStringOnce returns correct length -- {k:v} input` |
| 8 | `NewMapStringStringOncePtr returns correct length -- {k:v} input` |

### coreoncetests/AnyErrorOnce_testcases.go (12 titles)

| # | New Title |
|---|-----------|
| 1 | `AnyErrorOnce returns isDefined true and isNull false -- 'hello' value nil error` |
| 2 | `AnyErrorOnce returns isNull true and isEmpty true -- nil value nil error` |
| 3 | `AnyErrorOnce returns hasError true and isFailed true -- nil value error set` |
| 4 | `AnyErrorOnce.Value returns cached result -- initializer runs once` |
| 5 | `AnyErrorOnce.ValueMust returns value without panic -- no error` |
| 6 | `AnyErrorOnce.ValueMust returns panic -- error set` |
| 7 | `AnyErrorOnce.CastValueString returns 'cast' and castSuccess true -- string value` |
| 8 | `AnyErrorOnce.Serialize returns bytes without error -- 'json' value` |
| 9 | `AnyErrorOnce.Serialize returns error -- error set` |
| 10 | `NewAnyErrorOnce returns isNull false and noError true -- value input` |

### coreoncetests/BoolOnce_testcases.go (8 titles)

| # | New Title |
|---|-----------|
| 1 | `BoolOnce returns value true and string 'true' -- true input` |
| 2 | `BoolOnce returns value false and isEmpty true -- false input` |
| 3 | `BoolOnce.Value returns cached result -- initializer runs once` |
| 4 | `BoolOnce.MarshalJSON returns 'true' -- true input` |
| 5 | `BoolOnce.Serialize returns 'true' -- true input` |
| 6 | `NewBoolOnce returns correct value -- true input` |
| 7 | `NewBoolOncePtr returns correct value -- false input` |

### coreoncetests/AnyOnce_testcases.go (10 titles)

| # | New Title |
|---|-----------|
| 1 | `AnyOnce returns isDefined true and isNull false -- 'hello' input` |
| 2 | `AnyOnce returns isNull true and isEmpty true -- nil input` |
| 3 | `AnyOnce returns isDefined true -- int 42 input` |
| 4 | `AnyOnce.Value returns cached result -- initializer runs once` |
| 5 | `AnyOnce.CastString returns 'cast' and castSuccess true -- string value` |
| 6 | `AnyOnce.CastString returns castSuccess false -- int value` |
| 7 | `AnyOnce.MarshalJSON returns no error -- 'json' input` |
| 8 | `AnyOnce.Serialize returns no error -- 'ser' input` |
| 9 | `NewAnyOnce returns correct value -- 'val' input` |
| 10 | `NewAnyOncePtr returns correct value -- 'ptr' input` |

### coreoncetests/ByteOnce_testcases.go (6 titles)

| # | New Title |
|---|-----------|
| 1 | `ByteOnce returns value 42 and isPositive true -- input 42` |
| 2 | `ByteOnce returns isZero true and isEmpty true -- input 0` |
| 3 | `ByteOnce returns isPositive true -- input 255 max byte value` |
| 4 | `ByteOnce.MarshalJSON returns '99' -- input 99` |
| 5 | `ByteOnce.Serialize returns '77' -- input 77` |
| 6 | `NewByteOnce returns correct value -- input 5` |

### coreoncetests/BytesErrorOnce_testcases.go (34 titles)

| # | New Title |
|---|-----------|
| 1 | `BytesErrorOnce returns length 3 and isDefined true -- 'abc' input` |
| 2 | `BytesErrorOnce returns isEmpty true and isDefined false -- nil bytes nil error` |
| 3 | `BytesErrorOnce returns isNull false and isDefined true -- empty bytes nil error` |
| 4 | `BytesErrorOnce.Value returns cached error -- error 'test error'` |
| 5 | `BytesErrorOnce.Execute returns same result as Value -- 'exec' input` |
| 6 | `BytesErrorOnce.ValueOnly returns bytes without error -- 'only' input` |
| 7 | `BytesErrorOnce.ValueWithError returns same as Value -- 'vwe' input` |
| 8 | `BytesErrorOnce returns hasError true and isFailed true -- error 'fail'` |
| 9 | `BytesErrorOnce returns isValid true and isSuccess true -- 'ok' bytes no error` |
| 10 | `BytesErrorOnce returns hasIssuesOrEmpty true -- 'data' bytes with error` |
| 11 | `BytesErrorOnce returns hasIssuesOrEmpty true -- empty bytes no error` |
| 12 | `BytesErrorOnce returns hasSafeItems true -- 'ok' bytes no error` |
| 13 | `BytesErrorOnce.String returns 'str-val' -- 'str-val' bytes input` |
| 14 | `BytesErrorOnce.String returns empty and isStringEmpty true -- nil bytes` |
| 15 | `BytesErrorOnce returns isStringEmptyOrWhitespace true -- whitespace bytes` |
| 16 | `BytesErrorOnce.Deserialize returns no error -- valid JSON input` |
| 17 | `BytesErrorOnce.Deserialize returns error -- source has error` |
| 18 | `BytesErrorOnce.Deserialize returns error -- invalid JSON input` |
| 19 | `BytesErrorOnce.DeserializeMust returns value without panic -- valid JSON input` |
| 20 | `BytesErrorOnce.DeserializeMust returns panic -- source has error` |
| 21 | `BytesErrorOnce.MarshalJSON returns bytes -- '{"a":1}' input` |
| 22 | `BytesErrorOnce.Serialize returns bytes -- 'ser' input` |
| 23 | `BytesErrorOnce.HandleError returns no panic -- 'ok' bytes no error` |
| 24 | `BytesErrorOnce.HandleError returns panic -- error 'handle-err'` |
| 25 | `BytesErrorOnce.MustHaveSafeItems returns panic -- empty bytes` |
| 26 | `BytesErrorOnce.IsInitialized returns false before and true after Value call -- 'x' input` |
| 27 | `BytesErrorOnce.SerializeMust returns bytes without panic -- 'must-ser' input` |
| 28 | `BytesErrorOnce.SerializeMust returns panic -- error 'ser-fail'` |
| 29 | `NewBytesErrorOnce returns correct value -- 'val' input` |

### coreoncetests/BytesOnce_testcases.go (12 titles)

| # | New Title |
|---|-----------|
| 1 | `BytesOnce returns length 5 and isEmpty false -- 'hello' input` |
| 2 | `BytesOnce returns isEmpty true and isNil true -- nil input` |
| 3 | `BytesOnce returns isEmpty true and isNil false -- empty bytes input` |
| 4 | `BytesOnce returns isEmpty true -- nil initializer` |
| 5 | `BytesOnce.String returns 'test-string' -- 'test-string' input` |
| 6 | `BytesOnce returns isEmpty false and length 1 -- 'x' input` |
| 7 | `BytesOnce.Execute returns same result as Value -- 'data' input` |
| 8 | `BytesOnce.MarshalJSON returns data without error -- 'hello' input` |
| 9 | `BytesOnce.UnmarshalJSON returns 'replaced' -- overrides 'original' input` |
| 10 | `BytesOnce.Serialize returns JSON bytes -- 'serialize-me' input` |
| 11 | `NewBytesOnce returns correct value -- 'val' input` |

### coreoncetests/ErrorOnce_testcases.go (12 titles)

| # | New Title |
|---|-----------|
| 1 | `ErrorOnce returns hasError true and message 'fail' -- error 'fail'` |
| 2 | `ErrorOnce returns isValid true and message empty -- nil error` |
| 3 | `ErrorOnce returns isNullOrEmpty true -- nil error` |
| 4 | `ErrorOnce returns isNullOrEmpty true -- empty string error` |
| 5 | `ErrorOnce returns isNullOrEmpty false -- error 'msg'` |
| 6 | `ErrorOnce.IsMessageEqual returns true for 'match' and false for 'other' -- error 'match'` |
| 7 | `ErrorOnce.IsMessageEqual returns false for all -- nil error` |
| 8 | `ErrorOnce.ConcatNewString returns string containing 'base' and 'extra' -- error 'base'` |
| 9 | `ErrorOnce.ConcatNewString returns only additional message -- nil error` |
| 10 | `ErrorOnce.MarshalJSON returns '"marshal"' -- error 'marshal'` |
| 11 | `ErrorOnce.MarshalJSON returns '""' -- nil error` |

### coreoncetests/IntegerOnce_testcases.go (7 titles)

| # | New Title |
|---|-----------|
| 1 | `IntegerOnce returns isZero true and isEmpty true -- input 0` |
| 2 | `IntegerOnce returns isPositive true and isZero false -- input 42` |
| 3 | `IntegerOnce returns isNegative true and isZero false -- input -3` |
| 4 | `IntegerOnce returns isAbove true and isAboveEqual true -- input 10 compare 5` |
| 5 | `IntegerOnce returns isLessThan true and isLessThanEqual true -- input 3 compare 5` |
| 6 | `IntegerOnce.MarshalJSON returns '42' -- input 42` |

### coreoncetests/IntegersOnce_testcases.go (9 titles)

| # | New Title |
|---|-----------|
| 1 | `IntegersOnce returns length 3 and isEmpty false -- [3,1,2] input` |
| 2 | `IntegersOnce returns length 0 and isEmpty true -- empty input` |
| 3 | `IntegersOnce returns length 0 and isEmpty true -- nil input` |
| 4 | `IntegersOnce.Sorted returns [1,2,3] -- [3,1,2] input` |
| 5 | `IntegersOnce returns rangesMapLength 3 -- [10,20,30] input` |
| 6 | `IntegersOnce returns rangesMapLength 0 -- empty input` |
| 7 | `IntegersOnce.IsEqual returns true for same and false for different -- [1,2,3] input` |
| 8 | `IntegersOnce.MarshalJSON returns '[1,2]' -- [1,2] input` |
| 9 | `NewIntegersOnce returns correct length -- [1,2,3] input` |

### coreoncetests/StringOnce_testcases.go (13 titles)

| # | New Title |
|---|-----------|
| 1 | `StringOnce returns value 'hello' and isEmpty false -- 'hello' input` |
| 2 | `StringOnce returns isEmpty true and isEmptyOrWhitespace true -- empty input` |
| 3 | `StringOnce returns isEmpty false and isEmptyOrWhitespace true -- whitespace input` |
| 4 | `StringOnce.IsEqual returns true for 'abc' and false for 'xyz' -- 'abc' input` |
| 5 | `StringOnce.IsContains returns true for 'world' and false for 'xyz' -- 'hello world' input` |
| 6 | `StringOnce.HasPrefix returns true for 'prefix' and false for 'data' -- 'prefix-data' input` |
| 7 | `StringOnce.HasSuffix returns true for 'suffix' and false for 'data' -- 'data-suffix' input` |
| 8 | `StringOnce.SplitBy returns 3 parts -- 'a,b,c' split by ','` |
| 9 | `StringOnce.SplitLeftRight returns 'key' and 'value' -- 'key=value' split by '='` |
| 10 | `StringOnce.SplitLeftRight returns full left and empty right -- 'nosplit' no separator found` |
| 11 | `StringOnce.SplitLeftRightTrim returns trimmed 'key' and 'value' -- ' key = value ' split by '='` |
| 12 | `StringOnce.MarshalJSON returns '"json"' -- 'json' input` |

### coreoncetests/StringsOnce_testcases.go (10 titles)

| # | New Title |
|---|-----------|
| 1 | `StringsOnce returns length 3 and isEmpty false -- [a,b,c] input` |
| 2 | `StringsOnce returns length 0 and isEmpty true -- empty input` |
| 3 | `StringsOnce returns length 0 and isEmpty true -- nil input` |
| 4 | `StringsOnce returns hasX true and hasAllXY true -- [x,y,z] input` |
| 5 | `StringsOnce.Sorted returns [a,b,c] -- [c,a,b] input` |
| 6 | `StringsOnce returns uniqueLen 2 and rangesMapLen 2 -- [a,b,a] input` |
| 7 | `StringsOnce.IsEqual returns true for same and false for different -- [a,b] input` |
| 8 | `StringsOnce.MarshalJSON returns '["a","b"]' -- [a,b] input` |
| 9 | `NewStringsOnce returns correct length -- [x,y] input` |

### corejsontests/Result_IsEqual_testcases.go (6 titles)

| # | New Title |
|---|-----------|
| 1–6 | Renamed to `Result.IsEqual returns {true/false} -- {context}` pattern |

### corejsontests/Result_Unmarshal_testcases.go (4 titles)

| # | New Title |
|---|-----------|
| 1–4 | Renamed to `Result.Unmarshal returns {result} -- {context}` pattern |

### coreversiontests/VersionCompare_testcases.go (8 titles)

| # | New Title |
|---|-----------|
| 1–8 | Renamed to `VersionCompare returns {result} -- {context}` pattern |

### ostypetests/OsType_testcases.go (6 titles)

| # | New Title |
|---|-----------|
| 1–6 | Renamed to `OsType returns {result} -- {context}` pattern |

### reqtypetests/Request_testcases.go (4 titles)

| # | New Title |
|---|-----------|
| 1–4 | Renamed to `Request returns {result} -- {context}` pattern |

### issettertests/Value_testcases.go (6 titles)

| # | New Title |
|---|-----------|
| 1–6 | Renamed to `IsSetter.Value returns {result} -- {context}` pattern |

### stringslicetests/CloneIf_testcases.go (4 titles)

| # | New Title |
|---|-----------|
| 1–4 | Renamed to `CloneIf returns {result} -- {context}` pattern |

### codefuncstests/GenericWrappers_testcases.go (6 titles)

| # | New Title |
|---|-----------|
| 1–6 | Renamed to `GenericWrapper returns {result} -- {context}` pattern |

### codefuncstests/GetFuncName_testcases.go (3 titles)

| # | New Title |
|---|-----------|
| 1–3 | Renamed to `GetFuncName returns {result} -- {context}` pattern |

### codefuncstests/LegacyWrappers_testcases.go (4 titles)

| # | New Title |
|---|-----------|
| 1–4 | Renamed to `LegacyWrapper returns {result} -- {context}` pattern |

### codestacktests/FileWithLine_testcases.go (2 titles)

| # | New Title |
|---|-----------|
| 1 | `FileWithLine returns correct path and lineNumber -- '/src/main.go' line 42` |
| 2 | `FileWithLine returns empty path and lineNumber '0' -- empty file path` |

### coreappendtests/Append_testcases.go (2 titles)

| # | New Title |
|---|-----------|
| 1 | `PrependAppendAnyItemsToStringsSkipOnNil returns 3 items -- prepend and append provided` |
| 2 | `PrependAppendAnyItemsToStringsSkipOnNil returns 2 items -- nil prepend` |

### enumimpltests/enumTestCases.go (6 titles)

| # | New Title |
|---|-----------|
| 1 | `EnumByte returns min 0 and max 10 -- DynamicMap input` |
| 2 | `EnumInt8 returns min -2 and max 12 -- DynamicMap input` |
| 3 | `EnumInt16 returns min -3 and max 14 -- DynamicMap input` |
| 4 | `EnumInt32 returns min -4 and max 15 -- DynamicMap input` |
| 5 | `EnumUInt16 returns min 0 and max 20 -- DynamicMap input` |
| 6 | `EnumString returns min empty and max 'Something2' -- DynamicMap input lexicographic` |

---

## Batch 5 — Completed 2026-03-16

The following 5 packages were audited and renamed:

| Package | Files Renamed | Titles Renamed | Status |
|---------|:------------:|:--------------:|--------|
| corepayloadtests | 2 | 4 | ✅ Done |
| corecomparatortests | 2 | ~30 | ✅ Done |
| corecmptests | 2 | ~25 | ✅ Done |
| converterstests | 1 | ~18 | ✅ Done |
| regexnewtests | 5 | ~50 | ✅ Done |

---

## Pending Packages (Not Yet Audited)

All originally listed pending packages have been audited. Additional packages may have non-compliant titles but were not in the original pending list. To discover more, run:
```bash
grep -rn 'Title:' tests/integratedtests/*/testcases*.go | grep -v -E '"[A-Z][a-zA-Z0-9_.]+ returns '
```

## Related Docs

- [Test Case Naming Convention](/.lovable/memory/testing/test-case-naming-convention)
- [Test Audit Plan](/.lovable/memory/workflow/01-test-audit-plan.md)
- [Testing Patterns](/spec/01-app/13-testing-patterns.md)
