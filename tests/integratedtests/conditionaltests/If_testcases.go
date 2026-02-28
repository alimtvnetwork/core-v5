package conditionaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var ifStringTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue string",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: []string{
			"yes",
		},
	},
	{
		Title: "If false returns falseValue string",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: []string{
			"no",
		},
	},
}

var ifIntTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue int",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: []string{
			"10",
		},
	},
	{
		Title: "If false returns falseValue int",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: []string{
			"20",
		},
	},
}

var nilDefTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDef with nil pointer returns default",
		ArrangeInput: args.Map{
			"when":   "given nil pointer",
			"isNil":  true,
			"defVal": "default",
		},
		ExpectedInput: []string{
			"default",
		},
	},
	{
		Title: "NilDef with non-nil pointer returns value",
		ArrangeInput: args.Map{
			"when":   "given non-nil pointer",
			"isNil":  false,
			"value":  "actual",
			"defVal": "default",
		},
		ExpectedInput: []string{
			"actual",
		},
	},
}
