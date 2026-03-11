package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// StringOnce -- Core (Value + String queries)
// =============================================================================

type stringOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
}

var stringOnceCoreTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns value 'hello' and isEmpty false -- 'hello' input",
			ExpectedInput: args.Map{
				"value":               "hello",
				"string":              "hello",
				"isEmpty":             false,
				"isEmptyOrWhitespace": false,
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns isEmpty true and isEmptyOrWhitespace true -- empty input",
			ExpectedInput: args.Map{
				"value":               "",
				"string":              "",
				"isEmpty":             true,
				"isEmptyOrWhitespace": true,
			},
		},
		InitValue: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns isEmpty false and isEmptyOrWhitespace true -- whitespace input",
			ExpectedInput: args.Map{
				"value":               "   ",
				"string":              "   ",
				"isEmpty":             false,
				"isEmptyOrWhitespace": true,
			},
		},
		InitValue: "   ",
	},
}

// =============================================================================
// StringOnce -- Caching
// =============================================================================

var stringOnceCachingTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        "hello",
				"r2":        "hello",
				"r3":        "hello",
				"callCount": 1,
			},
		},
		InitValue: "hello",
	},
}

// =============================================================================
// StringOnce -- String matching
// =============================================================================

type stringOnceMatchTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	MatchArg  string
}

var stringOnceMatchTestCases = []stringOnceMatchTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.IsEqual returns true for 'abc' and false for 'xyz' -- 'abc' input",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "abc",
		MatchArg:  "abc",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'hello world' -- IsContains 'world' true, 'xyz' false",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "hello world",
		MatchArg:  "world",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'prefix-data' -- HasPrefix 'prefix' true, 'data' false",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "prefix-data",
		MatchArg:  "prefix",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'data-suffix' -- HasSuffix 'suffix' true, 'data' false",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "data-suffix",
		MatchArg:  "suffix",
	},
}

// =============================================================================
// StringOnce -- Split
// =============================================================================

type stringOnceSplitTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	Splitter  string
	Method    string // "splitBy", "splitLeftRight", "splitLeftRightTrim"
}

var stringOnceSplitTestCases = []stringOnceSplitTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'a,b,c' -- SplitBy ','",
			ExpectedInput: args.Map{
				"partsLength": 3,
				"firstPart":   "a",
				"lastPart":    "c",
			},
		},
		InitValue: "a,b,c",
		Splitter:  ",",
		Method:    "splitBy",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'key=value' -- SplitLeftRight '='",
			ExpectedInput: args.Map{
				"left":  "key",
				"right": "value",
			},
		},
		InitValue: "key=value",
		Splitter:  "=",
		Method:    "splitLeftRight",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'nosplit' -- SplitLeftRight '=' returns full left, empty right",
			ExpectedInput: args.Map{
				"left":  "nosplit",
				"right": "",
			},
		},
		InitValue: "nosplit",
		Splitter:  "=",
		Method:    "splitLeftRight",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce ' key = value ' -- SplitLeftRightTrim '='",
			ExpectedInput: args.Map{
				"left":  "key",
				"right": "value",
			},
		},
		InitValue: " key = value ",
		Splitter:  "=",
		Method:    "splitLeftRightTrim",
	},
}

// =============================================================================
// StringOnce -- JSON
// =============================================================================

var stringOnceJsonTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'json' -- MarshalJSON returns '\"json\"'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"json\"",
			},
		},
		InitValue: "json",
	},
}
