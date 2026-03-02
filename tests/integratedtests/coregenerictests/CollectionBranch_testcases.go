package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — ForEach
// ==========================================================================

var collectionForEachTestCases = []coretestcases.CaseV1{
	{
		Name: "ForEach visits all items with correct indices",
		WantLines: []string{
			"5",
			"0:1",
			"4:5",
		},
	},
	{
		Name: "ForEach on empty collection does nothing",
		WantLines: []string{
			"0",
		},
	},
}

// ==========================================================================
// Collection — ForEachBreak
// ==========================================================================

var collectionForEachBreakTestCases = []coretestcases.CaseV1{
	{
		Name: "ForEachBreak stops at first match",
		WantLines: []string{
			"2",
		},
	},
	{
		Name: "ForEachBreak visits all if no break",
		WantLines: []string{
			"5",
		},
	},
}

// ==========================================================================
// Collection — SortFunc
// ==========================================================================

var collectionSortFuncTestCases = []coretestcases.CaseV1{
	{
		Name: "SortFunc ascending",
		WantLines: []string{
			"1",
			"5",
		},
	},
	{
		Name: "SortFunc descending",
		WantLines: []string{
			"5",
			"1",
		},
	},
	{
		Name: "SortFunc single element",
		WantLines: []string{
			"42",
			"42",
		},
	},
}

// ==========================================================================
// Collection — AddIfMany
// ==========================================================================

var collectionAddIfManyTestCases = []coretestcases.CaseV1{
	{
		Name: "AddIfMany true adds all items",
		WantLines: []string{
			"3",
			"10",
			"30",
		},
	},
	{
		Name: "AddIfMany false adds nothing",
		WantLines: []string{
			"0",
		},
	},
}

// ==========================================================================
// Collection — AddFunc
// ==========================================================================

var collectionAddFuncTestCases = []coretestcases.CaseV1{
	{
		Name: "AddFunc appends result of function",
		WantLines: []string{
			"1",
			"42",
		},
	},
}

// ==========================================================================
// Collection — AddCollections (multiple)
// ==========================================================================

var collectionAddCollectionsTestCases = []coretestcases.CaseV1{
	{
		Name: "AddCollections merges multiple collections",
		WantLines: []string{
			"6",
			"1",
			"6",
		},
	},
	{
		Name: "AddCollections with nil collection skips it",
		WantLines: []string{
			"3",
			"1",
			"3",
		},
	},
}

// ==========================================================================
// Collection — Clone edge cases
// ==========================================================================

var collectionCloneEdgeTestCases = []coretestcases.CaseV1{
	{
		Name: "Clone empty returns empty",
		WantLines: []string{
			"0",
			"true",
		},
	},
}

// ==========================================================================
// Collection — Skip/Take boundary
// ==========================================================================

var collectionSkipTakeBoundaryTestCases = []coretestcases.CaseV1{
	{
		Name: "Skip all returns empty",
		WantLines: []string{
			"0",
		},
	},
	{
		Name: "Take more than length returns all",
		WantLines: []string{
			"3",
		},
	},
	{
		Name: "Skip 0 returns all, Take 0 returns empty",
		WantLines: []string{
			"3",
			"0",
		},
	},
}

// ==========================================================================
// Collection — Filter edge cases
// ==========================================================================

var collectionFilterEdgeTestCases = []coretestcases.CaseV1{
	{
		Name: "Filter no match returns empty",
		WantLines: []string{
			"0",
			"true",
		},
	},
	{
		Name: "Filter all match returns all",
		WantLines: []string{
			"3",
		},
	},
	{
		Name: "Filter empty collection returns empty",
		WantLines: []string{
			"0",
			"true",
		},
	},
}

// ==========================================================================
// Collection — CountFunc edge cases
// ==========================================================================

var collectionCountFuncEdgeTestCases = []coretestcases.CaseV1{
	{
		Name: "CountFunc no match returns 0",
		WantLines: []string{
			"0",
		},
	},
	{
		Name: "CountFunc empty collection returns 0",
		WantLines: []string{
			"0",
		},
	},
}

// ==========================================================================
// Collection — String output
// ==========================================================================

var collectionStringTestCases = []coretestcases.CaseV1{
	{
		Name: "String formats collection",
		WantLines: []string{
			"[1 2 3]",
		},
	},
	{
		Name: "String empty collection",
		WantLines: []string{
			"[]",
		},
	},
}

// ==========================================================================
// Collection — Lock variants
// ==========================================================================

var collectionLockTestCases = []coretestcases.CaseV1{
	{
		Name: "Lock variants work correctly",
		WantLines: []string{
			"3",
			"false",
			"3",
		},
	},
}

// ==========================================================================
// Collection — HasAnyItem / HasItems / HasIndex / LastIndex / Count / Capacity
// ==========================================================================

var collectionMetadataTestCases = []coretestcases.CaseV1{
	{
		Name: "Metadata methods on populated collection",
		WantLines: []string{
			"true",
			"true",
			"true",
			"false",
			"2",
			"3",
		},
	},
	{
		Name: "Metadata methods on empty collection",
		WantLines: []string{
			"false",
			"false",
			"false",
			"-1",
			"0",
		},
	},
}

// ==========================================================================
// Collection — RemoveAt single item
// ==========================================================================

var collectionRemoveAtSingleTestCases = []coretestcases.CaseV1{
	{
		Name: "RemoveAt single item leaves empty collection",
		WantLines: []string{
			"true",
			"0",
			"true",
		},
	},
}

// ==========================================================================
// Collection — AddCollection with nil/empty
// ==========================================================================

var collectionAddCollectionNilTestCases = []coretestcases.CaseV1{
	{
		Name: "AddCollection with empty collection does not change length",
		WantLines: []string{
			"3",
		},
	},
}

// ==========================================================================
// Hashmap — IsEquals updated (key-checking)
// ==========================================================================

var hashmapIsEqualsUpdatedTestCases = []coretestcases.CaseV1{
	{
		Name: "IsEquals same keys → true",
		WantLines: []string{"true"},
	},
	{
		Name: "IsEquals same length different keys → false",
		WantLines: []string{"false"},
	},
	{
		Name: "IsEquals different length → false",
		WantLines: []string{"false"},
	},
	{
		Name: "IsEquals both nil → true",
		WantLines: []string{"true"},
	},
	{
		Name: "IsEquals nil vs non-nil → false",
		WantLines: []string{"false"},
	},
	{
		Name: "IsEquals same pointer → true",
		WantLines: []string{"true"},
	},
}

// ==========================================================================
// Collection — CollectionLenCap
// ==========================================================================

var collectionLenCapTestCases = []coretestcases.CaseV1{
	{
		Name: "CollectionLenCap creates with pre-set length and capacity",
		ArrangeInput: args.Map{},
		WantLines: []string{
			"3",
			"10",
			"0",
		},
	},
}
