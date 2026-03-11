package argstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── Map test cases ──

var extMapGetTestCases = []coretestcases.CaseV1{
	{
		Title: "Map.Get existing key -- returns value and true",
		ArrangeInput: args.Map{
			"when":  "key exists",
			"key":   "name",
			"name":  "hello",
		},
		ExpectedInput: args.Map{
			"isValid": true,
			"value":   "hello",
		},
	},
	{
		Title: "Map.Get missing key -- returns nil and false",
		ArrangeInput: args.Map{
			"when": "key missing",
			"key":  "missing",
		},
		ExpectedInput: args.Map{
			"isValid": false,
		},
	},
}

var extMapTypedGetTestCases = []coretestcases.CaseV1{
	{
		Title: "Map.GetAsInt valid -- returns int and true",
		ArrangeInput: args.Map{
			"when":  "int value present",
			"count": 42,
		},
		ExpectedInput: args.Map{
			"intVal":   42,
			"intValid": true,
		},
	},
	{
		Title: "Map.GetAsBool valid -- returns bool and true",
		ArrangeInput: args.Map{
			"when":   "bool value present",
			"active": true,
		},
		ExpectedInput: args.Map{
			"boolVal":   true,
			"boolValid": true,
		},
	},
	{
		Title: "Map.GetAsString valid -- returns string and true",
		ArrangeInput: args.Map{
			"when": "string value present",
			"text": "hello",
		},
		ExpectedInput: args.Map{
			"strVal":   "hello",
			"strValid": true,
		},
	},
}

var extMapCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "Map.CompileToStrings -- sorted key-value lines",
		ArrangeInput: args.Map{
			"when": "map with two entries",
			"b":    2,
			"a":    1,
		},
		ExpectedInput: args.Map{
			"lineCount": 3,
		},
	},
	{
		Title: "Map.CompileToStrings -- empty map",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"lineCount": 0,
		},
	},
}

// ── One test cases ──

var extOneTestCases = []coretestcases.CaseV1{
	{
		Title: "One -- has first and expect",
		ArrangeInput: args.Map{
			"when":   "One with first and expect",
			"first":  "hello",
			"expect": 42,
		},
		ExpectedInput: args.Map{
			"hasFirst":   true,
			"hasExpect":  true,
			"argsCount":  1,
			"firstItem":  "hello",
			"expected":   42,
			"validCount": 1,
		},
	},
	{
		Title: "One -- nil first no expect",
		ArrangeInput: args.Map{
			"when": "One with nil first",
		},
		ExpectedInput: args.Map{
			"hasFirst":   false,
			"hasExpect":  false,
			"argsCount":  1,
			"validCount": 0,
		},
	},
}

// ── Two test cases ──

var extTwoTestCases = []coretestcases.CaseV1{
	{
		Title: "Two -- has both values",
		ArrangeInput: args.Map{
			"when":   "Two with first and second",
			"first":  "a",
			"second": "b",
		},
		ExpectedInput: args.Map{
			"hasFirst":   true,
			"hasSecond":  true,
			"argsCount":  2,
			"validCount": 2,
		},
	},
}

// ── LeftRight test cases ──

var extLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "LeftRight -- has left and right",
		ArrangeInput: args.Map{
			"when":  "LeftRight with both",
			"left":  10,
			"right": 20,
		},
		ExpectedInput: args.Map{
			"hasLeft":    true,
			"hasRight":   true,
			"hasFirst":   true,
			"hasSecond":  true,
			"argsCount":  2,
			"validCount": 2,
			"firstItem":  10,
			"secondItem": 20,
		},
	},
}

// ── String test cases ──

var extStringTestCases = []coretestcases.CaseV1{
	{
		Title: "String -- basic operations",
		ArrangeInput: args.Map{
			"when":  "String value hello",
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"length":              5,
			"isEmpty":             false,
			"hasCharacter":        true,
			"isDefined":           true,
			"isEmptyOrWhitespace": false,
			"asciiLength":         5,
		},
	},
	{
		Title: "String -- empty string",
		ArrangeInput: args.Map{
			"when":  "empty string",
			"input": "",
		},
		ExpectedInput: args.Map{
			"length":              0,
			"isEmpty":             true,
			"hasCharacter":        false,
			"isDefined":           false,
			"isEmptyOrWhitespace": true,
			"asciiLength":         0,
		},
	},
}
