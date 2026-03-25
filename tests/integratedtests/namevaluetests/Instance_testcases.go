package namevaluetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// StringAny tests
// ==========================================================================

var stringAnyStringTestCases = []coretestcases.CaseV1{
	{
		Title: "StringAny returns formatted string -- name 'host', value 'localhost'",
		ArrangeInput: args.Map{
			"when":  "given string name and string value",
			"name":  "host",
			"value": "localhost",
		},
		ExpectedInput: "host = localhost",
	},
	{
		Title: "StringAny returns formatted string -- name 'port', value 8080",
		ArrangeInput: args.Map{
			"when":  "given string name and int value",
			"name":  "port",
			"value": 8080,
		},
		ExpectedInput: "port = 8080",
	},
	{
		Title: "StringAny returns formatted string -- empty name, value 'something'",
		ArrangeInput: args.Map{
			"when":  "given empty name",
			"name":  "",
			"value": "something",
		},
		ExpectedInput: " = something",
	},
	{
		Title: "StringAny returns formatted string -- name 'key', nil value",
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
		Title: "StringString returns formatted string -- name 'env', value 'production'",
		ArrangeInput: args.Map{
			"when":  "given two strings",
			"name":  "env",
			"value": "production",
		},
		ExpectedInput: "env = production",
	},
	{
		Title: "StringString returns formatted string -- name 'env', empty value",
		ArrangeInput: args.Map{
			"when":  "given empty value",
			"name":  "env",
			"value": "",
		},
		ExpectedInput: "env = ",
	},
	{
		Title: "StringString returns formatted string -- both empty strings",
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
		Title: "StringInt returns formatted string -- name 'count', value 42",
		ArrangeInput: args.Map{
			"when":  "given name and positive int",
			"name":  "count",
			"value": 42,
		},
		ExpectedInput: "count = 42",
	},
	{
		Title: "StringInt returns formatted string -- name 'offset', value 0",
		ArrangeInput: args.Map{
			"when":  "given name and zero",
			"name":  "offset",
			"value": 0,
		},
		ExpectedInput: "offset = 0",
	},
	{
		Title: "StringInt returns formatted string -- name 'balance', value -100",
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
	Title: "StringMapAny returns valid JSON -- populated map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapAnyEmptyTestCase = coretestcases.CaseV1{
	Title: "StringMapAny returns valid JSON -- empty map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapAnyNilTestCase = coretestcases.CaseV1{
	Title: "StringMapAny returns valid JSON -- nil map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

// ==========================================================================
// StringMapString tests
// ==========================================================================

var stringMapStringPopulatedTestCase = coretestcases.CaseV1{
	Title: "StringMapString returns valid JSON -- populated map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

var stringMapStringNilTestCase = coretestcases.CaseV1{
	Title: "StringMapString returns valid JSON -- nil map",
	ExpectedInput: args.Map{
		"isValidJson":  true,
		"containsName": true,
	},
}

// ==========================================================================
// Dispose tests
// ==========================================================================

var disposeStringAnyTestCase = coretestcases.CaseV1{
	Title: "Dispose clears StringAny fields -- name and value reset",
	ExpectedInput: args.Map{
		"disposedName": "",
		"isNilValue":   true,
	},
}

var disposeStringStringTestCase = coretestcases.CaseV1{
	Title: "Dispose clears StringString fields -- name and value empty",
	ExpectedInput: args.Map{
		"disposedName":  "",
		"disposedValue": "",
	},
}

var disposeStringIntTestCase = coretestcases.CaseV1{
	Title: "Dispose clears StringInt fields -- name empty, value 0",
	ExpectedInput: args.Map{
		"disposedName":  "",
		"disposedValue": 0,
	},
}

// ==========================================================================
// JsonString tests
// ==========================================================================

var jsonStringStringAnyTestCase = coretestcases.CaseV1{
	Title: "JsonString returns valid JSON containing key -- StringAny input",
	ExpectedInput: args.Map{
		"isValidJson": true,
		"containsKey": true,
	},
}

var jsonStringStringIntTestCase = coretestcases.CaseV1{
	Title: "JsonString returns valid JSON containing number -- StringInt input",
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
		Title: "Collection returns correct length -- 3 StringAny items added",
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
		Title: "Collection returns length 0 -- no items added",
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
	Title: "VarNameValues returns collection containing name and value -- single StringAny",
	ExpectedInput: args.Map{
		"containsName":  true,
		"containsValue": true,
	},
}

var chmodMessageNameValuesTestCase = coretestcases.CaseV1{
	Title: "MessageNameValues returns collection with message and name-value -- StringAny input",
	ExpectedInput: args.Map{
		"containsMessage":   true,
		"containsNameValue": true,
	},
}

var chmodVarNameValuesEmptyTestCase = coretestcases.CaseV1{
	Title:         "VarNameValues returns empty -- empty StringAny slice",
	ExpectedInput: "",
}
