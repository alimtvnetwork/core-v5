# pagingutil — Pagination Utilities

## Overview

Package `pagingutil` provides simple pagination math: given a total length, page index, and page size, it computes skip/end offsets and validates whether paging is applicable. Used by collection types for `GetPagedCollection` operations.

## Architecture

```
pagingutil/
├── PagingRequest.go    # PagingRequest struct — input parameters
├── PagingInfo.go       # PagingInfo struct — computed output
├── GetPagesSize.go     # GetPagesSize — total pages calculation
├── GetPagingInfo.go    # GetPagingInfo — full paging computation
└── readme.md
```

## Types

### PagingRequest

```go
type PagingRequest struct {
    Length, PageIndex, EachPageSize int
}
```

- `Length` — total number of items
- `PageIndex` — 1-based page number
- `EachPageSize` — items per page

### PagingInfo

```go
type PagingInfo struct {
    PageIndex, SkipItems, EndingLength int
    IsPagingPossible                   bool
}
```

- `PageIndex` — echoed from request
- `SkipItems` — offset to skip (`EachPageSize * (PageIndex - 1)`)
- `EndingLength` — end index (clamped to `Length`)
- `IsPagingPossible` — `false` if `Length < EachPageSize`

## Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `GetPagesSize` | `(eachPageSize, totalLength int) int` | Ceiling division — total number of pages |
| `GetPagingInfo` | `(PagingRequest) PagingInfo` | Full paging computation with bounds checking |

### GetPagingInfo Behavior

- If `Length < EachPageSize`: returns `IsPagingPossible: false` with `SkipItems: 0`
- If `PageIndex <= 0`: panics via `errcore.CannotBeNegativeIndexType`
- Clamps `EndingLength` to `Length` when the last page is partial

## Usage

```go
import "gitlab.com/auk-go/core/pagingutil"

// Total pages
pages := pagingutil.GetPagesSize(10, 95) // 10

// Page info
info := pagingutil.GetPagingInfo(pagingutil.PagingRequest{
    Length:       95,
    PageIndex:    3,
    EachPageSize: 10,
})
// info.SkipItems: 20, info.EndingLength: 30, info.IsPagingPossible: true

// Last page (partial)
last := pagingutil.GetPagingInfo(pagingutil.PagingRequest{
    Length:       95,
    PageIndex:    10,
    EachPageSize: 10,
})
// last.SkipItems: 90, last.EndingLength: 95
```

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
