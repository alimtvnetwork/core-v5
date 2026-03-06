package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Length
// ==========================================================================

var idsLengthEmptyTestCase = coretestcases.CaseV1{
	Title:         "Length - empty returns 0",
	ExpectedInput: "0",
}

var idsLengthThreeItemsTestCase = coretestcases.CaseV1{
	Title:         "Length - 3 items returns 3",
	ExpectedInput: "3",
}

var idsLengthNilTestCase = coretestcases.CaseV1{
	Title:         "Length - nil receiver returns 0",
	ExpectedInput: "0",
}

// ==========================================================================
// GetById
// ==========================================================================

var idsGetByIdFoundTestCase = coretestcases.CaseV1{
	Title: "GetById - found returns item",
	ExpectedInput: args.Map{
		"found":    true,
		"id":       "beta",
		"isGlobal": true,
	},
}

var idsGetByIdMissingTestCase = coretestcases.CaseV1{
	Title:         "GetById - missing returns nil",
	ExpectedInput: "true",
}

var idsGetByIdEmptyTestCase = coretestcases.CaseV1{
	Title:         "GetById - empty id returns nil",
	ExpectedInput: "true",
}

// ==========================================================================
// Clone
// ==========================================================================

var idsCloneIndependenceTestCase = coretestcases.CaseV1{
	Title: "Clone - independence",
	ExpectedInput: args.Map{
		"originalLength": 2,
		"cloneLength":    3,
	},
}

var idsCloneEmptyTestCase = coretestcases.CaseV1{
	Title: "Clone - empty clones to empty",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"length":   0,
	},
}

var idsClonePreservesTestCase = coretestcases.CaseV1{
	Title: "Clone - preserves values",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"firstId":  "id-1",
		"isGlobal": false,
	},
}

// ==========================================================================
// Add
// ==========================================================================

var idsAddSingleTestCase = coretestcases.CaseV1{
	Title: "Add - single item",
	ExpectedInput: args.Map{
		"length":   1,
		"found":    true,
		"isGlobal": true,
	},
}

var idsAddEmptyIdTestCase = coretestcases.CaseV1{
	Title:         "Add - empty id ignored",
	ExpectedInput: "0",
}

var idsAddMultipleTestCase = coretestcases.CaseV1{
	Title: "Add - multiple accumulate",
	ExpectedInput: args.Map{
		"length":             3,
		"secondIsGlobal":     true,
		"thirdIsGlobal":      false,
	},
}

// ==========================================================================
// IsEmpty / HasAnyItem
// ==========================================================================

var idsIsEmptyTrueTestCase = coretestcases.CaseV1{
	Title: "IsEmpty - empty true",
	ExpectedInput: args.Map{
		"isEmpty":    true,
		"hasAnyItem": false,
	},
}

var idsIsEmptyFalseTestCase = coretestcases.CaseV1{
	Title: "IsEmpty - non-empty false",
	ExpectedInput: args.Map{
		"isEmpty":    false,
		"hasAnyItem": true,
	},
}

// ==========================================================================
// IndexOf
// ==========================================================================

var idsIndexOfFoundTestCase = coretestcases.CaseV1{
	Title: "IndexOf - found returns correct index",
	ExpectedInput: args.Map{
		"indexOfA": 0,
		"indexOfB": 1,
		"indexOfC": 2,
	},
}

var idsIndexOfMissingTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - missing returns -1",
	ExpectedInput: "-1",
}

var idsIndexOfEmptyStringTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - empty string returns -1",
	ExpectedInput: "-1",
}

var idsIndexOfEmptyCollectionTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - empty collection returns -1",
	ExpectedInput: "-1",
}

// ==========================================================================
// Adds
// ==========================================================================

var idsAddsBatchTestCase = coretestcases.CaseV1{
	Title: "Adds - batch add all items",
	ExpectedInput: args.Map{
		"length":   3,
		"foundOne": true,
		"foundTwo": true,
		"foundThree": true,
	},
}

var idsAddsEmptyTestCase = coretestcases.CaseV1{
	Title:         "Adds - empty ids no add",
	ExpectedInput: "0",
}

// ==========================================================================
// New edge
// ==========================================================================

var idsNewEdgeEmptyTestCase = coretestcases.CaseV1{
	Title: "New - no ids creates empty",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"length":   0,
	},
}
