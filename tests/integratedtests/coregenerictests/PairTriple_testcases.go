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
		Name: "Pair[string,string] valid",
		ArrangeInput: args.Map{
			"left":  "key",
			"right": "value",
		},
		WantLines: []string{"key", "value", "true", ""},
	},
	{
		Name: "Pair[string,string] empty strings valid",
		ArrangeInput: args.Map{
			"left":  "",
			"right": "",
		},
		WantLines: []string{"", "", "true", ""},
	},
}

// ==========================================
// Pair — InvalidPair
// ==========================================

var pairInvalidTestCases = []coretestcases.CaseV1{
	{
		Name: "InvalidPair with message",
		ArrangeInput: args.Map{
			"message": "something went wrong",
		},
		WantLines: []string{"", "", "false", "something went wrong"},
	},
	{
		Name: "InvalidPairNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		WantLines: []string{"", "", "false", ""},
	},
}

// ==========================================
// Pair — Clone independence
// ==========================================

var pairCloneTestCases = []coretestcases.CaseV1{
	{
		Name: "Clone produces independent copy",
		ArrangeInput: args.Map{
			"left":  "original-left",
			"right": "original-right",
		},
		WantLines: []string{"original-left", "original-right", "true", "mutated-left"},
	},
}

// ==========================================
// Pair — nil Clone
// ==========================================

var pairNilCloneTestCases = []coretestcases.CaseV1{
	{
		Name:         "Nil pair clone returns nil",
		ArrangeInput: args.Map{},
		WantLines:    []string{"true"},
	},
}

// ==========================================
// Pair — IsEqual
// ==========================================

var pairIsEqualSameTestCase = coretestcases.CaseV1{
	Name: "Equal pairs",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	WantLines: []string{"true"},
}

var pairIsEqualDiffLeftTestCase = coretestcases.CaseV1{
	Name: "Unequal pairs - different left",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	WantLines: []string{"false"},
}

var pairIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Name: "Nil vs non-nil",
	ArrangeInput: args.Map{
		"left":  "a",
		"right": "b",
	},
	WantLines: []string{"false"},
}

var pairIsEqualBothNilTestCase = coretestcases.CaseV1{
	Name:         "Both nil",
	ArrangeInput: args.Map{},
	WantLines:    []string{"true"},
}

// ==========================================
// Pair — Values()
// ==========================================

var pairValuesTestCases = []coretestcases.CaseV1{
	{
		Name: "Values returns left and right",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		WantLines: []string{"hello", "world"},
	},
}

// ==========================================
// Triple — NewTriple valid
// ==========================================

var tripleNewValidTestCases = []coretestcases.CaseV1{
	{
		Name: "Triple[string,string,string] valid",
		ArrangeInput: args.Map{
			"left":   "a",
			"middle": "b",
			"right":  "c",
		},
		WantLines: []string{"a", "b", "c", "true", ""},
	},
}

// ==========================================
// Triple — InvalidTriple
// ==========================================

var tripleInvalidTestCases = []coretestcases.CaseV1{
	{
		Name: "InvalidTriple with message",
		ArrangeInput: args.Map{
			"message": "bad input",
		},
		WantLines: []string{"", "", "", "false", "bad input"},
	},
	{
		Name: "InvalidTripleNoMessage",
		ArrangeInput: args.Map{
			"message": "",
		},
		WantLines: []string{"", "", "", "false", ""},
	},
}

// ==========================================
// Triple — Clone
// ==========================================

var tripleCloneTestCases = []coretestcases.CaseV1{
	{
		Name: "Clone produces independent copy",
		ArrangeInput: args.Map{
			"left":   "L",
			"middle": "M",
			"right":  "R",
		},
		WantLines: []string{"L", "M", "R", "true", "mutated"},
	},
}

// ==========================================
// Triple — nil Clone
// ==========================================

var tripleNilCloneTestCases = []coretestcases.CaseV1{
	{
		Name:         "Nil triple clone returns nil",
		ArrangeInput: args.Map{},
		WantLines:    []string{"true"},
	},
}

// ==========================================
// Triple — Values()
// ==========================================

var tripleValuesTestCases = []coretestcases.CaseV1{
	{
		Name: "Values returns all three",
		ArrangeInput: args.Map{
			"left":   "x",
			"middle": "y",
			"right":  "z",
		},
		WantLines: []string{"x", "y", "z"},
	},
}

// ==========================================
// Pair — Clear/Dispose
// ==========================================

var pairClearTestCases = []coretestcases.CaseV1{
	{
		Name: "Clear resets to zero values",
		ArrangeInput: args.Map{
			"left":  "non-empty",
			"right": "non-empty",
		},
		WantLines: []string{"", "", "false", ""},
	},
}

// ==========================================
// Triple — Clear/Dispose
// ==========================================

var tripleClearTestCases = []coretestcases.CaseV1{
	{
		Name: "Clear resets to zero values",
		ArrangeInput: args.Map{
			"left":   "non-empty",
			"middle": "non-empty",
			"right":  "non-empty",
		},
		WantLines: []string{"", "", "", "false", ""},
	},
}
