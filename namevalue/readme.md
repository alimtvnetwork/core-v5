# namevalue — Generic Name-Value Pairs & Collections

## Overview

Package `namevalue` provides a generic `Instance[K comparable, V any]` model for key-value pairs and a corresponding `Collection[K, V]` for managing ordered lists of pairs. Specialized type aliases cover common pairings (`StringAny`, `StringString`, `StringInt`, etc.) for ergonomic use without manual type parameter passing.

## Architecture

```
namevalue/
├── Instance.go                # Instance[K, V] — generic name-value pair
├── Collection.go              # Collection[K, V] — ordered collection with add/prepend/clone/join
├── NameValuesCollection.go    # NameValuesCollection (= Collection[string, any]) — backward-compatible alias
├── AppendsIf.go               # AppendsIf[K, V] — conditional append to slice
├── PrependsIf.go              # PrependsIf[K, V] — conditional prepend to slice
├── aliases.go                 # Type aliases: StringAny, StringString, StringInt, etc.
└── readme.md
```

## Instance[K, V]

A generic struct holding a `Name` (key) and `Value` (value).

```go
type Instance[K comparable, V any] struct {
    Name  K
    Value V
}
```

| Method | Description |
|--------|-------------|
| `IsNull()` | Nil pointer check |
| `String()` | `"Name: Value"` format |
| `JsonString()` | JSON serialization |
| `Dispose()` | Zero out fields |

## Collection[K, V]

An ordered collection of `Instance[K, V]` items with lazy string caching.

### Factory Functions

| Function | Description |
|----------|-------------|
| `NewGenericCollection[K, V](capacity)` | Create with capacity |
| `NewGenericCollectionDefault[K, V]()` | Create with default capacity (5) |
| `EmptyGenericCollection[K, V]()` | Create empty |
| `NewGenericCollectionUsing[K, V](isClone, ...items)` | Create from existing items |

### Mutation

| Method | Description |
|--------|-------------|
| `Add(item)` | Append single item |
| `Adds(...items)` | Append multiple items |
| `AddsPtr(...*items)` | Append from pointers (nil-safe) |
| `AddsIf(bool, ...items)` | Conditional append |
| `Append(...items)` | Alias for Adds |
| `AppendIf(bool, ...items)` | Conditional append |
| `AppendUsingFuncIf(bool, func)` | Lazy conditional append |
| `Prepend(...items)` | Prepend items |
| `PrependIf(bool, ...items)` | Conditional prepend |
| `PrependUsingFuncIf(bool, func)` | Lazy conditional prepend |
| `AppendPrependIf(bool, prepend, append)` | Combined conditional operation |
| `ConcatNew(...items)` | Clone + append (immutable) |
| `ConcatNewPtr(...*items)` | Clone + append pointers |

### Inspection

| Method | Description |
|--------|-------------|
| `Length()` / `Count()` | Item count |
| `IsEmpty()` / `HasAnyItem()` | Emptiness checks |
| `LastIndex()` / `HasIndex(int)` | Index operations |
| `IsEqualByString(*Collection)` | Equality via string comparison (bypasses non-comparable V) |

### Serialization

| Method | Description |
|--------|-------------|
| `String()` | Joined lines (lazy cached) |
| `Strings()` | String slice |
| `JsonString()` | JSON serialization |
| `JsonStrings()` | JSON string slice |
| `Join(joiner)` / `JoinLines()` / `JoinCsv()` | Various join formats |
| `CsvStrings()` / `JoinCsvLine()` | CSV formatting |
| `CompiledLazyString()` | Lazy-compiled string |
| `Error()` / `ErrorUsingMessage(msg)` | Convert to error |

### Lifecycle

| Method | Description |
|--------|-------------|
| `Clone()` / `ClonePtr()` | Deep copy |
| `Clear()` | Async dispose + reset |
| `Dispose()` | Clear + nil items |
| `InvalidateLazyString()` | Reset string cache |

## Package-Level Functions

| Function | Description |
|----------|-------------|
| `AppendsIf[K, V](bool, slice, ...items)` | Conditional append to raw slice |
| `PrependsIf[K, V](bool, slice, ...items)` | Conditional prepend to raw slice |

## Type Aliases

### Instance Aliases

| Alias | Expands To |
|-------|------------|
| `StringAny` | `Instance[string, any]` |
| `StringString` | `Instance[string, string]` |
| `StringInt` | `Instance[string, int]` |
| `StringMapAny` | `Instance[string, map[string]any]` |
| `StringMapString` | `Instance[string, map[string]string]` |

### Collection Aliases

| Alias | Expands To |
|-------|------------|
| `NameValuesCollection` | `Collection[string, any]` |
| `StringStringCollection` | `Collection[string, string]` |
| `StringIntCollection` | `Collection[string, int]` |
| `StringMapAnyCollection` | `Collection[string, map[string]any]` |
| `StringMapStringCollection` | `Collection[string, map[string]string]` |

### Legacy Factory Functions

| Function | Description |
|----------|-------------|
| `NewNameValuesCollection(capacity)` | `Collection[string, any]` with capacity |
| `NewCollection()` | `Collection[string, any]` with default capacity |
| `EmptyNameValuesCollection()` | Empty `Collection[string, any]` |
| `NewNewNameValuesCollectionUsing(isClone, ...items)` | From existing items |

## Usage

```go
import "gitlab.com/auk-go/core/namevalue"

// Generic usage
pair := namevalue.Instance[string, int]{Name: "count", Value: 42}
fmt.Println(pair.String()) // "count: 42"

// Using aliases
item := namevalue.StringAny{Name: "key", Value: "value"}

// Collection
col := namevalue.NewGenericCollectionDefault[string, string]()
col.Add(namevalue.StringString{Name: "host", Value: "localhost"})
col.Add(namevalue.StringString{Name: "port", Value: "8080"})
fmt.Println(col.JoinCsv()) // "host: localhost","port: 8080"
```

## How to Extend Safely

- **New type aliases**: Add to `aliases.go` — follow the existing naming pattern (`{KeyType}{ValueType}`).
- **New collection methods**: Add to `Collection.go` — ensure `InvalidateLazyString()` is called on mutations.
- **New utility functions**: Add as separate files (e.g., `FilterIf.go`) following the `AppendsIf`/`PrependsIf` pattern.

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
