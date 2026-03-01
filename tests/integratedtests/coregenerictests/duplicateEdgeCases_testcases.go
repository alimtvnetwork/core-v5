package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		ExpectedInput: []string{
			"1",
			"7",
			"7",
		},
	},
	{
		Title: "Distinct on single-element collection returns same",
		ArrangeInput: args.Map{
			"when":  "given collection with one element",
			"items": []int{42},
		},
		ExpectedInput: []string{
			"1",
			"42",
			"42",
		},
	},
	{
		Title: "Distinct on empty collection returns empty",
		ArrangeInput: args.Map{
			"when":  "given empty collection",
			"items": []int{},
		},
		ExpectedInput: []string{
			"0",
			"true",
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
		ExpectedInput: []string{
			"true",
			"4",
			"3",
			"3",
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
		ExpectedInput: []string{
			"5",
			"0",
			"true",
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
		ExpectedInput: []string{
			"true",
			"false",
			"true",
			"false",
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
		ExpectedInput: []string{
			"1",
			"true",
			"false",
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
		ExpectedInput: []string{
			"1",
			"true",
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
		ExpectedInput: []string{
			"true",
			"false",
			"false",
			"false",
			"1",
		},
	},
}

// ==========================================
// SimpleSlice: DistinctSimpleSlice all same
// ==========================================

var distinctSimpleSliceAllSameTestCases = []coretestcases.CaseV1{
	{
		Title: "DistinctSimpleSlice on all-same returns single element",
		ArrangeInput: args.Map{
			"when":  "given simple slice of five 8s",
			"items": []int{8, 8, 8, 8, 8},
		},
		ExpectedInput: []string{
			"1",
			"8",
		},
	},
	{
		Title: "DistinctSimpleSlice on empty returns empty",
		ArrangeInput: args.Map{
			"when":  "given empty simple slice",
			"items": []int{},
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}
