package coregenerictests

import (
	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Pair and Triple nil receiver test cases
// (migrated from inline tests in PairTripleExtended_test.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// =============================================================================

var pairNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Clear on nil Pair does not panic",
		Func: func(p *coregeneric.Pair[string, string]) {
			p.Clear()
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
}

var tripleNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Clear on nil Triple does not panic",
		Func: func(tr *coregeneric.Triple[string, string, string]) {
			tr.Clear()
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
}
