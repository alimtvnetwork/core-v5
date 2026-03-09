package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// MapAnyItems nil receiver test cases
// (migrated from standalone CaseV1 variables in MapAnyItemsEdge_testcases.go)
//
// Note: Some MapAnyItems nil tests require ArrangeInput for the right-side
// argument (e.g., IsEqualRaw), so they remain in CaseV1. Only pure
// nil-receiver zero-arg methods are migrated here.
// =============================================================================

var mapAnyItemsNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Length on nil returns 0",
		Func:  (*coredynamic.MapAnyItems).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coredynamic.MapAnyItems).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem on nil returns false",
		Func:  (*coredynamic.MapAnyItems).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
