package coreutilstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var isEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty returns true for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsEmpty returns false for non-empty string",
		ArrangeInput: args.Map{
			"when":  "given non-empty string",
			"input": "hello",
		},
		ExpectedInput: []string{"false"},
	},
}

var isBlankTestCases = []coretestcases.CaseV1{
	{
		Title: "IsBlank returns true for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsBlank returns true for whitespace only",
		ArrangeInput: args.Map{
			"when":  "given whitespace string",
			"input": "   ",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsBlank returns false for non-blank string",
		ArrangeInput: args.Map{
			"when":  "given non-blank string",
			"input": "hello",
		},
		ExpectedInput: []string{"false"},
	},
}
