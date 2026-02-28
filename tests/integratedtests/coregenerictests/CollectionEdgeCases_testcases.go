package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Collection — RemoveAt
// ==========================================

var collectionRemoveAtTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveAt removes item at valid index",
		ArrangeInput: args.Map{
			"when":  "given collection with 3 items, remove index 1",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"true",
			"2",
			"10",
			"30",
		},
	},
	{
		Title: "RemoveAt returns false for out-of-bounds index",
		ArrangeInput: args.Map{
			"when":  "given collection with 3 items, remove index 10",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"false",
			"3",
		},
	},
	{
		Title: "RemoveAt returns false for negative index",
		ArrangeInput: args.Map{
			"when":  "given collection with items, remove index -1",
			"items": []int{10, 20},
		},
		ExpectedInput: []string{
			"false",
			"2",
		},
	},
}

// ==========================================
// Collection — Reverse
// ==========================================

var collectionReverseTestCases = []coretestcases.CaseV1{
	{
		Title: "Reverse reverses collection in-place",
		ArrangeInput: args.Map{
			"when":  "given int collection",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"5",
			"5",
			"1",
		},
	},
	{
		Title: "Reverse single element is no-op",
		ArrangeInput: args.Map{
			"when":  "given single element",
			"items": []int{42},
		},
		ExpectedInput: []string{
			"1",
			"42",
			"42",
		},
	},
}

// ==========================================
// Collection — Skip / Take
// ==========================================

var collectionSkipTakeTestCases = []coretestcases.CaseV1{
	{
		Title: "Skip and Take return correct subsets",
		ArrangeInput: args.Map{
			"when":  "given 5-element collection",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"3",
			"3",
			"2",
			"1",
		},
	},
}

// ==========================================
// Collection — AddIf / AddIfMany
// ==========================================

var collectionAddIfTestCases = []coretestcases.CaseV1{
	{
		Title: "AddIf adds when condition is true, skips when false",
		ArrangeInput: args.Map{
			"when": "given conditional adds",
		},
		ExpectedInput: []string{
			"1",
			"100",
		},
	},
}

// ==========================================
// Collection — FirstOrDefault / LastOrDefault on empty
// ==========================================

var collectionDefaultsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstOrDefault and LastOrDefault return zero on empty",
		ArrangeInput: args.Map{
			"when": "given empty int collection",
		},
		ExpectedInput: []string{
			"0",
			"0",
			"true",
		},
	},
}

// ==========================================
// Collection — SafeAt
// ==========================================

var collectionSafeAtTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeAt returns element at valid index and zero at invalid",
		ArrangeInput: args.Map{
			"when":  "given collection with items",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"20",
			"0",
			"0",
		},
	},
}

// ==========================================
// Collection — ConcatNew
// ==========================================

var collectionConcatNewTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatNew creates new collection without modifying original",
		ArrangeInput: args.Map{
			"when":  "given collection concatenated with more items",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"5",
			"3",
			"1",
			"5",
		},
	},
}

// ==========================================
// Collection — CountFunc
// ==========================================

var collectionCountFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "CountFunc counts items matching predicate",
		ArrangeInput: args.Map{
			"when":  "given ints, count evens",
			"items": []int{1, 2, 3, 4, 5, 6},
		},
		ExpectedInput: []string{
			"3",
		},
	},
}

// ==========================================
// Collection — AddCollection / AddCollections
// ==========================================

var collectionAddCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollection merges another collection",
		ArrangeInput: args.Map{
			"when": "given two collections merged",
		},
		ExpectedInput: []string{
			"5",
			"1",
			"5",
		},
	},
}

// ==========================================
// Hashset — HasAll / HasAny
// ==========================================

var hashsetHasAllHasAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAll true when all present, HasAny true when any present",
		ArrangeInput: args.Map{
			"when":  "given hashset with a, b, c",
			"items": []string{"a", "b", "c"},
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
// Hashset — IsEquals
// ==========================================

var hashsetIsEqualsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEquals true for same content, false for different",
		ArrangeInput: args.Map{
			"when": "given two hashsets to compare",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

// ==========================================
// Hashset — AddBool
// ==========================================

var hashsetAddBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "AddBool returns false for new, true for existing",
		ArrangeInput: args.Map{
			"when": "given hashset with add bool",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

// ==========================================
// Hashmap — Remove
// ==========================================

var hashmapRemoveTestCases = []coretestcases.CaseV1{
	{
		Title: "Remove deletes key and returns existed status",
		ArrangeInput: args.Map{
			"when": "given hashmap with key to remove",
		},
		ExpectedInput: []string{
			"true",
			"1",
			"false",
		},
	},
}

// ==========================================
// Hashmap — GetOrDefault
// ==========================================

var hashmapGetOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "GetOrDefault returns value for existing, default for missing",
		ArrangeInput: args.Map{
			"when": "given hashmap with some keys",
		},
		ExpectedInput: []string{
			"100",
			"-1",
		},
	},
}

// ==========================================
// Hashmap — Clone independence
// ==========================================

var hashmapCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy",
		ArrangeInput: args.Map{
			"when": "given hashmap cloned then mutated",
		},
		ExpectedInput: []string{
			"2",
			"2",
			"true",
		},
	},
}

// ==========================================
// Hashmap — Keys / Values
// ==========================================

var hashmapKeysValuesTestCases = []coretestcases.CaseV1{
	{
		Title: "Keys and Values return correct counts",
		ArrangeInput: args.Map{
			"when": "given hashmap with entries",
		},
		ExpectedInput: []string{
			"3",
			"3",
		},
	},
}

// ==========================================
// Hashmap — IsEquals
// ==========================================

var hashmapIsEqualsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEquals true for same length, false for different",
		ArrangeInput: args.Map{
			"when": "given two hashmaps to compare",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

// ==========================================
// LinkedList — Items / IndexAt / edge cases
// ==========================================

var linkedListItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "Items returns all elements as slice",
		ArrangeInput: args.Map{
			"when":  "given linked list with items",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{
			"3",
			"a",
			"c",
		},
	},
}

var linkedListIndexAtTestCases = []coretestcases.CaseV1{
	{
		Title: "IndexAt returns correct node or nil for out-of-bounds",
		ArrangeInput: args.Map{
			"when":  "given linked list",
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: []string{
			"y",
			"true",
		},
	},
}

var linkedListEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty linked list returns correct defaults",
		ArrangeInput: args.Map{
			"when": "given empty linked list",
		},
		ExpectedInput: []string{
			"0",
			"true",
			"false",
			"",
		},
	},
}

// ==========================================
// SimpleSlice — Filter / Clone / Skip / Take
// ==========================================

var simpleSliceFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice.Filter returns matching items",
		ArrangeInput: args.Map{
			"when":  "given int slice, filter > 2",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"3",
			"3",
			"5",
		},
	},
}

var simpleSliceCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice.Clone creates independent copy",
		ArrangeInput: args.Map{
			"when":  "given int slice cloned then mutated",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"3",
			"true",
		},
	},
}

var simpleSliceSkipTakeTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice Skip and Take return correct subsets",
		ArrangeInput: args.Map{
			"when":  "given 5-element simple slice",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"3",
			"2",
		},
	},
}

// ==========================================
// FlatMapCollection
// ==========================================

var flatMapCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapCollection flattens mapped slices",
		ArrangeInput: args.Map{
			"when":  "given collection of ints mapped to repeated strings",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"6",
			"1",
			"3",
		},
	},
}

// ==========================================
// ReduceCollection
// ==========================================

var reduceCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "ReduceCollection sums all items",
		ArrangeInput: args.Map{
			"when":  "given int collection reduced to sum",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"15",
		},
	},
}

// ==========================================
// GroupByCollection
// ==========================================

var groupByCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupByCollection groups by even/odd",
		ArrangeInput: args.Map{
			"when":  "given int collection grouped by parity",
			"items": []int{1, 2, 3, 4, 5, 6},
		},
		ExpectedInput: []string{
			"2",
			"3",
			"3",
		},
	},
}
