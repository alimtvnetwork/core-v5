package namevaluetests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/namevalue"
)

// =============================================================================
// Collection nil receiver test cases (migrated from single CaseV1)
// =============================================================================

var collectionNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "ClonePtr on nil returns nil",
		Func:  (*namevalue.StringStringCollection).ClonePtr,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
}
