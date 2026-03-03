package errcoretests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var errTypeCombineTestCases = []coretestcases.CaseV1{
	{
		Title: "Combine with message and reference",
		ArrangeInput: args.Map{
			"when":    "given message and reference",
			"message": "some 2",
			"ref":     "alim-1",
		},
		ExpectedInput: ".*some 2.*alim-1.*",
	},
	{
		Title: "Combine with empty message keeps reference",
		ArrangeInput: args.Map{
			"when":    "given empty message with reference",
			"message": "",
			"ref":     "alim-2 no msg",
		},
		ExpectedInput: ".*alim-2 no msg.*",
	},
	{
		Title: "Combine with both empty returns type name only",
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
		Title: "MergeErrors with nil errors returns nil",
		ArrangeInput: args.Map{
			"when":     "given both nil errors",
			"hasError": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "MergeErrors with one non-nil error returns error",
		ArrangeInput: args.Map{
			"when":     "given one real error",
			"hasError": true,
		},
		ExpectedInput: "false",
	},
}

var errTypeErrorNoRefsTestCases = []coretestcases.CaseV1{
	{
		Title: "ErrorNoRefs with message returns non-nil error",
		ArrangeInput: args.Map{
			"when":    "given a message",
			"message": "something broke",
		},
		ExpectedInput: "true",
	},
	{
		Title: "ErrorNoRefs with empty message returns non-nil error",
		ArrangeInput: args.Map{
			"when":    "given empty message",
			"message": "",
		},
		ExpectedInput: "true",
	},
}

var errTypeErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Error with message and ref includes both",
		ArrangeInput: args.Map{
			"when":    "given message and ref",
			"message": "parsing failed",
			"ref":     "line-42",
		},
		ExpectedInput: ".*parsing failed.*line-42.*",
	},
	{
		Title: "Error with empty ref includes message",
		ArrangeInput: args.Map{
			"when":    "given message only",
			"message": "some error",
			"ref":     "",
		},
		ExpectedInput: ".*some error.*",
	},
}
