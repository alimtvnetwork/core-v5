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
		Name: "Standard three-part split",
		ArrangeInput: args.Map{
			"input": "a.b.c",
			"sep":   ".",
		},
		WantLines: []string{
			"a",
			"b",
			"c",
			"true",
			"",
		},
	},
	{
		Name: "No separator produces one part invalid",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ".",
		},
		WantLines: []string{
			"nosep",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Two parts only produces invalid triple",
		ArrangeInput: args.Map{
			"input": "a.b",
			"sep":   ".",
		},
		WantLines: []string{
			"a",
			"",
			"b",
			"false",
			"only two parts found",
		},
	},
	{
		Name: "Four parts uses first second and last",
		ArrangeInput: args.Map{
			"input": "a.b.c.d",
			"sep":   ".",
		},
		WantLines: []string{
			"a",
			"b",
			"d",
			"true",
			"",
		},
	},
	{
		Name: "Empty input produces one part invalid",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   ".",
		},
		WantLines: []string{
			"",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "x::y::z",
			"sep":   "::",
		},
		WantLines: []string{
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
		Name: "Trims whitespace from all three parts",
		ArrangeInput: args.Map{
			"input": "  a  .  b  .  c  ",
			"sep":   ".",
		},
		WantLines: []string{
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
		Name: "Third part gets remainder after second separator",
		ArrangeInput: args.Map{
			"input": "a:b:c:d:e",
			"sep":   ":",
		},
		WantLines: []string{
			"a",
			"b",
			"c:d:e",
			"true",
			"",
		},
	},
	{
		Name: "Exactly three parts",
		ArrangeInput: args.Map{
			"input": "x:y:z",
			"sep":   ":",
		},
		WantLines: []string{
			"x",
			"y",
			"z",
			"true",
			"",
		},
	},
	{
		Name: "Two parts produces invalid",
		ArrangeInput: args.Map{
			"input": "a:b",
			"sep":   ":",
		},
		WantLines: []string{
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
		Name: "Splits at most 3 with trimming",
		ArrangeInput: args.Map{
			"input": "  a  :  b  :  c : d : e  ",
			"sep":   ":",
		},
		WantLines: []string{
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
		Name: "Three-element slice produces valid triple",
		ArrangeInput: args.Map{
			"parts": []string{"L", "M", "R"},
		},
		WantLines: []string{
			"L",
			"M",
			"R",
			"true",
			"",
		},
	},
	{
		Name: "Empty slice produces invalid triple",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		WantLines: []string{
			"",
			"",
			"",
			"false",
			"empty input",
		},
	},
	{
		Name: "Single-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		WantLines: []string{
			"only",
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Two-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"a", "b"},
		},
		WantLines: []string{
			"a",
			"",
			"b",
			"false",
			"only two parts found",
		},
	},
}
