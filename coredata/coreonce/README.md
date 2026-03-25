# coreonce — Compute-Once Value Holders

## Overview

The `coreonce` package provides lazy-initialization wrappers that compute their value exactly once on first access and cache the result for all subsequent calls. Each type wraps a specific Go type with an initializer function.

## Types

| Type | Wrapped Type | Description |
|------|-------------|-------------|
| `AnyOnce` | `interface{}` | General-purpose once-holder with cast helpers and JSON serialization |
| `AnyErrorOnce` | `(interface{}, error)` | Once-holder that also captures an error |
| `StringOnce` | `string` | String once-holder with HasPrefix, HasSuffix, Contains, Split, and JSON support |
| `StringsOnce` | `[]string` | String slice once-holder |
| `BoolOnce` | `bool` | Boolean once-holder |
| `ByteOnce` | `byte` | Single byte once-holder |
| `BytesOnce` | `[]byte` | Byte slice once-holder |
| `BytesErrorOnce` | `([]byte, error)` | Byte slice once-holder with error |
| `ErrorOnce` | `error` | Error once-holder |
| `IntegerOnce` | `int` | Integer once-holder |
| `IntegersOnce` | `[]int` | Integer slice once-holder |
| `MapStringStringOnce` | `map[string]string` | Map once-holder |

## Usage Pattern

All types follow the same pattern:

```go
// Create with initializer function
once := coreonce.NewStringOnce(func() string {
    return expensiveComputation()
})

// First call executes the function
val := once.Value() // computes and caches

// Subsequent calls return cached value
val2 := once.Value() // returns cached, no recomputation
```

## Key Methods (common across types)

| Method | Description |
|--------|-------------|
| `Value()` | Returns the cached value, initializing on first call |
| `Execute()` | Alias for `Value()` |
| `IsInitialized()` | Whether the value has been computed |
| `String()` | String representation |
| `Serialize()` / `Deserialize()` | JSON marshalling support |
| `MarshalJSON()` / `UnmarshalJSON()` | Standard JSON interface |

### StringOnce extras

| Method | Description |
|--------|-------------|
| `IsEmpty()` / `IsEmptyOrWhitespace()` | Blank checks |
| `HasPrefix(s)` / `HasSuffix(s)` | String matching |
| `IsContains(s)` | Substring check |
| `SplitBy(sep)` | Split into `[]string` |
| `SplitLeftRight(sep)` | Split into left/right pair |
| `Bytes()` | Convert to `[]byte` |
| `Error()` | Convert to `error` |

### AnyOnce extras

| Method | Description |
|--------|-------------|
| `CastValueString()` | Type-assert to `string` |
| `CastValueStrings()` | Type-assert to `[]string` |
| `CastValueBytes()` | Type-assert to `[]byte` |
| `CastValueHashmapMap()` | Type-assert to `map[string]string` |
| `IsNull()` | Whether the value is nil |

## How to Extend Safely

- Add a new `XOnce` struct following the same pattern: `innerData`, `initializerFunc`, `isInitialized` fields.
- Provide both `NewXOnce` (value) and `NewXOncePtr` (pointer) constructors.
- Implement `Value()`, `Execute()`, `String()`, and JSON interfaces.

## Contributors

## Issues for Future Reference
