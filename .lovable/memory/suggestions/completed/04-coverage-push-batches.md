# Completed: Coverage Push — Batches 1-3 (24 packages)

## Completed: 2026-03-14 to 2026-03-15

### What Was Done

Created 24 Coverage test files across 3 batches to push packages toward 100% coverage.

### Batch 1 (2026-03-14) — 11 files
Targeted packages at 75-97% coverage:
- mapdiffinternal, coreversion, reqtype, coreproperty, corefuncs
- results, coretestcases, pathinternal, stringutil, convertinternal, jsoninternal

### Batch 2 (2026-03-14) — 6 files
Targeted packages at 0-57% coverage:
- chmodhelper, args, errcore, regexnew, corecmp, codestack

### Batch 3 (2026-03-15) — 7 files
Targeted packages needing comprehensive coverage:
- coregeneric (Collection, Hashset, Hashmap, LinkedList, SimpleSlice, orderedfuncs, comparablefuncs, numericfuncs)
- strutilinternal, isany, coremath, corecsv, coretaskinfo, simplewrap

### Pattern Used
All tests follow `args.Map` + `ShouldBeEqual` pattern with `Test_Cov{N}_` naming convention.

### Still Pending
12 packages still need coverage files. See active suggestions tracker.
