package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// BoolOnce -- Core (Value + String)
// =============================================================================

type boolOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue bool
}

var boolOnceCoreTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns true and String returns 'true' -- init true",
			ExpectedInput: args.Map{
				"value":  true,
				"string": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns false and String returns 'false' -- init false",
			ExpectedInput: args.Map{
				"value":  false,
				"string": "false",
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce -- Caching (call count verification)
// =============================================================================

var boolOnceCachingTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns cached true -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        true,
				"r2":        true,
				"r3":        true,
				"callCount": 1,
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns cached false -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        false,
				"r2":        false,
				"r3":        false,
				"callCount": 1,
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce -- JSON
// =============================================================================

var boolOnceJsonTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.MarshalJSON returns 'true' -- init true",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.MarshalJSON returns 'false' -- init false",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "false",
			},
		},
		InitValue: false,
	},
}
