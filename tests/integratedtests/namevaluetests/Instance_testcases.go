package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// region StringAny (backward-compat) tests

var stringAnyStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringAny formats name=value correctly",
		ArrangeInput: args.Map{
			"when":  "given string name and string value",
			"name":  "host",
			"value": "localhost",
		},
		ExpectedInput: []string{"host = localhost"},
	},
	{
		Title: "Positive: StringAny with integer value",
		ArrangeInput: args.Map{
			"when":  "given string name and int value",
			"name":  "port",
			"value": 8080,
		},
		ExpectedInput: []string{"port = 8080"},
	},
	{
		Title: "Negative: StringAny with empty name",
		ArrangeInput: args.Map{
			"when":  "given empty name",
			"name":  "",
			"value": "something",
		},
		ExpectedInput: []string{" = something"},
	},
	{
		Title: "Negative: StringAny with nil value",
		ArrangeInput: args.Map{
			"when":  "given nil value",
			"name":  "key",
			"value": nil,
		},
		ExpectedInput: []string{"key = <nil>"},
	},
}

// endregion

// region StringString tests

var stringStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringString formats both strings",
		ArrangeInput: args.Map{
			"when":  "given two strings",
			"name":  "env",
			"value": "production",
		},
		ExpectedInput: []string{"env = production"},
	},
	{
		Title: "Negative: StringString with empty value",
		ArrangeInput: args.Map{
			"when":  "given empty value",
			"name":  "env",
			"value": "",
		},
		ExpectedInput: []string{"env = "},
	},
	{
		Title: "Negative: StringString with both empty",
		ArrangeInput: args.Map{
			"when":  "given both empty",
			"name":  "",
			"value": "",
		},
		ExpectedInput: []string{" = "},
	},
}

// endregion

// region StringInt tests

var stringIntTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringInt formats correctly",
		ArrangeInput: args.Map{
			"when":  "given name and positive int",
			"name":  "count",
			"value": 42,
		},
		ExpectedInput: []string{"count = 42"},
	},
	{
		Title: "Positive: StringInt with zero",
		ArrangeInput: args.Map{
			"when":  "given name and zero",
			"name":  "offset",
			"value": 0,
		},
		ExpectedInput: []string{"offset = 0"},
	},
	{
		Title: "Negative: StringInt with negative value",
		ArrangeInput: args.Map{
			"when":  "given negative int",
			"name":  "balance",
			"value": -100,
		},
		ExpectedInput: []string{"balance = -100"},
	},
}

// endregion

// region StringMapAny tests

var stringMapAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringMapAny with populated map",
		ArrangeInput: args.Map{
			"when": "given name and populated map",
			"name": "config",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Negative: StringMapAny with empty map",
		ArrangeInput: args.Map{
			"when": "given name and empty map",
			"name": "empty",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Negative: StringMapAny with nil map",
		ArrangeInput: args.Map{
			"when": "given name and nil map",
			"name": "nothing",
		},
		ExpectedInput: []string{"true", "true"},
	},
}

// endregion

// region StringMapString tests

var stringMapStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringMapString with populated map",
		ArrangeInput: args.Map{
			"when": "given name and string map",
			"name": "headers",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Negative: StringMapString with nil map",
		ArrangeInput: args.Map{
			"when": "given name and nil map",
			"name": "nothing",
		},
		ExpectedInput: []string{"true", "true"},
	},
}

// endregion

// region Dispose tests

var genericDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Dispose clears StringAny fields",
		ArrangeInput: args.Map{
			"when":  "given StringAny with data",
			"type":  "stringany",
			"name":  "key",
			"value": "val",
		},
		ExpectedInput: []string{"", "true"},
	},
	{
		Title: "Positive: Dispose clears StringString fields",
		ArrangeInput: args.Map{
			"when":  "given StringString with data",
			"type":  "stringstring",
			"name":  "key",
			"value": "val",
		},
		ExpectedInput: []string{"", ""},
	},
	{
		Title: "Positive: Dispose clears StringInt to zero",
		ArrangeInput: args.Map{
			"when":  "given StringInt with data",
			"type":  "stringint",
			"name":  "count",
			"value": 42,
		},
		ExpectedInput: []string{"", "0"},
	},
}

// endregion

// region JsonString tests

var genericJsonStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringAny JsonString contains key",
		ArrangeInput: args.Map{
			"when":  "given valid StringAny",
			"name":  "server",
			"value": "api.example.com",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Positive: StringInt JsonString contains number",
		ArrangeInput: args.Map{
			"when":  "given valid StringInt",
			"name":  "port",
			"value": 443,
		},
		ExpectedInput: []string{"true", "true"},
	},
}

// endregion

// region Collection tests

var collectionTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Collection adds items and returns correct length",
		ArrangeInput: args.Map{
			"when":  "given multiple StringAny items",
			"count": 3,
		},
		ExpectedInput: []string{"3", "false"},
	},
	{
		Title: "Negative: Empty collection returns length 0",
		ArrangeInput: args.Map{
			"when":  "given no items",
			"count": 0,
		},
		ExpectedInput: []string{"0", "true"},
	},
}

// endregion

// region Chmod integration tests

var chmodIntegrationTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringAny works in errcore.VarNameValues",
		ArrangeInput: args.Map{
			"when": "given StringAny used in errcore formatting",
			"name": "Location",
			"path": "/tmp/test",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Positive: StringAny works in errcore.MessageNameValues",
		ArrangeInput: args.Map{
			"when":    "given StringAny used in message formatting",
			"message": "chmod verification failed",
			"name":    "Path",
			"path":    "/usr/local/bin",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Negative: Empty StringAny slice in VarNameValues returns empty",
		ArrangeInput: args.Map{
			"when": "given no name values",
		},
		ExpectedInput: []string{""},
	},
}

// endregion
