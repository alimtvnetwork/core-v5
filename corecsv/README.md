# corecsv — CSV String Compilation

## Overview

The `corecsv` package provides functions for converting slices of values (`string`, `any`, `fmt.Stringer`, `func() string`) into CSV-formatted strings with configurable joiners and quoting styles.

## Architecture

```
corecsv/
├── toString.go                                # Internal: any → string via convertinternal
├── StringsToCsvString.go                      # []string → joined CSV string
├── StringsToCsvStrings.go                     # []string → []string with quoting
├── StringsToCsvStringsDefault.go              # []string → []string (no quotes)
├── StringsToStringDefault.go                  # []string → joined string (default joiner)
├── AnyItemsToCsvString.go                     # []any → joined CSV string
├── AnyItemsToCsvStrings.go                    # []any → []string with quoting
├── AnyItemsToStringDefault.go                 # []any → joined string (default)
├── AnyToTypesCsvDefault.go                    # []any → type-aware CSV (default)
├── AnyToTypesCsvStrings.go                    # []any → []string with type info
├── AnyToValuesTypeString.go                   # []any → single type+value string
├── AnyToValuesTypeStrings.go                  # []any → []string with type+value
├── StringersToCsvStrings.go                   # []fmt.Stringer → []string with quoting
├── StringersToString.go                       # []fmt.Stringer → joined string
├── StringersToStringDefault.go                # []fmt.Stringer → joined string (default)
├── CompileStringersToCsvStrings.go            # []func()string → []string with quoting
├── CompileStringersToString.go                # []func()string → joined string
├── CompileStringersToStringDefault.go         # []func()string → joined string (default)
├── StringFunctionsToString.go                 # []func()string → []string with quoting
├── DefaultCsv.go                              # Shorthand: strings → CSV (comma+space, quoted)
├── DefaultCsvStrings.go                       # Shorthand: strings → []string (quoted)
├── DefaultCsvUsingJoiner.go                   # Shorthand: strings → CSV (custom joiner)
├── DefaultAnyCsv.go                           # Shorthand: any → CSV (comma+space)
├── DefaultAnyCsvStrings.go                    # Shorthand: any → []string
├── DefaultAnyCsvUsingJoiner.go                # Shorthand: any → CSV (custom joiner)
├── RangeNamesWithValuesIndexes.go             # []string → ["Name[0]", "Name[1]", ...]
├── RangeNamesWithValuesIndexesCsvString.go    # Range names → joined CSV
└── RangeNamesWithValuesIndexesString.go       # Range names → single string
```

## Quoting Modes

All parameterized functions accept `isIncludeQuote` and `isIncludeSingleQuote` flags:

| `isIncludeQuote` | `isIncludeSingleQuote` | Output Format |
|:-:|:-:|---|
| `true` | `true` | `'value'` (single-quoted) |
| `true` | `false` | `"value"` (double-quoted) |
| `false` | `false` | `value` (unquoted) |

## Function Categories

### String Input

| Function | Returns | Description |
|----------|---------|-------------|
| `StringsToCsvString(joiner, quote, singleQuote, items...)` | `string` | Join strings with quoting |
| `StringsToCsvStrings(quote, singleQuote, items...)` | `[]string` | Quote each string |
| `StringsToCsvStringsDefault(items...)` | `[]string` | No-quote passthrough |
| `StringsToStringDefault(items...)` | `string` | Default comma+space join |
| `DefaultCsv(items...)` | `string` | Shorthand: quoted, comma+space |
| `DefaultCsvStrings(items...)` | `[]string` | Shorthand: quoted strings |
| `DefaultCsvUsingJoiner(joiner, items...)` | `string` | Shorthand: quoted, custom joiner |

### Any Input

| Function | Returns | Description |
|----------|---------|-------------|
| `AnyItemsToCsvString(joiner, quote, singleQuote, items...)` | `string` | Join any values with quoting |
| `AnyItemsToCsvStrings(quote, singleQuote, items...)` | `[]string` | Quote each any value |
| `AnyItemsToStringDefault(items...)` | `string` | Default comma+space join |
| `AnyToTypesCsvDefault(items...)` | `string` | Type-aware CSV (default) |
| `AnyToTypesCsvStrings(quote, singleQuote, items...)` | `[]string` | Type-aware with quoting |
| `AnyToValuesTypeString(items...)` | `string` | Single type+value string |
| `AnyToValuesTypeStrings(items...)` | `[]string` | Type+value per item |
| `DefaultAnyCsv(items...)` | `string` | Shorthand: any, comma+space |
| `DefaultAnyCsvStrings(items...)` | `[]string` | Shorthand: any, quoted |
| `DefaultAnyCsvUsingJoiner(joiner, items...)` | `string` | Shorthand: any, custom joiner |

### Stringer / Function Input

| Function | Returns | Description |
|----------|---------|-------------|
| `StringersToCsvStrings(quote, singleQuote, stringers...)` | `[]string` | `fmt.Stringer` → quoted strings |
| `StringersToString(joiner, quote, singleQuote, stringers...)` | `string` | `fmt.Stringer` → joined string |
| `StringersToStringDefault(stringers...)` | `string` | `fmt.Stringer` → default join |
| `CompileStringersToCsvStrings(quote, singleQuote, funcs...)` | `[]string` | `func() string` → quoted strings |
| `CompileStringersToString(joiner, quote, singleQuote, funcs...)` | `string` | `func() string` → joined string |
| `CompileStringersToStringDefault(funcs...)` | `string` | `func() string` → default join |
| `StringFunctionsToString(quote, singleQuote, funcs...)` | `[]string` | `func() string` → quoted strings |

### Range Name Utilities

| Function | Returns | Description |
|----------|---------|-------------|
| `RangeNamesWithValuesIndexes(items...)` | `[]string` | `["Name[0]", "Name[1]", ...]` |
| `RangeNamesWithValuesIndexesCsvString(items...)` | `string` | Range names as CSV string |
| `RangeNamesWithValuesIndexesString(items...)` | `string` | Range names as single string |

## Usage Examples

```go
import "github.com/alimtvnetwork/core/corecsv"

// Simple default CSV
result := corecsv.DefaultCsv("a", "b", "c")
// → "'a', 'b', 'c'"

// Unquoted with custom joiner
result := corecsv.StringsToCsvString(" | ", false, false, "x", "y", "z")
// → "x | y | z"

// Any values
result := corecsv.DefaultAnyCsv(1, "hello", true)
// → "'1', 'hello', 'true'"

// Range names
names := corecsv.RangeNamesWithValuesIndexes("Red", "Green", "Blue")
// → ["Red[0]", "Green[1]", "Blue[2]"]
```

## Key Patterns

- Empty input always returns empty string or empty slice — no panics.
- Internal `toString()` delegates to `convertinternal.AnyTo.SmartString` for type-safe conversion.
- All quoting is applied uniformly via `fmt.Sprintf` with `constants` format strings.

## Related Docs

- [Repo Overview](../spec/01-app/00-repo-overview.md)
