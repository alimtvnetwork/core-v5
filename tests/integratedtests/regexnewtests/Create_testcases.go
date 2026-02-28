package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var createTestCases = []coretestcases.CaseV1{
	{
		Title: "Create with valid digit pattern compiles successfully",
		ArrangeInput: args.Map{
			"when":    "given a valid digit pattern",
			"pattern": "\\d+",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "Create with valid word boundary pattern compiles",
		ArrangeInput: args.Map{
			"when":    "given a word boundary pattern",
			"pattern": "\\bhello\\b",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "Create with invalid bracket pattern returns error",
		ArrangeInput: args.Map{
			"when":    "given an invalid pattern",
			"pattern": "[invalid",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "Create with empty pattern compiles as valid regex",
		ArrangeInput: args.Map{
			"when":    "given an empty pattern",
			"pattern": "",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

var isMatchLockTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatchLock returns true for matching digit pattern",
		ArrangeInput: args.Map{
			"when":    "given matching digit input",
			"pattern": "\\d+",
			"input":   "abc123",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsMatchLock returns false for non-matching pattern",
		ArrangeInput: args.Map{
			"when":    "given non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsMatchLock returns false for invalid regex",
		ArrangeInput: args.Map{
			"when":    "given invalid regex pattern",
			"pattern": "[bad",
			"input":   "anything",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsMatchLock returns true for exact full match",
		ArrangeInput: args.Map{
			"when":    "given exact match pattern",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

var isMatchFailedTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatchFailed returns false when pattern matches",
		ArrangeInput: args.Map{
			"when":    "given matching input",
			"pattern": "\\d+",
			"input":   "123",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsMatchFailed returns true when pattern does not match",
		ArrangeInput: args.Map{
			"when":    "given non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsMatchFailed returns true for invalid regex",
		ArrangeInput: args.Map{
			"when":    "given invalid regex",
			"pattern": "[broken",
			"input":   "test",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

var matchErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "MatchError returns no error on match",
		ArrangeInput: args.Map{
			"when":    "given matching input to MatchError",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "MatchError returns error on mismatch",
		ArrangeInput: args.Map{
			"when":    "given non-matching input to MatchError",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "MatchErrorLock returns no error on match",
		ArrangeInput: args.Map{
			"when":    "given matching input to MatchErrorLock",
			"pattern": "world",
			"input":   "hello world",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "MatchErrorLock returns error on mismatch",
		ArrangeInput: args.Map{
			"when":    "given non-matching input to MatchErrorLock",
			"pattern": "^xyz$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
