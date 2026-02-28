package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var identifierTestCases = []coretestcases.CaseV1{
	{
		Title: "NewIdentifier sets Id correctly",
		ArrangeInput: args.Map{
			"when": "given id 'test-123'",
			"id":   "test-123",
		},
		ExpectedInput: []string{
			"test-123",
			"false",
			"false",
		},
	},
	{
		Title: "NewIdentifier with empty id is empty",
		ArrangeInput: args.Map{
			"when": "given empty id",
			"id":   "",
		},
		ExpectedInput: []string{
			"",
			"true",
			"true",
		},
	},
}
