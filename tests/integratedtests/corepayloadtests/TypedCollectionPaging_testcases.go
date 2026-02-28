package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// GetPagesSize
// =============================================================================

var typedCollectionPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 1 when items fewer than page size",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 10",
			"pageSize": 10,
		},
		ExpectedInput: []string{
			"1",
		},
	},
	{
		Title: "GetPagesSize returns exact page count when evenly divisible",
		ArrangeInput: args.Map{
			"when":     "10 items with page size 5",
			"count":    10,
			"pageSize": 5,
		},
		ExpectedInput: []string{
			"2",
		},
	},
	{
		Title: "GetPagesSize returns ceiling when not evenly divisible",
		ArrangeInput: args.Map{
			"when":     "7 items with page size 3",
			"count":    7,
			"pageSize": 3,
		},
		ExpectedInput: []string{
			"3",
		},
	},
}

// =============================================================================
// GetSinglePageCollection
// =============================================================================

var typedCollectionSinglePageTestCases = []coretestcases.CaseV1{
	{
		Title: "GetSinglePageCollection returns first page correctly",
		ArrangeInput: args.Map{
			"when":      "page 1 of 2 items per page from 5 items",
			"count":     5,
			"pageSize":  2,
			"pageIndex": 1,
		},
		ExpectedInput: []string{
			"2",
			"user-0",
			"user-1",
		},
	},
	{
		Title: "GetSinglePageCollection returns last partial page",
		ArrangeInput: args.Map{
			"when":      "page 3 of 2 items per page from 5 items",
			"count":     5,
			"pageSize":  2,
			"pageIndex": 3,
		},
		ExpectedInput: []string{
			"1",
			"user-4",
		},
	},
	{
		Title: "GetSinglePageCollection returns entire collection when smaller than page size",
		ArrangeInput: args.Map{
			"when":      "3 items with page size 10",
			"count":     3,
			"pageSize":  10,
			"pageIndex": 1,
		},
		ExpectedInput: []string{
			"3",
			"user-0",
			"user-1",
			"user-2",
		},
	},
}

// =============================================================================
// GetPagedCollection
// =============================================================================

var typedCollectionPagedCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection splits into correct number of pages",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 2",
			"count":    5,
			"pageSize": 2,
		},
		ExpectedInput: []string{
			"3",
			"2",
			"2",
			"1",
		},
	},
	{
		Title: "GetPagedCollection returns single page when fewer items than page size",
		ArrangeInput: args.Map{
			"when":     "2 items with page size 10",
			"count":    2,
			"pageSize": 10,
		},
		ExpectedInput: []string{
			"1",
			"2",
		},
	},
	{
		Title: "GetPagedCollection with exact division has no partial pages",
		ArrangeInput: args.Map{
			"when":     "6 items with page size 3",
			"count":    6,
			"pageSize": 3,
		},
		ExpectedInput: []string{
			"2",
			"3",
			"3",
		},
	},
}

// =============================================================================
// GetPagedCollectionWithInfo
// =============================================================================

var typedCollectionPagedWithInfoTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollectionWithInfo returns correct PagingInfo per page",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 2",
			"count":    5,
			"pageSize": 2,
		},
		ExpectedInput: []string{
			"3",
			"1",
			"3",
			"2",
			"5",
			"2",
			"2",
			"3",
			"2",
			"5",
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
		ExpectedInput: []string{
			"1",
			"0",
		},
	},
}
