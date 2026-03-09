package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// LazyRegex.FullString
// =============================================================================

var lazyRegexFullStringTestCases = []coretestcases.CaseV1{
	{
		Title: "FullString on valid pattern returns non-empty JSON",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
		},
		ExpectedInput: args.Map{
			"isNotEmpty": true,
		},
	},
	{
		Title: "FullString on invalid pattern returns JSON with error info",
		ArrangeInput: args.Map{
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"isNotEmpty": true,
		},
	},
}

// =============================================================================
// LazyRegex.CompileMust
// =============================================================================

var lazyRegexCompileMustTestCases = []coretestcases.CaseV1{
	{
		Title: "CompileMust valid pattern returns regex",
		ArrangeInput: args.Map{
			"pattern": "\\w+",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"panicked":    false,
		},
	},
	{
		Title: "CompileMust invalid pattern panics",
		ArrangeInput: args.Map{
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"regexNotNil": false,
			"panicked":    true,
		},
	},
}

// =============================================================================
// LazyRegex.FirstMatchLine
// =============================================================================

var lazyRegexFirstMatchLineTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstMatchLine returns first match",
		ArrangeInput: args.Map{
			"pattern": "(\\d+)",
			"content": "abc123def456",
		},
		ExpectedInput: args.Map{
			"firstMatch":     "123",
			"isInvalidMatch": false,
		},
	},
	{
		Title: "FirstMatchLine no match returns empty and invalid",
		ArrangeInput: args.Map{
			"pattern": "^\\d+$",
			"content": "abc",
		},
		ExpectedInput: args.Map{
			"firstMatch":     "",
			"isInvalidMatch": true,
		},
	},
	{
		Title: "FirstMatchLine invalid regex returns empty and invalid",
		ArrangeInput: args.Map{
			"pattern": "[broken",
			"content": "test",
		},
		ExpectedInput: args.Map{
			"firstMatch":     "",
			"isInvalidMatch": true,
		},
	},
}

// =============================================================================
// LazyRegex.IsFailedMatchBytes
// =============================================================================

var lazyRegexIsFailedMatchBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "IsFailedMatchBytes false when bytes match",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
			"input":   "abc123",
		},
		ExpectedInput: args.Map{
			"isFailed": false,
		},
	},
	{
		Title: "IsFailedMatchBytes true when bytes do not match",
		ArrangeInput: args.Map{
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: args.Map{
			"isFailed": true,
		},
	},
	{
		Title: "IsFailedMatchBytes true for invalid regex",
		ArrangeInput: args.Map{
			"pattern": "[bad",
			"input":   "test",
		},
		ExpectedInput: args.Map{
			"isFailed": true,
		},
	},
}

// =============================================================================
// LazyRegex.MatchUsingFuncError
// =============================================================================

var lazyRegexMatchUsingFuncErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "MatchUsingFuncError nil on match",
		ArrangeInput: args.Map{
			"pattern":   "^hello$",
			"comparing": "hello",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "MatchUsingFuncError error on mismatch",
		ArrangeInput: args.Map{
			"pattern":   "^\\d+$",
			"comparing": "abc",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "MatchUsingFuncError error on invalid regex",
		ArrangeInput: args.Map{
			"pattern":   "[broken",
			"comparing": "test",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// LazyRegex.OnRequiredCompiledMust
// =============================================================================

var lazyRegexOnRequiredCompiledMustTestCases = []coretestcases.CaseV1{
	{
		Title: "OnRequiredCompiledMust valid pattern no panic",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
		},
		ExpectedInput: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "OnRequiredCompiledMust invalid pattern panics",
		ArrangeInput: args.Map{
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"panicked": true,
		},
	},
}

// =============================================================================
// LazyRegex.MustBeSafe
// =============================================================================

var lazyRegexMustBeSafeTestCases = []coretestcases.CaseV1{
	{
		Title: "MustBeSafe valid pattern no panic",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
		},
		ExpectedInput: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "MustBeSafe invalid pattern panics",
		ArrangeInput: args.Map{
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"panicked": true,
		},
	},
}
