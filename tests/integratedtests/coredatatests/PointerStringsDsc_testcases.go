package coredatatests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var pointerStringsDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "PointerStringsDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var pointerStringsDscLessTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc Less both non-nil descending",
		ExpectedInput: args.Map{
			"lessAB": false,
			"lessBA": true,
		},
	},
	{
		Title: "PointerStringsDsc Less nil-i returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "PointerStringsDsc Less nil-j returns true",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "PointerStringsDsc Less both nil returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var pointerStringsDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc sort.Sort places values descending, nil last",
		ExpectedInput: args.Map{
			"first":     "charlie",
			"lastIsNil": true,
		},
	},
}
