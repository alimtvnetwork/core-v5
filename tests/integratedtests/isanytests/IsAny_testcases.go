package isanytests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var isAnyDefinedNullTestCases = []coretestcases.CaseV1{
	{
		Title: "nil is Null and not Defined",
		ArrangeInput: args.Map{
			"when":  "given nil value",
			"input": nil,
		},
		ExpectedInput: args.Map{
			"isDefined": "false",
			"isNull":    "true",
		},
	},
	{
		Title: "non-nil error is Defined and not Null",
		ArrangeInput: args.Map{
			"when":     "given a non-nil error",
			"input":    "error-marker",
			"useError": true,
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
	{
		Title: "empty string is Defined (not nil)",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
	{
		Title: "integer zero is Defined",
		ArrangeInput: args.Map{
			"when":  "given integer zero",
			"input": 0,
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
}

var isAnyBothTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedBoth(nil, non-nil) returns false",
		ArrangeInput: args.Map{
			"when":   "given nil and non-nil",
			"first":  nil,
			"second": "something",
		},
		ExpectedInput: args.Map{
			"definedBoth": "false",
			"nullBoth":    "false",
		},
	},
	{
		Title: "NullBoth(nil, nil) returns true",
		ArrangeInput: args.Map{
			"when":   "given both nil",
			"first":  nil,
			"second": nil,
		},
		ExpectedInput: args.Map{
			"definedBoth": "false",
			"nullBoth":    "true",
		},
	},
	{
		Title: "DefinedBoth(string, string) returns true",
		ArrangeInput: args.Map{
			"when":   "given both defined strings",
			"first":  "a",
			"second": "b",
		},
		ExpectedInput: args.Map{
			"definedBoth": "true",
			"nullBoth":    "false",
		},
	},
}
