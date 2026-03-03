package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — ForEach
// ==========================================================================

var collectionForEachVisitsAllTestCase = coretestcases.CaseV1{
	Name:      "ForEach visits all items with correct indices",
	WantLines: []string{"5", "0:1", "4:5"},
}

var collectionForEachEmptyTestCase = coretestcases.CaseV1{
	Name:      "ForEach on empty collection does nothing",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — ForEachBreak
// ==========================================================================

var collectionForEachBreakStopsTestCase = coretestcases.CaseV1{
	Name:      "ForEachBreak stops at first match",
	WantLines: []string{"2"},
}

var collectionForEachBreakVisitsAllTestCase = coretestcases.CaseV1{
	Name:      "ForEachBreak visits all if no break",
	WantLines: []string{"5"},
}

// ==========================================================================
// Collection — SortFunc
// ==========================================================================

var collectionSortFuncAscTestCase = coretestcases.CaseV1{
	Name:      "SortFunc ascending",
	WantLines: []string{"1", "5"},
}

var collectionSortFuncDescTestCase = coretestcases.CaseV1{
	Name:      "SortFunc descending",
	WantLines: []string{"5", "1"},
}

var collectionSortFuncSingleTestCase = coretestcases.CaseV1{
	Name:      "SortFunc single element",
	WantLines: []string{"42", "42"},
}

// ==========================================================================
// Collection — AddIfMany
// ==========================================================================

var collectionAddIfManyTrueTestCase = coretestcases.CaseV1{
	Name:      "AddIfMany true adds all items",
	WantLines: []string{"3", "10", "30"},
}

var collectionAddIfManyFalseTestCase = coretestcases.CaseV1{
	Name:      "AddIfMany false adds nothing",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — AddFunc
// ==========================================================================

var collectionAddFuncTestCase = coretestcases.CaseV1{
	Name:      "AddFunc appends result of function",
	WantLines: []string{"1", "42"},
}

// ==========================================================================
// Collection — AddCollections (multiple)
// ==========================================================================

var collectionAddCollectionsMergeTestCase = coretestcases.CaseV1{
	Name:      "AddCollections merges multiple collections",
	WantLines: []string{"6", "1", "6"},
}

var collectionAddCollectionsNilTestCase = coretestcases.CaseV1{
	Name:      "AddCollections with nil collection skips it",
	WantLines: []string{"3", "1", "3"},
}

// ==========================================================================
// Collection — Clone edge cases
// ==========================================================================

var collectionCloneEmptyTestCase = coretestcases.CaseV1{
	Name:      "Clone empty returns empty",
	WantLines: []string{"0", "true"},
}

// ==========================================================================
// Collection — Skip/Take boundary
// ==========================================================================

var collectionSkipAllTestCase = coretestcases.CaseV1{
	Name:      "Skip all returns empty",
	WantLines: []string{"0"},
}

var collectionTakeMoreTestCase = coretestcases.CaseV1{
	Name:      "Take more than length returns all",
	WantLines: []string{"3"},
}

var collectionSkipZeroTakeZeroTestCase = coretestcases.CaseV1{
	Name:      "Skip 0 returns all, Take 0 returns empty",
	WantLines: []string{"3", "0"},
}

// ==========================================================================
// Collection — Filter edge cases
// ==========================================================================

var collectionFilterNoMatchTestCase = coretestcases.CaseV1{
	Name:      "Filter no match returns empty",
	WantLines: []string{"0", "true"},
}

var collectionFilterAllMatchTestCase = coretestcases.CaseV1{
	Name:      "Filter all match returns all",
	WantLines: []string{"3"},
}

var collectionFilterEmptyTestCase = coretestcases.CaseV1{
	Name:      "Filter empty collection returns empty",
	WantLines: []string{"0", "true"},
}

// ==========================================================================
// Collection — CountFunc edge cases
// ==========================================================================

var collectionCountFuncNoMatchTestCase = coretestcases.CaseV1{
	Name:      "CountFunc no match returns 0",
	WantLines: []string{"0"},
}

var collectionCountFuncEmptyTestCase = coretestcases.CaseV1{
	Name:      "CountFunc empty collection returns 0",
	WantLines: []string{"0"},
}

// ==========================================================================
// Collection — String output
// ==========================================================================

var collectionStringPopulatedTestCase = coretestcases.CaseV1{
	Name:      "String formats collection",
	WantLines: []string{"[1 2 3]"},
}

var collectionStringEmptyTestCase = coretestcases.CaseV1{
	Name:      "String empty collection",
	WantLines: []string{"[]"},
}

// ==========================================================================
// Collection — Lock variants
// ==========================================================================

var collectionLockVariantsTestCase = coretestcases.CaseV1{
	Name:      "Lock variants work correctly",
	WantLines: []string{"3", "false", "3"},
}

// ==========================================================================
// Collection — Metadata
// ==========================================================================

var collectionMetadataPopulatedTestCase = coretestcases.CaseV1{
	Name:      "Metadata methods on populated collection",
	WantLines: []string{"true", "true", "true", "false", "2", "3"},
}

var collectionMetadataEmptyTestCase = coretestcases.CaseV1{
	Name:      "Metadata methods on empty collection",
	WantLines: []string{"false", "false", "false", "-1", "0"},
}

// ==========================================================================
// Collection — RemoveAt single item
// ==========================================================================

var collectionRemoveAtSingleTestCase = coretestcases.CaseV1{
	Name:      "RemoveAt single item leaves empty collection",
	WantLines: []string{"true", "0", "true"},
}

// ==========================================================================
// Collection — AddCollection with nil/empty
// ==========================================================================

var collectionAddCollectionEmptyTestCase = coretestcases.CaseV1{
	Name:      "AddCollection with empty collection does not change length",
	WantLines: []string{"3"},
}

// ==========================================================================
// Hashmap — IsEquals updated
// ==========================================================================

var hashmapIsEqualsSameKeysTestCase = coretestcases.CaseV1{
	Name:      "IsEquals same keys → true",
	WantLines: []string{"true"},
}

var hashmapIsEqualsDiffKeysTestCase = coretestcases.CaseV1{
	Name:      "IsEquals same length different keys → false",
	WantLines: []string{"false"},
}

var hashmapIsEqualsDiffLengthTestCase = coretestcases.CaseV1{
	Name:      "IsEquals different length → false",
	WantLines: []string{"false"},
}

var hashmapIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Name:      "IsEquals both nil → true",
	WantLines: []string{"true"},
}

var hashmapIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Name:      "IsEquals nil vs non-nil → false",
	WantLines: []string{"false"},
}

var hashmapIsEqualsSamePtrTestCase = coretestcases.CaseV1{
	Name:      "IsEquals same pointer → true",
	WantLines: []string{"true"},
}

// ==========================================================================
// Collection — CollectionLenCap
// ==========================================================================

var collectionLenCapTestCase = coretestcases.CaseV1{
	Name:         "CollectionLenCap creates with pre-set length and capacity",
	ArrangeInput: args.Map{},
	WantLines:    []string{"3", "10", "0"},
}
