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
		Title: "Distinct returns unique items preserving order -- [a,b,a,c,b,a]",
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
		Title: "Distinct returns same items -- already unique [x,y,z]",
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
		Title: "Distinct returns empty -- empty input",
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
		Title: "DistinctCount returns 3 -- [a,b,a,c,b]",
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
		Title: "IsDistinct returns true -- unique [a,b,c]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsDistinct returns false -- duplicates [a,b,a]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a"},
		},
		ExpectedInput: "false",
	},
}
