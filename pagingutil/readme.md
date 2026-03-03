# pagingutil — Pagination Utilities

## Overview

Package `pagingutil` provides simple pagination math: given a total length, page index, and page size, it computes skip/end offsets and validates whether paging is applicable. Used by collection types for `GetPagedCollection` operations.

## Architecture

```
pagingutil/
├── PagingRequest.go      # PagingRequest struct — input parameters
├── PagingInfo.go         # PagingInfo struct — computed output
├── GetPagesSize.go       # GetPagesSize — total pages calculation
├── GetPagingInfo.go      # GetPagingInfo — full paging computation
├── pagingValidation.go   # Internal validation helpers and clamping
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
    PageIndex, SkipItems, EndingLength, TotalPages int
    IsPagingPossible                               bool
}
```

- `PageIndex` — the resolved page index (may differ from request due to clamping)
- `SkipItems` — offset to skip (`EachPageSize * (PageIndex - 1)`)
- `EndingLength` — end index (clamped to `Length`)
- `TotalPages` — total number of pages (`ceil(Length / EachPageSize)`)
- `IsPagingPossible` — `false` if items fit in a single page or input is invalid

## Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `GetPagesSize` | `(eachPageSize, totalLength int) int` | Ceiling division — total number of pages. Returns `0` for invalid page size. |
| `GetPagingInfo` | `(PagingRequest) PagingInfo` | Full paging computation with validation and clamping |

## Validation Rules

`GetPagingInfo` applies the following validation rules in order:

| Condition | Behavior | Result |
|-----------|----------|--------|
| `EachPageSize ≤ 0` | Invalid page size — returns zero-value `PagingInfo` | `PagingInfo{}` |
| `Length ≤ 0` | No items — returns empty result | `PageIndex: 0, IsPagingPossible: false` |
| `Length < EachPageSize` | All items fit in one page | `PageIndex: 1, SkipItems: 0, EndingLength: Length, IsPagingPossible: false` |
| `PageIndex < 1` | Clamped to first page | `PageIndex → 1` |
| `PageIndex > totalPages` | Clamped to last page | `PageIndex → totalPages` |
| Normal case | Standard pagination | `SkipItems` and `EndingLength` computed, `IsPagingPossible: true` |

`GetPagesSize` returns `0` when `eachPageSize ≤ 0`, avoiding division-by-zero.

## Edge Case Behavior

### Page Index Clamping

Rather than panicking on out-of-range page indices, the function **clamps** to valid bounds:

```go
// PageIndex 0 or negative → treated as page 1
info := GetPagingInfo(PagingRequest{Length: 50, PageIndex: 0, EachPageSize: 10})
// info.PageIndex: 1, info.SkipItems: 0

// PageIndex beyond last page → clamped to last page
info := GetPagingInfo(PagingRequest{Length: 50, PageIndex: 999, EachPageSize: 10})
// info.PageIndex: 5, info.SkipItems: 40, info.EndingLength: 50
```

### Single Page (No Paging Needed)

When all items fit in one page, `IsPagingPossible` is `false` and the full range is returned:

```go
info := GetPagingInfo(PagingRequest{Length: 5, PageIndex: 1, EachPageSize: 10})
// info.IsPagingPossible: false, info.SkipItems: 0, info.EndingLength: 5
```

### Empty or Invalid Input

```go
// Zero length → no items to page
info := GetPagingInfo(PagingRequest{Length: 0, PageIndex: 1, EachPageSize: 10})
// info.PageIndex: 0, info.IsPagingPossible: false

// Zero page size → zero-value PagingInfo
info := GetPagingInfo(PagingRequest{Length: 50, PageIndex: 1, EachPageSize: 0})
// info == PagingInfo{}

// Negative page size → same as zero
info := GetPagingInfo(PagingRequest{Length: 50, PageIndex: 1, EachPageSize: -5})
// info == PagingInfo{}
```

### Last Page (Partial)

The ending index is clamped so it never exceeds `Length`:

```go
info := GetPagingInfo(PagingRequest{Length: 95, PageIndex: 10, EachPageSize: 10})
// info.SkipItems: 90, info.EndingLength: 95 (not 100)
```

## Internal Validation Helpers

The `pagingValidation.go` file provides focused predicate functions used by `GetPagingInfo`:

| Helper | Purpose |
|--------|---------|
| `isLengthEmpty(length)` | `length ≤ 0` |
| `isPageSizeInvalid(size)` | `size ≤ 0` |
| `isPageIndexBelowMinimum(idx)` | `idx < 1` |
| `isPageIndexAboveMaximum(idx, total)` | `idx > total` |
| `isPagingOutOfRange(len, size)` | `len < size` |
| `hasNoItems(length)` | `length == 0` |
| `clampedPageIndex(idx, total)` | Clamps to `[1, totalPages]` |
| `calculateSkipItems(idx, size)` | `size * (idx - 1)` |
| `clampedEndingLength(end, len)` | `min(end, len)` |

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
