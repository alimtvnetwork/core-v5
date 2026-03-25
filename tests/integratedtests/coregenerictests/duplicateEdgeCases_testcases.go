package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// Collection: All same value — Distinct
// ==========================================

var distinctAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "Distinct on all-same-value collection returns single element",
		ArrangeInput: args.Map{
			"when":  "given collection where every element is 7",
			"items": []int{7, 7, 7, 7, 7},
		},
		ExpectedInput: args.Map{
			"length":         1,
			"isEmpty":        false,
			"firstOrDefault": 7,
			"lastOrDefault":  7,
		},
	},
	{
		Title: "Distinct on single-element collection returns same",
		ArrangeInput: args.Map{
			"when":  "given collection with one element",
			"items": []int{42},
		},
		ExpectedInput: args.Map{
			"length":         1,
			"isEmpty":        false,
			"firstOrDefault": 42,
			"lastOrDefault":  42,
		},
	},
	{
		Title: "Distinct on empty collection returns empty",
		ArrangeInput: args.Map{
			"when":  "given empty collection",
			"items": []int{},
		},
		ExpectedInput: args.Map{
			"length":         0,
			"isEmpty":        true,
			"firstOrDefault": 0,
			"lastOrDefault":  0,
		},
	},
}

// ==========================================
// Collection: All same value — RemoveItem
// ==========================================

var removeItemAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveItem on all-same removes only first occurrence",
		ArrangeInput: args.Map{
			"when":  "given collection of five 3s, remove 3",
			"items": []int{3, 3, 3, 3, 3},
		},
		ExpectedInput: args.Map{
			"removed": true,
			"length":  4,
			"first":   3,
			"last":    3,
		},
	},
}

// ==========================================
// Collection: All same value — RemoveAllItems
// ==========================================

var removeAllItemsAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveAllItems on all-same empties the collection",
		ArrangeInput: args.Map{
			"when":  "given collection of five 3s, remove all 3s",
			"items": []int{3, 3, 3, 3, 3},
		},
		ExpectedInput: args.Map{
			"removedCount": 5,
			"length":       0,
			"isEmpty":      true,
		},
	},
}

// ==========================================
// Collection: All same value — ContainsAll / ContainsAny
// ==========================================

var containsAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "ContainsAll on all-same: true for same value, false for different",
		ArrangeInput: args.Map{
			"when":  "given collection where every element is 5",
			"items": []int{5, 5, 5},
		},
		ExpectedInput: args.Map{
			"containsAllSingle":    true,
			"containsAllWithOther": false,
			"containsAnyWithMatch": true,
			"containsAnyNoMatch":   false,
		},
	},
}

// ==========================================
// Collection: All same value — ToHashset
// ==========================================

var toHashsetAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "ToHashset on all-same-value collection yields single-element set",
		ArrangeInput: args.Map{
			"when":  "given collection of five 9s",
			"items": []int{9, 9, 9, 9, 9},
		},
		ExpectedInput: args.Map{
			"length": 1,
			"has9":   true,
			"has99":  false,
		},
	},
}

// ==========================================
// Hashset: Add duplicates
// ==========================================

var hashsetAddDuplicatesTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset.From with all same values yields single element",
		ArrangeInput: args.Map{
			"when":  "given slice of repeated 'x' values",
			"items": []string{"x", "x", "x", "x"},
		},
		ExpectedInput: args.Map{
			"length": 1,
			"hasX":   true,
		},
	},
}

// ==========================================
// Hashset: AddBool with repeated adds
// ==========================================

var hashsetAddBoolDuplicatesTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset.AddBool returns false for all repeated adds after first",
		ArrangeInput: args.Map{
			"when": "adding same value 4 times",
		},
		ExpectedInput: args.Map{
			"add1":   false,
			"add2":   true,
			"add3":   true,
			"add4":   true,
			"length": 1,
		},
	},
}

// ==========================================
// SimpleSlice: DistinctSimpleSlice all same
// ==========================================

var distinctSimpleSliceAllSameNonEmptyTestCase = coretestcases.CaseV1{
	Title: "DistinctSimpleSlice on all-same returns single element",
	ArrangeInput: args.Map{
		"when":  "given simple slice of five 8s",
		"items": []int{8, 8, 8, 8, 8},
	},
	ExpectedInput: args.Map{
		"length": 1,
		"first":  8,
	},
}

var distinctSimpleSliceAllSameEmptyTestCase = coretestcases.CaseV1{
	Title: "DistinctSimpleSlice on empty returns empty",
	ArrangeInput: args.Map{
		"when":  "given empty simple slice",
		"items": []int{},
	},
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}
