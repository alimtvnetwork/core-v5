package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — ForEach
// ==========================================================================

var collectionForEachVisitsAllTestCase = coretestcases.CaseV1{
	Title:         "ForEach visits all items with correct indices",
	ExpectedInput: []string{"5", "0:1", "4:5"},
}

var collectionForEachEmptyTestCase = coretestcases.CaseV1{
	Title:         "ForEach on empty collection does nothing",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — ForEachBreak
// ==========================================================================

var collectionForEachBreakStopsTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops at first match",
	ExpectedInput: []string{"2"},
}

var collectionForEachBreakVisitsAllTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak visits all if no break",
	ExpectedInput: []string{"5"},
}

// ==========================================================================
// Collection — SortFunc
// ==========================================================================

var collectionSortFuncAscTestCase = coretestcases.CaseV1{
	Title:         "SortFunc ascending",
	ExpectedInput: []string{"1", "5"},
}

var collectionSortFuncDescTestCase = coretestcases.CaseV1{
	Title:         "SortFunc descending",
	ExpectedInput: []string{"5", "1"},
}

var collectionSortFuncSingleTestCase = coretestcases.CaseV1{
	Title:         "SortFunc single element",
	ExpectedInput: []string{"42", "42"},
}

// ==========================================================================
// Collection — AddIfMany
// ==========================================================================

var collectionAddIfManyTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany true adds all items",
	ExpectedInput: []string{"3", "10", "30"},
}

var collectionAddIfManyFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany false adds nothing",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — AddFunc
// ==========================================================================

var collectionAddFuncTestCase = coretestcases.CaseV1{
	Title:         "AddFunc appends result of function",
	ExpectedInput: []string{"1", "42"},
}

// ==========================================================================
// Collection — AddCollections (multiple)
// ==========================================================================

var collectionAddCollectionsMergeTestCase = coretestcases.CaseV1{
	Title:         "AddCollections merges multiple collections",
	ExpectedInput: []string{"6", "1", "6"},
}

var collectionAddCollectionsNilTestCase = coretestcases.CaseV1{
	Title:         "AddCollections with nil collection skips it",
	ExpectedInput: []string{"3", "1", "3"},
}

// ==========================================================================
// Collection — Clone edge cases
// ==========================================================================

var collectionCloneEmptyTestCase = coretestcases.CaseV1{
	Title:         "Clone empty returns empty",
	ExpectedInput: []string{"0", "true"},
}

// ==========================================================================
// Collection — Skip/Take boundary
// ==========================================================================

var collectionSkipAllTestCase = coretestcases.CaseV1{
	Title:         "Skip all returns empty",
	ExpectedInput: []string{"0"},
}

var collectionTakeMoreTestCase = coretestcases.CaseV1{
	Title:         "Take more than length returns all",
	ExpectedInput: []string{"3"},
}

var collectionSkipZeroTakeZeroTestCase = coretestcases.CaseV1{
	Title:         "Skip 0 returns all, Take 0 returns empty",
	ExpectedInput: []string{"3", "0"},
}

// ==========================================================================
// Collection — Filter edge cases
// ==========================================================================

var collectionFilterNoMatchTestCase = coretestcases.CaseV1{
	Title:         "Filter no match returns empty",
	ExpectedInput: []string{"0", "true"},
}

var collectionFilterAllMatchTestCase = coretestcases.CaseV1{
	Title:         "Filter all match returns all",
	ExpectedInput: []string{"3"},
}

var collectionFilterEmptyTestCase = coretestcases.CaseV1{
	Title:         "Filter empty collection returns empty",
	ExpectedInput: []string{"0", "true"},
}

// ==========================================================================
// Collection — CountFunc edge cases
// ==========================================================================

var collectionCountFuncNoMatchTestCase = coretestcases.CaseV1{
	Title:         "CountFunc no match returns 0",
	ExpectedInput: []string{"0"},
}

var collectionCountFuncEmptyTestCase = coretestcases.CaseV1{
	Title:         "CountFunc empty collection returns 0",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Collection — String output
// ==========================================================================

var collectionStringPopulatedTestCase = coretestcases.CaseV1{
	Title:         "String formats collection",
	ExpectedInput: []string{"[1 2 3]"},
}

var collectionStringEmptyTestCase = coretestcases.CaseV1{
	Title:         "String empty collection",
	ExpectedInput: []string{"[]"},
}

// ==========================================================================
// Collection — Lock variants
// ==========================================================================

var collectionLockVariantsTestCase = coretestcases.CaseV1{
	Title:         "Lock variants work correctly",
	ExpectedInput: []string{"3", "false", "3"},
}

// ==========================================================================
// Collection — Metadata
// ==========================================================================

var collectionMetadataPopulatedTestCase = coretestcases.CaseV1{
	Title:         "Metadata methods on populated collection",
	ExpectedInput: []string{"true", "true", "true", "false", "2", "3"},
}

var collectionMetadataEmptyTestCase = coretestcases.CaseV1{
	Title:         "Metadata methods on empty collection",
	ExpectedInput: []string{"false", "false", "false", "-1", "0"},
}

// ==========================================================================
// Collection — RemoveAt single item
// ==========================================================================

var collectionRemoveAtSingleTestCase = coretestcases.CaseV1{
	Title:         "RemoveAt single item leaves empty collection",
	ExpectedInput: []string{"true", "0", "true"},
}

// ==========================================================================
// Collection — AddCollection with nil/empty
// ==========================================================================

var collectionAddCollectionEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddCollection with empty collection does not change length",
	ExpectedInput: []string{"3"},
}

// ==========================================================================
// Hashmap — IsEquals updated
// ==========================================================================

var hashmapIsEqualsSameKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same keys → true",
	ExpectedInput: []string{"true"},
}

var hashmapIsEqualsDiffKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same length different keys → false",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualsDiffLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length → false",
	ExpectedInput: []string{"false"},
}

// Removed: hashmapIsEqualsBothNilTestCase — declared in Hashmap_testcases.go

var hashmapIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals nil vs non-nil → false",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualsSamePtrTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer → true",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Collection — CollectionLenCap
// ==========================================================================

var collectionLenCapTestCase = coretestcases.CaseV1{
	Title:         "CollectionLenCap creates with pre-set length and capacity",
	ArrangeInput:  args.Map{},
	ExpectedInput: []string{"3", "10", "0"},
}
