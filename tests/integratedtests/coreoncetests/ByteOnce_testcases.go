package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// ByteOnce -- Core
// =============================================================================

type byteOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue byte
}

var byteOnceCoreTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce returns value 42 and isPositive true -- input 42",
			ExpectedInput: args.Map{
				"value":      42,
				"int":        42,
				"string":     "42",
				"isEmpty":    false,
				"isZero":     false,
				"isNegative": false,
				"isPositive": true,
			},
		},
		InitValue: 42,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce 0 -- IsZero true, IsEmpty true",
			ExpectedInput: args.Map{
				"value":      0,
				"int":        0,
				"string":     "0",
				"isEmpty":    true,
				"isZero":     true,
				"isNegative": false,
				"isPositive": false,
			},
		},
		InitValue: 0,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce 255 -- max byte value",
			ExpectedInput: args.Map{
				"value":      255,
				"int":        255,
				"string":     "255",
				"isEmpty":    false,
				"isZero":     false,
				"isNegative": false,
				"isPositive": true,
			},
		},
		InitValue: 255,
	},
}

// =============================================================================
// ByteOnce -- Caching
// =============================================================================

var byteOnceCachingTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        10,
				"r2":        10,
				"callCount": 1,
			},
		},
		InitValue: 10,
	},
}

// =============================================================================
// ByteOnce -- JSON
// =============================================================================

var byteOnceJsonTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce 99 -- MarshalJSON returns '99'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "99",
			},
		},
		InitValue: 99,
	},
}

// =============================================================================
// ByteOnce -- Serialize
// =============================================================================

var byteOnceSerializeTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce 77 -- Serialize returns JSON bytes",
			ExpectedInput: args.Map{
				"noError":         true,
				"serializedValue": "77",
			},
		},
		InitValue: 77,
	},
}

// =============================================================================
// ByteOnce -- Constructor
// =============================================================================

var byteOnceConstructorTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewByteOnce (value) works correctly",
			ExpectedInput: args.Map{
				"constructedValue": 5,
			},
		},
		InitValue: 5,
	},
}
