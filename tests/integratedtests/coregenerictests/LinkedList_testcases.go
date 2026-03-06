package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Constructors
// ==========================================================================

var linkedListEmptyTestCase = coretestcases.CaseV1{
	Title: "EmptyLinkedList creates empty list",
	ExpectedInput: args.Map{
		"isEmpty":  true,
		"length":   0,
		"hasItems": false,
	},
}

var linkedListFromSliceTestCase = coretestcases.CaseV1{
	Title: "LinkedListFrom creates from slice",
	ExpectedInput: args.Map{
		"length": 3,
		"first":  "a",
		"last":   "c",
	},
}

var linkedListFromEmptySliceTestCase = coretestcases.CaseV1{
	Title:         "LinkedListFrom empty slice",
	ExpectedInput: "true",
}

// ==========================================================================
// Add
// ==========================================================================

var linkedListAddSingleTestCase = coretestcases.CaseV1{
	Title: "Add single sets head and tail",
	ExpectedInput: args.Map{
		"length": 1,
		"head":   42,
		"tail":   42,
	},
}

var linkedListAddMultipleTestCase = coretestcases.CaseV1{
	Title: "Add multiple appends to back",
	ExpectedInput: args.Map{
		"head":   1,
		"tail":   3,
		"length": 3,
	},
}

var linkedListAddFrontPrependsTestCase = coretestcases.CaseV1{
	Title: "AddFront prepends",
	ExpectedInput: args.Map{
		"head":   1,
		"tail":   3,
		"length": 3,
	},
}

var linkedListAddFrontEmptyTestCase = coretestcases.CaseV1{
	Title: "AddFront empty",
	ExpectedInput: args.Map{
		"head":   "first",
		"tail":   "first",
		"length": 1,
	},
}

var linkedListAddsTestCase = coretestcases.CaseV1{
	Title:         "Adds multiple",
	ExpectedInput: "3",
}

var linkedListAddSliceTestCase = coretestcases.CaseV1{
	Title:         "AddSlice appends",
	ExpectedInput: "2",
}

var linkedListAddIfTrueTestCase = coretestcases.CaseV1{
	Title:         "AddIf true adds",
	ExpectedInput: "1",
}

var linkedListAddIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddIf false skips",
	ExpectedInput: "true",
}

var linkedListAddsIfFalseTestCase = coretestcases.CaseV1{
	Title:         "AddsIf false skips",
	ExpectedInput: "true",
}

var linkedListAddFuncTestCase = coretestcases.CaseV1{
	Title:         "AddFunc adds result",
	ExpectedInput: "99",
}

var linkedListPushTestCase = coretestcases.CaseV1{
	Title:         "Push aliases work",
	ExpectedInput: "3",
}

// ==========================================================================
// FirstOrDefault / LastOrDefault
// ==========================================================================

var linkedListFirstDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault empty returns zero",
	ExpectedInput: "0",
}

var linkedListFirstDefaultNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "FirstOrDefault non-empty",
	ExpectedInput: "10",
}

var linkedListLastDefaultEmptyTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault empty returns zero",
	ExpectedInput: "",
}

var linkedListLastDefaultNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "LastOrDefault non-empty",
	ExpectedInput: "20",
}

// ==========================================================================
// Items / Collection / String
// ==========================================================================

var linkedListItemsAllTestCase = coretestcases.CaseV1{
	Title:         "Items returns all elements",
	ExpectedInput: "3",
}

var linkedListItemsEmptyTestCase = coretestcases.CaseV1{
	Title:         "Items empty returns empty",
	ExpectedInput: "0",
}

var linkedListCollectionTestCase = coretestcases.CaseV1{
	Title:         "Collection converts",
	ExpectedInput: "2",
}

var linkedListStringTestCase = coretestcases.CaseV1{
	Title:         "String representation",
	ExpectedInput: "[1 2 3]",
}

// ==========================================================================
// IndexAt
// ==========================================================================

var linkedListIndexAtValidTestCase = coretestcases.CaseV1{
	Title: "IndexAt valid returns node",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"value":    "b",
	},
}

var linkedListIndexAtFirstTestCase = coretestcases.CaseV1{
	Title:         "IndexAt first",
	ExpectedInput: "10",
}

var linkedListIndexAtLastTestCase = coretestcases.CaseV1{
	Title:         "IndexAt last",
	ExpectedInput: "30",
}

var linkedListIndexAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title: "IndexAt out of bounds",
	ExpectedInput: args.Map{
		"isNil":    true,
		"hasError": true,
	},
}

var linkedListIndexAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "IndexAt empty",
	ExpectedInput: "true",
}

// ==========================================================================
// ForEach
// ==========================================================================

var linkedListForEachVisitsAllTestCase = coretestcases.CaseV1{
	Title:         "ForEach visits all",
	ExpectedInput: "6",
}

var linkedListForEachEmptyTestCase = coretestcases.CaseV1{
	Title:         "ForEach empty noop",
	ExpectedInput: "false",
}

var linkedListForEachBreakStopsEarlyTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops early",
	ExpectedInput: "3",
}

var linkedListForEachBreakFirstTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak first element",
	ExpectedInput: "1",
}

// ==========================================================================
// Head / Tail
// ==========================================================================

var linkedListHeadTailTestCase = coretestcases.CaseV1{
	Title: "Head/Tail nodes",
	ExpectedInput: args.Map{
		"head":        1,
		"tail":        3,
		"headHasNext": true,
		"tailHasNext": false,
	},
}

var linkedListNodeNextTestCase = coretestcases.CaseV1{
	Title: "Node.Next traverses",
	ExpectedInput: args.Map{
		"first":   10,
		"second":  20,
		"third":   30,
		"hasMore": false,
	},
}

// ==========================================================================
// Lock variants
// ==========================================================================

var linkedListLengthLockTestCase = coretestcases.CaseV1{
	Title:         "LengthLock",
	ExpectedInput: "2",
}

var linkedListIsEmptyLockTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock",
	ExpectedInput: "true",
}

var linkedListAddLockTestCase = coretestcases.CaseV1{
	Title:         "AddLock",
	ExpectedInput: "2",
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var linkedListNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty nil receiver",
	ExpectedInput: "true",
}

// ==========================================================================
// AppendNode
// ==========================================================================

var linkedListAppendNodeAppendsTestCase = coretestcases.CaseV1{
	Title: "AppendNode appends",
	ExpectedInput: args.Map{
		"length":    3,
		"lastValue": 3,
	},
}

var linkedListAppendNodeEmptyTestCase = coretestcases.CaseV1{
	Title: "AppendNode empty",
	ExpectedInput: args.Map{
		"length": 1,
		"value":  99,
	},
}
