package corecmptests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var integerCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Integer returns Equal -- same values 5 and 5",
		ArrangeInput: args.Map{
			"when":  "given equal integers",
			"left":  5,
			"right": 5,
		},
		ExpectedInput: "Equal", // compareResult
	},
	{
		Title: "Integer returns LeftLess -- left 3 right 7",
		ArrangeInput: args.Map{
			"when":  "given left less than right",
			"left":  3,
			"right": 7,
		},
		ExpectedInput: "LeftLess", // compareResult
	},
	{
		Title: "Integer returns LeftGreater -- left 10 right 2",
		ArrangeInput: args.Map{
			"when":  "given left greater than right",
			"left":  10,
			"right": 2,
		},
		ExpectedInput: "LeftGreater", // compareResult
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
		ExpectedInput: "true", // isEqual
	},
	{
		Title: "IsStringsEqual returns false for different slices",
		ArrangeInput: args.Map{
			"when":  "given different string slices",
			"left":  []string{"a", "b"},
			"right": []string{"a", "c"},
		},
		ExpectedInput: "false", // isEqual
	},
	{
		Title: "IsStringsEqual returns false for different lengths",
		ArrangeInput: args.Map{
			"when":  "given slices of different length",
			"left":  []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: "false", // isEqual
	},
}
