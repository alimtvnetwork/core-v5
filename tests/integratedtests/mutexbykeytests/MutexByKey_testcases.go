package mutexbykeytests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var getAndDeleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Get returns non-nil mutex for key",
		ArrangeInput: args.Map{
			"when": "given a new key",
			"key":  "test-key-1",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Get returns same mutex for same key",
		ArrangeInput: args.Map{
			"when": "given same key twice",
			"key":  "test-key-same",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

var deleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Delete returns true for existing key",
		ArrangeInput: args.Map{
			"when": "given existing key to delete",
			"key":  "test-key-del",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Delete returns false for non-existing key",
		ArrangeInput: args.Map{
			"when": "given non-existing key to delete",
			"key":  "test-key-nonexistent",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
