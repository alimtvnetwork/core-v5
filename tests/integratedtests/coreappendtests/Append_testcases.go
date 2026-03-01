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
		ExpectedInput: []string{
			"3",
			"start",
			"end",
		},
	},
	{
		Title: "PrependAppend with nil prepend skips it",
		ArrangeInput: args.Map{
			"when":   "given nil prepend",
			"append": "end",
		},
		ExpectedInput: []string{
			"2",
			"middle",
			"end",
		},
	},
}
