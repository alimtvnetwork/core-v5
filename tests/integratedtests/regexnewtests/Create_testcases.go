package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var createTestCases = []coretestcases.StringGherkins{
	{
		Title: "Create with valid digit pattern compiles successfully",
		When:  "given a valid digit pattern",
		Input: "\\d+",
		ExpectedLines: []string{"true", "false"},
	},
	{
		Title: "Create with valid word boundary pattern compiles",
		When:  "given a word boundary pattern",
		Input: "\\bhello\\b",
		ExpectedLines: []string{"true", "false"},
	},
	{
		Title: "Create with invalid bracket pattern returns error",
		When:  "given an invalid pattern",
		Input: "[invalid",
		ExpectedLines: []string{"false", "true"},
	},
	{
		Title: "Create with empty pattern compiles as valid regex",
		When:  "given an empty pattern",
		Input: "",
		ExpectedLines: []string{"true", "false"},
	},
}

var createIsMatchLockTestCases = []coretestcases.StringGherkins{
	{
		Title:     "IsMatchLock returns true for matching digit pattern",
		When:      "given matching digit input",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "abc123"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:     "IsMatchLock returns false for non-matching pattern",
		When:      "given non-matching input",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:     "IsMatchLock returns false for invalid regex",
		When:      "given invalid regex pattern",
		Input:     "[bad",
		ExtraArgs: map[string]any{"compareInput": "anything"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:     "IsMatchLock returns true for exact full match",
		When:      "given exact match pattern",
		Input:     "^hello$",
		ExtraArgs: map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{"true"},
	},
}

var createIsMatchFailedTestCases = []coretestcases.StringGherkins{
	{
		Title:     "IsMatchFailed returns false when pattern matches",
		When:      "given matching input",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "123"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:     "IsMatchFailed returns true when pattern does not match",
		When:      "given non-matching input",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:     "IsMatchFailed returns true for invalid regex",
		When:      "given invalid regex",
		Input:     "[broken",
		ExtraArgs: map[string]any{"compareInput": "test"},
		ExpectedLines: []string{"true"},
	},
}

var matchErrorMatchTestCase = coretestcases.StringGherkins{
	Title:     "MatchError returns no error on match",
	When:      "given matching input to MatchError",
	Input:     "^hello$",
	ExtraArgs: map[string]any{"compareInput": "hello"},
	ExpectedLines: []string{"true"},
}

var matchErrorMismatchTestCase = coretestcases.StringGherkins{
	Title:     "MatchError returns error on mismatch",
	When:      "given non-matching input to MatchError",
	Input:     "^\\d+$",
	ExtraArgs: map[string]any{"compareInput": "abc"},
	ExpectedLines: []string{"false"},
}

var matchErrorLockMatchTestCase = coretestcases.StringGherkins{
	Title:     "MatchErrorLock returns no error on match",
	When:      "given matching input to MatchErrorLock",
	Input:     "world",
	ExtraArgs: map[string]any{"compareInput": "hello world"},
	ExpectedLines: []string{"true"},
}

var matchErrorLockMismatchTestCase = coretestcases.StringGherkins{
	Title:     "MatchErrorLock returns error on mismatch",
	When:      "given non-matching input to MatchErrorLock",
	Input:     "^xyz$",
	ExtraArgs: map[string]any{"compareInput": "abc"},
	ExpectedLines: []string{"false"},
}
