package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Collection — RemoveAt edge cases
// ==========================================================================

var collectionRemoveAtMiddleTestCase = coretestcases.CaseV1{
	Title: "RemoveAt middle index",
	ExpectedInput: args.Map{
		"removed": true,
		"length":  4,
		"first":   1,
		"last":    5,
	},
}

var collectionRemoveAtFirstTestCase = coretestcases.CaseV1{
	Title: "RemoveAt first index",
	ExpectedInput: args.Map{
		"removed": true,
		"length":  4,
		"first":   2,
		"last":    5,
	},
}

var collectionRemoveAtLastTestCase = coretestcases.CaseV1{
	Title: "RemoveAt last index",
	ExpectedInput: args.Map{
		"removed": true,
		"length":  4,
		"first":   1,
		"last":    4,
	},
}

var collectionRemoveAtNegativeTestCase = coretestcases.CaseV1{
	Title: "RemoveAt negative index returns false",
	ExpectedInput: args.Map{
		"removed": false,
		"length":  5,
	},
}

var collectionRemoveAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "RemoveAt out-of-bounds index returns false",
	ExpectedInput: args.Map{
		"removed": false,
		"length":  5,
	},
}

var collectionRemoveAtEmptyTestCase = coretestcases.CaseV1{
	Title: "RemoveAt on empty collection returns false",
	ExpectedInput: args.Map{
		"removed": false,
		"length":  0,
	},
}

// ==========================================================================
// Collection — Reverse
// ==========================================================================

var collectionReversePopulatedTestCase = coretestcases.CaseV1{
	Title: "Reverse populated collection",
	ExpectedInput: args.Map{
		"length": 5,
		"first":  5,
		"last":   1,
	},
}

var collectionReverseSingleTestCase = coretestcases.CaseV1{
	Title:         "Reverse single element",
	ExpectedInput: args.Map{"first": 42},
}

var collectionReverseEmptyTestCase = coretestcases.CaseV1{
	Title:         "Reverse empty collection",
	ExpectedInput: args.Map{"length": 0},
}

// ==========================================================================
// Collection — FirstOrDefault / LastOrDefault
// ==========================================================================

var collectionFirstOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault on populated returns first",
	ExpectedInput: args.Map{"result": 10},
}

var collectionFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault on empty returns zero",
	ExpectedInput: args.Map{"result": 0},
}

var collectionLastOrDefaultPopulatedTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault on populated returns last",
	ExpectedInput: args.Map{"result": 30},
}

var collectionLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault on empty returns zero",
	ExpectedInput: args.Map{"result": 0},
}

// ==========================================================================
// Collection — SafeAt
// ==========================================================================

var collectionSafeAtValidTestCase = coretestcases.CaseV1{
	Title:         "SafeAt valid index returns item",
	ExpectedInput: args.Map{"result": 20},
}

var collectionSafeAtNegativeTestCase = coretestcases.CaseV1{
	Title:         "SafeAt negative index returns zero",
	ExpectedInput: args.Map{"result": 0},
}

var collectionSafeAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "SafeAt out-of-bounds returns zero",
	ExpectedInput: args.Map{"result": 0},
}

var collectionSafeAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "SafeAt on empty returns zero",
	ExpectedInput: args.Map{"result": 0},
}

// ==========================================================================
// Collection — ConcatNew
// ==========================================================================

var collectionConcatNewPopulatedTestCase = coretestcases.CaseV1{
	Title: "ConcatNew creates new collection with appended items",
	ExpectedInput: args.Map{
		"resultLength": 5,
		"resultFirst":  1,
		"resultLast":   5,
		"origLength":   3,
	},
}

var collectionConcatNewEmptyTestCase = coretestcases.CaseV1{
	Title: "ConcatNew on empty with items",
	ExpectedInput: args.Map{
		"length": 2,
		"first":  10,
		"last":   20,
	},
}

// ==========================================================================
// Collection — AddIf
// ==========================================================================

var collectionAddIfTrueTestCase = coretestcases.CaseV1{
	Title: "AddIf true adds item",
	ExpectedInput: args.Map{
		"length": 1,
		"first":  42,
	},
}

var collectionAddIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIf false does not add",
	ExpectedInput: args.Map{"length": 0},
}

// ==========================================================================
// Collection — ForEachBreak on empty
// ==========================================================================

var collectionForEachBreakEmptyTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak on empty does nothing",
	ExpectedInput: args.Map{"visited": 0},
}

// ==========================================================================
// Collection — AddSlice
// ==========================================================================

var collectionAddSlicePopulatedTestCase = coretestcases.CaseV1{
	Title: "AddSlice appends all items from slice",
	ExpectedInput: args.Map{
		"length": 3,
		"first":  10,
		"last":   30,
	},
}

var collectionAddSliceEmptyTestCase = coretestcases.CaseV1{
	Title:         "AddSlice with empty slice does nothing",
	ExpectedInput: args.Map{"length": 0},
}

// ==========================================================================
// Collection — Items / ItemsPtr
// ==========================================================================

var collectionItemsSliceTestCase = coretestcases.CaseV1{
	Title: "Items returns underlying slice",
	ExpectedInput: args.Map{
		"length": 3,
		"first":  1,
	},
}

var collectionItemsPtrTestCase = coretestcases.CaseV1{
	Title:         "ItemsPtr returns non-nil pointer",
	ExpectedInput: args.Map{"isNotNil": true},
}

// ==========================================================================
// Collection — RemoveAtLock
// ==========================================================================

var collectionRemoveAtLockTestCase = coretestcases.CaseV1{
	Title: "RemoveAtLock removes item thread-safely",
	ExpectedInput: args.Map{
		"removed": true,
		"length":  2,
	},
}
