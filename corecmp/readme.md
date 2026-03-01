# corecmp

## Overview

Core comparison functions that produce `corecomparator.Compare` results for built-in Go types. Each function follows a consistent pattern: compare two values and return a typed comparison result.

## Functions

### Primitive Comparators

| Function                  | Signature                                        | Description                              |
|---------------------------|--------------------------------------------------|------------------------------------------|
| `Integer(left, right)`    | `(int, int) → Compare`                          | Compares two `int` values                |
| `Integer8(left, right)`   | `(int8, int8) → Compare`                        | Compares two `int8` values               |
| `Integer16(left, right)`  | `(int16, int16) → Compare`                      | Compares two `int16` values              |
| `Integer32(left, right)`  | `(int32, int32) → Compare`                      | Compares two `int32` values              |
| `Integer64(left, right)`  | `(int64, int64) → Compare`                      | Compares two `int64` values              |
| `Byte(left, right)`       | `(byte, byte) → Compare`                        | Compares two `byte` values               |
| `Time(left, right)`       | `(time.Time, time.Time) → Compare`              | Compares two `time.Time` values          |

### Pointer Comparators

| Function                  | Signature                                        | Description                              |
|---------------------------|--------------------------------------------------|------------------------------------------|
| `IntegerPtr(left, right)` | `(*int, *int) → Compare`                        | Nil-safe pointer comparison for `int`    |
| `Integer8Ptr(left, right)`| `(*int8, *int8) → Compare`                      | Nil-safe pointer comparison for `int8`   |
| `Integer16Ptr(left, right)`| `(*int16, *int16) → Compare`                   | Nil-safe pointer comparison for `int16`  |
| `Integer32Ptr(left, right)`| `(*int32, *int32) → Compare`                   | Nil-safe pointer comparison for `int32`  |
| `Integer64Ptr(left, right)`| `(*int64, *int64) → Compare`                   | Nil-safe pointer comparison for `int64`  |
| `BytePtr(left, right)`    | `(*byte, *byte) → Compare`                      | Nil-safe pointer comparison for `byte`   |
| `TimePtr(left, right)`    | `(*time.Time, *time.Time) → Compare`            | Nil-safe pointer comparison for `time`   |

### Any-Type Comparator

| Function                  | Signature                                        | Description                              |
|---------------------------|--------------------------------------------------|------------------------------------------|
| `AnyItem(left, right)`    | `(any, any) → Compare`                          | Interface equality using `==` operator   |

#### `AnyItem` Comparison Logic

1. Both `nil` → `Equal`
2. One `nil`, other not → `NotEqual`
3. `left == right` → `Equal` (works for comparable built-in types)
4. Otherwise → `Inconclusive` (types may not be deeply comparable via `==`)

> **Note:** `AnyItem` does NOT perform deep comparison. Slices, maps, and structs without comparable implementation return `Inconclusive` when they differ. Use `reflect.DeepEqual` or typed comparators for those cases.

### Equality Checks

| Function                          | Signature                          | Description                              |
|-----------------------------------|------------------------------------|------------------------------------------|
| `IsIntegersEqual(left, right)`    | `(int, int) → bool`               | Shorthand equality check                 |
| `IsIntegersEqualPtr(left, right)` | `(*int, *int) → bool`             | Nil-safe pointer equality                |
| `IsStringsEqual(left, right)`     | `(string, string) → bool`         | String equality check                    |
| `IsStringsEqualPtr(left, right)`  | `(*string, *string) → bool`       | Nil-safe string pointer equality         |
| `IsStringsEqualWithoutOrder(...)`  | `([]string, []string) → bool`    | Order-independent string slice equality  |

### Version Comparators

| Function                          | Signature                          | Description                              |
|-----------------------------------|------------------------------------|------------------------------------------|
| `VersionSliceByte(left, right)`   | `([]byte, []byte) → Compare`      | Byte-slice version comparison            |
| `VersionSliceInteger(left, right)`| `([]int, []int) → Compare`        | Integer-slice version comparison         |

## Return Values

All comparator functions return `corecomparator.Compare`:
- `Equal` — values are equal
- `LeftGreater` — left value is greater
- `LeftLess` — left value is less
- `NotEqual` — values are not equal (pointer: one nil)
- `Inconclusive` — comparison cannot be determined

## Usage

```go
import (
    "gitlab.com/auk-go/core/corecmp"
    "gitlab.com/auk-go/core/corecomparator"
)

result := corecmp.Integer(5, 3)
if result.IsLeftGreater() {
    // 5 > 3
}

ptrResult := corecmp.IntegerPtr(&a, nil)
if ptrResult.IsNotEqual() {
    // one is nil
}

anyResult := corecmp.AnyItem("hello", "hello")
if anyResult.IsEqual() {
    // equal
}
```

## Related Docs

- [corecomparator readme](../corecomparator/readme.md)
- [Comparison & Sorting spec](../spec/01-app/folders/10-remaining-packages.md)
