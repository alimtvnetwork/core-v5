package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset — Add / AddBool edge cases
// ==========================================================================

var hashsetAddDuplicateTestCase = coretestcases.CaseV1{
	Name:      "Add duplicate does not increase length",
	WantLines: []string{"3"},
}

var hashsetAddBoolTestCase = coretestcases.CaseV1{
	Name:      "AddBool returns false for new, true for existing",
	WantLines: []string{"false", "true", "1"},
}

var hashsetAddsVariadicTestCase = coretestcases.CaseV1{
	Name:      "Adds variadic adds all unique items",
	WantLines: []string{"4"},
}

var hashsetAddSliceTestCase = coretestcases.CaseV1{
	Name:      "AddSlice adds all items from slice",
	WantLines: []string{"3", "true", "true", "true"},
}

// ==========================================================================
// Hashset — AddIf / AddIfMany
// ==========================================================================

var hashsetAddIfTrueTestCase = coretestcases.CaseV1{
	Name:      "AddIf true adds item",
	WantLines: []string{"1", "true"},
}

var hashsetAddIfFalseTestCase = coretestcases.CaseV1{
	Name:      "AddIf false skips item",
	WantLines: []string{"0"},
}

var hashsetAddIfManyTrueTestCase = coretestcases.CaseV1{
	Name:      "AddIfMany true adds all",
	WantLines: []string{"3"},
}

var hashsetAddIfManyFalseTestCase = coretestcases.CaseV1{
	Name:      "AddIfMany false adds none",
	WantLines: []string{"0"},
}

// ==========================================================================
// Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

var hashsetMergeOtherSetTestCase = coretestcases.CaseV1{
	Name:      "AddHashsetItems merges other set",
	WantLines: []string{"4", "true", "true"},
}

var hashsetMergeNilOtherTestCase = coretestcases.CaseV1{
	Name:      "AddHashsetItems with nil other does nothing",
	WantLines: []string{"2"},
}

var hashsetMergeEmptyOtherTestCase = coretestcases.CaseV1{
	Name:      "AddHashsetItems with empty other does nothing",
	WantLines: []string{"2"},
}

var hashsetAddItemsMapTestCase = coretestcases.CaseV1{
	Name:      "AddItemsMap adds only true entries",
	WantLines: []string{"2", "true", "false"},
}

// ==========================================================================
// Hashset — Remove edge cases
// ==========================================================================

var hashsetRemoveExistingTestCase = coretestcases.CaseV1{
	Name:      "Remove existing returns true and decreases length",
	WantLines: []string{"true", "2", "false"},
}

var hashsetRemoveNonExistingTestCase = coretestcases.CaseV1{
	Name:      "Remove non-existing returns false",
	WantLines: []string{"false", "3"},
}

// ==========================================================================
// Hashset — Has / Contains
// ==========================================================================

var hashsetHasTestCase = coretestcases.CaseV1{
	Name:      "Has returns true for existing, false for missing",
	WantLines: []string{"true", "false"},
}

var hashsetContainsAliasTestCase = coretestcases.CaseV1{
	Name:      "Contains is alias for Has",
	WantLines: []string{"true", "false"},
}

// ==========================================================================
// Hashset — HasAll / HasAny
// ==========================================================================

var hashsetHasAllTrueTestCase = coretestcases.CaseV1{
	Name:      "HasAll true when all present",
	WantLines: []string{"true"},
}

var hashsetHasAllFalseTestCase = coretestcases.CaseV1{
	Name:      "HasAll false when one missing",
	WantLines: []string{"false"},
}

var hashsetHasAnyTrueTestCase = coretestcases.CaseV1{
	Name:      "HasAny true when one present",
	WantLines: []string{"true"},
}

var hashsetHasAnyFalseTestCase = coretestcases.CaseV1{
	Name:      "HasAny false when none present",
	WantLines: []string{"false"},
}

var hashsetHasAllEmptyArgsTestCase = coretestcases.CaseV1{
	Name:      "HasAll with empty args returns true",
	WantLines: []string{"true"},
}

var hashsetHasAnyEmptyArgsTestCase = coretestcases.CaseV1{
	Name:      "HasAny with empty args returns false",
	WantLines: []string{"false"},
}

// ==========================================================================
// Hashset — IsEquals
// ==========================================================================

var hashsetIsEqualsSameItemsTestCase = coretestcases.CaseV1{
	Name:      "IsEquals same items → true",
	WantLines: []string{"true"},
}

var hashsetIsEqualsDifferentItemsTestCase = coretestcases.CaseV1{
	Name:      "IsEquals different items → false",
	WantLines: []string{"false"},
}

var hashsetIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Name:      "IsEquals different length → false",
	WantLines: []string{"false"},
}

var hashsetIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Name:      "IsEquals both nil → true",
	WantLines: []string{"true"},
}

var hashsetIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Name:      "IsEquals nil vs non-nil → false",
	WantLines: []string{"false"},
}

var hashsetIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Name:      "IsEquals same pointer → true",
	WantLines: []string{"true"},
}

var hashsetIsEqualsBothEmptyTestCase = coretestcases.CaseV1{
	Name:      "IsEquals both empty → true",
	WantLines: []string{"true"},
}

// ==========================================================================
// Hashset — Resize
// ==========================================================================

var hashsetResizeLargerTestCase = coretestcases.CaseV1{
	Name:      "Resize to larger capacity preserves items",
	WantLines: []string{"3", "true", "true", "true"},
}

var hashsetResizeSmallerTestCase = coretestcases.CaseV1{
	Name:      "Resize to smaller than length does nothing",
	WantLines: []string{"3"},
}

// ==========================================================================
// Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

var hashsetOutputListTestCase = coretestcases.CaseV1{
	Name:      "List returns all items",
	WantLines: []string{"3"},
}

var hashsetOutputListEmptyTestCase = coretestcases.CaseV1{
	Name:      "List on empty returns empty slice",
	WantLines: []string{"0"},
}

var hashsetOutputListPtrTestCase = coretestcases.CaseV1{
	Name:      "ListPtr returns non-nil pointer",
	WantLines: []string{"true"},
}

var hashsetOutputMapTestCase = coretestcases.CaseV1{
	Name:      "Map returns underlying map",
	WantLines: []string{"3"},
}

var hashsetOutputCollectionTestCase = coretestcases.CaseV1{
	Name:      "Collection returns Collection[T] with same items",
	WantLines: []string{"3"},
}

// ==========================================================================
// Hashset — Lock variants
// ==========================================================================

var hashsetLockAddContainsTestCase = coretestcases.CaseV1{
	Name:      "AddLock and ContainsLock work thread-safely",
	WantLines: []string{"2", "true", "true", "false"},
}

var hashsetLockAddSliceTestCase = coretestcases.CaseV1{
	Name:      "AddSliceLock adds items thread-safely",
	WantLines: []string{"3"},
}

var hashsetLockRemoveTestCase = coretestcases.CaseV1{
	Name:      "RemoveLock removes item thread-safely",
	WantLines: []string{"true", "2", "false"},
}

var hashsetLockIsEmptyLengthTestCase = coretestcases.CaseV1{
	Name:      "IsEmptyLock and LengthLock return correct values",
	WantLines: []string{"true", "0", "false", "2"},
}

// ==========================================================================
// Hashset — Constructors
// ==========================================================================

var hashsetConstructorEmptyTestCase = coretestcases.CaseV1{
	Name:      "EmptyHashset creates empty set",
	WantLines: []string{"0", "true"},
}

var hashsetConstructorNewCapTestCase = coretestcases.CaseV1{
	Name:      "NewHashset with capacity creates empty set",
	WantLines: []string{"0", "true"},
}

var hashsetConstructorFromTestCase = coretestcases.CaseV1{
	Name:      "HashsetFrom creates populated set",
	WantLines: []string{"3", "true", "true", "true"},
}

var hashsetConstructorFromMapTestCase = coretestcases.CaseV1{
	Name:      "HashsetFromMap creates set from map",
	WantLines: []string{"2", "true", "true"},
}

var hashsetConstructorHasItemsTestCase = coretestcases.CaseV1{
	Name:      "HasItems returns true for populated, false for empty",
	WantLines: []string{"true", "false"},
}
