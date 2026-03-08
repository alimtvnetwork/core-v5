package casenilsafetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// Nil-safe pointer receiver methods (should NOT panic)
// =============================================================================

var nilSafePointerReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsValid on nil returns false",
		Func:  (*sampleStruct).IsValid,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
			"isSafe":   true,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*sampleStruct).Length,
		Expected: args.Map{
			"value":    "0",
			"panicked": false,
			"isSafe":   true,
		},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*sampleStruct).String,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
			"isSafe":   true,
		},
	},
}

// =============================================================================
// Void methods (no return values)
// =============================================================================

var nilSafeVoidTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Reset on nil does not panic",
		Func:  (*sampleStruct).Reset,
		Expected: args.Map{
			"panicked":    false,
			"returnCount": 0,
		},
	},
}

// =============================================================================
// Multi-return methods
// =============================================================================

var nilSafeMultiReturnTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Parse on nil returns (0, nil)",
		Func:  (*sampleStruct).Parse,
		Args:  []any{"hello"},
		Expected: args.Map{
			"value":       "0",
			"panicked":    false,
			"hasError":    false,
			"returnCount": 2,
		},
	},
	{
		Title: "Lookup on nil returns empty false",
		Func:  (*sampleStruct).Lookup,
		Args:  []any{"key"},
		Expected: args.Map{
			"value":       "",
			"panicked":    false,
			"returnCount": 2,
		},
	},
}

// =============================================================================
// Unsafe methods (SHOULD panic on nil)
// =============================================================================

var nilUnsafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "UnsafeMethod on nil panics",
		Func:  (*sampleStruct).UnsafeMethod,
		Expected: args.Map{
			"panicked": true,
		},
	},
	{
		Title: "ValueString on nil panics (value receiver)",
		Func:  (*sampleStruct).ValueString,
		Expected: args.Map{
			"panicked": true,
		},
	},
}

// =============================================================================
// MethodName extraction
// =============================================================================

var methodNameTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "MethodName extracts IsValid",
		Func:  (*sampleStruct).IsValid,
		Expected: args.Map{
			"methodName": "IsValid",
		},
	},
	{
		Title: "MethodName extracts Parse",
		Func:  (*sampleStruct).Parse,
		Args:  []any{"x"},
		Expected: args.Map{
			"methodName": "Parse",
		},
	},
	{
		Title: "MethodName extracts Reset",
		Func:  (*sampleStruct).Reset,
		Expected: args.Map{
			"methodName": "Reset",
		},
	},
}
