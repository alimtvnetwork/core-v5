# Module Spec: chmodhelper

## What Problem It Solves

Managing Unix file permissions programmatically in Go is verbose and error-prone. `chmodhelper` provides:
- Human-readable rwx string parsing (`"rwxr-xr--"` → `os.FileMode`)
- Permission verification against expectations
- Safe file I/O with automatic directory creation and permission control
- Recursive path operations with permission metadata

## API Usage Examples

### Reading Existing Permissions

```go
rwx, err := chmodhelper.GetExistingChmodRwxWrapper("/path/to/file")
// rwx.String() → "rwxr-xr--"
```

### Writing Files with Permissions

```go
writer := chmodhelper.New.SimpleFileReaderWriter.Options(
    true,  // overwrite
    true,  // create dirs
    true,  // create file
    "/path/to/output.txt",
)
err := writer.Write([]byte("content"))
```

### Verifying Permissions

```go
isMatch := chmodhelper.ChmodVerify.IsChmodEqualUsingRwxOwnerGroupOther(
    "/path/to/file",
    "rwx", "r-x", "r--",
)
```

### Parsing RWX Instructions

```go
fileMode, err := chmodhelper.ParseRwxInstructionToFileMode("rwxr-xr--")
```

## Error Handling Expectations

- All operations that touch the filesystem return `error`.
- Use `PathExistStat` to check existence before operations.
- `GetRwxLengthError` validates rwx string lengths.
- File write operations use `globalMutex` for thread safety.

## Logging Expectations

- No explicit logging — errors are returned to the caller.
- Stack traces can be added via `errcore.StackEnhance`.

## Unit Test Examples

Tests are in `tests/integratedtests/chmodhelpertests/`.

```go
func Test_ParseRwx(t *testing.T) {
    for i, tc := range parseRwxTestCases {
        // Arrange
        input := tc.ArrangeInput.(string)
        
        // Act
        result, err := chmodhelper.ParseRwxInstructionToFileMode(input)
        
        // Assert
        tc.ShouldBeEqual(t, i, result)
    }
}
```

## Key Types

| Type | Purpose |
|------|---------|
| `RwxWrapper` | Full rwx permission representation (owner, group, other) |
| `SimpleFileReaderWriter` | File I/O with permission control |
| `VarAttribute` | Variable permission attribute |
| `SingleRwx` | Single r/w/x permission set |
| `PathExistStat` | Path existence check result with file info |
| `RwxInstructionExecutor` | Executable chmod instruction |

## Related Docs

- [chmodhelper folder spec](../folders/01-chmodhelper.md)
- [Repo Overview](../00-repo-overview.md)
