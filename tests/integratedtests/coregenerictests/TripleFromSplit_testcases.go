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
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "b",
			"right":   "c",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "No separator produces one part invalid",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ".",
		},
		ExpectedInput: args.Map{
			"left":    "nosep",
			"middle":  "",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Two parts only produces invalid triple",
		ArrangeInput: args.Map{
			"input": "a.b",
			"sep":   ".",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "",
			"right":   "b",
			"isValid": false,
			"message": "only two parts found",
		},
	},
	{
		Title: "Four parts uses first second and last",
		ArrangeInput: args.Map{
			"input": "a.b.c.d",
			"sep":   ".",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "b",
			"right":   "d",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Empty input produces one part invalid",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   ".",
		},
		ExpectedInput: args.Map{
			"left":    "",
			"middle":  "",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "x::y::z",
			"sep":   "::",
		},
		ExpectedInput: args.Map{
			"left":    "x",
			"middle":  "y",
			"right":   "z",
			"isValid": true,
			"message": "",
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
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "b",
			"right":   "c",
			"isValid": true,
			"message": "",
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
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "b",
			"right":   "c:d:e",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Exactly three parts",
		ArrangeInput: args.Map{
			"input": "x:y:z",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "x",
			"middle":  "y",
			"right":   "z",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Two parts produces invalid",
		ArrangeInput: args.Map{
			"input": "a:b",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "",
			"right":   "b",
			"isValid": false,
			"message": "only two parts found",
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
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "b",
			"right":   "c : d : e",
			"isValid": true,
			"message": "",
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
		ExpectedInput: args.Map{
			"left":    "L",
			"middle":  "M",
			"right":   "R",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Empty slice produces invalid triple",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		ExpectedInput: args.Map{
			"left":    "",
			"middle":  "",
			"right":   "",
			"isValid": false,
			"message": "empty input",
		},
	},
	{
		Title: "Single-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		ExpectedInput: args.Map{
			"left":    "only",
			"middle":  "",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Two-element produces invalid",
		ArrangeInput: args.Map{
			"parts": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"middle":  "",
			"right":   "b",
			"isValid": false,
			"message": "only two parts found",
		},
	},
}
