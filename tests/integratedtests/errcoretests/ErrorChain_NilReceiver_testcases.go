package errcoretests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// ConcatMessageWithErr nil passthrough test case
// (migrated from CaseV1 in ErrorChain_testcases.go)
//
// Note: ConcatMessageWithErr is a package function, not a method.
// We use a function literal wrapper that calls it with nil error.
// =============================================================================

var concatMessageNilSafeTestCases_v2 = []coretestcases.CaseNilSafe{
	{
		Title: "ConcatMessageWithErr nil returns nil",
		Func: func(_ *struct{}) bool {
			return errcore.ConcatMessageWithErr("should not appear", nil) == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
