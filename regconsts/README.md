# regconsts — Regex Pattern Constants

Package `regconsts` provides pre-defined regex pattern strings for common validation and parsing tasks. These are raw strings — compile them with `regexp.Compile` or use with [`regexnew`](/regexnew/README.md) for lazy compilation.

## Categories

### Validation Patterns

| Constant | Validates |
|----------|-----------|
| `Alpha` | Letters only |
| `AlphaNumeric` | Letters and digits |
| `AlphaDash` | Letters, digits, underscore, dash |
| `AlphaSpace` | Letters, digits, underscore, space, dash |
| `Numeric` | Signed integers |
| `Int` | Strict signed integers |
| `FloatingPointNumberOnly` | Floating-point numbers |
| `Hexadecimal` | Hex characters |
| `Base64` | Base64 encoded strings |
| `ASCII` | ASCII-only characters |

### Format Patterns

| Constant | Format |
|----------|--------|
| `UUID` / `UUID3` / `UUID4` / `UUID5` | UUID versions |
| `Date` | `yyyy-mm-dd` |
| `DateDDMMYY` | `dd/mm/yyyy` |
| `Semver` | Semantic versioning |
| `ISBN10` / `ISBN13` | Book identifiers |
| `IMEI` / `IMSI` | Device identifiers |

### Network Patterns

| Constant | Pattern |
|----------|---------|
| `Url` | HTTP(S) URLs |
| `IpSimple` / `IpEthernet` / `IpWithSubnet` / `IpWithPort` | IP addresses |
| `Ipv6WithSubnet` | IPv6 addresses |
| `MacAddress` / `ValidMac` | MAC addresses |
| `DNSName` | DNS hostnames |
| `SSN` | US Social Security Numbers |

### Code/Config Patterns

| Constant | Pattern |
|----------|---------|
| `HashComment` | `# comment` lines |
| `SlashComment` | `// comment` lines |
| `AllWhitespaces` | Whitespace sequences |
| `EachWordsWithDollarSymbolDefinition` | `$var` or `${var}` |
| `ThirdBracketHeader` | `[section]` headers |

## Usage

```go
import (
    "regexp"
    "github.com/alimtvnetwork/core/regconsts"
)

re := regexp.MustCompile(regconsts.UUID4)
if re.MatchString(input) {
    // valid UUID v4
}
```

## Related Docs

- [regexnew README](/regexnew/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
