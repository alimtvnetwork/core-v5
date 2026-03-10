package coredatatests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var pointerIntegersDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegersDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "PointerIntegersDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"count": 2,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var pointerIntegersDscLessTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegersDsc Less both non-nil descending",
		ExpectedInput: args.Map{
			"lessIJ": false,
			"lessJI": true,
		},
	},
	{
		Title: "PointerIntegersDsc Less nil-i returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "PointerIntegersDsc Less nil-j returns true",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "PointerIntegersDsc Less both nil returns false (nil-i check first)",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var pointerIntegersDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerIntegersDsc sort.Sort places values descending, nil last",
		ExpectedInput: args.Map{
			"first":     5,
			"lastIsNil": true,
		},
	},
}
