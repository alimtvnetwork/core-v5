package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var instanceStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Instance String formats name=value",
		ArrangeInput: args.Map{
			"when":  "given name and value",
			"name":  "host",
			"value": "localhost",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "Instance with integer value formats correctly",
		ArrangeInput: args.Map{
			"when":  "given name and integer value",
			"name":  "port",
			"value": 8080,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "Instance with empty name still formats",
		ArrangeInput: args.Map{
			"when":  "given empty name",
			"name":  "",
			"value": "something",
		},
		ExpectedInput: []string{"true"},
	},
}

var instanceJsonStringTestCases = []coretestcases.CaseV1{
	{
		Title: "JsonString returns valid JSON",
		ArrangeInput: args.Map{
			"when":  "given name and value",
			"name":  "key",
			"value": "val",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "JsonString with integer value returns valid JSON",
		ArrangeInput: args.Map{
			"when":  "given name and integer value",
			"name":  "count",
			"value": 42,
		},
		ExpectedInput: []string{"true", "true"},
	},
}

var instanceDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Dispose clears name and value",
		ArrangeInput: args.Map{
			"when":  "given instance with data",
			"name":  "key",
			"value": "val",
		},
		ExpectedInput: []string{"", "true"},
	},
}
