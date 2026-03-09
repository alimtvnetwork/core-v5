package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// CastedResult nil receiver test cases
// (migrated from first element of CaseV1 slices in CastedResult_testcases.go)
// =============================================================================

var castedResultNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsInvalid true on nil receiver",
		Func:  (*coredynamic.CastedResult).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsNotNull false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotNull,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsNotPointer false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotPointer,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsNotMatchingAcceptedType false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotMatchingAcceptedType,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasError false on nil receiver",
		Func:  (*coredynamic.CastedResult).HasError,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyIssues true on nil receiver",
		Func:  (*coredynamic.CastedResult).HasAnyIssues,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
