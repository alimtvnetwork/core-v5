# typesconv — Pointer-Value Type Conversions

## Overview

Package `typesconv` provides nil-safe pointer↔value conversion functions for Go primitive types (`string`, `int`, `bool`, `byte`, `float32`). Each type follows a consistent pattern: `Ptr`, `PtrToSimple`, `PtrToSimpleDef`, `PtrToDefPtr`, and `PtrDefValFunc`. Also includes `string↔bool` conversion utilities.

## Architecture

```
typesconv/
├── string.go      # String pointer/value conversions + StringToBool
├── interger.go    # Int pointer/value conversions (note: filename typo preserved)
├── bool.go        # Bool pointer/value conversions
├── byte.go        # Byte pointer/value conversions
├── float.go       # Float32 pointer/value conversions
└── readme.md
```

## Function Pattern (Per Type)

Each type `T` (string, int, bool, byte, float32) has these functions:

| Function | Signature | Description |
|----------|-----------|-------------|
| `{T}Ptr` | `(val T) *T` | Wrap value as pointer |
| `{T}PtrToSimple` | `(val *T) T` | Dereference pointer; returns zero-value if nil |
| `{T}PtrToSimpleDef` | `(val *T, defVal T) T` | Dereference with custom default |
| `{T}PtrToDefPtr` | `(val *T, defVal T) *T` | Returns `val` if non-nil, else `&defVal` |
| `{T}PtrDefValFunc` | `(val *T, fn func() T) *T` | Lazy default via function if nil |

## String-Specific Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `StringToBool` | `(string) bool` | Parses `"yes"/"no"/"true"/"false"/etc.` |
| `StringPointerToBool` | `(*string) bool` | Nil-safe string-to-bool |
| `StringPointerToBoolPtr` | `(*string) *bool` | Nil-safe string-to-bool pointer |
| `StringToBoolPtr` | `(string) *bool` | String-to-bool pointer |

## Complete Function List

### string (`string.go`)

`StringPtr`, `StringPtrToSimple`, `StringPtrToSimpleDef`, `StringPtrToDefPtr`, `StringPtrDefValFunc`, `StringToBool`, `StringPointerToBool`, `StringPointerToBoolPtr`, `StringToBoolPtr`

### int (`interger.go`)

`IntPtr`, `IntPtrToSimple`, `IntPtrToSimpleDef`, `IntPtrToDefPtr`, `IntPtrDefValFunc`

### bool (`bool.go`)

`BoolPtr`, `BoolPtrToSimple`, `BoolPtrToSimpleDef`, `BoolPtrToDefPtr`, `BoolPtrDefValFunc`

### byte (`byte.go`)

`BytePtr`, `BytePtrToSimple`, `BytePtrToSimpleDef`, `BytePtrToDefPtr`, `BytePtrDefValFunc`

### float32 (`float.go`)

`FloatPtr`, `FloatPtrToSimple`, `FloatPtrToSimpleDef`, `FloatPtrToDefPtr`, `FloatPtrDefValFunc`

## Usage

```go
import "gitlab.com/auk-go/core/typesconv"

// Pointer creation
p := typesconv.IntPtr(42)  // *int pointing to 42

// Safe dereference with default
val := typesconv.IntPtrToSimpleDef(nil, -1)  // -1

// String to bool
typesconv.StringToBool("yes")   // true
typesconv.StringToBool("false") // false
typesconv.StringToBool("")      // false

// Lazy default
result := typesconv.StringPtrDefValFunc(nil, func() string {
    return "computed-default"
})
// result: *string pointing to "computed-default"
```

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
