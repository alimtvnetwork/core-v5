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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3", // distinctCount
			Second: "a", // item0
			Third:  "b", // item1
			Fourth: "c", // item2
		},
	},
	{
		Title: "Distinct on already unique returns same items",
		ArrangeInput: args.Map{
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3", // distinctCount
			Second: "x", // item0
			Third:  "y", // item1
			Fourth: "z", // item2
		},
	},
	{
		Title: "Distinct on empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0", // distinctCount
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
		ExpectedInput: "3", // distinctCount
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
		ExpectedInput: "true", // isDistinct
	},
	{
		Title: "IsDistinct false for duplicates",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a"},
		},
		ExpectedInput: "false", // isDistinct
	},
}
