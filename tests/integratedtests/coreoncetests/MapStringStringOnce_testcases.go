package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// MapStringStringOnce — Core
// =============================================================================

type mapSSOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue map[string]string
}

var mapSSOnceCoreTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce {a:1,b:2} — Length 2, not empty",
			ExpectedInput: args.Map{
				"length":     2,
				"isEmpty":    false,
				"hasAnyItem": true,
			},
		},
		InitValue: map[string]string{"a": "1", "b": "2"},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce empty — Length 0, isEmpty true",
			ExpectedInput: args.Map{
				"length":     0,
				"isEmpty":    true,
				"hasAnyItem": false,
			},
		},
		InitValue: map[string]string{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce nil — Length 0, isEmpty true",
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
// MapStringStringOnce — Lookup (Has, IsContains, IsMissing, GetValue)
// =============================================================================

var mapSSOnceContainsTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce {k1:v1,k2:v2} — Has/IsContains/GetValue",
			ExpectedInput: args.Map{
				"hasK1":      true,
				"containsK2": true,
				"isMissingX": true,
				"hasAllK1K2": true,
				"getK1":      "v1",
			},
		},
		InitValue: map[string]string{"k1": "v1", "k2": "v2"},
	},
}

// =============================================================================
// MapStringStringOnce — Keys / Values / Sorted
// =============================================================================

var mapSSOnceKeysValuesTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce {b:2,a:1} — AllKeysSorted [a,b], AllValuesSorted [1,2]",
			ExpectedInput: args.Map{
				"keysLen":          2,
				"valuesLen":        2,
				"sortedFirstKey":   "a",
				"sortedLastKey":    "b",
				"sortedFirstValue": "1",
				"sortedLastValue":  "2",
			},
		},
		InitValue: map[string]string{"b": "2", "a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce — IsEqual
// =============================================================================

var mapSSOnceIsEqualTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce {a:1} — IsEqual same true, different false",
			ExpectedInput: args.Map{
				"isEqualSame":    true,
				"isEqualDiffVal": false,
				"isEqualDiffKey": false,
				"isEqualDiffLen": false,
			},
		},
		InitValue: map[string]string{"a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce — Caching
// =============================================================================

var mapSSOnceCachingTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.Value caches — initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
				"length":    2,
			},
		},
		InitValue: map[string]string{"x": "1", "y": "2"},
	},
}

// =============================================================================
// MapStringStringOnce — JSON
// =============================================================================

var mapSSOnceJsonTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce {a:1} — MarshalJSON succeeds",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitValue: map[string]string{"a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce — Constructor
// =============================================================================

var mapSSOnceConstructorTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewMapStringStringOnce (value) works correctly",
			ExpectedInput: args.Map{
				"length": 1,
			},
		},
		InitValue: map[string]string{"k": "v"},
	},
}
