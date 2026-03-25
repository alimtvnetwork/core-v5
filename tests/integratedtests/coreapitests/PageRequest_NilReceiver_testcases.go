package coreapitests

import (
	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// PageRequest nil receiver test cases
// (migrated from first element of CaseV1 slices in PageRequest_testcases.go)
// =============================================================================

var pageRequestNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsPageSizeEmpty on nil returns true",
		Func:  (*coreapi.PageRequest).IsPageSizeEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsPageIndexEmpty on nil returns true",
		Func:  (*coreapi.PageRequest).IsPageIndexEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasPageSize on nil returns false",
		Func:  (*coreapi.PageRequest).HasPageSize,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasPageIndex on nil returns false",
		Func:  (*coreapi.PageRequest).HasPageIndex,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Clone on nil returns nil",
		Func: func(r *coreapi.PageRequest) bool {
			return r.Clone() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
