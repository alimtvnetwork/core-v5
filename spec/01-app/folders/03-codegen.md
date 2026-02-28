# codegen (DEPRECATED)

## Folder Purpose

Generates unit test boilerplate code from Go function signatures. It inspects functions via reflection, builds test templates with arrange/act/assert patterns, and writes `.go` test files.

## Responsibilities

1. Reflect on function signatures to extract input/output argument types.
2. Generate formatted test code using templates.
3. Write generated test files to disk.
4. Support multiple generation modes (simple, multiple arranges).

## Key Files and Entrypoints

| File | Purpose |
|------|---------|
| `GenerateFunc.go` | Main generator: builds test code from function metadata |
| `FinalCode.go` | Compiled output holder with write capability |
| `GoCode.go` | Go code representation with package header |
| `all-interfaces.go` | `BaseGenerator` interface contract |
| `vars.go` | Package singletons, templates, utility references |
| `template_func.go` | Test function templates |
| `testCaseGenerator.go` | Test case compilation |
| `variablesGenerator.go` | Variable setup generation |

## Deprecation Status

This package is **scheduled for removal**. See [Codegen Deprecation Plan](../10-codegen-deprecation-plan.md).

## Related Docs

- [Codegen Deprecation Plan](../10-codegen-deprecation-plan.md)
- [Repo Overview](../00-repo-overview.md)
