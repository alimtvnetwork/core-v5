package corevalidatortests

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
	"gitlab.com/auk-go/core/corevalidator"
)

// =============================================================================
// LineValidator nil receiver test cases
// (migrated from inline t.Error tests in LineValidator_test.go)
// =============================================================================

var lineValidatorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsMatchMany on nil returns true",
		Func:  (*corevalidator.LineValidator).IsMatchMany,
		Args:  []any{false, true, corestr.TextWithLineNumber{Text: "x"}},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
