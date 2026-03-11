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
			params.pattern: "\\d+",
		},
		ExpectedInput: args.Map{
			params.isNotEmpty: true,
		},
	},
	{
		Title: "LazyRegex.FullString returns non-empty -- invalid pattern '[bad'",
		ArrangeInput: args.Map{
			params.pattern: "[bad",
		},
		ExpectedInput: args.Map{
			params.isNotEmpty: true,
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
			params.pattern: "\\w+",
		},
		ExpectedInput: args.Map{
			params.regexNotNil: true,
			params.panicked:    false,
		},
	},
	{
		Title: "LazyRegex.CompileMust returns panic -- invalid pattern '[bad'",
		ArrangeInput: args.Map{
			params.pattern: "[bad",
		},
		ExpectedInput: args.Map{
			params.regexNotNil: false,
			params.panicked:    true,
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
			params.pattern: "(\\d+)",
			params.content: "abc123def456",
		},
		ExpectedInput: args.Map{
			params.firstMatch:     "123",
			params.isInvalidMatch: false,
		},
	},
	{
		Title: "LazyRegex.FirstMatchLine returns empty and invalid -- no match",
		ArrangeInput: args.Map{
			params.pattern: "^\\d+$",
			params.content: "abc",
		},
		ExpectedInput: args.Map{
			params.firstMatch:     "",
			params.isInvalidMatch: true,
		},
	},
	{
		Title: "LazyRegex.FirstMatchLine returns empty and invalid -- invalid regex '[broken'",
		ArrangeInput: args.Map{
			params.pattern: "[broken",
			params.content: "test",
		},
		ExpectedInput: args.Map{
			params.firstMatch:     "",
			params.isInvalidMatch: true,
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
			params.pattern: "\\d+",
			params.input:   "abc123",
		},
		ExpectedInput: args.Map{
			params.isFailed: false,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatchBytes returns true -- non-matching bytes",
		ArrangeInput: args.Map{
			params.pattern: "^\\d+$",
			params.input:   "abc",
		},
		ExpectedInput: args.Map{
			params.isFailed: true,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatchBytes returns true -- invalid regex '[bad'",
		ArrangeInput: args.Map{
			params.pattern: "[bad",
			params.input:   "test",
		},
		ExpectedInput: args.Map{
			params.isFailed: true,
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
			params.pattern:   "^hello$",
			params.comparing: "hello",
		},
		ExpectedInput: args.Map{
			params.hasError: false,
		},
	},
	{
		Title: "LazyRegex.MatchUsingFuncError returns error -- non-matching input",
		ArrangeInput: args.Map{
			params.pattern:   "^\\d+$",
			params.comparing: "abc",
		},
		ExpectedInput: args.Map{
			params.hasError: true,
		},
	},
	{
		Title: "LazyRegex.MatchUsingFuncError returns error -- invalid regex '[broken'",
		ArrangeInput: args.Map{
			params.pattern:   "[broken",
			params.comparing: "test",
		},
		ExpectedInput: args.Map{
			params.hasError: true,
		},
	},
}

// =============================================================================
// LazyRegex.OnRequiredCompiledMust
// =============================================================================

var lazyRegexOnRequiredCompiledMustTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.OnRequiredCompiledMust returns no panic -- valid pattern",
		ArrangeInput: args.Map{
			params.pattern: "\\d+",
		},
		ExpectedInput: args.Map{
			params.panicked: false,
		},
	},
	{
		Title: "LazyRegex.OnRequiredCompiledMust returns panic -- invalid pattern '[bad'",
		ArrangeInput: args.Map{
			params.pattern: "[bad",
		},
		ExpectedInput: args.Map{
			params.panicked: true,
		},
	},
}

// =============================================================================
// LazyRegex.MustBeSafe
// =============================================================================

var lazyRegexMustBeSafeTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.MustBeSafe returns no panic -- valid pattern",
		ArrangeInput: args.Map{
			params.pattern: "\\d+",
		},
		ExpectedInput: args.Map{
			params.panicked: false,
		},
	},
	{
		Title: "LazyRegex.MustBeSafe returns panic -- invalid pattern '[bad'",
		ArrangeInput: args.Map{
			params.pattern: "[bad",
		},
		ExpectedInput: args.Map{
			params.panicked: true,
		},
	},
}
