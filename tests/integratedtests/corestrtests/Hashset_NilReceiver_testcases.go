package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coredata/corestr"
)

// =============================================================================
// Hashset nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var hashsetNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*corestr.Hashset).IsEmpty,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*corestr.Hashset).Length,
		Expected: args.Map{
			"value":    "0",
			"panicked": false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func:  (*corestr.Hashset).HasItems,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "HasAnyItem on nil returns false",
		Func:  (*corestr.Hashset).HasAnyItem,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "Clear on nil returns nil",
		Func:  (*corestr.Hashset).Clear,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
}
