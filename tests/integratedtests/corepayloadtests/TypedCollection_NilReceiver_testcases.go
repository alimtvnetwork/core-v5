package corepayloadtests

import (
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		Title: "Length on nil returns 0",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) int {
			return c.Length()
		},
		Expected: args.Map{
			"value":    "0",
			"panicked": false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) bool {
			return c.IsEmpty()
		},
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "HasItems on nil returns false",
		Func: func(c *corepayload.TypedPayloadCollection[testUser]) bool {
			return c.HasItems()
		},
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
}
