package defaulterrtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var defaultErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Marshalling error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking Marshalling error",
			"error": "Marshalling",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "UnMarshalling error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnMarshalling error",
			"error": "UnMarshalling",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
	{
		Title: "OutOfRange error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking OutOfRange error",
			"error": "OutOfRange",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}
