# Codegen Deprecation Plan

## Status: ✅ COMPLETE — Removed in v1.6.0

The `codegen/` package and all sub-packages (`aukast/`, `codegentype/`, `corecreator/`, `coreproperty/`, `fmtcodegentype/`) have been fully removed from the repository.

### What Was Removed

| Item | Description |
|------|-------------|
| `codegen/` (entire tree) | Test boilerplate generator and all sub-packages |
| `cmd/main/unitTestGenerator.go` | CLI consumer of codegen |
| `tests/integratedtests/codegentests/` | Integration tests for codegen |
| `tests/integratedtests/corepropertytests/` | Tests for `codegen/coreproperty` |

### What Was Retained

- `coretests/args.FuncWrap` — function reflection wrapper (lives in `coretests/`, unaffected)
- `coretests/coretestcases.CaseV1` — test case struct (lives in `coretests/`, unaffected)

### Audit Result

- No external repos import `codegen/` (confirmed via workspace search)
- `coretests/` has zero imports from `codegen/` (confirmed)
- All remaining tests pass after removal
