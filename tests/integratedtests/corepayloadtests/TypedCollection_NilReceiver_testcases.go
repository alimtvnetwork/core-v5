package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// TypedPayloadCollection nil receiver test cases
// (migrated from CaseV1 in TypedCollectionFlatMap_testcases.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// Renaming the method still causes a build error at the call site.
// =============================================================================

var typedPayloadCollectionNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "TypedPayloadCollection.Length returns 0 -- nil receiver",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) int {
			return c.Length()
		},
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "TypedPayloadCollection.IsEmpty returns true -- nil receiver",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) bool {
			return c.IsEmpty()
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) bool {
			return c.HasItems()
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
