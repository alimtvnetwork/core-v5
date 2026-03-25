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
			Title: "AnyOnce.Value returns string and IsNull false -- 'hello' input",
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
			Title: "AnyOnce.Value returns IsNull true and String empty -- nil input",
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
			Title: "AnyOnce.Value returns int and IsNull false -- 42 input",
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
		Title: "AnyOnce.CastValueString returns 'cast-me' and castSuccess true -- string input",
		ExpectedInput: args.Map{
			"castValue":   "cast-me",
			"castSuccess": true,
		},
	},
	InitValue: "cast-me",
}

var anyOnceCastStringsTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce.CastValueStrings returns length 2 and castSuccess true -- []string input",
		ExpectedInput: args.Map{
			"castLen":     2,
			"castSuccess": true,
		},
	},
	InitValue: []string{"a", "b"},
}

var anyOnceCastBytesTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce.CastValueBytes returns length 5 and castSuccess true -- []byte input",
		ExpectedInput: args.Map{
			"castLen":     5,
			"castSuccess": true,
		},
	},
	InitValue: []byte("bytes"),
}

var anyOnceCastMapTestCase = anyOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "AnyOnce.CastValueHashmapMap returns length 1 and castSuccess true -- map input",
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
			Title: "AnyOnce.Value returns cached result -- initializer runs once",
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
			Title: "AnyOnce.Serialize returns bytes without error -- 'json' value",
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
			Title: "NewAnyOnce returns isNull false -- value input",
			ExpectedInput: args.Map{
				"isNull": false,
			},
		},
		InitValue: "val",
	},
}
