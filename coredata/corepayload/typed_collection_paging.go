package corepayload

import (
	"math"
	"sync"

	"github.com/alimtvnetwork/core/pagingutil"
)

// =============================================================================
// Paging
// =============================================================================

// GetPagesSize returns the number of pages for the given page size.
// Returns 0 if eachPageSize is zero or negative.
func (it *TypedPayloadCollection[T]) GetPagesSize(eachPageSize int) int {
	if eachPageSize <= 0 {
		return 0
	}

	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

// GetPagingInfo returns paging metadata for the given page size and 1-based page index.
func (it *TypedPayloadCollection[T]) GetPagingInfo(
	eachPageSize int,
	pageIndex int,
) pagingutil.PagingInfo {
	return pagingutil.GetPagingInfo(
		pagingutil.PagingRequest{
			Length:       it.Length(),
			PageIndex:    pageIndex,
			EachPageSize: eachPageSize,
		},
	)
}

// GetSinglePageCollection returns a single page of items. pageIndex is 1-based.
//
// If the total length is less than eachPageSize, the entire collection is returned.
//
// Usage:
//
//	page := collection.GetSinglePageCollection(10, 2) // second page of 10
func (it *TypedPayloadCollection[T]) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *TypedPayloadCollection[T] {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	pageInfo := it.GetPagingInfo(eachPageSize, pageIndex)

	return &TypedPayloadCollection[T]{
		items: it.items[pageInfo.SkipItems:pageInfo.EndingLength],
	}
}

// GetPagedCollection splits the collection into pages of the given size.
//
// Pages are populated concurrently using goroutines.
//
// Usage:
//
//	pages := collection.GetPagedCollection(10)
//	for _, page := range pages {
//	    page.ForEachData(func(i int, data User) { ... })
//	}
func (it *TypedPayloadCollection[T]) GetPagedCollection(
	eachPageSize int,
) []*TypedPayloadCollection[T] {
	length := it.Length()

	if length < eachPageSize {
		return []*TypedPayloadCollection[T]{it}
	}

	pageCount := it.GetPagesSize(eachPageSize)
	pages := make([]*TypedPayloadCollection[T], pageCount)

	wg := sync.WaitGroup{}
	wg.Add(pageCount)

	for i := 1; i <= pageCount; i++ {
		go func(pageIndex int) {
			pages[pageIndex-1] = it.GetSinglePageCollection(eachPageSize, pageIndex)
			wg.Done()
		}(i)
	}

	wg.Wait()

	return pages
}

// GetPagedCollectionWithInfo returns paged collections alongside a PagingInfo
// for each page, useful for APIs that return paginated responses.
func (it *TypedPayloadCollection[T]) GetPagedCollectionWithInfo(
	eachPageSize int,
) []TypedPayloadPage[T] {
	pages := it.GetPagedCollection(eachPageSize)
	totalPages := len(pages)
	totalItems := it.Length()

	result := make([]TypedPayloadPage[T], totalPages)

	for i, page := range pages {
		result[i] = TypedPayloadPage[T]{
			Collection: page,
			Paging: PagingInfo{
				CurrentPageIndex: i + 1,
				TotalPages:       totalPages,
				PerPageItems:     eachPageSize,
				TotalItems:       totalItems,
			},
		}
	}

	return result
}

// TypedPayloadPage wraps a single page of TypedPayloadCollection[T] with paging metadata.
type TypedPayloadPage[T any] struct {
	Collection *TypedPayloadCollection[T]
	Paging     PagingInfo
}
