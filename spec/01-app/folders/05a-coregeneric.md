# coregeneric

## Folder Purpose

Generic data structures package providing type-parameterized versions of all core collection types. This is the foundational layer that `corestr` and other type-specific packages build upon.

## Responsibilities

1. **Collection[T]** — Generic slice-backed collection with embedded mutex, supporting Add/Remove/Filter/Sort/Clone and thread-safe `*Lock` variants.
2. **Hashset[T comparable]** — Generic set backed by `map[T]bool` with Add/Has/Remove/Resize and thread-safe variants.
3. **Hashmap[K comparable, V any]** — Generic map wrapper with Set/Get/Has/Remove/Keys/Values and thread-safe variants.
4. **SimpleSlice[T]** — Thin generic slice wrapper (`[]T`) with Add/Filter/Skip/Take convenience methods.
5. **LinkedList[T]** — Generic singly-linked list with head/tail pointers, embedded mutex, and full traversal support.
6. **LinkedListNode[T]** — Generic linked list node with chain traversal.
7. **Type Aliases** — Pre-defined aliases for all common primitive types (e.g., `StringCollection`, `IntHashset`, `StringStringHashmap`).
8. **Generic Functions** (`funcs.go`) — Cross-type transformations requiring `any`: `MapCollection`, `FlatMapCollection`, `ReduceCollection`, `GroupByCollection`, `ContainsFunc`, `IndexOfFunc`, `MapSimpleSlice`.
9. **Comparable Constraint Functions** (`comparablefuncs.go`) — Equality-based operations requiring `comparable`: `ContainsAll`, `ContainsAny`, `RemoveItem`, `RemoveAllItems`, `ToHashset`, `DistinctSimpleSlice`, `ContainsSimpleSliceItem`.
10. **Ordered Constraint Functions** (`orderedfuncs.go`) — Functions requiring `cmp.Ordered` for sorting, min/max, clamping, and sum across all data structures.

## Comparable Constraint Functions

Package-level functions in `comparablefuncs.go` require `comparable` type constraints. They provide equality-based operations without custom predicates.

### Collection[T comparable]

| Function | Description |
|----------|-------------|
| `ContainsAll(col, items...)` | True if collection contains all given items |
| `ContainsAny(col, items...)` | True if collection contains any given item |
| `RemoveItem(col, item)` | Remove first occurrence; returns `true` if found |
| `RemoveAllItems(col, item)` | Remove all occurrences; returns count removed |
| `ToHashset(col)` | Convert `Collection[T]` → `Hashset[T]` |

### SimpleSlice[T comparable]

| Function | Description |
|----------|-------------|
| `DistinctSimpleSlice(ss)` | New slice with duplicates removed |
| `ContainsSimpleSliceItem(ss, item)` | Check if item exists |

### Usage Examples

```go
col := coregeneric.New.Collection.Int.Items(1, 2, 3, 2, 1)

// Check membership
coregeneric.ContainsAll(col, 1, 3)   // true
coregeneric.ContainsAny(col, 5, 3)   // true

// Remove
coregeneric.RemoveItem(col, 2)       // true, removes first 2
coregeneric.RemoveAllItems(col, 1)   // 2 (removed both 1s)

// Convert to set
set := coregeneric.ToHashset(col)    // Hashset[int]{3}

// SimpleSlice dedup
ss := coregeneric.New.SimpleSlice.String.Items("a", "b", "a")
unique := coregeneric.DistinctSimpleSlice(ss) // ["a", "b"]
```

## Ordered Constraint Functions

Package-level functions in `orderedfuncs.go` require `cmp.Ordered` type constraints. They are organized by data structure:

### Collection[T cmp.Ordered]

| Function | Description |
|----------|-------------|
| `SortCollection(col)` | In-place ascending sort |
| `SortCollectionDesc(col)` | In-place descending sort |
| `MinCollection(col)` | Minimum element (panics on empty) |
| `MaxCollection(col)` | Maximum element (panics on empty) |
| `MinCollectionOrDefault(col, def)` | Minimum or default if empty |
| `MaxCollectionOrDefault(col, def)` | Maximum or default if empty |
| `IsSortedCollection(col)` | Check if ascending sorted |
| `SumCollection(col)` | Sum all elements |
| `ClampCollection(col, min, max)` | Clamp all values to `[min, max]` |

### Hashset[T cmp.Ordered]

| Function | Description |
|----------|-------------|
| `SortedListHashset(hs)` | Items as sorted ascending slice |
| `SortedListDescHashset(hs)` | Items as sorted descending slice |
| `SortedCollectionHashset(hs)` | Items as sorted `Collection[T]` |
| `MinHashset(hs)` | Minimum element (panics on empty) |
| `MaxHashset(hs)` | Maximum element (panics on empty) |
| `MinHashsetOrDefault(hs, def)` | Minimum or default if empty |
| `MaxHashsetOrDefault(hs, def)` | Maximum or default if empty |

### Hashmap — Key-Ordered [K cmp.Ordered, V any]

| Function | Description |
|----------|-------------|
| `SortedKeysHashmap(hm)` | Keys sorted ascending |
| `SortedKeysDescHashmap(hm)` | Keys sorted descending |
| `MinKeyHashmap(hm)` | Minimum key (panics on empty) |
| `MaxKeyHashmap(hm)` | Maximum key (panics on empty) |
| `MinKeyHashmapOrDefault(hm, def)` | Minimum key or default if empty |
| `MaxKeyHashmapOrDefault(hm, def)` | Maximum key or default if empty |

### Hashmap — Value-Ordered [K comparable, V cmp.Ordered]

| Function | Description |
|----------|-------------|
| `SortedValuesHashmap(hm)` | Values sorted ascending |
| `MinValueHashmap(hm)` | Minimum value (panics on empty) |
| `MaxValueHashmap(hm)` | Maximum value (panics on empty) |
| `MinValueHashmapOrDefault(hm, def)` | Minimum value or default if empty |
| `MaxValueHashmapOrDefault(hm, def)` | Maximum value or default if empty |

### SimpleSlice[T cmp.Ordered]

| Function | Description |
|----------|-------------|
| `SortSimpleSlice(ss)` | In-place ascending sort |
| `MinSimpleSlice(ss)` | Minimum element (panics on empty) |
| `MaxSimpleSlice(ss)` | Maximum element (panics on empty) |

## Relationship to corestr

`corestr` remains the string-specific package with string-only methods (Join, EqualFold, Trim, Split, etc.). Internally, `corestr` types can delegate to `coregeneric` for common operations. The public API of `corestr` stays the same.

```
coregeneric.Collection[T]  ← generic foundation
    ↑
corestr.Collection         ← string-specific (Join, EqualFold, etc.)
```

## New Creator Pattern

Uses `typedXCreator[T]` — one generic struct definition covers all primitive types:

```go
// One definition...
type typedCollectionCreator[T any] struct{}
func (it typedCollectionCreator[T]) Empty() *Collection[T] { ... }
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] { ... }

// ...instantiated for every type
type newCollectionCreator struct {
    String  typedCollectionCreator[string]
    Int     typedCollectionCreator[int]
    Float64 typedCollectionCreator[float64]
    // ... 16 total types
}
```

### Usage Examples

```go
// Collection
col := coregeneric.New.Collection.String.Cap(10)
col := coregeneric.New.Collection.Int.Items(1, 2, 3)
col := coregeneric.New.Collection.Float64.From(existingSlice)

// Hashset
set := coregeneric.New.Hashset.String.Items("a", "b", "c")
set := coregeneric.New.Hashset.Int.Cap(100)

// Hashmap
hm := coregeneric.New.Hashmap.StringString.Cap(20)
hm := coregeneric.New.Hashmap.StringAny.From(existingMap)

// SimpleSlice
ss := coregeneric.New.SimpleSlice.Int.Items(1, 2, 3)

// LinkedList
ll := coregeneric.New.LinkedList.String.Items("a", "b", "c")
```

## Key Patterns

- All types use embedded `sync.Mutex` for thread-safe `*Lock` variants.
- Zero-nil safety: nil receivers return safe defaults (0, empty, false).
- Generic functions (`MapCollection`, `ReduceCollection`, etc.) are package-level because Go doesn't allow additional type parameters on generic methods.
- Type aliases in `types.go` provide convenient shorthand (e.g., `StringCollection = Collection[string]`).

## Supported Primitive Types

| Category | Types |
|----------|-------|
| Signed integers | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned integers | `uint`, `uint8`, `uint16`, `uint32`, `uint64` |
| Floats | `float32`, `float64` |
| Other | `byte`, `bool`, `string`, `any` |

## File Organization

```
coredata/coregeneric/
├── vars.go                        # New = &newCreator{}
├── newCreator.go                  # Root aggregator struct
├── newCollectionCreator.go        # Collection sub-creator (16 type fields)
├── newHashsetCreator.go           # Hashset sub-creator (14 type fields)
├── newHashmapCreator.go           # Hashmap sub-creator (9 key-value combos)
├── newSimpleSliceCreator.go       # SimpleSlice sub-creator (16 type fields)
├── newLinkedListCreator.go        # LinkedList sub-creator (12 type fields)
├── typedCollectionCreator.go      # Generic typed creator for Collection
├── typedHashsetCreator.go         # Generic typed creator for Hashset
├── typedHashmapCreator.go         # Generic typed creator for Hashmap
├── typedSimpleSliceCreator.go     # Generic typed creator for SimpleSlice
├── typedLinkedListCreator.go      # Generic typed creator for LinkedList
├── Collection.go                  # Collection[T] type + methods
├── Hashset.go                     # Hashset[T] type + methods
├── Hashmap.go                     # Hashmap[K,V] type + methods
├── SimpleSlice.go                 # SimpleSlice[T] type + methods
├── LinkedList.go                  # LinkedList[T] type + methods
├── LinkedListNode.go              # LinkedListNode[T] type + methods
├── types.go                       # Type aliases for all primitives
├── funcs.go                       # Generic cross-type functions
├── orderedfuncs.go                # cmp.Ordered constraint functions (Sort, Min, Max, Sum, Clamp)
├── comparablefuncs.go             # comparable constraint functions (ContainsAll, Distinct, ToHashset)
```

## How to Extend Safely

- **New primitive type**: Add a field to each `newXCreator` struct and a type alias in `types.go`.
- **New data structure**: Create the generic type, a `typedXCreator[T]`, and a `newXCreator` aggregator. Add the field to `newCreator`.
- **New generic function**: Add to `funcs.go` as a package-level function.
- **Type-specific methods**: Keep in the type-specific package (e.g., `corestr` for string methods).

## Related Docs

- [Repo Overview](../00-repo-overview.md)
- [Folder Map](../01-folder-map.md)
- [New Creator Pattern](../21-new-creator-pattern.md)
- [coredata Overview](./05-coredata.md)
