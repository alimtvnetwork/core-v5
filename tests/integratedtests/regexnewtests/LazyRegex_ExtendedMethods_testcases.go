package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// LazyRegex.FullString
// =============================================================================

var lazyRegexFullStringTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.FullString returns non-empty -- valid pattern '\\d+'",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
		},
		ExpectedInput: args.Map{
			"isNotEmpty": true,
		},
	},
	{
		Title: "LazyRegex.FullString returns non-empty -- invalid pattern '[bad'",
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
		Title: "LazyRegex.CompileMust returns regex without panic -- valid pattern '\\w+'",
		ArrangeInput: args.Map{
			"pattern": "\\w+",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"panicked":    false,
		},
	},
	{
		Title: "LazyRegex.CompileMust returns panic -- invalid pattern '[bad'",
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
		Title: "LazyRegex.FirstMatchLine returns '123' -- pattern '(\\d+)' content 'abc123def456'",
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
		Title: "LazyRegex.FirstMatchLine returns empty and invalid -- no match",
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
		Title: "LazyRegex.FirstMatchLine returns empty and invalid -- invalid regex '[broken'",
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
		Title: "LazyRegex.IsFailedMatchBytes returns false -- matching bytes '\\d+'",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
			"input":   "abc123",
		},
		ExpectedInput: args.Map{
			"isFailed": false,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatchBytes returns true -- non-matching bytes",
		ArrangeInput: args.Map{
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: args.Map{
			"isFailed": true,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatchBytes returns true -- invalid regex '[bad'",
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
		Title: "LazyRegex.MatchUsingFuncError returns no error -- matching input",
		ArrangeInput: args.Map{
			"pattern":   "^hello$",
			"comparing": "hello",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "LazyRegex.MatchUsingFuncError returns error -- non-matching input",
		ArrangeInput: args.Map{
			"pattern":   "^\\d+$",
			"comparing": "abc",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "LazyRegex.MatchUsingFuncError returns error -- invalid regex '[broken'",
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
