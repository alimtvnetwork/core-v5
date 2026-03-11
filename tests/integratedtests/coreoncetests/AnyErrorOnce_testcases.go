package coreoncetests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// AnyErrorOnce -- Core
// =============================================================================

type anyErrorOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue any
	InitErr   error
}

var anyErrorOnceCoreTestCases = []anyErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce returns isDefined true and isNull false -- 'hello' value, nil error",
			ExpectedInput: args.Map{
				"hasError":   false,
				"isValid":    true,
				"isSuccess":  true,
				"isInvalid":  false,
				"isFailed":   false,
				"isNull":     false,
				"isEmpty":    false,
				"hasAnyItem": true,
				"isDefined":  true,
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce returns isNull true and isEmpty true -- nil value, nil error",
			ExpectedInput: args.Map{
				"hasError":   false,
				"isValid":    true,
				"isSuccess":  true,
				"isInvalid":  false,
				"isFailed":   false,
				"isNull":     true,
				"isEmpty":    true,
				"hasAnyItem": false,
				"isDefined":  false,
			},
		},
		InitValue: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce returns hasError true and isFailed true -- nil value, error set",
			ExpectedInput: args.Map{
				"hasError":   true,
				"isValid":    false,
				"isSuccess":  false,
				"isInvalid":  true,
				"isFailed":   true,
				"isNull":     true,
				"isEmpty":    false,
				"hasAnyItem": true,
				"isDefined":  true,
			},
		},
		InitErr: errors.New("fail"),
	},
}

// =============================================================================
// AnyErrorOnce -- Caching
// =============================================================================

var anyErrorOnceCachingTestCases = []anyErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce.Value returns cached result -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount":      1,
				"executeEqValue": true,
			},
		},
		InitValue: "cached",
	},
}

// =============================================================================
// AnyErrorOnce -- ValueMust / ExecuteMust
// =============================================================================

var anyErrorOnceMustSuccessTestCase = anyErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyErrorOnce.ValueMust returns value without panic -- no error",
		ExpectedInput: args.Map{
			"didPanic": false,
		},
	},
	InitValue: "ok",
}

var anyErrorOnceMustPanicTestCase = anyErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyErrorOnce.ValueMust returns panic -- error set",
		ExpectedInput: args.Map{
			"didPanic": true,
		},
	},
	InitErr: errors.New("must-fail"),
}

// =============================================================================
// AnyErrorOnce -- Cast
// =============================================================================

var anyErrorOnceCastStringTestCase = anyErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyErrorOnce.CastValueString returns 'cast' and castSuccess true -- string value",
		ExpectedInput: args.Map{
			"castValue":   "cast",
			"castSuccess": true,
			"noError":     true,
		},
	},
	InitValue: "cast",
}

// =============================================================================
// AnyErrorOnce -- JSON
// =============================================================================

var anyErrorOnceJsonTestCases = []anyErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce.Serialize returns bytes without error -- 'json' value",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitValue: "json",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce.Serialize returns error -- error set",
			ExpectedInput: args.Map{
				"hasError": true,
			},
		},
		InitErr: errors.New("ser-fail"),
	},
}

// =============================================================================
// AnyErrorOnce -- Constructor
// =============================================================================

var anyErrorOnceConstructorTestCases = []anyErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewAnyErrorOnce returns isNull false and noError true -- value input",
			ExpectedInput: args.Map{
				"isNull":  false,
				"noError": true,
			},
		},
		InitValue: "val",
	},
}