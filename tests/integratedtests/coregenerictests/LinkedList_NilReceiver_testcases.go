package coregenerictests

import (
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// LinkedList nil receiver test cases (migrated from single CaseV1)
// =============================================================================

var linkedListNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coregeneric.LinkedList[int]).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func:  (*coregeneric.LinkedList[int]).HasItems,
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*coregeneric.LinkedList[int]).Length,
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
}
