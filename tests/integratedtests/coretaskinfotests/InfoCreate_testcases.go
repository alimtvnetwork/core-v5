package coretaskinfotests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var infoCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Default creates info with name, desc, url",
		ArrangeInput: args.Map{
			"when": "given default info creation",
			"name": "some name",
			"desc": "some desc",
			"url":  "some url",
		},
		ExpectedInput: []string{
			"some name",
			"some desc",
			"some url",
			"false",
			"true",
		},
	},
	{
		Title: "New.Info.Examples creates info with examples",
		ArrangeInput: args.Map{
			"when":     "given info with examples",
			"name":     "example name",
			"desc":     "example desc",
			"url":      "example url",
			"examples": []string{"ex1", "ex2"},
		},
		ExpectedInput: []string{
			"example name",
			"example desc",
			"example url",
			"false",
			"true",
		},
	},
	{
		Title: "Nil info returns safe defaults",
		ArrangeInput: args.Map{
			"when":  "given nil info",
			"isNil": true,
		},
		ExpectedInput: []string{
			"",
			"",
			"",
			"true",
			"false",
		},
	},
}

var infoSerializeTestCases = []coretestcases.CaseV1{
	{
		Title: "Info serializes and deserializes correctly",
		ArrangeInput: args.Map{
			"when": "given round-trip serialize/deserialize",
			"name": "round-trip name",
			"desc": "round-trip desc",
			"url":  "round-trip url",
		},
		ExpectedInput: []string{
			"round-trip name",
			"round-trip desc",
			"round-trip url",
			"true",
		},
	},
}
