package coredatatests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var stringsDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "StringsDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "StringsDsc empty slice Len returns 0",
		ArrangeInput: args.Map{
			"values": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "StringsDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"values": []string{"charlie", "alpha", "beta"},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var stringsDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "StringsDsc sorts descending",
		ArrangeInput: args.Map{
			"values": []string{"alpha", "charlie", "beta"},
		},
		ExpectedInput: args.Map{
			"first": "charlie",
			"last":  "alpha",
		},
	},
	{
		Title: "StringsDsc single element unchanged",
		ArrangeInput: args.Map{
			"values": []string{"only"},
		},
		ExpectedInput: args.Map{
			"first": "only",
			"last":  "only",
		},
	},
}
