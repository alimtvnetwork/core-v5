package corefuncstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// GetFuncName / GetFuncFullName
// ==========================================

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncName returns short name of a function",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

// ==========================================
// ActionReturnsErrorFuncWrapper
// ==========================================

var actionErrWrapperSuccessTestCases = []coretestcases.CaseV1{
	{
		Title: "ActionReturnsErrorFuncWrapper.Exec returns nil on success",
		ArrangeInput: args.Map{
			"when": "given action that succeeds",
		},
		ExpectedInput: []string{
			"true",
			"cleanup",
		},
	},
}

var actionErrWrapperFailureTestCases = []coretestcases.CaseV1{
	{
		Title: "ActionReturnsErrorFuncWrapper.Exec returns error on failure",
		ArrangeInput: args.Map{
			"when": "given action that fails",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

// ==========================================
// IsSuccessFuncWrapper
// ==========================================

var isSuccessWrapperTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSuccessFuncWrapper.Exec returns true on success",
		ArrangeInput: args.Map{
			"when": "given action that returns true",
		},
		ExpectedInput: []string{
			"true",
			"checker",
		},
	},
	{
		Title: "IsSuccessFuncWrapper.Exec returns false on failure",
		ArrangeInput: args.Map{
			"when": "given action that returns false",
		},
		ExpectedInput: []string{
			"false",
			"checker",
		},
	},
}

// ==========================================
// InOutErrFuncWrapperOf (generic)
// ==========================================

var inOutErrWrapperOfSuccessTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapperOf.Exec returns output on success",
		ArrangeInput: args.Map{
			"when":  "given string-to-int wrapper",
			"input": "hello",
		},
		ExpectedInput: []string{
			"5",
			"true",
		},
	},
}

var inOutErrWrapperOfFailureTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapperOf.Exec returns error on failure",
		ArrangeInput: args.Map{
			"when":  "given wrapper that returns error",
			"input": "",
		},
		ExpectedInput: []string{
			"0",
			"false",
		},
	},
}

// ==========================================
// NewCreator factory methods
// ==========================================

var newCreatorActionErrTestCases = []coretestcases.CaseV1{
	{
		Title: "New.ActionErr creates named wrapper correctly",
		ArrangeInput: args.Map{
			"when": "given New.ActionErr factory",
		},
		ExpectedInput: []string{
			"my-action",
			"true",
		},
	},
}

var newCreatorIsSuccessTestCases = []coretestcases.CaseV1{
	{
		Title: "New.IsSuccess creates named wrapper correctly",
		ArrangeInput: args.Map{
			"when": "given New.IsSuccess factory",
		},
		ExpectedInput: []string{
			"my-check",
			"true",
		},
	},
}
