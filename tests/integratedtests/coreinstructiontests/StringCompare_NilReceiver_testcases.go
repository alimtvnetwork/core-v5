package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// StringCompare nil receiver test cases (migrated from CaseV1 string-dispatch)
// =============================================================================

var stringCompareNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Nil receiver - IsMatch returns true (vacuous truth)",
		Func:  (*coreinstruction.StringCompare).IsMatch,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsMatchFailed returns false",
		Func:  (*coreinstruction.StringCompare).IsMatchFailed,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsInvalid returns true",
		Func:  (*coreinstruction.StringCompare).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsDefined returns false",
		Func:  (*coreinstruction.StringCompare).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - VerifyError returns nil",
		Func:  (*coreinstruction.StringCompare).VerifyError,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
