package corefuncstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// GetFuncName
// ==========================================================================

var ext2GetFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetFuncName returns short name of named function",
		ArrangeInput:  args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// GetFuncFullName
// ==========================================================================

var ext2GetFuncFullNameTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetFuncFullName returns full qualified name",
		ArrangeInput:  args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// NewInOutErrWrapper
// ==========================================================================

var ext2NewInOutErrWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewInOutErrWrapper.Exec returns output and nil error -- success",
		ArrangeInput: args.Map{"when": "given valid input"},
		ExpectedInput: args.Map{
			"result": 5,
			"isNil":  true,
			"name":   "parse",
		},
	},
}

var ext2NewInOutErrWrapperFailTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewInOutErrWrapper.Exec returns error -- failure",
		ArrangeInput: args.Map{"when": "given empty input"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================================================
// InOutErrFuncWrapperOf.ToLegacy
// ==========================================================================

var ext2InOutErrToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "InOutErrFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given generic wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name":   "legacy-parse",
			"result": "5",
			"isNil":  true,
		},
	},
}

// ==========================================================================
// SerializeOutputFuncWrapperOf.ToLegacy
// ==========================================================================

var ext2SerializeToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "SerializeOutputFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given serialize wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name":     "legacy-serialize",
			"hasBytes": true,
			"isNil":    true,
		},
	},
}
