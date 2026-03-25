# chmodhelper

## Folder Purpose

Provides a comprehensive toolkit for Unix-style file permission management in Go. Handles parsing rwx strings, converting to/from `os.FileMode`, verifying permissions, applying chmod operations, reading/writing files with permission control, and recursive path operations.

## Responsibilities

1. **Parse** rwx instruction strings (e.g. `"rwxr-xr--"`) into structured types (`RwxWrapper`, `VarAttribute`, `SingleRwx`).
2. **Verify** existing file/directory permissions against expected values.
3. **Apply** chmod changes to files and directories.
4. **Read/Write** files with built-in permission management (`SimpleFileReaderWriter`).
5. **Recursive** path enumeration with permission metadata.
6. **Merge** wildcard rwx patterns with fixed permissions.
7. **Filter** paths by existence and permission criteria.

## Key Files and Entrypoints

| File | Purpose |
|------|---------|
| `vars.go` | Package-level singletons: `New`, `ChmodApply`, `ChmodVerify`, `SimpleFileWriter` |
| `RwxWrapper.go` | Core rwx data structure |
| `SimpleFileReaderWriter.go` | File I/O with permission options (overwrite, create dirs) |
| `chmodApplier.go` | Apply chmod to paths |
| `chmodVerifier.go` | Verify permissions match expectations |
| `ParseRwxInstructionToFileMode.go` | Parse rwx string → `os.FileMode` |
| `GetExistingChmod*.go` | Read current permissions from filesystem |
| `GetRecursivePaths*.go` | Recursive directory traversal |
| `RwxInstructionExecutor.go` | Execute chmod instructions |

## How to Extend Safely

- Add new parsers as separate files following the `Parse*.go` naming pattern.
- New verification logic should implement or compose with existing `chmodVerifier`.
- Always handle `PathExistStat` errors — the package uses stat-based checks extensively.
- Use `SimpleFileReaderWriter` options pattern for new file operations.

## Common Mistakes and Prevention

| Mistake | Prevention |
|---------|-----------|
| Ignoring `PathExistStat` errors | Always check `IsPathExists` before operations |
| Assuming Unix-only | Package has some cross-platform awareness but test on target OS |
| Not using `globalMutex` for concurrent writes | Use `SimpleFileReaderWriter` which handles locking |

## Related Docs

- [Module Spec: chmod-helper](../modules/01-chmod-helper.md)
- [Repo Overview](../00-repo-overview.md)
