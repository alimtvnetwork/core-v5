package bytetypetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var newVariantTestCases = []coretestcases.CaseV1{
	{
		Title: "New creates Variant with correct value",
		ArrangeInput: args.Map{
			"when":  "given byte value 5",
			"input": 5,
		},
		ExpectedInput: []string{
			"5",
			"false",
			"false",
			"true",
		},
	},
	{
		Title: "New creates Zero Variant",
		ArrangeInput: args.Map{
			"when":  "given byte value 0",
			"input": 0,
		},
		ExpectedInput: []string{
			"0",
			"true",
			"true",
			"false",
		},
	},
}
