package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		ExpectedInput: []string{
			"1",
			"1",
		},
	},
	{
		Title: "GetPagedCollection page size 1 creates one page per item",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 1",
			"count":    3,
			"pageSize": 1,
		},
		ExpectedInput: []string{
			"3",
			"1",
			"1",
			"1",
		},
	},
	{
		Title: "GetPagedCollection page size equals item count",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 5",
			"count":    5,
			"pageSize": 5,
		},
		ExpectedInput: []string{
			"1",
			"5",
		},
	},
	{
		Title: "GetPagedCollection page size larger than item count",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 100",
			"count":    3,
			"pageSize": 100,
		},
		ExpectedInput: []string{
			"1",
			"3",
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
		ExpectedInput: []string{
			"3",
			"user-3",
			"user-4",
			"user-5",
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
		ExpectedInput: []string{
			"1",
			"user-1",
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
		ExpectedInput: []string{
			"1",
			"1",   // CurrentPageIndex
			"1",   // TotalPages
			"5",   // PerPageItems
			"1",   // TotalItems
		},
	},
	{
		Title: "GetPagedCollectionWithInfo exact division metadata",
		ArrangeInput: args.Map{
			"when":     "4 items with page size 2",
			"count":    4,
			"pageSize": 2,
		},
		ExpectedInput: []string{
			"2",
			"1",   // page 1 CurrentPageIndex
			"2",   // TotalPages
			"2",   // PerPageItems
			"4",   // TotalItems
			"2",   // page 2 CurrentPageIndex
			"2",   // TotalPages
			"2",   // PerPageItems
			"4",   // TotalItems
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
		ExpectedInput: []string{"5"},
	},
	{
		Title: "GetPagesSize with single item",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 10",
			"count":    1,
			"pageSize": 10,
		},
		ExpectedInput: []string{"1"},
	},
}
