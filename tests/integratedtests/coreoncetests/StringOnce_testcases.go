package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/args"
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
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "hello", // value
				Second: "hello", // stringMethod
				Third:  "false", // isEmpty
				Fourth: "false", // isEmptyOrWhitespace
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce empty — IsEmpty true, IsEmptyOrWhitespace true",
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "",     // value
				Second: "",     // stringMethod
				Third:  "true", // isEmpty
				Fourth: "true", // isEmptyOrWhitespace
			},
		},
		InitValue: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce whitespace — IsEmpty false, IsEmptyOrWhitespace true",
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "   ",   // value
				Second: "   ",   // stringMethod
				Third:  "false", // isEmpty
				Fourth: "true",  // isEmptyOrWhitespace
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
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "hello", // r1
				Second: "hello", // r2
				Third:  "hello", // r3
				Fourth: "1",     // callCount
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
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // isEqualMatch
				Second: "false", // isEqualOther
			},
		},
		InitValue: "abc",
		MatchArg:  "abc",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'hello world' — IsContains 'world' true, 'xyz' false",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // isContainsMatch
				Second: "false", // isContainsOther
			},
		},
		InitValue: "hello world",
		MatchArg:  "world",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'prefix-data' — HasPrefix 'prefix' true, 'data' false",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // hasPrefixMatch
				Second: "false", // hasPrefixOther
			},
		},
		InitValue: "prefix-data",
		MatchArg:  "prefix",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'data-suffix' — HasSuffix 'suffix' true, 'data' false",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // hasSuffixMatch
				Second: "false", // hasSuffixOther
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
			ExpectedInput: args.Three[string, string, string]{
				First:  "3", // partsLength
				Second: "a", // firstPart
				Third:  "c", // lastPart
			},
		},
		InitValue: "a,b,c",
		Splitter:  ",",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'key=value' — SplitLeftRight '='",
			ExpectedInput: args.Two[string, string]{
				First:  "key",   // left
				Second: "value", // right
			},
		},
		InitValue: "key=value",
		Splitter:  "=",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce 'nosplit' — SplitLeftRight '=' returns full left, empty right",
			ExpectedInput: args.Two[string, string]{
				First:  "nosplit", // left
				Second: "",        // right
			},
		},
		InitValue: "nosplit",
		Splitter:  "=",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce ' key = value ' — SplitLeftRightTrim '='",
			ExpectedInput: args.Two[string, string]{
				First:  "key",   // leftTrimmed
				Second: "value", // rightTrimmed
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
			ExpectedInput: args.Two[string, string]{
				First:  "true",     // noError
				Second: "\"json\"", // marshaledValue
			},
		},
		InitValue: "json",
	},
}
