# Memory: testing/test-compliance-audit-plan
Updated: 2026-03-26

## Audit Summary

### Issue 1: Tests Inside Source Packages (WRONG LOCATION)
**92 files** across **14 packages** need to move to `tests/integratedtests/{pkg}tests/`.
Internal packages are excluded per rules.

| Source Package | Files | Target Directory |
|---|---|---|
| `coredata/corejson/` | 25 | `tests/integratedtests/corejsontests/` |
| `coredata/corestr/` | 20 | `tests/integratedtests/corestrtests/` |
| `coredata/coredynamic/` | 11 | `tests/integratedtests/coredynamictests/` |
| `errcore/` | 11 | `tests/integratedtests/errcoretests/` |
| `codestack/` | 5 | `tests/integratedtests/codestacktests/` |
| `coredata/corepayload/` | 5 | `tests/integratedtests/corepayloadtests/` |
| `coretests/args/` | 3 | `tests/integratedtests/argstests/` |
| `coredata/stringslice/` | 2 | `tests/integratedtests/stringslicetests/` |
| `corecmp/` | 2 | `tests/integratedtests/corecmptests/` |
| `chmodhelper/` | 2 | `tests/integratedtests/chmodhelpertests/` |
| `regexnew/` | 2 | `tests/integratedtests/regexnewtests/` |
| `reflectcore/reflectmodel/` | 2 | `tests/integratedtests/reflectmodeltests/` |
| `coretests/` | 1 | `tests/integratedtests/coreteststests/` |
| `coreinstruction/` | 1 | `tests/integratedtests/coreinstructiontests/` |

**Moving steps per file:**
1. Change `package` declaration from source pkg to `{pkg}tests`
2. Update imports to reference source package explicitly
3. Prefix any direct (unexported) access with exported alternatives or helper wrappers
4. Move file to target directory
5. Delete original from source package

### Issue 2: Tests Using `t.Fatal`/`t.Error` (WRONG ASSERTION STYLE)
**299 files** in `tests/integratedtests/` + **92 files** in source packages use raw Go testing assertions.

**Required change:** Replace `t.Fatal("msg")`, `t.Errorf(...)`, `if x != y { t.Fatal(...) }` patterns with:
- `CaseV1` + `ShouldBeEqualMap` for multi-field checks
- `CaseV1` + `ShouldBeEqual` for single-value checks  
- `convey.Convey` + `convey.So` for inline assertions
- Separate test data into `_testcases.go` files

### Issue 3: Missing AAA Comments
**570 files** in `tests/integratedtests/` missing `// Arrange`, `// Act`, `// Assert` comments.

### Issue 4: args.Map Values on Single Line
Many test cases have `args.Map{"key1": val1, "key2": val2}` on one line instead of multi-line format.

## Execution Plan (Iterative, 2 packages per "next")

### Phase 1: Move In-Package Tests (Priority — blocks coverage measurement)
Move files from source packages → integrated tests directory.
Order by file count (largest first):
1. `corejson` (25 files)
2. `corestr` (20 files)
3. `coredynamic` (11 files)
4. `errcore` (11 files)
5. `codestack` (5 files)
6. `corepayload` (5 files)
7. `coretests/args` (3 files)
8. Remaining small packages (2 files each): `stringslice`, `corecmp`, `chmodhelper`, `regexnew`, `reflectmodel`, `coretests`, `coreinstruction`

### Phase 2: Fix Assertion Style (t.Fatal → CaseV1/GoConvey)
Convert raw assertions to framework style. This is the most labor-intensive phase.
Process 2 packages per iteration.

### Phase 3: Add AAA Comments
Add `// Arrange`, `// Act`, `// Assert` to all test functions missing them.

### Phase 4: Format args.Map to Multi-Line
Reformat single-line `args.Map` to one key-value per line.

## Rules Reminder
- Never modify production code unless fixing a blocker
- `internal/` packages excluded
- Existing working tests must not be deleted
- Test titles: `"{Function} returns {Result} -- {Input Context}"`
- Use `params.go` for map key constants
- Use native types in `args.Map` (not stringified)
