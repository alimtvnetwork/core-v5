# coredata

## Folder Purpose

Umbrella package for rich data structures and serialization utilities. Contains sub-packages for string collections, JSON handling, dynamic reflection-based data, ranges, payloads, and once-computed values.

## Responsibilities

1. **corestr/** — String-centric collections: `Collection`, `Hashmap`, `Hashset`, `LinkedList`, `SimpleSlice`, `CharCollectionMap`, etc.
2. **corejson/** — JSON serialization/deserialization with builder patterns (`Serialize.*`, `Deserialize.*`).
3. **coredynamic/** — Reflection-based dynamic data manipulation, type checking, casting.
4. **coreonce/** — Compute-once value holders.
5. **corepayload/** — Enhanced payload structures.
6. **corerange/** — Range data types.
7. **coreapi/** — API-related data models.
8. **stringslice/** — String slice utilities.
9. Root-level files — Integer/String/Pointer slice types and descending variants.

## Key Patterns

- All sub-packages use a `newCreator` struct pattern with a package-level `New` variable.
- Builder pattern: `corestr.New.SimpleSlice.Cap(10)`.
- `Empty` creator pattern: `corestr.Empty.SimpleSlice()`.

## How to Extend Safely

- New collection types go in `corestr/`.
- New serialization formats go in parallel to `corejson/`.
- Always provide a `newCreator` factory; avoid bare struct construction.

## Related Docs

- [Repo Overview](../00-repo-overview.md)
- [Folder Map](../01-folder-map.md)
