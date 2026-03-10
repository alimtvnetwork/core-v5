package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		ExpectedInput: args.Map{
			"groupA": "a:2",
			"groupB": "b:2",
			"groupC": "c:1",
		},
	},
	{
		Title: "GroupBy on empty returns empty map",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0",
	},
	{
		Title: "GroupBy single group",
		ArrangeInput: args.Map{
			"items": []string{"ant", "ape", "ace"},
		},
		ExpectedInput: "a:3",
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
		ExpectedInput: args.Map{
			"blueCount":  "blue:2",
			"greenCount": "green:1",
			"redCount":   "red:3",
		},
	},
	{
		Title: "GroupByCount empty returns empty",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0",
	},
}
