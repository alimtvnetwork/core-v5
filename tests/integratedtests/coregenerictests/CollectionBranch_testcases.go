package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — ForEach
// ==========================================================================

var collectionForEachVisitsAllTestCase = coretestcases.CaseV1{
	Title: "ForEach visits all items with correct indices",
	ExpectedInput: args.Map{
		"visited":    5,
		"firstEntry": "0:1",
		"lastEntry":  "4:5",
	},
}

var collectionForEachEmptyTestCase = coretestcases.CaseV1{
	Title:         "ForEach on empty collection does nothing",
	ExpectedInput: args.Map{"visited": 0},
}

// ==========================================================================
// Collection — ForEachBreak
// ==========================================================================

var collectionForEachBreakStopsTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops at first match",
	ExpectedInput: args.Map{"visited": 3},
}

var collectionForEachBreakVisitsAllTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak visits all if no break",
	ExpectedInput: args.Map{"visited": 5},
}

// ==========================================================================
// Collection — SortFunc
// ==========================================================================

var collectionSortFuncAscTestCase = coretestcases.CaseV1{
	Title: "SortFunc ascending",
	ExpectedInput: args.Map{
		"first": 1,
		"last":  5,
	},
}

var collectionSortFuncDescTestCase = coretestcases.CaseV1{
	Title: "SortFunc descending",
	ExpectedInput: args.Map{
		"first": 5,
		"last":  1,
	},
}

var collectionSortFuncSingleTestCase = coretestcases.CaseV1{
	Title: "SortFunc single element",
	ExpectedInput: args.Map{
		"first": 42,
		"last":  42,
	},
}

// ==========================================================================
// Collection — AddIfMany
// ==========================================================================

var collectionAddIfManyTrueTestCase = coretestcases.CaseV1{
	Title: "AddIfMany true adds all items",
	ExpectedInput: args.Map{
		"length": 3,
		"first":  10,
		"last":   30,
	},
}

var collectionAddIfManyFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIfMany false adds nothing",
	ExpectedInput: args.Map{"length": 0},
}

// ==========================================================================
// Collection — AddFunc
// ==========================================================================

var collectionAddFuncTestCase = coretestcases.CaseV1{
	Title: "AddFunc appends result of function",
	ExpectedInput: args.Map{
		"length": 1,
		"first":  42,
	},
}

// ==========================================================================
// Collection — AddCollections (multiple)
// ==========================================================================

var collectionAddCollectionsMergeTestCase = coretestcases.CaseV1{
	Title: "AddCollections merges multiple collections",
	ExpectedInput: args.Map{
		"length": 6,
		"first":  1,
		"last":   6,
	},
}

var collectionAddCollectionsNilTestCase = coretestcases.CaseV1{
	Title: "AddCollections with nil collection skips it",
	ExpectedInput: args.Map{
		"length": 3,
		"first":  1,
		"last":   3,
	},
}

// ==========================================================================
// Collection — Clone edge cases
// ==========================================================================

var collectionCloneEmptyTestCase = coretestcases.CaseV1{
	Title: "Clone empty returns empty",
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}

// ==========================================================================
// Collection — Skip/Take boundary
// ==========================================================================

var collectionSkipAllTestCase = coretestcases.CaseV1{
	Title:         "Skip all returns empty",
	ExpectedInput: args.Map{"length": 0},
}

var collectionTakeMoreTestCase = coretestcases.CaseV1{
	Title:         "Take more than length returns all",
	ExpectedInput: args.Map{"length": 3},
}

var collectionSkipZeroTakeZeroTestCase = coretestcases.CaseV1{
	Title: "Skip 0 returns all, Take 0 returns empty",
	ExpectedInput: args.Map{
		"skipLength": 3,
		"takeLength": 0,
	},
}

// ==========================================================================
// Collection — Filter edge cases
// ==========================================================================

var collectionFilterNoMatchTestCase = coretestcases.CaseV1{
	Title: "Filter no match returns empty",
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}

var collectionFilterAllMatchTestCase = coretestcases.CaseV1{
	Title:         "Filter all match returns all",
	ExpectedInput: args.Map{"length": 3},
}

var collectionFilterEmptyTestCase = coretestcases.CaseV1{
	Title: "Filter empty collection returns empty",
	ExpectedInput: args.Map{
		"length":  0,
		"isEmpty": true,
	},
}

// ==========================================================================
// Collection — CountFunc edge cases
// ==========================================================================

var collectionCountFuncNoMatchTestCase = coretestcases.CaseV1{
	Title:         "CountFunc no match returns 0",
	ExpectedInput: args.Map{"count": 0},
}

var collectionCountFuncEmptyTestCase = coretestcases.CaseV1{
	Title:         "CountFunc empty collection returns 0",
	ExpectedInput: args.Map{"count": 0},
}

// ==========================================================================
// Collection — String output
// ==========================================================================

var collectionStringPopulatedTestCase = coretestcases.CaseV1{
	Title:         "String formats collection",
	ExpectedInput: args.Map{"result": "[1 2 3]"},
}

var collectionStringEmptyTestCase = coretestcases.CaseV1{
	Title:         "String empty collection",
	ExpectedInput: args.Map{"result": "[]"},
}

// ==========================================================================
// Collection — Lock variants
// ==========================================================================

var collectionLockVariantsTestCase = coretestcases.CaseV1{
	Title: "Lock variants work correctly",
	ExpectedInput: args.Map{
		"lengthLock":  3,
		"isEmptyLock": false,
		"length":      3,
	},
}

// ==========================================================================
// Collection — Metadata
// ==========================================================================

var collectionMetadataPopulatedTestCase = coretestcases.CaseV1{
	Title: "Metadata methods on populated collection",
	ExpectedInput: args.Map{
		"hasAnyItem": true,
		"hasItems":   true,
		"hasIndex2":  true,
		"hasIndex5":  false,
		"lastIndex":  2,
		"count":      3,
	},
}

var collectionMetadataEmptyTestCase = coretestcases.CaseV1{
	Title: "Metadata methods on empty collection",
	ExpectedInput: args.Map{
		"hasAnyItem": false,
		"hasItems":   false,
		"hasIndex0":  false,
		"lastIndex":  -1,
		"count":      0,
	},
}

// ==========================================================================
// Collection — RemoveAt single item
// ==========================================================================

var collectionRemoveAtSingleTestCase = coretestcases.CaseV1{
	Title: "RemoveAt single item leaves empty collection",
	ExpectedInput: args.Map{
		"removed": true,
		"length":  0,
		"isEmpty": true,
	},
}

// ==========================================================================
// Collection — AddCollection with nil/empty
// ==========================================================================

var collectionAddCollectionEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddCollection with empty collection does not change length",
	ExpectedInput: args.Map{"length": 3},
}

// ==========================================================================
// Hashmap — IsEquals updated
// ==========================================================================

var hashmapIsEqualsSameKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same keys → true",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashmapIsEqualsDiffKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same length different keys → false",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashmapIsEqualsDiffLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length → false",
	ExpectedInput: args.Map{"isEquals": false},
}

// Removed: hashmapIsEqualsBothNilTestCase — declared in Hashmap_testcases.go

var hashmapIsEqualsNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals nil vs non-nil → false",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashmapIsEqualsSamePtrTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer → true",
	ExpectedInput: args.Map{"isEquals": true},
}

// ==========================================================================
// Collection — CollectionLenCap
// ==========================================================================

var collectionLenCapTestCase = coretestcases.CaseV1{
	Title:        "CollectionLenCap creates with pre-set length and capacity",
	ArrangeInput: args.Map{},
	ExpectedInput: args.Map{
		"length":   3,
		"capacity": 10,
		"first":    0,
	},
}
