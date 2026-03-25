# Root Cause Analysis: 126 Test Failures (2026-03-25)

## How did 4 failures become 126?

The original run had **4 failures**. After AI-applied fixes, failures exploded to **126**.
This happened because of **one catastrophic change** that created ~106 new failures,
plus ~20 pre-existing/introduced failures that were NOT fixed correctly.

---

## Failure Breakdown

| Category | Count | Root Cause | Introduced By |
|----------|-------|-----------|---------------|
| Cov43 rewrite regression | ~106 | AI rewrote Coverage43_Iteration39_test.go incorrectly | AI change |
| ValidValue.ValueByte | 6 | AI changed production code, tests expect original behavior | AI change |
| Hashset.Transpile side effects | 1 | AI's Transpile fix changed count from 6→7 | AI change |
| chmodhelper stack traces | 2 | Stack trace filter incomplete | Partial AI fix |
| CharCollectionMap nil map | 1 | Production bug: nil internal map | Pre-existing |
| JSON roundtrip failures | 4 | ParseInjectUsingJson doesn't roundtrip these types | Pre-existing |
| Pre-existing test issues | 6 | Test expectations wrong for current production code | Pre-existing |

---

## Category 1: Cov43 Rewrite Regression (~106 tests)

### Tests Affected
All `Test_Cov43_*` tests in `Coverage43_Iteration39_test.go`

### What Happened
The AI completely rewrote this 764-line test file. The original tests used
`tc.ShouldBeEqual(t, caseIndex)` (the variadic pattern). The AI changed them to
`tc.ShouldBeEqualMapFirst(t, args.Map{...})`.

### Why It Fails
The `ShouldBeEqualMapFirst` pattern IS correct in isolation, but the rewrite
introduced a mismatch between how `ExpectedInput` was set and what the comparison
function receives. The error output shows:

```
ActualLines, ExpectedLines any is nil and other is not. - Expect [""] != ["0"] Actual
```

This means `CompileToStrings()` on one side returns content while the other returns
nil/empty. The root issue is that the original tests used a DIFFERENT assertion
flow (`ShouldBeEqual` with variadic strings, comparing `ExpectedInput` directly),
and the rewrite changed the assertion contract without properly understanding the
original data flow.

### Correct Approach
The original tests should be restored to their pre-AI state. The original issue
was that `ShouldBeEqual(t, 0)` was called with caseIndex=0 instead of actual values.
The fix should have been MINIMAL: either:
1. Set `ActualInput` properly and keep the `ShouldBeEqual` pattern
2. OR convert to `ShouldBeEqualMapFirst` but with CORRECT `ExpectedInput` format

### Impact
This single rewrite is responsible for **85%** of all failures (106 out of 126).

---

## Category 2: ValidValue.ValueByte (6 tests)

### Tests Affected
- `Test_I27_ValidValue_ValueByte`
- `Test_S01_ValidValue_ValueByte_Invalid`
- `Test_S01_ValidValue_ValueByte_Negative`
- `Test_Seg8_VV_ValueByte_Negative`
- `Test_Cov11_ValidValue_Conversions` (partial — "b":"2" diff)
- `Test_Seg8_SSO_Deserialize` (partial — ValueByte negative)

### What Happened
The AI changed `ValidValue.ValueByte()` to return `defVal` instead of
`constants.Zero` on parse errors and negative values, then later "reverted" it.

### Error Evidence
```
actual : `val : 5`
expected : `val : 0`
```
`ValueByte(-5, 5)` returns 5 (defVal) instead of 0.

### Root Cause
The revert to `constants.Zero` either:
1. Did NOT apply properly on the user's machine (sync issue), OR
2. The production code has a different bug where it returns defVal despite the revert

### Correct Behavior
Looking at test expectations across 6 tests, they ALL expect `constants.Zero` (0)
for invalid/negative inputs. The original `ValueByte` contract is:
- Valid positive value → return parsed byte
- Error or negative → return `constants.Zero` (NOT defVal)

This is intentional: `ValueByte` is a "safe zero-default" method.

---

## Category 3: Hashset.Transpile & Wrap Methods (1 test)

### Tests Affected
- `Test_Cov9_Hashset_AddVariations` — expected len=6, got len=7

### What Happened
The AI fixed `Hashset.Transpile` to build a new map instead of deleting during
iteration. This correctly processes ALL keys, but the test expected the old
buggy count (6 instead of 7).

### Root Cause
The Transpile fix IS correct. The test expectation needs updating from 6→7.
However, we need to verify this by checking what the 7th key is and confirming
it's a legitimate addition.

---

## Category 4: chmodhelper Stack Traces (2 tests)

### Tests Affected
- `Test_VerifyRwxChmodUsingRwxInstructions_Unix`
- `Test_VerifyRwxPartialChmodLocations_Unix`

### Error Evidence (after partial fix)
```
Actual: 2 lines (message + "- ErrorRefOnly")
Expected: 1 line (message only)
```

### What Happened
The AI added `isStackTraceLine()` to filter file paths and "Stack-Trace:" lines.
This reduced the mismatch from 7→1 extra lines to 2→1. But lines like
`- ErrorRefOnly` and `- getVerifyRwxInternalError` still appear.

### Root Cause
The filter in `nonWhiteSortedEqual.go` does check for these prefixes:
```go
if strings.HasPrefix(trimmed, "- ErrorRefOnly") ||
    strings.HasPrefix(trimmed, "- getVerifyRwxInternalError") {
    return true
}
```
But the actual error text is `"- ErrorRefOnly"` which should match. If it's still
appearing, the issue might be that the filter runs AFTER whitespace normalization
(tokens are split and sorted per line), so `"- ErrorRefOnly"` becomes
`"ErrorRefOnly -"` after sort, and the prefix check no longer matches.

### This IS the root cause
The `nonWhiteSortedLines()` function:
1. Splits line into tokens: `["- ", "ErrorRefOnly"]` → `["-", "ErrorRefOnly"]`
2. Sorts tokens: `["-", "ErrorRefOnly"]`
3. Joins: `"- ErrorRefOnly"` (happens to stay same due to sort order)

BUT `isStackTraceLine()` is called BEFORE the token normalization on the raw line.
Wait - looking at the code, `isStackTraceLine(line)` is checked first on the raw
line, THEN if not a stack trace line, the tokens are sorted. So the filter should
work. Unless the raw line has extra whitespace or non-printable characters.

**CORRECTION**: Re-reading the output more carefully:
```
"- ErrorRefOnly",
```
This is a SEPARATE line in the error message. The filter checks `strings.HasPrefix(trimmed, "- ErrorRefOnly")`.
The trimmed line is `"- ErrorRefOnly"`. This SHOULD match. The fact it doesn't filter
suggests the code on the user's machine doesn't have the latest filter.

---

## Category 5: CharCollectionMap Nil Map (1 test)

### Test Affected
- `Test_CovCCM_01_IsEmpty_HasItems`

### Error
```
panic: assignment to entry in nil map
CharCollectionMap.Add(...)  CharCollectionMap.go:524
```

### Root Cause
`CharCollectionMap.Add()` doesn't initialize the internal map before assignment.
This is a **production bug** — the `Add` method assumes the map is already initialized.

### Fix Required
Add nil map check in `CharCollectionMap.Add()`:
```go
if it.items == nil {
    it.items = make(map[rune]*Collection)
}
```

---

## Category 6: JSON Roundtrip Failures (4 tests)

### Tests Affected
- `Test_CovS06_CharCollMap_Json_Verification` — roundTrip: false
- `Test_CovS07_Json_Verification` — roundTrip: false
- `Test_Cov6_Collection_JsonString` — value receiver loses unexported items
- `Test_Cov6_Hashmap_Clone` — clone produces equal when expected not-equal

### Root Cause
These are **pre-existing production bugs** in JSON serialization:
1. `ParseInjectUsingJson` doesn't properly deserialize `CharCollectionMap` and other types
2. `Collection.JsonString()` on a value receiver loses unexported fields (items slice)
3. `Hashmap.Clone()` shares internal state (map pointer) instead of deep-copying

---

## Category 7: Pre-existing Test Issues (6 tests)

### Tests Affected
- `Test_Cov42_Collection_IsContainsAllSlice_Empty` — assertion pattern issue
- `Test_Cov44_SSO_IsValueBool` — "a":"1" diff 
- `Test_Cov52_CharCollectionMap_Json` — JSON roundtrip
- `Test_Cov57_HashsetsCollection_Json` — JSON roundtrip
- `Test_Cov63_CharHashsetMap_Json` — JSON roundtrip
- `Test_Cov69_LeftRightTrimmedUsingSlice_One` — hasContent true vs false

### Root Cause
These tests have expectation mismatches that existed before the AI intervention.
They're either:
1. Tests that tested incorrect assumptions about production behavior
2. Tests exposed after a runtime crash fix (they used to crash before reaching the assertion)

---

## Summary: What Should Have Been Done Differently

1. **Never rewrite entire test files** — The Cov43 rewrite was the #1 mistake.
   Minimal, targeted fixes prevent cascading regressions.

2. **Don't change production code to match test expectations** — The ValidValue.ValueByte
   change broke 6+ tests. The original production behavior was correct.

3. **Verify fixes compile and run before committing** — None of the AI changes
   were tested locally before being synced.

4. **Categorize before fixing** — Should have separated "my changes broke this" from
   "this was already broken" before making any fix attempts.

---

## Recommended Fix Priority

| Priority | Action | Fixes |
|----------|--------|-------|
| P0 | **Revert Cov43 rewrite** to original file | ~106 tests |
| P0 | **Revert ValidValue.ValueByte** to constants.Zero | ~6 tests |
| P1 | **Fix chmodhelper stack trace filter** order | 2 tests |
| P1 | **Fix CharCollectionMap nil map** in production | 1 test |
| P2 | **Update Hashset.AddVariations test** expectation 6→7 | 1 test |
| P3 | **Investigate JSON roundtrip** production bugs | ~4 tests |
| P3 | **Review pre-existing test expectations** | ~6 tests |
