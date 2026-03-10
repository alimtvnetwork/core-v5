package keymktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var keyLegendGroupIntRangeTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupIntRange returns correct key range -- range 5 to 10",
		ArrangeInput: args.Map{
			"when":    "given range 5 to 10",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"startId": 5,
			"endId":   10,
		},
		ExpectedInput: args.Map{
			"count":    "6",
			"firstKey": "cimux-main-5-stateName",
			"lastKey":  "cimux-main-10-stateName",
		},
	},
}

var keyLegendUserStringWithoutStateTestCases = []coretestcases.CaseV1{
	{
		Title: "UserStringWithoutState returns root-package-group-user -- user 'mynewuser1'",
		ArrangeInput: args.Map{
			"when":    "given user mynewuser1",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"user":    "mynewuser1",
		},
		ExpectedInput: "cimux-main-myg-mynewuser1",
	},
}

var keyLegendUpToStateTestCases = []coretestcases.CaseV1{
	{
		Title: "UpToState returns root-package-group-state-user -- user 'my-user'",
		ArrangeInput: args.Map{
			"when":    "given user my-user",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"user":    "my-user",
		},
		ExpectedInput: "cimux-main-myg-stateName-my-user",
	},
}
