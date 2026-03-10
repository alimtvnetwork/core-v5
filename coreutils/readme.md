# coreutils — String Utility Functions

## Overview

Package `coreutils` contains the `stringutil` sub-package, which provides a comprehensive set of string inspection, conversion, splitting, masking, and template replacement functions.

## Architecture

```
coreutils/
└── stringutil/
    ├── vars.go                              # ReplaceTemplate singleton
    ├── consts.go                            # ExpectedLeftRightLength
    ├── all-interfaces.go                    # namer interface
    ├── replaceTemplate.go                   # Template replacement engine
    │
    │  — Type Conversion —
    ├── ToBool.go                            # ToBool(string) bool
    ├── ToByte.go                            # ToByte(string) (byte, error)
    ├── ToByteDefault.go                     # ToByteDefault(string, byte) byte
    ├── ToInt.go                             # ToInt(string) (int, error)
    ├── ToIntDef.go                          # ToIntDef(string, int) int
    ├── ToIntDefault.go                      # ToIntDefault(string) int
    ├── ToIntUsingRegexMatch.go              # ToIntUsingRegexMatch
    ├── ToInt8.go / ToInt8Def.go             # int8 conversion
    ├── ToInt16.go / ToInt16Default.go       # int16 conversion
    ├── ToInt32.go / ToInt32Def.go           # int32 conversion
    ├── ToUint16Default.go                   # uint16 conversion
    ├── ToUint32Default.go                   # uint32 conversion
    │
    │  — String Inspection —
    ├── IsEmpty.go / IsEmptyPtr.go           # Empty checks
    ├── IsNotEmpty.go                        # Non-empty check
    ├── IsBlank.go / IsBlankPtr.go           # Whitespace-only checks
    ├── IsDefined.go / IsDefinedPtr.go       # Non-nil, non-empty checks
    ├── IsNullOrEmptyPtr.go                  # Nil or empty pointer check
    ├── IsEmptyOrWhitespace.go / Ptr         # Empty or whitespace checks
    │
    │  — Prefix / Suffix —
    ├── IsStarts.go / IsStartsWith.go        # String prefix checks
    ├── IsStartsChar.go / IsStartsRune.go    # Char/rune prefix checks
    ├── IsEnds.go / IsEndsWith.go            # String suffix checks
    ├── IsEndsChar.go / IsEndsRune.go        # Char/rune suffix checks
    ├── IsStartsAndEndsChar.go               # Both prefix and suffix char
    ├── IsStartsAndEndsWith.go               # Both prefix and suffix string
    ├── IsAnyStartsWith.go                   # Any of multiple prefixes
    ├── IsAnyEndsWith.go                     # Any of multiple suffixes
    │
    │  — Contains —
    ├── IsContains.go / IsContainsPtr.go     # Substring checks
    ├── IsContainsPtrSimple.go               # Simplified pointer contains
    │
    │  — Any-to-String —
    ├── AnyToString.go                       # Convert any to string
    ├── AnyToStringNameField.go              # Convert using Name field
    ├── AnyToTypeString.go                   # Type-annotated string
    │
    │  — Splitting —
    ├── SplitFirstLast.go                    # Split into first and last
    ├── SplitLeftRight.go                    # Split into left and right
    ├── SplitLeftRightTrimmed.go             # Split + trim
    ├── SplitLeftRightType.go                # Split with type info
    ├── SplitLeftRightTypeTrimmed.go         # Split with type + trim
    ├── SplitLeftRightsTrims.go              # Multiple split + trim
    ├── SplitContentsByWhitespaceConditions.go  # Conditional whitespace split
    │
    │  — Substring —
    ├── SafeSubstring.go                     # Safe substring (no panic)
    ├── SafeSubstringStarts.go               # Safe substring from start
    ├── SafeSubstringEnds.go                 # Safe substring from end
    ├── FirstChar.go                         # First character extraction
    │
    │  — Masking —
    ├── MaskLine.go                          # Mask single line
    ├── MaskLines.go                         # Mask multiple lines
    ├── MaskTrimLine.go                      # Mask + trim single line
    ├── MaskTrimLines.go                     # Mask + trim multiple lines
    │
    │  — Transformation —
    ├── ClonePtr.go                          # Clone string pointer
    ├── SafeClonePtr.go                      # Nil-safe clone pointer
    ├── RemoveMany.go                        # Remove multiple substrings
    ├── RemoveManyBySplitting.go             # Remove by splitting
    ├── TrimKeepSingleSpaceOnly.go           # Collapse whitespace
    ├── KeyValReplacer.go                    # Key-value template replacement
    └── readme.md
```

## Key Function Categories

### Type Conversion

| Function | Description |
|----------|-------------|
| `ToInt(string)` | Parse int with error |
| `ToIntDef(string, int)` / `ToIntDefault(string)` | Parse int with default |
| `ToInt8` / `ToInt16` / `ToInt32` + `Def` variants | Typed integer parsing |
| `ToUint16Default` / `ToUint32Default` | Unsigned integer parsing |
| `ToBool(string)` | Parse boolean |
| `ToByte(string)` / `ToByteDefault(string, byte)` | Parse byte |

### String Inspection

| Function | Description |
|----------|-------------|
| `IsEmpty` / `IsEmptyPtr` / `IsNotEmpty` | Emptiness checks |
| `IsBlank` / `IsBlankPtr` | Whitespace-only checks |
| `IsDefined` / `IsDefinedPtr` | Non-nil, non-empty |
| `IsNullOrEmptyPtr` | Nil or empty pointer |

### Prefix / Suffix

| Function | Description |
|----------|-------------|
| `IsStarts` / `IsStartsWith` / `IsStartsChar` / `IsStartsRune` | Prefix checks |
| `IsEnds` / `IsEndsWith` / `IsEndsChar` / `IsEndsRune` | Suffix checks |
| `IsStartsAndEndsChar` / `IsStartsAndEndsWith` | Both prefix + suffix |
| `IsAnyStartsWith` / `IsAnyEndsWith` | Any of multiple prefixes/suffixes |

### Splitting

| Function | Description |
|----------|-------------|
| `SplitFirstLast(string, sep)` | First and last segments |
| `SplitLeftRight(string, sep)` | Left and right of separator |
| `SplitLeftRightTrimmed` / `SplitLeftRightType` | Trimmed/typed variants |

## Usage

```go
import "github.com/alimtvnetwork/core/coreutils/stringutil"

// Inspection
stringutil.IsEmpty("")           // true
stringutil.IsBlank("   ")       // true
stringutil.IsDefined("hello")   // true

// Conversion
val := stringutil.ToIntDef("42", 0) // 42
b := stringutil.ToBool("yes")       // true

// Splitting
left, right := stringutil.SplitLeftRight("key=value", "=")
// left: "key", right: "value"

// Safe substring
sub := stringutil.SafeSubstring("hello", 1, 3) // "el"
```

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
