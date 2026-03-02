package stringslicetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// stringslice.CloneIf
// =============================================================================

var cloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneIf clones with extra capacity when isClone true",
		ArrangeInput: args.Map{
			"when":          "given isClone true with extra cap",
			"input":         []string{"a", "b"},
			"isClone":       true,
			"additionalCap": 5,
		},
		ExpectedInput: []string{"2", "a", "b", "true"},
	},
	{
		Title: "CloneIf returns original slice when isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []string{"x", "y"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: []string{"2", "x", "y", "false"},
	},
	{
		Title: "CloneIf returns empty on nil input when isClone false",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone false",
			"isNil":         true,
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: []string{"0", "false"},
	},
	{
		Title: "CloneIf clones nil input when isClone true",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone true",
			"isNil":         true,
			"isClone":       true,
			"additionalCap": 3,
		},
		ExpectedInput: []string{"0", "true"},
	},
}

// =============================================================================
// stringslice.AnyItemsCloneIf
// =============================================================================

var anyItemsCloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyItemsCloneIf clones when true",
		ArrangeInput: args.Map{
			"when":          "given isClone true",
			"input":         []any{"a", 1, true},
			"isClone":       true,
			"additionalCap": 2,
		},
		ExpectedInput: []string{"3", "a", "1", "true", "true"},
	},
	{
		Title: "AnyItemsCloneIf returns original when false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []any{"x"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: []string{"1", "x", "false"},
	},
}
