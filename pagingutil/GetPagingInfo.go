package pagingutil

// GetPagingInfo calculates paging metadata from a PagingRequest.
//
// Validation rules:
//   - Zero or negative EachPageSize → empty PagingInfo (no paging possible).
//   - Zero Length → empty PagingInfo with PageIndex 0.
//   - PageIndex below 1 → clamped to first page (1).
//   - PageIndex above total pages → clamped to last page.
//   - Length < EachPageSize → single page, IsPagingPossible = false.
func GetPagingInfo(request PagingRequest) PagingInfo {
	// Guard: invalid page size
	if isPageSizeInvalid(request.EachPageSize) {
		return PagingInfo{}
	}

	length := request.Length

	// Guard: no items
	if isLengthEmpty(length) {
		return PagingInfo{
			PageIndex:        0,
			SkipItems:        0,
			EndingLength:     0,
			IsPagingPossible: false,
		}
	}

	// Guard: everything fits in one page
	if isPagingNotPossible(length, request.EachPageSize) {
		return PagingInfo{
			PageIndex:        1,
			SkipItems:        0,
			EndingLength:     length,
			IsPagingPossible: false,
		}
	}

	// Calculate total pages for clamping
	totalPages := GetPagesSize(request.EachPageSize, length)

	// Clamp page index to valid range
	pageIndex := clampedPageIndex(request.PageIndex, totalPages)

	// Calculate offsets
	skipItems := calculateSkipItems(pageIndex, request.EachPageSize)
	endingIndex := clampedEndingLength(skipItems+request.EachPageSize, length)

	return PagingInfo{
		PageIndex:        pageIndex,
		SkipItems:        skipItems,
		EndingLength:     endingIndex,
		IsPagingPossible: true,
	}
}
