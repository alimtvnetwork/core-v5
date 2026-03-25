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
		Title: "NewPair returns valid Pair -- left 'key', right 'value'",
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
		Title: "NewPair returns valid Pair -- both empty strings",
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
		Title: "InvalidPair returns invalid with message -- 'something went wrong'",
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
		Title: "InvalidPair returns invalid with empty message -- empty string",
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
		Title: "Clone returns independent copy -- left 'original-left', right 'original-right'",
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
		Title:         "Clone returns nil -- nil pair input",
		ArrangeInput:  args.Map{},
		ExpectedInput: "true",
	},
}

// ==========================================
// Pair — IsEqual
// ==========================================

var pairIsEqualSameTestCase = coretestcases.CaseV1{
	Title: "IsEqual returns true -- same left 'a' and right 'b'",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: "true",
}

var pairIsEqualDiffLeftTestCase = coretestcases.CaseV1{
	Title: "IsEqual returns false -- different left values",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: "false",
}

var pairIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Title: "IsEqual returns false -- nil vs non-nil pair",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	ExpectedInput: "false",
}

var pairIsEqualBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqual returns true -- both nil pairs",
	ArrangeInput:  args.Map{},
	ExpectedInput: "true",
}

// ==========================================
// Pair — Values()
// ==========================================

var pairValuesTestCases = []coretestcases.CaseV1{
	{
		Title: "Values returns left and right -- 'hello', 'world'",
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
		Title: "NewTriple returns valid Triple -- 'a', 'b', 'c'",
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
		Title: "InvalidTriple returns invalid with message -- 'bad input'",
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
		Title: "InvalidTriple returns invalid with empty message -- empty string",
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
		Title: "Clone returns independent copy -- Triple('L','M','R')",
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
		Title:         "Clone returns nil -- nil triple input",
		ArrangeInput:  args.Map{},
		ExpectedInput: "true",
	},
}

// ==========================================
// Triple — Values()
// ==========================================

var tripleValuesTestCases = []coretestcases.CaseV1{
	{
		Title: "Values returns all three -- 'x', 'y', 'z'",
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
		Title: "Clear resets Pair to zero values -- from non-empty left and right",
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
		Title: "Clear resets Triple to zero values -- from non-empty left, middle, right",
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
