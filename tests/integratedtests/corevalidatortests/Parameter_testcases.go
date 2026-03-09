package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var parameterIsIgnoreCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Parameter IsCaseSensitive=true IsIgnoreCase returns false",
		ArrangeInput: args.Map{
			"isCaseSensitive": true,
		},
		ExpectedInput: args.Map{
			"isIgnoreCase": false,
		},
	},
	{
		Title: "Parameter IsCaseSensitive=false IsIgnoreCase returns true",
		ArrangeInput: args.Map{
			"isCaseSensitive": false,
		},
		ExpectedInput: args.Map{
			"isIgnoreCase": true,
		},
	},
}

var parameterDefaultValuesTestCase = coretestcases.CaseV1{
	Title: "Parameter zero value has expected defaults",
	ExpectedInput: args.Map{
		"caseIndex":                  0,
		"header":                     "",
		"isSkipCompareOnActualEmpty": false,
		"isAttachUserInputs":         false,
		"isCaseSensitive":            false,
	},
}
