# osconsts — OS-Specific Constants

Package `osconsts` provides runtime-detected OS constants and architecture information for cross-platform code.

## Variables

| Variable | Type | Description |
|----------|------|-------------|
| `IsX32Architecture` | `bool` | True if running on 32-bit architecture |
| `IsX64Architecture` | `bool` | True if running on 64-bit architecture |
| `CurrentSystemArchitecture` | `string` | `runtime.GOARCH` value |
| `X32Architectures` | `[]string` | Known 32-bit architecture names |
| `X32ArchitecturesMap` | `map[string]bool` | O(1) lookup for 32-bit architectures |

### OS-Specific Constants

Platform-specific values (line separators, paths, shell commands) are provided via `internal/osconstsinternal` and exposed through package-level variables.

## Usage

```go
import "github.com/alimtvnetwork/core/osconsts"

if osconsts.IsX64Architecture {
    // use 64-bit optimized path
}
```

## Related Docs

- [ostype README](/ostype/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
