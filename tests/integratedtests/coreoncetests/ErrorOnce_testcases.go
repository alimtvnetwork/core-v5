package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// ErrorOnce — Core (Value + state queries)
// =============================================================================

type errorOnceTestCase struct {
	Case       coretestcases.CaseV1
	InitError  string // empty means nil error
}

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
			ExpectedInput: []string{
				"fail", // r1
				"fail", // r2
				"fail", // r3
				"1",    // callCount
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
			Title: "ErrorOnce nil — IsNullOrEmpty true",
			ExpectedInput: []string{
				"true",
			},
		},
		InitError: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce empty string — IsNullOrEmpty true",
			ExpectedInput: []string{
				"true",
			},
		},
		InitError: "empty-marker", // special: will create errors.New("")
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce with message — IsNullOrEmpty false",
			ExpectedInput: []string{
				"false",
			},
		},
		InitError: "msg",
	},
}

// =============================================================================
// ErrorOnce — IsMessageEqual
// =============================================================================

type errorOnceMessageEqualTestCase struct {
	Case       coretestcases.CaseV1
	InitError  string
	MatchMsg   string
}

var errorOnceMessageEqualTestCases = []errorOnceMessageEqualTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'match' — IsMessageEqual 'match' true, 'other' false",
			ExpectedInput: []string{
				"true",  // IsMessageEqual(matchMsg)
				"false", // IsMessageEqual("other")
			},
		},
		InitError: "match",
		MatchMsg:  "match",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — IsMessageEqual always false",
			ExpectedInput: []string{
				"false",
				"false",
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
			ExpectedInput: []string{
				"true", // contains base
				"true", // contains extra
			},
		},
		InitError: "base",
		ExtraMsg:  "extra",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — ConcatNewString returns only additional message",
			ExpectedInput: []string{
				"only", // result equals extra message
			},
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
			ExpectedInput: []string{
				"true",       // no error
				"\"marshal\"", // marshaled value
			},
		},
		InitError: "marshal",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil — MarshalJSON returns '\"\"'",
			ExpectedInput: []string{
				"true", // no error
				"\"\"", // marshaled empty
			},
		},
		InitError: "",
	},
}
