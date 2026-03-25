# cmd

## Folder Purpose

Contains CLI entry-points for running, testing, and demonstrating the core package. Each sub-directory under `cmd/` is an independent `main` package.

## Responsibilities

1. Provide runnable executables for development and testing.
2. Serve as example usage of the core library.

## Key Files and Entrypoints

| Path | Purpose |
|------|---------|
| `cmd/main/` | Default package functionality testing |
| `cmd/README.md` | Documentation for CLI conventions |

## How to Extend Safely

- Create a new directory under `cmd/` with `package main`.
- Add a corresponding `make` target in the root `makefile`.
- Follow the naming pattern: `cmd/{name}/{name}.go`.

## Common Mistakes and Prevention

| Mistake | Prevention |
|---------|-----------|
| Forgetting `make` target | Always add `run-{name}` target |
| Using non-main package | All `cmd/` sub-dirs must be `package main` |

## Related Docs

- [CMD Entrypoints Spec](../12-cmd-entrypoints.md)
