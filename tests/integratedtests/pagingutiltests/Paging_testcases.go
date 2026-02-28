package pagingutiltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var getPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 1 for exact fit",
		ArrangeInput: args.Map{
			"when":         "given 10 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  10,
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "GetPagesSize returns 2 for partial overflow",
		ArrangeInput: args.Map{
			"when":         "given 11 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  11,
		},
		ExpectedInput: []string{"2"},
	},
	{
		Title: "GetPagesSize returns 3 for 25 items page 10",
		ArrangeInput: args.Map{
			"when":         "given 25 items with page size 10",
			"eachPageSize": 10,
			"totalLength":  25,
		},
		ExpectedInput: []string{"3"},
	},
}

var getPagingInfoTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagingInfo returns correct skip and ending for page 2",
		ArrangeInput: args.Map{
			"when":         "given page 2 with page size 10 and 25 items",
			"length":       25,
			"pageIndex":    2,
			"eachPageSize": 10,
		},
		ExpectedInput: []string{
			"2",
			"10",
			"20",
			"true",
		},
	},
	{
		Title: "GetPagingInfo returns not possible when length < page size",
		ArrangeInput: args.Map{
			"when":         "given 5 items with page size 10",
			"length":       5,
			"pageIndex":    1,
			"eachPageSize": 10,
		},
		ExpectedInput: []string{
			"1",
			"0",
			"10",
			"false",
		},
	},
}
