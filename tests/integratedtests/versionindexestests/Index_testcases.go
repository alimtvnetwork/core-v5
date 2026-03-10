package versionindexestests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var jsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Patch JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: args.Map{
			"indexName":  "Patch",
			"indexValue": "2",
		},
	},
	{
		Title: "Major JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Major index",
			"index": "Major",
		},
		ExpectedInput: args.Map{
			"indexName":  "Major",
			"indexValue": "0",
		},
	},
	{
		Title: "Build JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Build index",
			"index": "Build",
		},
		ExpectedInput: args.Map{
			"indexName":  "Build",
			"indexValue": "3",
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
		ExpectedInput: args.Map{
			"name":      "Minor",
			"nameValue": "Minor[1]",
		},
	},
	{
		Title: "Patch Name returns Patch",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: args.Map{
			"name":      "Patch",
			"nameValue": "Patch[2]",
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
		ExpectedInput: args.Map{
			"resultName":      "Patch",
			"resultNameValue": "Patch[2]",
		},
	},
	{
		Title: "JsonParseSelfInject overwrites Build with Major JSON",
		ArrangeInput: args.Map{
			"when":   "given Major JSON injected into Build",
			"source": "Major",
			"target": "Build",
		},
		ExpectedInput: args.Map{
			"resultName":      "Major",
			"resultNameValue": "Major[0]",
		},
	},
}
