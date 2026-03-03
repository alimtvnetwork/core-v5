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
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "MapCollection empty source",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// FlatMapCollection
// ==========================================================================

var flatMapCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "FlatMapCollection flattens",
		ExpectedInput: []string{"6"},
	},
	{
		Title:         "FlatMapCollection nil",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// ReduceCollection
// ==========================================================================

var reduceCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:         "ReduceCollection sum",
		ExpectedInput: []string{"10"},
	},
	{
		Title:         "ReduceCollection nil returns initial",
		ExpectedInput: []string{"99"},
	},
	{
		Title:         "ReduceCollection string concat",
		ExpectedInput: []string{"abc"},
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
		ExpectedInput: []string{"0"},
	},
}

// ==========================================================================
// ContainsFunc
// ==========================================================================

var containsFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "ContainsFunc found",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "ContainsFunc not found",
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "ContainsFunc nil",
		ExpectedInput: []string{"false"},
	},
}

// ==========================================================================
// ContainsItem
// ==========================================================================

var containsItemTestCases = []coretestcases.CaseV1{
	{
		Title:         "ContainsItem found",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "ContainsItem not found",
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "ContainsItem nil",
		ExpectedInput: []string{"false"},
	},
}

// ==========================================================================
// IndexOfFunc
// ==========================================================================

var indexOfFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "IndexOfFunc found",
		ExpectedInput: []string{"1"},
	},
	{
		Title:         "IndexOfFunc not found",
		ExpectedInput: []string{"-1"},
	},
	{
		Title:         "IndexOfFunc nil",
		ExpectedInput: []string{"-1"},
	},
}

// ==========================================================================
// IndexOfItem
// ==========================================================================

var indexOfItemTestCases = []coretestcases.CaseV1{
	{
		Title:         "IndexOfItem found",
		ExpectedInput: []string{"2"},
	},
	{
		Title:         "IndexOfItem not found",
		ExpectedInput: []string{"-1"},
	},
}

// ==========================================================================
// Distinct
// ==========================================================================

var distinctTestCases = []coretestcases.CaseV1{
	{
		Title:         "Distinct removes duplicates",
		ExpectedInput: []string{"3"},
	},
	{
		Title:         "Distinct nil",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "Distinct no duplicates",
		ExpectedInput: []string{"3"},
	},
}

// ==========================================================================
// MapSimpleSlice
// ==========================================================================

var mapSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapSimpleSlice transforms",
		ExpectedInput: []string{"3"},
	},
	{
		Title:         "MapSimpleSlice nil",
		ExpectedInput: []string{"true"},
	},
}
