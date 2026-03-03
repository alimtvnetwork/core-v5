package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty
// ==========================================================================

var hashsetAddNonEmptyAddsTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty with non-empty string adds item",
	ExpectedInput: []string{"1", "true"},
}

var hashsetAddNonEmptySkipsEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty with empty string does not add",
	ExpectedInput: []string{"0"},
}

var hashsetAddNonEmptyChainedTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty chained adds multiple items",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// SimpleSlice.InsertAt
// ==========================================================================

var simpleSliceInsertAtMiddleTestCase = coretestcases.CaseV1{
	Title:         "InsertAt middle persists and shifts items",
	ExpectedInput: []string{"4", "a", "X", "b", "c"},
}

var simpleSliceInsertAtPrependTestCase = coretestcases.CaseV1{
	Title:         "InsertAt index 0 prepends",
	ExpectedInput: []string{"4", "X", "a", "b", "c"},
}

var simpleSliceInsertAtAppendTestCase = coretestcases.CaseV1{
	Title:         "InsertAt end appends",
	ExpectedInput: []string{"4", "a", "b", "c", "X"},
}

var simpleSliceInsertAtNegativeTestCase = coretestcases.CaseV1{
	Title:         "InsertAt negative index does nothing",
	ExpectedInput: []string{"3", "a", "b", "c"},
}

var simpleSliceInsertAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "InsertAt out-of-bounds index does nothing",
	ExpectedInput: []string{"3", "a", "b", "c"},
}

// ==========================================================================
// Collection.RemoveAt
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt valid middle index succeeds",
	ExpectedInput: []string{"true", "2"},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt index 0 succeeds",
	ExpectedInput: []string{"true", "2", "b"},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt last index succeeds",
	ExpectedInput: []string{"true", "2", "b"},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt negative index returns false",
	ExpectedInput: []string{"false", "3"},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt out-of-bounds returns false",
	ExpectedInput: []string{"false", "3"},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt on empty returns false",
	ExpectedInput: []string{"false", "0"},
}

// ==========================================================================
// Hashmap.IsEqualPtr
// ==========================================================================

var hashmapIsEqualPtrSameTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr same keys same values → true",
	ExpectedInput: []string{"true"},
}

var hashmapIsEqualPtrDiffValTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr same keys different values → false",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualPtrDiffKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr different keys → false",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualPtrBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr both empty → true",
	ExpectedInput: []string{"true"},
}

var hashmapIsEqualPtrNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr nil vs non-nil → false",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Caching removal
// ==========================================================================

var cachingRemovalFreshHashsetTestCase = coretestcases.CaseV1{
	Title:         "Fresh Hashset IsEmpty returns true, Length returns 0",
	ExpectedInput: []string{"true", "0"},
}

var cachingRemovalHashsetAfterAddTestCase = coretestcases.CaseV1{
	Title:         "Hashset IsEmpty false after Add, Length correct",
	ExpectedInput: []string{"false", "2"},
}

var cachingRemovalFreshHashmapTestCase = coretestcases.CaseV1{
	Title:         "Fresh Hashmap IsEmpty returns true, Length returns 0",
	ExpectedInput: []string{"true", "0"},
}

var cachingRemovalHashmapAfterSetTestCase = coretestcases.CaseV1{
	Title:         "Hashmap IsEmpty false after Set, Length correct",
	ExpectedInput: []string{"false", "2"},
}

// ==========================================================================
// SimpleSlice.Skip/Take
// ==========================================================================

var simpleSliceSkipBeyondTestCase = coretestcases.CaseV1{
	Title:         "Skip beyond length returns empty",
	ExpectedInput: []string{"0"},
}

var simpleSliceTakeBeyondTestCase = coretestcases.CaseV1{
	Title:         "Take beyond length returns all",
	ExpectedInput: []string{"3"},
}

var simpleSliceSkipZeroTestCase = coretestcases.CaseV1{
	Title:         "Skip 0 returns all",
	ExpectedInput: []string{"3"},
}

var simpleSliceTakeZeroTestCase = coretestcases.CaseV1{
	Title:         "Take 0 returns empty",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// HasIndex
// ==========================================================================

var hasIndexNegativeSimpleSliceTestCase = coretestcases.CaseV1{
	Title:         "SimpleSlice.HasIndex negative returns false",
	ExpectedInput: []string{"false"},
}

var hasIndexNegativeCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection.HasIndex negative returns false",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Hashmap.Clear nil safety
// ==========================================================================

var hashmapClearNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "Clear on nil Hashmap returns nil without panic",
	ExpectedInput: []string{"true"},
}

var hashmapClearPopulatedTestCase = coretestcases.CaseV1{
	Title:         "Clear on populated Hashmap resets to empty",
	ExpectedInput: []string{"0", "true"},
}

var hashmapClearChainableTestCase = coretestcases.CaseV1{
	Title:         "Clear preserves chainability",
	ExpectedInput: []string{"0", "1"},
}

// ==========================================================================
// Hashset.AddBool cache invalidation
// ==========================================================================

var hashsetAddBoolNewItemTestCase = coretestcases.CaseV1{
	Title:         "AddBool new item invalidates cache and Items reflects it",
	ExpectedInput: []string{"false", "1", "true"},
}

var hashsetAddBoolExistingTestCase = coretestcases.CaseV1{
	Title:         "AddBool existing item does not change length",
	ExpectedInput: []string{"true", "1"},
}

var hashsetAddBoolMultipleTestCase = coretestcases.CaseV1{
	Title:         "AddBool multiple new items all appear in Items",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// Hashmap.AddOrUpdateCollection length mismatch
// ==========================================================================

var hashmapAddOrUpdateMismatchedTestCase = coretestcases.CaseV1{
	Title:         "Mismatched lengths returns unchanged hashmap",
	ExpectedInput: []string{"0"},
}

var hashmapAddOrUpdateEqualTestCase = coretestcases.CaseV1{
	Title:         "Equal lengths adds all pairs",
	ExpectedInput: []string{"2", "v1", "v2"},
}

var hashmapAddOrUpdateNilKeysTestCase = coretestcases.CaseV1{
	Title:         "Nil keys returns unchanged",
	ExpectedInput: []string{"0"},
}

var hashmapAddOrUpdateEmptyKeysTestCase = coretestcases.CaseV1{
	Title:         "Empty keys returns unchanged",
	ExpectedInput: []string{"0"},
}
