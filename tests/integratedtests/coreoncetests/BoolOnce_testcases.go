package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// BoolOnce — Core (Value + String)
// =============================================================================

type boolOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue bool
}

var boolOnceCoreTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce true — Value returns true, String returns 'true'",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // value
				Second: "true", // stringRepresentation
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false — Value returns false, String returns 'false'",
			ExpectedInput: args.Two[string, string]{
				First:  "false", // value
				Second: "false", // stringRepresentation
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce — Caching (call count verification)
// =============================================================================

var boolOnceCachingTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value caches — initializer runs exactly once",
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "true", // result1
				Second: "true", // result2
				Third:  "true", // result3
				Fourth: "1",    // callCount
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value false — caches false, initializer runs once",
			ExpectedInput: args.Four[string, string, string, string]{
				First:  "false", // result1
				Second: "false", // result2
				Third:  "false", // result3
				Fourth: "1",     // callCount
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce — JSON
// =============================================================================

var boolOnceJsonTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce true — MarshalJSON returns 'true'",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "true", // marshaledValue
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false — MarshalJSON returns 'false'",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // noError
				Second: "false", // marshaledValue
			},
		},
		InitValue: false,
	},
}
