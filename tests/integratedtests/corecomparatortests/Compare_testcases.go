package corecomparatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var compareStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal compare has correct name and symbol",
		ArrangeInput: args.Map{
			"when":  "given Equal compare",
			"value": 0,
		},
		ExpectedInput: []string{
			"Equal",
			"=",
			"eq",
			"true",
			"true",
		},
	},
	{
		Title: "LeftGreater compare has correct name and symbol",
		ArrangeInput: args.Map{
			"when":  "given LeftGreater compare",
			"value": 1,
		},
		ExpectedInput: []string{
			"LeftGreater",
			">",
			"gt",
			"false",
			"true",
		},
	},
	{
		Title: "Inconclusive compare is invalid",
		ArrangeInput: args.Map{
			"when":  "given Inconclusive compare",
			"value": 6,
		},
		ExpectedInput: []string{
			"Inconclusive",
			"?!",
			"i",
			"false",
			"false",
		},
	},
}
