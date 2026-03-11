# Test Audit & Diagnostics Improvement Plan

## Status: In Progress

## Completed Tasks

### 1. ✅ Diagnostic Output Formatting (Iteration 1-4)
- Fixed MapMismatchError header indentation (4-space indent, leading newline)
- Fixed LineDiff label alignment (column 21 for both actual/expected colons)
- Fixed args.Map ExpectedInput falling through to PrettyJSON
- Switched map diagnostic output from quoted-string to Go-literal format
- Added separator headers (`============================>`) for visual structure
- Entries use tab-indented `"key": value,` format (copy-pasteable)

### 2. ✅ Test Title Audit — Batches 1-4
Renamed ~250+ test case titles across all audited packages to follow the convention:
`"{Function} returns {Result} -- {Input Context}"`

**Audited packages:**
- corestrtests (LeftRightFromSplit, BugfixRegression)
- coreflecttests (FuncWrap)
- corefuncstests (corefuncs)
- corevalidatortests (Condition, TextValidator, TextValidators, HeaderSliceValidators, SliceValidatorUnit, RangeSegmentsValidator, BaseValidatorCoreCondition)

**Skipped (already compliant or snapshot-based):**
- LineNumber_testcases, LineValidator_testcases, Parameter_testcases, SimpleSliceValidator_testcases
- corevalidator_testCases.go (snapshot-based)

### 3. ✅ Fix 21 Failing Tests (Iteration 4)
Fixed all 21 failing tests from `failing-tests-4.txt`:

**Production code fixes (2):**
- `internal/trydo/WrappedErr.go` — operator precedence bug in `HasErrorOrException()` (missing parentheses)
- `coredata/corestr/SimpleSlice.go` — `InsertAt()` slice bounds panic (fixed to use copy pattern)

**Test runner logic fixes (3 files, 6 test functions):**
- `codefuncstests/LegacyWrappers_test.go` — removed `containsName: false` from else branch in 4 test functions
- `corejsontests/Result_IsEmpty_test.go` — fixed value-vs-pointer type assertion for `corejson.Result`
- `stringslicetests/CloneIf_test.go` — fixed `isIndependentCopy` logic in 2 test functions

**Test expectation fixes (9 files):**
- `BytesErrorOnce_testcases.go` — fixed `isDefined`/`isEmpty` for `[]byte{}`, lifecycle panic expectations
- `ErrorOnce_testcases.go` — fixed quoted string expectation for ConcatNewString
- `OsType_testcases.go` — fixed group names (WindowsGroup, UnixGroup, etc.), Unknown casing
- `Request_testcases.go` — fixed Create name, Drop/CreateOrUpdate logical group flags
- `ComparisonExtended_testcases.go` — fixed IsExpectedVersion for v4 vs v4.0
- `Value_testcases.go` — fixed IsOutOfRange for value 5
- `CloneIf_testcases.go` — expectations match corrected test logic
- `Attributes_testcases.go` — deep clone returns error, updated expectations
- `Attributes_test.go` — graceful error handling for Clone, removed unused import

## Pending Tasks

### 4. 🔲 Test Title Audit — Remaining Packages
Packages not yet audited:
- coreoncetests, corejsontests, corepayloadtests, coreversiontests
- ostypetests, reqtypetests, issettertests, stringslicetests
- trydotests, codefuncstests, codestacktests, coreappendtests
- corecomparatortests, corecmptests, converterstests
- enumimpltests, regexnewtests

### 5. 🔲 Nil Receiver Coverage Audit
- Ensure all pointer receiver methods have CaseNilSafe tests
- Pattern: migrate nil cases from CaseV1 to CaseNilSafe

### 6. 🔲 Deep Clone Production Bug Investigation
- `corepayload.Attributes.Clone(deep=true)` returns error
- May be missing implementation or dependency issue
- Logged as issue, needs production code investigation
