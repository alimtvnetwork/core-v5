package mutexbykeytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var extGetSameKeyTwiceTestCases = []coretestcases.CaseV1{
	{
		Title: "Get returns same mutex instance -- same key requested twice",
		ArrangeInput: args.Map{
			"when": "given same key requested twice",
			"key":  "test-key-ext-same",
		},
		ExpectedInput: args.Map{
			"isSame": true,
		},
	},
}
