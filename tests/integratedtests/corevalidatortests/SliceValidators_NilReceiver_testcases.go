package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
)

// =============================================================================
// SliceValidators nil receiver test cases
// (migrated from inline t.Error tests in SliceValidators_test.go)
// =============================================================================

var sliceValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Length on nil returns 0",
		Func:  (*corevalidator.SliceValidators).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*corevalidator.SliceValidators).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsMatch on nil returns true",
		Func:  (*corevalidator.SliceValidators).IsMatch,
		Args:  []any{true},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
