# corerange — Range Data Types & Boundary Validation

## Overview

The `corerange` package provides range types and boundary-checking utilities for numeric and string values. It includes min/max bounds, typed range structs, and the `Within` validator for clamping and validating values against numeric ranges.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `Within` | `*within` | Range validation and clamping utilities |

## Range Types

| Type | Description |
|------|-------------|
| `BaseRange` | Base range with raw input, separator, validity flags |
| `RangeInt` | Integer range with start/end and min/max bounds |
| `RangeByte` | Byte range |
| `RangeInt8` | Int8 range |
| `RangeInt16` | Int16 range |
| `RangeString` | String-based range |
| `RangeAny` | Generic range for any type |
| `StartEndInt` | Simple start/end integer pair |
| `StartEndString` | Simple start/end string pair |
| `StartEndSimpleString` | Simplified start/end string |

## MinMax Types

| Type | Description |
|------|-------------|
| `MinMaxInt` | Min/max integer bounds |
| `MinMaxInt8` | Min/max int8 bounds |
| `MinMaxInt16` | Min/max int16 bounds |
| `MinMaxInt64` | Min/max int64 bounds |
| `MinMaxByte` | Min/max byte bounds |

## Within — Range Validation

The `Within` variable provides boundary validation and clamping:

```go
// Validate and clamp integer to range
val, inRange := corerange.Within.RangeInteger(true, 0, 100, input)

// Parse string to int32 with range check
val, inRange := corerange.Within.StringRangeInt32("42")

// Parse string to float with range
val, inRange := corerange.Within.StringRangeFloat(true, 0.0, 1.0, "0.5")

// Byte range validation
val, inRange := corerange.Within.RangeByteDefault(200)
```

### Within Methods

| Method | Description |
|--------|-------------|
| `RangeInteger(boundary, min, max, input)` | Validate int against range, optionally clamp |
| `RangeDefaultInteger(min, max, input)` | Integer range with clamping enabled |
| `RangeByte(boundary, input)` | Validate byte (0–255) |
| `RangeByteDefault(input)` | Byte range with clamping |
| `RangeUint16(boundary, input)` | Validate uint16 |
| `RangeFloat(boundary, min, max, input)` | Validate float32 |
| `RangeFloat64(boundary, min, max, input)` | Validate float64 |
| `StringRangeInt8(input)` | Parse string → int8 with bounds |
| `StringRangeInt16(input)` | Parse string → int16 with bounds |
| `StringRangeInt32(input)` | Parse string → int32 with bounds |
| `StringRangeByte(input)` | Parse string → byte with bounds |
| `StringRangeUint16(input)` | Parse string → uint16 with bounds |
| `StringRangeUint32(input)` | Parse string → uint32 with bounds |
| `StringRangeInteger(boundary, min, max, input)` | Parse string → int with custom range |
| `StringRangeFloat(boundary, min, max, input)` | Parse string → float32 with range |
| `StringRangeFloat64(boundary, min, max, input)` | Parse string → float64 with range |

### Boundary Behavior

When `isUsageMinMaxBoundary` is `true`, out-of-range values are clamped to the nearest bound. When `false`, the original value is returned with `isInRange = false`.

## How to Extend Safely

- Add new `RangeX` / `MinMaxX` types following existing patterns.
- Add new `Within` methods for additional type conversions.
- Keep boundary clamping behavior consistent.

## Contributors

## Issues for Future Reference
