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
		Title: "New.Lazy returns matching result -- simple word pattern",
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
		Title: "New.Lazy returns matching result -- digit pattern",
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
		Title: "New.Lazy returns no match -- non-matching input",
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
		Title: "New.Lazy returns error -- invalid regex pattern",
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
		Title: "New.Lazy returns undefined -- empty pattern",
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
		Title: "New.LazyLock returns matching result -- word pattern thread-safe",
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
		Title: "New.LazyLock returns matching result -- email pattern",
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
	Title: "LazyRegex.IsMatch returns true -- full string digit match",
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
	Title: "LazyRegex.IsMatch returns false -- partial digit mismatch",
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
	Title: "LazyRegex.IsFailedMatch returns false -- matching input (inverse of IsMatch)",
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
	Title: "LazyRegex.FirstMatchLine returns first submatch -- capture group match",
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
	Title: "LazyRegex.FirstMatchLine returns empty -- no match found",
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
