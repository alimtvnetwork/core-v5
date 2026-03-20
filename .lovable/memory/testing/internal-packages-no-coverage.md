# Memory: Internal Packages — No Coverage Tests

## Rule (STRICT — Never Override)

**Do NOT write code-coverage tests for any package under the `internal/` folder.**

This applies to ALL current and future internal sub-packages:

- `internal/convertinternal`
- `internal/csvinternal`
- `internal/fsinternal`
- `internal/internalinterface`
- `internal/jsoninternal`
- `internal/mapdiffinternal`
- `internal/messages`
- `internal/msgcreator`
- `internal/msgformats`
- `internal/osconstsinternal`
- `internal/pathinternal`
- `internal/reflectinternal`
- `internal/strutilinternal`
- `internal/trydo`

## What This Means

1. **No coverage test files** — never create `Coverage*_test.go` files targeting internal packages.
2. **No coverage iteration work** — never include internal packages in coverage push plans.
3. **Existing tests stay** — do not remove tests that already exist; they may serve business/integration purposes.
4. **Business-critical tests are OK** — writing tests for important business logic in internal packages is fine, but these must NOT be motivated by coverage percentage goals.
5. **Coverage reporting** — internal packages should be excluded or ignored in coverage reports and plans.

## Rationale

Internal packages are private implementation details. Spending time on coverage for them is not productive. Coverage effort must focus exclusively on public/non-internal packages.
