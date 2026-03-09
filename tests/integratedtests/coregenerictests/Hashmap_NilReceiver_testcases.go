package coregenerictests

import (
	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Generic Hashmap nil receiver test cases
// (migrated from CaseV1 in Hashmap_testcases.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// =============================================================================

var hashmapNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty on nil returns true",
		Func: func(hm *coregeneric.Hashmap[string, int]) bool {
			return hm.IsEmpty()
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Length on nil returns 0",
		Func: func(hm *coregeneric.Hashmap[string, int]) int {
			return hm.Length()
		},
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func: func(hm *coregeneric.Hashmap[string, int]) bool {
			return hm.HasItems()
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
