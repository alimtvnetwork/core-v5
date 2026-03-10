package errcoretests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var errTypeCombineTestCases = []coretestcases.CaseV1{
	{
		Title: "Combine returns formatted string -- message 'some 2' and ref 'alim-1'",
		ArrangeInput: args.Map{
			"when":    "given message and reference",
			"message": "some 2",
			"ref":     "alim-1",
		},
		ExpectedInput: ".*some 2.*alim-1.*",
	},
	{
		Title: "Combine returns string with reference -- empty message, ref 'alim-2 no msg'",
		ArrangeInput: args.Map{
			"when":    "given empty message with reference",
			"message": "",
			"ref":     "alim-2 no msg",
		},
		ExpectedInput: ".*alim-2 no msg.*",
	},
	{
		Title: "Combine returns type name only -- empty message and empty ref",
		ArrangeInput: args.Map{
			"when":    "given both empty",
			"message": "",
			"ref":     "",
		},
		ExpectedInput: ".*BytesAreNilOrEmpty.*",
	},
}

var errMergeTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors returns nil -- both errors nil",
		ArrangeInput: args.Map{
			"when":     "given both nil errors",
			"hasError": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "MergeErrors returns error -- one non-nil error",
		ArrangeInput: args.Map{
			"when":     "given one real error",
			"hasError": true,
		},
		ExpectedInput: "false",
	},
}

var errTypeErrorNoRefsTestCases = []coretestcases.CaseV1{
	{
		Title: "ErrorNoRefs returns non-nil error -- message 'something broke'",
		ArrangeInput: args.Map{
			"when":    "given a message",
			"message": "something broke",
		},
		ExpectedInput: "true",
	},
	{
		Title: "ErrorNoRefs returns non-nil error -- empty message",
		ArrangeInput: args.Map{
			"when":    "given empty message",
			"message": "",
		},
		ExpectedInput: "true",
	},
}

var errTypeErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Error returns string with both -- message 'parsing failed' and ref 'line-42'",
		ArrangeInput: args.Map{
			"when":    "given message and ref",
			"message": "parsing failed",
			"ref":     "line-42",
		},
		ExpectedInput: ".*parsing failed.*line-42.*",
	},
	{
		Title: "Error returns string with message -- message 'some error', empty ref",
		ArrangeInput: args.Map{
			"when":    "given message only",
			"message": "some error",
			"ref":     "",
		},
		ExpectedInput: ".*some error.*",
	},
}
