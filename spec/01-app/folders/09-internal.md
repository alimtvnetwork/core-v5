# internal

## Folder Purpose

Private implementation helpers not intended for external consumption. Go enforces `internal/` as non-importable by packages outside the module.

## Sub-Packages

| Package | Purpose |
|---------|---------|
| `reflectinternal/` | Reflection helpers: type name extraction, looper, isChecker, mapConverter |
| `convertinteranl/` | Type conversion internals (note: **typo** in folder name — should be `convertinternal`) |
| `pathinternal/` | Path manipulation: `Join()`, `ParentDir()`, `GetTemp()` |
| `csvinternal/` | CSV formatting internals |
| `jsoninternal/` | JSON processing internals |
| `fsinternal/` | File-system internals |
| `mapdiffinternal/` | Map diff logic |
| `messages/` | Internal message templates |
| `msgcreator/` | Message creator utilities |
| `msgformats/` | Message format strings |
| `osconstsinternal/` | OS constants internals |
| `strutilinternal/` | String utility internals |
| `trydo/` | Try-do pattern utilities |
| `internalinterface/` | Internal-only interface contracts |

## Known Issues

- `convertinteranl` — typo; should be `convertinternal`. See issues.
- Some internal packages are quite large and could benefit from refactoring.

## Related Docs

- [Repo Overview](../00-repo-overview.md)
