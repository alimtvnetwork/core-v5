package corefuncstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// NamedActionFuncWrapper
// ==========================================

var namedActionWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "NamedActionFuncWrapper.Exec calls action with name",
		ArrangeInput: args.Map{"when": "given named action wrapper"},
		ExpectedInput: args.Map{
			"called": true,
			"name":   "my-named",
		},
	},
}

var namedActionAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:        "NamedActionFuncWrapper.AsActionFunc returns callable",
		ArrangeInput: args.Map{"when": "given wrapper converted to ActionFunc"},
		ExpectedInput: args.Map{
			"called": true,
		},
	},
}

var namedActionAsErrFuncTestCases = []coretestcases.CaseV1{
	{
		Title:        "NamedActionFuncWrapper.AsActionReturnsErrorFunc returns nil",
		ArrangeInput: args.Map{"when": "given wrapper converted to error func"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var namedActionNextTestCases = []coretestcases.CaseV1{
	{
		Title:        "NamedActionFuncWrapper.Next calls both actions",
		ArrangeInput: args.Map{"when": "given two named actions chained"},
		ExpectedInput: args.Map{
			"firstCalled":  true,
			"secondCalled": true,
		},
	},
}

// ==========================================
// InOutErrFuncWrapper (legacy)
// ==========================================

var legacyInOutErrExecTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapper.Exec returns output -- valid input",
		ArrangeInput: args.Map{
			"when":  "given legacy wrapper with valid input",
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"result": "5",
			"isNil":  true,
		},
	},
}

var legacyInOutErrAsErrFuncSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:        "InOutErrFuncWrapper.AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{"when": "given wrapper that succeeds"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var legacyInOutErrAsErrFuncFailTestCases = []coretestcases.CaseV1{
	{
		Title:        "InOutErrFuncWrapper.AsActionReturnsErrorFunc returns error on failure",
		ArrangeInput: args.Map{"when": "given wrapper that fails"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================
// ResultDelegatingFuncWrapper (legacy)
// ==========================================

var legacyResultDelegatingExecTestCases = []coretestcases.CaseV1{
	{
		Title:        "ResultDelegatingFuncWrapper.Exec delegates to target -- success",
		ArrangeInput: args.Map{"when": "given wrapper that succeeds"},
		ExpectedInput: args.Map{
			"isNil":  true,
			"result": "delegated",
		},
	},
}

var legacyResultDelegatingAsErrFuncSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:        "ResultDelegatingFuncWrapper.AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{"when": "given wrapper that succeeds"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var legacyResultDelegatingAsErrFuncFailTestCases = []coretestcases.CaseV1{
	{
		Title:        "ResultDelegatingFuncWrapper.AsActionReturnsErrorFunc returns error on failure",
		ArrangeInput: args.Map{"when": "given wrapper that fails"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================
// GetFunc / GetFuncFullName
// ==========================================

var getFuncRuntimeTestCases = []coretestcases.CaseV1{
	{
		Title:        "GetFunc returns non-nil runtime.Func for named function",
		ArrangeInput: args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{
			"notNil": true,
		},
	},
}

// ==========================================
// Generic Wrapper constructors - additional
// ==========================================

var newInOutWrapperTestCases = []coretestcases.CaseV1{
	{
		Title: "NewInOutWrapper.Exec returns output -- valid input",
		ArrangeInput: args.Map{
			"when":  "given string-to-int wrapper",
			"input": "test",
		},
		ExpectedInput: args.Map{
			"result": 4,
			"name":   "strlen",
		},
	},
}

var newInActionErrWrapperSuccessTestCases = []coretestcases.CaseV1{
	{
		Title: "NewInActionErrWrapper.Exec returns nil -- valid input",
		ArrangeInput: args.Map{
			"when":  "given validate wrapper with valid input",
			"input": "valid",
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var newInActionErrWrapperFailTestCases = []coretestcases.CaseV1{
	{
		Title: "NewInActionErrWrapper.Exec returns error -- empty input",
		ArrangeInput: args.Map{
			"when":  "given validate wrapper with empty input",
			"input": "",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

var newResultDelegatingWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewResultDelegatingWrapper.Exec delegates -- success",
		ArrangeInput: args.Map{"when": "given typed result delegating wrapper"},
		ExpectedInput: args.Map{
			"isNil":  true,
			"result": "typed-delegated",
		},
	},
}

var newSerializeWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewSerializeWrapper.Exec returns bytes -- success",
		ArrangeInput: args.Map{"when": "given serialize wrapper"},
		ExpectedInput: args.Map{
			"hasBytes": true,
			"isNil":    true,
		},
	},
}

// ==========================================
// IsSuccessFuncWrapper additional methods
// ==========================================

var isSuccessAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:        "IsSuccessFuncWrapper.AsActionFunc returns callable",
		ArrangeInput: args.Map{"when": "given success wrapper"},
		ExpectedInput: args.Map{
			"called": true,
		},
	},
}

var isSuccessAsErrFuncSuccessTestCases = []coretestcases.CaseV1{
	{
		Title:        "IsSuccessFuncWrapper.AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{"when": "given wrapper that returns true"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var isSuccessAsErrFuncFailTestCases = []coretestcases.CaseV1{
	{
		Title:        "IsSuccessFuncWrapper.AsActionReturnsErrorFunc returns error on failure",
		ArrangeInput: args.Map{"when": "given wrapper that returns false"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================
// ActionReturnsErrorFuncWrapper additional
// ==========================================

var actionErrAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:        "ActionReturnsErrorFuncWrapper.AsActionReturnsErrorFunc nil on success",
		ArrangeInput: args.Map{"when": "given wrapper that succeeds"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}

var actionErrAsErrFuncWithErrorTestCases = []coretestcases.CaseV1{
	{
		Title:        "ActionReturnsErrorFuncWrapper.AsActionReturnsErrorFunc wraps error",
		ArrangeInput: args.Map{"when": "given wrapper that fails"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================
// Generic ToLegacy converters
// ==========================================

var inOutWrapperToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "InOutFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given generic wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name":   "legacy-conv",
			"result": "5",
		},
	},
}

var inActionErrToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "InActionReturnsErrFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given generic wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name": "legacy-validate",
		},
	},
}

var resultDelegatingToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "ResultDelegatingFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given generic wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name": "legacy-unmarshal",
		},
	},
}
