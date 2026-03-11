package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var createMustTestCases = []coretestcases.MapGherkins{
	{
		Title: "CreateMust with valid digit pattern returns compiled regex",
		When:  "given a valid digit pattern",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "abc123",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
	{
		Title: "CreateMust with valid word pattern returns compiled regex",
		When:  "given a valid word pattern",
		Input: args.Map{
			params.pattern:      "\\w+",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
	{
		Title: "CreateMust with anchored pattern matches correctly",
		When:  "given an anchored pattern with non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     false,
		},
	},
}

var createMustLockIfTestCases = []coretestcases.MapGherkins{
	{
		Title: "CreateMustLockIf with lock true compiles valid pattern",
		When:  "given valid pattern with lock true",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "99",
			params.isLock:       true,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
	{
		Title: "CreateMustLockIf with lock false compiles valid pattern",
		When:  "given valid pattern with lock false",
		Input: args.Map{
			params.pattern:      "[a-z]+",
			params.compareInput: "hello",
			params.isLock:       false,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
}

var createLockIfTestCases = []coretestcases.MapGherkins{
	{
		Title: "CreateLockIf with lock true compiles valid pattern",
		When:  "given valid pattern with lock true",
		Input: args.Map{
			params.pattern: "\\d+",
			params.isLock:  true,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "CreateLockIf with lock false compiles valid pattern",
		When:  "given valid pattern with lock false",
		Input: args.Map{
			params.pattern: "[a-z]+",
			params.isLock:  false,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "CreateLockIf with lock true returns error for invalid pattern",
		When:  "given invalid pattern with lock true",
		Input: args.Map{
			params.pattern: "[bad",
			params.isLock:  true,
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
	{
		Title: "CreateLockIf with lock false returns error for invalid pattern",
		When:  "given invalid pattern with lock false",
		Input: args.Map{
			params.pattern: "(unclosed",
			params.isLock:  false,
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
}

var createApplicableLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "CreateApplicableLock with valid pattern is applicable",
		When:  "given a valid pattern",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil:  true,
			params.hasError:     false,
			params.isApplicable: true,
		},
	},
	{
		Title: "CreateApplicableLock with invalid pattern is not applicable",
		When:  "given an invalid pattern",
		Input: args.Map{
			params.pattern: "[bad",
		},
		Expected: args.Map{
			params.regexNotNil:  false,
			params.hasError:     true,
			params.isApplicable: false,
		},
	},
	{
		Title: "CreateApplicableLock with empty pattern is applicable",
		When:  "given an empty pattern",
		Input: args.Map{
			params.pattern: "",
		},
		Expected: args.Map{
			params.regexNotNil:  true,
			params.hasError:     false,
			params.isApplicable: true,
		},
	},
}

var newMustLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "NewMustLock with valid pattern returns compiled regex",
		When:  "given a valid digit pattern",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
	{
		Title: "NewMustLock with word boundary pattern matches",
		When:  "given a word boundary pattern",
		Input: args.Map{
			params.pattern:      "\\bhello\\b",
			params.compareInput: "hello world",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     true,
		},
	},
	{
		Title: "NewMustLock with anchored pattern rejects mismatch",
		When:  "given an anchored pattern with non-matching input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.isMatch:     false,
		},
	},
}

var matchUsingFuncErrorLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "MatchUsingFuncErrorLock returns nil on match",
		When:  "given matching input with MatchString func",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isNoError: true,
		},
	},
	{
		Title: "MatchUsingFuncErrorLock returns error on mismatch",
		When:  "given non-matching input with MatchString func",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
	{
		Title: "MatchUsingFuncErrorLock returns error for invalid pattern",
		When:  "given invalid pattern with MatchString func",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
}

var matchUsingCustomizeErrorFuncLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "CustomizeErrorFunc returns nil on match with nil customizer",
		When:  "given matching input with nil customizer",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
			params.customizer:   "nil",
		},
		Expected: args.Map{
			params.isNoError:    true,
			params.isCustomError: false,
		},
	},
	{
		Title: "CustomizeErrorFunc returns default error on mismatch with nil customizer",
		When:  "given non-matching input with nil customizer",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
			params.customizer:   "nil",
		},
		Expected: args.Map{
			params.isNoError:    false,
			params.isCustomError: false,
		},
	},
	{
		Title: "CustomizeErrorFunc returns custom error on mismatch",
		When:  "given non-matching input with custom error func",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
			params.customizer:   "custom",
		},
		Expected: args.Map{
			params.isNoError:    false,
			params.isCustomError: true,
		},
	},
	{
		Title: "CustomizeErrorFunc returns nil on match with custom error func",
		When:  "given matching input with custom error func",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
			params.customizer:   "custom",
		},
		Expected: args.Map{
			params.isNoError:    true,
			params.isCustomError: false,
		},
	},
}
