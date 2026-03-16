package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isMatchLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.IsMatchLock returns true -- matching digit pattern",
		When:  "given digit pattern and numeric input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
		},
		Expected: args.Map{
			params.isMatch: true,
		},
	},
	{
		Title: "LazyRegex.IsMatchLock returns false -- non-matching pattern",
		When:  "given digit-only pattern and alpha input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
	{
		Title: "LazyRegex.IsMatchLock returns false -- invalid pattern",
		When:  "given invalid regex pattern",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
	{
		Title: "LazyRegex.IsMatchLock returns true -- email-like pattern",
		When:  "given email-like pattern",
		Input: args.Map{
			params.pattern:      "^[a-zA-Z0-9]+@[a-zA-Z]+\\.[a-zA-Z]+$",
			params.compareInput: "user@example.com",
		},
		Expected: args.Map{
			params.isMatch: true,
		},
	},
	{
		Title: "LazyRegex.IsMatchLock returns false -- empty input with required pattern",
		When:  "given required pattern and empty input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
}

var isMatchFailedTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.IsFailedMatch returns false -- pattern matches",
		When:  "given matching pattern",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "42",
		},
		Expected: args.Map{
			params.isFailed: false,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatch returns true -- pattern does not match",
		When:  "given non-matching pattern",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isFailed: true,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatch returns true -- invalid pattern",
		When:  "given invalid pattern",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isFailed: true,
		},
	},
}

var isMatchLockLazyIsMatchTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.IsMatch returns true -- matching pattern",
		When:  "given valid pattern with matching input",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isMatch: true,
		},
	},
	{
		Title: "LazyRegex.IsMatch returns false -- non-matching input",
		When:  "given valid pattern with non-matching input",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "world",
		},
		Expected: args.Map{
			params.isMatch: false,
		},
	},
}

var isMatchLockCompileTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.Compile returns no error -- valid pattern",
		When:  "given valid pattern",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil:  true,
			params.hasError:     false,
			params.isApplicable: true,
		},
	},
}

var isMatchLockIsFailedMatchTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.IsFailedMatch returns false -- matching input",
		When:  "given matching input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
		},
		Expected: args.Map{
			params.isFailedMatch: false,
		},
	},
	{
		Title: "LazyRegex.IsFailedMatch returns true -- non-matching input",
		When:  "given non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isFailedMatch: true,
		},
	},
}

var isMatchLockPatternStringTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.Pattern returns original pattern -- valid pattern",
		When:  "given a pattern",
		Input: args.Map{
			params.pattern: "^test\\d+$",
		},
		Expected: args.Map{
			params.patternResult: "^test\\d+$",
		},
	},
}

var isMatchLockMatchErrorTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.MatchError returns nil -- matching input",
		When:  "given matching input",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isNoError: true,
		},
	},
	{
		Title: "LazyRegex.MatchError returns error -- non-matching input",
		When:  "given non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
}
