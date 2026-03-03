package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// StringOnce — Core (Value + String queries)
// =============================================================================

type stringOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
}

var stringOnceCoreTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'hello' — Value, String, IsEmpty false",
			ExpectedInput: []string{
				"hello", // Value
				"hello", // String
				"false", // IsEmpty
				"false", // IsEmptyOrWhitespace
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce empty — IsEmpty true, IsEmptyOrWhitespace true",
			ExpectedInput: []string{
				"",     // Value
				"",     // String
				"true", // IsEmpty
				"true", // IsEmptyOrWhitespace
			},
		},
		InitValue: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce whitespace — IsEmpty false, IsEmptyOrWhitespace true",
			ExpectedInput: []string{
				"   ",   // Value
				"   ",   // String
				"false", // IsEmpty
				"true",  // IsEmptyOrWhitespace
			},
		},
		InitValue: "   ",
	},
}

// =============================================================================
// StringOnce — Caching
// =============================================================================

var stringOnceCachingTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.Value caches — initializer runs exactly once",
			ExpectedInput: []string{
				"hello", // r1
				"hello", // r2
				"hello", // r3
				"1",     // callCount
			},
		},
		InitValue: "hello",
	},
}

// =============================================================================
// StringOnce — String matching
// =============================================================================

type stringOnceMatchTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	MatchArg  string
}

var stringOnceMatchTestCases = []stringOnceMatchTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'abc' — IsEqual 'abc' true, 'xyz' false",
			ExpectedInput: []string{
				"true",  // IsEqual(matchArg)
				"false", // IsEqual("xyz")
			},
		},
		InitValue: "abc",
		MatchArg:  "abc",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'hello world' — IsContains 'world' true, 'xyz' false",
			ExpectedInput: []string{
				"true",  // IsContains(matchArg)
				"false", // IsContains("xyz")
			},
		},
		InitValue: "hello world",
		MatchArg:  "world",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'prefix-data' — HasPrefix 'prefix' true, 'data' false",
			ExpectedInput: []string{
				"true",  // HasPrefix(matchArg)
				"false", // HasPrefix("data")
			},
		},
		InitValue: "prefix-data",
		MatchArg:  "prefix",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'data-suffix' — HasSuffix 'suffix' true, 'data' false",
			ExpectedInput: []string{
				"true",  // HasSuffix(matchArg)
				"false", // HasSuffix("data")
			},
		},
		InitValue: "data-suffix",
		MatchArg:  "suffix",
	},
}

// =============================================================================
// StringOnce — Split
// =============================================================================

type stringOnceSplitTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	Splitter  string
}

var stringOnceSplitTestCases = []stringOnceSplitTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'a,b,c' — SplitBy ','",
			ExpectedInput: []string{
				"3", // parts length
				"a", // parts[0]
				"c", // parts[2]
			},
		},
		InitValue: "a,b,c",
		Splitter:  ",",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'key=value' — SplitLeftRight '='",
			ExpectedInput: []string{
				"key",   // left
				"value", // right
			},
		},
		InitValue: "key=value",
		Splitter:  "=",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'nosplit' — SplitLeftRight '=' returns full left, empty right",
			ExpectedInput: []string{
				"nosplit", // left
				"",        // right
			},
		},
		InitValue: "nosplit",
		Splitter:  "=",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce ' key = value ' — SplitLeftRightTrim '='",
			ExpectedInput: []string{
				"key",   // left trimmed
				"value", // right trimmed
			},
		},
		InitValue: " key = value ",
		Splitter:  "=",
	},
}

// =============================================================================
// StringOnce — JSON
// =============================================================================

var stringOnceJsonTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'json' — MarshalJSON returns '\"json\"'",
			ExpectedInput: []string{
				"true",     // no error
				"\"json\"", // marshaled value
			},
		},
		InitValue: "json",
	},
}
