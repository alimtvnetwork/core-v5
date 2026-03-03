package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// MapCollection
// ==========================================================================

var mapCollectionIntToStringTestCase = coretestcases.CaseV1{
	Title:         "MapCollection int to string",
	ExpectedInput: []string{"3", "v1", "v3"},
}

var mapCollectionNilSourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection nil source",
	ExpectedInput: []string{"true"},
}

var mapCollectionEmptySourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection empty source",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// FlatMapCollection
// ==========================================================================

var flatMapCollectionFlattensTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection flattens",
	ExpectedInput: []string{"6"},
}

var flatMapCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection nil",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// ReduceCollection
// ==========================================================================

var reduceCollectionSumTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection sum",
	ExpectedInput: []string{"10"},
}

var reduceCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection nil returns initial",
	ExpectedInput: []string{"99"},
}

var reduceCollectionConcatTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection string concat",
	ExpectedInput: []string{"abc"},
}

// ==========================================================================
// GroupByCollection
// ==========================================================================

var groupByCollectionGroupsTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection groups",
	ExpectedInput: []string{"2", "3", "3"},
}

var groupByCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection nil",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// ContainsFunc
// ==========================================================================

var containsFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc found",
	ExpectedInput: []string{"true"},
}

var containsFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc not found",
	ExpectedInput: []string{"false"},
}

var containsFuncNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc nil",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// ContainsItem
// ==========================================================================

var containsItemFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem found",
	ExpectedInput: []string{"true"},
}

var containsItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem not found",
	ExpectedInput: []string{"false"},
}

var containsItemNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem nil",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// IndexOfFunc
// ==========================================================================

var indexOfFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc found",
	ExpectedInput: []string{"1"},
}

var indexOfFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc not found",
	ExpectedInput: []string{"-1"},
}

var indexOfFuncNilTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc nil",
	ExpectedInput: []string{"-1"},
}

// ==========================================================================
// IndexOfItem
// ==========================================================================

var indexOfItemFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem found",
	ExpectedInput: []string{"2"},
}

var indexOfItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem not found",
	ExpectedInput: []string{"-1"},
}

// ==========================================================================
// Distinct
// ==========================================================================

var distinctRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct removes duplicates",
	ExpectedInput: []string{"3"},
}

var distinctNilTestCase = coretestcases.CaseV1{
	Title:         "Distinct nil",
	ExpectedInput: []string{"true"},
}

var distinctNoDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct no duplicates",
	ExpectedInput: []string{"3"},
}

// ==========================================================================
// MapSimpleSlice
// ==========================================================================

var mapSimpleSliceTransformsTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice transforms",
	ExpectedInput: []string{"3"},
}

var mapSimpleSliceNilTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice nil",
	ExpectedInput: []string{"true"},
}
