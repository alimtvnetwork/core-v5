package pagingutil

type PagingInfo struct {
	PageIndex, SkipItems, EndingLength, TotalPages int
	IsPagingPossible                               bool
}
