package coreoncetests

import (
	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
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
