package corestrtests

import (
	"github.com/alimtvnetwork/core/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
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
