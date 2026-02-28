# coregeneric Architecture

## Overview

The `coredata/coregeneric` package provides **type-parameterized** (generic) versions of all core collection data structures in the auk-go ecosystem. It serves as the foundational layer that type-specific packages like `corestr` build upon.

## Motivation

Previously, each data type required hand-written, type-specific implementations:
- `corestr.Collection` → `[]string` only
- `corestr.Hashset` → `map[string]bool` only
- `coredynamic.Collection[T]` → generic but limited to that package

`coregeneric` consolidates all generic data structures into one package, providing:
1. **Code reuse**: One implementation serves all types
2. **Consistency**: Same API across `Collection[int]`, `Collection[string]`, `Collection[float64]`
3. **Type safety**: Compile-time type checking via Go generics
4. **IDE discoverability**: Full New Creator pattern support

## Generic Types

| Type | Constraint | Replaces |
|------|-----------|----------|
| `Collection[T]` | `any` | `corestr.Collection`, `coredynamic.Collection[T]` |
| `Hashset[T]` | `comparable` | `corestr.Hashset` |
| `Hashmap[K, V]` | `K: comparable, V: any` | `corestr.Hashmap` |
| `SimpleSlice[T]` | `any` | `corestr.SimpleSlice` |
| `LinkedList[T]` | `any` | `corestr.LinkedList` |

## Typed Creator Innovation

The key architectural innovation is the `typedXCreator[T]` pattern. Instead of writing 16 separate creator structs (one per primitive type), a single generic struct handles all:

```go
// ONE definition covers ALL types
type typedCollectionCreator[T any] struct{}

func (it typedCollectionCreator[T]) Empty() *Collection[T] { ... }
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] { ... }
func (it typedCollectionCreator[T]) From(items []T) *Collection[T] { ... }
func (it typedCollectionCreator[T]) Clone(items []T) *Collection[T] { ... }
func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] { ... }

// Instantiated per-type via struct fields
type newCollectionCreator struct {
    String  typedCollectionCreator[string]
    Int     typedCollectionCreator[int]
    Int64   typedCollectionCreator[int64]
    Float64 typedCollectionCreator[float64]
    // ... all 16 types
}
```

This eliminates massive code duplication while preserving the New Creator pattern's IDE discoverability.

## Generic Functions (funcs.go)

Go does not allow additional type parameters on methods of generic types. Therefore, cross-type transformations are package-level functions:

```go
// Transform Collection[T] → Collection[U]
result := coregeneric.MapCollection(source, func(item string) int { ... })

// Reduce Collection[T] → U
sum := coregeneric.ReduceCollection(ints, 0, func(acc, item int) int { return acc + item })

// Group Collection[T] by key K
groups := coregeneric.GroupByCollection(items, func(item T) K { ... })

// Deduplicate (requires comparable)
unique := coregeneric.Distinct(collection)
```

## Migration Path: corestr → coregeneric

`corestr` retains its public API for backward compatibility. String-specific methods (Join, EqualFold, Trim, Split, etc.) remain in `corestr`. Common operations can internally delegate to `coregeneric`:

```
Phase 1 (current): coregeneric created alongside corestr — both exist independently
Phase 2 (future):  corestr types embed/delegate to coregeneric internally
Phase 3 (future):  corestr becomes a thin wrapper around coregeneric.Collection[string]
```

## Related Docs

- [New Creator Pattern](./21-new-creator-pattern.md)
- [coredata Overview](./folders/05-coredata.md)
- [coregeneric Folder Spec](./folders/05a-coregeneric.md)
