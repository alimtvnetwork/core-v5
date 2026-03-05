package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var collectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 0 for zero eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=0",
			"items":        []int{1, 2, 3, 4, 5},
			"eachPageSize": 0,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 0 for negative eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=-3",
			"items":        []int{1, 2, 3},
			"eachPageSize": -3,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 0 for empty collection",
		ArrangeInput: args.Map{
			"when":         "given empty collection with valid page size",
			"items":        []int{},
			"eachPageSize": 5,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 1 when items fit in one page",
		ArrangeInput: args.Map{
			"when":         "given 3 items with eachPageSize=5",
			"items":        []int{1, 2, 3},
			"eachPageSize": 5,
		},
		ExpectedInput: "1",
	},
	{
		Title: "GetPagesSize returns 1 when items exactly fill one page",
		ArrangeInput: args.Map{
			"when":         "given 5 items with eachPageSize=5",
			"items":        []int{1, 2, 3, 4, 5},
			"eachPageSize": 5,
		},
		ExpectedInput: "1",
	},
	{
		Title: "GetPagesSize returns 2 when items spill into second page",
		ArrangeInput: args.Map{
			"when":         "given 6 items with eachPageSize=5",
			"items":        []int{1, 2, 3, 4, 5, 6},
			"eachPageSize": 5,
		},
		ExpectedInput: "2",
	},
	{
		Title: "GetPagesSize returns 3 for 10 items with eachPageSize=4",
		ArrangeInput: args.Map{
			"when":         "given 10 items with eachPageSize=4",
			"items":        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			"eachPageSize": 4,
		},
		ExpectedInput: "3",
	},
	{
		Title: "GetPagesSize returns count when eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 4 items with eachPageSize=1",
			"items":        []int{1, 2, 3, 4},
			"eachPageSize": 1,
		},
		ExpectedInput: "4",
	},
	{
		Title: "GetPagesSize returns 1 when single item with eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 1 item with eachPageSize=1",
			"items":        []int{42},
			"eachPageSize": 1,
		},
		ExpectedInput: "1",
	},
}
