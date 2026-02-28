package corecmptests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var integerCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Integer returns Equal for same values",
		ArrangeInput: args.Map{
			"when":  "given equal integers",
			"left":  5,
			"right": 5,
		},
		ExpectedInput: []string{
			"Equal",
		},
	},
	{
		Title: "Integer returns LeftLess when left < right",
		ArrangeInput: args.Map{
			"when":  "given left less than right",
			"left":  3,
			"right": 7,
		},
		ExpectedInput: []string{
			"LeftLess",
		},
	},
	{
		Title: "Integer returns LeftGreater when left > right",
		ArrangeInput: args.Map{
			"when":  "given left greater than right",
			"left":  10,
			"right": 2,
		},
		ExpectedInput: []string{
			"LeftGreater",
		},
	},
}

var isStringsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStringsEqual returns true for identical slices",
		ArrangeInput: args.Map{
			"when":  "given identical string slices",
			"left":  []string{"a", "b", "c"},
			"right": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsStringsEqual returns false for different slices",
		ArrangeInput: args.Map{
			"when":  "given different string slices",
			"left":  []string{"a", "b"},
			"right": []string{"a", "c"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsStringsEqual returns false for different lengths",
		ArrangeInput: args.Map{
			"when":  "given slices of different length",
			"left":  []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
