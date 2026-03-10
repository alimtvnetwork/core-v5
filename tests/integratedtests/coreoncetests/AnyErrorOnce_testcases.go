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
			Title: "AnyErrorOnce 'hello'/nil -- no error, not null, isDefined",
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
			Title: "AnyErrorOnce nil/nil -- isNull, isEmpty",
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
			Title: "AnyErrorOnce nil/error -- hasError, isFailed",
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
			Title: "AnyErrorOnce.Value caches -- initializer runs once",
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
		Title: "AnyErrorOnce.ValueMust succeeds without panic",
		ExpectedInput: args.Map{
			"didPanic": false,
		},
	},
	InitValue: "ok",
}

var anyErrorOnceMustPanicTestCase = anyErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyErrorOnce.ValueMust panics on error",
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
		Title: "AnyErrorOnce 'cast' -- CastValueString succeeds",
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
			Title: "AnyErrorOnce 'json' -- Serialize succeeds",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitValue: "json",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyErrorOnce with error -- Serialize returns error",
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
			Title: "NewAnyErrorOnce (value) works correctly",
			ExpectedInput: args.Map{
				"isNull":  false,
				"noError": true,
			},
		},
		InitValue: "val",
	},
}
