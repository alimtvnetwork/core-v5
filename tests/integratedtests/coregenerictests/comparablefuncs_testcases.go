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
	ExpectedInput: "true",
}

var containsAllFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAll false when item missing",
	ArrangeInput: args.Map{
		"when":        "given collection missing one search item",
		"items":       []int{1, 2, 3},
		"searchItems": []int{1, 2, 99},
	},
	ExpectedInput: "false",
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
	ExpectedInput: "true",
}

var containsAnyFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAny false when none present",
	ArrangeInput: args.Map{
		"when":        "given collection with no matching items",
		"items":       []int{1, 2, 3},
		"searchItems": []int{88, 99, 100},
	},
	ExpectedInput: "false",
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
	ExpectedInput: args.Two[string, string]{
		First:  "true", // removed
		Second: "4",    // newLength
	},
}

var removeItemMissingTestCase = coretestcases.CaseV1{
	Title: "RemoveItem returns false for missing item",
	ArrangeInput: args.Map{
		"when":       "given collection without target item",
		"items":      []int{1, 3, 5},
		"removeItem": 99,
	},
	ExpectedInput: args.Two[string, string]{
		First:  "false", // removed
		Second: "3",     // newLength
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
		ExpectedInput: args.Two[string, string]{
			First:  "3", // removedCount
			Second: "3", // newLength
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",     // uniqueCount
			Second: "true",  // has1
			Third:  "true",  // has2
			Fourth: "true",  // has3
			Fifth:  "false", // has99
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "4", // uniqueCount
			Second: "3", // firstElement
			Third:  "4", // lastElement
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
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // containsExisting
			Second: "false", // containsMissing
		},
	},
}
