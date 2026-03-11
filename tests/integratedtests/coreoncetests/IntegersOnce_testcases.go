package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IntegersOnce -- Core
// =============================================================================

type integersOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue []int
}

var integersOnceCoreTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns length 3 and isEmpty false -- [3,1,2] input",
			ExpectedInput: args.Map{
				"length":  3,
				"isEmpty": false,
				"isZero":  false,
			},
		},
		InitValue: []int{3, 1, 2},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns length 0 and isEmpty true -- empty input",
			ExpectedInput: args.Map{
				"length":  0,
				"isEmpty": true,
				"isZero":  true,
			},
		},
		InitValue: []int{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce nil -- Length 0, isEmpty true",
			ExpectedInput: args.Map{
				"length":  0,
				"isEmpty": true,
				"isZero":  true,
			},
		},
		InitValue: nil,
	},
}

// =============================================================================
// IntegersOnce -- Sorted
// =============================================================================

var integersOnceSortedTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce [3,1,2] -- Sorted returns [1,2,3]",
			ExpectedInput: args.Map{
				"first": 1,
				"last":  3,
			},
		},
		InitValue: []int{3, 1, 2},
	},
}

// =============================================================================
// IntegersOnce -- RangesMap / RangesBoolMap
// =============================================================================

var integersOnceRangesTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce [10,20,30] -- RangesMap and RangesBoolMap",
			ExpectedInput: args.Map{
				"rangesMapLength":    3,
				"rangesBoolMapLen":   3,
				"uniqueMapLen":       3,
				"rangesMapHas10":     true,
				"rangesBoolMapHas20": true,
			},
		},
		InitValue: []int{10, 20, 30},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce empty -- RangesMap empty",
			ExpectedInput: args.Map{
				"rangesMapLength":    0,
				"rangesBoolMapLen":   0,
				"uniqueMapLen":       0,
				"rangesMapHas10":     false,
				"rangesBoolMapHas20": false,
			},
		},
		InitValue: []int{},
	},
}

// =============================================================================
// IntegersOnce -- IsEqual
// =============================================================================

var integersOnceIsEqualTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce [1,2,3] -- IsEqual same true, different false",
			ExpectedInput: args.Map{
				"isEqualSame":    true,
				"isEqualDiff":    false,
				"isEqualDiffLen": false,
			},
		},
		InitValue: []int{1, 2, 3},
	},
}

// =============================================================================
// IntegersOnce -- Caching
// =============================================================================

var integersOnceCachingTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce.Value caches -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
				"length":    2,
			},
		},
		InitValue: []int{5, 10},
	},
}

// =============================================================================
// IntegersOnce -- JSON
// =============================================================================

var integersOnceJsonTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce [1,2] -- MarshalJSON returns '[1,2]'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "[1,2]",
			},
		},
		InitValue: []int{1, 2},
	},
}

// =============================================================================
// IntegersOnce -- Constructor
// =============================================================================

var integersOnceConstructorTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewIntegersOnce (value) works correctly",
			ExpectedInput: args.Map{
				"length": 3,
			},
		},
		InitValue: []int{1, 2, 3},
	},
}
