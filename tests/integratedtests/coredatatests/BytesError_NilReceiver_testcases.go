package coredatatests

import (
	"gitlab.com/auk-go/core/coredata"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// BytesError nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var bytesErrorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "String on nil returns empty",
		Func:  (*coredata.BytesError).String,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "HasError on nil returns false",
		Func:  (*coredata.BytesError).HasError,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsEmptyError on nil returns true",
		Func:  (*coredata.BytesError).IsEmptyError,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*coredata.BytesError).Length,
		Expected: args.Map{
			"value":    "0",
			"panicked": false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coredata.BytesError).IsEmpty,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "HandleError on nil does not panic",
		Func:  (*coredata.BytesError).HandleError,
		Expected: args.Map{
			"panicked":    false,
			"returnCount": 0,
		},
	},
}
