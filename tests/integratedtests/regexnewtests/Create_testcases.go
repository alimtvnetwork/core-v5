package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var createTestCases = []coretestcases.MapGherkins{
	{
		Title: "Create returns compiled regex -- valid digit pattern",
		When:  "given a valid digit pattern",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "Create returns compiled regex -- valid word boundary pattern",
		When:  "given a word boundary pattern",
		Input: args.Map{
			params.pattern: "\\bhello\\b",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "Create returns error -- invalid bracket pattern",
		When:  "given an invalid pattern",
		Input: args.Map{
			params.pattern: "[invalid",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
	{
		Title: "Create returns compiled regex -- empty pattern",
		When:  "given an empty pattern",
		Input: args.Map{
			params.pattern: "",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
}

var createIsMatchLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "IsMatchLock returns true -- matching digit pattern",
		When:  "given matching digit input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "abc123",
		},
		Expected: args.Map{
			params.isMatch: true,
		},
	},
	{
		Title: "IsMatchLock returns false -- non-matching pattern",
		When:  "given non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
	{
		Title: "IsMatchLock returns false -- invalid regex",
		When:  "given invalid regex pattern",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "anything",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
	{
		Title: "IsMatchLock returns true -- exact full match",
		When:  "given exact match pattern",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isMatch: true,
		},
	},
}

var createIsMatchFailedTestCases = []coretestcases.MapGherkins{
	{
		Title: "IsMatchFailed returns false -- pattern matches",
		When:  "given matching input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
		},
		Expected: args.Map{
			params.isFailed: false,
		},
	},
	{
		Title: "IsMatchFailed returns true -- pattern does not match",
		When:  "given non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isFailed: true,
		},
	},
	{
		Title: "IsMatchFailed returns true -- invalid regex",
		When:  "given invalid regex",
		Input: args.Map{
			params.pattern:      "[broken",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isFailed: true,
		},
	},
}

var matchErrorMatchTestCase = coretestcases.MapGherkins{
	Title: "MatchError returns no error -- matching input",
	When:  "given matching input to MatchError",
	Input: args.Map{
		params.pattern:      "^hello$",
		params.compareInput: "hello",
	},
	Expected: args.Map{
		params.isNoError: true,
	},
}

var matchErrorMismatchTestCase = coretestcases.MapGherkins{
	Title: "MatchError returns error -- non-matching input",
	When:  "given non-matching input to MatchError",
	Input: args.Map{
		params.pattern:      "^\\d+$",
		params.compareInput: "abc",
	},
	Expected: args.Map{
		params.isNoError: false,
	},
}

var matchErrorLockMatchTestCase = coretestcases.MapGherkins{
	Title: "MatchErrorLock returns no error -- matching input",
	When:  "given matching input to MatchErrorLock",
	Input: args.Map{
		params.pattern:      "world",
		params.compareInput: "hello world",
	},
	Expected: args.Map{
		params.isNoError: true,
	},
}

var matchErrorLockMismatchTestCase = coretestcases.MapGherkins{
	Title: "MatchErrorLock returns error -- non-matching input",
	When:  "given non-matching input to MatchErrorLock",
	Input: args.Map{
		params.pattern:      "^xyz$",
		params.compareInput: "abc",
	},
	Expected: args.Map{
		params.isNoError: false,
	},
}
