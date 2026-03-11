package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// New.Lazy / New.LazyLock — table-driven test cases
// ==========================================================================

var lazyRegexNewTestCases = []coretestcases.MapGherkins{
	{
		Title: "New.Lazy with simple word pattern matches correctly",
		When:  "given a simple word pattern",
		Input: args.Map{
			params.pattern:      "hello",
			params.compareInput: "hello world",
		},
		Expected: args.Map{
			params.patternResult: "hello",
			params.isDefined:     true,
			params.isApplicable:  true,
			params.isMatch:       true,
			params.isFailedMatch: false,
		},
	},
	{
		Title: "New.Lazy with digit pattern matches digits",
		When:  "given a digit pattern",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "abc123def",
		},
		Expected: args.Map{
			params.patternResult: "\\d+",
			params.isDefined:     true,
			params.isApplicable:  true,
			params.isMatch:       true,
			params.isFailedMatch: false,
		},
	},
	{
		Title: "New.Lazy with no-match input returns false",
		When:  "given input that does not match",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.patternResult: "^\\d+$",
			params.isDefined:     true,
			params.isApplicable:  true,
			params.isMatch:       false,
			params.isFailedMatch: true,
		},
	},
	{
		Title: "New.Lazy with invalid regex pattern has error",
		When:  "given an invalid regex pattern",
		Input: args.Map{
			params.pattern:      "[invalid",
			params.compareInput: "anything",
		},
		Expected: args.Map{
			params.patternResult: "[invalid",
			params.isDefined:     true,
			params.isApplicable:  false,
			params.isMatch:       false,
			params.isFailedMatch: true,
		},
	},
	{
		Title: "New.Lazy with empty pattern matches everything",
		When:  "given an empty pattern",
		Input: args.Map{
			params.pattern:      "",
			params.compareInput: "anything",
		},
		Expected: args.Map{
			params.patternResult: "",
			params.isDefined:     false,
			params.isApplicable:  false,
			params.isMatch:       false,
			params.isFailedMatch: true,
		},
	},
}

var lazyRegexLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "New.LazyLock with word pattern is thread-safe",
		When:  "given a word pattern via LazyLock",
		Input: args.Map{
			params.pattern:      "world",
			params.compareInput: "hello world",
		},
		Expected: args.Map{
			params.patternResult: "world",
			params.isDefined:     true,
			params.isApplicable:  true,
			params.isMatch:       true,
			params.isFailedMatch: false,
		},
	},
	{
		Title: "New.LazyLock with email pattern matches email",
		When:  "given an email-like pattern",
		Input: args.Map{
			params.pattern:      `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			params.compareInput: "user@example.com",
		},
		Expected: args.Map{
			params.patternResult: `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
			params.isDefined:     true,
			params.isApplicable:  true,
			params.isMatch:       true,
			params.isFailedMatch: false,
		},
	},
}

// ==========================================================================
// PatternMatch — named test cases
// ==========================================================================

var lazyRegexIsMatchFullDigitTestCase = coretestcases.MapGherkins{
	Title: "IsMatch returns true for full string digit match",
	When:  "given full string digit pattern",
	Input: args.Map{
		params.pattern:      "^\\d+$",
		params.compareInput: "12345",
	},
	Expected: args.Map{
		params.isMatch: true,
	},
}

var lazyRegexIsMatchPartialMismatchTestCase = coretestcases.MapGherkins{
	Title: "IsMatch returns false for partial digit mismatch",
	When:  "given full string digit pattern with letters",
	Input: args.Map{
		params.pattern:      "^\\d+$",
		params.compareInput: "123abc",
	},
	Expected: args.Map{
		params.isMatch: false,
	},
}

var lazyRegexIsFailedMatchTestCase = coretestcases.MapGherkins{
	Title: "IsFailedMatch is inverse of IsMatch",
	When:  "given matching input to IsFailedMatch",
	Input: args.Map{
		params.pattern:      "^hello$",
		params.compareInput: "hello",
	},
	Expected: args.Map{
		params.isFailedMatch: false,
	},
}

var lazyRegexFirstMatchLineFoundTestCase = coretestcases.MapGherkins{
	Title: "FirstMatchLine returns first submatch",
	When:  "given a pattern with capture group",
	Input: args.Map{
		params.pattern:      "(\\d+)",
		params.compareInput: "abc 123 def 456",
	},
	Expected: args.Map{
		params.firstMatch:     "123",
		params.isInvalidMatch: false,
	},
}

var lazyRegexFirstMatchLineNotFoundTestCase = coretestcases.MapGherkins{
	Title: "FirstMatchLine returns empty on no match",
	When:  "given a pattern that does not match",
	Input: args.Map{
		params.pattern:      "(\\d+)",
		params.compareInput: "no digits here",
	},
	Expected: args.Map{
		params.firstMatch:     "",
		params.isInvalidMatch: true,
	},
}
