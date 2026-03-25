# constants — Shared Constants & Capacity Values

Package `constants` defines project-wide shared constants: strings, characters, runes, bytes, numeric values, format strings, capacity presets, OS commands, path separators, and runtime variables. Used by virtually every package in the module.

## Architecture

```
constants/
├── constants.go               # Primary constants: strings, chars, runes, bytes, symbols, OS commands, comparisons
├── capacity-constants.go      # ArbitraryCapacity*, WordCapacity*, LineCapacity*, FileCapacity*, Capacity* presets
├── messages-formats.go        # Sprint format strings: SprintValueFormat, TitleValueFormat, EnumOnlySupportedFormat, etc.
├── vars.go                    # Runtime variables: pointer copies of constants, NewLineBytes, CpuNumber, MaxWorker
├── varsptr.go                 # Pointer variables for optional/nullable constant usage
├── arrayvars.go               # Array/slice constant variables
├── arrayvarsptr.go            # Pointer array/slice variables
├── line_darwin.go             # OS-specific: NewLine for macOS
├── line_linux.go              # OS-specific: NewLine for Linux
├── line_windows.go            # OS-specific: NewLine for Windows
├── bitsize/                   # Bit-size constants
│   └── Of8, Of16, ..., Of256G
├── percentconst/              # Percentage constants
└── regkeysconsts/             # Registry key constants (Windows)
```

## Constant Categories

### Strings & Symbols

Common string constants for consistent usage across the codebase:

```go
constants.EmptyString       // ""
constants.Space             // " "
constants.Hyphen            // "-"
constants.Comma             // ","
constants.CommaSpace        // ", "
constants.Dot               // "."
constants.Colon             // ":"
constants.ForwardSlash      // "/"
constants.NewLineUnix       // "\n"
constants.NewLineWindows    // "\r\n"
constants.DefaultLine       // "\n"
constants.Tab               // "\t"
constants.NilAngelBracket   // "<nil>"
constants.WildcardSymbol    // "*"
```

### Byte & Rune Variants

Every string symbol has corresponding `byte` and `rune` variants:

```go
constants.DotChar           // byte '.'
constants.DotRune           // rune '.'
constants.CommaChar         // byte ','
constants.CommaRune         // rune ','
constants.SpaceChar         // byte ' '
constants.SpaceRune         // rune ' '
```

### Numeric Constants

```go
constants.Zero              // 0
constants.One               // 1
constants.InvalidIndex      // -1
constants.MinusOne          // -1
constants.MaxInt            // platform max int
constants.MinInt            // platform min int
```

### Capacity Presets

Pre-defined capacity values for `make()` calls to avoid magic numbers:

```go
constants.ArbitraryCapacity10   // 10
constants.ArbitraryCapacity30   // 30
constants.ArbitraryCapacity100  // 100
constants.Capacity64            // 64
constants.Capacity256           // 256
constants.Capacity1G            // 1024
```

### Format Strings

Printf-style format templates:

```go
constants.SprintValueFormat               // "%v"
constants.SprintPropertyNameValueFormat   // "%+v"
constants.SprintTypeFormat                // "%T"
constants.SprintDoubleQuoteFormat         // "%q"
constants.TitleValueFormat                // "%s : %v"
constants.FromToFormat                    // `{From : %q, To: %q}`
```

### Runtime Variables

```go
constants.CpuNumber         // runtime.NumCPU()
constants.ProcessorCount    // alias for CpuNumber
constants.MaxWorker         // CpuNumber * 5
constants.SafeWorker        // CpuNumber * 3
constants.NewLineBytes      // []byte(NewLine)
```

## Sub-Packages

| Package | Description |
|---------|-------------|
| `bitsize/` | Bit-size constants: `Of8`, `Of16`, `Of32`, `Of64`, up to `Of256G` |
| `percentconst/` | Percentage value constants |
| `regkeysconsts/` | Windows registry key path constants |

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
