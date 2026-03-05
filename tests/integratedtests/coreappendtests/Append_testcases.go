package coreappendtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var prependAppendTestCases = []coretestcases.CaseV1{
	{
		Title: "PrependAppend skips nil items",
		ArrangeInput: args.Map{
			"when":    "given prepend, append and middle items with nil",
			"prepend": "start",
			"append":  "end",
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "3",     // totalCount
			Second: "start", // firstItem
			Third:  "end",   // lastItem
		},
	},
	{
		Title: "PrependAppend with nil prepend skips it",
		ArrangeInput: args.Map{
			"when":   "given nil prepend",
			"append": "end",
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",      // totalCount
			Second: "middle", // firstItem
			Third:  "end",    // lastItem
		},
	},
}
