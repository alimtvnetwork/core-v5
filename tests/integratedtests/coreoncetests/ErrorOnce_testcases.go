package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// ErrorOnce — Core (Value + state queries)
// =============================================================================

type errorOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitError string // empty means nil error
}

// Note: 9 fields — exceeds args.Six, kept as []string for accuracy.
var errorOnceCoreTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce with error — HasError true, IsValid false, Message 'fail'",
			ExpectedInput: []string{
				"true",  // HasError
				"false", // IsValid
				"false", // IsSuccess
				"false", // IsEmpty
				"true",  // IsInvalid
				"true",  // IsFailed
				"true",  // HasAnyItem
				"true",  // IsDefined
				"fail",  // Message
			},
		},
		InitError: "fail",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — HasError false, IsValid true, Message empty",
			ExpectedInput: []string{
				"false", // HasError
				"true",  // IsValid
				"true",  // IsSuccess
				"true",  // IsEmpty
				"false", // IsInvalid
				"false", // IsFailed
				"false", // HasAnyItem
				"false", // IsDefined
				"",      // Message
			},
		},
		InitError: "",
	},
}

// =============================================================================
// ErrorOnce — Caching
// =============================================================================

var errorOnceCachingTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.Value caches — initializer runs exactly once",
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "fail", // r1
				Second: "fail", // r2
				Third:  "fail", // r3
				Fourth: "1",    // callCount
			},
		},
		InitError: "fail",
	},
}

// =============================================================================
// ErrorOnce — IsNullOrEmpty
// =============================================================================

var errorOnceNullOrEmptyTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "ErrorOnce nil — IsNullOrEmpty true",
			ExpectedInput: "true", // isNullOrEmpty
		},
		InitError: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "ErrorOnce empty string — IsNullOrEmpty true",
			ExpectedInput: "true", // isNullOrEmpty
		},
		InitError: "empty-marker", // special: will create errors.New("")
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "ErrorOnce with message — IsNullOrEmpty false",
			ExpectedInput: "false", // isNullOrEmpty
		},
		InitError: "msg",
	},
}

// =============================================================================
// ErrorOnce — IsMessageEqual
// =============================================================================

type errorOnceMessageEqualTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	MatchMsg  string
}

var errorOnceMessageEqualTestCases = []errorOnceMessageEqualTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'match' — IsMessageEqual 'match' true, 'other' false",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // isMessageEqualMatch
				Second: "false", // isMessageEqualOther
			},
		},
		InitError: "match",
		MatchMsg:  "match",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — IsMessageEqual always false",
			ExpectedInput: args.Two[string, string]{
				First:  "false", // isMessageEqualMatch
				Second: "false", // isMessageEqualOther
			},
		},
		InitError: "",
		MatchMsg:  "anything",
	},
}

// =============================================================================
// ErrorOnce — ConcatNew
// =============================================================================

type errorOnceConcatTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	ExtraMsg  string
}

var errorOnceConcatTestCases = []errorOnceConcatTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'base' — ConcatNewString contains both 'base' and 'extra'",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // containsBase
				Second: "true", // containsExtra
			},
		},
		InitError: "base",
		ExtraMsg:  "extra",
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "ErrorOnce nil — ConcatNewString returns only additional message",
			ExpectedInput: "only", // resultEqualsExtraMessage
		},
		InitError: "",
		ExtraMsg:  "only",
	},
}

// =============================================================================
// ErrorOnce — JSON
// =============================================================================

var errorOnceJsonTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'marshal' — MarshalJSON returns '\"marshal\"'",
			ExpectedInput: args.Two[string, string]{
				First:  "true",       // noError
				Second: "\"marshal\"", // marshaledValue
			},
		},
		InitError: "marshal",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — MarshalJSON returns '\"\"'",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "\"\"", // marshaledEmpty
			},
		},
		InitError: "",
	},
}
