package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// FromTo nil receiver test cases
// (migrated from CaseV1 variables in FromTo_testcases.go)
// =============================================================================

var fromToNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "ClonePtr on nil returns nil",
		Func: func(ft *coreinstruction.FromTo) bool {
			return ft.ClonePtr() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsNull on nil returns true",
		Func:  (*coreinstruction.FromTo).IsNull,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsFromEmpty on nil returns true",
		Func:  (*coreinstruction.FromTo).IsFromEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "SetFromName on nil does not panic",
		Func: func(ft *coreinstruction.FromTo) {
			ft.SetFromName("x")
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "SourceDestination on nil returns nil",
		Func: func(ft *coreinstruction.FromTo) bool {
			return ft.SourceDestination() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Rename on nil returns nil",
		Func: func(ft *coreinstruction.FromTo) bool {
			return ft.Rename() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
