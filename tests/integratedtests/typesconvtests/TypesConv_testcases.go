package typesconvtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var stringToBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "StringToBool returns true for 'true'",
		ArrangeInput: args.Map{
			"when":  "given 'true'",
			"input": "true",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "StringToBool returns true for 'yes'",
		ArrangeInput: args.Map{
			"when":  "given 'yes'",
			"input": "yes",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "StringToBool returns false for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "StringToBool returns false for 'no'",
		ArrangeInput: args.Map{
			"when":  "given 'no'",
			"input": "no",
		},
		ExpectedInput: []string{"false"},
	},
}

var intPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "IntPtrToSimple returns value for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 42,
		},
		ExpectedInput: []string{"42"},
	},
	{
		Title: "IntPtrToSimple returns 0 for nil",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"0"},
	},
}
