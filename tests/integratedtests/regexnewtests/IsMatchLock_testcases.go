package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isMatchLockTestCases = []coretestcases.StringGherkins{
	{
		Title:         "IsMatchLock true for matching digit pattern",
		When:          "given digit pattern and numeric input",
		Input:         "\\d+",
		ExtraArgs:     map[string]any{"compareInput": "123"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:         "IsMatchLock false for non-matching pattern",
		When:          "given digit-only pattern and alpha input",
		Input:         "^\\d+$",
		ExtraArgs:     map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:         "IsMatchLock false for invalid pattern",
		When:          "given invalid regex pattern",
		Input:         "[bad",
		ExtraArgs:     map[string]any{"compareInput": "test"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:         "IsMatchLock true for email-like pattern",
		When:          "given email-like pattern",
		Input:         "^[a-zA-Z0-9]+@[a-zA-Z]+\\.[a-zA-Z]+$",
		ExtraArgs:     map[string]any{"compareInput": "user@example.com"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:         "IsMatchLock false for empty input with required pattern",
		When:          "given required pattern and empty input",
		Input:         "^\\d+$",
		ExtraArgs:     map[string]any{"compareInput": ""},
		ExpectedLines: []string{"false"},
	},
}

var isMatchFailedTestCases = []coretestcases.StringGherkins{
	{
		Title:         "IsMatchFailed false when pattern matches",
		When:          "given matching pattern",
		Input:         "\\d+",
		ExtraArgs:     map[string]any{"compareInput": "42"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:         "IsMatchFailed true when pattern does not match",
		When:          "given non-matching pattern",
		Input:         "^\\d+$",
		ExtraArgs:     map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:         "IsMatchFailed true for invalid pattern",
		When:          "given invalid pattern",
		Input:         "[bad",
		ExtraArgs:     map[string]any{"compareInput": "test"},
		ExpectedLines: []string{"true"},
	},
}

var isMatchLockLazyIsMatchTestCases = []coretestcases.StringGherkins{
	{
		Title:         "LazyRegex.IsMatch true for matching pattern",
		When:          "given valid pattern with matching input",
		Input:         "^hello$",
		ExtraArgs:     map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:         "LazyRegex.IsMatch false for non-matching input",
		When:          "given valid pattern with non-matching input",
		Input:         "^hello$",
		ExtraArgs:     map[string]any{"compareInput": "world"},
		ExpectedLines: []string{"false"},
	},
}

var isMatchLockCompileTestCases = []coretestcases.StringGherkins{
	{
		Title:         "LazyRegex compiles valid pattern without error",
		When:          "given valid pattern",
		Input:         "\\d+",
		ExpectedLines: []string{"true", "false", "true"},
	},
}

var isMatchLockIsFailedMatchTestCases = []coretestcases.StringGherkins{
	{
		Title:         "LazyRegex.IsFailedMatch false when matches",
		When:          "given matching input",
		Input:         "\\d+",
		ExtraArgs:     map[string]any{"compareInput": "123"},
		ExpectedLines: []string{"false"},
	},
	{
		Title:         "LazyRegex.IsFailedMatch true when not matches",
		When:          "given non-matching input",
		Input:         "^\\d+$",
		ExtraArgs:     map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"true"},
	},
}

var isMatchLockPatternStringTestCases = []coretestcases.StringGherkins{
	{
		Title:         "LazyRegex.Pattern returns original pattern",
		When:          "given a pattern",
		Input:         "^test\\d+$",
		ExpectedLines: []string{"^test\\d+$"},
	},
}

var isMatchLockMatchErrorTestCases = []coretestcases.StringGherkins{
	{
		Title:         "LazyRegex.MatchError nil for matching input",
		When:          "given matching input",
		Input:         "^hello$",
		ExtraArgs:     map[string]any{"compareInput": "hello"},
		ExpectedLines: []string{"true"},
	},
	{
		Title:         "LazyRegex.MatchError returns error for non-matching",
		When:          "given non-matching input",
		Input:         "^\\d+$",
		ExtraArgs:     map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{"false"},
	},
}
