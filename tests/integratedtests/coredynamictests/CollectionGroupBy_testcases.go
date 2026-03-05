package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// GroupBy — group by first character
// Each expected line is "key:count" sorted alphabetically
// ==========================================

var collectionGroupByTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupBy groups strings by first character",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado", "banana", "blueberry", "cherry"},
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "a:2", // groupA
			Second: "b:2", // groupB
			Third:  "c:1", // groupC
		},
	},
	{
		Title: "GroupBy on empty returns empty map",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: []string{},
	},
	{
		Title: "GroupBy single group",
		ArrangeInput: args.Map{
			"items": []string{"ant", "ape", "ace"},
		},
		ExpectedInput: "a:3", // singleGroup
	},
}

// ==========================================
// GroupByCount
// Each expected line is "key:count" sorted alphabetically
// ==========================================

var collectionGroupByCountTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupByCount counts occurrences",
		ArrangeInput: args.Map{
			"items": []string{"red", "blue", "red", "green", "blue", "red"},
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "blue:2",  // blueCount
			Second: "green:1", // greenCount
			Third:  "red:3",   // redCount
		},
	},
	{
		Title: "GroupByCount empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: []string{},
	},
}
