package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset — Add / AddBool edge cases
// ==========================================================================

var hashsetAddDuplicateTestCase = coretestcases.CaseV1{
	Title:         "Add duplicate does not increase length",
	ExpectedInput: []string{"3"},
}

var hashsetAddBoolTestCase = coretestcases.CaseV1{
	Title:         "AddBool returns false for new, true for existing",
	ExpectedInput: []string{"false", "true", "1"},
}

var hashsetAddsVariadicTestCase = coretestcases.CaseV1{
	Title:         "Adds variadic adds all unique items",
	ExpectedInput: []string{"4"},
}

var hashsetAddSliceTestCase = coretestcases.CaseV1{
	Title:         "AddSlice adds all items from slice",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// Hashset — AddIf / AddIfMany
// ==========================================================================

var hashsetAddIfTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIf true adds item",
	ExpectedInput: []string{"1", "true"},
}

var hashsetAddIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIf false skips item",
	ExpectedInput: []string{"0"},
}

var hashsetAddIfManyTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany true adds all",
	ExpectedInput: []string{"3"},
}

var hashsetAddIfManyFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany false adds none",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

var hashsetMergeOtherSetTestCase = coretestcases.CaseV1{
	Title:         "AddHashsetItems merges other set",
	ExpectedInput: []string{"4", "true", "true"},
}

var hashsetMergeNilOtherTestCase = coretestcases.CaseV1{
	Title:         "AddHashsetItems with nil other does nothing",
	ExpectedInput: []string{"2"},
}

var hashsetMergeEmptyOtherTestCase = coretestcases.CaseV1{
	Title:         "AddHashsetItems with empty other does nothing",
	ExpectedInput: []string{"2"},
}

var hashsetAddItemsMapTestCase = coretestcases.CaseV1{
	Title:         "AddItemsMap adds only true entries",
	ExpectedInput: []string{"2", "true", "false"},
}

// ==========================================================================
// Hashset — Remove edge cases
// ==========================================================================

var hashsetRemoveExistingTestCase = coretestcases.CaseV1{
	Title:         "Remove existing returns true and decreases length",
	ExpectedInput: []string{"true", "2", "false"},
}

var hashsetRemoveNonExistingTestCase = coretestcases.CaseV1{
	Title:         "Remove non-existing returns false",
	ExpectedInput: []string{"false", "3"},
}

// ==========================================================================
// Hashset — Has / Contains
// ==========================================================================

var hashsetHasTestCase = coretestcases.CaseV1{
	Title:         "Has returns true for existing, false for missing",
	ExpectedInput: []string{"true", "false"},
}

var hashsetContainsAliasTestCase = coretestcases.CaseV1{
	Title:         "Contains is alias for Has",
	ExpectedInput: []string{"true", "false"},
}

// ==========================================================================
// Hashset — HasAll / HasAny
// ==========================================================================

var hashsetHasAllTrueTestCase = coretestcases.CaseV1{
	Title:         "HasAll true when all present",
	ExpectedInput: []string{"true"},
}

var hashsetHasAllFalseTestCase = coretestcases.CaseV1{
	Title:         "HasAll false when one missing",
	ExpectedInput: []string{"false"},
}

var hashsetHasAnyTrueTestCase = coretestcases.CaseV1{
	Title:         "HasAny true when one present",
	ExpectedInput: []string{"true"},
}

var hashsetHasAnyFalseTestCase = coretestcases.CaseV1{
	Title:         "HasAny false when none present",
	ExpectedInput: []string{"false"},
}

var hashsetHasAllEmptyArgsTestCase = coretestcases.CaseV1{
	Title:         "HasAll with empty args returns true",
	ExpectedInput: []string{"true"},
}

var hashsetHasAnyEmptyArgsTestCase = coretestcases.CaseV1{
	Title:         "HasAny with empty args returns false",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Hashset — IsEquals
// ==========================================================================

var hashsetIsEqualsSameItemsTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same items → true",
	ExpectedInput: []string{"true"},
}

var hashsetIsEqualsDifferentItemsTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different items → false",
	ExpectedInput: []string{"false"},
}

var hashsetIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length → false",
	ExpectedInput: []string{"false"},
}

var hashsetIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both nil → true",
	ExpectedInput: []string{"true"},
}

var hashsetIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals nil vs non-nil → false",
	ExpectedInput: []string{"false"},
}

var hashsetIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer → true",
	ExpectedInput: []string{"true"},
}

var hashsetIsEqualsBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both empty → true",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Hashset — Resize
// ==========================================================================

var hashsetResizeLargerTestCase = coretestcases.CaseV1{
	Title:         "Resize to larger capacity preserves items",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

var hashsetResizeSmallerTestCase = coretestcases.CaseV1{
	Title:         "Resize to smaller than length does nothing",
	ExpectedInput: []string{"3"},
}

// ==========================================================================
// Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

var hashsetOutputListTestCase = coretestcases.CaseV1{
	Title:         "List returns all items",
	ExpectedInput: []string{"3"},
}

var hashsetOutputListEmptyTestCase = coretestcases.CaseV1{
	Title:         "List on empty returns empty slice",
	ExpectedInput: []string{"0"},
}

var hashsetOutputListPtrTestCase = coretestcases.CaseV1{
	Title:         "ListPtr returns non-nil pointer",
	ExpectedInput: []string{"true"},
}

var hashsetOutputMapTestCase = coretestcases.CaseV1{
	Title:         "Map returns underlying map",
	ExpectedInput: []string{"3"},
}

var hashsetOutputCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection returns Collection[T] with same items",
	ExpectedInput: []string{"3"},
}

// ==========================================================================
// Hashset — Lock variants
// ==========================================================================

var hashsetLockAddContainsTestCase = coretestcases.CaseV1{
	Title:         "AddLock and ContainsLock work thread-safely",
	ExpectedInput: []string{"2", "true", "true", "false"},
}

var hashsetLockAddSliceTestCase = coretestcases.CaseV1{
	Title:         "AddSliceLock adds items thread-safely",
	ExpectedInput: []string{"3"},
}

var hashsetLockRemoveTestCase = coretestcases.CaseV1{
	Title:         "RemoveLock removes item thread-safely",
	ExpectedInput: []string{"true", "2", "false"},
}

var hashsetLockIsEmptyLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock and LengthLock return correct values",
	ExpectedInput: []string{"true", "0", "false", "2"},
}

// ==========================================================================
// Hashset — Constructors
// ==========================================================================

var hashsetConstructorEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyHashset creates empty set",
	ExpectedInput: []string{"0", "true"},
}

var hashsetConstructorNewCapTestCase = coretestcases.CaseV1{
	Title:         "NewHashset with capacity creates empty set",
	ExpectedInput: []string{"0", "true"},
}

var hashsetConstructorFromTestCase = coretestcases.CaseV1{
	Title:         "HashsetFrom creates populated set",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

var hashsetConstructorFromMapTestCase = coretestcases.CaseV1{
	Title:         "HashsetFromMap creates set from map",
	ExpectedInput: []string{"2", "true", "true"},
}

var hashsetConstructorHasItemsTestCase = coretestcases.CaseV1{
	Title:         "HasItems returns true for populated, false for empty",
	ExpectedInput: []string{"true", "false"},
}
