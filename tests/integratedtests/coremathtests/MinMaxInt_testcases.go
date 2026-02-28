package coremathtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var maxIntTestCases = []coretestcases.CaseV1{
	{
		Title: "MaxInt returns larger of two positives",
		ArrangeInput: args.Map{
			"when": "given 3 and 7",
			"a":    3,
			"b":    7,
		},
		ExpectedInput: []string{"7"},
	},
	{
		Title: "MaxInt returns equal when same",
		ArrangeInput: args.Map{
			"when": "given 5 and 5",
			"a":    5,
			"b":    5,
		},
		ExpectedInput: []string{"5"},
	},
	{
		Title: "MaxInt handles negatives",
		ArrangeInput: args.Map{
			"when": "given -3 and -7",
			"a":    -3,
			"b":    -7,
		},
		ExpectedInput: []string{"-3"},
	},
}

var minIntTestCases = []coretestcases.CaseV1{
	{
		Title: "MinInt returns smaller of two positives",
		ArrangeInput: args.Map{
			"when": "given 3 and 7",
			"a":    3,
			"b":    7,
		},
		ExpectedInput: []string{"3"},
	},
	{
		Title: "MinInt returns equal when same",
		ArrangeInput: args.Map{
			"when": "given 5 and 5",
			"a":    5,
			"b":    5,
		},
		ExpectedInput: []string{"5"},
	},
}
