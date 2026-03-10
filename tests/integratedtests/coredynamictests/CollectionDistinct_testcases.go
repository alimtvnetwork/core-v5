package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		ExpectedInput: args.Map{
			"distinctCount": 3,
			"item0":         "a",
			"item1":         "b",
			"item2":         "c",
		},
	},
	{
		Title: "Distinct on already unique returns same items",
		ArrangeInput: args.Map{
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Map{
			"distinctCount": 3,
			"item0":         "x",
			"item1":         "y",
			"item2":         "z",
		},
	},
	{
		Title: "Distinct on empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0",
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
		ExpectedInput: "3",
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
		ExpectedInput: "true",
	},
	{
		Title: "IsDistinct false for duplicates",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a"},
		},
		ExpectedInput: "false",
	},
}
