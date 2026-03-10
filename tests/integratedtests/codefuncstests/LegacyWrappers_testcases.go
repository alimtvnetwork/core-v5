package codefuncstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IsSuccessFuncWrapper — Exec
// =============================================================================

var isSuccessExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns true when action succeeds",
		ArrangeInput: args.Map{
			"actionResult": true,
			"name":         "check-health",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "Exec returns false when action fails",
		ArrangeInput: args.Map{
			"actionResult": false,
			"name":         "check-health",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

// =============================================================================
// IsSuccessFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var isSuccessAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"actionResult": true,
			"name":         "check-ok",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns error on failure",
		ArrangeInput: args.Map{
			"actionResult": false,
			"name":         "check-fail",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
}

// =============================================================================
// NamedActionFuncWrapper — Exec
// =============================================================================

var namedActionExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec calls action with wrapper name",
		ArrangeInput: args.Map{
			"name": "my-task",
		},
		ExpectedInput: args.Map{
			"calledWith": "my-task",
		},
	},
}

// =============================================================================
// NamedActionFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var namedActionAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc always returns nil",
		ArrangeInput: args.Map{
			"name": "log-step",
		},
		ExpectedInput: args.Map{
			"hasError":   false,
			"calledWith": "log-step",
		},
	},
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — Exec
// =============================================================================

var actionReturnsErrorExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns nil when action succeeds",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "cleanup",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "Exec returns error when action fails",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "cleanup",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var actionReturnsErrorAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name on failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "deploy",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "deploy",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapper — Exec
// =============================================================================

var inOutErrExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns output on success",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": false,
			"name":         "transform",
		},
		ExpectedInput: args.Map{
			"output":   "HELLO",
			"hasError": false,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": true,
			"name":         "transform",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var inOutErrAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": true,
			"name":         "process",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": false,
			"name":         "process",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapper — Exec
// =============================================================================

var resultDelegatingExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns nil on success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "unmarshal",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "unmarshal",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var resultDelegatingAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "deserialize",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "deserialize",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}
