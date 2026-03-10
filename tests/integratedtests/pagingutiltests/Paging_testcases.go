package pagingutiltests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var getPagesSizeTestCases = []coretestcases.CaseV1{
	// === Positive cases ===
	{
		Title: "GetPagesSize returns 1 for exact fit",
		ArrangeInput: args.Map{
			"when":         "given 10 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  10,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
	{
		Title: "GetPagesSize returns 2 for partial overflow",
		ArrangeInput: args.Map{
			"when":         "given 11 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  11,
		},
		ExpectedInput: args.Map{
			"pagesSize": 2,
		},
	},
	{
		Title: "GetPagesSize returns 3 for 25 items page 10",
		ArrangeInput: args.Map{
			"when":         "given 25 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  25,
		},
		ExpectedInput: args.Map{
			"pagesSize": 3,
		},
	},
	{
		Title: "GetPagesSize returns 1 for single item",
		ArrangeInput: args.Map{
			"when":         "given 1 item with page size 10",
			"eachPageSize": 10,
			"totalLength":  1,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
	{
		Title: "GetPagesSize returns 10 for 100 items page 10",
		ArrangeInput: args.Map{
			"when":         "given 100 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  100,
		},
		ExpectedInput: args.Map{
			"pagesSize": 10,
		},
	},
	{
		Title: "GetPagesSize returns 1 for page size 1 and 1 item",
		ArrangeInput: args.Map{
			"when":         "given 1 item with page size 1",
			"eachPageSize": 1,
			"totalLength":  1,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
	{
		Title: "GetPagesSize returns 5 for page size 1 and 5 items",
		ArrangeInput: args.Map{
			"when":         "given 5 items with page size 1",
			"eachPageSize": 1,
			"totalLength":  5,
		},
		ExpectedInput: args.Map{
			"pagesSize": 5,
		},
	},

	// === Boundary / edge cases ===
	{
		Title: "GetPagesSize returns 0 for zero total length",
		ArrangeInput: args.Map{
			"when":         "given 0 items",
			"eachPageSize": 10,
			"totalLength":  0,
		},
		ExpectedInput: args.Map{
			"pagesSize": 0,
		},
	},

	// === Negative / guard cases ===
	{
		Title: "GetPagesSize returns 0 for zero page size",
		ArrangeInput: args.Map{
			"when":         "given page size 0 (division by zero guard)",
			"eachPageSize": 0,
			"totalLength":  25,
		},
		ExpectedInput: args.Map{
			"pagesSize": 0,
		},
	},
	{
		Title: "GetPagesSize returns 0 for negative page size",
		ArrangeInput: args.Map{
			"when":         "given negative page size",
			"eachPageSize": -5,
			"totalLength":  25,
		},
		ExpectedInput: args.Map{
			"pagesSize": 0,
		},
	},
}

var getPagingInfoTestCases = []coretestcases.CaseV1{
	// === Positive: standard paging ===
	{
		Title: "GetPagingInfo returns correct skip and ending for page 2",
		ArrangeInput: args.Map{
			"when":         "given page 2 with page size 10 and 25 items",
			"length":       25,
			"pageIndex":    2,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        2,
			"skipItems":        10,
			"endingLength":     20,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo page 1 starts at 0",
		ArrangeInput: args.Map{
			"when":         "given page 1 with page size 10 and 25 items",
			"length":       25,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     10,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo last page ending is clamped to length",
		ArrangeInput: args.Map{
			"when":         "given page 3 with page size 10 and 25 items",
			"length":       25,
			"pageIndex":    3,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        3,
			"skipItems":        20,
			"endingLength":     25,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo exact fit page ending equals length",
		ArrangeInput: args.Map{
			"when":         "given page 2 with page size 10 and 20 items",
			"length":       20,
			"pageIndex":    2,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        2,
			"skipItems":        10,
			"endingLength":     20,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo page size 1 single element pages",
		ArrangeInput: args.Map{
			"when":         "given page 3 with page size 1 and 5 items",
			"length":       5,
			"pageIndex":    3,
			"eachPageSize": 1,
		},
		ExpectedInput: args.Map{
			"pageIndex":        3,
			"skipItems":        2,
			"endingLength":     3,
			"isPagingPossible": true,
		},
	},

	// === Boundary: not pageable (length < eachPageSize) ===
	{
		Title: "GetPagingInfo returns not possible when length < page size",
		ArrangeInput: args.Map{
			"when":         "given 5 items with page size 10",
			"length":       5,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     5,
			"isPagingPossible": false,
		},
	},
	{
		Title: "GetPagingInfo returns not possible for 1 item page size 5",
		ArrangeInput: args.Map{
			"when":         "given 1 item with page size 5",
			"length":       1,
			"pageIndex":    1,
			"eachPageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     1,
			"isPagingPossible": false,
		},
	},

	// === Boundary: exact fit (length == eachPageSize) ===
	{
		Title: "GetPagingInfo exact fit length equals page size",
		ArrangeInput: args.Map{
			"when":         "given 10 items with page size 10",
			"length":       10,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     10,
			"isPagingPossible": true,
		},
	},

	// === NEW: zero length ===
	{
		Title: "GetPagingInfo zero length returns empty with pageIndex 0",
		ArrangeInput: args.Map{
			"when":         "given 0 items with page size 10",
			"length":       0,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        0,
			"skipItems":        0,
			"endingLength":     0,
			"isPagingPossible": false,
		},
	},

	// === NEW: negative page index clamped to first page ===
	{
		Title: "GetPagingInfo negative pageIndex clamped to first page",
		ArrangeInput: args.Map{
			"when":         "given negative pageIndex with 25 items",
			"length":       25,
			"pageIndex":    -3,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     10,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo zero pageIndex clamped to first page",
		ArrangeInput: args.Map{
			"when":         "given zero pageIndex with 25 items",
			"length":       25,
			"pageIndex":    0,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     10,
			"isPagingPossible": true,
		},
	},

	// === NEW: pageIndex too large clamped to last page ===
	{
		Title: "GetPagingInfo pageIndex beyond total pages clamped to last",
		ArrangeInput: args.Map{
			"when":         "given pageIndex 100 with 25 items and page size 10",
			"length":       25,
			"pageIndex":    100,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        3,
			"skipItems":        20,
			"endingLength":     25,
			"isPagingPossible": true,
		},
	},
	{
		Title: "GetPagingInfo pageIndex 5 with only 2 pages clamped to last",
		ArrangeInput: args.Map{
			"when":         "given pageIndex 5 with 15 items and page size 10",
			"length":       15,
			"pageIndex":    5,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        2,
			"skipItems":        10,
			"endingLength":     15,
			"isPagingPossible": true,
		},
	},

	// === NEW: invalid page size ===
	{
		Title: "GetPagingInfo zero page size returns empty PagingInfo",
		ArrangeInput: args.Map{
			"when":         "given zero page size",
			"length":       25,
			"pageIndex":    1,
			"eachPageSize": 0,
		},
		ExpectedInput: args.Map{
			"pageIndex":        0,
			"skipItems":        0,
			"endingLength":     0,
			"isPagingPossible": false,
		},
	},
	{
		Title: "GetPagingInfo negative page size returns empty PagingInfo",
		ArrangeInput: args.Map{
			"when":         "given negative page size",
			"length":       25,
			"pageIndex":    1,
			"eachPageSize": -5,
		},
		ExpectedInput: args.Map{
			"pageIndex":        0,
			"skipItems":        0,
			"endingLength":     0,
			"isPagingPossible": false,
		},
	},

	// === NEW: negative pageIndex with unpageable length ===
	{
		Title: "GetPagingInfo negative pageIndex with small length clamped to page 1",
		ArrangeInput: args.Map{
			"when":         "given negative pageIndex with 3 items and page size 10",
			"length":       3,
			"pageIndex":    -1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     3,
			"isPagingPossible": false,
		},
	},

	// === NEW: large pageIndex with unpageable length ===
	{
		Title: "GetPagingInfo large pageIndex with small length clamped to page 1",
		ArrangeInput: args.Map{
			"when":         "given pageIndex 99 with 3 items and page size 10",
			"length":       3,
			"pageIndex":    99,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        1,
			"skipItems":        0,
			"endingLength":     3,
			"isPagingPossible": false,
		},
	},

	// === NEW: negative length treated as empty ===
	{
		Title: "GetPagingInfo negative length treated as empty",
		ArrangeInput: args.Map{
			"when":         "given negative length",
			"length":       -5,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: args.Map{
			"pageIndex":        0,
			"skipItems":        0,
			"endingLength":     0,
			"isPagingPossible": false,
		},
	},
}
