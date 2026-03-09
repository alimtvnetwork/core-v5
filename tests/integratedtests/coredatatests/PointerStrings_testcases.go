package coredatatests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var pointerStringsLenTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStrings nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "PointerStrings with elements Len returns count",
		ArrangeInput: args.Map{
			"count": 2,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var pointerStringsLessTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStrings Less both non-nil alpha < beta",
		ExpectedInput: args.Map{
			"less01": true,
			"less10": false,
		},
	},
	{
		Title: "PointerStrings Less nil-first returns true",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "PointerStrings Less nil-second returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "PointerStrings Less both-nil returns true (first nil check)",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var pointerStringsSortTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStrings sort.Sort nil first then ascending",
		ExpectedInput: args.Map{
			"firstIsNil": true,
			"second":     "alpha",
			"third":      "beta",
			"fourth":     "charlie",
		},
	},
}
