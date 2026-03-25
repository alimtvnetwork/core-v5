# Issue: Remaining Test Failures Analysis

## Status: ANALYSIS (documented for follow-up)

### 1. chmodhelper Stack Trace Mismatch (2 tests)
- `Test_VerifyRwxChmodUsingRwxInstructions_Unix`
- `Test_VerifyRwxPartialChmodLocations_Unix`
- **Root Cause**: Error messages include `- ErrorRefOnly` or `- getVerifyRwxInternalError` 
  source reference lines that the expected values don't include. The `isStackTraceLine` 
  filter in `nonWhiteSortedEqual.go` catches full stack traces but production error 
  output now includes a single source-ref line (e.g., `- ErrorRefOnly`) rather than 
  a full stack trace.
- **Fix**: Expected values in test cases need updating to include the source-ref lines,
  OR the filter needs to also strip single `- <functionName>` lines.

### 2. Hashset.AddVariations (1 test)
- `Test_Cov9_Hashset_AddVariations` — expected len 6, got 7
- **Root Cause**: The Transpile fix (building new map instead of delete-during-iteration)
  correctly processes ALL keys, while the old buggy behavior would skip some keys during 
  iteration mutation. The test expected the buggy count.
- **Fix**: Update test expectation to 7, or investigate if the add-variations test 
  exercises a code path where 6 was genuinely correct.

### 3. Hashmap.Clone / JSON roundtrip (2 tests)  
- `Test_Cov6_Hashmap_Clone` — expects sameLen/sameVal=false but gets true
- `Test_Cov6_Collection_JsonString` — map key mismatch
- **Root Cause**: Pre-existing test expectation errors, not regressions.

### 4. CharCollectionMap nil map panic (1 test)
- `Test_CovCCM_01_IsEmpty_HasItems` — `assignment to entry in nil map`
- **Root Cause**: `CharCollectionMap.Add` doesn't check for nil internal map.

### 5. JSON roundtrip failures (2 tests)
- `Test_CovS06_CharCollMap_Json_Verification` — roundTrip: false
- `Test_CovS07_Json_Verification` — roundTrip: false  
- **Root Cause**: ParseInjectUsingJson doesn't properly deserialize these types.

### 6. Cov42_Collection_IsContainsAllSlice_Empty (1 test)
- Pre-existing assertion pattern issue (same as Cov43 pattern).

### 7. Cov44_SSO_IsValueBool / Cov52/57/63/69 (misc)
- Various pre-existing test issues exposed after runtime crash was fixed.
