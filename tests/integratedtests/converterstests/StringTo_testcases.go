package converterstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var stringToIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "StringTo.Integer parses valid integer",
		ArrangeInput: args.Map{
			"when":  "given valid integer string",
			"input": "42",
		},
		ExpectedInput: []string{
			"42",
			"false",
		},
	},
	{
		Title: "StringTo.Integer fails on non-numeric string",
		ArrangeInput: args.Map{
			"when":  "given non-numeric string",
			"input": "abc",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
	{
		Title: "StringTo.Integer parses negative integer",
		ArrangeInput: args.Map{
			"when":  "given negative integer string",
			"input": "-5",
		},
		ExpectedInput: []string{
			"-5",
			"false",
		},
	},
}

var bytesToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "BytesTo.String converts bytes to string",
		ArrangeInput: args.Map{
			"when":  "given valid byte slice",
			"input": "hello",
		},
		ExpectedInput: []string{
			"hello",
		},
	},
	{
		Title: "BytesTo.String returns empty for empty bytes",
		ArrangeInput: args.Map{
			"when":  "given empty byte slice",
			"input": "",
		},
		ExpectedInput: []string{
			"",
		},
	},
}
