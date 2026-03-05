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
	ExpectedInput: args.Three[string, string, string]{
		First:  "true", // found
		Second: "beta", // id
		Third:  "true", // isGlobal
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
	ExpectedInput: args.Two[string, string]{
		First:  "2", // originalLength
		Second: "3", // cloneLength
	},
}

var idsCloneEmptyTestCase = coretestcases.CaseV1{
	Title: "Clone - empty clones to empty",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isEmpty
		Second: "0",    // length
	},
}

var idsClonePreservesTestCase = coretestcases.CaseV1{
	Title: "Clone - preserves values",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true", // isNotNil
		Second: "id-1", // firstId
		Third:  "false", // isEmpty
	},
}

// ==========================================================================
// Add
// ==========================================================================

var idsAddSingleTestCase = coretestcases.CaseV1{
	Title: "Add - single item",
	ExpectedInput: args.Three[string, string, string]{
		First:  "1",    // length
		Second: "true", // found
		Third:  "true", // isGlobal
	},
}

var idsAddEmptyIdTestCase = coretestcases.CaseV1{
	Title:         "Add - empty id ignored",
	ExpectedInput: "0",
}

var idsAddMultipleTestCase = coretestcases.CaseV1{
	Title: "Add - multiple accumulate",
	ExpectedInput: args.Three[string, string, string]{
		First:  "3",     // length
		Second: "true",  // foundFirst
		Third:  "false", // foundMissing
	},
}

// ==========================================================================
// IsEmpty / HasAnyItem
// ==========================================================================

var idsIsEmptyTrueTestCase = coretestcases.CaseV1{
	Title: "IsEmpty - empty true",
	ExpectedInput: args.Two[string, string]{
		First:  "true",  // isEmpty
		Second: "false", // hasAnyItem
	},
}

var idsIsEmptyFalseTestCase = coretestcases.CaseV1{
	Title: "IsEmpty - non-empty false",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isEmpty
		Second: "true",  // hasAnyItem
	},
}

// ==========================================================================
// IndexOf
// ==========================================================================

var idsIndexOfFoundTestCase = coretestcases.CaseV1{
	Title: "IndexOf - found returns correct index",
	ExpectedInput: args.Three[string, string, string]{
		First:  "0", // indexOf first
		Second: "1", // indexOf second
		Third:  "2", // indexOf third
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
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "3",    // length
		Second: "true", // found1
		Third:  "true", // found2
		Fourth: "true", // found3
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
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isEmpty
		Second: "0",    // length
	},
}
