package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Length
// ==========================================================================

var idsLengthTestCases = []coretestcases.CaseV1{
	{
		Title:         "Length - empty returns 0",
		ExpectedInput: []string{"0"},
	},
	{
		Title:         "Length - 3 items returns 3",
		ExpectedInput: []string{"3"},
	},
	{
		Title:         "Length - nil receiver returns 0",
		ExpectedInput: []string{"0"},
	},
}

// ==========================================================================
// GetById
// ==========================================================================

var idsGetByIdTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetById - found returns item",
		ExpectedInput: []string{"true", "beta", "true"},
	},
	{
		Title:         "GetById - missing returns nil",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "GetById - empty id returns nil",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// Clone
// ==========================================================================

var idsCloneTestCases = []coretestcases.CaseV1{
	{
		Title:         "Clone - independence",
		ExpectedInput: []string{"2", "3"},
	},
	{
		Title:         "Clone - empty clones to empty",
		ExpectedInput: []string{"true", "0"},
	},
	{
		Title:         "Clone - preserves values",
		ExpectedInput: []string{"true", "id-1", "false"},
	},
}

// ==========================================================================
// Add
// ==========================================================================

var idsAddTestCases = []coretestcases.CaseV1{
	{
		Title:         "Add - single item",
		ExpectedInput: []string{"1", "true", "true"},
	},
	{
		Title:         "Add - empty id ignored",
		ExpectedInput: []string{"0"},
	},
	{
		Title:         "Add - multiple accumulate",
		ExpectedInput: []string{"3", "true", "false"},
	},
}

// ==========================================================================
// IsEmpty / HasAnyItem
// ==========================================================================

var idsIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsEmpty - empty true",
		ExpectedInput: []string{"true", "false"},
	},
	{
		Title:         "IsEmpty - non-empty false",
		ExpectedInput: []string{"false", "true"},
	},
}

// ==========================================================================
// IndexOf
// ==========================================================================

var idsIndexOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "IndexOf - found returns correct index",
		ExpectedInput: []string{"0", "1", "2"},
	},
	{
		Title:         "IndexOf - missing returns -1",
		ExpectedInput: []string{"-1"},
	},
	{
		Title:         "IndexOf - empty string returns -1",
		ExpectedInput: []string{"-1"},
	},
	{
		Title:         "IndexOf - empty collection returns -1",
		ExpectedInput: []string{"-1"},
	},
}

// ==========================================================================
// Adds
// ==========================================================================

var idsAddsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Adds - batch add all items",
		ExpectedInput: []string{"3", "true", "true", "true"},
	},
	{
		Title:         "Adds - empty ids no add",
		ExpectedInput: []string{"0"},
	},
}

// ==========================================================================
// New edge
// ==========================================================================

var idsNewEdgeTestCases = []coretestcases.CaseV1{
	{
		Title:         "New - no ids creates empty",
		ExpectedInput: []string{"true", "0"},
	},
}
