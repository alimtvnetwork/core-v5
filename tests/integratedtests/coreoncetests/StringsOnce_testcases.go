package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// StringsOnce -- Core
// =============================================================================

type stringsOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue []string
}

var stringsOnceCoreTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce returns length 3 and isEmpty false -- [a,b,c] input",
			ExpectedInput: args.Map{
				"length":     3,
				"isEmpty":    false,
				"hasAnyItem": true,
			},
		},
		InitValue: []string{"a", "b", "c"},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce returns length 0 and isEmpty true -- empty input",
			ExpectedInput: args.Map{
				"length":     0,
				"isEmpty":    true,
				"hasAnyItem": false,
			},
		},
		InitValue: []string{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce returns length 0 and isEmpty true -- nil input",
			ExpectedInput: args.Map{
				"length":     0,
				"isEmpty":    true,
				"hasAnyItem": false,
			},
		},
		InitValue: nil,
	},
}

// =============================================================================
// StringsOnce -- Contains / Has / HasAll
// =============================================================================

var stringsOnceContainsTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce returns hasX true and hasAllXY true -- [x,y,z] input",
			ExpectedInput: args.Map{
				"hasX":       true,
				"containsY":  true,
				"hasMissing": false,
				"hasAllXY":   true,
				"hasAllXW":   false,
			},
		},
		InitValue: []string{"x", "y", "z"},
	},
}

// =============================================================================
// StringsOnce -- Sorted
// =============================================================================

var stringsOnceSortedTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce.Sorted returns [a,b,c] -- [c,a,b] input",
			ExpectedInput: args.Map{
				"first": "a",
				"last":  "c",
			},
		},
		InitValue: []string{"c", "a", "b"},
	},
}

// =============================================================================
// StringsOnce -- UniqueMap / RangesMap
// =============================================================================

var stringsOnceMapTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce returns uniqueLen 2 and rangesMapLen 2 -- [a,b,a] input",
			ExpectedInput: args.Map{
				"uniqueLen":    2,
				"rangesMapLen": 2,
				"uniqueHasA":   true,
				"uniqueHasB":   true,
			},
		},
		InitValue: []string{"a", "b", "a"},
	},
}

// =============================================================================
// StringsOnce -- IsEqual
// =============================================================================

var stringsOnceIsEqualTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce [a,b] -- IsEqual same true, different false",
			ExpectedInput: args.Map{
				"isEqualSame":    true,
				"isEqualDiff":    false,
				"isEqualDiffLen": false,
			},
		},
		InitValue: []string{"a", "b"},
	},
}

// =============================================================================
// StringsOnce -- Caching
// =============================================================================

var stringsOnceCachingTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce.Value caches -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
				"length":    2,
			},
		},
		InitValue: []string{"hello", "world"},
	},
}

// =============================================================================
// StringsOnce -- JSON
// =============================================================================

var stringsOnceJsonTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringsOnce [a,b] -- MarshalJSON returns '[\"a\",\"b\"]'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "[\"a\",\"b\"]",
			},
		},
		InitValue: []string{"a", "b"},
	},
}

// =============================================================================
// StringsOnce -- Constructor
// =============================================================================

var stringsOnceConstructorTestCases = []stringsOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewStringsOnce (value) works correctly",
			ExpectedInput: args.Map{
				"length": 2,
			},
		},
		InitValue: []string{"x", "y"},
	},
}
