package coreindexestests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var hasIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "HasIndex returns true when index exists",
		ArrangeInput: args.Map{
			"when":    "given matching index",
			"indexes": []int{1, 3, 5},
			"current": 3,
		},
		ExpectedInput: "true",
	},
	{
		Title: "HasIndex returns false when index missing",
		ArrangeInput: args.Map{
			"when":    "given non-matching index",
			"indexes": []int{1, 3, 5},
			"current": 4,
		},
		ExpectedInput: "false",
	},
}

var lastIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndex returns length minus one",
		ArrangeInput: args.Map{
			"when":   "given length 5",
			"length": 5,
		},
		ExpectedInput: "4",
	},
	{
		Title: "LastIndex of length 1 is 0",
		ArrangeInput: args.Map{
			"when":   "given length 1",
			"length": 1,
		},
		ExpectedInput: "0",
	},
}

var isWithinIndexRangeTestCases = []coretestcases.CaseV1{
	{
		Title: "IsWithinIndexRange true for valid index",
		ArrangeInput: args.Map{
			"when":   "given index within range",
			"index":  2,
			"length": 5,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsWithinIndexRange false for out-of-range index",
		ArrangeInput: args.Map{
			"when":   "given index beyond range",
			"index":  5,
			"length": 5,
		},
		ExpectedInput: "false",
	},
}
