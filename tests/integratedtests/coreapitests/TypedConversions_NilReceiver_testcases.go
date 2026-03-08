package coreapitests

import (
	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// TypedSimpleGenericRequest nil receiver test cases
// (migrated from CaseV1 string-dispatch in TypedConversions_testcases.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// Renaming the method still causes a build error at the call site.
// =============================================================================

var typedSimpleGenericRequestNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Nil receiver IsValid returns false",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) bool {
			return r.IsValid()
		},
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver IsInvalid returns true",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) bool {
			return r.IsInvalid()
		},
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver Message returns empty string",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) string {
			return r.Message()
		},
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver InvalidError returns nil",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) error {
			return r.InvalidError()
		},
		Expected: args.Map{
			"panicked": false,
			"hasError": false,
		},
	},
}
