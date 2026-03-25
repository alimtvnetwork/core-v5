package codefuncstests

import (
	"fmt"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// GetFuncName — positive + negative
// =============================================================================

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncName returns non-empty name -- named function input",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: "true",
	},
}

var getFuncNameNilTestCase = coretestcases.CaseV1{
	Title: "GetFuncName returns empty string -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncName returns empty string -- non-function (int) input",
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
		Title: "GetFuncFullName returns full package-qualified name -- named function input",
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
	Title: "GetFuncFullName returns empty string -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncFullNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncFullName returns empty string -- non-function (string) input",
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
		Title: "GetFunc returns non-nil -- named function input",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

var getFuncNilTestCase = coretestcases.CaseV1{
	Title: "GetFunc returns nil -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"isNil":    true,
		"panicked": false,
	},
}

var getFuncNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFunc returns nil -- non-function (struct) input",
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
		Title: "New.ActionErr returns wrapper with correct name -- 'test-action' factory",
		ArrangeInput: args.Map{
			"method": "ActionErr",
			"name":   "test-action",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "New.IsSuccess returns wrapper that returns true -- 'test-check' factory",
		ArrangeInput: args.Map{
			"method": "IsSuccess",
			"name":   "test-check",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "New.NamedAction returns wrapper that calls with name -- 'test-named' factory",
		ArrangeInput: args.Map{
			"method": "NamedAction",
			"name":   "test-named",
		},
		ExpectedInput: args.Map{
			"calledWith": "test-named",
		},
	},
	{
		Title: "New.LegacyInOutErr returns wrapper with output 'processed' -- 'test-inout' factory",
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
		Title: "New.LegacyResultDelegating returns wrapper without error -- 'test-delegate' factory",
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
