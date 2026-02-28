package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// SortCollection — ascending
// ==========================================

var sortCollectionAscTestCases = []coretestcases.CaseV1{
	{
		Title: "SortCollection sorts integers ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted int collection",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: []string{
			"5",
			"1",
			"5",
			"true",
		},
	},
	{
		Title: "SortCollection on already sorted is no-op",
		ArrangeInput: args.Map{
			"when":  "given already sorted collection",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"3",
			"1",
			"3",
			"true",
		},
	},
	{
		Title: "SortCollection single element",
		ArrangeInput: args.Map{
			"when":  "given single-element collection",
			"items": []int{42},
		},
		ExpectedInput: []string{
			"1",
			"42",
			"42",
			"true",
		},
	},
}

// ==========================================
// SortCollectionDesc — descending
// ==========================================

var sortCollectionDescTestCases = []coretestcases.CaseV1{
	{
		Title: "SortCollectionDesc sorts integers descending",
		ArrangeInput: args.Map{
			"when":  "given unsorted int collection",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: []string{
			"5",
			"5",
			"1",
		},
	},
}

// ==========================================
// MinCollection / MaxCollection
// ==========================================

var minMaxCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "MinCollection and MaxCollection return correct values",
		ArrangeInput: args.Map{
			"when":  "given int collection with various values",
			"items": []int{7, 2, 9, 1, 5},
		},
		ExpectedInput: []string{
			"1",
			"9",
		},
	},
	{
		Title: "MinCollection and MaxCollection on single element",
		ArrangeInput: args.Map{
			"when":  "given single-element collection",
			"items": []int{42},
		},
		ExpectedInput: []string{
			"42",
			"42",
		},
	},
}

// ==========================================
// MinCollectionOrDefault / MaxCollectionOrDefault
// ==========================================

var minMaxCollectionOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "OrDefault returns values for non-empty collection",
		ArrangeInput: args.Map{
			"when":  "given non-empty int collection",
			"items": []int{3, 1, 4},
		},
		ExpectedInput: []string{
			"1",
			"4",
		},
	},
}

var minMaxCollectionOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "OrDefault returns default for empty collection",
		ArrangeInput: args.Map{
			"when": "given empty int collection with default -1",
		},
		ExpectedInput: []string{
			"-1",
			"-1",
		},
	},
}

// ==========================================
// IsSortedCollection
// ==========================================

var isSortedCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSortedCollection true for sorted",
		ArrangeInput: args.Map{
			"when":  "given ascending sorted collection",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsSortedCollection false for unsorted",
		ArrangeInput: args.Map{
			"when":  "given unsorted collection",
			"items": []int{3, 1, 2},
		},
		ExpectedInput: []string{
			"false",
		},
	},
}

// ==========================================
// SumCollection
// ==========================================

var sumCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "SumCollection returns correct sum",
		ArrangeInput: args.Map{
			"when":  "given int collection",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"15",
		},
	},
	{
		Title: "SumCollection empty returns zero",
		ArrangeInput: args.Map{
			"when":  "given empty int collection",
			"items": []int{},
		},
		ExpectedInput: []string{
			"0",
		},
	},
}

// ==========================================
// ClampCollection
// ==========================================

var clampCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "ClampCollection clamps values to range",
		ArrangeInput: args.Map{
			"when":  "given ints clamped to [2, 4]",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"2",
			"2",
			"3",
			"4",
			"4",
		},
	},
}

// ==========================================
// Hashset ordered: SortedListHashset
// ==========================================

var sortedListHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedListHashset returns sorted items",
		ArrangeInput: args.Map{
			"when":  "given int hashset with unordered items",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: []string{
			"5",
			"1",
			"5",
		},
	},
}

// ==========================================
// Hashset ordered: MinHashset / MaxHashset
// ==========================================

var minMaxHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashset and MaxHashset return correct values",
		ArrangeInput: args.Map{
			"when":  "given int hashset",
			"items": []int{7, 2, 9, 1, 5},
		},
		ExpectedInput: []string{
			"1",
			"9",
		},
	},
}

var minMaxHashsetOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashsetOrDefault returns default on empty",
		ArrangeInput: args.Map{
			"when": "given empty int hashset with default -1",
		},
		ExpectedInput: []string{
			"-1",
			"-1",
		},
	},
}

// ==========================================
// Hashmap ordered: SortedKeysHashmap
// ==========================================

var sortedKeysHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedKeysHashmap returns keys in ascending order",
		ArrangeInput: args.Map{
			"when": "given string-int hashmap with unordered keys",
		},
		ExpectedInput: []string{
			"3",
			"alpha",
			"gamma",
		},
	},
}

// ==========================================
// Hashmap ordered: MinKeyHashmap / MaxKeyHashmap
// ==========================================

var minMaxKeyHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "MinKeyHashmap and MaxKeyHashmap return correct keys",
		ArrangeInput: args.Map{
			"when": "given string-int hashmap",
		},
		ExpectedInput: []string{
			"alpha",
			"gamma",
		},
	},
}

// ==========================================
// Hashmap ordered: MinValueHashmap / MaxValueHashmap
// ==========================================

var minMaxValueHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "MinValueHashmap and MaxValueHashmap return correct values",
		ArrangeInput: args.Map{
			"when": "given string-int hashmap with numeric values",
		},
		ExpectedInput: []string{
			"1",
			"30",
		},
	},
}

// ==========================================
// SimpleSlice ordered: SortSimpleSlice
// ==========================================

var sortSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "SortSimpleSlice sorts ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted int simple slice",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: []string{
			"5",
			"1",
			"5",
		},
	},
}

var minMaxSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "MinSimpleSlice and MaxSimpleSlice return correct values",
		ArrangeInput: args.Map{
			"when":  "given int simple slice",
			"items": []int{7, 2, 9, 1, 5},
		},
		ExpectedInput: []string{
			"1",
			"9",
		},
	},
}
