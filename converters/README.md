# converters — Type Conversion Utilities

Package `converters` provides struct-as-namespace type conversion utilities for strings, bytes, integers, maps, JSON formatting, and dynamic `any` values. Delegates to internal implementations for core logic.

## Architecture

```
converters/
├── vars.go                    # Package singletons: StringTo, BytesTo, StringsTo, AnyTo, Map, PrettyJson, etc.
├── stringTo.go                # StringTo.* — string → int, float64, byte, JSON bytes
├── stringsTo.go               # StringsTo.* — []string → map conversions
├── bytesTo.go                 # BytesTo.* — []byte → string, pointer conversions
├── unsafeBytesTo.go           # Unsafe byte conversion (zero-copy)
├── anyItemConverter.go        # AnyTo.* — any → string, []string conversions
├── StringsToMapConverter.go   # StringsToMapConverter — []string → map[string]string
└── coreconverted/             # Result types for conversion outputs
    ├── Integers.go            # Integers — converted int slice with combined error
    └── Bytes.go               # Bytes — converted byte slice with combined error
```

## Entry Points

| Namespace | Description |
|-----------|-------------|
| `converters.StringTo.*` | Convert string to int, float64, byte, JSON bytes |
| `converters.BytesTo.*` | Convert []byte / *[]byte to string |
| `converters.StringsTo.*` | Convert []string to map types |
| `converters.AnyTo.*` | Convert any value to string or []string |
| `converters.Map.*` | Map conversion utilities (delegated from internal) |
| `converters.PrettyJson.*` | JSON pretty-print formatting |
| `converters.JsonString.*` | JSON string formatting |
| `converters.Integers.*` | Integer conversion utilities (delegated from internal) |
| `converters.KeyValuesTo.*` | Key-value pair conversions |
| `converters.CodeFormatter.*` | Code formatting utilities |

## Usage

### String Conversions

```go
import "gitlab.com/auk-go/core/converters"

// String to integer
val, err := converters.StringTo.Integer("42")

// String to integer with default
val, ok := converters.StringTo.IntegerWithDefault("abc", -1) // -1, false

// String to float64
f, err := converters.StringTo.Float64("3.14")

// String to byte
b, err := converters.StringTo.Byte("255")

// Batch string to integers
result := converters.StringTo.IntegersWithDefaults("1,2,3", ",", 0)
// result.Values = [1, 2, 3], result.CombinedError = nil
```

### Byte Conversions

```go
// Bytes to string
s := converters.BytesTo.String([]byte("hello"))

// Pointer bytes to string
s = converters.BytesTo.PtrString(&rawBytes)
```

### JSON Formatting

```go
// Pretty-print JSON
prettyStr := converters.PrettyJson.String(jsonBytes)
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
