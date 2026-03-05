package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty
// ==========================================================================

var hashsetAddNonEmptyAddsTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty with non-empty string adds item",
	ExpectedInput: args.Two[string, string]{
		First:  "1",    // length
		Second: "true", // containsItem
	},
}

var hashsetAddNonEmptySkipsEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty with empty string does not add",
	ExpectedInput: "0", // length
}

var hashsetAddNonEmptyChainedTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty chained adds multiple items",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "3",    // length
		Second: "true", // containsItem1
		Third:  "true", // containsItem2
		Fourth: "true", // containsItem3
	},
}

// ==========================================================================
// SimpleSlice.InsertAt
// ==========================================================================

var simpleSliceInsertAtMiddleTestCase = coretestcases.CaseV1{
	Title: "InsertAt middle persists and shifts items",
	ExpectedInput: args.Five[string, string, string, string, string]{
		First:  "4", // length
		Second: "a", // item0
		Third:  "X", // item1
		Fourth: "b", // item2
		Fifth:  "c", // item3
	},
}

var simpleSliceInsertAtPrependTestCase = coretestcases.CaseV1{
	Title: "InsertAt index 0 prepends",
	ExpectedInput: args.Five[string, string, string, string, string]{
		First:  "4", // length
		Second: "X", // item0
		Third:  "a", // item1
		Fourth: "b", // item2
		Fifth:  "c", // item3
	},
}

var simpleSliceInsertAtAppendTestCase = coretestcases.CaseV1{
	Title: "InsertAt end appends",
	ExpectedInput: args.Five[string, string, string, string, string]{
		First:  "4", // length
		Second: "a", // item0
		Third:  "b", // item1
		Fourth: "c", // item2
		Fifth:  "X", // item3
	},
}

var simpleSliceInsertAtNegativeTestCase = coretestcases.CaseV1{
	Title: "InsertAt negative index does nothing",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "3", // length
		Second: "a", // item0
		Third:  "b", // item1
		Fourth: "c", // item2
	},
}

var simpleSliceInsertAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "InsertAt out-of-bounds index does nothing",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "3", // length
		Second: "a", // item0
		Third:  "b", // item1
		Fourth: "c", // item2
	},
}

// ==========================================================================
// Collection.RemoveAt
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Title: "RemoveAt valid middle index succeeds",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isRemoved
		Second: "2",    // remainingLength
	},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title: "RemoveAt index 0 succeeds",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true", // isRemoved
		Second: "2",    // remainingLength
		Third:  "b",    // newFirstItem
	},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title: "RemoveAt last index succeeds",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true", // isRemoved
		Second: "2",    // remainingLength
		Third:  "b",    // lastItem
	},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title: "RemoveAt negative index returns false",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isRemoved
		Second: "3",     // remainingLength
	},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "RemoveAt out-of-bounds returns false",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isRemoved
		Second: "3",     // remainingLength
	},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title: "RemoveAt on empty returns false",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isRemoved
		Second: "0",     // remainingLength
	},
}

// ==========================================================================
// Hashmap.IsEqualPtr
// ==========================================================================

var hashmapIsEqualPtrSameTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr same keys same values → true",
	ExpectedInput: "true",
}

var hashmapIsEqualPtrDiffValTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr same keys different values → false",
	ExpectedInput: "false",
}

var hashmapIsEqualPtrDiffKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr different keys → false",
	ExpectedInput: "false",
}

var hashmapIsEqualPtrBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr both empty → true",
	ExpectedInput: "true",
}

var hashmapIsEqualPtrNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr nil vs non-nil → false",
	ExpectedInput: "false",
}

// ==========================================================================
// Caching removal
// ==========================================================================

var cachingRemovalFreshHashsetTestCase = coretestcases.CaseV1{
	Title: "Fresh Hashset IsEmpty returns true, Length returns 0",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isEmpty
		Second: "0",    // length
	},
}

var cachingRemovalHashsetAfterAddTestCase = coretestcases.CaseV1{
	Title: "Hashset IsEmpty false after Add, Length correct",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isEmpty
		Second: "2",     // length
	},
}

var cachingRemovalFreshHashmapTestCase = coretestcases.CaseV1{
	Title: "Fresh Hashmap IsEmpty returns true, Length returns 0",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isEmpty
		Second: "0",    // length
	},
}

var cachingRemovalHashmapAfterSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap IsEmpty false after Set, Length correct",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isEmpty
		Second: "2",     // length
	},
}

// ==========================================================================
// SimpleSlice.Skip/Take
// ==========================================================================

var simpleSliceSkipBeyondTestCase = coretestcases.CaseV1{
	Title:         "Skip beyond length returns empty",
	ExpectedInput: "0", // resultLength
}

var simpleSliceTakeBeyondTestCase = coretestcases.CaseV1{
	Title:         "Take beyond length returns all",
	ExpectedInput: "3", // resultLength
}

var simpleSliceSkipZeroTestCase = coretestcases.CaseV1{
	Title:         "Skip 0 returns all",
	ExpectedInput: "3", // resultLength
}

var simpleSliceTakeZeroTestCase = coretestcases.CaseV1{
	Title:         "Take 0 returns empty",
	ExpectedInput: "0", // resultLength
}

// ==========================================================================
// HasIndex
// ==========================================================================

var hasIndexNegativeSimpleSliceTestCase = coretestcases.CaseV1{
	Title:         "SimpleSlice.HasIndex negative returns false",
	ExpectedInput: "false",
}

var hasIndexNegativeCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection.HasIndex negative returns false",
	ExpectedInput: "false",
}

// ==========================================================================
// Hashmap.Clear nil safety
// ==========================================================================

var hashmapClearNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "Clear on nil Hashmap returns nil without panic",
	ExpectedInput: "true", // isNil
}

var hashmapClearPopulatedTestCase = coretestcases.CaseV1{
	Title: "Clear on populated Hashmap resets to empty",
	ExpectedInput: args.Two[string, string]{
		First:  "0",    // length
		Second: "true", // isEmpty
	},
}

var hashmapClearChainableTestCase = coretestcases.CaseV1{
	Title: "Clear preserves chainability",
	ExpectedInput: args.Two[string, string]{
		First:  "0", // lengthAfterClear
		Second: "1", // lengthAfterReAdd
	},
}

// ==========================================================================
// Hashset.AddBool cache invalidation
// ==========================================================================

var hashsetAddBoolNewItemTestCase = coretestcases.CaseV1{
	Title: "AddBool new item invalidates cache and Items reflects it",
	ExpectedInput: args.Three[string, string, string]{
		First:  "false", // existedBefore
		Second: "1",     // lengthAfter
		Third:  "true",  // itemsContains
	},
}

var hashsetAddBoolExistingTestCase = coretestcases.CaseV1{
	Title: "AddBool existing item does not change length",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // existedBefore
		Second: "1",    // lengthAfter
	},
}

var hashsetAddBoolMultipleTestCase = coretestcases.CaseV1{
	Title: "AddBool multiple new items all appear in Items",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "3",    // length
		Second: "true", // containsItem1
		Third:  "true", // containsItem2
		Fourth: "true", // containsItem3
	},
}

// ==========================================================================
// Hashmap.AddOrUpdateCollection length mismatch
// ==========================================================================

var hashmapAddOrUpdateMismatchedTestCase = coretestcases.CaseV1{
	Title:         "Mismatched lengths returns unchanged hashmap",
	ExpectedInput: "0", // length
}

var hashmapAddOrUpdateEqualTestCase = coretestcases.CaseV1{
	Title: "Equal lengths adds all pairs",
	ExpectedInput: args.Three[string, string, string]{
		First:  "2",  // length
		Second: "v1", // value1
		Third:  "v2", // value2
	},
}

var hashmapAddOrUpdateNilKeysTestCase = coretestcases.CaseV1{
	Title:         "Nil keys returns unchanged",
	ExpectedInput: "0", // length
}

var hashmapAddOrUpdateEmptyKeysTestCase = coretestcases.CaseV1{
	Title:         "Empty keys returns unchanged",
	ExpectedInput: "0", // length
}
