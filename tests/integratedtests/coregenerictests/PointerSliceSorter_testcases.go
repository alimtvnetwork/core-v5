package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
	Title: "Sort asc then switch to desc and re-sort",
	ExpectedInput: args.Map{
		"firstAfterAsc":  "1",
		"lastAfterAsc":   "5",
		"firstAfterDesc": "5",
		"lastAfterDesc":  "1",
	},
}

// ==========================================================================
// PointerSliceSorter — IsSorted
// ==========================================================================

var ptrSorterIsSortedTestCase = coretestcases.CaseV1{
	Title: "IsSorted true after sort, false before",
	ExpectedInput: args.Map{
		"beforeSort": false,
		"afterSort":  true,
	},
}

// ==========================================================================
// PointerSliceSorter — Empty / single element
// ==========================================================================

var ptrSorterEmptyTestCase = coretestcases.CaseV1{
	Title: "Empty slice: Len=0, IsSorted=true",
	ExpectedInput: args.Map{
		"length":   0,
		"isSorted": true,
	},
}

var ptrSorterSingleTestCase = coretestcases.CaseV1{
	Title: "Single element: IsSorted=true after sort",
	ExpectedInput: args.Map{
		"length":   1,
		"isSorted": true,
		"value":    "42",
	},
}

var ptrSorterNilSliceTestCase = coretestcases.CaseV1{
	Title:         "Nil items slice: Len=0",
	ExpectedInput: "0",
}

// ==========================================================================
// PointerSliceSorter — SetItems / Items
// ==========================================================================

var ptrSorterSetItemsTestCase = coretestcases.CaseV1{
	Title: "SetItems replaces slice and sorts new data",
	ExpectedInput: args.Map{
		"length": 3,
		"item0":  "10",
		"item1":  "20",
		"item2":  "30",
	},
}

// ==========================================================================
// PointerSliceSorter — Chaining
// ==========================================================================

var ptrSorterChainingTestCase = coretestcases.CaseV1{
	Title:         "Chained SetDesc.SetNilFirst.Sort produces correct order",
	ExpectedInput: []string{"<nil>", "5", "3", "1"},
}
