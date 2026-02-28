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
8. **Generic Functions** — Cross-type transformations: `MapCollection`, `FlatMapCollection`, `ReduceCollection`, `GroupByCollection`, `Distinct`, `ContainsItem`, `IndexOfItem`.

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
└── funcs.go                       # Generic cross-type functions
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
