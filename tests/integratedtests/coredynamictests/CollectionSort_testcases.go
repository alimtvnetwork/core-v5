package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// SortAsc — strings
// ==========================================

var sortAscStringTestCases = []coretestcases.CaseV1{
	{
		Title: "SortAsc sorts strings alphabetically",
		ArrangeInput: args.Map{
			"when":  "given unsorted strings",
			"items": []string{"cherry", "apple", "banana"},
		},
		ExpectedInput: []string{"apple", "banana", "cherry"},
	},
}

// ==========================================
// SortDesc — strings
// ==========================================

var sortDescStringTestCases = []coretestcases.CaseV1{
	{
		Title: "SortDesc sorts strings reverse alphabetically",
		ArrangeInput: args.Map{
			"when":  "given unsorted strings",
			"items": []string{"cherry", "apple", "banana"},
		},
		ExpectedInput: []string{"cherry", "banana", "apple"},
	},
}

// ==========================================
// SortAsc — ints
// ==========================================

var sortAscIntTestCases = []coretestcases.CaseV1{
	{
		Title: "SortAsc sorts ints ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted ints",
			"items": []int{30, 10, 20, 5},
		},
		ExpectedInput: []string{"5", "10", "20", "30"},
	},
}

// ==========================================
// SortDesc — ints
// ==========================================

var sortDescIntTestCases = []coretestcases.CaseV1{
	{
		Title: "SortDesc sorts ints descending",
		ArrangeInput: args.Map{
			"when":  "given unsorted ints",
			"items": []int{30, 10, 20, 5},
		},
		ExpectedInput: []string{"30", "20", "10", "5"},
	},
}

// ==========================================
// SortedAsc — non-mutating
// ==========================================

var sortedAscNonMutatingTestCases = []coretestcases.CaseV1{
	{
		Title: "SortedAsc returns new sorted collection without mutating original",
		ArrangeInput: args.Map{
			"when":  "given unsorted strings",
			"items": []string{"c", "a", "b"},
		},
		ExpectedInput: []string{"a", "b", "c", "c", "a", "b"},
	},
}

// ==========================================
// Sort — empty collection
// ==========================================

var sortEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "SortAsc on empty collection is safe",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// Sort — single element
// ==========================================

var sortSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "SortAsc on single element returns same",
		ArrangeInput: args.Map{
			"when":  "given single element",
			"items": []string{"only"},
		},
		ExpectedInput: args.Map{
			"length": 1,
			"first":  "only",
		},
	},
}

// ==========================================
// IsSortedAsc
// ==========================================

var isSortedAscTrueTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSortedAsc returns true for sorted collection",
		ArrangeInput: args.Map{
			"when":  "given sorted ints",
			"items": []int{1, 2, 3, 4},
		},
		ExpectedInput: "true",
	},
}

var isSortedAscFalseTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSortedAsc returns false for unsorted collection",
		ArrangeInput: args.Map{
			"when":  "given unsorted ints",
			"items": []int{3, 1, 2},
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// SortFunc — custom comparator
// ==========================================

var sortFuncCustomTestCases = []coretestcases.CaseV1{
	{
		Title: "SortFunc with custom less sorts by string length",
		ArrangeInput: args.Map{
			"when":  "given strings sorted by length",
			"items": []string{"hello", "hi", "hey"},
		},
		ExpectedInput: []string{"hi", "hey", "hello"},
	},
}

// ==========================================
// SortAsc — float64
// ==========================================

var sortAscFloat64TestCases = []coretestcases.CaseV1{
	{
		Title: "SortAsc sorts float64 ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted floats",
			"items": []float64{3.14, 1.0, 2.71},
		},
		ExpectedInput: []string{"1", "2.71", "3.14"},
	},
}
