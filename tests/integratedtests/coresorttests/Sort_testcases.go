package coresorttests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// intsort.Quick — ascending
// =============================================================================

var intSortQuickTestCases = []coretestcases.CaseV1{
	{
		Title: "Quick sorts integers ascending",
		ArrangeInput: args.Map{
			"input": []int{3, 1, 4, 1, 5},
		},
		ExpectedInput: "[1 1 3 4 5]",
	},
	{
		Title: "Quick handles already sorted",
		ArrangeInput: args.Map{
			"input": []int{1, 2, 3},
		},
		ExpectedInput: "[1 2 3]",
	},
	{
		Title: "Quick handles empty slice",
		ArrangeInput: args.Map{
			"input": []int{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "Quick handles single element",
		ArrangeInput: args.Map{
			"input": []int{42},
		},
		ExpectedInput: "[42]",
	},
	{
		Title: "Quick handles negative numbers",
		ArrangeInput: args.Map{
			"input": []int{-3, 0, 2, -1},
		},
		ExpectedInput: "[-3 -1 0 2]",
	},
	{
		Title: "Quick handles all duplicates",
		ArrangeInput: args.Map{
			"input": []int{5, 5, 5},
		},
		ExpectedInput: "[5 5 5]",
	},
	{
		Title: "Quick handles reverse sorted",
		ArrangeInput: args.Map{
			"input": []int{5, 4, 3, 2, 1},
		},
		ExpectedInput: "[1 2 3 4 5]",
	},
}

// =============================================================================
// intsort.QuickDsc — descending
// =============================================================================

var intSortQuickDscTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickDsc sorts integers descending",
		ArrangeInput: args.Map{
			"input": []int{3, 1, 4, 1, 5},
		},
		ExpectedInput: "[5 4 3 1 1]",
	},
	{
		Title: "QuickDsc handles already descending",
		ArrangeInput: args.Map{
			"input": []int{5, 4, 3},
		},
		ExpectedInput: "[5 4 3]",
	},
	{
		Title: "QuickDsc handles empty slice",
		ArrangeInput: args.Map{
			"input": []int{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickDsc handles single element",
		ArrangeInput: args.Map{
			"input": []int{7},
		},
		ExpectedInput: "[7]",
	},
	{
		Title: "QuickDsc handles negative numbers",
		ArrangeInput: args.Map{
			"input": []int{-3, 0, 2, -1},
		},
		ExpectedInput: "[2 0 -1 -3]",
	},
	{
		Title: "QuickDsc handles all duplicates",
		ArrangeInput: args.Map{
			"input": []int{5, 5, 5},
		},
		ExpectedInput: "[5 5 5]",
	},
}

// =============================================================================
// intsort.QuickPtr — ascending pointer sort
// =============================================================================

var intSortQuickPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickPtr sorts integer pointers ascending",
		ArrangeInput: args.Map{
			"input": []int{3, 1, 4, 1, 5},
		},
		ExpectedInput: "[1 1 3 4 5]",
	},
	{
		Title: "QuickPtr handles empty slice",
		ArrangeInput: args.Map{
			"input": []int{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickPtr handles single element",
		ArrangeInput: args.Map{
			"input": []int{42},
		},
		ExpectedInput: "[42]",
	},
	{
		Title: "QuickPtr handles negative numbers",
		ArrangeInput: args.Map{
			"input": []int{-5, 3, -1, 0},
		},
		ExpectedInput: "[-5 -1 0 3]",
	},
	{
		Title: "QuickPtr handles all duplicates",
		ArrangeInput: args.Map{
			"input": []int{2, 2, 2},
		},
		ExpectedInput: "[2 2 2]",
	},
}

// =============================================================================
// intsort.QuickDscPtr — descending pointer sort
// =============================================================================

var intSortQuickDscPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickDscPtr sorts integer pointers descending",
		ArrangeInput: args.Map{
			"input": []int{3, 1, 4, 1, 5},
		},
		ExpectedInput: "[5 4 3 1 1]",
	},
	{
		Title: "QuickDscPtr handles empty slice",
		ArrangeInput: args.Map{
			"input": []int{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickDscPtr handles single element",
		ArrangeInput: args.Map{
			"input": []int{42},
		},
		ExpectedInput: "[42]",
	},
	{
		Title: "QuickDscPtr handles negative numbers",
		ArrangeInput: args.Map{
			"input": []int{-5, 3, -1, 0},
		},
		ExpectedInput: "[3 0 -1 -5]",
	},
}

// =============================================================================
// strsort.Quick — ascending
// =============================================================================

var strSortQuickTestCases = []coretestcases.CaseV1{
	{
		Title: "Quick sorts strings ascending",
		ArrangeInput: args.Map{
			"input": []string{"banana", "apple", "cherry"},
		},
		ExpectedInput: "[apple banana cherry]",
	},
	{
		Title: "Quick handles already sorted strings",
		ArrangeInput: args.Map{
			"input": []string{"a", "b", "c"},
		},
		ExpectedInput: "[a b c]",
	},
	{
		Title: "Quick handles empty string slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "Quick handles single string",
		ArrangeInput: args.Map{
			"input": []string{"only"},
		},
		ExpectedInput: "[only]",
	},
	{
		Title: "Quick handles duplicate strings",
		ArrangeInput: args.Map{
			"input": []string{"x", "a", "x", "a"},
		},
		ExpectedInput: "[a a x x]",
	},
	{
		Title: "Quick handles case-sensitive ordering",
		ArrangeInput: args.Map{
			"input": []string{"Banana", "apple", "Cherry"},
		},
		ExpectedInput: "[Banana Cherry apple]",
	},
}

// =============================================================================
// strsort.QuickDsc — descending
// =============================================================================

var strSortQuickDscTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickDsc sorts strings descending",
		ArrangeInput: args.Map{
			"input": []string{"banana", "apple", "cherry"},
		},
		ExpectedInput: "[cherry banana apple]",
	},
	{
		Title: "QuickDsc handles empty string slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickDsc handles single string",
		ArrangeInput: args.Map{
			"input": []string{"only"},
		},
		ExpectedInput: "[only]",
	},
	{
		Title: "QuickDsc handles duplicate strings",
		ArrangeInput: args.Map{
			"input": []string{"x", "a", "x"},
		},
		ExpectedInput: "[x x a]",
	},
	{
		Title: "QuickDsc handles case-sensitive ordering",
		ArrangeInput: args.Map{
			"input": []string{"Banana", "apple", "Cherry"},
		},
		ExpectedInput: "[apple Cherry Banana]",
	},
}

// =============================================================================
// strsort.QuickPtr — ascending pointer sort
// =============================================================================

var strSortQuickPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickPtr sorts string pointers ascending",
		ArrangeInput: args.Map{
			"input": []string{"banana", "apple", "cherry"},
		},
		ExpectedInput: "[apple banana cherry]",
	},
	{
		Title: "QuickPtr handles empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickPtr handles single element",
		ArrangeInput: args.Map{
			"input": []string{"only"},
		},
		ExpectedInput: "[only]",
	},
	{
		Title: "QuickPtr handles duplicates",
		ArrangeInput: args.Map{
			"input": []string{"z", "a", "z"},
		},
		ExpectedInput: "[a z z]",
	},
}

// =============================================================================
// strsort.QuickDscPtr — descending pointer sort
// =============================================================================

var strSortQuickDscPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "QuickDscPtr sorts string pointers descending",
		ArrangeInput: args.Map{
			"input": []string{"banana", "apple", "cherry"},
		},
		ExpectedInput: "[cherry banana apple]",
	},
	{
		Title: "QuickDscPtr handles empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: "[]",
	},
	{
		Title: "QuickDscPtr handles single element",
		ArrangeInput: args.Map{
			"input": []string{"only"},
		},
		ExpectedInput: "[only]",
	},
	{
		Title: "QuickDscPtr handles duplicates",
		ArrangeInput: args.Map{
			"input": []string{"z", "a", "z"},
		},
		ExpectedInput: "[z z a]",
	},
}
