package coredatatests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var integersLenTestCases = []coretestcases.CaseV1{
	{
		Title: "Integers nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "Integers empty slice Len returns 0",
		ArrangeInput: args.Map{
			"values": []int{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "Integers with elements Len returns count",
		ArrangeInput: args.Map{
			"values": []int{3, 1, 2},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var integersLessTestCases = []coretestcases.CaseV1{
	{
		Title: "Integers Less 3 < 5 returns true",
		ExpectedInput: args.Map{
			"less10": true,
			"less01": false,
			"less00": false,
		},
	},
}

var integersSortTestCases = []coretestcases.CaseV1{
	{
		Title: "Integers sort.Sort ascending",
		ArrangeInput: args.Map{
			"values": []int{5, 1, 4, 2, 3},
		},
		ExpectedInput: args.Map{
			"first": 1,
			"last":  5,
		},
	},
}
