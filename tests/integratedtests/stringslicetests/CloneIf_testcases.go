package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// stringslice.CloneIf
// =============================================================================

var cloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneIf returns independent copy with length 2 -- isClone true with extra cap",
		ArrangeInput: args.Map{
			"when":          "given isClone true with extra cap",
			"input":         []string{"a", "b"},
			"isClone":       true,
			"additionalCap": 5,
		},
		ExpectedInput: args.Map{
			"resultLength":      "2",
			"item0":             "a",
			"item1":             "b",
			"isIndependentCopy": "true",
		},
	},
	{
		Title: "CloneIf returns original slice reference -- isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []string{"x", "y"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "2",
			"item0":             "x",
			"item1":             "y",
			"isIndependentCopy": "false",
		},
	},
	{
		Title: "CloneIf returns empty non-independent -- nil input with isClone false",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone false",
			"isNil":         true,
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "0",
			"isIndependentCopy": "false",
		},
	},
	{
		Title: "CloneIf returns empty independent copy -- nil input with isClone true",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone true",
			"isNil":         true,
			"isClone":       true,
			"additionalCap": 3,
		},
		ExpectedInput: args.Map{
			"resultLength":      "0",
			"isIndependentCopy": "true",
		},
	},
}

// =============================================================================
// stringslice.AnyItemsCloneIf
// =============================================================================

var anyItemsCloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyItemsCloneIf returns independent copy with length 3 -- isClone true",
		ArrangeInput: args.Map{
			"when":          "given isClone true",
			"input":         []any{"a", 1, true},
			"isClone":       true,
			"additionalCap": 2,
		},
		ExpectedInput: args.Map{
			"resultLength":      "3",
			"item0":             "a",
			"item1":             "1",
			"item2":             "true",
			"isIndependentCopy": "true",
		},
	},
	{
		Title: "AnyItemsCloneIf returns original slice reference -- isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []any{"x"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "1",
			"item0":             "x",
			"isIndependentCopy": "false",
		},
	},
}