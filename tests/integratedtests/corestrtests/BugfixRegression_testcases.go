package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty — regression for no-op bug
// ==========================================================================

var hashsetAddNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Name:      "AddNonEmpty with non-empty string adds item",
		WantLines: []string{"1", "true"},
	},
	{
		Name:      "AddNonEmpty with empty string does not add",
		WantLines: []string{"0"},
	},
	{
		Name:      "AddNonEmpty chained adds multiple items",
		WantLines: []string{"3", "true", "true", "true"},
	},
}

// ==========================================================================
// SimpleSlice.InsertAt — regression for not-persisting + no bounds
// ==========================================================================

var simpleSliceInsertAtTestCases = []coretestcases.CaseV1{
	{
		Name:      "InsertAt middle persists and shifts items",
		WantLines: []string{"4", "a", "X", "b", "c"},
	},
	{
		Name:      "InsertAt index 0 prepends",
		WantLines: []string{"4", "X", "a", "b", "c"},
	},
	{
		Name:      "InsertAt end appends",
		WantLines: []string{"4", "a", "b", "c", "X"},
	},
	{
		Name:      "InsertAt negative index does nothing",
		WantLines: []string{"3", "a", "b", "c"},
	},
	{
		Name:      "InsertAt out-of-bounds index does nothing",
		WantLines: []string{"3", "a", "b", "c"},
	},
}

// ==========================================================================
// Collection.RemoveAt — regression for inverted guard
// ==========================================================================

var collectionRemoveAtTestCases = []coretestcases.CaseV1{
	{
		Name:      "RemoveAt valid middle index succeeds",
		WantLines: []string{"true", "2"},
	},
	{
		Name:      "RemoveAt index 0 succeeds",
		WantLines: []string{"true", "2", "b"},
	},
	{
		Name:      "RemoveAt last index succeeds",
		WantLines: []string{"true", "2", "b"},
	},
	{
		Name:      "RemoveAt negative index returns false",
		WantLines: []string{"false", "3"},
	},
	{
		Name:      "RemoveAt out-of-bounds returns false",
		WantLines: []string{"false", "3"},
	},
	{
		Name:      "RemoveAt on empty returns false",
		WantLines: []string{"false", "0"},
	},
}

// ==========================================================================
// Hashmap.IsEqualPtr — regression for inverted comparison
// ==========================================================================

var hashmapIsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Name:      "IsEqualPtr same keys same values → true",
		WantLines: []string{"true"},
	},
	{
		Name:      "IsEqualPtr same keys different values → false",
		WantLines: []string{"false"},
	},
	{
		Name:      "IsEqualPtr different keys → false",
		WantLines: []string{"false"},
	},
	{
		Name:      "IsEqualPtr both empty → true",
		WantLines: []string{"true"},
	},
	{
		Name:      "IsEqualPtr nil vs non-nil → false",
		WantLines: []string{"false"},
	},
}

// ==========================================================================
// Caching removal — IsEmpty/Length on fresh instances
// ==========================================================================

var cachingRemovalTestCases = []coretestcases.CaseV1{
	{
		Name:      "Fresh Hashset IsEmpty returns true, Length returns 0",
		WantLines: []string{"true", "0"},
	},
	{
		Name:      "Hashset IsEmpty false after Add, Length correct",
		WantLines: []string{"false", "2"},
	},
	{
		Name:      "Fresh Hashmap IsEmpty returns true, Length returns 0",
		WantLines: []string{"true", "0"},
	},
	{
		Name:      "Hashmap IsEmpty false after Set, Length correct",
		WantLines: []string{"false", "2"},
	},
}

// ==========================================================================
// SimpleSlice.Skip/Take — regression for bounds protection
// ==========================================================================

var simpleSliceSkipTakeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Skip beyond length returns empty",
		WantLines: []string{"0"},
	},
	{
		Name:      "Take beyond length returns all",
		WantLines: []string{"3"},
	},
	{
		Name:      "Skip 0 returns all",
		WantLines: []string{"3"},
	},
	{
		Name:      "Take 0 returns empty",
		WantLines: []string{"0"},
	},
}

// ==========================================================================
// HasIndex — regression for negative index guard
// ==========================================================================

var hasIndexNegativeTestCases = []coretestcases.CaseV1{
	{
		Name:      "SimpleSlice.HasIndex negative returns false",
		WantLines: []string{"false"},
	},
	{
		Name:      "Collection.HasIndex negative returns false",
		WantLines: []string{"false"},
	},
}

// ==========================================================================
// Hashmap.Clear nil safety — regression for nil panic
// ==========================================================================

var hashmapClearNilTestCases = []coretestcases.CaseV1{
	{
		Name:      "Clear on nil Hashmap returns nil without panic",
		WantLines: []string{"true"},
	},
	{
		Name:      "Clear on populated Hashmap resets to empty",
		WantLines: []string{"0", "true"},
	},
	{
		Name:      "Clear preserves chainability",
		WantLines: []string{"0", "1"},
	},
}

// ==========================================================================
// Hashset.AddBool cache invalidation — regression for stale cache
// ==========================================================================

var hashsetAddBoolCacheTestCases = []coretestcases.CaseV1{
	{
		Name:      "AddBool new item invalidates cache and Items reflects it",
		WantLines: []string{"false", "1", "true"},
	},
	{
		Name:      "AddBool existing item does not change length",
		WantLines: []string{"true", "1"},
	},
	{
		Name:      "AddBool multiple new items all appear in Items",
		WantLines: []string{"3", "true", "true", "true"},
	},
}

// ==========================================================================
// Hashmap.AddOrUpdateCollection length mismatch — regression for silent corruption
// ==========================================================================

var hashmapAddOrUpdateCollectionMismatchTestCases = []coretestcases.CaseV1{
	{
		Name:      "Mismatched lengths returns unchanged hashmap",
		WantLines: []string{"0"},
	},
	{
		Name:      "Equal lengths adds all pairs",
		WantLines: []string{"2", "v1", "v2"},
	},
	{
		Name:      "Nil keys returns unchanged",
		WantLines: []string{"0"},
	},
	{
		Name:      "Empty keys returns unchanged",
		WantLines: []string{"0"},
	},
}
