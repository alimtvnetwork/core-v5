package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var anyCollectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyCollection GetPagesSize returns 0 for zero eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=0",
			"count":        5,
			"eachPageSize": 0,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "AnyCollection GetPagesSize returns 0 for negative eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=-2",
			"count":        3,
			"eachPageSize": -2,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "AnyCollection GetPagesSize returns 0 for empty collection",
		ArrangeInput: args.Map{
			"when":         "given empty collection",
			"count":        0,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "AnyCollection GetPagesSize returns 1 when items fit one page",
		ArrangeInput: args.Map{
			"when":         "given 3 items with eachPageSize=5",
			"count":        3,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "AnyCollection GetPagesSize returns 1 for exact page fill",
		ArrangeInput: args.Map{
			"when":         "given 5 items with eachPageSize=5",
			"count":        5,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "AnyCollection GetPagesSize returns 2 for spill into second page",
		ArrangeInput: args.Map{
			"when":         "given 6 items with eachPageSize=5",
			"count":        6,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"2"},
	},
	{
		Title: "AnyCollection GetPagesSize returns count when eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 4 items with eachPageSize=1",
			"count":        4,
			"eachPageSize": 1,
		},
		ExpectedInput: []string{"4"},
	},
}

var dynamicCollectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "DynamicCollection GetPagesSize returns 0 for zero eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=0",
			"count":        5,
			"eachPageSize": 0,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "DynamicCollection GetPagesSize returns 0 for negative eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=-1",
			"count":        3,
			"eachPageSize": -1,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "DynamicCollection GetPagesSize returns 0 for empty collection",
		ArrangeInput: args.Map{
			"when":         "given empty collection",
			"count":        0,
			"eachPageSize": 4,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "DynamicCollection GetPagesSize returns 1 for exact page fill",
		ArrangeInput: args.Map{
			"when":         "given 4 items with eachPageSize=4",
			"count":        4,
			"eachPageSize": 4,
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "DynamicCollection GetPagesSize returns 3 for ceiling division",
		ArrangeInput: args.Map{
			"when":         "given 10 items with eachPageSize=4",
			"count":        10,
			"eachPageSize": 4,
		},
		ExpectedInput: []string{"3"},
	},
	{
		Title: "DynamicCollection GetPagesSize returns count when eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 3 items with eachPageSize=1",
			"count":        3,
			"eachPageSize": 1,
		},
		ExpectedInput: []string{"3"},
	},
}

var keyValCollectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "KeyValCollection GetPagesSize returns 0 for zero eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=0",
			"count":        5,
			"eachPageSize": 0,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "KeyValCollection GetPagesSize returns 0 for negative eachPageSize",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=-5",
			"count":        3,
			"eachPageSize": -5,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "KeyValCollection GetPagesSize returns 0 for empty collection",
		ArrangeInput: args.Map{
			"when":         "given empty collection",
			"count":        0,
			"eachPageSize": 3,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "KeyValCollection GetPagesSize returns 1 when items under page size",
		ArrangeInput: args.Map{
			"when":         "given 2 items with eachPageSize=5",
			"count":        2,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "KeyValCollection GetPagesSize returns 2 for spill into second page",
		ArrangeInput: args.Map{
			"when":         "given 7 items with eachPageSize=5",
			"count":        7,
			"eachPageSize": 5,
		},
		ExpectedInput: []string{"2"},
	},
	{
		Title: "KeyValCollection GetPagesSize returns count when eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 5 items with eachPageSize=1",
			"count":        5,
			"eachPageSize": 1,
		},
		ExpectedInput: []string{"5"},
	},
}
