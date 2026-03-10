package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
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
