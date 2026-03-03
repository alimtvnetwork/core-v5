package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Constructors
// ==========================================================================

var linkedListEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "EmptyLinkedList creates empty list",
		ExpectedInput: []string{"true", "0", "false"},
	},
}

var linkedListFromTestCases = []coretestcases.CaseV1{
	{
		Title:         "LinkedListFrom creates from slice",
		ExpectedInput: []string{"3", "a", "c"},
	},
	{
		Title:         "LinkedListFrom empty slice",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// Add
// ==========================================================================

var linkedListAddSingleTestCases = []coretestcases.CaseV1{
	{
		Title:         "Add single sets head and tail",
		ExpectedInput: []string{"1", "42", "42"},
	},
}

var linkedListAddMultipleTestCases = []coretestcases.CaseV1{
	{
		Title:         "Add multiple appends to back",
		ExpectedInput: []string{"1", "3", "3"},
	},
}

var linkedListAddFrontTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddFront prepends",
		ExpectedInput: []string{"1", "3", "3"},
	},
	{
		Title:         "AddFront empty",
		ExpectedInput: []string{"first", "first", "1"},
	},
}

var linkedListAddsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Adds multiple",
		ExpectedInput: []string{"3"},
	},
}

var linkedListAddSliceTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddSlice appends",
		ExpectedInput: []string{"2"},
	},
}

var linkedListAddIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddIf true adds",
		ExpectedInput: []string{"1"},
	},
	{
		Title:         "AddIf false skips",
		ExpectedInput: []string{"true"},
	},
}

var linkedListAddsIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddsIf false skips",
		ExpectedInput: []string{"true"},
	},
}

var linkedListAddFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddFunc adds result",
		ExpectedInput: []string{"99"},
	},
}

var linkedListPushTestCases = []coretestcases.CaseV1{
	{
		Title:         "Push aliases work",
		ExpectedInput: []string{"3"},
	},
}

// ==========================================================================
// FirstOrDefault / LastOrDefault
// ==========================================================================

var linkedListFirstDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "FirstOrDefault empty returns zero",
		ExpectedInput: []string{"0"},
	},
	{
		Title:         "FirstOrDefault non-empty",
		ExpectedInput: []string{"10"},
	},
}

var linkedListLastDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "LastOrDefault empty returns zero",
		ExpectedInput: []string{""},
	},
	{
		Title:         "LastOrDefault non-empty",
		ExpectedInput: []string{"20"},
	},
}

// ==========================================================================
// Items / Collection / String
// ==========================================================================

var linkedListItemsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Items returns all elements",
		ExpectedInput: []string{"3"},
	},
	{
		Title:         "Items empty returns empty",
		ExpectedInput: []string{"0"},
	},
}

var linkedListCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "Collection converts",
		ExpectedInput: []string{"2"},
	},
}

var linkedListStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "String representation",
		ExpectedInput: []string{"[1 2 3]"},
	},
}

// ==========================================================================
// IndexAt
// ==========================================================================

var linkedListIndexAtValidTestCase = coretestcases.CaseV1{
	Title:         "IndexAt valid returns node",
	ExpectedInput: []string{"true", "b"},
}

var linkedListIndexAtFirstTestCase = coretestcases.CaseV1{
	Title:         "IndexAt first",
	ExpectedInput: []string{"10"},
}

var linkedListIndexAtLastTestCase = coretestcases.CaseV1{
	Title:         "IndexAt last",
	ExpectedInput: []string{"30"},
}

var linkedListIndexAtOutOfBoundsTestCase = coretestcases.CaseV1{
	Title:         "IndexAt out of bounds",
	ExpectedInput: []string{"true", "true"},
}

var linkedListIndexAtEmptyTestCase = coretestcases.CaseV1{
	Title:         "IndexAt empty",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// ForEach
// ==========================================================================

var linkedListForEachTestCases = []coretestcases.CaseV1{
	{
		Title:         "ForEach visits all",
		ExpectedInput: []string{"6"},
	},
	{
		Title:         "ForEach empty noop",
		ExpectedInput: []string{"false"},
	},
}

var linkedListForEachBreakTestCases = []coretestcases.CaseV1{
	{
		Title:         "ForEachBreak stops early",
		ExpectedInput: []string{"3"},
	},
	{
		Title:         "ForEachBreak first element",
		ExpectedInput: []string{"1"},
	},
}

// ==========================================================================
// Head / Tail
// ==========================================================================

var linkedListHeadTailTestCases = []coretestcases.CaseV1{
	{
		Title:         "Head/Tail nodes",
		ExpectedInput: []string{"1", "3", "true", "false"},
	},
}

var linkedListNodeNextTestCases = []coretestcases.CaseV1{
	{
		Title:         "Node.Next traverses",
		ExpectedInput: []string{"10", "20", "30", "false"},
	},
}

// ==========================================================================
// Lock variants
// ==========================================================================

var linkedListLockTestCases = []coretestcases.CaseV1{
	{
		Title:         "LengthLock",
		ExpectedInput: []string{"2"},
	},
	{
		Title:         "IsEmptyLock",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "AddLock",
		ExpectedInput: []string{"2"},
	},
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var linkedListNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsEmpty nil receiver",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// AppendNode
// ==========================================================================

var linkedListAppendNodeTestCases = []coretestcases.CaseV1{
	{
		Title:         "AppendNode appends",
		ExpectedInput: []string{"3", "3"},
	},
	{
		Title:         "AppendNode empty",
		ExpectedInput: []string{"1", "99"},
	},
}
