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
		Title:         "GetFuncName returns short name of a function",
		ArrangeInput:  args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{"hasShortName": true, "fullLongerThanShort": true},
	},
}

// ==========================================
// ActionReturnsErrorFuncWrapper
// ==========================================

var actionErrWrapperSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:         "ActionReturnsErrorFuncWrapper.Exec returns nil on success",
		ArrangeInput:  args.Map{"when": "given action that succeeds"},
		ExpectedInput: args.Map{"isNil": true, "name": "cleanup"},
	},
}

var actionErrWrapperFailureTestCases = []coretestcases.CaseV1{
	{
		Title:         "ActionReturnsErrorFuncWrapper.Exec returns error on failure",
		ArrangeInput:  args.Map{"when": "given action that fails"},
		ExpectedInput: args.Map{"isNil": false, "hasError": true},
	},
}

// ==========================================
// IsSuccessFuncWrapper
// ==========================================

var isSuccessWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsSuccessFuncWrapper.Exec returns true on success",
		ArrangeInput:  args.Map{"when": "given action that returns true"},
		ExpectedInput: args.Map{"result": true, "name": "checker"},
	},
	{
		Title:         "IsSuccessFuncWrapper.Exec returns false on failure",
		ArrangeInput:  args.Map{"when": "given action that returns false"},
		ExpectedInput: args.Map{"result": false, "name": "checker"},
	},
}

// ==========================================
// InOutErrFuncWrapperOf (generic)
// ==========================================

var inOutErrWrapperOfSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:         "InOutErrFuncWrapperOf.Exec returns output on success",
		ArrangeInput:  args.Map{"when": "given string-to-int wrapper", "input": "hello"},
		ExpectedInput: args.Map{"result": 5, "isNil": true},
	},
}

var inOutErrWrapperOfFailureTestCases = []coretestcases.CaseV1{
	{
		Title:         "InOutErrFuncWrapperOf.Exec returns error on failure",
		ArrangeInput:  args.Map{"when": "given wrapper that returns error", "input": ""},
		ExpectedInput: args.Map{"result": 0, "isNil": false},
	},
}

// ==========================================
// NewCreator factory methods
// ==========================================

var newCreatorActionErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "New.ActionErr creates named wrapper correctly",
		ArrangeInput:  args.Map{"when": "given New.ActionErr factory"},
		ExpectedInput: args.Map{"name": "my-action", "hasAction": true},
	},
}

var newCreatorIsSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:         "New.IsSuccess creates named wrapper correctly",
		ArrangeInput:  args.Map{"when": "given New.IsSuccess factory"},
		ExpectedInput: args.Map{"name": "my-check", "hasAction": true},
	},
}
