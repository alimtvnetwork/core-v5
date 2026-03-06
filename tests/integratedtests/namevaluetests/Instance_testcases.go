package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// StringAny tests
// ==========================================================================

var stringAnyStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringAny formats name=value correctly",
		ArrangeInput: args.Map{
			"when":  "given string name and string value",
			"name":  "host",
			"value": "localhost",
		},
		ExpectedInput: "host = localhost",
	},
	{
		Title: "Positive: StringAny with integer value",
		ArrangeInput: args.Map{
			"when":  "given string name and int value",
			"name":  "port",
			"value": 8080,
		},
		ExpectedInput: "port = 8080",
	},
	{
		Title: "Negative: StringAny with empty name",
		ArrangeInput: args.Map{
			"when":  "given empty name",
			"name":  "",
			"value": "something",
		},
		ExpectedInput: " = something",
	},
	{
		Title: "Negative: StringAny with nil value",
		ArrangeInput: args.Map{
			"when":  "given nil value",
			"name":  "key",
			"value": nil,
		},
		ExpectedInput: "key = <nil>",
	},
}

// ==========================================================================
// StringString tests
// ==========================================================================

var stringStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringString formats both strings",
		ArrangeInput: args.Map{
			"when":  "given two strings",
			"name":  "env",
			"value": "production",
		},
		ExpectedInput: "env = production",
	},
	{
		Title: "Negative: StringString with empty value",
		ArrangeInput: args.Map{
			"when":  "given empty value",
			"name":  "env",
			"value": "",
		},
		ExpectedInput: "env = ",
	},
	{
		Title: "Negative: StringString with both empty",
		ArrangeInput: args.Map{
			"when":  "given both empty",
			"name":  "",
			"value": "",
		},
		ExpectedInput: " = ",
	},
}

// ==========================================================================
// StringInt tests
// ==========================================================================

var stringIntTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringInt formats correctly",
		ArrangeInput: args.Map{
			"when":  "given name and positive int",
			"name":  "count",
			"value": 42,
		},
		ExpectedInput: "count = 42",
	},
	{
		Title: "Positive: StringInt with zero",
		ArrangeInput: args.Map{
			"when":  "given name and zero",
			"name":  "offset",
			"value": 0,
		},
		ExpectedInput: "offset = 0",
	},
	{
		Title: "Negative: StringInt with negative value",
		ArrangeInput: args.Map{
			"when":  "given negative int",
			"name":  "balance",
			"value": -100,
		},
		ExpectedInput: "balance = -100",
	},
}

// ==========================================================================
// StringMapAny tests
// ==========================================================================

var stringMapAnyPopulatedTestCase = coretestcases.CaseV1{
	Title: "Positive: StringMapAny with populated map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapAnyEmptyTestCase = coretestcases.CaseV1{
	Title: "Negative: StringMapAny with empty map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapAnyNilTestCase = coretestcases.CaseV1{
	Title: "Negative: StringMapAny with nil map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

// ==========================================================================
// StringMapString tests
// ==========================================================================

var stringMapStringPopulatedTestCase = coretestcases.CaseV1{
	Title: "Positive: StringMapString with populated map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapStringNilTestCase = coretestcases.CaseV1{
	Title: "Negative: StringMapString with nil map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

// ==========================================================================
// Dispose tests
// ==========================================================================

var disposeStringAnyTestCase = coretestcases.CaseV1{
	Title: "Positive: Dispose clears StringAny fields",
	ExpectedInput: args.Map{
		"disposedName":  "",
		"isNilValue":    true,
	},
}

var disposeStringStringTestCase = coretestcases.CaseV1{
	Title: "Positive: Dispose clears StringString fields",
	ExpectedInput: args.Map{
		"disposedName":  "",
		"disposedValue": "",
	},
}

var disposeStringIntTestCase = coretestcases.CaseV1{
	Title: "Positive: Dispose clears StringInt to zero",
	ExpectedInput: args.Map{
		"disposedName":  "",
		"disposedValue": 0,
	},
}

// ==========================================================================
// JsonString tests
// ==========================================================================

var jsonStringStringAnyTestCase = coretestcases.CaseV1{
	Title: "Positive: StringAny JsonString contains key",
	ExpectedInput: args.Map{
		"isValidJson": true,
		"containsKey": true,
	},
}

var jsonStringStringIntTestCase = coretestcases.CaseV1{
	Title: "Positive: StringInt JsonString contains number",
	ExpectedInput: args.Map{
		"isValidJson":    true,
		"containsNumber": true,
	},
}

// ==========================================================================
// Collection tests
// ==========================================================================

var collectionTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Collection adds items and returns correct length",
		ArrangeInput: args.Map{
			"when":  "given multiple StringAny items",
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length":  3,
			"isEmpty": false,
		},
	},
	{
		Title: "Negative: Empty collection returns length 0",
		ArrangeInput: args.Map{
			"when":  "given no items",
			"count": 0,
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================================================
// Chmod integration tests
// ==========================================================================

var chmodVarNameValuesSingleTestCase = coretestcases.CaseV1{
	Title: "Positive: StringAny works in errcore.VarNameValues",
	ExpectedInput: args.Map{
		"containsName":  true,
		"containsValue": true,
	},
}

var chmodMessageNameValuesTestCase = coretestcases.CaseV1{
	Title: "Positive: StringAny works in errcore.MessageNameValues",
	ExpectedInput: args.Map{
		"containsMessage":   true,
		"containsNameValue": true,
	},
}

var chmodVarNameValuesEmptyTestCase = coretestcases.CaseV1{
	Title:         "Negative: Empty StringAny slice in VarNameValues returns empty",
	ExpectedInput: "",
}
