package corestrtests

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Hashset nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var hashsetNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*corestr.Hashset).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*corestr.Hashset).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func:  (*corestr.Hashset).HasItems,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem on nil returns false",
		Func:  (*corestr.Hashset).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Clear on nil returns nil",
		Func:  (*corestr.Hashset).Clear,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
}
