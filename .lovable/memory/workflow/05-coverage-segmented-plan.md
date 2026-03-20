# Coverage 100% Push — Segmented Iteration Plan

## Status: 🔧 Active
## Last Updated: 2026-03-20
## Source Data: `data/coverage/per-package-coverage-17.json`

---

## Rules

1. **Internal packages are EXCLUDED** — never write coverage tests for `internal/*`
2. **Large packages (>1000 uncovered stmts)** are split into ~200-line segments
3. **Small packages (<200 uncovered)** are grouped 2-4 per iteration
4. **Each iteration** = one "next" command from the user
5. **Follow AAA pattern**, `CaseV1` / `args.Map` / `ShouldBeEqual` conventions
6. **Read source files before writing tests** — never assume APIs
7. **Test file naming**: `Coverage{N}_test.go` in `tests/integratedtests/{pkg}tests/`
8. **Function naming**: `Test_Cov{N}_{Method}_{Context}`

---

## Package Summary (Non-Internal, Below 100%)

| Package | Coverage | Uncovered Stmts | Segments |
|---------|----------|-----------------|----------|
| `namevalue` | 0.0% | 188 | 1 |
| `reflectcore/reflectmodel` | 0.8% | 251 | 2 |
| `coredata/corestr` | 3.6% | 5553 | **28** |
| `coredata/coredynamic` | 3.7% | 2191 | **11** |
| `regexnew` | 87.0% | 29 | 1 |
| `chmodhelper` | 89.6% | 170 | 1 |
| `coretests` | 90.5% | 35 | 1 |
| `coretests/args` | 90.9% | 156 | 1 |
| `coredata/corepayload` | 92.9% | 117 | 1 |
| `errcore` | 93.6% | 53 | 1 |
| `coredata/coregeneric` | 94.7% | 57 | 1 |
| `coredata/corejson` | 95.0% | 106 | 1 |
| `corecmp` | 95.1% | 9 | 1 |
| `codestack` | 95.2% | 24 | 1 |
| `corevalidator` | 95.4% | 33 | 1 |
| `coretests/coretestcases` | 95.9% | 11 | 1 |
| `coreinstruction` | 95.9% | 16 | 1 |
| `codegen/coreproperty` | 96.2% | 2 | 1 |
| `coreimpl/enumimpl` | 96.3% | 55 | 1 |
| `coretests/results` | 96.6% | 5 | 1 |
| `iserror` | 97.4% | 1 | 1 |
| `coreutils/stringutil` | 98.0% | 9 | 1 |
| `simplewrap` | 98.1% | 2 | 1 |
| `keymk` | 98.5% | 6 | 1 |
| `coremath` | 98.5% | 1 | 1 |
| `enums/versionindexes` | 98.6% | 1 | 1 |
| `reqtype` | 99.1% | 2 | 1 |
| `coreversion` | 99.2% | 3 | 1 |
| `coredata/stringslice` | 99.2% | 4 | 1 |
| `coretaskinfo` | 99.2% | 2 | 1 |
| `coredata/coreonce` | 99.3% | 5 | 1 |
| `isany` | 99.4% | 1 | 1 |
| `ostype` | 99.4% | 1 | 1 |
| `issetter` | 99.6% | 1 | 1 |
| `coredata/corerange` | 99.7% | 2 | 1 |
| `converters` | 99.8% | 1 | 1 |

**Total uncovered: ~9,179 stmts across 36 packages**

---

## Iteration Plan

### Phase 1 — Quick Wins (≤10 uncovered stmts each) — Iterations 1-3

**Iteration 1**: `corecmp` (9), `coretests/results` (5), `coredata/coreonce` (5), `coredata/stringslice` (4)
→ 23 stmts total

**Iteration 2**: `coreversion` (3), `coretaskinfo` (2), `codegen/coreproperty` (2), `simplewrap` (2), `coredata/corerange` (2), `reqtype` (2)
→ 14 stmts total

**Iteration 3**: `keymk` (6), `iserror` (1), `isany` (1), `ostype` (1), `issetter` (1), `converters` (1), `coremath` (1), `enums/versionindexes` (1)
→ 13 stmts total

### Phase 2 — Small-Medium Packages — Iterations 4-8

**Iteration 4**: `coretests/coretestcases` (11), `coreinstruction` (16)
→ 27 stmts total

**Iteration 5**: `codestack` (24), `regexnew` (29)
→ 53 stmts total

**Iteration 6**: `corevalidator` (33), `coretests` (35)
→ 68 stmts total

**Iteration 7**: `errcore` (53), `coreimpl/enumimpl` (55)
→ 108 stmts total

**Iteration 8**: `coredata/coregeneric` (57), `coredata/corejson` (106)
→ 163 stmts total

### Phase 3 — Medium Packages — Iterations 9-12

**Iteration 9**: `coretests/args` (156)
→ 156 stmts

**Iteration 10**: `chmodhelper` (170)
→ 170 stmts

**Iteration 11**: `coredata/corepayload` (117), `namevalue` (188)
→ 305 stmts total (namevalue = new test file from scratch)

**Iteration 12**: `reflectcore/reflectmodel` — Segment A (first ~200 of 251)
→ ~200 stmts

### Phase 4 — reflectmodel remainder — Iteration 13

**Iteration 13**: `reflectcore/reflectmodel` — Segment B (remaining ~51)
→ ~51 stmts

### Phase 5 — coredata/coredynamic (2191 uncovered) — Iterations 14-24

**Iteration 14**: `coredynamic` Segment A — lines/stmts 1-200
**Iteration 15**: `coredynamic` Segment B — lines/stmts 201-400
**Iteration 16**: `coredynamic` Segment C — lines/stmts 401-600
**Iteration 17**: `coredynamic` Segment D — lines/stmts 601-800
**Iteration 18**: `coredynamic` Segment E — lines/stmts 801-1000
**Iteration 19**: `coredynamic` Segment F — lines/stmts 1001-1200
**Iteration 20**: `coredynamic` Segment G — lines/stmts 1201-1400
**Iteration 21**: `coredynamic` Segment H — lines/stmts 1401-1600
**Iteration 22**: `coredynamic` Segment I — lines/stmts 1601-1800
**Iteration 23**: `coredynamic` Segment J — lines/stmts 1801-2000
**Iteration 24**: `coredynamic` Segment K — lines/stmts 2001-2191

### Phase 6 — coredata/corestr (5553 uncovered) — Iterations 25-52

**Iteration 25**: `corestr` Segment A — stmts 1-200
**Iteration 26**: `corestr` Segment B — stmts 201-400
**Iteration 27**: `corestr` Segment C — stmts 401-600
**Iteration 28**: `corestr` Segment D — stmts 601-800
**Iteration 29**: `corestr` Segment E — stmts 801-1000
**Iteration 30**: `corestr` Segment F — stmts 1001-1200
**Iteration 31**: `corestr` Segment G — stmts 1201-1400
**Iteration 32**: `corestr` Segment H — stmts 1401-1600
**Iteration 33**: `corestr` Segment I — stmts 1601-1800
**Iteration 34**: `corestr` Segment J — stmts 1801-2000
**Iteration 35**: `corestr` Segment K — stmts 2001-2200
**Iteration 36**: `corestr` Segment L — stmts 2201-2400
**Iteration 37**: `corestr` Segment M — stmts 2401-2600
**Iteration 38**: `corestr` Segment N — stmts 2601-2800
**Iteration 39**: `corestr` Segment O — stmts 2801-3000
**Iteration 40**: `corestr` Segment P — stmts 3001-3200
**Iteration 41**: `corestr` Segment Q — stmts 3201-3400
**Iteration 42**: `corestr` Segment R — stmts 3401-3600
**Iteration 43**: `corestr` Segment S — stmts 3601-3800
**Iteration 44**: `corestr` Segment T — stmts 3801-4000
**Iteration 45**: `corestr` Segment U — stmts 4001-4200
**Iteration 46**: `corestr` Segment V — stmts 4201-4400
**Iteration 47**: `corestr` Segment W — stmts 4401-4600
**Iteration 48**: `corestr` Segment X — stmts 4601-4800
**Iteration 49**: `corestr` Segment Y — stmts 4801-5000
**Iteration 50**: `corestr` Segment Z — stmts 5001-5200
**Iteration 51**: `corestr` Segment AA — stmts 5201-5400
**Iteration 52**: `corestr` Segment AB — stmts 5401-5553

---

## Execution Notes

- For large packages (`corestr`, `coredynamic`), each segment will:
  1. Read the uncovered source files relevant to that ~200-stmt slice
  2. Identify uncovered branches/paths from `coverage.out`
  3. Write tests in a new or appended `Coverage{N}_test.go`
  4. Segment boundaries are approximate — will align to logical file/function boundaries
- Segment letters (A, B, C...) will map to specific source files once we begin each phase
- Total estimated iterations: **52**
