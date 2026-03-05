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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "2",    // resultLength
			Second: "a",    // item0
			Third:  "b",    // item1
			Fourth: "true", // isIndependentCopy
		},
	},
	{
		Title: "CloneIf returns original slice when isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []string{"x", "y"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "2",     // resultLength
			Second: "x",     // item0
			Third:  "y",     // item1
			Fourth: "false", // isIndependentCopy
		},
	},
	{
		Title: "CloneIf returns empty on nil input when isClone false",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone false",
			"isNil":         true,
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",     // resultLength
			Second: "false", // isIndependentCopy
		},
	},
	{
		Title: "CloneIf clones nil input when isClone true",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone true",
			"isNil":         true,
			"isClone":       true,
			"additionalCap": 3,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",    // resultLength
			Second: "true", // isIndependentCopy
		},
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",    // resultLength
			Second: "a",    // item0
			Third:  "1",    // item1
			Fourth: "true", // item2
			Fifth:  "true", // isIndependentCopy
		},
	},
	{
		Title: "AnyItemsCloneIf returns original when false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []any{"x"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "1",     // resultLength
			Second: "x",     // item0
			Third:  "false", // isIndependentCopy
		},
	},
}
