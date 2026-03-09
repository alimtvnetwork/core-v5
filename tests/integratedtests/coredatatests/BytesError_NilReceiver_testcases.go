package coredatatests

import (
	"gitlab.com/auk-go/core/coredata"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// BytesError nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var bytesErrorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "String on nil returns empty",
		Func:  (*coredata.BytesError).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "HasError on nil returns false",
		Func:  (*coredata.BytesError).HasError,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyError on nil returns true",
		Func:  (*coredata.BytesError).IsEmptyError,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*coredata.BytesError).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coredata.BytesError).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HandleError on nil does not panic",
		Func:  (*coredata.BytesError).HandleError,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
}
