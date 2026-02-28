# corecomparator

## Overview

Core comparison abstraction providing the `Compare` enum type used throughout the auk-go ecosystem for expressing relational comparison results.

## Type: `Compare`

A `byte`-backed enum representing the result of comparing two values.

### Values

| Constant           | Value | Meaning                           |
|--------------------|-------|-----------------------------------|
| `Equal`            | 0     | Left and right are equal          |
| `LeftGreater`      | 1     | Left is strictly greater          |
| `LeftGreaterEqual` | 2     | Left is greater than or equal     |
| `LeftLess`         | 3     | Left is strictly less             |
| `LeftLessEqual`    | 4     | Left is less than or equal        |
| `NotEqual`         | 5     | Values are not equal              |
| `Inconclusive`     | 6     | Comparison could not be determined|

### Key Methods

| Method                        | Returns  | Description                                         |
|-------------------------------|----------|-----------------------------------------------------|
| `IsEqual()`                   | `bool`   | True if result is `Equal`                            |
| `IsLeftGreater()`             | `bool`   | True if result is `LeftGreater`                      |
| `IsLeftLess()`                | `bool`   | True if result is `LeftLess`                         |
| `IsInconclusive()`            | `bool`   | True if result is `Inconclusive`                     |
| `IsLeftGreaterEqualLogically()` | `bool` | True if `Equal` or `LeftGreater` or `LeftGreaterEqual` |
| `IsLeftLessEqualLogically()`  | `bool`   | True if `Equal` or `LeftLess` or `LeftLessEqual`     |
| `IsNotEqualLogically()`       | `bool`   | True if result is anything except `Equal`            |
| `IsAnyOf(...Compare)`         | `bool`   | True if result matches any of the given values       |
| `IsCompareEqualLogically(Compare)` | `bool` | Semantic equality check with logical expansion  |
| `OperatorSymbol()`            | `string` | Returns operator symbol (e.g., `==`, `>`, `<`)      |
| `Name()`                      | `string` | Human-readable name                                  |
| `MarshalJSON()`               | `[]byte` | JSON serialization by name                           |
| `UnmarshalJSON([]byte)`       | `error`  | JSON deserialization from name                       |

### Supporting Types

| File                      | Type / Purpose                                    |
|---------------------------|---------------------------------------------------|
| `BaseIsCaseSensitive.go`  | Embeddable struct for case-sensitive flag          |
| `BaseIsIgnoreCase.go`     | Embeddable struct for case-insensitive flag        |
| `Min.go` / `Max.go`       | Min/Max value helpers                             |
| `MinLength.go`            | Minimum length constraint                          |
| `Ranges.go`               | Full range of `Compare` values                    |
| `RangeNamesCsv.go`        | CSV representation of all range names             |

## Usage

```go
import "gitlab.com/auk-go/core/corecomparator"

result := someCompareFunc(a, b)

if result.IsEqual() {
    // ...
}

if result.IsLeftGreaterEqualLogically() {
    // includes Equal, LeftGreater, and LeftGreaterEqual
}
```

## Related Docs

- [corecmp Readme](../corecmp/Readme.md)
- [Comparison & Sorting spec](../../spec/01-app/folders/10-remaining-packages.md)
