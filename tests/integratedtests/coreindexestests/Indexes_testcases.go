package coreindexestests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var hasIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "HasIndex returns true -- index 3 in [1,3,5]",
		ArrangeInput: args.Map{
			"when":    "given matching index",
			"indexes": []int{1, 3, 5},
			"current": 3,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "HasIndex returns false -- index 4 not in [1,3,5]",
		ArrangeInput: args.Map{
			"when":    "given non-matching index",
			"indexes": []int{1, 3, 5},
			"current": 4,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var lastIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndex returns 4 -- length 5",
		ArrangeInput: args.Map{
			"when":   "given length 5",
			"length": 5,
		},
		ExpectedInput: args.Map{"result": 4},
	},
	{
		Title: "LastIndex returns 0 -- length 1",
		ArrangeInput: args.Map{
			"when":   "given length 1",
			"length": 1,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var isWithinIndexRangeTestCases = []coretestcases.CaseV1{
	{
		Title: "IsWithinIndexRange returns true -- index 2 in length 5",
		ArrangeInput: args.Map{
			"when":   "given index within range",
			"index":  2,
			"length": 5,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IsWithinIndexRange returns false -- index 5 in length 5",
		ArrangeInput: args.Map{
			"when":   "given index beyond range",
			"index":  5,
			"length": 5,
		},
		ExpectedInput: args.Map{"result": false},
	},
}
