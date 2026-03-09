package casenilsafetests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Nil-safe pointer receiver methods (should NOT panic)
// =============================================================================

var nilSafePointerReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsValid on nil returns false",
		Func:  (*sampleStruct).IsValid,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*sampleStruct).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*sampleStruct).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
}

// =============================================================================
// Void methods (no return values)
// =============================================================================

var nilSafeVoidTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Reset on nil does not panic",
		Func:  (*sampleStruct).Reset,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
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
		Expected: results.ResultAny{
			Value:       "0",
			Panicked:    false,
			ReturnCount: 2,
		},
		CompareFields: []string{"value", "panicked", "hasError", "returnCount"},
	},
	{
		Title: "Lookup on nil returns empty false",
		Func:  (*sampleStruct).Lookup,
		Args:  []any{"key"},
		Expected: results.ResultAny{
			Value:       "",
			Panicked:    false,
			ReturnCount: 2,
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
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
	{
		Title: "ValueString on nil panics (value receiver)",
		Func:  (*sampleStruct).ValueString,
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
}

// =============================================================================
// MethodName extraction — uses CompareFields for non-standard assertions
// =============================================================================

var methodNameTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "MethodName extracts IsValid",
		Func:  (*sampleStruct).IsValid,
		Expected: results.ResultAny{
			Value: "IsValid",
		},
	},
	{
		Title: "MethodName extracts Parse",
		Func:  (*sampleStruct).Parse,
		Args:  []any{"x"},
		Expected: results.ResultAny{
			Value: "Parse",
		},
	},
	{
		Title: "MethodName extracts Reset",
		Func:  (*sampleStruct).Reset,
		Expected: results.ResultAny{
			Value: "Reset",
		},
	},
}
