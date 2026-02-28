package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Generic AddLock
// ==========================================

var genericAddLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic AddLock appends item thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddLock calls on generic collection",
			"count": 100,
		},
		ExpectedInput: []string{
			"100",
		},
	},
}

// ==========================================
// Generic AddsLock
// ==========================================

var genericAddsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic AddsLock appends multiple items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddsLock calls on generic collection",
			"count": 50,
			"batch": 3,
		},
		ExpectedInput: []string{
			"150",
		},
	},
}

// ==========================================
// Generic LengthLock
// ==========================================

var genericLengthLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic LengthLock returns correct count",
		ArrangeInput: args.Map{
			"when":  "given generic collection with items",
			"items": []any{"x", 42, true},
		},
		ExpectedInput: []string{
			"3",
		},
	},
}

// ==========================================
// Generic IsEmptyLock — empty
// ==========================================

var genericIsEmptyLockEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic IsEmptyLock returns true for empty collection",
		ArrangeInput: args.Map{
			"when": "given empty generic collection",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// Generic IsEmptyLock — non-empty
// ==========================================

var genericIsEmptyLockNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic IsEmptyLock returns false for non-empty collection",
		ArrangeInput: args.Map{
			"when": "given non-empty generic collection",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}

// ==========================================
// Generic ItemsLock
// ==========================================

var genericItemsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic ItemsLock returns independent copy",
		ArrangeInput: args.Map{
			"when":  "given generic collection with items",
			"items": []any{"a", "b"},
		},
		ExpectedInput: []string{
			"2",
			"a",
			"b",
			"true",
		},
	},
}

// ==========================================
// Generic ClearLock
// ==========================================

var genericClearLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic ClearLock removes all items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given generic collection then ClearLock",
			"items": []any{"x", "y", "z"},
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// Generic AddCollectionLock
// ==========================================

var genericAddCollectionLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic AddCollectionLock merges thread-safely",
		ArrangeInput: args.Map{
			"when":   "given two generic collections merged with lock",
			"first":  []any{"a"},
			"second": []any{"b", "c"},
		},
		ExpectedInput: []string{
			"3",
			"a",
			"c",
		},
	},
}

// ==========================================
// Generic FilterLock
// ==========================================

var genericFilterLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic FilterLock filters items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent reads while filtering generic collection",
			"items": []any{"alpha", "beta", "gamma", "delta"},
		},
		ExpectedInput: []string{
			"2",
			"alpha",
			"delta",
		},
	},
}

// ==========================================
// Generic LoopLock
// ==========================================

var genericLoopLockTestCases = []coretestcases.CaseV1{
	{
		Title: "Generic LoopLock iterates over snapshot thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent writes during LoopLock on generic collection",
			"count": 50,
		},
		ExpectedInput: []string{
			"50",
		},
	},
}
