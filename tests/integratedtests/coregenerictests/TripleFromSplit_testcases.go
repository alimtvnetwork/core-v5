package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// TripleFromSplit
// ==========================================

var tripleFromSplitTestCases = []coretestcases.CaseV1{
	{
		Title: "Standard three-part split",
		ArrangeInput: args.Map{
			"input": "a.b.c",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"a",
			"b",
			"c",
			"true",
			"",
		},
	},
	{
		Title: "No separator produces one part invalid",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"nosep",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Two parts only produces invalid triple",
		ArrangeInput: args.Map{
			"input": "a.b",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"a",
			"",
			"b",
			"false",
			"only two parts found",
		},
	},
	{
		Title: "Four parts uses first second and last",
		ArrangeInput: args.Map{
			"input": "a.b.c.d",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"a",
			"b",
			"d",
			"true",
			"",
		},
	},
	{
		Title: "Empty input produces one part invalid",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "x::y::z",
			"sep":   "::",
		},
		ExpectedInput: []string{
			"x",
			"y",
			"z",
			"true",
			"",
		},
	},
}

// ==========================================
// TripleFromSplitTrimmed
// ==========================================

var tripleFromSplitTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Trims whitespace from all three parts",
		ArrangeInput: args.Map{
			"input": "  a  .  b  .  c  ",
			"sep":   ".",
		},
		ExpectedInput: []string{
			"a",
			"b",
			"c",
			"true",
			"",
		},
	},
}

// ==========================================
// TripleFromSplitN
// ==========================================

var tripleFromSplitNTestCases = []coretestcases.CaseV1{
	{
		Title: "Third part gets remainder after second separator",
		ArrangeInput: args.Map{
			"input": "a:b:c:d:e",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"a",
			"b",
			"c:d:e",
			"true",
			"",
		},
	},
	{
		Title: "Exactly three parts",
		ArrangeInput: args.Map{
			"input": "x:y:z",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"x",
			"y",
			"z",
			"true",
			"",
		},
	},
	{
		Title: "Two parts produces invalid",
		ArrangeInput: args.Map{
			"input": "a:b",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"a",
			"",
			"b",
			"false",
			"only two parts found",
		},
	},
}

// ==========================================
// TripleFromSplitNTrimmed
// ==========================================

var tripleFromSplitNTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Splits at most 3 with trimming",
		ArrangeInput: args.Map{
			"input": "  a  :  b  :  c : d : e  ",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"a",
			"b",
			"c : d : e",
			"true",
			"",
		},
	},
}

// ==========================================
// TripleFromSlice
// ==========================================

var tripleFromSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "Three-element slice produces valid triple",
		ArrangeInput: args.Map{
			"parts": []string{"L", "M", "R"},
		},
		ExpectedInput: []string{
			"L",
			"M",
			"R",
			"true",
			"",
		},
	},
	{
		Title: "Empty slice produces invalid triple",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		ExpectedInput: []string{
			"",
			"",
			"",
			"false",
			"empty input",
		},
	},
	{
		Title: "Single-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		ExpectedInput: []string{
			"only",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Two-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"a", "b"},
		},
		ExpectedInput: []string{
			"a",
			"",
			"b",
			"false",
			"only two parts found",
		},
	},
}
