package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty
// ==========================================================================

var hashsetAddNonEmptyAddsTestCase = coretestcases.CaseV1{
	Name:      "AddNonEmpty with non-empty string adds item",
	WantLines: []string{"1", "true"},
}

var hashsetAddNonEmptySkipsEmptyTestCase = coretestcases.CaseV1{
	Name:      "AddNonEmpty with empty string does not add",
	WantLines: []string{"0"},
}

var hashsetAddNonEmptyChainedTestCase = coretestcases.CaseV1{
	Name:      "AddNonEmpty chained adds multiple items",
	WantLines: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// SimpleSlice.InsertAt
// ==========================================================================

var simpleSliceInsertAtMiddleTestCase = coretestcases.CaseV1{
	Name:      "InsertAt middle persists and shifts items",
	WantLines: []string{"4", "a", "X", "b", "c"},
}

var simpleSliceInsertAtPrependTestCase = coretestcases.CaseV1{
	Name:      "InsertAt index 0 prepends",
	WantLines: []string{"4", "X", "a", "b", "c"},
}

var simpleSliceInsertAtAppendTestCase = coretestcases.CaseV1{
	Name:      "InsertAt end appends",
	WantLines: []string{"4", "a", "b", "c", "X"},
}

var simpleSliceInsertAtNegativeTestCase = coretestcases.CaseV1{
	Name:      "InsertAt negative index does nothing",
	WantLines: []string{"3", "a", "b", "c"},
}

var simpleSliceInsertAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Name:      "InsertAt out-of-bounds index does nothing",
	WantLines: []string{"3", "a", "b", "c"},
}

// ==========================================================================
// Collection.RemoveAt
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt valid middle index succeeds",
	WantLines: []string{"true", "2"},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt index 0 succeeds",
	WantLines: []string{"true", "2", "b"},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt last index succeeds",
	WantLines: []string{"true", "2", "b"},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt negative index returns false",
	WantLines: []string{"false", "3"},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt out-of-bounds returns false",
	WantLines: []string{"false", "3"},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt on empty returns false",
	WantLines: []string{"false", "0"},
}

// ==========================================================================
// Hashmap.IsEqualPtr
// ==========================================================================

var hashmapIsEqualPtrSameTestCase = coretestcases.CaseV1{
	Name:      "IsEqualPtr same keys same values → true",
	WantLines: []string{"true"},
}

var hashmapIsEqualPtrDiffValTestCase = coretestcases.CaseV1{
	Name:      "IsEqualPtr same keys different values → false",
	WantLines: []string{"false"},
}

var hashmapIsEqualPtrDiffKeysTestCase = coretestcases.CaseV1{
	Name:      "IsEqualPtr different keys → false",
	WantLines: []string{"false"},
}

var hashmapIsEqualPtrBothEmptyTestCase = coretestcases.CaseV1{
	Name:      "IsEqualPtr both empty → true",
	WantLines: []string{"true"},
}

var hashmapIsEqualPtrNilVsNonNilTestCase = coretestcases.CaseV1{
	Name:      "IsEqualPtr nil vs non-nil → false",
	WantLines: []string{"false"},
}

// ==========================================================================
// Caching removal
// ==========================================================================

var cachingRemovalFreshHashsetTestCase = coretestcases.CaseV1{
	Name:      "Fresh Hashset IsEmpty returns true, Length returns 0",
	WantLines: []string{"true", "0"},
}

var cachingRemovalHashsetAfterAddTestCase = coretestcases.CaseV1{
	Name:      "Hashset IsEmpty false after Add, Length correct",
	WantLines: []string{"false", "2"},
}

var cachingRemovalFreshHashmapTestCase = coretestcases.CaseV1{
	Name:      "Fresh Hashmap IsEmpty returns true, Length returns 0",
	WantLines: []string{"true", "0"},
}

var cachingRemovalHashmapAfterSetTestCase = coretestcases.CaseV1{
	Name:      "Hashmap IsEmpty false after Set, Length correct",
	WantLines: []string{"false", "2"},
}

// ==========================================================================
// SimpleSlice.Skip/Take
// ==========================================================================

var simpleSliceSkipBeyondTestCase = coretestcases.CaseV1{
	Name:      "Skip beyond length returns empty",
	WantLines: []string{"0"},
}

var simpleSliceTakeBeyondTestCase = coretestcases.CaseV1{
	Name:      "Take beyond length returns all",
	WantLines: []string{"3"},
}

var simpleSliceSkipZeroTestCase = coretestcases.CaseV1{
	Name:      "Skip 0 returns all",
	WantLines: []string{"3"},
}

var simpleSliceTakeZeroTestCase = coretestcases.CaseV1{
	Name:      "Take 0 returns empty",
	WantLines: []string{"0"},
}

// ==========================================================================
// HasIndex
// ==========================================================================

var hasIndexNegativeSimpleSliceTestCase = coretestcases.CaseV1{
	Name:      "SimpleSlice.HasIndex negative returns false",
	WantLines: []string{"false"},
}

var hasIndexNegativeCollectionTestCase = coretestcases.CaseV1{
	Name:      "Collection.HasIndex negative returns false",
	WantLines: []string{"false"},
}

// ==========================================================================
// Hashmap.Clear nil safety
// ==========================================================================

var hashmapClearNilReceiverTestCase = coretestcases.CaseV1{
	Name:      "Clear on nil Hashmap returns nil without panic",
	WantLines: []string{"true"},
}

var hashmapClearPopulatedTestCase = coretestcases.CaseV1{
	Name:      "Clear on populated Hashmap resets to empty",
	WantLines: []string{"0", "true"},
}

var hashmapClearChainableTestCase = coretestcases.CaseV1{
	Name:      "Clear preserves chainability",
	WantLines: []string{"0", "1"},
}

// ==========================================================================
// Hashset.AddBool cache invalidation
// ==========================================================================

var hashsetAddBoolNewItemTestCase = coretestcases.CaseV1{
	Name:      "AddBool new item invalidates cache and Items reflects it",
	WantLines: []string{"false", "1", "true"},
}

var hashsetAddBoolExistingTestCase = coretestcases.CaseV1{
	Name:      "AddBool existing item does not change length",
	WantLines: []string{"true", "1"},
}

var hashsetAddBoolMultipleTestCase = coretestcases.CaseV1{
	Name:      "AddBool multiple new items all appear in Items",
	WantLines: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// Hashmap.AddOrUpdateCollection length mismatch
// ==========================================================================

var hashmapAddOrUpdateMismatchedTestCase = coretestcases.CaseV1{
	Name:      "Mismatched lengths returns unchanged hashmap",
	WantLines: []string{"0"},
}

var hashmapAddOrUpdateEqualTestCase = coretestcases.CaseV1{
	Name:      "Equal lengths adds all pairs",
	WantLines: []string{"2", "v1", "v2"},
}

var hashmapAddOrUpdateNilKeysTestCase = coretestcases.CaseV1{
	Name:      "Nil keys returns unchanged",
	WantLines: []string{"0"},
}

var hashmapAddOrUpdateEmptyKeysTestCase = coretestcases.CaseV1{
	Name:      "Empty keys returns unchanged",
	WantLines: []string{"0"},
}
