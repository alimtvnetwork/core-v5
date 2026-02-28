package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var lazyRegexNewTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Lazy with simple word pattern matches correctly",
		ArrangeInput: args.Map{
			"when":     "given a simple word pattern",
			"pattern":  "hello",
			"input":    "hello world",
			"isMatch":  true,
		},
		ExpectedInput: []string{
			"hello",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title: "New.Lazy with digit pattern matches digits",
		ArrangeInput: args.Map{
			"when":     "given a digit pattern",
			"pattern":  "\\d+",
			"input":    "abc123def",
			"isMatch":  true,
		},
		ExpectedInput: []string{
			"\\d+",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title: "New.Lazy with no-match input returns false",
		ArrangeInput: args.Map{
			"when":     "given input that does not match",
			"pattern":  "^\\d+$",
			"input":    "abc",
			"isMatch":  false,
		},
		ExpectedInput: []string{
			"^\\d+$",
			"true",
			"true",
			"false",
			"true",
		},
	},
	{
		Title: "New.Lazy with invalid regex pattern has error",
		ArrangeInput: args.Map{
			"when":     "given an invalid regex pattern",
			"pattern":  "[invalid",
			"input":    "anything",
			"isMatch":  false,
		},
		ExpectedInput: []string{
			"[invalid",
			"true",
			"false",
			"false",
			"true",
		},
	},
	{
		Title: "New.Lazy with empty pattern matches everything",
		ArrangeInput: args.Map{
			"when":     "given an empty pattern",
			"pattern":  "",
			"input":    "anything",
			"isMatch":  false,
		},
		ExpectedInput: []string{
			"",
			"false",
			"false",
			"false",
			"false",
		},
	},
}

var lazyRegexLockTestCases = []coretestcases.CaseV1{
	{
		Title: "New.LazyLock with word pattern is thread-safe",
		ArrangeInput: args.Map{
			"when":     "given a word pattern via LazyLock",
			"pattern":  "world",
			"input":    "hello world",
			"isMatch":  true,
		},
		ExpectedInput: []string{
			"world",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title: "New.LazyLock with email pattern matches email",
		ArrangeInput: args.Map{
			"when":     "given an email-like pattern",
			"pattern":  `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			"input":    "user@example.com",
			"isMatch":  true,
		},
		ExpectedInput: []string{
			`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			"true",
			"true",
			"true",
			"false",
		},
	},
}

var lazyRegexPatternMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatch returns true for full string digit match",
		ArrangeInput: args.Map{
			"when":    "given full string digit pattern",
			"pattern": "^\\d+$",
			"input":   "12345",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsMatch returns false for partial digit mismatch",
		ArrangeInput: args.Map{
			"when":    "given full string digit pattern with letters",
			"pattern": "^\\d+$",
			"input":   "123abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsFailedMatch is inverse of IsMatch",
		ArrangeInput: args.Map{
			"when":    "given matching input to IsFailedMatch",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "FirstMatchLine returns first submatch",
		ArrangeInput: args.Map{
			"when":    "given a pattern with capture group",
			"pattern": "(\\d+)",
			"input":   "abc 123 def 456",
		},
		ExpectedInput: []string{
			"123",
			"false",
		},
	},
	{
		Title: "FirstMatchLine returns empty on no match",
		ArrangeInput: args.Map{
			"when":    "given a pattern that does not match",
			"pattern": "(\\d+)",
			"input":   "no digits here",
		},
		ExpectedInput: []string{
			"",
			"true",
		},
	},
}
