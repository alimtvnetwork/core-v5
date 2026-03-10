package corecomparatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var compareStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal compare has correct name and symbol",
		ArrangeInput: args.Map{
			"when":  "given Equal compare",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"name":      "Equal",
			"symbol":    "=",
			"shortName": "eq",
			"isEqual":   true,
			"isValid":   true,
		},
	},
	{
		Title: "LeftGreater compare has correct name and symbol",
		ArrangeInput: args.Map{
			"when":  "given LeftGreater compare",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"name":      "LeftGreater",
			"symbol":    ">",
			"shortName": "gt",
			"isEqual":   false,
			"isValid":   true,
		},
	},
	{
		Title: "Inconclusive compare is invalid",
		ArrangeInput: args.Map{
			"when":  "given Inconclusive compare",
			"value": 6,
		},
		ExpectedInput: args.Map{
			"name":      "Inconclusive",
			"symbol":    "?!",
			"shortName": "i",
			"isEqual":   false,
			"isValid":   false,
		},
	},
}
