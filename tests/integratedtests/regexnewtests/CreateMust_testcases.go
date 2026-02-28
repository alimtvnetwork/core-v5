package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var createMustTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateMust with valid digit pattern returns compiled regex",
		ArrangeInput: args.Map{
			"when":    "given a valid digit pattern",
			"pattern": "\\d+",
			"input":   "abc123",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "CreateMust with valid word pattern returns compiled regex",
		ArrangeInput: args.Map{
			"when":    "given a valid word pattern",
			"pattern": "\\w+",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "CreateMust with anchored pattern matches correctly",
		ArrangeInput: args.Map{
			"when":    "given an anchored pattern with non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

var createMustLockIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateMustLockIf with lock true compiles valid pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with lock true",
			"pattern": "\\d+",
			"input":   "99",
			"isLock":  true,
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "CreateMustLockIf with lock false compiles valid pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with lock false",
			"pattern": "[a-z]+",
			"input":   "hello",
			"isLock":  false,
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var createLockIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateLockIf with lock true compiles valid pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with lock true",
			"pattern": "\\d+",
			"isLock":  true,
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "CreateLockIf with lock false compiles valid pattern",
		ArrangeInput: args.Map{
			"when":    "given valid pattern with lock false",
			"pattern": "[a-z]+",
			"isLock":  false,
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "CreateLockIf with lock true returns error for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid pattern with lock true",
			"pattern": "[bad",
			"isLock":  true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "CreateLockIf with lock false returns error for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid pattern with lock false",
			"pattern": "(unclosed",
			"isLock":  false,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

var createApplicableLockTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateApplicableLock with valid pattern is applicable",
		ArrangeInput: args.Map{
			"when":    "given a valid pattern",
			"pattern": "\\d+",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
		},
	},
	{
		Title: "CreateApplicableLock with invalid pattern is not applicable",
		ArrangeInput: args.Map{
			"when":    "given an invalid pattern",
			"pattern": "[bad",
		},
		ExpectedInput: []string{
			"false",
			"true",
			"false",
		},
	},
	{
		Title: "CreateApplicableLock with empty pattern is applicable",
		ArrangeInput: args.Map{
			"when":    "given an empty pattern",
			"pattern": "",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
		},
	},
}

var newMustLockTestCases = []coretestcases.CaseV1{
	{
		Title: "NewMustLock with valid pattern returns compiled regex",
		ArrangeInput: args.Map{
			"when":    "given a valid digit pattern",
			"pattern": "\\d+",
			"input":   "123",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "NewMustLock with word boundary pattern matches",
		ArrangeInput: args.Map{
			"when":    "given a word boundary pattern",
			"pattern": "\\bhello\\b",
			"input":   "hello world",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "NewMustLock with anchored pattern rejects mismatch",
		ArrangeInput: args.Map{
			"when":    "given an anchored pattern with non-matching input",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

var matchUsingFuncErrorLockTestCases = []coretestcases.CaseV1{
	{
		Title: "MatchUsingFuncErrorLock returns nil on match",
		ArrangeInput: args.Map{
			"when":    "given matching input with MatchString func",
			"pattern": "^hello$",
			"input":   "hello",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "MatchUsingFuncErrorLock returns error on mismatch",
		ArrangeInput: args.Map{
			"when":    "given non-matching input with MatchString func",
			"pattern": "^\\d+$",
			"input":   "abc",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "MatchUsingFuncErrorLock returns error for invalid pattern",
		ArrangeInput: args.Map{
			"when":    "given invalid pattern with MatchString func",
			"pattern": "[bad",
			"input":   "test",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}

var matchUsingCustomizeErrorFuncLockTestCases = []coretestcases.CaseV1{
	{
		Title: "CustomizeErrorFunc returns nil on match with nil customizer",
		ArrangeInput: args.Map{
			"when":       "given matching input with nil customizer",
			"pattern":    "^hello$",
			"input":      "hello",
			"customizer": "nil",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "CustomizeErrorFunc returns default error on mismatch with nil customizer",
		ArrangeInput: args.Map{
			"when":       "given non-matching input with nil customizer",
			"pattern":    "^\\d+$",
			"input":      "abc",
			"customizer": "nil",
		},
		ExpectedInput: []string{
			"false",
			"false",
		},
	},
	{
		Title: "CustomizeErrorFunc returns custom error on mismatch",
		ArrangeInput: args.Map{
			"when":       "given non-matching input with custom error func",
			"pattern":    "^\\d+$",
			"input":      "abc",
			"customizer": "custom",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "CustomizeErrorFunc returns nil on match with custom error func",
		ArrangeInput: args.Map{
			"when":       "given matching input with custom error func",
			"pattern":    "\\d+",
			"input":      "123",
			"customizer": "custom",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}
