package coreappendtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var prependAppendTestCases = []coretestcases.CaseV1{
	{
		Title: "PrependAppendAnyItemsToStringsSkipOnNil returns 3 items -- prepend and append provided",
		ArrangeInput: args.Map{
			"when":    "given prepend, append and middle items with nil",
			"prepend": "start",
			"append":  "end",
		},
		ExpectedInput: args.Map{
			"totalCount": "3",
			"firstItem":  "start",
			"lastItem":   "end",
		},
	},
	{
		Title: "PrependAppendAnyItemsToStringsSkipOnNil returns 2 items -- nil prepend",
		ArrangeInput: args.Map{
			"when":   "given nil prepend",
			"append": "end",
		},
		ExpectedInput: args.Map{
			"totalCount": "2",
			"firstItem":  "middle",
			"lastItem":   "end",
		},
	},
}
