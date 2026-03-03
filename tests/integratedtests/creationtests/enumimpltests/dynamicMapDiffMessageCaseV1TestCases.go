package enumimpltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var dynamicMapDiffMessageCaseV1TestCases = []coretestcases.CaseV1{
	{
		Title: "Dynamic map diff string compiled must be same",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"exist":                        1,
				"not-exist-in-right":           3,
				"exist-in-left-right-diff-val": 5,
			},
			"right": map[string]any{
				"exist":                        1,
				"not-exist-in-left":            2,
				"exist-in-left-right-diff-val": 6,
			},
		},
		ExpectedInput: []string{
			"Dynamic map diff string compiled must be same",
			"",
			"Difference Between Map:",
			"",
			"{{",
			"",
			"  \"not-exist-in-left\":2,",
			"  \"not-exist-in-right\":3,",
			"  \"exist-in-left-right-diff-val\":5",
			"",
			"}}",
		},
	},
}
