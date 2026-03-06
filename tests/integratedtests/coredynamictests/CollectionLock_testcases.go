package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// AddLock
// ==========================================

var collectionAddLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddLock appends item thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddLock calls",
			"count": 100,
		},
		ExpectedInput: "100",
	},
}

// ==========================================
// AddsLock
// ==========================================

var collectionAddsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddsLock appends multiple items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddsLock calls",
			"count": 50,
			"batch": 2,
		},
		ExpectedInput: "100",
	},
}

// ==========================================
// LengthLock
// ==========================================

var collectionLengthLockTestCases = []coretestcases.CaseV1{
	{
		Title: "LengthLock returns correct count under concurrency",
		ArrangeInput: args.Map{
			"when":  "given items added then LengthLock called",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "3",
	},
}

// ==========================================
// IsEmptyLock
// ==========================================

var collectionIsEmptyLockTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyLock returns true for empty collection",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: "true",
	},
}

var collectionIsEmptyLockNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyLock returns false for non-empty collection",
		ArrangeInput: args.Map{
			"when": "given non-empty collection",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// ItemsLock
// ==========================================

var collectionItemsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "ItemsLock returns independent copy",
		ArrangeInput: args.Map{
			"when":  "given collection with items",
			"items": []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length":        2,
			"first":         "x",
			"last":          "y",
			"isIndependent": true,
		},
	},
}

// ==========================================
// ClearLock
// ==========================================

var collectionClearLockTestCases = []coretestcases.CaseV1{
	{
		Title: "ClearLock removes all items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given collection then ClearLock",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// AddCollectionLock
// ==========================================

var collectionAddCollectionLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollectionLock merges thread-safely",
		ArrangeInput: args.Map{
			"when":   "given two collections merged with lock",
			"first":  []string{"a"},
			"second": []string{"b", "c"},
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "a",
			"last":   "c",
		},
	},
}
