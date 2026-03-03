package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// PointerSliceSorter — Ascending sort
// ==========================================================================

var ptrSorterAscIntTestCase = coretestcases.CaseV1{
	Name:      "Asc sort int pointers",
	WantLines: []string{"1", "2", "3", "4", "5"},
}

var ptrSorterAscStringTestCase = coretestcases.CaseV1{
	Name:      "Asc sort strings",
	WantLines: []string{"apple", "banana", "cherry"},
}

// ==========================================================================
// PointerSliceSorter — Descending sort
// ==========================================================================

var ptrSorterDescIntTestCase = coretestcases.CaseV1{
	Name:      "Desc sort int pointers",
	WantLines: []string{"5", "4", "3", "2", "1"},
}

// ==========================================================================
// PointerSliceSorter — Nil handling
// ==========================================================================

var ptrSorterNilsToEndTestCase = coretestcases.CaseV1{
	Name:      "Asc sort with nils pushed to end",
	WantLines: []string{"1", "3", "5", "<nil>", "<nil>"},
}

var ptrSorterNilFirstTestCase = coretestcases.CaseV1{
	Name:      "NilFirst=true pushes nils to beginning",
	WantLines: []string{"<nil>", "<nil>", "1", "3", "5"},
}

var ptrSorterAllNilTestCase = coretestcases.CaseV1{
	Name:      "All nil slice stays stable",
	WantLines: []string{"<nil>", "<nil>", "<nil>"},
}

// ==========================================================================
// PointerSliceSorter — Custom Less function
// ==========================================================================

var ptrSorterCustomLessTestCase = coretestcases.CaseV1{
	Name:      "Custom less: reverse absolute distance from 3",
	WantLines: []string{"3", "2", "4", "1", "5"},
}

// ==========================================================================
// PointerSliceSorter — SetAsc / SetDesc switching
// ==========================================================================

var ptrSorterSwitchTestCase = coretestcases.CaseV1{
	Name:      "Sort asc then switch to desc and re-sort",
	WantLines: []string{"1", "5", "5", "1"},
}

// ==========================================================================
// PointerSliceSorter — IsSorted
// ==========================================================================

var ptrSorterIsSortedTestCase = coretestcases.CaseV1{
	Name:      "IsSorted true after sort, false before",
	WantLines: []string{"false", "true"},
}

// ==========================================================================
// PointerSliceSorter — Empty / single element
// ==========================================================================

var ptrSorterEmptyTestCase = coretestcases.CaseV1{
	Name:      "Empty slice: Len=0, IsSorted=true",
	WantLines: []string{"0", "true"},
}

var ptrSorterSingleTestCase = coretestcases.CaseV1{
	Name:      "Single element: IsSorted=true after sort",
	WantLines: []string{"1", "true", "42"},
}

var ptrSorterNilSliceTestCase = coretestcases.CaseV1{
	Name:      "Nil items slice: Len=0",
	WantLines: []string{"0"},
}

// ==========================================================================
// PointerSliceSorter — SetItems / Items
// ==========================================================================

var ptrSorterSetItemsTestCase = coretestcases.CaseV1{
	Name:      "SetItems replaces slice and sorts new data",
	WantLines: []string{"3", "10", "20", "30"},
}

// ==========================================================================
// PointerSliceSorter — Chaining
// ==========================================================================

var ptrSorterChainingTestCase = coretestcases.CaseV1{
	Name:      "Chained SetDesc.SetNilFirst.Sort produces correct order",
	WantLines: []string{"<nil>", "5", "3", "1"},
}
