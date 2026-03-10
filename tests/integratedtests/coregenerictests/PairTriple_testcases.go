package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		ExpectedInput: args.Map{
			"left":         "key",
			"right":        "value",
			"isValid":      true,
			"errorMessage": "",
		},
	},
	{
		Title: "Pair[string,string] empty strings valid",
		ArrangeInput: args.Map{
			"left":  "",
			"right": "",
		},
		ExpectedInput: args.Map{
			"left":         "",
			"right":        "",
			"isValid":      true,
			"errorMessage": "",
		},
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
		ExpectedInput: args.Map{
			"left":         "",
			"right":        "",
			"isValid":      false,
			"errorMessage": "something went wrong",
		},
	},
	{
		Title: "InvalidPairNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: args.Map{
			"left":         "",
			"right":        "",
			"isValid":      false,
			"errorMessage": "",
		},
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
		ExpectedInput: args.Map{
			"clonedLeft":            "original-left",
			"clonedRight":           "original-right",
			"isValid":               true,
			"originalAfterMutation": "mutated-left",
		},
	},
}

// ==========================================
// Pair — nil Clone
// ==========================================

var pairNilCloneTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil pair clone returns nil",
		ArrangeInput:  args.Map{},
		ExpectedInput: "true",
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
	ExpectedInput: "true",
}

var pairIsEqualDiffLeftTestCase = coretestcases.CaseV1{
	Title: "Unequal pairs - different left",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: "false",
}

var pairIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Title: "Nil vs non-nil",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: "false",
}

var pairIsEqualBothNilTestCase = coretestcases.CaseV1{
	Title:         "Both nil",
	ArrangeInput:  args.Map{},
	ExpectedInput: "true",
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
		ExpectedInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
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
		ExpectedInput: args.Map{
			"left":         "a",
			"middle":       "b",
			"right":        "c",
			"isValid":      true,
			"errorMessage": "",
		},
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
		ExpectedInput: args.Map{
			"left":         "",
			"middle":       "",
			"right":        "",
			"isValid":      false,
			"errorMessage": "bad input",
		},
	},
	{
		Title: "InvalidTripleNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: args.Map{
			"left":         "",
			"middle":       "",
			"right":        "",
			"isValid":      false,
			"errorMessage": "",
		},
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
		ExpectedInput: args.Map{
			"clonedLeft":            "L",
			"clonedMiddle":          "M",
			"clonedRight":           "R",
			"isValid":               true,
			"originalAfterMutation": "mutated",
		},
	},
}

// ==========================================
// Triple — nil Clone
// ==========================================

var tripleNilCloneTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil triple clone returns nil",
		ArrangeInput:  args.Map{},
		ExpectedInput: "true",
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
		ExpectedInput: args.Map{
			"left":   "x",
			"middle": "y",
			"right":  "z",
		},
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
		ExpectedInput: args.Map{
			"clearedLeft":  "",
			"clearedRight": "",
			"isValid":      false,
			"errorMessage": "",
		},
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
		ExpectedInput: args.Map{
			"clearedLeft":   "",
			"clearedMiddle": "",
			"clearedRight":  "",
			"isValid":       false,
			"errorMessage":  "",
		},
	},
}
