package coreoncetests

import (
	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// BytesErrorOnce nil receiver test cases
// (migrated from IsNilReceiver flag in bytesErrorOnceTestCase wrapper)
// =============================================================================

var bytesErrorOnceNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasIssuesOrEmpty on nil returns true",
		Func:  (*coreonce.BytesErrorOnce).HasIssuesOrEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasSafeItems on nil returns false",
		Func:  (*coreonce.BytesErrorOnce).HasSafeItems,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coreonce.BytesErrorOnce).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsDefined on nil returns false",
		Func:  (*coreonce.BytesErrorOnce).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
