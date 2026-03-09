package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// IdentifiersWithGlobals nil receiver test cases
// (migrated from CaseV1 in IdentifiersWithGlobals_testcases.go)
// =============================================================================

var identifiersWithGlobalsNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Length on nil returns 0",
		Func:  (*coreinstruction.IdentifiersWithGlobals).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
}
