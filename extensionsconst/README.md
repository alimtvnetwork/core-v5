# extensionsconst — File Extension Constants

Package `extensionsconst` provides string constants for common file extensions in two forms: with dot prefix (`DotJson`) and without (`Json`).

## Constants

### With Dot Prefix (`dot-extensions.go`)

| Constant | Value |
|----------|-------|
| `DotJson` | `.json` |
| `DotYaml` | `.yaml` |
| `DotTxt` | `.txt` |
| `DotPdf` | `.pdf` |
| `DotSql` | `.sql` |
| `DotSh` | `.sh` |
| `DotLog` | `.log` |
| `DotZip` | `.zip` |
| `DotPem` | `.pem` |
| ... | 50+ extensions |

### Without Dot Prefix (`extensions.go`)

Same set without the leading dot (e.g., `Json = "json"`).

### Wildcards

| Constant | Value | Description |
|----------|-------|-------------|
| `AllFiles` | `*.*` | Match all files |
| `DotAny` | `.*` | Any extension |
| `AllExtensionFilterStart` | `*.` | Glob prefix |

## Usage

```go
import "github.com/alimtvnetwork/core/extensionsconst"

filename := "config" + extensionsconst.DotJson  // "config.json"
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
