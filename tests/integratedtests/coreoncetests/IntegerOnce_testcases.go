package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IntegerOnce — Core (Value + String + comparisons)
// =============================================================================

type integerOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue int
}

// Note: 8 fields — exceeds args.Six, kept as []string for accuracy.
var integerOnceCoreTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 0 — IsZero, IsEmpty, String '0'",
			ExpectedInput: []string{
				"0",     // Value
				"0",     // String
				"true",  // IsZero
				"true",  // IsEmpty
				"false", // IsAboveZero
				"false", // IsPositive
				"false", // IsLessThanZero
				"false", // IsNegative
			},
		},
		InitValue: 0,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 42 — positive, not zero",
			ExpectedInput: []string{
				"42",    // Value
				"42",    // String
				"false", // IsZero
				"false", // IsEmpty
				"true",  // IsAboveZero
				"true",  // IsPositive
				"false", // IsLessThanZero
				"false", // IsNegative
			},
		},
		InitValue: 42,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce -3 — negative, not zero",
			ExpectedInput: []string{
				"-3",    // Value
				"-3",    // String
				"false", // IsZero
				"false", // IsEmpty
				"false", // IsAboveZero
				"false", // IsPositive
				"true",  // IsLessThanZero
				"true",  // IsNegative
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
			ExpectedInput: args.Three[string, string, string]{
				First:  "42", // r1
				Second: "42", // r2
				Third:  "1",  // callCount
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
			ExpectedInput: args.Three[string, string, string]{
				First:  "true",  // isAbove5
				Second: "false", // isAbove10
				Third:  "true",  // isAboveEqual10
			},
		},
		InitValue:    10,
		CompareValue: 5,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce 3 — IsLessThan(5), IsLessThanEqual(3)",
			ExpectedInput: args.Three[string, string, string]{
				First:  "true",  // isLessThan5
				Second: "false", // isLessThan3
				Third:  "true",  // isLessThanEqual3
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
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "42",   // marshaledValue
			},
		},
		InitValue: 42,
	},
}
