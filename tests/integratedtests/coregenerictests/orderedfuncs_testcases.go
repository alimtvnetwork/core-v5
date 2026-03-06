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
		ExpectedInput: args.Map{
			"length":   5,
			"first":    1,
			"last":     5,
			"isSorted": true,
		},
	},
	{
		Title: "SortCollection on already sorted is no-op",
		ArrangeInput: args.Map{
			"when":  "given already sorted collection",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: args.Map{
			"length":   3,
			"first":    1,
			"last":     3,
			"isSorted": true,
		},
	},
	{
		Title: "SortCollection single element",
		ArrangeInput: args.Map{
			"when":  "given single-element collection",
			"items": []int{42},
		},
		ExpectedInput: args.Map{
			"length":   1,
			"first":    42,
			"last":     42,
			"isSorted": true,
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
		ExpectedInput: args.Map{
			"length": 5,
			"first":  5,
			"last":   1,
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
		ExpectedInput: args.Map{
			"min": 1,
			"max": 9,
		},
	},
	{
		Title: "MinCollection and MaxCollection on single element",
		ArrangeInput: args.Map{
			"when":  "given single-element collection",
			"items": []int{42},
		},
		ExpectedInput: args.Map{
			"min": 42,
			"max": 42,
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
		ExpectedInput: args.Map{
			"min": 1,
			"max": 4,
		},
	},
}

var minMaxCollectionOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "OrDefault returns default for empty collection",
		ArrangeInput: args.Map{
			"when": "given empty int collection with default -1",
		},
		ExpectedInput: args.Map{
			"min": -1,
			"max": -1,
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
		ExpectedInput: args.Map{"isSorted": true},
	},
	{
		Title: "IsSortedCollection false for unsorted",
		ArrangeInput: args.Map{
			"when":  "given unsorted collection",
			"items": []int{3, 1, 2},
		},
		ExpectedInput: args.Map{"isSorted": false},
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
		ExpectedInput: args.Map{"sum": 15},
	},
	{
		Title: "SumCollection empty returns zero",
		ArrangeInput: args.Map{
			"when":  "given empty int collection",
			"items": []int{},
		},
		ExpectedInput: args.Map{"sum": 0},
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
		ExpectedInput: args.Map{
			"val0": 2,
			"val1": 2,
			"val2": 3,
			"val3": 4,
			"val4": 4,
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
		ExpectedInput: args.Map{
			"length": 5,
			"first":  1,
			"last":   5,
		},
	},
}

// ==========================================
// Hashset ordered: SortedListDescHashset
// ==========================================

var sortedListDescHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedListDescHashset returns items in descending order",
		ArrangeInput: args.Map{
			"when":  "given int hashset with unordered items",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: args.Map{
			"length": 5,
			"first":  5,
			"last":   1,
		},
	},
}

// ==========================================
// Hashset ordered: SortedCollectionHashset
// ==========================================

var sortedCollectionHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedCollectionHashset returns sorted collection",
		ArrangeInput: args.Map{
			"when":  "given int hashset with unordered items",
			"items": []int{5, 3, 1, 4, 2},
		},
		ExpectedInput: args.Map{
			"length": 5,
			"first":  1,
			"last":   5,
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
		ExpectedInput: args.Map{
			"min": 1,
			"max": 9,
		},
	},
}

var minMaxHashsetOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashsetOrDefault returns default on empty",
		ArrangeInput: args.Map{
			"when": "given empty int hashset with default -1",
		},
		ExpectedInput: args.Map{
			"min": -1,
			"max": -1,
		},
	},
}

var minMaxHashsetOrDefaultNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashsetOrDefault returns values for non-empty hashset",
		ArrangeInput: args.Map{
			"when":  "given non-empty int hashset with default -1",
			"items": []int{3, 1, 4},
		},
		ExpectedInput: args.Map{
			"min": 1,
			"max": 4,
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "alpha",
			"last":   "gamma",
		},
	},
}

// ==========================================
// Hashmap ordered: SortedKeysDescHashmap
// ==========================================

var sortedKeysDescHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedKeysDescHashmap returns keys in descending order",
		ArrangeInput: args.Map{
			"when": "given string-int hashmap with unordered keys",
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "gamma",
			"last":   "alpha",
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
		ExpectedInput: args.Map{
			"minKey": "alpha",
			"maxKey": "gamma",
		},
	},
}

// ==========================================
// Hashmap ordered: MinKeyHashmapOrDefault / MaxKeyHashmapOrDefault
// ==========================================

var minMaxKeyHashmapOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "MinKeyHashmapOrDefault returns default on empty",
		ArrangeInput: args.Map{
			"when": "given empty string-int hashmap with default 'none'",
		},
		ExpectedInput: args.Map{
			"minKey": "none",
			"maxKey": "none",
		},
	},
}

var minMaxKeyHashmapOrDefaultNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "MinKeyHashmapOrDefault returns values for non-empty hashmap",
		ArrangeInput: args.Map{
			"when": "given non-empty string-int hashmap with default 'none'",
		},
		ExpectedInput: args.Map{
			"minKey": "alpha",
			"maxKey": "gamma",
		},
	},
}

// ==========================================
// Hashmap ordered: SortedValuesHashmap
// ==========================================

var sortedValuesHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedValuesHashmap returns values in ascending order",
		ArrangeInput: args.Map{
			"when": "given string-int hashmap with numeric values",
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  1,
			"last":   30,
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
		ExpectedInput: args.Map{
			"minValue": 1,
			"maxValue": 30,
		},
	},
}

// ==========================================
// Hashmap ordered: MinValueHashmapOrDefault / MaxValueHashmapOrDefault
// ==========================================

var minMaxValueHashmapOrDefaultEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "MinValueHashmapOrDefault returns default on empty",
		ArrangeInput: args.Map{
			"when": "given empty hashmap with default -1",
		},
		ExpectedInput: args.Map{
			"minValue": -1,
			"maxValue": -1,
		},
	},
}

var minMaxValueHashmapOrDefaultNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "MinValueHashmapOrDefault returns values for non-empty",
		ArrangeInput: args.Map{
			"when": "given non-empty string-int hashmap with default -1",
		},
		ExpectedInput: args.Map{
			"minValue": 1,
			"maxValue": 30,
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
		ExpectedInput: args.Map{
			"length": 5,
			"first":  1,
			"last":   5,
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
		ExpectedInput: args.Map{
			"min": 1,
			"max": 9,
		},
	},
}

// ==========================================================================
// EDGE CASES: Empty collections, single elements, negative numbers
// ==========================================================================

// ==========================================
// Edge: SortCollection — empty
// ==========================================

var sortCollectionEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "SortCollection on empty collection produces empty",
		ArrangeInput: args.Map{
			"when":  "given empty int collection",
			"items": []int{},
		},
		ExpectedInput: args.Map{
			"length":   0,
			"isSorted": true,
		},
	},
}

// ==========================================
// Edge: SortCollection — negative numbers
// ==========================================

var sortCollectionNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "SortCollection sorts negative numbers correctly",
		ArrangeInput: args.Map{
			"when":  "given collection with negative and positive values",
			"items": []int{3, -5, 0, -1, 7, -10},
		},
		ExpectedInput: args.Map{
			"length":   6,
			"first":    -10,
			"last":     7,
			"isSorted": true,
		},
	},
}

// ==========================================
// Edge: MinCollection / MaxCollection — negative numbers
// ==========================================

var minMaxCollectionNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "MinCollection and MaxCollection with all negative values",
		ArrangeInput: args.Map{
			"when":  "given collection with only negative values",
			"items": []int{-3, -7, -1, -9, -5},
		},
		ExpectedInput: args.Map{
			"min": -9,
			"max": -1,
		},
	},
	{
		Title: "MinCollection and MaxCollection with mixed positive and negative",
		ArrangeInput: args.Map{
			"when":  "given collection with mixed signs",
			"items": []int{-100, 0, 50, -25, 100},
		},
		ExpectedInput: args.Map{
			"min": -100,
			"max": 100,
		},
	},
}

// ==========================================
// Edge: SumCollection — negative numbers
// ==========================================

var sumCollectionNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "SumCollection with negative numbers",
		ArrangeInput: args.Map{
			"when":  "given collection with negative values",
			"items": []int{-5, 10, -3, 8, -10},
		},
		ExpectedInput: args.Map{"sum": 0},
	},
	{
		Title: "SumCollection with all negative numbers",
		ArrangeInput: args.Map{
			"when":  "given collection with all negative values",
			"items": []int{-1, -2, -3},
		},
		ExpectedInput: args.Map{"sum": -6},
	},
}

// ==========================================
// Edge: ClampCollection — negative range
// ==========================================

var clampCollectionNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "ClampCollection with negative range",
		ArrangeInput: args.Map{
			"when":  "given ints clamped to [-5, -1]",
			"items": []int{-10, -3, 0, -1, -7},
		},
		ExpectedInput: args.Map{
			"val0": -5,
			"val1": -3,
			"val2": -1,
			"val3": -1,
			"val4": -5,
		},
	},
}

// ==========================================
// Edge: IsSortedCollection — single and empty
// ==========================================

var isSortedCollectionEdgeTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSortedCollection true for empty collection",
		ArrangeInput: args.Map{
			"when":  "given empty collection",
			"items": []int{},
		},
		ExpectedInput: args.Map{"isSorted": true},
	},
	{
		Title: "IsSortedCollection true for single element",
		ArrangeInput: args.Map{
			"when":  "given single-element collection",
			"items": []int{99},
		},
		ExpectedInput: args.Map{"isSorted": true},
	},
}

// ==========================================
// Edge: SortedListHashset — single element
// ==========================================

var sortedListHashsetSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedListHashset with single element",
		ArrangeInput: args.Map{
			"when":  "given hashset with single item",
			"items": []int{42},
		},
		ExpectedInput: args.Map{
			"length": 1,
			"first":  42,
			"last":   42,
		},
	},
}

// ==========================================
// Edge: MinHashset / MaxHashset — single element
// ==========================================

var minMaxHashsetSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashset and MaxHashset on single-element hashset",
		ArrangeInput: args.Map{
			"when":  "given hashset with one item",
			"items": []int{7},
		},
		ExpectedInput: args.Map{
			"min": 7,
			"max": 7,
		},
	},
}

// ==========================================
// Edge: MinHashset / MaxHashset — negative numbers
// ==========================================

var minMaxHashsetNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "MinHashset and MaxHashset with negative numbers",
		ArrangeInput: args.Map{
			"when":  "given hashset with negative values",
			"items": []int{-3, -7, 0, 5, -1},
		},
		ExpectedInput: args.Map{
			"min": -7,
			"max": 5,
		},
	},
}

// ==========================================
// Edge: SortedListHashset — negative numbers
// ==========================================

var sortedListHashsetNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedListHashset with negative numbers sorts correctly",
		ArrangeInput: args.Map{
			"when":  "given hashset with mixed signs",
			"items": []int{3, -2, 0, -5, 1},
		},
		ExpectedInput: args.Map{
			"length": 5,
			"first":  -5,
			"last":   3,
		},
	},
}

// ==========================================
// Edge: Hashmap — single entry
// ==========================================

var sortedKeysHashmapSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedKeysHashmap with single entry",
		ArrangeInput: args.Map{
			"when": "given hashmap with one key-value pair",
		},
		ExpectedInput: args.Map{
			"length": 1,
			"first":  "only",
			"last":   "only",
		},
	},
}

var minMaxKeyHashmapSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "MinKeyHashmap and MaxKeyHashmap on single-entry hashmap",
		ArrangeInput: args.Map{
			"when": "given hashmap with one entry",
		},
		ExpectedInput: args.Map{
			"minKey": "only",
			"maxKey": "only",
		},
	},
}

// ==========================================
// Edge: Hashmap — negative values
// ==========================================

var minMaxValueHashmapNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "MinValueHashmap and MaxValueHashmap with negative values",
		ArrangeInput: args.Map{
			"when": "given hashmap with negative integer values",
		},
		ExpectedInput: args.Map{
			"minValue": -20,
			"maxValue": 5,
		},
	},
}

var sortedValuesHashmapNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedValuesHashmap with negative values sorts correctly",
		ArrangeInput: args.Map{
			"when": "given hashmap with mixed sign values",
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  -20,
			"last":   5,
		},
	},
}
