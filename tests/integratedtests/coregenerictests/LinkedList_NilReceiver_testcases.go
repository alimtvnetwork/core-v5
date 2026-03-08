package coregenerictests

import (
	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// LinkedList nil receiver test cases (migrated from single CaseV1)
// =============================================================================

var linkedListNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coregeneric.LinkedList[int]).IsEmpty,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func:  (*coregeneric.LinkedList[int]).HasItems,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*coregeneric.LinkedList[int]).Length,
		Expected: args.Map{
			"value":    "0",
			"panicked": false,
		},
	},
}
