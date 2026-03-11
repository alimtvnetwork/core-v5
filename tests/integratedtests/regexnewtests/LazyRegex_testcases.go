package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// New.Lazy / New.LazyLock — table-driven test cases
// ==========================================================================

var lazyRegexNewTestCases = []coretestcases.StringBoolGherkins{
	{
		Title:      "New.Lazy with simple word pattern matches correctly",
		When:       "given a simple word pattern",
		Input:      "hello",
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "hello world"},
		ExpectedLines: []string{
			"hello",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title:      "New.Lazy with digit pattern matches digits",
		When:       "given a digit pattern",
		Input:      "\\d+",
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "abc123def"},
		ExpectedLines: []string{
			"\\d+",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title:      "New.Lazy with no-match input returns false",
		When:       "given input that does not match",
		Input:      "^\\d+$",
		IsMatching: false,
		ExtraArgs:  map[string]any{"compareInput": "abc"},
		ExpectedLines: []string{
			"^\\d+$",
			"true",
			"true",
			"false",
			"true",
		},
	},
	{
		Title:      "New.Lazy with invalid regex pattern has error",
		When:       "given an invalid regex pattern",
		Input:      "[invalid",
		IsMatching: false,
		ExtraArgs:  map[string]any{"compareInput": "anything"},
		ExpectedLines: []string{
			"[invalid",
			"true",
			"false",
			"false",
			"true",
		},
	},
	{
		Title:      "New.Lazy with empty pattern matches everything",
		When:       "given an empty pattern",
		Input:      "",
		IsMatching: false,
		ExtraArgs:  map[string]any{"compareInput": "anything"},
		ExpectedLines: []string{
			"",
			"false",
			"false",
			"false",
			"true",
		},
	},
}

var lazyRegexLockTestCases = []coretestcases.StringBoolGherkins{
	{
		Title:      "New.LazyLock with word pattern is thread-safe",
		When:       "given a word pattern via LazyLock",
		Input:      "world",
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "hello world"},
		ExpectedLines: []string{
			"world",
			"true",
			"true",
			"true",
			"false",
		},
	},
	{
		Title:      "New.LazyLock with email pattern matches email",
		When:       "given an email-like pattern",
		Input:      `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
		IsMatching: true,
		ExtraArgs:  map[string]any{"compareInput": "user@example.com"},
		ExpectedLines: []string{
			`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			"true",
			"true",
			"true",
			"false",
		},
	},
}

// ==========================================================================
// PatternMatch — named test cases
// ==========================================================================

var lazyRegexIsMatchFullDigitTestCase = coretestcases.StringBoolGherkins{
	Title:      "IsMatch returns true for full string digit match",
	When:       "given full string digit pattern",
	Input:      "^\\d+$",
	Expected:   true,
	IsMatching: true,
	ExtraArgs:  map[string]any{"compareInput": "12345"},
	ExpectedLines: []string{
		"true",
	},
}

var lazyRegexIsMatchPartialMismatchTestCase = coretestcases.StringBoolGherkins{
	Title:      "IsMatch returns false for partial digit mismatch",
	When:       "given full string digit pattern with letters",
	Input:      "^\\d+$",
	Expected:   false,
	IsMatching: false,
	ExtraArgs:  map[string]any{"compareInput": "123abc"},
	ExpectedLines: []string{
		"false",
	},
}

var lazyRegexIsFailedMatchTestCase = coretestcases.StringBoolGherkins{
	Title:      "IsFailedMatch is inverse of IsMatch",
	When:       "given matching input to IsFailedMatch",
	Input:      "^hello$",
	Expected:   false,
	IsMatching: true,
	ExtraArgs:  map[string]any{"compareInput": "hello"},
	ExpectedLines: []string{
		"false",
	},
}

var lazyRegexFirstMatchLineFoundTestCase = coretestcases.StringGherkins{
	Title:     "FirstMatchLine returns first submatch",
	When:      "given a pattern with capture group",
	Input:     "(\\d+)",
	Expected:  "123",
	ExtraArgs: map[string]any{"compareInput": "abc 123 def 456"},
	ExpectedLines: []string{
		"123",
		"false",
	},
}

var lazyRegexFirstMatchLineNotFoundTestCase = coretestcases.StringGherkins{
	Title:     "FirstMatchLine returns empty on no match",
	When:      "given a pattern that does not match",
	Input:     "(\\d+)",
	Expected:  "",
	ExtraArgs: map[string]any{"compareInput": "no digits here"},
	ExpectedLines: []string{
		"",
		"true",
	},
}
