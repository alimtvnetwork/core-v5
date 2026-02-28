# coregeneric — Generic Data Structures

## Overview

The `coregeneric` package provides **type-parameterized** versions of all core collection data structures. It serves as the foundational layer that type-specific packages like `corestr` build upon, offering code reuse, consistency, compile-time type safety, and IDE discoverability via the New Creator pattern.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `New` | `*newCreator` | Root aggregator for the New Creator pattern |

## Generic Types

| Type | Constraint | Description |
|------|-----------|-------------|
| `Collection[T]` | `any` | Thread-safe slice-backed list with Add/Remove/Filter/Sort/Clone and `*Lock` variants |
| `Hashset[T]` | `comparable` | Thread-safe set backed by `map[T]bool` with Add/Has/Remove/Resize |
| `Hashmap[K, V]` | `K: comparable, V: any` | Thread-safe map wrapper with Set/Get/Has/Remove/Keys/Values |
| `SimpleSlice[T]` | `any` | Lightweight typed slice wrapper with Add/Filter/Skip/Take |
| `LinkedList[T]` | `any` | Singly-linked list with head/tail pointers and embedded mutex |

## Typed Creator Pattern

Instead of writing 16 separate creator structs, a single generic struct handles all primitive types:

```go
// One generic definition covers ALL types
type typedCollectionCreator[T any] struct{}
func (it typedCollectionCreator[T]) Empty() *Collection[T] { ... }
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] { ... }
func (it typedCollectionCreator[T]) From(items []T) *Collection[T] { ... }
func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] { ... }

// Instantiated per-type via struct fields
type newCollectionCreator struct {
    String  typedCollectionCreator[string]
    Int     typedCollectionCreator[int]
    Float64 typedCollectionCreator[float64]
    // ... all 16 primitive types
}
```

### Usage Examples

```go
col := coregeneric.New.Collection.String.Cap(10)
set := coregeneric.New.Hashset.Int.Items(1, 2, 3)
hm  := coregeneric.New.Hashmap.StringString.Cap(20)
ss  := coregeneric.New.SimpleSlice.Float64.Items(1.0, 2.5)
ll  := coregeneric.New.LinkedList.String.Empty()
```

## Generic Functions

Go does not allow additional type parameters on methods of generic types. Cross-type transformations are package-level functions in `funcs.go`, `comparablefuncs.go`, and `orderedfuncs.go`:

| Function | Constraint | Description |
|----------|-----------|-------------|
| `MapCollection[T, U]` | `any` | Transform `Collection[T]` → `Collection[U]` |
| `FlatMapCollection[T, U]` | `any` | Map + flatten `Collection[T]` → `Collection[U]` |
| `ReduceCollection[T, U]` | `any` | Reduce `Collection[T]` → `U` |
| `GroupByCollection[T, K]` | `K: comparable` | Group into `map[K]*Collection[T]` |
| `Distinct[T]` | `comparable` | Deduplicate collection |
| `ContainsItem[T]` | `comparable` | Check if item exists |
| `ContainsAll[T]` | `comparable` | Check all items exist |
| `ToHashset[T]` | `comparable` | Convert collection to hashset |
| `Sort[T]` | `cmp.Ordered` | Sort collection ascending |
| `SortDesc[T]` | `cmp.Ordered` | Sort collection descending |
| `Min[T]` / `Max[T]` | `cmp.Ordered` | Find min/max element |
| `Sum[T]` | `cmp.Ordered` | Sum all elements |

## Type Aliases

Pre-defined aliases for all common primitives (e.g., `StringCollection`, `IntHashset`, `StringStringHashmap`) are in `types.go`.

## Supported Primitive Types

| Category | Types |
|----------|-------|
| Signed integers | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned integers | `uint`, `uint8`, `uint16`, `uint32`, `uint64` |
| Floats | `float32`, `float64` |
| Other | `byte`, `bool`, `string`, `any` |

## File Organization

| File | Responsibility |
|------|---------------|
| `vars.go` | `New = &newCreator{}` |
| `newCreator.go` | Root aggregator struct |
| `new*Creator.go` | Sub-creators for each data structure |
| `typed*Creator.go` | Generic typed creator implementations |
| `Collection.go` | `Collection[T]` type + methods |
| `Hashset.go` | `Hashset[T]` type + methods |
| `Hashmap.go` | `Hashmap[K,V]` type + methods |
| `SimpleSlice.go` | `SimpleSlice[T]` type + methods |
| `LinkedList.go` / `LinkedListNode.go` | LinkedList types |
| `types.go` | Type aliases for all primitives |
| `funcs.go` | Generic cross-type functions (any constraint) |
| `comparablefuncs.go` | Functions requiring `comparable` constraint |
| `orderedfuncs.go` | Functions requiring `cmp.Ordered` constraint |

## Key Patterns

- All types use embedded `sync.Mutex` for thread-safe `*Lock` variants.
- Zero-nil safety: nil receivers return safe defaults (0, empty, false).
- Constraint hierarchy: `cmp.Ordered ⊂ comparable ⊂ any`.

## How to Extend Safely

- **New primitive type**: Add a field to each `newXCreator` struct and a type alias in `types.go`.
- **New data structure**: Create the generic type, a `typedXCreator[T]`, and a `newXCreator` aggregator. Add the field to `newCreator`.
- **New generic function**: Add to `funcs.go` / `comparablefuncs.go` / `orderedfuncs.go` based on constraint.

## Contributors

## Issues for Future Reference
