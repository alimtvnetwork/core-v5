package enumimpltests

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var dynamicMapDiffCaseV1TestCases = []coretestcases.CaseV1{
	{
		Title: "dynamic map must yield diff properly.",
		ArrangeInput: args.Map{
			"left": enumimpl.DynamicMap{
				"exist":                        1,
				"not-exist-in-right":           3,
				"exist-in-left-right-diff-val": 5,
			},
			"right": enumimpl.DynamicMap{
				"exist":                        1,
				"not-exist-in-left":            2,
				"exist-in-left-right-diff-val": 6,
			},
		},
		ExpectedInput: []string{
			"exist-in-left-right-diff-val : 5",
			"not-exist-in-left : 2",
			"not-exist-in-right : 3",
		},
	},
}
