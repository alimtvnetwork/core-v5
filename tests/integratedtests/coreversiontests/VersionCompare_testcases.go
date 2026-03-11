package coreversiontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var versionCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "VersionCompare returns Equal -- identical versions v0.0.1",
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
		Title: "VersionCompare returns LeftGreater -- left major v3.0 vs right v0.2.1",
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
		Title: "VersionCompare returns LeftLess -- left minor v0.0.2 vs right v0.2.1",
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
		Title: "VersionCompare returns Equal -- zero-padded v4 vs v4.0",
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
