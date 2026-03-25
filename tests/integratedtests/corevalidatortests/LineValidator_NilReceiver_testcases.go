package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
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
