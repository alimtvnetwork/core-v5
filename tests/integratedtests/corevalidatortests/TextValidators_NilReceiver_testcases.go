package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
	"gitlab.com/auk-go/core/corevalidator"
)

// =============================================================================
// TextValidators nil receiver test cases
// (migrated from inline t.Error tests in TextValidators_test.go)
// =============================================================================

var textValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Dispose on nil does not panic",
		Func:  (*corevalidator.TextValidators).Dispose,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*corevalidator.TextValidators).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "VerifyErrorMany on nil returns nil",
		Func:  (*corevalidator.TextValidators).VerifyErrorMany,
		Args:  []any{true, &corevalidator.Parameter{CaseIndex: 0}, "a"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
