# Codegen Deprecation Plan

## Status: DEPRECATED — Active removal in progress

## Current State

### What Codegen Does Today

The `codegen/` package generates unit-test boilerplate from Go function signatures:

1. Reflects on a function to extract input/output parameters.
2. Applies templates to produce Arrange-Act-Assert test code.
3. Writes generated `.go` test files to disk.
4. Supports multiple generation modes: `Simple` (single arrange) and `MultipleArranges` (loop-based).

### What Depends on It

#### Internal consumers (within `github.com/alimtvnetwork/core`)

| File | Imports |
|------|---------|
| `cmd/main/unitTestGenerator.go` | `codegen`, `codegen/aukast`, `codegen/codegentype`, `codegen/coreproperty`, `codegen/fmtcodegentype` |
| `tests/integratedtests/codegentests/corepropertytests/Write_test.go` | `codegen/coreproperty` |

#### Sub-packages (internal to codegen)

| Package | Purpose |
|---------|---------|
| `codegen/aukast/` | Go AST reader and walker for code analysis |
| `codegen/codegentype/` | Enum for generation modes (`Simple`, `MultipleArranges`) |
| `codegen/corecreator/` | Template-based code creation helpers |
| `codegen/coreproperty/` | Property writer for code generation output |
| `codegen/fmtcodegentype/` | Enum for format types (`Default`, `WithExpect`, `WithFuncExpect`, `WithFuncError`) |

#### Shared types NOT affected by deprecation

These are used by codegen but live in `coretests/` and remain stable:

- `coretests/args.FuncWrap` — function reflection wrapper
- `coretests/args.One` — single test case arg container
- `coretests/coretestcases.CaseV1` — test case struct

#### External consumer audit

**Audit method**: Search all repos in the `auk-go` GitLab org for imports containing `codegen`.

**Known external consumers**: None confirmed. The `codegen` package appears to be internal-only tooling.

**Recommended audit commands**:
```bash
# Run against each auk-go repo:
grep -r "auk-go/core/codegen" --include="*.go" .

# Or use GitLab search API across the group:
# GET /groups/auk-go/search?scope=blobs&search=auk-go/core/codegen
```

## Two-Track Migration Plan

### Track A: Keep Buildable (Current — Temporary)

1. ✅ Mark the `codegen/` package and all sub-packages with deprecation notices via `doc.go`.
2. ✅ Ensure `codegen/` continues to compile with the rest of the module.
3. Do **not** add new features to codegen.
4. Maintain existing codegen tests so they pass.

### Track B: Remove Codegen Completely

#### Prerequisites (must complete before deletion)

- [ ] Confirm no external repos import `codegen/` (run audit commands above).
- [ ] Verify `coretests/args/FuncWrap` and `coretests/coretestcases/CaseV1` have no circular dependency on codegen.
- [ ] Create a git tag (e.g. `v0.x.y-pre-codegen-removal`) for rollback.

#### Removal steps

1. **Remove `cmd/main/unitTestGenerator.go`** — the only cmd consumer. This is a developer tool, not a production entrypoint.
2. **Remove test directory**: `tests/integratedtests/codegentests/` — integration tests for codegen only.
3. **Remove codegen sub-packages**: `aukast/`, `codegentype/`, `corecreator/`, `coreproperty/`, `fmtcodegentype/`.
4. **Remove codegen root**: All files in `codegen/`.
5. **Clean go.mod**: Run `go mod tidy` to remove any dependencies only needed by codegen.
6. **Update docs**: Remove codegen references from README, CHANGELOG, specs.
7. **Update CI**: Remove any make targets that depend on codegen.

## Exit Criteria

Before codegen can be deleted:

- [ ] No external package imports `codegen/` or its sub-packages.
- [ ] All shared types (`FuncWrap`, `CaseV1`) remain accessible via `coretests/`.
- [ ] All remaining tests pass after removal.
- [ ] README and specs are updated.
- [ ] A release tag is created before removal for rollback.
- [ ] `cmd/main/unitTestGenerator.go` is removed or refactored.

## Risks and Mitigations

| Risk | Mitigation |
|------|-----------|
| Breaking external consumers | Audit imports across all auk-go repos first |
| Missing generated test files that were committed | Verify generated files are not referenced at runtime |
| CI expects codegen to exist | Update CI pipelines and make targets |
| Loss of test generation capability | Document the testing patterns manually in specs (done: see `13-testing-patterns.md`) |
| `cmd/main/unitTestGenerator.go` stops compiling | Remove it as part of Track B step 1 |

## Dependency Graph

```
cmd/main/unitTestGenerator.go
  └── codegen/
        ├── codegen/aukast/
        ├── codegen/codegentype/
        ├── codegen/corecreator/
        ├── codegen/coreproperty/
        ├── codegen/fmtcodegentype/
        ├── chmodhelper/
        ├── coredata/corestr/
        ├── coretests/args/          ← shared, NOT deprecated
        ├── coretests/coretestcases/  ← shared, NOT deprecated
        ├── errcore/
        ├── internal/convertinternal/
        └── internal/reflectinternal/
```
