package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — RemoveAt edge cases
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt middle index",
	WantLines: []string{"true", "4", "1", "5"},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt first index",
	WantLines: []string{"true", "4", "2", "5"},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt last index",
	WantLines: []string{"true", "4", "1", "4"},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt negative index returns false",
	WantLines: []string{"false", "5"},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt out-of-bounds index returns false",
	WantLines: []string{"false", "5"},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt on empty collection returns false",
	WantLines: []string{"false", "0"},
}

// ==========================================================================
// Collection — Reverse
// ==========================================================================

var collectionReversePopulatedTestCase = coretestcases.CaseV1{
	Name:      "Reverse populated collection",
	WantLines: []string{"5", "4", "3", "2", "1"},
}

var collectionReverseSingleTestCase = coretestcases.CaseV1{
	Name:      "Reverse single element",
	WantLines: []string{"42"},
}

var collectionReverseEmptyTestCase = coretestcases.CaseV1{
	Name:      "Reverse empty collection",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — FirstOrDefault / LastOrDefault
// ==========================================================================

var collectionFirstOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Name:      "FirstOrDefault on populated returns first",
	WantLines: []string{"10"},
}

var collectionFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Name:      "FirstOrDefault on empty returns zero",
	WantLines: []string{"0"},
}

var collectionLastOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Name:      "LastOrDefault on populated returns last",
	WantLines: []string{"30"},
}

var collectionLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Name:      "LastOrDefault on empty returns zero",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — SafeAt
// ==========================================================================

var collectionSafeAtValidTestCase = coretestcases.CaseV1{
	Name:      "SafeAt valid index returns item",
	WantLines: []string{"20"},
}

var collectionSafeAtNegativeTestCase = coretestcases.CaseV1{
	Name:      "SafeAt negative index returns zero",
	WantLines: []string{"0"},
}

var collectionSafeAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Name:      "SafeAt out-of-bounds returns zero",
	WantLines: []string{"0"},
}

var collectionSafeAtEmptyTestCase = coretestcases.CaseV1{
	Name:      "SafeAt on empty returns zero",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — ConcatNew
// ==========================================================================

var collectionConcatNewPopulatedTestCase = coretestcases.CaseV1{
	Name:      "ConcatNew creates new collection with appended items",
	WantLines: []string{"5", "1", "5", "3"},
}

var collectionConcatNewEmptyTestCase = coretestcases.CaseV1{
	Name:      "ConcatNew on empty with items",
	WantLines: []string{"2", "10", "20"},
}

// ==========================================================================
// Collection — AddIf
// ==========================================================================

var collectionAddIfTrueTestCase = coretestcases.CaseV1{
	Name:      "AddIf true adds item",
	WantLines: []string{"1", "42"},
}

var collectionAddIfFalseTestCase = coretestcases.CaseV1{
	Name:      "AddIf false does not add",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — ForEachBreak on empty
// ==========================================================================

var collectionForEachBreakEmptyTestCase = coretestcases.CaseV1{
	Name:      "ForEachBreak on empty does nothing",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — AddSlice
// ==========================================================================

var collectionAddSlicePopulatedTestCase = coretestcases.CaseV1{
	Name:      "AddSlice appends all items from slice",
	WantLines: []string{"3", "10", "30"},
}

var collectionAddSliceEmptyTestCase = coretestcases.CaseV1{
	Name:      "AddSlice with empty slice does nothing",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — Items / ItemsPtr
// ==========================================================================

var collectionItemsSliceTestCase = coretestcases.CaseV1{
	Name:      "Items returns underlying slice",
	WantLines: []string{"3", "1"},
}

var collectionItemsPtrTestCase = coretestcases.CaseV1{
	Name:      "ItemsPtr returns non-nil pointer",
	WantLines: []string{"true"},
}

// ==========================================================================
// Collection — RemoveAtLock
// ==========================================================================

var collectionRemoveAtLockTestCase = coretestcases.CaseV1{
	Name:      "RemoveAtLock removes item thread-safely",
	WantLines: []string{"true", "2"},
}
