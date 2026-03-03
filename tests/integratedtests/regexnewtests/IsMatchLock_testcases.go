package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var isMatchLockTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatchLock true for matching digit pattern",
		ArrangeInput: args.Map{
			"when":    "given digit pattern and numeric input",
			"pattern": "\\d+",
			"input":   "123",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsMatchLock false for non-matching pattern",
		ArrangeInput: args.Map{
			"when":    "given digit-only pattern and alpha input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsMatchLock false for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid regex pattern",
			"pattern": "[bad",
			"input":   "test",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsMatchLock true for email-like pattern",
		ArrangeInput: args.Map{
			"when":    "given email-like pattern",
			"pattern": "^[a-zA-Z0-9]+@[a-zA-Z]+\\.[a-zA-Z]+$",
			"input":   "user@example.com",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsMatchLock false for empty input with required pattern",
		ArrangeInput: args.Map{
			"when":    "given required pattern and empty input",
			"pattern": "^\\d+$",
			"input":   "",
		},
		ExpectedInput: "false",
	},
}

var isMatchFailedTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatchFailed false when pattern matches",
		ArrangeInput: args.Map{
			"when":    "given matching pattern",
			"pattern": "\\d+",
			"input":   "42",
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsMatchFailed true when pattern does not match",
		ArrangeInput: args.Map{
			"when":    "given non-matching pattern",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsMatchFailed true for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid pattern",
			"pattern": "[bad",
			"input":   "test",
		},
		ExpectedInput: "true",
	},
}

var lazyRegexIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.IsMatch true for matching pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with matching input",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "LazyRegex.IsMatch false for non-matching input",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with non-matching input",
			"pattern": "^hello$",
			"input":   "world",
		},
		ExpectedInput: "false",
	},
}

var lazyRegexCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex compiles valid pattern without error",
		ArrangeInput: args.Map{
			"when":    "given valid pattern",
			"pattern": "\\d+",
		},
		ExpectedInput: []string{"true", "false", "true"},
	},
}

var lazyRegexIsFailedMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.IsFailedMatch false when matches",
		ArrangeInput: args.Map{
			"when":    "given matching input",
			"pattern": "\\d+",
			"input":   "123",
		},
		ExpectedInput: "false",
	},
	{
		Title: "LazyRegex.IsFailedMatch true when not matches",
		ArrangeInput: args.Map{
			"when":    "given non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: "true",
	},
}

var lazyRegexPatternStringTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.Pattern returns original pattern",
		ArrangeInput: args.Map{
			"when":    "given a pattern",
			"pattern": "^test\\d+$",
		},
		ExpectedInput: "^test\\d+$",
	},
}

var lazyRegexMatchErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.MatchError nil for matching input",
		ArrangeInput: args.Map{
			"when":    "given matching input",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "LazyRegex.MatchError returns error for non-matching",
		ArrangeInput: args.Map{
			"when":    "given non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: "false",
	},
}
