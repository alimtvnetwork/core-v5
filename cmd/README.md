# Command Line Interfaces

All CLI entry-points are located under `cmd/`. Each sub-directory contains a `package main`.

## Available Commands

| Directory | Make Target | Description |
|-----------|-------------|-------------|
| `cmd/main/` | `make` or `make run-main` | Default package functionality testing |
| `cmd/server/` | `make run-server` | Server runner |
| `cmd/client/` | `make run-client` | Client runner |
| `cmd/sample/` | `make run-sample` | Sample/demo runner |

## Building

```bash
make build          # builds cmd/main/ → ./build/cli
make build-run      # build and run immediately
```

## Creating a New Command

1. Create a new directory: `cmd/mycommand/`
2. Create `cmd/mycommand/main.go` with `package main`
3. Add a make target to the root `makefile`:
   ```makefile
   run-mycommand:
       go run cmd/mycommand/*.go
   ```

## Testing

```bash
make run-tests      # runs tests/ directory
```

## Links

- [Spec: CMD Entrypoints](/spec/01-app/12-cmd-entrypoints.md)
- [Spec: Repo Overview](/spec/01-app/00-repo-overview.md)

## Contributors

## Issues for Future Reference
