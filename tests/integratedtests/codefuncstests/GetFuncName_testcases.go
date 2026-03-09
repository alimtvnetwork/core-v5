package codefuncstests

import (
	"fmt"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// GetFuncName — positive + negative
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

var getFuncNameNilTestCase = coretestcases.CaseV1{
	Title: "GetFuncName with nil returns empty string",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncName with non-function returns empty string",
	ArrangeInput: args.Map{
		"when":  "given an integer instead of a function",
		"input": 42,
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

// =============================================================================
// GetFuncFullName — positive + negative
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

var getFuncFullNameNilTestCase = coretestcases.CaseV1{
	Title: "GetFuncFullName with nil returns empty string",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncFullNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncFullName with non-function returns empty string",
	ArrangeInput: args.Map{
		"when":  "given a string instead of a function",
		"input": "not-a-func",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

// =============================================================================
// GetFunc — positive + negative
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

var getFuncNilTestCase = coretestcases.CaseV1{
	Title: "GetFunc with nil returns nil",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"isNil":    true,
		"panicked": false,
	},
}

var getFuncNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFunc with non-function returns nil",
	ArrangeInput: args.Map{
		"when":  "given a struct instead of a function",
		"input": struct{ Name string }{"test"},
	},
	ExpectedInput: args.Map{
		"isNil":    true,
		"panicked": false,
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
