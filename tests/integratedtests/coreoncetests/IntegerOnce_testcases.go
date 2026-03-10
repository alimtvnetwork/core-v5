package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IntegerOnce — Core (Value + String + comparisons)
// =============================================================================

type integerOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue int
}

var integerOnceCoreTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 0 — IsZero, IsEmpty, String '0'",
			ExpectedInput: args.Map{
				"value":          0,
				"string":         "0",
				"isZero":         true,
				"isEmpty":        true,
				"isAboveZero":    false,
				"isPositive":     false,
				"isLessThanZero": false,
				"isNegative":     false,
			},
		},
		InitValue: 0,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 42 — positive, not zero",
			ExpectedInput: args.Map{
				"value":          42,
				"string":         "42",
				"isZero":         false,
				"isEmpty":        false,
				"isAboveZero":    true,
				"isPositive":     true,
				"isLessThanZero": false,
				"isNegative":     false,
			},
		},
		InitValue: 42,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce -3 — negative, not zero",
			ExpectedInput: args.Map{
				"value":          -3,
				"string":         "-3",
				"isZero":         false,
				"isEmpty":        false,
				"isAboveZero":    false,
				"isPositive":     false,
				"isLessThanZero": true,
				"isNegative":     true,
			},
		},
		InitValue: -3,
	},
}

// =============================================================================
// IntegerOnce — Caching
// =============================================================================

var integerOnceCachingTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce.Value caches — initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        42,
				"r2":        42,
				"callCount": 1,
			},
		},
		InitValue: 42,
	},
}

// =============================================================================
// IntegerOnce — Comparisons (IsAbove, IsLessThan)
// =============================================================================

type integerOnceCompareTestCase struct {
	Case         coretestcases.CaseV1
	InitValue    int
	CompareValue int
}

var integerOnceCompareTestCases = []integerOnceCompareTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 10 — IsAbove(5), IsAboveEqual(10)",
			ExpectedInput: args.Map{
				"isAboveCompare":   true,
				"isAboveSelf":      false,
				"isAboveEqualSelf": true,
			},
		},
		InitValue:    10,
		CompareValue: 5,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 3 — IsLessThan(5), IsLessThanEqual(3)",
			ExpectedInput: args.Map{
				"isLessThanCompare":   true,
				"isLessThanSelf":      false,
				"isLessThanEqualSelf": true,
			},
		},
		InitValue:    3,
		CompareValue: 5,
	},
}

// =============================================================================
// IntegerOnce — JSON
// =============================================================================

var integerOnceJsonTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 42 — MarshalJSON returns '42'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "42",
			},
		},
		InitValue: 42,
	},
}
