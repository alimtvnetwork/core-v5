package coreoncetests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// ErrorOnce -- Core (Value + state queries)
// =============================================================================

type errorOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitError string // empty means nil error
}

var errorOnceCoreTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns hasError true and message 'fail' -- error 'fail'",
			ExpectedInput: args.Map{
				"hasError":   true,
				"isValid":    false,
				"isSuccess":  false,
				"isEmpty":    false,
				"isInvalid":  true,
				"isFailed":   true,
				"hasAnyItem": true,
				"isDefined":  true,
				"message":    "fail",
			},
		},
		InitError: "fail",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isValid true and message empty -- nil error",
			ExpectedInput: args.Map{
				"hasError":   false,
				"isValid":    true,
				"isSuccess":  true,
				"isEmpty":    true,
				"isInvalid":  false,
				"isFailed":   false,
				"hasAnyItem": false,
				"isDefined":  false,
				"message":    "",
			},
		},
		InitError: "",
	},
}

// =============================================================================
// ErrorOnce -- Caching
// =============================================================================

var errorOnceCachingTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        "fail",
				"r2":        "fail",
				"r3":        "fail",
				"callCount": 1,
			},
		},
		InitError: "fail",
	},
}

// =============================================================================
// ErrorOnce -- IsNullOrEmpty
// =============================================================================

var errorOnceNullOrEmptyTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty true -- nil error",
			ExpectedInput: args.Map{
				"isNullOrEmpty": true,
			},
		},
		InitError: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty true -- empty string error",
			ExpectedInput: args.Map{
				"isNullOrEmpty": true,
			},
		},
		InitError: "empty-marker", // special: will create errors.New("")
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty false -- error 'msg'",
			ExpectedInput: args.Map{
				"isNullOrEmpty": false,
			},
		},
		InitError: "msg",
	},
}

// =============================================================================
// ErrorOnce -- IsMessageEqual
// =============================================================================

type errorOnceMessageEqualTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	MatchMsg  string
}

var errorOnceMessageEqualTestCases = []errorOnceMessageEqualTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'match' -- IsMessageEqual 'match' true, 'other' false",
			ExpectedInput: args.Map{
				"isMessageEqualMatch": true,
				"isMessageEqualOther": false,
			},
		},
		InitError: "match",
		MatchMsg:  "match",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil -- IsMessageEqual always false",
			ExpectedInput: args.Map{
				"isMessageEqualMatch": false,
				"isMessageEqualOther": false,
			},
		},
		InitError: "",
		MatchMsg:  "anything",
	},
}

// =============================================================================
// ErrorOnce -- ConcatNew
// =============================================================================

type errorOnceConcatTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	ExtraMsg  string
}

var errorOnceConcatTestCases = []errorOnceConcatTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'base' -- ConcatNewString contains both 'base' and 'extra'",
			ExpectedInput: args.Map{
				"containsBase":  true,
				"containsExtra": true,
			},
		},
		InitError: "base",
		ExtraMsg:  "extra",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil -- ConcatNewString returns only additional message",
			ExpectedInput: args.Map{
				"result": "\"only\"",
			},
		},
		InitError: "",
		ExtraMsg:  "only",
	},
}

// =============================================================================
// ErrorOnce -- JSON
// =============================================================================

var errorOnceJsonTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce 'marshal' -- MarshalJSON returns '\"marshal\"'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"marshal\"",
			},
		},
		InitError: "marshal",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce nil -- MarshalJSON returns '\"\"'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"\"",
			},
		},
		InitError: "",
	},
}

// unused import guard
var _ = errors.New
