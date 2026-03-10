package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// Edge cases for paging — single item, page size 1, large page size
// =============================================================================

var typedCollectionPagingEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection single item with page size 1",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 1",
			"count":    1,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 1,
		},
	},
	{
		Title: "GetPagedCollection page size 1 creates one page per item",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 1",
			"count":    3,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pageCount":  3,
			"page1Items": 1,
			"page2Items": 1,
			"page3Items": 1,
		},
	},
	{
		Title: "GetPagedCollection page size equals item count",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 5",
			"count":    5,
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 5,
		},
	},
	{
		Title: "GetPagedCollection page size larger than item count",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 100",
			"count":    3,
			"pageSize": 100,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 3,
		},
	},
}

// =============================================================================
// GetSinglePageCollection edge cases
// =============================================================================

var typedCollectionSinglePageEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetSinglePageCollection middle page returns correct items",
		ArrangeInput: args.Map{
			"when":      "page 2 of 3 items per page from 9 items",
			"count":     9,
			"pageSize":  3,
			"pageIndex": 2,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 3,
			"item0":         "user-3",
			"item1":         "user-4",
			"item2":         "user-5",
		},
	},
	{
		Title: "GetSinglePageCollection page size 1 returns single item",
		ArrangeInput: args.Map{
			"when":      "page 2 of page size 1 from 5 items",
			"count":     5,
			"pageSize":  1,
			"pageIndex": 2,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 1,
			"item0":         "user-1",
		},
	},
}

// =============================================================================
// GetPagedCollectionWithInfo edge cases
// =============================================================================

var typedCollectionPagedWithInfoEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollectionWithInfo single item has correct metadata",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 5",
			"count":    1,
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageCount":          1,
			"p1CurrentPageIndex": 1,
			"p1TotalPages":       1,
			"p1PerPageItems":     5,
			"p1TotalItems":       1,
		},
	},
	{
		Title: "GetPagedCollectionWithInfo exact division metadata",
		ArrangeInput: args.Map{
			"when":     "4 items with page size 2",
			"count":    4,
			"pageSize": 2,
		},
		ExpectedInput: args.Map{
			"pageCount":          2,
			"p1CurrentPageIndex": 1,
			"p1TotalPages":       2,
			"p1PerPageItems":     2,
			"p1TotalItems":       4,
			"p2CurrentPageIndex": 2,
			"p2TotalPages":       2,
			"p2PerPageItems":     2,
			"p2TotalItems":       4,
		},
	},
}

// =============================================================================
// GetPagesSize edge cases
// =============================================================================

var typedCollectionPagesSizeEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize with page size 1",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 1",
			"count":    5,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pagesSize": 5,
		},
	},
	{
		Title: "GetPagesSize with single item",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 10",
			"count":    1,
			"pageSize": 10,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
}

// ==========================================================================
// Paging empty with GetPagedCollectionWithInfo
// ==========================================================================

var typedCollectionPagingWithInfoEmptyTestCase = coretestcases.CaseV1{
	Title: "PagingWithInfo on empty collection returns 1 page with 0 items",
	ExpectedInput: args.Map{
		"pageCount":    1,
		"firstPageLen": 0,
	},
}
