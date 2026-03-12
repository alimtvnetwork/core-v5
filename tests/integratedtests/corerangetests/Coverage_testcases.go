package corerangetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── MinMaxByte ──

var minMaxByteCases = []coretestcases.CaseV1{
	{
		Title: "MinMaxByte 3-7 -- difference and range length",
		ArrangeInput: args.Map{
			"when": "Min=3, Max=7",
			"min":  3,
			"max":  7,
		},
		ExpectedInput: args.Map{
			"difference":    4,
			"rangeLength":   5,
			"rangeLengthInt": 5,
			"isWithin5":     true,
			"isWithin10":    false,
			"isInvalid10":   true,
			"isOutOfRange10": true,
		},
	},
	{
		Title: "MinMaxByte 0-0 -- zero range",
		ArrangeInput: args.Map{
			"when": "Min=0, Max=0",
			"min":  0,
			"max":  0,
		},
		ExpectedInput: args.Map{
			"difference":    0,
			"rangeLength":   1,
			"rangeLengthInt": 1,
			"isWithin5":     false,
			"isWithin10":    false,
			"isInvalid10":   true,
			"isOutOfRange10": true,
		},
	},
}

var minMaxByteComparisonCases = []coretestcases.CaseV1{
	{
		Title: "MinMaxByte 2-8 -- min/max comparisons",
		ArrangeInput: args.Map{
			"when": "Min=2, Max=8",
			"min":  2,
			"max":  8,
		},
		ExpectedInput: args.Map{
			"isMinEqual2":      true,
			"isMinAboveEqual2": true,
			"isMinAbove1":      true,
			"isMinLess3":       true,
			"isMinLessEqual2":  true,
			"isMaxEqual8":      true,
			"isMaxAboveEqual8": true,
			"isMaxAbove7":      true,
			"isMaxLess9":       true,
			"isMaxLessEqual8":  true,
		},
	},
}

// ── MinMaxInt ──

var minMaxIntCases = []coretestcases.CaseV1{
	{
		Title: "MinMaxInt 3-7 -- difference and range",
		ArrangeInput: args.Map{
			"when": "Min=3, Max=7",
			"min":  3,
			"max":  7,
		},
		ExpectedInput: args.Map{
			"difference":    4,
			"diffAbsolute":  4,
			"rangeLength":   5,
			"isWithin5":     true,
			"isWithin10":    false,
			"isOutOfRange2": true,
			"stringVal":     "3-7",
		},
	},
	{
		Title: "MinMaxInt 7-3 -- negative difference",
		ArrangeInput: args.Map{
			"when": "Min=7, Max=3 (inverted)",
			"min":  7,
			"max":  3,
		},
		ExpectedInput: args.Map{
			"difference":    -4,
			"diffAbsolute":  4,
			"rangeLength":   5,
			"isWithin5":     true,
			"isWithin10":    false,
			"isOutOfRange2": true,
			"stringVal":     "7-3",
		},
	},
}

// ── RangeInt ──

var rangeIntCases = []coretestcases.CaseV1{
	{
		Title: "NewRangeInt valid -- '3:7' within 0-10",
		ArrangeInput: args.Map{
			"when":      "raw=3:7, sep=:, min=0, max=10",
			"raw":       "3:7",
			"separator": ":",
			"min":       0,
			"max":       10,
		},
		ExpectedInput: args.Map{
			"isValid":     true,
			"start":       3,
			"end":         7,
			"rangeLength": 5,
			"difference":  4,
		},
	},
	{
		Title: "NewRangeInt invalid -- '7:3' (end < start)",
		ArrangeInput: args.Map{
			"when":      "raw=7:3, sep=:, min=0, max=10",
			"raw":       "7:3",
			"separator": ":",
			"min":       0,
			"max":       10,
		},
		ExpectedInput: args.Map{
			"isValid":     false,
			"start":       7,
			"end":         3,
			"rangeLength": 5,
			"difference":  -4,
		},
	},
	{
		Title: "NewRangeInt invalid -- single value 'abc'",
		ArrangeInput: args.Map{
			"when":      "raw=abc, sep=:, min=0, max=10",
			"raw":       "abc",
			"separator": ":",
			"min":       0,
			"max":       10,
		},
		ExpectedInput: args.Map{
			"isValid": false,
		},
	},
}

// ── StartEndInt ──

var startEndIntCases = []coretestcases.CaseV1{
	{
		Title: "StartEndInt 3-7 -- basic methods",
		ArrangeInput: args.Map{
			"when":  "Start=3, End=7",
			"start": 3,
			"end":   7,
		},
		ExpectedInput: args.Map{
			"hasStart":           true,
			"hasEnd":             true,
			"isInvalidStart":     false,
			"isInvalidEnd":       false,
			"isStartEndBoth":     true,
			"isInvalidBoth":      false,
			"isInvalidAny":       false,
			"diff":               4,
			"diffAbs":            4,
			"rangeLength":        5,
			"stringVal":          "3-7",
			"stringSpace":        "3 7",
			"stringHyphen":       "3-7",
			"stringColon":        "3:7",
			"isStartGraterThan2": true,
			"isEndGraterThan5":   true,
		},
	},
	{
		Title: "StartEndInt 0-0 -- zero values",
		ArrangeInput: args.Map{
			"when":  "Start=0, End=0",
			"start": 0,
			"end":   0,
		},
		ExpectedInput: args.Map{
			"hasStart":           false,
			"hasEnd":             false,
			"isInvalidStart":     true,
			"isInvalidEnd":       true,
			"isStartEndBoth":     false,
			"isInvalidBoth":      true,
			"isInvalidAny":       true,
			"diff":               0,
			"diffAbs":            0,
			"rangeLength":        1,
			"stringVal":          "0-0",
			"stringSpace":        "0 0",
			"stringHyphen":       "0-0",
			"stringColon":        "0:0",
			"isStartGraterThan2": false,
			"isEndGraterThan5":   false,
		},
	},
}

// ── Within ──

var withinIntegerCases = []coretestcases.CaseV1{
	{
		Title: "RangeInteger in range -- 5 within 1-10",
		ArrangeInput: args.Map{
			"when":  "input=5, min=1, max=10, boundary=true",
			"input": 5,
			"min":   1,
			"max":   10,
		},
		ExpectedInput: args.Map{
			"value":     5,
			"isInRange": true,
		},
	},
	{
		Title: "RangeInteger below min with boundary -- 0 within 1-10",
		ArrangeInput: args.Map{
			"when":  "input=0, min=1, max=10, boundary=true",
			"input": 0,
			"min":   1,
			"max":   10,
		},
		ExpectedInput: args.Map{
			"value":     1,
			"isInRange": false,
		},
	},
	{
		Title: "RangeInteger above max with boundary -- 15 within 1-10",
		ArrangeInput: args.Map{
			"when":  "input=15, min=1, max=10, boundary=true",
			"input": 15,
			"min":   1,
			"max":   10,
		},
		ExpectedInput: args.Map{
			"value":     10,
			"isInRange": false,
		},
	},
}
