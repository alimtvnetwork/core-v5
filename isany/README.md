# isany — Reflection-Based Type & Null Checking

## Overview

The `isany` package provides reflection-based predicate functions for checking nullability, type identity, equality, and value state of `any`-typed values. It is the foundational type-inspection layer used by `anycmp`, `corecomparator`, `corecmp`, and assertion frameworks throughout the codebase.

## Null & Defined Checks

### Single-Value Checks

| Function | Description |
|----------|-------------|
| `Null(item any) bool` | Returns `true` for `nil` or nil-able kinds (chan, func, map, ptr, slice) that are nil |
| `NotNull(item any) bool` | Inverse of `Null` |
| `Defined(item any) bool` | Alias for `NotNull` — returns `true` for non-nil values |
| `ReflectNull(rv reflect.Value) bool` | `Null` check on a `reflect.Value` directly |
| `ReflectNotNull(rv reflect.Value) bool` | Inverse of `ReflectNull` |
| `ReflectValueNull(rv reflect.Value) bool` | Nil check on `reflect.Value` for nil-able kinds |

### Pair Checks

| Function | Description |
|----------|-------------|
| `NullLeftRight(left, right any) (isLeftNull, isRightNull bool)` | Check both values for null simultaneously |
| `NullBoth(left, right any) bool` | Returns `true` only if **both** are null |
| `DefinedBoth(left, right any) bool` | Returns `true` only if **both** are defined |
| `DefinedLeftRight(left, right any) (isLeftDefined, isRightDefined bool)` | Check both for defined |

### Batch Checks

| Function | Description |
|----------|-------------|
| `AllNull(items ...any) bool` | Returns `true` if **all** items are null (false for empty) |
| `AnyNull(items ...any) bool` | Returns `true` if **any** item is null |
| `DefinedAllOf(items ...any) bool` | Returns `true` if **all** items are defined |
| `DefinedAnyOf(items ...any) bool` | Returns `true` if **any** item is defined |
| `DefinedItems(items ...any) (allDefined bool, definedItems []any)` | Filter and return only defined items |

## Zero & Value Checks

| Function | Description |
|----------|-------------|
| `Zero(item any) bool` | Returns `true` if null or reflect zero value |
| `AllZero(items ...any) bool` | Returns `true` if **all** items are zero |
| `AnyZero(items ...any) bool` | Returns `true` if **any** item is zero |

## Equality Checks

| Function | Description |
|----------|-------------|
| `DeepEqual(left, right any) bool` | Reflection-based deep equality via `reflectinternal` |
| `DeepEqualAllItems(items ...any) bool` | Returns `true` if all items are deeply equal |
| `NotDeepEqual(left, right any) bool` | Inverse of `DeepEqual` |
| `JsonEqual(left, right any) bool` | Equality via JSON marshal comparison (fast-path for strings/ints) |
| `JsonMismatch(left, right any) bool` | Inverse of `JsonEqual` |
| `StringEqual(left, right any) bool` | Equality via string conversion |
| `Conclusive(left, right any) (isEqual, isConclusive bool)` | Prioritized nil/pointer check returning conclusive/inconclusive result |

## Type Checks

| Function | Description |
|----------|-------------|
| `TypeSame(left, right any) bool` | Returns `true` if both values have the same `reflect.Type` |
| `Pointer(item any) bool` | Returns `true` if the value is a pointer |
| `Function(item any) (isFunc bool, name string)` | Returns `true` + function name if value is a function |
| `FuncOnly(item any) bool` | Returns `true` if value is a function (no name extraction) |
| `PrimitiveType(item any) bool` | Returns `true` for bool, int*, uint*, float*, string |
| `PrimitiveTypeRv(kind reflect.Kind) bool` | `PrimitiveType` on a `reflect.Kind` directly |
| `NumberType(item any) bool` | Returns `true` for int*, uint*, float* types |
| `NumberTypeRv(rv reflect.Value) bool` | `NumberType` on a `reflect.Value` directly |
| `FloatingPointType(item any) bool` | Returns `true` for float32, float64 |
| `FloatingPointTypeRv(rv reflect.Value) bool` | On a `reflect.Value` directly |
| `PositiveIntegerType(item any) bool` | Returns `true` for uint* types |

## Usage Examples

```go
// Null checking
if isany.Null(someValue) {
    // handle nil
}

// Batch defined check
if isany.DefinedAllOf(a, b, c) {
    // all three are non-nil
}

// Type inspection
if isany.PrimitiveType(val) {
    // val is bool, int, float, string, etc.
}

// Conclusive comparison
isEqual, isConclusive := isany.Conclusive(left, right)
if isConclusive {
    // nil/pointer check was sufficient
} else {
    // need deeper comparison (DeepEqual, JsonEqual, etc.)
}

// JSON equality for mixed types
if isany.JsonEqual(mapA, mapB) {
    // structurally equivalent when serialized
}
```

## File Organization

| File | Responsibility |
|------|---------------|
| `Null.go` | Core `Null` check with reflect nil-able kind handling |
| `NotNull.go` | Inverse of `Null` |
| `Defined.go` | Alias for `NotNull` |
| `DefinedBoth.go`, `DefinedLeftRight.go` | Pair-wise defined checks |
| `DefinedAllOf.go`, `DefinedAnyOf.go` | Batch defined checks |
| `DefinedItems.go` | Filter to only defined items |
| `AllNull.go`, `AnyNull.go` | Batch null checks |
| `NullBoth.go`, `NullLeftRight.go` | Pair-wise null checks |
| `ReflectNull.go`, `ReflectNotNull.go`, `ReflectValueNull.go` | Reflect-level null checks |
| `Zero.go`, `AllZero.go`, `AnyZero.go` | Zero-value checks |
| `DeepEqual.go`, `NotDeepEqual.go`, `DeepEqualAllItems.go` | Deep equality |
| `JsonEqual.go`, `JsonMismatch.go` | JSON-based equality |
| `StringEqual.go` | String-conversion equality |
| `Conclusive.go` | Prioritized conclusive comparison |
| `TypeSame.go` | Type identity check |
| `Pointer.go` | Pointer type check |
| `Function.go`, `FuncOnly.go` | Function type detection |
| `PrimitiveType.go`, `PrimitiveTypeRv.go` | Primitive type detection |
| `NumberType.go`, `NumberTypeRv.go` | Numeric type detection |
| `FloatingPointType.go`, `FloatingPointTypeRv.go` | Float type detection |
| `PositiveIntegerType.go` | Unsigned integer detection |

## Design Principles

- **Nil-safe**: All functions handle `nil` inputs gracefully without panicking.
- **Reflect-aware**: Uses `reflect.Value.Kind()` to detect nil-able types (chan, func, map, ptr, slice) that Go's `==` operator cannot check.
- **Layered**: `Null` is the foundation — `Defined`, `NullBoth`, `AllNull`, etc. compose on top of it.
- **Fast-path**: Functions like `JsonEqual` optimize for common types (string, int) before falling back to marshal comparison.

## How to Extend Safely

- **New predicate**: Create a dedicated file (`IsXxx.go`) with a single exported function.
- **New batch variant**: Follow the `All*` / `Any*` naming convention.
- **Reflect-level variant**: Add an `*Rv` suffixed function that accepts `reflect.Value`.

## Related Docs

- [Repo Overview](../spec/01-app/00-repo-overview.md)
- [anycmp package](../anycmp/)
