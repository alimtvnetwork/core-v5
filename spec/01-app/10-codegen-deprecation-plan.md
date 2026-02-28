# Codegen Deprecation Plan

## Current State

### What Codegen Does Today

The `codegen/` package generates unit-test boilerplate from Go function signatures:

1. Reflects on a function to extract input/output parameters.
2. Applies templates to produce Arrange-Act-Assert test code.
3. Writes generated `.go` test files to disk.
4. Supports multiple generation modes: `Simple` (single arrange) and `MultipleArranges` (loop-based).

### What Depends on It

- `codegen/` sub-packages: `codegentype/`, `fmtcodegentype/`, `corecreator/`, `coreproperty/`, `aukast/`
- `tests/integratedtests/codegentests/` — integration tests for codegen
- `coretests/args/` — `FuncWrap` is used by codegen but also independently useful
- `coretests/coretestcases/` — `CaseV1` is used by codegen but also independently useful
- No downstream external packages are known to depend on codegen directly (needs confirmation).

## Two-Track Migration Plan

### Track A: Keep Buildable (Temporary)

1. Mark the `codegen/` package with a deprecation notice in its doc comments.
2. Ensure `codegen/` continues to compile with the rest of the module.
3. Do **not** add new features to codegen.
4. Maintain existing codegen tests so they pass.

### Track B: Remove Codegen Completely

1. **Audit consumers**: Search all repos in the auk-go org for imports of `codegen/`.
2. **Decouple shared types**: Move `coretests/args/FuncWrap` and `coretests/coretestcases/CaseV1` out of codegen's dependency tree (they are already in `coretests/` — confirm no circular deps).
3. **Remove codegen sub-packages**: Delete `codegentype/`, `fmtcodegentype/`, `corecreator/`, `coreproperty/`, `aukast/`.
4. **Remove codegen root**: Delete all files in `codegen/`.
5. **Remove codegen tests**: Delete `tests/integratedtests/codegentests/`.
6. **Clean go.mod**: Remove any dependencies only needed by codegen.
7. **Update docs**: Remove codegen references from README, CHANGELOG, specs.

## Exit Criteria

Before codegen can be deleted:

- [ ] No external package imports `codegen/` or its sub-packages.
- [ ] All shared types (`FuncWrap`, `CaseV1`) remain accessible via `coretests/`.
- [ ] All remaining tests pass after removal.
- [ ] README and specs are updated.
- [ ] A release tag is created before removal for rollback.

## Risks and Mitigations

| Risk | Mitigation |
|------|-----------|
| Breaking external consumers | Audit imports across all auk-go repos first |
| Missing generated test files that were committed | Verify generated files are not referenced at runtime |
| CI expects codegen to exist | Update CI pipelines and make targets |
| Loss of test generation capability | Document the testing patterns manually in specs |
