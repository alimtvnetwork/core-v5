package versionindexestests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var jsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Patch JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: []string{
			"Patch",
			"2",
		},
	},
	{
		Title: "Major JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Major index",
			"index": "Major",
		},
		ExpectedInput: []string{
			"Major",
			"0",
		},
	},
	{
		Title: "Build JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Build index",
			"index": "Build",
		},
		ExpectedInput: []string{
			"Build",
			"3",
		},
	},
}

var nameAndNameValueTestCases = []coretestcases.CaseV1{
	{
		Title: "Minor Name returns Minor",
		ArrangeInput: args.Map{
			"when":  "given Minor index",
			"index": "Minor",
		},
		ExpectedInput: []string{
			"Minor",
			"Minor : 1",
		},
	},
	{
		Title: "Patch Name returns Patch",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: []string{
			"Patch",
			"Patch : 2",
		},
	},
}

var jsonParseSelfInjectTestCases = []coretestcases.CaseV1{
	{
		Title: "JsonParseSelfInject overwrites Minor with Patch JSON",
		ArrangeInput: args.Map{
			"when":       "given Patch JSON injected into Minor",
			"source":     "Patch",
			"target":     "Minor",
		},
		ExpectedInput: []string{
			"Patch",
			"Patch : 2",
		},
	},
	{
		Title: "JsonParseSelfInject overwrites Build with Major JSON",
		ArrangeInput: args.Map{
			"when":       "given Major JSON injected into Build",
			"source":     "Major",
			"target":     "Build",
		},
		ExpectedInput: []string{
			"Major",
			"Major : 0",
		},
	},
}
