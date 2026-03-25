# coredata — Data Structures & Serialization

Package `coredata` is the umbrella for rich data structures, generic collections, serialization utilities, and string processing. Contains sub-packages for generic and string-specific collections, JSON handling, dynamic reflection-based data, ranges, payloads, and once-computed values.

## Architecture

```
coredata/
├── coregeneric/           # Generic data structures: Collection[T], Hashset[T], Hashmap[K,V], SimpleSlice[T], LinkedList[T]
├── corestr/               # String-specialized collections with Join, EqualFold, Trim, Split, and 16+ types
├── corejson/              # JSON serialization/deserialization with builder patterns (Serialize.*, Deserialize.*)
├── coredynamic/           # Reflection-based dynamic data manipulation, type checking, casting
├── coreonce/              # Compute-once lazy value holders for all common types
├── corepayload/           # Enhanced payload structures
├── corerange/             # Range data types, min/max bounds, and boundary validation
├── coreapi/               # Typed API request/response models with generics
├── stringslice/           # 80+ pure functions for raw []string manipulation
├── Integers.go            # Integer slice type
├── IntegersDsc.go         # Descending integer slice
├── PointerIntegers.go     # Pointer integer slice
├── PointerIntegersDsc.go  # Descending pointer integer slice
├── PointerStrings.go      # Pointer string slice
├── PointerStringsDsc.go   # Descending pointer string slice
├── StringsDsc.go          # Descending string slice
└── BytesError.go          # Bytes with error pair
```

## Sub-Package Index

| Package | Description | README |
|---------|-------------|--------|
| [`coregeneric/`](./coregeneric/) | Generic data structures: `Collection[T]`, `Hashset[T]`, `Hashmap[K,V]`, `SimpleSlice[T]`, `LinkedList[T]` | [README](./coregeneric/README.md) |
| [`corestr/`](./corestr/) | String-specialized collections with Join, EqualFold, Trim, Split, and 16+ types | [README](./corestr/README.md) |
| [`corejson/`](./corejson/) | JSON serialization/deserialization with builder patterns (`Serialize.*`, `Deserialize.*`) | [README](./corejson/README.md) |
| [`coredynamic/`](./coredynamic/) | Reflection-based dynamic data manipulation, type checking, casting | [README](./coredynamic/README.md) |
| [`coreonce/`](./coreonce/) | Compute-once lazy value holders for all common types | [README](./coreonce/README.md) |
| [`corepayload/`](./corepayload/) | Enhanced payload structures | [README](./corepayload/README.md) |
| [`corerange/`](./corerange/) | Range data types, min/max bounds, and boundary validation | [README](./corerange/README.md) |
| [`coreapi/`](./coreapi/) | Typed API request/response models with generics | [README](./coreapi/README.md) |
| [`stringslice/`](./stringslice/) | 80+ pure functions for raw `[]string` manipulation | [README](./stringslice/README.md) |

## Root-Level Types

| File | Type |
|------|------|
| `Integers.go` | Integer slice type |
| `IntegersDsc.go` | Descending integer slice |
| `PointerIntegers.go` | Pointer integer slice |
| `PointerIntegersDsc.go` | Descending pointer integer slice |
| `PointerStrings.go` | Pointer string slice |
| `PointerStringsDsc.go` | Descending pointer string slice |
| `StringsDsc.go` | Descending string slice |
| `BytesError.go` | Bytes with error pair |

## Key Patterns

- **New Creator Pattern**: All sub-packages expose a `New` variable for builder-style construction (e.g., `corestr.New.Collection.Cap(10)`).
- **Empty Creator Pattern**: Quick empty instances via `Empty` variable (e.g., `corestr.Empty.SimpleSlice()`).
- **Thread Safety**: Collection types embed `sync.Mutex` with `*Lock` method variants.
- **Nil Safety**: Nil receivers return safe zero-value defaults.

## How to Extend Safely

- New collection types go in `corestr/` (string-specific) or `coregeneric/` (type-agnostic).
- New serialization formats go parallel to `corejson/`.
- Always provide a `newCreator` factory; avoid bare struct construction.
- New raw slice utilities go in `stringslice/`.

## Related Docs

- [Folder Spec](/spec/01-app/folders/05-coredata.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
