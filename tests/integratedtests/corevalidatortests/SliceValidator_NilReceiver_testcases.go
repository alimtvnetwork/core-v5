package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
)

// =============================================================================
// SliceValidator nil receiver test cases
// (migrated from inline t.Error tests in SliceValidatorUnit_test.go
//  and SliceValidatorExtra_test.go)
// =============================================================================

var sliceValidatorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsValid on nil returns true",
		Func:  (*corevalidator.SliceValidator).IsValid,
		Args:  []any{true},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ActualLinesLength on nil returns 0",
		Func:  (*corevalidator.SliceValidator).ActualLinesLength,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "AllVerifyError on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyError,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "VerifyFirstError on nil returns nil",
		Func:  (*corevalidator.SliceValidator).VerifyFirstError,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "Dispose on nil does not panic",
		Func:  (*corevalidator.SliceValidator).Dispose,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "AllVerifyErrorExceptLast on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorExceptLast,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "AllVerifyErrorQuick on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorQuick,
		Args:  []any{0, "test", "a"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "AllVerifyErrorTestCase on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorTestCase,
		Args:  []any{0, "test", true},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "ActualLinesString on nil returns empty",
		Func:  (*corevalidator.SliceValidator).ActualLinesString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExpectingLinesString on nil returns empty",
		Func:  (*corevalidator.SliceValidator).ExpectingLinesString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "IsUsedAlready on nil returns false",
		Func:  (*corevalidator.SliceValidator).IsUsedAlready,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
