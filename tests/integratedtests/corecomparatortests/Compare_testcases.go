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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "Equal", // name
			Second: "=",     // symbol
			Third:  "eq",    // shortName
			Fourth: "true",  // isEqual
			Fifth:  "true",  // isValid
		},
	},
	{
		Title: "LeftGreater compare has correct name and symbol",
		ArrangeInput: args.Map{
			"when":  "given LeftGreater compare",
			"value": 1,
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "LeftGreater", // name
			Second: ">",           // symbol
			Third:  "gt",          // shortName
			Fourth: "false",       // isEqual
			Fifth:  "true",        // isValid
		},
	},
	{
		Title: "Inconclusive compare is invalid",
		ArrangeInput: args.Map{
			"when":  "given Inconclusive compare",
			"value": 6,
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "Inconclusive", // name
			Second: "?!",           // symbol
			Third:  "i",            // shortName
			Fourth: "false",        // isEqual
			Fifth:  "false",        // isValid
		},
	},
}
