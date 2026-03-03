package coreoncetests

import (
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
			ExpectedInput: []string{
				"true",
				"true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false — Value returns false, String returns 'false'",
			ExpectedInput: []string{
				"false",
				"false",
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
			ExpectedInput: []string{
				"true",  // r1
				"true",  // r2
				"true",  // r3
				"1",     // callCount
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value false — caches false, initializer runs once",
			ExpectedInput: []string{
				"false",
				"false",
				"false",
				"1",
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
			ExpectedInput: []string{
				"true", // no error
				"true", // marshaled value
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce false — MarshalJSON returns 'false'",
			ExpectedInput: []string{
				"true",  // no error
				"false", // marshaled value
			},
		},
		InitValue: false,
	},
}
