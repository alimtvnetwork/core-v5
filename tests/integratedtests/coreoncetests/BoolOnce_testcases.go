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
			Title: "BoolOnce true -- Value returns true, String returns 'true'",
			ExpectedInput: args.Map{
				"value":  true,
				"string": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false -- Value returns false, String returns 'false'",
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
			Title: "BoolOnce.Value caches -- initializer runs exactly once",
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
			Title: "BoolOnce.Value false -- caches false, initializer runs once",
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
			Title: "BoolOnce true -- MarshalJSON returns 'true'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false -- MarshalJSON returns 'false'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "false",
			},
		},
		InitValue: false,
	},
}
