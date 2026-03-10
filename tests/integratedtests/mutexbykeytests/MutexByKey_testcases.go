package mutexbykeytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var getAndDeleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Get returns non-nil mutex -- new key 'test-key-1'",
		ArrangeInput: args.Map{
			"when": "given a new key",
			"key":  "test-key-1",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Get returns same mutex -- same key 'test-key-same' requested twice",
		ArrangeInput: args.Map{
			"when": "given same key twice",
			"key":  "test-key-same",
		},
		ExpectedInput: "true",
	},
}

var deleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Delete returns true -- existing key 'test-key-del'",
		ArrangeInput: args.Map{
			"when": "given existing key to delete",
			"key":  "test-key-del",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Delete returns false -- non-existing key 'test-key-nonexistent'",
		ArrangeInput: args.Map{
			"when": "given non-existing key to delete",
			"key":  "test-key-nonexistent",
		},
		ExpectedInput: "false",
	},
}
