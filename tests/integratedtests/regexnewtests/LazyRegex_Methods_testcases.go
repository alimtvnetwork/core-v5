package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var lazyRegexCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "Compile valid pattern returns no error",
		ArrangeInput: args.Map{
			"when":    "given a valid pattern to Compile",
			"pattern": "\\d+",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
		},
	},
	{
		Title: "Compile invalid pattern returns error",
		ArrangeInput: args.Map{
			"when":    "given an invalid pattern to Compile",
			"pattern": "[bad",
		},
		ExpectedInput: []string{
			"false",
			"true",
			"false",
		},
	},
	{
		Title: "Compile empty pattern on undefined lazy returns error",
		ArrangeInput: args.Map{
			"when":    "given empty pattern to Compile",
			"pattern": "",
		},
		ExpectedInput: []string{
			"false",
			"false",
			"false",
		},
	},
}

var lazyRegexHasErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "HasError returns false for valid pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern for HasError",
			"pattern": "hello",
		},
		ExpectedInput: []string{
			"false",
			"false",
		},
	},
	{
		Title: "HasError returns true for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid pattern for HasError",
			"pattern": "[broken",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var lazyRegexMatchBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatchBytes returns true for matching bytes",
		ArrangeInput: args.Map{
			"when":    "given matching byte input",
			"pattern": "\\d+",
			"input":   "abc123",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "IsMatchBytes returns false for non-matching bytes",
		ArrangeInput: args.Map{
			"when":    "given non-matching byte input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

var lazyRegexMatchErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex.MatchError returns nil on match",
		ArrangeInput: args.Map{
			"when":    "given matching input to LazyRegex.MatchError",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "LazyRegex.MatchError returns error on mismatch",
		ArrangeInput: args.Map{
			"when":    "given non-matching input to LazyRegex.MatchError",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "LazyRegex.MatchError returns error for invalid regex",
		ArrangeInput: args.Map{
			"when":    "given invalid regex to LazyRegex.MatchError",
			"pattern": "[bad",
			"input":   "test",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}

var lazyRegexStringTestCases = []coretestcases.CaseV1{
	{
		Title: "String returns pattern for valid LazyRegex",
		ArrangeInput: args.Map{
			"when":    "given valid pattern for String",
			"pattern": "\\d+",
		},
		ExpectedInput: []string{
			"\\d+",
		},
	},
	{
		Title: "Pattern returns the original pattern",
		ArrangeInput: args.Map{
			"when":    "given email pattern for Pattern",
			"pattern": `[a-z]+@[a-z]+\.[a-z]+`,
		},
		ExpectedInput: []string{
			`[a-z]+@[a-z]+\.[a-z]+`,
		},
	},
}
