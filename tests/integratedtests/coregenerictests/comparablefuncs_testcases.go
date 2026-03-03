package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// ContainsAll
// ==========================================

var containsAllTrueTestCase = coretestcases.CaseV1{
	Title: "ContainsAll true when all items present",
	ArrangeInput: args.Map{
		"when":        "given collection containing all search items",
		"items":       []int{1, 2, 3, 4, 5},
		"searchItems": []int{1, 3, 5},
	},
	ExpectedInput: []string{
		"true",
	},
}

var containsAllFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAll false when item missing",
	ArrangeInput: args.Map{
		"when":        "given collection missing one search item",
		"items":       []int{1, 2, 3},
		"searchItems": []int{1, 2, 99},
	},
	ExpectedInput: []string{
		"false",
	},
}

// ==========================================
// ContainsAny
// ==========================================

var containsAnyTrueTestCase = coretestcases.CaseV1{
	Title: "ContainsAny true when at least one present",
	ArrangeInput: args.Map{
		"when":        "given collection with one matching item",
		"items":       []int{1, 2, 3},
		"searchItems": []int{99, 3, 100},
	},
	ExpectedInput: []string{
		"true",
	},
}

var containsAnyFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAny false when none present",
	ArrangeInput: args.Map{
		"when":        "given collection with no matching items",
		"items":       []int{1, 2, 3},
		"searchItems": []int{88, 99, 100},
	},
	ExpectedInput: []string{
		"false",
	},
}

// ==========================================
// RemoveItem
// ==========================================

var removeItemFoundTestCase = coretestcases.CaseV1{
	Title: "RemoveItem removes first occurrence",
	ArrangeInput: args.Map{
		"when":       "given collection with duplicates, remove first 2",
		"items":      []int{1, 2, 3, 2, 4},
		"removeItem": 2,
	},
	ExpectedInput: []string{
		"true",
		"4",
	},
}

var removeItemMissingTestCase = coretestcases.CaseV1{
	Title: "RemoveItem returns false for missing item",
	ArrangeInput: args.Map{
		"when":       "given collection without target item",
		"items":      []int{1, 3, 5},
		"removeItem": 99,
	},
	ExpectedInput: []string{
		"false",
		"3",
	},
}

// ==========================================
// RemoveAllItems
// ==========================================

var removeAllItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveAllItems removes all occurrences",
		ArrangeInput: args.Map{
			"when":  "given collection with multiple 2s",
			"items": []int{1, 2, 3, 2, 4, 2},
		},
		ExpectedInput: []string{
			"3",
			"3",
		},
	},
}

// ==========================================
// ToHashset
// ==========================================

var toHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "ToHashset converts collection to hashset with unique items",
		ArrangeInput: args.Map{
			"when":  "given collection with duplicates",
			"items": []int{1, 2, 3, 2, 1},
		},
		ExpectedInput: []string{
			"3",
			"true",
			"true",
			"true",
			"false",
		},
	},
}

// ==========================================
// DistinctSimpleSlice
// ==========================================

var distinctSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "DistinctSimpleSlice removes duplicates preserving order",
		ArrangeInput: args.Map{
			"when":  "given simple slice with duplicates",
			"items": []int{3, 1, 2, 1, 3, 4},
		},
		ExpectedInput: []string{
			"4",
			"3",
			"4",
		},
	},
}

// ==========================================
// ContainsSimpleSliceItem
// ==========================================

var containsSimpleSliceItemTestCases = []coretestcases.CaseV1{
	{
		Title: "ContainsSimpleSliceItem true for existing item",
		ArrangeInput: args.Map{
			"when":  "given simple slice containing target",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}
