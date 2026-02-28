# CMD Entrypoints

## Overview

The `cmd/` directory contains CLI entry-points. Each sub-directory is an independent `package main`.

## Available Commands

| Command | Make Target | Description |
|---------|-------------|-------------|
| `cmd/main/` | `make run-main` or `make` | Default package functionality testing |
| `cmd/server/` | `make run-server` | Server runner (if exists) |
| `cmd/client/` | `make run-client` | Client runner (if exists) |
| `cmd/sample/` | `make run-sample` | Sample/demo runner (if exists) |

## Build

```bash
make build   # builds to ./build/cli from cmd/main/
```

## Inputs and Outputs

- Currently no CLI flags or arguments are documented.
- Output goes to stdout/stderr.
- Build output goes to `./build/cli`.

## Examples

```bash
# Run default
make

# Build and run
make build-run

# Run tests
make run-tests
```

## Common Errors and Fixes

| Error | Fix |
|-------|-----|
| `No go in PATH` | Install Go and ensure it's in PATH |
| `cmd/server/ does not exist` | Not all make targets have corresponding directories |
| Build fails on Windows | Use Windows-specific paths in makefile |

## Related Docs

- [cmd/README.md](/cmd/README.md)
- [Repo Overview](./00-repo-overview.md)
