package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset — Add / AddBool edge cases
// ==========================================================================

var hashsetAddDuplicateTestCase = coretestcases.CaseV1{
	Title:         "Add duplicate does not increase length",
	ExpectedInput: args.Map{"length": 3},
}

var hashsetAddBoolTestCase = coretestcases.CaseV1{
	Title: "AddBool returns false for new, true for existing",
	ExpectedInput: args.Map{
		"firstExisted":  false,
		"secondExisted": true,
		"length":        1,
	},
}

var hashsetAddsVariadicTestCase = coretestcases.CaseV1{
	Title:         "Adds variadic adds all unique items",
	ExpectedInput: args.Map{"length": 4},
}

var hashsetAddSliceTestCase = coretestcases.CaseV1{
	Title: "AddSlice adds all items from slice",
	ExpectedInput: args.Map{
		"length": 3,
		"hasX":   true,
		"hasY":   true,
		"hasZ":   true,
	},
}

// ==========================================================================
// Hashset — AddIf / AddIfMany
// ==========================================================================

var hashsetAddIfTrueTestCase = coretestcases.CaseV1{
	Title: "AddIf true adds item",
	ExpectedInput: args.Map{
		"length":  1,
		"hasItem": true,
	},
}

var hashsetAddIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIf false skips item",
	ExpectedInput: args.Map{"length": 0},
}

var hashsetAddIfManyTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany true adds all",
	ExpectedInput: args.Map{"length": 3},
}

var hashsetAddIfManyFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany false adds none",
	ExpectedInput: args.Map{"length": 0},
}

// ==========================================================================
// Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

var hashsetMergeOtherSetTestCase = coretestcases.CaseV1{
	Title: "AddHashsetItems merges other set",
	ExpectedInput: args.Map{
		"length": 4,
		"has3":   true,
		"has4":   true,
	},
}

var hashsetMergeNilOtherTestCase = coretestcases.CaseV1{
	Title:         "AddHashsetItems with nil other does nothing",
	ExpectedInput: args.Map{"length": 2},
}

var hashsetMergeEmptyOtherTestCase = coretestcases.CaseV1{
	Title:         "AddHashsetItems with empty other does nothing",
	ExpectedInput: args.Map{"length": 2},
}

var hashsetAddItemsMapTestCase = coretestcases.CaseV1{
	Title: "AddItemsMap adds only true entries",
	ExpectedInput: args.Map{
		"length":  2,
		"hasYes":  true,
		"hasNope": false,
	},
}

// ==========================================================================
// Hashset — Remove edge cases
// ==========================================================================

var hashsetRemoveExistingTestCase = coretestcases.CaseV1{
	Title: "Remove existing returns true and decreases length",
	ExpectedInput: args.Map{
		"existed":  true,
		"length":   2,
		"stillHas": false,
	},
}

var hashsetRemoveNonExistingTestCase = coretestcases.CaseV1{
	Title: "Remove non-existing returns false",
	ExpectedInput: args.Map{
		"existed": false,
		"length":  3,
	},
}

// ==========================================================================
// Hashset — Has / Contains
// ==========================================================================

var hashsetHasTestCase = coretestcases.CaseV1{
	Title: "Has returns true for existing, false for missing",
	ExpectedInput: args.Map{
		"hasExisting": true,
		"hasMissing":  false,
	},
}

var hashsetContainsAliasTestCase = coretestcases.CaseV1{
	Title: "Contains is alias for Has",
	ExpectedInput: args.Map{
		"containsExisting": true,
		"containsMissing":  false,
	},
}

// ==========================================================================
// Hashset — HasAll / HasAny
// ==========================================================================

var hashsetHasAllTrueTestCase = coretestcases.CaseV1{
	Title:         "HasAll true when all present",
	ExpectedInput: args.Map{"hasAll": true},
}

var hashsetHasAllFalseTestCase = coretestcases.CaseV1{
	Title:         "HasAll false when one missing",
	ExpectedInput: args.Map{"hasAll": false},
}

var hashsetHasAnyTrueTestCase = coretestcases.CaseV1{
	Title:         "HasAny true when one present",
	ExpectedInput: args.Map{"hasAny": true},
}

var hashsetHasAnyFalseTestCase = coretestcases.CaseV1{
	Title:         "HasAny false when none present",
	ExpectedInput: args.Map{"hasAny": false},
}

var hashsetHasAllEmptyArgsTestCase = coretestcases.CaseV1{
	Title:         "HasAll with empty args returns true",
	ExpectedInput: args.Map{"hasAll": true},
}

var hashsetHasAnyEmptyArgsTestCase = coretestcases.CaseV1{
	Title:         "HasAny with empty args returns false",
	ExpectedInput: args.Map{"hasAny": false},
}

// ==========================================================================
// Hashset — IsEquals
// ==========================================================================

var hashsetIsEqualsSameItemsTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same items → true",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashsetIsEqualsDifferentItemsTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different items → false",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashsetIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length → false",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashsetIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both nil → true",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashsetIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals nil vs non-nil → false",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashsetIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer → true",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashsetIsEqualsBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both empty → true",
	ExpectedInput: args.Map{"isEquals": true},
}

// ==========================================================================
// Hashset — Resize
// ==========================================================================

var hashsetResizeLargerTestCase = coretestcases.CaseV1{
	Title: "Resize to larger capacity preserves items",
	ExpectedInput: args.Map{
		"length": 3,
		"has1":   true,
		"has2":   true,
		"has3":   true,
	},
}

var hashsetResizeSmallerTestCase = coretestcases.CaseV1{
	Title:         "Resize to smaller than length does nothing",
	ExpectedInput: args.Map{"length": 3},
}

// ==========================================================================
// Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

var hashsetOutputListTestCase = coretestcases.CaseV1{
	Title:         "List returns all items",
	ExpectedInput: args.Map{"listLen": 3},
}

var hashsetOutputListEmptyTestCase = coretestcases.CaseV1{
	Title:         "List on empty returns empty slice",
	ExpectedInput: args.Map{"listLen": 0},
}

var hashsetOutputListPtrTestCase = coretestcases.CaseV1{
	Title:         "ListPtr returns non-nil pointer",
	ExpectedInput: args.Map{"isNotNil": true},
}

var hashsetOutputMapTestCase = coretestcases.CaseV1{
	Title:         "Map returns underlying map",
	ExpectedInput: args.Map{"mapLen": 3},
}

var hashsetOutputCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection returns Collection[T] with same items",
	ExpectedInput: args.Map{"collectionLen": 3},
}

// ==========================================================================
// Hashset — Lock variants
// ==========================================================================

var hashsetLockAddContainsTestCase = coretestcases.CaseV1{
	Title: "AddLock and ContainsLock work thread-safely",
	ExpectedInput: args.Map{
		"length":    2,
		"containsA": true,
		"containsB": true,
		"containsZ": false,
	},
}

var hashsetLockAddSliceTestCase = coretestcases.CaseV1{
	Title:         "AddSliceLock adds items thread-safely",
	ExpectedInput: args.Map{"length": 3},
}

var hashsetLockRemoveTestCase = coretestcases.CaseV1{
	Title: "RemoveLock removes item thread-safely",
	ExpectedInput: args.Map{
		"existed":  true,
		"length":   2,
		"stillHas": false,
	},
}

var hashsetLockIsEmptyLengthTestCase = coretestcases.CaseV1{
	Title: "IsEmptyLock and LengthLock return correct values",
	ExpectedInput: args.Map{
		"emptyBefore":  true,
		"lengthBefore": 0,
		"emptyAfter":   false,
		"lengthAfter":  2,
	},
}

// ==========================================================================
// Hashset — Constructors
// ==========================================================================

var hashsetConstructorEmptyTestCase = coretestcases.CaseV1{
	Title: "EmptyHashset creates empty set",
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}

var hashsetConstructorNewCapTestCase = coretestcases.CaseV1{
	Title: "NewHashset with capacity creates empty set",
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}

var hashsetConstructorFromTestCase = coretestcases.CaseV1{
	Title: "HashsetFrom creates populated set",
	ExpectedInput: args.Map{
		"length": 3,
		"hasA":   true,
		"hasB":   true,
		"hasC":   true,
	},
}

var hashsetConstructorFromMapTestCase = coretestcases.CaseV1{
	Title: "HashsetFromMap creates set from map",
	ExpectedInput: args.Map{
		"length": 2,
		"has10":  true,
		"has20":  true,
	},
}

var hashsetConstructorHasItemsTestCase = coretestcases.CaseV1{
	Title: "HasItems returns true for populated, false for empty",
	ExpectedInput: args.Map{
		"populatedHasItems": true,
		"emptyHasItems":     false,
	},
}
