package errcorretests

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
		ExpectedInput: []string{
			".*some 2.*alim-1.*",
		},
	},
	{
		Title: "Combine with empty message keeps reference",
		ArrangeInput: args.Map{
			"when":    "given empty message with reference",
			"message": "",
			"ref":     "alim-2 no msg",
		},
		ExpectedInput: []string{
			".*alim-2 no msg.*",
		},
	},
	{
		Title: "Combine with both empty returns type name only",
		ArrangeInput: args.Map{
			"when":    "given both empty",
			"message": "",
			"ref":     "",
		},
		ExpectedInput: []string{
			".*BytesAreNilOrEmpty.*",
		},
	},
}

var errMergeTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors with nil errors returns nil",
		ArrangeInput: args.Map{
			"when":     "given both nil errors",
			"hasError": false,
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "MergeErrors with one non-nil error returns error",
		ArrangeInput: args.Map{
			"when":     "given one real error",
			"hasError": true,
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
