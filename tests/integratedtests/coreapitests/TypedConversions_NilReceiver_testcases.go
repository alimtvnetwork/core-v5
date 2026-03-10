package coreapitests

import (
	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
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
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver IsInvalid returns true",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) bool {
			return r.IsInvalid()
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver Message returns empty string",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) string {
			return r.Message()
		},
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver InvalidError returns nil",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) error {
			return r.InvalidError()
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
