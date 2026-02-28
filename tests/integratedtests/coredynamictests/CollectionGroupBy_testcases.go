package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// GroupBy — group by first character
// ==========================================

var collectionGroupByTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupBy groups strings by first character",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado", "banana", "blueberry", "cherry"},
		},
		ExpectedInput: []string{"a:2", "b:2", "c:1"},
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
		ExpectedInput: []string{"a:3"},
	},
}

// ==========================================
// GroupByCount
// ==========================================

var collectionGroupByCountTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupByCount counts occurrences",
		ArrangeInput: args.Map{
			"items": []string{"red", "blue", "red", "green", "blue", "red"},
		},
		ExpectedInput: []string{"blue:2", "green:1", "red:3"},
	},
	{
		Title: "GroupByCount empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: []string{},
	},
}
