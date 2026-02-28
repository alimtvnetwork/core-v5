package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var instanceStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Instance String formats name=value",
		ArrangeInput: args.Map{
			"when":  "given name and value",
			"name":  "host",
			"value": "localhost",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Instance with integer value formats correctly",
		ArrangeInput: args.Map{
			"when":  "given name and integer value",
			"name":  "port",
			"value": 8080,
		},
		ExpectedInput: []string{
			"true",
		},
	},
}
