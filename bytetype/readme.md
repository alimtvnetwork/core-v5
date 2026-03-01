# bytetype — Byte-Backed Enum Type

## Overview

Package `bytetype` provides a `Variant` enum type backed by `byte` (0–255). It implements the `BasicEnumContractsBinder` interface from `enuminf`, supports JSON marshal/unmarshal, comparison operations, range validation, and named string representations via `BasicEnumImpl`.

## Architecture

```
bytetype/
├── bytetype.go            # Variant type — constants, enum methods, JSON, comparison
├── vars.go                # BasicEnumImpl — enum implementation via enumimpl
├── New.go                 # New(byte) Variant — factory function
├── String.go              # String([]byte) string — byte slice to string
├── GetSet.go              # GetSet(bool, Variant, Variant) — ternary for Variant
├── GetSetVariant.go       # GetSetVariant(bool, byte, byte) — ternary returning Variant
├── IsCompareResult.go     # IsCompareResult — comparison dispatch using corecomparator
└── readme.md
```

## Variant Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `Zero` / `Min` | `0` | Minimum value |
| `One` | `1` | — |
| `Two` | `2` | — |
| `Three` | `3` | — |
| `Max` | `255` | Maximum (`math.MaxUint8`) |

## Variant Methods

### Identity & Display

| Method | Description |
|--------|-------------|
| `Name()` | Named string via `BasicEnumImpl` |
| `String()` | Same as `Name()` |
| `NameValue()` | `"Name(Value)"` format |
| `ValueString()` / `StringValue()` / `ToNumberString()` | Numeric string |
| `JsonString()` | JSON string representation |

### Value Access

| Method | Description |
|--------|-------------|
| `Value()` / `ValueByte()` | Raw `byte` |
| `ValueInt()` / `ValueInt8()` / `ValueInt16()` / `ValueInt32()` | Typed integer conversions |
| `ValueUInt16()` | `uint16` conversion |

### Identity Checks

| Method | Description |
|--------|-------------|
| `IsZero()` / `IsOne()` / `IsTwo()` / `IsThree()` | Named constant checks |
| `IsMin()` / `IsMax()` | Boundary checks |
| `IsValid()` / `IsInvalid()` | Non-zero / zero check |
| `Is(Variant)` | Equality with another Variant |

### Comparison

| Method | Description |
|--------|-------------|
| `IsEqual(byte)` / `IsEqualInt(int)` | Equality |
| `IsGreater(byte)` / `IsGreaterInt(int)` | `v > n` |
| `IsGreaterEqual(byte)` / `IsGreaterEqualInt(int)` | `v >= n` |
| `IsLess(byte)` / `IsLessInt(int)` | `v < n` |
| `IsLessEqual(byte)` / `IsLessEqualInt(int)` | `v <= n` |
| `IsBetween(start, end byte)` / `IsBetweenInt(start, end int)` | Range check |
| `IsValueEqual(byte)` | Raw byte equality |
| `IsCompareResult(byte, Compare)` | Dispatch comparison by `corecomparator.Compare` |

### Arithmetic

| Method | Description |
|--------|-------------|
| `Add(byte)` | `v + n` |
| `Subtract(byte)` | `v - n` |

### Enum Interface

| Method | Description |
|--------|-------------|
| `IsEnumEqual(BasicEnumer)` | Compare with any enum |
| `IsAnyEnumsEqual(...BasicEnumer)` | Any enum matches |
| `IsNameEqual(string)` / `IsAnyNamesOf(...string)` | Name-based matching |
| `IsValidRange()` / `IsInvalidRange()` | Range validation |
| `AllNameValues()` / `StringRanges()` / `IntegerEnumRanges()` | Enum metadata |
| `RangeNamesCsv()` / `RangesInvalidMessage()` / `RangesInvalidErr()` | Error/display helpers |
| `TypeName()` / `EnumType()` | Type metadata |
| `Format(string)` | Custom format |
| `HasIndexInStrings(...string)` | Use Variant value as index into string slice |
| `MarshalJSON()` / `UnmarshalJSON([]byte)` | JSON serialization |
| `AsBasicEnumContractsBinder()` | Interface cast |
| `ToPtr()` | Pointer conversion |

## Package-Level Functions

| Function | Description |
|----------|-------------|
| `New(byte)` | Create `Variant` from raw byte |
| `String([]byte)` | Convert byte slice to string (empty-safe) |
| `GetSet(bool, Variant, Variant)` | Ternary selection between two Variants |
| `GetSetVariant(bool, byte, byte)` | Ternary selection returning Variant from raw bytes |

## Dependencies

| Package | Usage |
|---------|-------|
| `coreimpl/enumimpl` | `BasicEnumImpl` creation |
| `corecomparator` | Comparison dispatch |
| `corejson` | JSON serialization |
| `enuminf` | Enum interface contracts |

## Usage

```go
import "gitlab.com/auk-go/core/bytetype"

v := bytetype.New(2)
fmt.Println(v.Name())       // "Two"
fmt.Println(v.IsGreater(1)) // true
fmt.Println(v.IsBetween(1, 3)) // true

// Ternary
result := bytetype.GetSet(true, bytetype.One, bytetype.Three)
fmt.Println(result) // "One"

// JSON
data, _ := bytetype.Two.MarshalJSON()
fmt.Println(string(data)) // "Two"
```

## Related Docs

- [Enum Interface Contracts](../coreinterface/enuminf/README.md)
- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
