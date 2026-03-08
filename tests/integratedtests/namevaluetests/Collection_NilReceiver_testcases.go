package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/namevalue"
)

// =============================================================================
// Collection nil receiver test cases (migrated from single CaseV1)
// =============================================================================

var collectionNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "ClonePtr on nil returns nil",
		Func:  (*namevalue.StringStringCollection).ClonePtr,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
}
