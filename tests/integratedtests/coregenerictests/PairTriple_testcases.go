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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "key",   // left
			Second: "value", // right
			Third:  "true",  // isValid
			Fourth: "",      // errorMessage
		},
	},
	{
		Title: "Pair[string,string] empty strings valid",
		ArrangeInput: args.Map{
			"left":  "",
			"right": "",
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "",     // left
			Second: "",     // right
			Third:  "true", // isValid
			Fourth: "",     // errorMessage
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "",                      // left
			Second: "",                      // right
			Third:  "false",                 // isValid
			Fourth: "something went wrong",  // errorMessage
		},
	},
	{
		Title: "InvalidPairNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "",      // left
			Second: "",      // right
			Third:  "false", // isValid
			Fourth: "",      // errorMessage
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "original-left",  // clonedLeft
			Second: "original-right", // clonedRight
			Third:  "true",           // isValid
			Fourth: "mutated-left",   // originalAfterMutation
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
		ExpectedInput: "true", // isNil
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
		ExpectedInput: args.Two[string, string]{
			First:  "hello", // left
			Second: "world", // right
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "a",    // left
			Second: "b",    // middle
			Third:  "c",    // right
			Fourth: "true", // isValid
			Fifth:  "",     // errorMessage
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "",          // left
			Second: "",          // middle
			Third:  "",          // right
			Fourth: "false",     // isValid
			Fifth:  "bad input", // errorMessage
		},
	},
	{
		Title: "InvalidTripleNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "",      // left
			Second: "",      // middle
			Third:  "",      // right
			Fourth: "false", // isValid
			Fifth:  "",      // errorMessage
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "L",       // clonedLeft
			Second: "M",       // clonedMiddle
			Third:  "R",       // clonedRight
			Fourth: "true",    // isValid
			Fifth:  "mutated", // originalAfterMutation
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
		ExpectedInput: "true", // isNil
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "x", // left
			Second: "y", // middle
			Third:  "z", // right
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "",      // clearedLeft
			Second: "",      // clearedRight
			Third:  "false", // isValid
			Fourth: "",      // errorMessage
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "",      // clearedLeft
			Second: "",      // clearedMiddle
			Third:  "",      // clearedRight
			Fourth: "false", // isValid
			Fifth:  "",      // errorMessage
		},
	},
}
