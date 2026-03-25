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
		Title: "LazyRegex.IsNull returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsNull,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsUndefined returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsUndefined,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsDefined returns false -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsCompiled returns false -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsCompiled,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.String returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.Pattern returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).Pattern,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.HasAnyIssues returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).HasAnyIssues,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsInvalid returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.OnRequiredCompiled returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).OnRequiredCompiled,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "LazyRegex.FullString returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).FullString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.CompiledError returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).CompiledError,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "LazyRegex.Error returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).Error,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
}
