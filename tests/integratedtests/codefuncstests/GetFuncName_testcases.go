package codefuncstests

import (
	"fmt"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// GetFuncName
// =============================================================================

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncName returns name for a function",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// GetFuncFullName
// =============================================================================

var getFuncFullNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncFullName returns full name for a function",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: args.Map{
			"isNotEmpty":      true,
			"containsPackage": true,
		},
	},
}

// =============================================================================
// GetFunc
// =============================================================================

var getFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFunc returns non-nil for a function",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

// =============================================================================
// newCreator — factory methods
// =============================================================================

var newCreatorTestCases = []coretestcases.CaseV1{
	{
		Title: "New.ActionErr creates wrapper with correct name",
		ArrangeInput: args.Map{
			"method": "ActionErr",
			"name":   "test-action",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "New.IsSuccess creates wrapper that returns success",
		ArrangeInput: args.Map{
			"method": "IsSuccess",
			"name":   "test-check",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "New.NamedAction creates wrapper that calls with name",
		ArrangeInput: args.Map{
			"method": "NamedAction",
			"name":   "test-named",
		},
		ExpectedInput: args.Map{
			"calledWith": "test-named",
		},
	},
	{
		Title: "New.LegacyInOutErr creates wrapper that returns output",
		ArrangeInput: args.Map{
			"method": "LegacyInOutErr",
			"name":   "test-inout",
		},
		ExpectedInput: args.Map{
			"output":   "processed",
			"hasError": false,
		},
	},
	{
		Title: "New.LegacyResultDelegating creates wrapper that delegates",
		ArrangeInput: args.Map{
			"method": "LegacyResultDelegating",
			"name":   "test-delegate",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// Ensure fmt is used
var _ = fmt.Sprintf
