package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// LazyRegex nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var lazyRegexNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsNull on nil returns true",
		Func:  (*regexnew.LazyRegex).IsNull,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsUndefined on nil returns true",
		Func:  (*regexnew.LazyRegex).IsUndefined,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsDefined on nil returns false",
		Func:  (*regexnew.LazyRegex).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsCompiled on nil returns false",
		Func:  (*regexnew.LazyRegex).IsCompiled,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*regexnew.LazyRegex).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "Pattern on nil returns empty",
		Func:  (*regexnew.LazyRegex).Pattern,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyIssues on nil returns true",
		Func:  (*regexnew.LazyRegex).HasAnyIssues,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*regexnew.LazyRegex).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "OnRequiredCompiled on nil returns error",
		Func:  (*regexnew.LazyRegex).OnRequiredCompiled,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "FullString on nil returns empty",
		Func:  (*regexnew.LazyRegex).FullString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "CompiledError on nil returns error",
		Func:  (*regexnew.LazyRegex).CompiledError,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "Error on nil returns error",
		Func:  (*regexnew.LazyRegex).Error,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
}
