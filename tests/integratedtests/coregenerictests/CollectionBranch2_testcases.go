package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — RemoveAt edge cases
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt middle index",
	ExpectedInput: []string{"true", "4", "1", "5"},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt first index",
	ExpectedInput: []string{"true", "4", "2", "5"},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt last index",
	ExpectedInput: []string{"true", "4", "1", "4"},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt negative index returns false",
	ExpectedInput: []string{"false", "5"},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt out-of-bounds index returns false",
	ExpectedInput: []string{"false", "5"},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt on empty collection returns false",
	ExpectedInput: []string{"false", "0"},
}

// ==========================================================================
// Collection — Reverse
// ==========================================================================

var collectionReversePopulatedTestCase = coretestcases.CaseV1{
	Title:         "Reverse populated collection",
	ExpectedInput: []string{"5", "4", "3", "2", "1"},
}

var collectionReverseSingleTestCase = coretestcases.CaseV1{
	Title:         "Reverse single element",
	ExpectedInput: []string{"42"},
}

var collectionReverseEmptyTestCase = coretestcases.CaseV1{
	Title:         "Reverse empty collection",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — FirstOrDefault / LastOrDefault
// ==========================================================================

var collectionFirstOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault on populated returns first",
	ExpectedInput: []string{"10"},
}

var collectionFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault on empty returns zero",
	ExpectedInput: []string{"0"},
}

var collectionLastOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault on populated returns last",
	ExpectedInput: []string{"30"},
}

var collectionLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault on empty returns zero",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — SafeAt
// ==========================================================================

var collectionSafeAtValidTestCase = coretestcases.CaseV1{
	Title:         "SafeAt valid index returns item",
	ExpectedInput: []string{"20"},
}

var collectionSafeAtNegativeTestCase = coretestcases.CaseV1{
	Title:         "SafeAt negative index returns zero",
	ExpectedInput: []string{"0"},
}

var collectionSafeAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "SafeAt out-of-bounds returns zero",
	ExpectedInput: []string{"0"},
}

var collectionSafeAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "SafeAt on empty returns zero",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — ConcatNew
// ==========================================================================

var collectionConcatNewPopulatedTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew creates new collection with appended items",
	ExpectedInput: []string{"5", "1", "5", "3"},
}

var collectionConcatNewEmptyTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew on empty with items",
	ExpectedInput: []string{"2", "10", "20"},
}

// ==========================================================================
// Collection — AddIf
// ==========================================================================

var collectionAddIfTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIf true adds item",
	ExpectedInput: []string{"1", "42"},
}

var collectionAddIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIf false does not add",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — ForEachBreak on empty
// ==========================================================================

var collectionForEachBreakEmptyTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak on empty does nothing",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — AddSlice
// ==========================================================================

var collectionAddSlicePopulatedTestCase = coretestcases.CaseV1{
	Title:         "AddSlice appends all items from slice",
	ExpectedInput: []string{"3", "10", "30"},
}

var collectionAddSliceEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddSlice with empty slice does nothing",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — Items / ItemsPtr
// ==========================================================================

var collectionItemsSliceTestCase = coretestcases.CaseV1{
	Title:         "Items returns underlying slice",
	ExpectedInput: []string{"3", "1"},
}

var collectionItemsPtrTestCase = coretestcases.CaseV1{
	Title:         "ItemsPtr returns non-nil pointer",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Collection — RemoveAtLock
// ==========================================================================

var collectionRemoveAtLockTestCase = coretestcases.CaseV1{
	Title:         "RemoveAtLock removes item thread-safely",
	ExpectedInput: []string{"true", "2"},
}
