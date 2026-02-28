package keymktests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var keyCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "Default key compiles with hyphen joiner",
		ArrangeInput: args.Map{
			"when":   "given main and chain items",
			"main":   "root",
			"chains": []string{"sub", "item"},
		},
		ExpectedInput: []string{
			"root-sub-item",
		},
	},
	{
		Title: "Default key compiles with main only",
		ArrangeInput: args.Map{
			"when":   "given main only",
			"main":   "solo",
			"chains": []string{},
		},
		ExpectedInput: []string{
			"solo",
		},
	},
}
