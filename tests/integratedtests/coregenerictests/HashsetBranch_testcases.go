package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset — Add / AddBool edge cases
// ==========================================================================

var hashsetAddEdgeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Add duplicate does not increase length",
		WantLines: []string{"3"},
	},
	{
		Name:      "AddBool returns false for new, true for existing",
		WantLines: []string{"false", "true", "1"},
	},
	{
		Name:      "Adds variadic adds all unique items",
		WantLines: []string{"4"},
	},
	{
		Name:      "AddSlice adds all items from slice",
		WantLines: []string{"3", "true", "true", "true"},
	},
}

// ==========================================================================
// Hashset — AddIf / AddIfMany
// ==========================================================================

var hashsetAddIfTestCases = []coretestcases.CaseV1{
	{
		Name:      "AddIf true adds item",
		WantLines: []string{"1", "true"},
	},
	{
		Name:      "AddIf false skips item",
		WantLines: []string{"0"},
	},
	{
		Name:      "AddIfMany true adds all",
		WantLines: []string{"3"},
	},
	{
		Name:      "AddIfMany false adds none",
		WantLines: []string{"0"},
	},
}

// ==========================================================================
// Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

var hashsetMergeTestCases = []coretestcases.CaseV1{
	{
		Name:      "AddHashsetItems merges other set",
		WantLines: []string{"4", "true", "true"},
	},
	{
		Name:      "AddHashsetItems with nil other does nothing",
		WantLines: []string{"2"},
	},
	{
		Name:      "AddHashsetItems with empty other does nothing",
		WantLines: []string{"2"},
	},
	{
		Name:      "AddItemsMap adds only true entries",
		WantLines: []string{"2", "true", "false"},
	},
}

// ==========================================================================
// Hashset — Remove edge cases
// ==========================================================================

var hashsetRemoveEdgeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Remove existing returns true and decreases length",
		WantLines: []string{"true", "2", "false"},
	},
	{
		Name:      "Remove non-existing returns false",
		WantLines: []string{"false", "3"},
	},
}

// ==========================================================================
// Hashset — Has / Contains
// ==========================================================================

var hashsetHasEdgeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Has returns true for existing, false for missing",
		WantLines: []string{"true", "false"},
	},
	{
		Name:      "Contains is alias for Has",
		WantLines: []string{"true", "false"},
	},
}

// ==========================================================================
// Hashset — HasAll / HasAny
// ==========================================================================

var hashsetHasAllAnyTestCases = []coretestcases.CaseV1{
	{
		Name:      "HasAll true when all present",
		WantLines: []string{"true"},
	},
	{
		Name:      "HasAll false when one missing",
		WantLines: []string{"false"},
	},
	{
		Name:      "HasAny true when one present",
		WantLines: []string{"true"},
	},
	{
		Name:      "HasAny false when none present",
		WantLines: []string{"false"},
	},
	{
		Name:      "HasAll with empty args returns true",
		WantLines: []string{"true"},
	},
	{
		Name:      "HasAny with empty args returns false",
		WantLines: []string{"false"},
	},
}

// ==========================================================================
// Hashset — IsEquals
// ==========================================================================

var hashsetIsEqualsTestCases = []coretestcases.CaseV1{
	{
		Name:      "IsEquals same items → true",
		WantLines: []string{"true"},
	},
	{
		Name:      "IsEquals different items → false",
		WantLines: []string{"false"},
	},
	{
		Name:      "IsEquals different length → false",
		WantLines: []string{"false"},
	},
	{
		Name:      "IsEquals both nil → true",
		WantLines: []string{"true"},
	},
	{
		Name:      "IsEquals nil vs non-nil → false",
		WantLines: []string{"false"},
	},
	{
		Name:      "IsEquals same pointer → true",
		WantLines: []string{"true"},
	},
	{
		Name:      "IsEquals both empty → true",
		WantLines: []string{"true"},
	},
}

// ==========================================================================
// Hashset — Resize
// ==========================================================================

var hashsetResizeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Resize to larger capacity preserves items",
		WantLines: []string{"3", "true", "true", "true"},
	},
	{
		Name:      "Resize to smaller than length does nothing",
		WantLines: []string{"3"},
	},
}

// ==========================================================================
// Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

var hashsetOutputTestCases = []coretestcases.CaseV1{
	{
		Name:      "List returns all items",
		WantLines: []string{"3"},
	},
	{
		Name:      "List on empty returns empty slice",
		WantLines: []string{"0"},
	},
	{
		Name:      "ListPtr returns non-nil pointer",
		WantLines: []string{"true"},
	},
	{
		Name:      "Map returns underlying map",
		WantLines: []string{"3"},
	},
	{
		Name:      "Collection returns Collection[T] with same items",
		WantLines: []string{"3"},
	},
}

// ==========================================================================
// Hashset — Lock variants
// ==========================================================================

var hashsetLockTestCases = []coretestcases.CaseV1{
	{
		Name:      "AddLock and ContainsLock work thread-safely",
		WantLines: []string{"2", "true", "true", "false"},
	},
	{
		Name:      "AddSliceLock adds items thread-safely",
		WantLines: []string{"3"},
	},
	{
		Name:      "RemoveLock removes item thread-safely",
		WantLines: []string{"true", "2", "false"},
	},
	{
		Name:      "IsEmptyLock and LengthLock return correct values",
		WantLines: []string{"true", "0", "false", "2"},
	},
}

// ==========================================================================
// Hashset — Constructors
// ==========================================================================

var hashsetConstructorTestCases = []coretestcases.CaseV1{
	{
		Name:      "EmptyHashset creates empty set",
		WantLines: []string{"0", "true"},
	},
	{
		Name:      "NewHashset with capacity creates empty set",
		WantLines: []string{"0", "true"},
	},
	{
		Name:      "HashsetFrom creates populated set",
		WantLines: []string{"3", "true", "true", "true"},
	},
	{
		Name:      "HashsetFromMap creates set from map",
		WantLines: []string{"2", "true", "true"},
	},
	{
		Name:      "HasItems returns true for populated, false for empty",
		WantLines: []string{"true", "false"},
	},
}
