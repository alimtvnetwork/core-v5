# New Creator Pattern

## Overview

The **New Creator** pattern is the standardized object construction approach used across the auk-go ecosystem. It replaces flat `NewX()` factory functions with a hierarchical, namespace-like factory that improves IDE discoverability, breaks down complex initialization into smaller logical chunks, and makes code self-documenting.

## Architecture

The pattern uses a three-layer factory hierarchy:

```
Layer 1: Root Aggregator    →  package-level `New` variable
Layer 2: Sub-Creators       →  fields grouping related factories (e.g., `New.Collection`)
Layer 3: Factory Methods    →  specific creation methods (e.g., `.Cap(10)`, `.Empty()`)
```

### Diagram

```
New (root aggregator)
├── Collection (sub-creator: groups all collection types)
│   ├── String (type creator)
│   │   ├── Empty()
│   │   ├── Cap(n)
│   │   ├── From(items)
│   │   ├── Clone(items)
│   │   └── LenCap(len, cap)
│   ├── Int (type creator)
│   │   ├── Empty()
│   │   ├── Cap(n)
│   │   └── ...
│   └── AnyMap (type creator)
│       └── ...
├── Info (sub-creator: another domain)
│   ├── Plain (variant creator)
│   │   ├── Default(name, desc, url)
│   │   ├── NameDescUrlExamples(...)
│   │   └── ...
│   └── Secure (variant creator)
│       ├── Default(name, desc, url)
│       └── ...
└── ... (more sub-creators)
```

## Implementation Rules

### 1. Root Aggregator — `vars.go`

The root aggregator is a package-level variable. It is always named `New` and is an unexported struct.

```go
// vars.go
var New = &newCreator{}
```

### 2. Root Creator Struct — `newCreator.go`

Contains only sub-creator fields. No methods.

```go
// newCreator.go
type newCreator struct {
    Collection newCollectionCreator
    Info       newInfoCreator
}
```

### 3. Sub-Creator — `newCollectionCreator.go`

Groups related type-specific creators. Acts as a namespace.

```go
// newCollectionCreator.go
type newCollectionCreator struct {
    String    newStringCollectionCreator
    Int       newIntCollectionCreator
    Float64   newFloat64CollectionCreator
}
```

### 4. Type-Specific Creator — One File Per Type

Each type creator lives in its own file and contains only factory methods for that type.

```go
// newStringCollectionCreator.go
type newStringCollectionCreator struct{}

func (it newStringCollectionCreator) Empty() *StringCollection {
    return EmptyCollection[string]()
}

func (it newStringCollectionCreator) Cap(capacity int) *StringCollection {
    return NewCollection[string](capacity)
}

func (it newStringCollectionCreator) From(items []string) *StringCollection {
    return CollectionFrom[string](items)
}

func (it newStringCollectionCreator) Clone(items []string) *StringCollection {
    return CollectionClone[string](items)
}
```

### 5. Generic Typed Creator (coregeneric innovation)

When a package uses Go generics, a **single generic struct** replaces all per-type creator files:

```go
// typedCollectionCreator.go — ONE definition covers ALL types
type typedCollectionCreator[T any] struct{}

func (it typedCollectionCreator[T]) Empty() *Collection[T] {
    return EmptyCollection[T]()
}

func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] {
    return NewCollection[T](capacity)
}

func (it typedCollectionCreator[T]) From(items []T) *Collection[T] {
    return CollectionFrom[T](items)
}

func (it typedCollectionCreator[T]) Clone(items []T) *Collection[T] {
    return CollectionClone[T](items)
}

func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] {
    return CollectionFrom[T](items)
}
```

The sub-creator then instantiates it per primitive type via struct fields:

```go
// newCollectionCreator.go
type newCollectionCreator struct {
    String  typedCollectionCreator[string]
    Int     typedCollectionCreator[int]
    Int64   typedCollectionCreator[int64]
    Float64 typedCollectionCreator[float64]
    Byte    typedCollectionCreator[byte]
    Bool    typedCollectionCreator[bool]
    Any     typedCollectionCreator[any]
    // ... all 16 primitive types
}
```

This eliminates massive code duplication while preserving full IDE discoverability.

## Standard Factory Methods

Every type creator should implement these methods where applicable:

| Method | Purpose |
|--------|---------|
| `Empty()` | Zero-capacity instance |
| `Cap(n)` | Instance with pre-allocated capacity |
| `From(items)` | Wrap existing slice (no copy) |
| `Clone(items)` | Copy slice into new instance |
| `Create(items)` | Alias for `From` |
| `Items(items...)` | Variadic convenience |
| `LenCap(len, cap)` | Specific length and capacity |
| `Default(...)` | Domain-specific default creation |

## Usage Examples

### Collection Creation (coredynamic)

```go
// Before (flat functions)
col := coredynamic.NewStringCollection(10)
col := coredynamic.EmptyIntCollection()

// After (New Creator pattern)
col := coredynamic.New.Collection.String.Cap(10)
col := coredynamic.New.Collection.Int.Empty()
col := coredynamic.New.Collection.AnyMap.From(existingSlice)
col := coredynamic.New.Collection.Float64.Items(1.0, 2.5, 3.7)
```

### Generic Collection Creation (coregeneric)

```go
// Collection — works for any primitive type
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

### String Utilities (corestr)

```go
// String collection
col := corestr.New.Collection.Cap(20)
col := corestr.New.Collection.Strings(mySlice)

// Simple slice
slice := corestr.New.SimpleSlice.Cap(10)
slice := corestr.New.SimpleSlice.Lines("a", "b", "c")

// Hashset
set := corestr.New.Hashset.Strings(mySlice)
```

### Task Info (coretaskinfo)

```go
// Plain text info
info := coretaskinfo.New.Info.Plain.Default("name", "desc", "url")
info := coretaskinfo.New.Info.Plain.AllUrlExamples("name", "desc", "url", "hint", "err", "ex1")

// Secure info
info := coretaskinfo.New.Info.Secure.Default("name", "desc", "url")
info := coretaskinfo.New.Info.Secure.NameDescUrlExamples("name", "desc", "url", "ex1", "ex2")
```

## Dividing Complex Creations

When a domain has multiple creation variants (e.g., plain vs secure, with/without options), split them into separate sub-creators:

```
New.Info
├── Plain   → newInfoPlainTextCreator   (creates Info without ExcludeOptions)
├── Secure  → newInfoSecureTextCreator  (creates Info with IsSecureText: true)
├── Default(...)                         (simple convenience method)
└── Deserialized(bytes)                  (parsing from JSON)
```

This approach:
- **Reduces cognitive load**: Each creator handles one variant
- **Improves IDE autocomplete**: `New.Info.` shows Plain/Secure/Default immediately
- **Keeps files small**: Each sub-creator is its own file (~30-50 lines)

## File Organization

```
package/
├── vars.go                          # New = &newCreator{}
├── newCreator.go                    # Root aggregator struct
├── newCollectionCreator.go          # Sub-creator (aggregates type creators)
├── newStringCollectionCreator.go    # Type-specific factory
├── newIntCollectionCreator.go       # Type-specific factory
├── newFloat64CollectionCreator.go   # Type-specific factory
├── Collection.go                    # Collection[T] type definition
├── CollectionMethods.go             # Additional Collection[T] methods
└── CollectionTypes.go               # Type aliases + legacy factory shortcuts
```

## Design Principles

1. **One file per creator**: Each `newXCreator` struct gets its own file
2. **Struct-as-namespace**: Unexported structs group related factory methods
3. **No logic in aggregators**: Root and sub-creators only hold fields, never methods with logic
4. **Backward compatibility**: Legacy flat functions remain and delegate to the same underlying constructors
5. **Generics integration**: Type-specific creators wrap generic `NewCollection[T]` / `EmptyCollection[T]` calls

## Related Docs

- [Repo Overview](./00-repo-overview.md)
- [Code Strengths](./19-code-strengths.md)
- [Improvement Plan](./20-improvement-plan.md)
