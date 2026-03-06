package coresorttests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var intSortQuickTestCases = []coretestcases.CaseV1{
	{
		Title: "Quick sorts integers ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted integers",
			"input": []int{3, 1, 4, 1, 5},
		},
		ExpectedInput: "[1 1 3 4 5]",
	},
	{
		Title: "Quick handles already sorted",
		ArrangeInput: args.Map{
			"when":  "given sorted integers",
			"input": []int{1, 2, 3},
		},
		ExpectedInput: "[1 2 3]",
	},
}

var strSortQuickTestCases = []coretestcases.CaseV1{
	{
		Title: "Quick sorts strings ascending",
		ArrangeInput: args.Map{
			"when":  "given unsorted strings",
			"input": []string{"banana", "apple", "cherry"},
		},
		ExpectedInput: "[apple banana cherry]",
	},
}
