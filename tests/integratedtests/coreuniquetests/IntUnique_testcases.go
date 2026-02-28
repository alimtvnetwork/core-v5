package coreuniquetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var intUniqueGetTestCases = []coretestcases.CaseV1{
	{
		Title: "Get removes duplicates",
		ArrangeInput: args.Map{
			"when":  "given slice with duplicates",
			"input": []int{1, 2, 2, 3, 3, 3},
		},
		ExpectedInput: []string{
			"3",
		},
	},
	{
		Title: "Get returns same for already unique",
		ArrangeInput: args.Map{
			"when":  "given slice without duplicates",
			"input": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"3",
		},
	},
	{
		Title: "Get handles nil",
		ArrangeInput: args.Map{
			"when":  "given nil slice",
			"isNil": true,
		},
		ExpectedInput: []string{
			"true",
		},
	},
}
