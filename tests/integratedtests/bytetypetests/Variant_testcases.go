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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "5",     // value
			Second: "false", // isZero
			Third:  "false", // isEmpty
			Fourth: "true",  // isNonZero
		},
	},
	{
		Title: "New creates Zero Variant",
		ArrangeInput: args.Map{
			"when":  "given byte value 0",
			"input": 0,
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "0",    // value
			Second: "true", // isZero
			Third:  "true", // isEmpty
			Fourth: "false", // isNonZero
		},
	},
}
