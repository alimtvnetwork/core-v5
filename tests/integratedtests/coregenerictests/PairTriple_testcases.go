package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Pair — NewPair valid
// ==========================================

var pairNewValidTestCases = []coretestcases.CaseV1{
	{
		Title: "Pair[string,string] valid",
		ArrangeInput: args.Map{
			"left":  "key",
			"right": "value",
		},
		ExpectedInput: []string{"key", "value", "true", ""},
	},
	{
		Title: "Pair[string,string] empty strings valid",
		ArrangeInput: args.Map{
			"left":  "",
			"right": "",
		},
		ExpectedInput: []string{"", "", "true", ""},
	},
}

// ==========================================
// Pair — InvalidPair
// ==========================================

var pairInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidPair with message",
		ArrangeInput: args.Map{
			"message": "something went wrong",
		},
		ExpectedInput: []string{"", "", "false", "something went wrong"},
	},
	{
		Title: "InvalidPairNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: []string{"", "", "false", ""},
	},
}

// ==========================================
// Pair — Clone independence
// ==========================================

var pairCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone produces independent copy",
		ArrangeInput: args.Map{
			"left":  "original-left",
			"right": "original-right",
		},
		ExpectedInput: []string{"original-left", "original-right", "true", "mutated-left"},
	},
}

// ==========================================
// Pair — nil Clone
// ==========================================

var pairNilCloneTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil pair clone returns nil",
		ArrangeInput:  args.Map{},
		ExpectedInput: []string{"true"},
	},
}

// ==========================================
// Pair — IsEqual
// ==========================================

var pairIsEqualSameTestCase = coretestcases.CaseV1{
	Title: "Equal pairs",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: []string{"true"},
}

var pairIsEqualDiffLeftTestCase = coretestcases.CaseV1{
	Title: "Unequal pairs - different left",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: []string{"false"},
}

var pairIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Title: "Nil vs non-nil",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: []string{"false"},
}

var pairIsEqualBothNilTestCase = coretestcases.CaseV1{
	Title:         "Both nil",
	ArrangeInput:  args.Map{},
	ExpectedInput: []string{"true"},
}

// ==========================================
// Pair — Values()
// ==========================================

var pairValuesTestCases = []coretestcases.CaseV1{
	{
		Title: "Values returns left and right",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: []string{"hello", "world"},
	},
}

// ==========================================
// Triple — NewTriple valid
// ==========================================

var tripleNewValidTestCases = []coretestcases.CaseV1{
	{
		Title: "Triple[string,string,string] valid",
		ArrangeInput: args.Map{
			"left":   "a",
			"middle": "b",
			"right":  "c",
		},
		ExpectedInput: []string{"a", "b", "c", "true", ""},
	},
}

// ==========================================
// Triple — InvalidTriple
// ==========================================

var tripleInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTriple with message",
		ArrangeInput: args.Map{
			"message": "bad input",
		},
		ExpectedInput: []string{"", "", "", "false", "bad input"},
	},
	{
		Title: "InvalidTripleNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: []string{"", "", "", "false", ""},
	},
}

// ==========================================
// Triple — Clone
// ==========================================

var tripleCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone produces independent copy",
		ArrangeInput: args.Map{
			"left":   "L",
			"middle": "M",
			"right":  "R",
		},
		ExpectedInput: []string{"L", "M", "R", "true", "mutated"},
	},
}

// ==========================================
// Triple — nil Clone
// ==========================================

var tripleNilCloneTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil triple clone returns nil",
		ArrangeInput:  args.Map{},
		ExpectedInput: []string{"true"},
	},
}

// ==========================================
// Triple — Values()
// ==========================================

var tripleValuesTestCases = []coretestcases.CaseV1{
	{
		Title: "Values returns all three",
		ArrangeInput: args.Map{
			"left":   "x",
			"middle": "y",
			"right":  "z",
		},
		ExpectedInput: []string{"x", "y", "z"},
	},
}

// ==========================================
// Pair — Clear/Dispose
// ==========================================

var pairClearTestCases = []coretestcases.CaseV1{
	{
		Title: "Clear resets to zero values",
		ArrangeInput: args.Map{
			"left":  "non-empty",
			"right": "non-empty",
		},
		ExpectedInput: []string{"", "", "false", ""},
	},
}

// ==========================================
// Triple — Clear/Dispose
// ==========================================

var tripleClearTestCases = []coretestcases.CaseV1{
	{
		Title: "Clear resets to zero values",
		ArrangeInput: args.Map{
			"left":   "non-empty",
			"middle": "non-empty",
			"right":  "non-empty",
		},
		ExpectedInput: []string{"", "", "", "false", ""},
	},
}
