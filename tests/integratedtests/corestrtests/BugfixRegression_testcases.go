package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset.AddNonEmpty
// ==========================================================================

var hashsetAddNonEmptyAddsTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty returns length 1 -- non-empty string added",
	ExpectedInput: args.Map{
		"length":       "1",
		"containsItem": "true",
	},
}

var hashsetAddNonEmptySkipsEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddNonEmpty returns length 0 -- empty string skipped",
	ExpectedInput: "0", // length
}

var hashsetAddNonEmptyChainedTestCase = coretestcases.CaseV1{
	Title: "AddNonEmpty returns length 3 -- chained three items",
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
	Title: "InsertAt returns shifted items -- middle index insertion",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "a",
		"item1":  "X",
		"item2":  "b",
		"item3":  "c",
	},
}

var simpleSliceInsertAtPrependTestCase = coretestcases.CaseV1{
	Title: "InsertAt returns prepended item -- index 0",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "X",
		"item1":  "a",
		"item2":  "b",
		"item3":  "c",
	},
}

var simpleSliceInsertAtAppendTestCase = coretestcases.CaseV1{
	Title: "InsertAt returns appended item -- end index",
	ExpectedInput: args.Map{
		"length": "4",
		"item0":  "a",
		"item1":  "b",
		"item2":  "c",
		"item3":  "X",
	},
}

var simpleSliceInsertAtNegativeTestCase = coretestcases.CaseV1{
	Title: "InsertAt returns unchanged slice -- negative index",
	ExpectedInput: args.Map{
		"length": "3",
		"item0":  "a",
		"item1":  "b",
		"item2":  "c",
	},
}

var simpleSliceInsertAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "InsertAt returns unchanged slice -- out-of-bounds index",
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
	Title: "RemoveAt returns true -- valid middle index",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
	},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title: "RemoveAt returns true -- index 0",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
		"newFirstItem":    "b",
	},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title: "RemoveAt returns true -- last index",
	ExpectedInput: args.Map{
		"isRemoved":       "true",
		"remainingLength": "2",
		"lastItem":        "b",
	},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title: "RemoveAt returns false -- negative index",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "3",
	},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "RemoveAt returns false -- out-of-bounds index",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "3",
	},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title: "RemoveAt returns false -- empty collection",
	ExpectedInput: args.Map{
		"isRemoved":       "false",
		"remainingLength": "0",
	},
}

// ==========================================================================
// Hashmap.IsEqualPtr
// ==========================================================================

var hashmapIsEqualPtrSameTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr returns true -- same keys same values",
	ExpectedInput: "true",
}

var hashmapIsEqualPtrDiffValTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr returns false -- same keys different values",
	ExpectedInput: "false",
}

var hashmapIsEqualPtrDiffKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr returns false -- different keys",
	ExpectedInput: "false",
}

var hashmapIsEqualPtrBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr returns true -- both empty",
	ExpectedInput: "true",
}

var hashmapIsEqualPtrNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqualPtr returns false -- nil vs non-nil",
	ExpectedInput: "false",
}

// ==========================================================================
// Caching removal
// ==========================================================================

var cachingRemovalFreshHashsetTestCase = coretestcases.CaseV1{
	Title: "Hashset returns isEmpty true length 0 -- fresh instance",
	ExpectedInput: args.Map{
		"isEmpty": "true",
		"length":  "0",
	},
}

var cachingRemovalHashsetAfterAddTestCase = coretestcases.CaseV1{
	Title: "Hashset returns isEmpty false length 2 -- after Add",
	ExpectedInput: args.Map{
		"isEmpty": "false",
		"length":  "2",
	},
}

var cachingRemovalFreshHashmapTestCase = coretestcases.CaseV1{
	Title: "Hashmap returns isEmpty true length 0 -- fresh instance",
	ExpectedInput: args.Map{
		"isEmpty": "true",
		"length":  "0",
	},
}

var cachingRemovalHashmapAfterSetTestCase = coretestcases.CaseV1{
	Title: "Hashmap returns isEmpty false length 2 -- after Set",
	ExpectedInput: args.Map{
		"isEmpty": "false",
		"length":  "2",
	},
}

// ==========================================================================
// SimpleSlice.Skip/Take
// ==========================================================================

var simpleSliceSkipBeyondTestCase = coretestcases.CaseV1{
	Title:         "Skip returns empty -- count beyond length",
	ExpectedInput: "0", // resultLength
}

var simpleSliceTakeBeyondTestCase = coretestcases.CaseV1{
	Title:         "Take returns all items -- count beyond length",
	ExpectedInput: "3", // resultLength
}

var simpleSliceSkipZeroTestCase = coretestcases.CaseV1{
	Title:         "Skip returns all items -- count 0",
	ExpectedInput: "3", // resultLength
}

var simpleSliceTakeZeroTestCase = coretestcases.CaseV1{
	Title:         "Take returns empty -- count 0",
	ExpectedInput: "0", // resultLength
}

// ==========================================================================
// HasIndex
// ==========================================================================

var hasIndexNegativeSimpleSliceTestCase = coretestcases.CaseV1{
	Title:         "SimpleSlice.HasIndex returns false -- negative index",
	ExpectedInput: "false",
}

var hasIndexNegativeCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection.HasIndex returns false -- negative index",
	ExpectedInput: "false",
}

// ==========================================================================
// Hashmap.Clear nil safety
// Note: Migrated to BugfixRegression_NilReceiver_testcases.go using CaseNilSafe pattern.
// ==========================================================================

var hashmapClearNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "Clear returns nil -- nil Hashmap receiver",
	ExpectedInput: "true", // isNil
}

var hashmapClearPopulatedTestCase = coretestcases.CaseV1{
	Title: "Clear returns empty hashmap -- populated receiver",
	ExpectedInput: args.Map{
		"length":  "0",
		"isEmpty": "true",
	},
}

var hashmapClearChainableTestCase = coretestcases.CaseV1{
	Title: "Clear returns chainable instance -- re-add after clear",
	ExpectedInput: args.Map{
		"lengthAfterClear": "0",
		"lengthAfterReAdd": "1",
	},
}

// ==========================================================================
// Hashset.AddBool cache invalidation
// ==========================================================================

var hashsetAddBoolNewItemTestCase = coretestcases.CaseV1{
	Title: "AddBool returns false existed, length 1 -- new item added",
	ExpectedInput: args.Map{
		"existedBefore": "false",
		"lengthAfter":   "1",
		"itemsContains": "true",
	},
}

var hashsetAddBoolExistingTestCase = coretestcases.CaseV1{
	Title: "AddBool returns true existed, same length -- existing item",
	ExpectedInput: args.Map{
		"existedBefore": "true",
		"lengthAfter":   "1",
	},
}

var hashsetAddBoolMultipleTestCase = coretestcases.CaseV1{
	Title: "AddBool returns length 3 -- three new items added",
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
	Title:         "AddOrUpdateCollection returns length 0 -- mismatched key-value lengths",
	ExpectedInput: "0", // length
}

var hashmapAddOrUpdateEqualTestCase = coretestcases.CaseV1{
	Title: "AddOrUpdateCollection returns length 2 -- equal key-value lengths",
	ExpectedInput: args.Map{
		"length": "2",
		"value1": "v1",
		"value2": "v2",
	},
}

var hashmapAddOrUpdateNilKeysTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateCollection returns length 0 -- nil keys",
	ExpectedInput: "0", // length
}

var hashmapAddOrUpdateEmptyKeysTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateCollection returns length 0 -- empty keys",
	ExpectedInput: "0", // length
}
