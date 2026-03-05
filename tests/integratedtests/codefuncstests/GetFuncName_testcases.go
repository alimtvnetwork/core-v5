package codefuncstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncName returns name for a function",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: "true", // containsFuncName
	},
}
