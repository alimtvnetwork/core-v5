package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var lineNumberHasLineNumberTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 5 HasLineNumber returns true",
		ArrangeInput: args.Map{
			"lineNumber": 5,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 0 HasLineNumber returns true",
		ArrangeInput: args.Map{
			"lineNumber": 0,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 HasLineNumber returns false",
		ArrangeInput: args.Map{
			"lineNumber": -1,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "LineNumber -2 HasLineNumber returns false",
		ArrangeInput: args.Map{
			"lineNumber": -2,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var lineNumberIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 3 IsMatch 3 returns true",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      3,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 3 IsMatch 5 returns false",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "LineNumber 3 IsMatch -1 returns true (skip check)",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      -1,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 IsMatch 5 returns true (skip check)",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 IsMatch -1 returns true",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      -1,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 0 IsMatch 0 returns true",
		ArrangeInput: args.Map{
			"lineNumber": 0,
			"input":      0,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var lineNumberVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 2 VerifyError 2 returns nil",
		ArrangeInput: args.Map{
			"lineNumber": 2,
			"input":      2,
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "LineNumber 2 VerifyError 5 returns error",
		ArrangeInput: args.Map{
			"lineNumber": 2,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "LineNumber -1 VerifyError 5 returns nil (skip)",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}
