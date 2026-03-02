package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// PointerSliceSorter — Ascending sort
// ==========================================================================

var ptrSorterAscTestCases = []coretestcases.CaseV1{
	{
		Name:      "Asc sort int pointers",
		WantLines: []string{"1", "2", "3", "4", "5"},
	},
	{
		Name:      "Asc sort strings",
		WantLines: []string{"apple", "banana", "cherry"},
	},
}

// ==========================================================================
// PointerSliceSorter — Descending sort
// ==========================================================================

var ptrSorterDescTestCases = []coretestcases.CaseV1{
	{
		Name:      "Desc sort int pointers",
		WantLines: []string{"5", "4", "3", "2", "1"},
	},
}

// ==========================================================================
// PointerSliceSorter — Nil handling
// ==========================================================================

var ptrSorterNilHandlingTestCases = []coretestcases.CaseV1{
	{
		Name:      "Asc sort with nils pushed to end",
		WantLines: []string{"1", "3", "5", "<nil>", "<nil>"},
	},
	{
		Name:      "NilFirst=true pushes nils to beginning",
		WantLines: []string{"<nil>", "<nil>", "1", "3", "5"},
	},
	{
		Name:      "All nil slice stays stable",
		WantLines: []string{"<nil>", "<nil>", "<nil>"},
	},
}

// ==========================================================================
// PointerSliceSorter — Custom Less function
// ==========================================================================

var ptrSorterCustomLessTestCases = []coretestcases.CaseV1{
	{
		Name:      "Custom less: reverse absolute distance from 3",
		WantLines: []string{"3", "2", "4", "1", "5"},
	},
}

// ==========================================================================
// PointerSliceSorter — SetAsc / SetDesc switching
// ==========================================================================

var ptrSorterSwitchTestCases = []coretestcases.CaseV1{
	{
		Name:      "Sort asc then switch to desc and re-sort",
		WantLines: []string{"1", "5", "5", "1"},
	},
}

// ==========================================================================
// PointerSliceSorter — IsSorted
// ==========================================================================

var ptrSorterIsSortedTestCases = []coretestcases.CaseV1{
	{
		Name:      "IsSorted true after sort, false before",
		WantLines: []string{"false", "true"},
	},
}

// ==========================================================================
// PointerSliceSorter — Empty / single element
// ==========================================================================

var ptrSorterEdgeTestCases = []coretestcases.CaseV1{
	{
		Name:      "Empty slice: Len=0, IsSorted=true",
		WantLines: []string{"0", "true"},
	},
	{
		Name:      "Single element: IsSorted=true after sort",
		WantLines: []string{"1", "true", "42"},
	},
	{
		Name:      "Nil items slice: Len=0",
		WantLines: []string{"0"},
	},
}

// ==========================================================================
// PointerSliceSorter — SetItems / Items
// ==========================================================================

var ptrSorterSetItemsTestCases = []coretestcases.CaseV1{
	{
		Name:      "SetItems replaces slice and sorts new data",
		WantLines: []string{"3", "10", "20", "30"},
	},
}

// ==========================================================================
// PointerSliceSorter — Chaining
// ==========================================================================

var ptrSorterChainingTestCases = []coretestcases.CaseV1{
	{
		Name:      "Chained SetDesc.SetNilFirst.Sort produces correct order",
		WantLines: []string{"<nil>", "5", "3", "1"},
	},
}
