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
		ExpectedInput: "1",
	},
	{
		Title: "GetPagesSize returns exact page count when evenly divisible",
		ArrangeInput: args.Map{
			"when":     "10 items with page size 5",
			"count":    10,
			"pageSize": 5,
		},
		ExpectedInput: "2",
	},
	{
		Title: "GetPagesSize returns ceiling when not evenly divisible",
		ArrangeInput: args.Map{
			"when":     "7 items with page size 3",
			"count":    7,
			"pageSize": 3,
		},
		ExpectedInput: "3",
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",      // pageItemCount
			Second: "user-0", // firstItem
			Third:  "user-1", // lastItem
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
		ExpectedInput: args.Two[string, string]{
			First:  "1",      // pageItemCount
			Second: "user-4", // firstItem
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3",      // pageItemCount
			Second: "user-0", // firstItem
			Third:  "user-1", // secondItem
			Fourth: "user-2", // lastItem
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3", // pageCount
			Second: "2", // page1Items
			Third:  "2", // page2Items
			Fourth: "1", // page3Items
		},
	},
	{
		Title: "GetPagedCollection returns single page when fewer items than page size",
		ArrangeInput: args.Map{
			"when":     "2 items with page size 10",
			"count":    2,
			"pageSize": 10,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "1", // pageCount
			Second: "2", // page1Items
		},
	},
	{
		Title: "GetPagedCollection with exact division has no partial pages",
		ArrangeInput: args.Map{
			"when":     "6 items with page size 3",
			"count":    6,
			"pageSize": 3,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "2", // pageCount
			Second: "3", // page1Items
			Third:  "3", // page2Items
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
		// 10 elements exceeds args.Six — keep as []string
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
		ExpectedInput: args.Two[string, string]{
			First:  "1", // pageCount
			Second: "0", // page1Items
		},
	},
}
