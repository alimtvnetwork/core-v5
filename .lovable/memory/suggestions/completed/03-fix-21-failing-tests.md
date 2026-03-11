# Completed: Fix 21 Failing Tests

## Completed: 2026-03-11

### Production Bugs Fixed
1. **WrappedErr.HasErrorOrException** — operator precedence: `&&` vs `||` without parentheses caused nil pointer panic
2. **SimpleSlice.InsertAt** — `append(s[:index+1], s[index:]...)` caused slice bounds panic; fixed to use copy-based insert

### Test Logic Bugs Fixed
3. **codefuncstests** — 4 test functions added `containsName: false` unconditionally, causing extra map entry mismatch
4. **corejsontests** — `corejson.NewResult.UsingBytes()` returns value, not pointer; fixed type assertion
5. **stringslicetests** — `isIndependentCopy` logic conflated "returned original" with "independent copy"

### Expectation Corrections (16 test cases across 9 files)
- BytesErrorOnce: `[]byte{}` is defined/non-empty; lifecycle panics cascade
- ErrorOnce: ConcatNewString wraps nil-error result in quotes
- OsType: group names include "Group" suffix; Unknown is capitalized
- ReqType: Create name is "CreateUsingAliasMap"; Drop/CreateOrUpdate flags updated
- Version: v4 vs v4.0 with LeftGreater returns true
- IsOutOfRange: value 5 is out of range
- Attributes: deep clone returns error (graceful handling added)
