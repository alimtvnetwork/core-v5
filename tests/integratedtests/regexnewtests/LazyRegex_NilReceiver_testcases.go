package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/regexnew"
)

// =============================================================================
// LazyRegex nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var lazyRegexNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsNull on nil returns true",
		Func:  (*regexnew.LazyRegex).IsNull,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "IsUndefined on nil returns true",
		Func:  (*regexnew.LazyRegex).IsUndefined,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "IsDefined on nil returns false",
		Func:  (*regexnew.LazyRegex).IsDefined,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsCompiled on nil returns false",
		Func:  (*regexnew.LazyRegex).IsCompiled,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*regexnew.LazyRegex).String,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "Pattern on nil returns empty",
		Func:  (*regexnew.LazyRegex).Pattern,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "HasAnyIssues on nil returns true",
		Func:  (*regexnew.LazyRegex).HasAnyIssues,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*regexnew.LazyRegex).IsInvalid,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "OnRequiredCompiled on nil returns error",
		Func:  (*regexnew.LazyRegex).OnRequiredCompiled,
		Expected: args.Map{
			"panicked": false,
			"hasError": true,
		},
	},
	{
		Title: "FullString on nil returns empty",
		Func:  (*regexnew.LazyRegex).FullString,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
}
