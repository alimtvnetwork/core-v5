package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Distinct
// ==========================================

var collectionDistinctTestCases = []coretestcases.CaseV1{
	{
		Title: "Distinct removes duplicates preserving order",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a", "c", "b", "a"},
		},
		ExpectedInput: []string{"3", "a", "b", "c"},
	},
	{
		Title: "Distinct on already unique returns same items",
		ArrangeInput: args.Map{
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: []string{"3", "x", "y", "z"},
	},
	{
		Title: "Distinct on empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: []string{"0"},
	},
}

// ==========================================
// DistinctCount
// ==========================================

var collectionDistinctCountTestCases = []coretestcases.CaseV1{
	{
		Title: "DistinctCount returns unique count",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a", "c", "b"},
		},
		ExpectedInput: []string{"3"},
	},
}

// ==========================================
// IsDistinct
// ==========================================

var collectionIsDistinctTestCases = []coretestcases.CaseV1{
	{
		Title: "IsDistinct true for unique items",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsDistinct false for duplicates",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a"},
		},
		ExpectedInput: []string{"false"},
	},
}
