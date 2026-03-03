package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var lazyRegexCompileTestCases = []coretestcases.StringGherkins{
	{
		Title: "Compile valid pattern returns no error",
		When:  "given a valid pattern to Compile",
		Input: "\\d+",
		ExpectedLines: []string{
			"true",
			"false",
			"true",
		},
	},
	{
		Title: "Compile invalid pattern returns error",
		When:  "given an invalid pattern to Compile",
		Input: "[bad",
		ExpectedLines: []string{
			"false",
			"true",
			"false",
		},
	},
	{
		Title: "Compile empty pattern on undefined lazy returns error",
		When:  "given empty pattern to Compile",
		Input: "",
		ExpectedLines: []string{
			"false",
			"false",
			"false",
		},
	},
}

var lazyRegexHasErrorTestCases = []coretestcases.StringGherkins{
	{
		Title: "HasError returns false for valid pattern",
		When:  "given valid pattern for HasError",
		Input: "hello",
		ExpectedLines: []string{
			"false",
			"false",
		},
	},
	{
		Title: "HasError returns true for invalid pattern",
		When:  "given invalid pattern for HasError",
		Input: "[broken",
		ExpectedLines: []string{
			"true",
			"true",
		},
	},
}

var lazyRegexMatchBytesTestCases = []coretestcases.StringGherkins{
	{
		Title:     "IsMatchBytes returns true for matching bytes",
		When:      "given matching byte input",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "abc123"},
		ExpectedLines: []string{
			"true",
			"false",
		},
	},
	{
		Title:     "IsMatchBytes returns false for non-matching bytes",
		When:      "given non-matching byte input",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{
			"false",
			"true",
		},
	},
}

var lazyRegexMatchErrorTestCases = []coretestcases.StringGherkins{
	{
		Title:     "LazyRegex.MatchError returns nil on match",
		When:      "given matching input to LazyRegex.MatchError",
		Input:     "^hello$",
		ExtraArgs: map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{
			"true",
		},
	},
	{
		Title:     "LazyRegex.MatchError returns error on mismatch",
		When:      "given non-matching input to LazyRegex.MatchError",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{
			"false",
		},
	},
	{
		Title:     "LazyRegex.MatchError returns error for invalid regex",
		When:      "given invalid regex to LazyRegex.MatchError",
		Input:     "[bad",
		ExtraArgs: map[string]any{"compareInput": "test"},
		ExpectedLines: []string{
			"false",
		},
	},
}

var lazyRegexStringTestCases = []coretestcases.StringGherkins{
	{
		Title: "String returns pattern for valid LazyRegex",
		When:  "given valid pattern for String",
		Input: "\\d+",
		ExpectedLines: []string{
			"\\d+",
		},
	},
	{
		Title: "Pattern returns the original pattern",
		When:  "given email pattern for Pattern",
		Input: `[a-z]+@[a-z]+\.[a-z]+`,
		ExpectedLines: []string{
			`[a-z]+@[a-z]+\.[a-z]+`,
		},
	},
}
