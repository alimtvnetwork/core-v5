package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Contains
// ==========================================

var collectionContainsTestCases = []coretestcases.CaseV1{
	{
		Title: "Contains returns true for existing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c"},
			"search": "b",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "Contains returns false for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c"},
			"search": "z",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Contains returns false for empty collection",
		ArrangeInput: args.Map{
			"items":  []string{},
			"search": "a",
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// IndexOf
// ==========================================

var collectionIndexOfTestCases = []coretestcases.CaseV1{
	{
		Title: "IndexOf returns correct index",
		ArrangeInput: args.Map{
			"items":  []string{"x", "y", "z"},
			"search": "y",
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "IndexOf returns -1 for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"x", "y", "z"},
			"search": "w",
		},
		ExpectedInput: []string{"-1"},
	},
}

// ==========================================
// HasAll
// ==========================================

var collectionHasAllTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAll returns true when all present",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c", "d"},
			"search": []string{"b", "d"},
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "HasAll returns false when one missing",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b"},
			"search": []string{"a", "z"},
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================
// LastIndexOf
// ==========================================

var collectionLastIndexOfTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndexOf returns last occurrence",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "a", "c"},
			"search": "a",
		},
		ExpectedInput: []string{"2"},
	},
}

// ==========================================
// Count
// ==========================================

var collectionCountTestCases = []coretestcases.CaseV1{
	{
		Title: "Count returns correct occurrence count",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "a", "a", "c"},
			"search": "a",
		},
		ExpectedInput: []string{"3"},
	},
	{
		Title: "Count returns 0 for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b"},
			"search": "z",
		},
		ExpectedInput: []string{"0"},
	},
}
