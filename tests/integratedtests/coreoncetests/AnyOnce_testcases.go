package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// AnyOnce -- Core
// =============================================================================

type anyOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue any
}

var anyOnceCoreTestCases = []anyOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyOnce 'hello' -- Value returns string, IsNull false",
			ExpectedInput: args.Map{
				"isNull":                    false,
				"isStringEmpty":             false,
				"isStringEmptyOrWhitespace": false,
				"isInitialized":             true,
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyOnce nil -- IsNull true, String empty",
			ExpectedInput: args.Map{
				"isNull":                    true,
				"isStringEmpty":             true,
				"isStringEmptyOrWhitespace": true,
				"isInitialized":             true,
			},
		},
		InitValue: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "AnyOnce 42 -- Value returns int, not null",
			ExpectedInput: args.Map{
				"isNull":                    false,
				"isStringEmpty":             false,
				"isStringEmptyOrWhitespace": false,
				"isInitialized":             true,
			},
		},
		InitValue: 42,
	},
}

// =============================================================================
// AnyOnce -- Cast methods
// =============================================================================

var anyOnceCastStringTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce 'cast-me' -- CastValueString succeeds",
		ExpectedInput: args.Map{
			"castValue":   "cast-me",
			"castSuccess": true,
		},
	},
	InitValue: "cast-me",
}

var anyOnceCastStringsTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce []string -- CastValueStrings succeeds",
		ExpectedInput: args.Map{
			"castLen":     2,
			"castSuccess": true,
		},
	},
	InitValue: []string{"a", "b"},
}

var anyOnceCastBytesTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce []byte -- CastValueBytes succeeds",
		ExpectedInput: args.Map{
			"castLen":     5,
			"castSuccess": true,
		},
	},
	InitValue: []byte("bytes"),
}

var anyOnceCastMapTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce map[string]string -- CastValueHashmapMap succeeds",
		ExpectedInput: args.Map{
			"castLen":     1,
			"castSuccess": true,
		},
	},
	InitValue: map[string]string{"k": "v"},
}

// =============================================================================
// AnyOnce -- Caching
// =============================================================================

var anyOnceCachingTestCases = []anyOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyOnce.Value caches -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
			},
		},
		InitValue: "cached",
	},
}

// =============================================================================
// AnyOnce -- JSON
// =============================================================================

var anyOnceJsonTestCases = []anyOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "AnyOnce 'json' -- Serialize succeeds",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitValue: "json",
	},
}

// =============================================================================
// AnyOnce -- Constructor
// =============================================================================

var anyOnceConstructorTestCases = []anyOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewAnyOnce (value) works correctly",
			ExpectedInput: args.Map{
				"isNull": false,
			},
		},
		InitValue: "val",
	},
}
