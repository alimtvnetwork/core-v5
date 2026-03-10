package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty
// ==========================================================================

var hashsetAddNonEmptyAddsTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty with non-empty string adds item",
	ExpectedInput: args.Map{
		"length":       "1",
		"containsItem": "true",
	},
}

var hashsetAddNonEmptySkipsEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty with empty string does not add",
	ExpectedInput: "0", // length
}

var hashsetAddNonEmptyChainedTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty chained adds multiple items",
	ExpectedInput: args.Map{
		"length":        "3",
		"containsItem1": "true",
		"containsItem2": "true",
		"containsItem3": "true",
	},
}

// ==========================================================================
// SimpleSlice.InsertAt
// ==========================================================================

var simpleSliceInsertAtMiddleTestCase = coretestcases.CaseV1{
	Title: "InsertAt middle persists and shifts items",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "a",
		"item1":  "X",
		"item2":  "b",
		"item3":  "c",
	},
}

var simpleSliceInsertAtPrependTestCase = coretestcases.CaseV1{
	Title: "InsertAt index 0 prepends",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "X",
		"item1":  "a",
		"item2":  "b",
		"item3":  "c",
	},
}

var simpleSliceInsertAtAppendTestCase = coretestcases.CaseV1{
	Title: "InsertAt end appends",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "a",
		"item1":  "b",
		"item2":  "c",
		"item3":  "X",
	},
}

var simpleSliceInsertAtNegativeTestCase = coretestcases.CaseV1{
	Title: "InsertAt negative index does nothing",
	ExpectedInput: args.Map{
		"length": "3",
		"item0":  "a",
		"item1":  "b",
		"item2":  "c",
	},
}

var simpleSliceInsertAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "InsertAt out-of-bounds index does nothing",
	ExpectedInput: args.Map{
		"length": "3",
		"item0":  "a",
		"item1":  "b",
		"item2":  "c",
	},
}

// ==========================================================================
// Collection.RemoveAt
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Title: "RemoveAt valid middle index succeeds",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
	},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title: "RemoveAt index 0 succeeds",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
		"newFirstItem":    "b",
	},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title: "RemoveAt last index succeeds",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
		"lastItem":        "b",
	},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title: "RemoveAt negative index returns false",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "3",
	},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "RemoveAt out-of-bounds returns false",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "3",
	},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title: "RemoveAt on empty returns false",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "0",
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
	ExpectedInput: args.Map{
		"isEmpty": "true",
		"length":  "0",
	},
}

var cachingRemovalHashsetAfterAddTestCase = coretestcases.CaseV1{
	Title: "Hashset IsEmpty false after Add, Length correct",
	ExpectedInput: args.Map{
		"isEmpty": "false",
		"length":  "2",
	},
}

var cachingRemovalFreshHashmapTestCase = coretestcases.CaseV1{
	Title: "Fresh Hashmap IsEmpty returns true, Length returns 0",
	ExpectedInput: args.Map{
		"isEmpty": "true",
		"length":  "0",
	},
}

var cachingRemovalHashmapAfterSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap IsEmpty false after Set, Length correct",
	ExpectedInput: args.Map{
		"isEmpty": "false",
		"length":  "2",
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
// Note: Migrated to BugfixRegression_NilReceiver_testcases.go using CaseNilSafe pattern.
// ==========================================================================

var hashmapClearNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "Clear on nil Hashmap returns nil without panic",
	ExpectedInput: "true", // isNil
}

var hashmapClearPopulatedTestCase = coretestcases.CaseV1{
	Title: "Clear on populated Hashmap resets to empty",
	ExpectedInput: args.Map{
		"length":  "0",
		"isEmpty": "true",
	},
}

var hashmapClearChainableTestCase = coretestcases.CaseV1{
	Title: "Clear preserves chainability",
	ExpectedInput: args.Map{
		"lengthAfterClear": "0",
		"lengthAfterReAdd": "1",
	},
}

// ==========================================================================
// Hashset.AddBool cache invalidation
// ==========================================================================

var hashsetAddBoolNewItemTestCase = coretestcases.CaseV1{
	Title: "AddBool new item invalidates cache and Items reflects it",
	ExpectedInput: args.Map{
		"existedBefore": "false",
		"lengthAfter":   "1",
		"itemsContains": "true",
	},
}

var hashsetAddBoolExistingTestCase = coretestcases.CaseV1{
	Title: "AddBool existing item does not change length",
	ExpectedInput: args.Map{
		"existedBefore": "true",
		"lengthAfter":   "1",
	},
}

var hashsetAddBoolMultipleTestCase = coretestcases.CaseV1{
	Title: "AddBool multiple new items all appear in Items",
	ExpectedInput: args.Map{
		"length":        "3",
		"containsItem1": "true",
		"containsItem2": "true",
		"containsItem3": "true",
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
	ExpectedInput: args.Map{
		"length": "2",
		"value1": "v1",
		"value2": "v2",
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
