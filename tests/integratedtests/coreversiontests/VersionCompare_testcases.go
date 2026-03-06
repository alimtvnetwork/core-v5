package coreversiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var versionCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal versions return Equal",
		ArrangeInput: args.Map{
			"when":  "given equal versions",
			"left":  "v0.0.1",
			"right": "v0.0.1",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
	{
		Title: "Left major greater returns LeftGreater",
		ArrangeInput: args.Map{
			"when":  "given left major version greater",
			"left":  "v3.0",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftGreater",
		},
	},
	{
		Title: "Left minor less returns LeftLess",
		ArrangeInput: args.Map{
			"when":  "given left minor version less",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftLess",
		},
	},
	{
		Title: "Same major with zero-padded returns Equal",
		ArrangeInput: args.Map{
			"when":  "given v4 vs v4.0",
			"left":  "v4",
			"right": "v4.0",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
}
