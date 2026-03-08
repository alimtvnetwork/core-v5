package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// StringCompare nil receiver test cases (migrated from CaseV1 string-dispatch)
// =============================================================================

var stringCompareNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Nil receiver - IsMatch returns true (vacuous truth)",
		Func:  (*coreinstruction.StringCompare).IsMatch,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver - IsMatchFailed returns false",
		Func:  (*coreinstruction.StringCompare).IsMatchFailed,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver - IsInvalid returns true",
		Func:  (*coreinstruction.StringCompare).IsInvalid,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver - IsDefined returns false",
		Func:  (*coreinstruction.StringCompare).IsDefined,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "Nil receiver - VerifyError returns nil",
		Func:  (*coreinstruction.StringCompare).VerifyError,
		Expected: args.Map{
			"panicked": false,
			"hasError": false,
		},
	},
}
