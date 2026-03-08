package coreoncetests

import (
	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// BytesErrorOnce nil receiver test cases
// (migrated from IsNilReceiver flag in bytesErrorOnceTestCase wrapper)
// =============================================================================

var bytesErrorOnceNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasIssuesOrEmpty on nil returns true",
		Func:  (*coreonce.BytesErrorOnce).HasIssuesOrEmpty,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "HasSafeItems on nil returns false",
		Func:  (*coreonce.BytesErrorOnce).HasSafeItems,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coreonce.BytesErrorOnce).IsEmpty,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "IsDefined on nil returns false",
		Func:  (*coreonce.BytesErrorOnce).IsDefined,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
}
