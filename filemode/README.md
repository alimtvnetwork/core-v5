# filemode — File Permission Constants

Package `filemode` provides named `os.FileMode` constants for common Unix permission patterns, replacing raw octal literals with self-documenting names.

## Constants

### Named Permissions

| Constant | Octal | Description |
|----------|-------|-------------|
| `FileDefault` | `0644` | Owner read/write, group/other read |
| `DirDefault` | `0755` | Owner full, group/other read+execute |
| `FullAccess` | `0777` | All permissions for all |
| `OwnerFullAccessOnly` | `0700` | Owner full, no group/other |
| `OwnerGroupFullAccessOnly` | `0770` | Owner+group full, no other |
| `OwnerCanReadWriteGroupOtherCanReadOnly` | `0644` | Standard file permissions |
| `OwnerCanDoAllExecuteGroupOtherCanReadExecute` | `0755` | Standard directory permissions |

### Numeric Shortcuts

`X100` through `X777` — direct octal-to-name mappings (e.g., `X644 = 0644`).

## Usage

```go
import "github.com/alimtvnetwork/core/filemode"

os.WriteFile("config.json", data, filemode.FileDefault)
os.MkdirAll("/var/app", filemode.DirDefault)
```

## Related Docs

- [chmodhelper README](/chmodhelper/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
