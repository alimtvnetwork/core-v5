# coreunique

Deduplication utilities for primitive slices. Currently provides integer uniqueness via the `intunique` sub-package.

## Architecture

```
coreunique/
└── intunique/
    ├── Get.go      # Deduplicate []int, return unique slice
    └── GetMap.go   # Deduplicate []int, return map[int]bool
```

## Features

### intunique

| Function | Description |
|---|---|
| `Get` | Returns a new `*[]int` containing only unique values from the input slice |
| `GetMap` | Returns a `*map[int]bool` representing the unique set of integers |

## Usage Examples

```go
import "gitlab.com/auk-go/core/coreunique/intunique"

nums := []int{1, 2, 2, 3, 3, 3}

// Get unique slice
unique := intunique.Get(&nums)
// *unique contains [1, 2, 3] (order not guaranteed)

// Get as map/set
set := intunique.GetMap(&nums)
// *set = map[int]bool{1: true, 2: true, 3: true}
```

## Design Notes

- Both functions accept `*[]int` and return pointers for consistency with the broader codebase.
- `Get` returns the original pointer unchanged for nil, empty, or single-element inputs.
- `GetMap` returns nil for nil input and an empty map for empty input.

## Related Docs

- [coreindexes](../coreindexes/readme.md) — index constants and utilities
