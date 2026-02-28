# corecomparator — Core Comparison Abstraction

## Overview

Package `corecomparator` provides the `Compare` enum type — a `byte`-backed comparison result used throughout the auk-go ecosystem. It includes 7 comparison outcomes, logical grouping methods, operator symbols, SQL operators, JSON marshalling, and embeddable case-sensitivity structs.

## Type: `Compare`

A `byte`-backed enum representing the result of comparing two values.

### Constants

| Constant | Value | Symbol | Short | SQL | Meaning |
|----------|-------|--------|-------|-----|---------|
| `Equal` | 0 | `=` | `eq` | `=` | Left and right are equal |
| `LeftGreater` | 1 | `>` | `gt` | `>` | Left is strictly greater |
| `LeftGreaterEqual` | 2 | `>=` | `ge` | `>=` | Left is greater than or equal |
| `LeftLess` | 3 | `<` | `lt` | `<` | Left is strictly less |
| `LeftLessEqual` | 4 | `<=` | `le` | `<=` | Left is less than or equal |
| `NotEqual` | 5 | `!=` | `ne` | `<>` | Values are not equal |
| `Inconclusive` | 6 | `?!` | `i` | `i` | Comparison could not be determined |

### Identity & Validity Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `Is(other)` | `bool` | Exact match |
| `IsEqual()` | `bool` | True if `Equal` |
| `IsLeftGreater()` | `bool` | True if `LeftGreater` |
| `IsLeftGreaterEqual()` | `bool` | True if `LeftGreaterEqual` |
| `IsLeftLess()` | `bool` | True if `LeftLess` |
| `IsLeftLessEqual()` | `bool` | True if `LeftLessEqual` |
| `IsNotEqual()` | `bool` | True if `NotEqual` |
| `IsInconclusive()` | `bool` | True if `Inconclusive` |
| `IsNotInconclusive()` | `bool` | True if not `Inconclusive` |
| `IsDefined()` / `IsValid()` / `IsDefinedProperly()` | `bool` | True if not `Inconclusive` |
| `IsInvalid()` | `bool` | True if `Inconclusive` |

### Logical Grouping Methods

| Method | Returns | True When |
|--------|---------|-----------|
| `IsLeftGreaterEqualLogically()` | `bool` | `Equal`, `LeftGreater`, or `LeftGreaterEqual` |
| `IsLeftLessEqualLogically()` | `bool` | `Equal`, `LeftLess`, or `LeftLessEqual` |
| `IsNotEqualLogically()` | `bool` | Anything except `Equal` |
| `IsLeftGreaterOrGreaterEqualOrEqual()` | `bool` | Same as `IsLeftGreaterEqualLogically` |
| `IsLeftLessOrLessEqualOrEqual()` | `bool` | Same as `IsLeftLessEqualLogically` |
| `IsInconclusiveOrNotEqual()` | `bool` | `Inconclusive` or `NotEqual` |
| `IsDefinedPlus(right)` | `bool` | Not `Inconclusive` AND equals `right` |
| `IsCompareEqualLogically(expected)` | `bool` | Semantic equality with logical expansion |

### Matching Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `IsAnyOf(...Compare)` | `bool` | True if matches any of the given values (empty = true) |
| `IsAnyNamesOf(...string)` | `bool` | True if name matches any given string |
| `IsNameEqual(name)` | `bool` | True if `Name()` equals the given string |
| `IsValueEqual(value)` | `bool` | True if byte value matches |
| `IsGreater()` | `bool` | Shorthand: `LeftGreater` |
| `IsGreaterEqual()` | `bool` | Shorthand: `LeftGreater` or `Equal` |
| `IsLess()` | `bool` | Shorthand: `LeftLess` |
| `IsLessEqual()` | `bool` | Shorthand: `LeftLess` or `Equal` |

### Formatting & Serialization

| Method | Returns | Description |
|--------|---------|-------------|
| `Name()` / `String()` | `string` | Human-readable name (e.g., `"Equal"`) |
| `NameValue()` | `string` | `"Equal(0)"` format |
| `OperatorSymbol()` | `string` | `=`, `>`, `<`, `>=`, `<=`, `!=`, `?!` |
| `OperatorShortForm()` | `string` | `eq`, `gt`, `lt`, `ge`, `le`, `ne`, `i` |
| `SqlOperatorSymbol()` | `string` | SQL operator (`<>` for NotEqual) |
| `ToNumberString()` / `NumberString()` / `ValueString()` | `string` | Numeric string representation |
| `NumberJsonString()` | `string` | JSON-quoted number string |
| `Value()` / `ValueByte()` | `byte` | Raw byte value |
| `ValueInt()` / `ValueInt8()` / `ValueInt16()` / `ValueInt32()` | various | Typed numeric value |
| `StringValue()` | `string` | Raw string conversion of byte |
| `MarshalJSON()` | `([]byte, error)` | JSON serialization by name |
| `UnmarshalJSON([]byte)` | `error` | JSON deserialization from name, number, JSON-number, operator, or short form |
| `RangeNamesCsv()` | `string` | CSV of all range names with indexes |
| `CsvStrings(...Compare)` | `[]string` | Slice of `NameValue()` for given values |
| `CsvString(...Compare)` | `string` | Joined CSV of `NameValue()` for given values |
| `OnlySupportedErr(msg, ...Compare)` | `error` | Error if value not in supported set (with message) |
| `OnlySupportedDirectErr(...Compare)` | `error` | Error if value not in supported set |

### Validation

| Method | Returns | Description |
|--------|---------|-------------|
| `OnlySupportedErr(message, ...Compare)` | `error` | Returns error if current value is not in the supported set; nil otherwise |
| `OnlySupportedDirectErr(...Compare)` | `error` | Same without message prefix |

## Supporting Types

### `BaseIsCaseSensitive`

Embeddable struct with `IsCaseSensitive bool` field. Provides `IsIgnoreCase()`, conversion to `BaseIsIgnoreCase`, and `Clone()`/`ClonePtr()`.

### `BaseIsIgnoreCase`

Embeddable struct with `IsIgnoreCase bool` field. Provides `IsCaseSensitive()`, conversion to `BaseIsCaseSensitive`, and `Clone()`/`ClonePtr()`.

### Package-Level Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `Min()` | `Compare` | Returns `Equal` (minimum valid value) |
| `Max()` | `Compare` | Returns `NotEqual` (maximum valid value, excluding `Inconclusive`) |
| `MinLength(left, right)` | `int` | Returns the smaller of two integers |
| `Ranges()` | `[]string` | All `CompareNames` as a string slice |
| `RangeNamesCsv()` | `string` | CSV representation of all range names with indexes |

### Package Variables

| Variable | Type | Description |
|----------|------|-------------|
| `CompareNames` | `[7]string` | Name lookup array |
| `CompareOperatorsSymbols` | `[7]string` | Operator symbol lookup (`=`, `>`, etc.) |
| `CompareOperatorsShotNames` | `[7]string` | Short form lookup (`eq`, `gt`, etc.) |
| `SqlCompareOperators` | `[7]string` | SQL operator lookup (`<>` for NotEqual) |
| `RangesMap` | `map[string]Compare` | Reverse lookup: name/number/operator/short → `Compare` |

## File Organization

| File | Purpose |
|------|---------|
| `Compare.go` | `Compare` type definition, all methods, JSON support |
| `consts.go` | Package-level constants |
| `vars.go` | `CompareNames`, `CompareOperatorsSymbols`, `SqlCompareOperators`, `RangesMap` |
| `Ranges.go` | `Ranges()` — full range of Compare values as strings |
| `RangeNamesCsv.go` | `RangeNamesCsv()` — CSV name representation |
| `BaseIsCaseSensitive.go` | Embeddable case-sensitive flag struct |
| `BaseIsIgnoreCase.go` | Embeddable case-insensitive flag struct |
| `Min.go` | `Min()` → `Equal` |
| `Max.go` | `Max()` → `NotEqual` |
| `MinLength.go` | `MinLength(left, right int)` helper |

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

// Operator symbols for display
fmt.Println(result.OperatorSymbol()) // "=", ">", "<", etc.

// SQL operators
fmt.Println(result.SqlOperatorSymbol()) // "=", "<>", etc.

// Validation
err := result.OnlySupportedErr("sort direction", corecomparator.LeftGreater, corecomparator.LeftLess)
```

## Related Docs

- [corecmp Readme](../corecmp/Readme.md)
- [anycmp Readme](../anycmp/README.md)
- [Comparison & Sorting spec](../spec/01-app/folders/10-remaining-packages.md)
- [Folder Map](../spec/01-app/01-folder-map.md)
