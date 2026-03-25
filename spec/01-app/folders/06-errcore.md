# errcore

## Folder Purpose

Rich error construction, formatting, merging, and stack-trace enhancement. Provides structured error messages with context, expectations, variable maps, and Gherkins-style output.

## Responsibilities

1. Build errors with meaningful context (source/destination, variable maps, references).
2. Merge multiple errors into single error or string.
3. Format errors with stack traces.
4. Provide expectation-based error messages for testing.
5. Handle compiled errors and raw error collections.

## Key Files and Entrypoints

| File | Purpose |
|------|---------|
| `vars.go` | `ShouldBe`, `Expected`, `StackEnhance` singletons |
| `MergeErrors.go` | Combine multiple errors |
| `StackTracesCompiled.go` | Stack trace formatting |
| `VarMap.go`, `VarTwo.go`, `VarThree.go` | Variable-context error builders |
| `Expecting*.go` | Expectation error messages |
| `GherkinsString*.go` | Gherkins-style formatted output |

## How to Extend Safely

- New error types: create a new file with the pattern `{ErrorType}.go`.
- Compose with existing types; don't duplicate formatting logic.

## Logging Convention

All diagnostic output uses `log/slog` with structured key-value pairs. No centralized slog config lives in this library — consumer apps are responsible for configuring their own `slog.Handler` (JSON for production, text for development).

## Related Docs

- [Repo Overview](../00-repo-overview.md)
