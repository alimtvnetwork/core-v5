# Coverage & Testing Master Plan

## Status: In Progress
## Last Updated: 2026-03-16

## Critical Root Cause Checkpoint
Coverage work has been repeatedly invalidated by assumed APIs, broad unverified coverage-file generation, and skipping the compile-first gate. Do **not** treat newly written coverage files as successful until `./run.ps1 PC` and then `./run.ps1 TC` confirm the result.

See finalized postmortem memory: `.lovable/memory/workflow/completed/02-coverage-remediation-root-cause.md`
See issue record: `issues/repeated-coverage-remediation-failure-root-cause.md`

---

## Completed Tasks

### 1. ✅ 100% Coverage Push — Batch 1 (11 packages)
**Completed: 2026-03-14**

Created Coverage test files for 11 packages that were between 75-97% coverage:
- `mapdiffinternal` (97.2% → targeting 100%)
- `coreversion` (97.1%)
- `reqtype` (97%)
- `coreproperty` (96.2%)
- `keymk` (96.2%) — *needs separate file, not yet created*
- `corefuncs` (96%)
- `coretests/results` (95.2%)
- `coredata/corerange` (94.9%) — *needs separate file, not yet created*
- `coredata/coreonce` (94%) — *needs separate file, not yet created*
- `coreimpl/enumimpl` (93.2%) — *needs separate file, not yet created*
- `convertinternal` (89.8%)
- `jsoninternal` (89.5%)
- `coredata/stringslice` (83.1%) — *needs separate file, not yet created*
- `corevalidator` (82.8%) — *needs separate file, not yet created*
- `coreutils/stringutil` (82%)
- `coretests/coretestcases` (79.6%)
- `internal/pathinternal` (75.6%)

**Files Created:**
- `tests/integratedtests/mapdiffinternaltests/Coverage3_test.go`
- `tests/integratedtests/coreversiontests/Coverage3_test.go`
- `tests/integratedtests/reqtypetests/Coverage3_test.go`
- `tests/integratedtests/corepropertytests/Coverage_test.go`
- `tests/integratedtests/corefuncstests/Coverage3_test.go`
- `tests/integratedtests/pathinternaltests/Coverage3_test.go`
- `tests/integratedtests/stringutiltests/Coverage4_test.go`
- `tests/integratedtests/convertinternaltests/Coverage4_test.go`
- `tests/integratedtests/jsoninternaltests/Coverage4_test.go`
- `tests/integratedtests/resultstests/Coverage_test.go`
- `tests/integratedtests/coretestcasestests/Coverage5_test.go`

### 2. ✅ 100% Coverage Push — Batch 2 (6 packages)
**Completed: 2026-03-14**

Created Coverage test files for 6 packages with lower coverage (0-57%):
- `chmodhelper` (57%)
- `coretests/args` (43.7%)
- `errcore` (25.7%)
- `regexnew` (14.3%)
- `corecmp` (16.8%)
- `codestack` (0%)

**Files Created:**
- `tests/integratedtests/chmodhelpertests/Coverage3_test.go`
- `tests/integratedtests/argstests/Coverage4_test.go`
- `tests/integratedtests/errcoretests/Coverage7_test.go`
- `tests/integratedtests/regexnewtests/Coverage5_test.go`
- `tests/integratedtests/corecmptests/Coverage7_test.go`
- `tests/integratedtests/codestacktests/Coverage6_test.go`

### 3. ✅ 100% Coverage Push — Batch 3 (7 packages)
**Completed: 2026-03-15**

Created Coverage test files for 7 more packages:
- `coregeneric` — Collection, Hashset, Hashmap, LinkedList, SimpleSlice, orderedfuncs, comparablefuncs, numericfuncs
- `strutilinternal` — AnyToString, MaskLine, SplitLeftRight, CurlyWrapIf, Clone, ReflectInterfaceVal, NonEmpty
- `isany` — ReflectNull, ReflectValueNull, NullBoth, DefinedBoth, DefinedAllOf/AnyOf, AllNull, AnyNull
- `coremath` — MaxInt/MinInt equal edges, Integer boundary checks
- `corecsv` — All quote branches, empty inputs, Stringers
- `coretaskinfo` — InfoJson, InfoMap, Plain/Secure creators, Deserialized, ExcludeOptions
- `simplewrap` — TitleCurlyWrap, TitleSquare, With/WithPtr, ConditionalWrapWith, DoubleQuoteWrapElements

**Files Created:**
- `tests/integratedtests/coregenerictests/Coverage3_test.go`
- `tests/integratedtests/strutilinternaltests/Coverage3_test.go`
- `tests/integratedtests/isanytests/Coverage8_test.go`
- `tests/integratedtests/coremathtests/Coverage2_test.go`
- `tests/integratedtests/corecsvtests/Coverage2_test.go`
- `tests/integratedtests/coretaskinfotests/Coverage2_test.go`
- `tests/integratedtests/simplewraptests/Coverage6_test.go`

### 4. ✅ Coverage Prompt Generator System
**Completed: 2026-03-15**

Built a PowerShell-based system that auto-generates AI-friendly prompt files for coverage gaps:
- **`scripts/coverage/Generate-CoveragePrompts.ps1`** — Main script: parses coverage.out + `go tool cover -func`, identifies all functions <100%, extracts uncovered line ranges, writes batched prompt files (500 functions/file) to `data/prompts/`
- **`scripts/coverage/Get-UncoveredLines.ps1`** — Utility: extracts uncovered lines for a specific source file
- **`scripts/coverage/Get-FunctionCoverage.ps1`** — Utility: filters function coverage by threshold
- **Integration**: Auto-called at end of `./run.ps1 TC` after coverage summary

---

## Pending Tasks

### 5. 🔲 Remaining Packages Still Below 100%
The following packages from Batch 1 were noted but coverage files were NOT created:
- `keymk` (96.2%)
- `coredata/corerange` (94.9%)
- `coredata/coreonce` (94%)
- `coreimpl/enumimpl` (93.2%)
- `coredata/stringslice` (83.1%)
- `corevalidator` (82.8%)

Additional packages from Batch 2 not yet addressed:
- `corepayload` (42.6%)
- `reflectinternal` (22.9%)
- `corejson` (20%)
- `corestr` (2.8%)
- `coredynamic` (2.3%)
- `reflectmodel` (0.4%)

### 6. 🔲 Verify All New Tests Compile & Pass
Run `./run.ps1 TC` to verify all 24 new Coverage test files compile and produce expected coverage improvements. Check for blocked packages.

### 7. 🔲 Test Title Audit — Remaining Packages
Packages not yet audited for title convention:
- coreoncetests, corejsontests, corepayloadtests, coreversiontests
- ostypetests, reqtypetests, issettertests, stringslicetests
- trydotests, codefuncstests, codestacktests, coreappendtests
- corecomparatortests, corecmptests, converterstests
- enumimpltests, regexnewtests

### 8. 🔲 Nil Receiver Coverage Audit
Systematically migrate nil receiver test cases from CaseV1 to CaseNilSafe pattern across all packages.

### 9. 🔲 Deep Clone Production Bug Investigation
`corepayload.Attributes.Clone(deep=true)` returns error. Needs production code investigation.

### 10. 🔲 Test Runner Hardening
Review all test runners for:
- Unconditional map key insertion (like the `containsName: false` pattern)
- Value vs pointer type assertions
- Incorrect independence/equality check logic

### 11. 🔲 Diagnostic Output Regression Tests
Create snapshot tests for diagnostic output formatting to prevent regressions.
