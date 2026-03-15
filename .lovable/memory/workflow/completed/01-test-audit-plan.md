# Test Audit & Diagnostics Improvement Plan

## Status: ✅ Completed (2026-03-11)

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
- corevalidatortests (7 files)

### 3. ✅ Fix 21 Failing Tests (Iteration 4)
Fixed all 21 failing tests. See `completed/03-fix-21-failing-tests.md` for full details.

## Moved To Pending
The following tasks from this plan remain open and are tracked in the active workflow:
- Test title audit for remaining packages → `01-coverage-and-testing-plan.md`
- Nil receiver coverage audit → `01-coverage-and-testing-plan.md`
- Deep clone production bug → `01-coverage-and-testing-plan.md`
