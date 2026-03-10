package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
)

// =============================================================================
// BaseLinesValidators nil receiver test cases
// (migrated from inline t.Error tests in BaseLinesValidators_test.go)
// =============================================================================

var baseLinesValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "LinesValidatorsLength on nil returns 0",
		Func:  (*corevalidator.BaseLinesValidators).LinesValidatorsLength,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyLinesValidators on nil returns true",
		Func:  (*corevalidator.BaseLinesValidators).IsEmptyLinesValidators,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasLinesValidators on nil returns false",
		Func:  (*corevalidator.BaseLinesValidators).HasLinesValidators,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}

// =============================================================================
// LinesValidators nil receiver test cases
// (migrated from inline t.Error tests in BaseLinesValidators_test.go)
// =============================================================================

var linesValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Length on nil returns 0",
		Func:  (*corevalidator.LinesValidators).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*corevalidator.LinesValidators).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
