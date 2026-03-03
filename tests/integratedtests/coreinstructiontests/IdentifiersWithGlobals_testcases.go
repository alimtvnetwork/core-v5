package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Length
// ==========================================================================

var idsLengthEmptyTestCase = coretestcases.CaseV1{
	Title:         "Length - empty returns 0",
	ExpectedInput: []string{"0"},
}

var idsLengthThreeItemsTestCase = coretestcases.CaseV1{
	Title:         "Length - 3 items returns 3",
	ExpectedInput: []string{"3"},
}

var idsLengthNilTestCase = coretestcases.CaseV1{
	Title:         "Length - nil receiver returns 0",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// GetById
// ==========================================================================

var idsGetByIdFoundTestCase = coretestcases.CaseV1{
	Title:         "GetById - found returns item",
	ExpectedInput: []string{"true", "beta", "true"},
}

var idsGetByIdMissingTestCase = coretestcases.CaseV1{
	Title:         "GetById - missing returns nil",
	ExpectedInput: []string{"true"},
}

var idsGetByIdEmptyTestCase = coretestcases.CaseV1{
	Title:         "GetById - empty id returns nil",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Clone
// ==========================================================================

var idsCloneIndependenceTestCase = coretestcases.CaseV1{
	Title:         "Clone - independence",
	ExpectedInput: []string{"2", "3"},
}

var idsCloneEmptyTestCase = coretestcases.CaseV1{
	Title:         "Clone - empty clones to empty",
	ExpectedInput: []string{"true", "0"},
}

var idsClonePreservesTestCase = coretestcases.CaseV1{
	Title:         "Clone - preserves values",
	ExpectedInput: []string{"true", "id-1", "false"},
}

// ==========================================================================
// Add
// ==========================================================================

var idsAddSingleTestCase = coretestcases.CaseV1{
	Title:         "Add - single item",
	ExpectedInput: []string{"1", "true", "true"},
}

var idsAddEmptyIdTestCase = coretestcases.CaseV1{
	Title:         "Add - empty id ignored",
	ExpectedInput: []string{"0"},
}

var idsAddMultipleTestCase = coretestcases.CaseV1{
	Title:         "Add - multiple accumulate",
	ExpectedInput: []string{"3", "true", "false"},
}

// ==========================================================================
// IsEmpty / HasAnyItem
// ==========================================================================

var idsIsEmptyTrueTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty - empty true",
	ExpectedInput: []string{"true", "false"},
}

var idsIsEmptyFalseTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty - non-empty false",
	ExpectedInput: []string{"false", "true"},
}

// ==========================================================================
// IndexOf
// ==========================================================================

var idsIndexOfFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - found returns correct index",
	ExpectedInput: []string{"0", "1", "2"},
}

var idsIndexOfMissingTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - missing returns -1",
	ExpectedInput: []string{"-1"},
}

var idsIndexOfEmptyStringTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - empty string returns -1",
	ExpectedInput: []string{"-1"},
}

var idsIndexOfEmptyCollectionTestCase = coretestcases.CaseV1{
	Title:         "IndexOf - empty collection returns -1",
	ExpectedInput: []string{"-1"},
}

// ==========================================================================
// Adds
// ==========================================================================

var idsAddsBatchTestCase = coretestcases.CaseV1{
	Title:         "Adds - batch add all items",
	ExpectedInput: []string{"3", "true", "true", "true"},
}

var idsAddsEmptyTestCase = coretestcases.CaseV1{
	Title:         "Adds - empty ids no add",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// New edge
// ==========================================================================

var idsNewEdgeEmptyTestCase = coretestcases.CaseV1{
	Title:         "New - no ids creates empty",
	ExpectedInput: []string{"true", "0"},
}
