package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// GetPagesSize
// =============================================================================

var typedCollectionPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 1 -- 3 items page size 10",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 10",
			"pageSize": 10,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
	{
		Title: "GetPagesSize returns 2 -- 10 items page size 5 evenly divisible",
		ArrangeInput: args.Map{
			"when":     "10 items with page size 5",
			"count":    10,
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pagesSize": 2,
		},
	},
	{
		Title: "GetPagesSize returns 3 -- 7 items page size 3 ceiling",
		ArrangeInput: args.Map{
			"when":     "7 items with page size 3",
			"count":    7,
			"pageSize": 3,
		},
		ExpectedInput: args.Map{
			"pagesSize": 3,
		},
	},
}

// =============================================================================
// GetSinglePageCollection
// =============================================================================

var typedCollectionSinglePageTestCases = []coretestcases.CaseV1{
	{
		Title: "GetSinglePageCollection returns 2 items -- page 1 of 5 items page size 2",
		ArrangeInput: args.Map{
			"when":      "page 1 of 2 items per page from 5 items",
			"count":     5,
			"pageSize":  2,
			"pageIndex": 1,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 2,
			"item0":         "user-0",
			"item1":         "user-1",
		},
	},
	{
		Title: "GetSinglePageCollection returns 1 item -- last partial page 3 of 5 items",
		ArrangeInput: args.Map{
			"when":      "page 3 of 2 items per page from 5 items",
			"count":     5,
			"pageSize":  2,
			"pageIndex": 3,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 1,
			"item0":         "user-4",
		},
	},
	{
		Title: "GetSinglePageCollection returns all 3 items -- page size 10 larger than count",
		ArrangeInput: args.Map{
			"when":      "3 items with page size 10",
			"count":     3,
			"pageSize":  10,
			"pageIndex": 1,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 3,
			"item0":         "user-0",
			"item1":         "user-1",
			"item2":         "user-2",
		},
	},
}

// =============================================================================
// GetPagedCollection
// =============================================================================

var typedCollectionPagedCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection returns 3 pages -- 5 items page size 2",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 2",
			"count":    5,
			"pageSize": 2,
		},
		ExpectedInput: args.Map{
			"pageCount":  3,
			"page1Items": 2,
			"page2Items": 2,
			"page3Items": 1,
		},
	},
	{
		Title: "GetPagedCollection returns 1 page -- 2 items page size 10",
		ArrangeInput: args.Map{
			"when":     "2 items with page size 10",
			"count":    2,
			"pageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 2,
		},
	},
	{
		Title: "GetPagedCollection returns 2 pages -- 6 items page size 3 exact division",
		ArrangeInput: args.Map{
			"when":     "6 items with page size 3",
			"count":    6,
			"pageSize": 3,
		},
		ExpectedInput: args.Map{
			"pageCount":  2,
			"page1Items": 3,
			"page2Items": 3,
		},
	},
}

// =============================================================================
// GetPagedCollectionWithInfo
// =============================================================================

var typedCollectionPagedWithInfoTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollectionWithInfo returns correct PagingInfo -- 5 items page size 2",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 2",
			"count":    5,
			"pageSize": 2,
		},
		ExpectedInput: args.Map{
			"pageCount":          3,
			"p1CurrentPageIndex": 1,
			"p1TotalPages":       3,
			"p1PerPageItems":     2,
			"p1TotalItems":       5,
			"p2CurrentPageIndex": 2,
			"p2TotalPages":       3,
			"p2PerPageItems":     2,
			"p2TotalItems":       5,
		},
	},
}

// =============================================================================
// Edge: paging on empty collection
// =============================================================================

var typedCollectionPagingEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection on empty collection returns single empty page",
		ArrangeInput: args.Map{
			"when":     "empty collection with page size 5",
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 0,
		},
	},
}
