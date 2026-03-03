package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var createMustTestCases = []coretestcases.StringGherkins{
	{
		Title:     "CreateMust with valid digit pattern returns compiled regex",
		When:      "given a valid digit pattern",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "abc123"},
		ExpectedLines: []string{"true", "true"},
	},
	{
		Title:     "CreateMust with valid word pattern returns compiled regex",
		When:      "given a valid word pattern",
		Input:     "\\w+",
		ExtraArgs: map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{"true", "true"},
	},
	{
		Title:     "CreateMust with anchored pattern matches correctly",
		When:      "given an anchored pattern with non-matching input",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"true", "false"},
	},
}

var createMustLockIfTestCases = []coretestcases.StringBoolGherkins{
	{
		Title:      "CreateMustLockIf with lock true compiles valid pattern",
		When:       "given valid pattern with lock true",
		Input:      "\\d+",
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "99", "isLock": true},
		ExpectedLines: []string{"true", "true"},
	},
	{
		Title:      "CreateMustLockIf with lock false compiles valid pattern",
		When:       "given valid pattern with lock false",
		Input:      "[a-z]+",
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "hello", "isLock": false},
		ExpectedLines: []string{"true", "true"},
	},
}

var createLockIfTestCases = []coretestcases.StringBoolGherkins{
	{
		Title: "CreateLockIf with lock true compiles valid pattern",
		When:  "given valid pattern with lock true",
		Input: "\\d+",
		ExtraArgs: map[string]any{"isLock": true},
		ExpectedLines: []string{"true", "false"},
	},
	{
		Title: "CreateLockIf with lock false compiles valid pattern",
		When:  "given valid pattern with lock false",
		Input: "[a-z]+",
		ExtraArgs: map[string]any{"isLock": false},
		ExpectedLines: []string{"true", "false"},
	},
	{
		Title: "CreateLockIf with lock true returns error for invalid pattern",
		When:  "given invalid pattern with lock true",
		Input: "[bad",
		ExtraArgs: map[string]any{"isLock": true},
		ExpectedLines: []string{"false", "true"},
	},
	{
		Title: "CreateLockIf with lock false returns error for invalid pattern",
		When:  "given invalid pattern with lock false",
		Input: "(unclosed",
		ExtraArgs: map[string]any{"isLock": false},
		ExpectedLines: []string{"false", "true"},
	},
}

var createApplicableLockTestCases = []coretestcases.StringGherkins{
	{
		Title: "CreateApplicableLock with valid pattern is applicable",
		When:  "given a valid pattern",
		Input: "\\d+",
		ExpectedLines: []string{"true", "false", "true"},
	},
	{
		Title: "CreateApplicableLock with invalid pattern is not applicable",
		When:  "given an invalid pattern",
		Input: "[bad",
		ExpectedLines: []string{"false", "true", "false"},
	},
	{
		Title: "CreateApplicableLock with empty pattern is applicable",
		When:  "given an empty pattern",
		Input: "",
		ExpectedLines: []string{"true", "false", "true"},
	},
}

var newMustLockTestCases = []coretestcases.StringGherkins{
	{
		Title:     "NewMustLock with valid pattern returns compiled regex",
		When:      "given a valid digit pattern",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "123"},
		ExpectedLines: []string{"true", "true"},
	},
	{
		Title:     "NewMustLock with word boundary pattern matches",
		When:      "given a word boundary pattern",
		Input:     "\\bhello\\b",
		ExtraArgs: map[string]any{"compareInput": "hello world"},
		ExpectedLines: []string{"true", "true"},
	},
	{
		Title:     "NewMustLock with anchored pattern rejects mismatch",
		When:      "given an anchored pattern with non-matching input",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"true", "false"},
	},
}

var matchUsingFuncErrorLockTestCases = []coretestcases.StringGherkins{
	{
		Title:     "MatchUsingFuncErrorLock returns nil on match",
		When:      "given matching input with MatchString func",
		Input:     "^hello$",
		ExtraArgs: map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:     "MatchUsingFuncErrorLock returns error on mismatch",
		When:      "given non-matching input with MatchString func",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:     "MatchUsingFuncErrorLock returns error for invalid pattern",
		When:      "given invalid pattern with MatchString func",
		Input:     "[bad",
		ExtraArgs: map[string]any{"compareInput": "test"},
		ExpectedLines: []string{"false"},
	},
}

var matchUsingCustomizeErrorFuncLockTestCases = []coretestcases.StringGherkins{
	{
		Title:     "CustomizeErrorFunc returns nil on match with nil customizer",
		When:      "given matching input with nil customizer",
		Input:     "^hello$",
		ExtraArgs: map[string]any{"compareInput": "hello", "customizer": "nil"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:     "CustomizeErrorFunc returns default error on mismatch with nil customizer",
		When:      "given non-matching input with nil customizer",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc", "customizer": "nil"},
		ExpectedLines: []string{"false", "false"},
	},
	{
		Title:     "CustomizeErrorFunc returns custom error on mismatch",
		When:      "given non-matching input with custom error func",
		Input:     "^\\d+$",
		ExtraArgs: map[string]any{"compareInput": "abc", "customizer": "custom"},
		ExpectedLines: []string{"false", "true"},
	},
	{
		Title:     "CustomizeErrorFunc returns nil on match with custom error func",
		When:      "given matching input with custom error func",
		Input:     "\\d+",
		ExtraArgs: map[string]any{"compareInput": "123", "customizer": "custom"},
		ExpectedLines: []string{"true"},
	},
}
