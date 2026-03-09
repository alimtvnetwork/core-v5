package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
	"gitlab.com/auk-go/core/corestr"
)

// =============================================================================
// Hashmap.Clear nil receiver test case
// (migrated from CaseV1 in BugfixRegression_testcases.go)
// =============================================================================

var hashmapClearNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Clear on nil Hashmap returns nil without panic",
		Func: func(hm *corestr.Hashmap) bool {
			return hm.Clear() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
