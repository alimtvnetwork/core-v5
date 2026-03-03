package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// PointerSliceSorter — Ascending sort
// ==========================================================================

var ptrSorterAscIntTestCase = coretestcases.CaseV1{
	Title:         "Asc sort int pointers",
	ExpectedInput: []string{"1", "2", "3", "4", "5"},
}

var ptrSorterAscStringTestCase = coretestcases.CaseV1{
	Title:         "Asc sort strings",
	ExpectedInput: []string{"apple", "banana", "cherry"},
}

// ==========================================================================
// PointerSliceSorter — Descending sort
// ==========================================================================

var ptrSorterDescIntTestCase = coretestcases.CaseV1{
	Title:         "Desc sort int pointers",
	ExpectedInput: []string{"5", "4", "3", "2", "1"},
}

// ==========================================================================
// PointerSliceSorter — Nil handling
// ==========================================================================

var ptrSorterNilsToEndTestCase = coretestcases.CaseV1{
	Title:         "Asc sort with nils pushed to end",
	ExpectedInput: []string{"1", "3", "5", "<nil>", "<nil>"},
}

var ptrSorterNilFirstTestCase = coretestcases.CaseV1{
	Title:         "NilFirst=true pushes nils to beginning",
	ExpectedInput: []string{"<nil>", "<nil>", "1", "3", "5"},
}

var ptrSorterAllNilTestCase = coretestcases.CaseV1{
	Title:         "All nil slice stays stable",
	ExpectedInput: []string{"<nil>", "<nil>", "<nil>"},
}

// ==========================================================================
// PointerSliceSorter — Custom Less function
// ==========================================================================

var ptrSorterCustomLessTestCase = coretestcases.CaseV1{
	Title:         "Custom less: reverse absolute distance from 3",
	ExpectedInput: []string{"3", "2", "4", "1", "5"},
}

// ==========================================================================
// PointerSliceSorter — SetAsc / SetDesc switching
// ==========================================================================

var ptrSorterSwitchTestCase = coretestcases.CaseV1{
	Title:         "Sort asc then switch to desc and re-sort",
	ExpectedInput: []string{"1", "5", "5", "1"},
}

// ==========================================================================
// PointerSliceSorter — IsSorted
// ==========================================================================

var ptrSorterIsSortedTestCase = coretestcases.CaseV1{
	Title:         "IsSorted true after sort, false before",
	ExpectedInput: []string{"false", "true"},
}

// ==========================================================================
// PointerSliceSorter — Empty / single element
// ==========================================================================

var ptrSorterEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty slice: Len=0, IsSorted=true",
	ExpectedInput: []string{"0", "true"},
}

var ptrSorterSingleTestCase = coretestcases.CaseV1{
	Title:         "Single element: IsSorted=true after sort",
	ExpectedInput: []string{"1", "true", "42"},
}

var ptrSorterNilSliceTestCase = coretestcases.CaseV1{
	Title:         "Nil items slice: Len=0",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// PointerSliceSorter — SetItems / Items
// ==========================================================================

var ptrSorterSetItemsTestCase = coretestcases.CaseV1{
	Title:         "SetItems replaces slice and sorts new data",
	ExpectedInput: []string{"3", "10", "20", "30"},
}

// ==========================================================================
// PointerSliceSorter — Chaining
// ==========================================================================

var ptrSorterChainingTestCase = coretestcases.CaseV1{
	Title:         "Chained SetDesc.SetNilFirst.Sort produces correct order",
	ExpectedInput: []string{"<nil>", "5", "3", "1"},
}
