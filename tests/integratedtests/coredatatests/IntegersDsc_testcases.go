package coredatatests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var integersDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegersDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "IntegersDsc empty slice Len returns 0",
		ArrangeInput: args.Map{
			"values": []int{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "IntegersDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"values": []int{5, 3, 8},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var integersDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegersDsc sorts descending",
		ArrangeInput: args.Map{
			"values": []int{1, 5, 3, 2, 4},
		},
		ExpectedInput: args.Map{
			"first": 5,
			"last":  1,
		},
	},
	{
		Title: "IntegersDsc single element unchanged",
		ArrangeInput: args.Map{
			"values": []int{42},
		},
		ExpectedInput: args.Map{
			"first": 42,
			"last":  42,
		},
	},
	{
		Title: "IntegersDsc already sorted stays same",
		ArrangeInput: args.Map{
			"values": []int{5, 4, 3, 2, 1},
		},
		ExpectedInput: args.Map{
			"first": 5,
			"last":  1,
		},
	},
}
