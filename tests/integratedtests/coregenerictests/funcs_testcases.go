package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// MapCollection
// ==========================================================================

var mapCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapCollection int to string",
		ExpectedInput: []string{"3", "v1", "v3"},
	},
	{
		Title:         "MapCollection nil source",
		ExpectedInput: "true",
	},
	{
		Title:         "MapCollection empty source",
		ExpectedInput: "true",
	},
}

// ==========================================================================
// FlatMapCollection
// ==========================================================================

var flatMapCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "FlatMapCollection flattens",
		ExpectedInput: "6",
	},
	{
		Title:         "FlatMapCollection nil",
		ExpectedInput: "true",
	},
}

// ==========================================================================
// ReduceCollection
// ==========================================================================

var reduceCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "ReduceCollection sum",
		ExpectedInput: "10",
	},
	{
		Title:         "ReduceCollection nil returns initial",
		ExpectedInput: "99",
	},
	{
		Title:         "ReduceCollection string concat",
		ExpectedInput: "abc",
	},
}

// ==========================================================================
// GroupByCollection
// ==========================================================================

var groupByCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "GroupByCollection groups",
		ExpectedInput: []string{"2", "3", "3"},
	},
	{
		Title:         "GroupByCollection nil",
		ExpectedInput: "0",
	},
}

// ==========================================================================
// ContainsFunc
// ==========================================================================

var containsFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "ContainsFunc found",
		ExpectedInput: "true",
	},
	{
		Title:         "ContainsFunc not found",
		ExpectedInput: "false",
	},
	{
		Title:         "ContainsFunc nil",
		ExpectedInput: "false",
	},
}

// ==========================================================================
// ContainsItem
// ==========================================================================

var containsItemTestCases = []coretestcases.CaseV1{
	{
		Title:         "ContainsItem found",
		ExpectedInput: "true",
	},
	{
		Title:         "ContainsItem not found",
		ExpectedInput: "false",
	},
	{
		Title:         "ContainsItem nil",
		ExpectedInput: "false",
	},
}

// ==========================================================================
// IndexOfFunc
// ==========================================================================

var indexOfFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "IndexOfFunc found",
		ExpectedInput: "1",
	},
	{
		Title:         "IndexOfFunc not found",
		ExpectedInput: "-1",
	},
	{
		Title:         "IndexOfFunc nil",
		ExpectedInput: "-1",
	},
}

// ==========================================================================
// IndexOfItem
// ==========================================================================

var indexOfItemTestCases = []coretestcases.CaseV1{
	{
		Title:         "IndexOfItem found",
		ExpectedInput: "2",
	},
	{
		Title:         "IndexOfItem not found",
		ExpectedInput: "-1",
	},
}

// ==========================================================================
// Distinct
// ==========================================================================

var distinctTestCases = []coretestcases.CaseV1{
	{
		Title:         "Distinct removes duplicates",
		ExpectedInput: "3",
	},
	{
		Title:         "Distinct nil",
		ExpectedInput: "true",
	},
	{
		Title:         "Distinct no duplicates",
		ExpectedInput: "3",
	},
}

// ==========================================================================
// MapSimpleSlice
// ==========================================================================

var mapSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapSimpleSlice transforms",
		ExpectedInput: "3",
	},
	{
		Title:         "MapSimpleSlice nil",
		ExpectedInput: "true",
	},
}
