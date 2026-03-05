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
		ExpectedInput: args.Two[string, string]{
			First:  "Patch", // indexName
			Second: "2",     // indexValue
		},
	},
	{
		Title: "Major JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Major index",
			"index": "Major",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Major", // indexName
			Second: "0",     // indexValue
		},
	},
	{
		Title: "Build JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Build index",
			"index": "Build",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Build", // indexName
			Second: "3",     // indexValue
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
		ExpectedInput: args.Two[string, string]{
			First:  "Minor",    // name
			Second: "Minor[1]", // nameValue
		},
	},
	{
		Title: "Patch Name returns Patch",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Patch",    // name
			Second: "Patch[2]", // nameValue
		},
	},
}

var jsonParseSelfInjectTestCases = []coretestcases.CaseV1{
	{
		Title: "JsonParseSelfInject overwrites Minor with Patch JSON",
		ArrangeInput: args.Map{
			"when":   "given Patch JSON injected into Minor",
			"source": "Patch",
			"target": "Minor",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Patch",    // resultName
			Second: "Patch[2]", // resultNameValue
		},
	},
	{
		Title: "JsonParseSelfInject overwrites Build with Major JSON",
		ArrangeInput: args.Map{
			"when":   "given Major JSON injected into Build",
			"source": "Major",
			"target": "Build",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Major",    // resultName
			Second: "Major[0]", // resultNameValue
		},
	},
}
