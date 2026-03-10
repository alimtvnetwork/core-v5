package coredatatests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var pointerIntegersLenTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegers nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "PointerIntegers with elements Len returns count",
		ArrangeInput: args.Map{
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var pointerIntegersLessTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegers Less both non-nil ascending",
		ExpectedInput: args.Map{
			"lessIJ": true,
			"lessJI": false,
		},
	},
	{
		Title: "PointerIntegers Less nil-i returns true",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "PointerIntegers Less nil-j returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "PointerIntegers Less both nil returns true (nil-i check first)",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var pointerIntegersSortTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegers sort.Sort places nil first then ascending",
		ExpectedInput: args.Map{
			"firstIsNil": true,
			"second":     1,
			"last":       5,
		},
	},
}
