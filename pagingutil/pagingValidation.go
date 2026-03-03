package pagingutil

// isLengthEmpty returns true when there are no items to page through.
func isLengthEmpty(length int) bool {
	return length <= 0
}

// isPageSizeInvalid returns true when page size is zero or negative.
func isPageSizeInvalid(eachPageSize int) bool {
	return eachPageSize <= 0
}

// isPageIndexBelowMinimum returns true when page index is less than 1.
func isPageIndexBelowMinimum(pageIndex int) bool {
	return pageIndex < 1
}

// isPageIndexAboveMaximum returns true when page index exceeds total pages.
func isPageIndexAboveMaximum(pageIndex, totalPages int) bool {
	return pageIndex > totalPages
}

// isPagingOutOfRange returns true when item count fits within a single page.
func isPagingOutOfRange(length, eachPageSize int) bool {
	return length < eachPageSize
}

// hasNoItems returns true when length is zero — distinct from negative guard.
func hasNoItems(length int) bool {
	return length == 0
}

// clampedPageIndex returns a valid page index:
//   - negative or zero → 1 (first page)
//   - beyond total pages → totalPages (last page)
//   - otherwise → original value
func clampedPageIndex(pageIndex, totalPages int) int {
	if isPageIndexBelowMinimum(pageIndex) {
		return 1
	}

	if isPageIndexAboveMaximum(pageIndex, totalPages) {
		return totalPages
	}

	return pageIndex
}

// calculateSkipItems returns the number of items to skip for a given page.
func calculateSkipItems(pageIndex, eachPageSize int) int {
	return eachPageSize * (pageIndex - 1)
}

// clampedEndingLength ensures ending index does not exceed total length.
func clampedEndingLength(endingIndex, length int) int {
	if endingIndex > length {
		return length
	}

	return endingIndex
}
