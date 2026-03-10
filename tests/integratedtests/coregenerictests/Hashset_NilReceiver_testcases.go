package coregenerictests

import (
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// Generic Hashset nil receiver test cases
// (migrated from inline t.Error tests in Hashset_test.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// =============================================================================

var hashsetNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func: func(hs *coregeneric.Hashset[string]) bool {
			return hs.IsEmpty()
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func: func(hs *coregeneric.Hashset[string]) int {
			return hs.Length()
		},
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func: func(hs *coregeneric.Hashset[string]) bool {
			return hs.HasItems()
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
