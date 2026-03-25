package corefuncstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// GetFuncName / GetFuncFullName
// ==========================================

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title:        "GetFuncName returns short name -- named function input",
		ArrangeInput: args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{
			"hasShortName":        true,
			"fullLongerThanShort": true,
		},
	},
}

// ==========================================
// ActionReturnsErrorFuncWrapper
// ==========================================

var actionErrWrapperSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:        "ActionReturnsErrorFuncWrapper.Exec returns nil -- successful action",
		ArrangeInput: args.Map{"when": "given action that succeeds"},
		ExpectedInput: args.Map{
			"isNil": true,
			"name":  "cleanup",
		},
	},
}

var actionErrWrapperFailureTestCases = []coretestcases.CaseV1{
	{
		Title:        "ActionReturnsErrorFuncWrapper.Exec returns error -- failing action",
		ArrangeInput: args.Map{"when": "given action that fails"},
		ExpectedInput: args.Map{
			"isNil":    false,
			"hasError": true,
		},
	},
}

// ==========================================
// IsSuccessFuncWrapper
// ==========================================

var isSuccessWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "IsSuccessFuncWrapper.Exec returns true -- action returns true",
		ArrangeInput: args.Map{"when": "given action that returns true"},
		ExpectedInput: args.Map{
			"result": true,
			"name":   "checker",
		},
	},
	{
		Title:        "IsSuccessFuncWrapper.Exec returns false -- action returns false",
		ArrangeInput: args.Map{"when": "given action that returns false"},
		ExpectedInput: args.Map{
			"result": false,
			"name":   "checker",
		},
	},
}

// ==========================================
// InOutErrFuncWrapperOf (generic)
// ==========================================

var inOutErrWrapperOfSuccessTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapperOf.Exec returns output 5 -- string 'hello' input",
		ArrangeInput: args.Map{
			"when":  "given string-to-int wrapper",
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"result": 5,
			"isNil":  true,
		},
	},
}

var inOutErrWrapperOfFailureTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapperOf.Exec returns error -- empty string input",
		ArrangeInput: args.Map{
			"when":  "given wrapper that returns error",
			"input": "",
		},
		ExpectedInput: args.Map{
			"result": 0,
			"isNil":  false,
		},
	},
}

// ==========================================
// NewCreator factory methods
// ==========================================

var newCreatorActionErrTestCases = []coretestcases.CaseV1{
	{
		Title:        "New.ActionErr returns named wrapper -- 'my-action' factory",
		ArrangeInput: args.Map{"when": "given New.ActionErr factory"},
		ExpectedInput: args.Map{
			"name":      "my-action",
			"hasAction": true,
		},
	},
}

var newCreatorIsSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:        "New.IsSuccess returns named wrapper -- 'my-check' factory",
		ArrangeInput: args.Map{"when": "given New.IsSuccess factory"},
		ExpectedInput: args.Map{
			"name":      "my-check",
			"hasAction": true,
		},
	},
}
