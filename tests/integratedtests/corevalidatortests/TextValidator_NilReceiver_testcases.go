package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
)

// =============================================================================
// TextValidator nil receiver test cases
// (migrated from inline t.Error tests in TextValidator_test.go)
// =============================================================================

var textValidatorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsMatchMany on nil returns true",
		Func:  (*corevalidator.TextValidator).IsMatchMany,
		Args:  []any{false, true, "anything"},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "VerifyDetailError on nil returns nil",
		Func:  (*corevalidator.TextValidator).VerifyDetailError,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}, "anything"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "VerifySimpleError on nil returns nil",
		Func:  (*corevalidator.TextValidator).VerifySimpleError,
		Args:  []any{0, &corevalidator.Parameter{CaseIndex: 0}, "anything"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
