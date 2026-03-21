# coremath — Numeric Utilities

Package `coremath` provides min/max functions for numeric types and range-boundary checking utilities.

## API

### Range Checking (Primary API)

```go
import "github.com/alimtvnetwork/core/coremath"

// Check if a value is within a range
coremath.IsRangeWithin.Integer(5, 1, 10)   // true
coremath.IsRangeWithin.Integer64(100, 0, 50) // false

// Check if a value is out of range
coremath.IsOutOfRange.Integer(15, 1, 10)   // true
coremath.IsOutOfRange.Integer64(5, 0, 50)  // false
```

### Min/Max Functions

> **Note:** For `byte` and `float32`, deprecated wrappers exist. Prefer Go's built-in `min()`/`max()` (Go 1.21+).

```go
// Use built-in min/max for most types (Go 1.21+)
result := max(a, b)

// Package provides MaxInt/MinInt for legacy compatibility
coremath.MaxInt(3, 7) // 7
coremath.MinInt(3, 7) // 3
```

## Files

| File | Purpose |
|------|---------|
| `isRangeWithin.go` | `IsRangeWithin` struct-as-namespace |
| `isOutOfRange.go` | `IsOutOfRange` struct-as-namespace |
| `integer*Within.go` | Range-within checks for int, int16, int32, int64 |
| `integer*OutOfRange.go` | Out-of-range checks for int, int64 |
| `unsigned*Within.go` | Range checks for unsigned int16 |
| `MaxInt.go` / `MinInt.go` | Int min/max |
| `vars.go` | Package-level singleton variables |

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
