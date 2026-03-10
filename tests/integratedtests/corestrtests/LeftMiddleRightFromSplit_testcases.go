package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "Normal three-part split",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTwoPartsTestCase = coretestcases.CaseV1{
	Title: "Two parts only — middle empty, invalid",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitSinglePartTestCase = coretestcases.CaseV1{
	Title: "Single part — no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitFourPlusTestCase = coretestcases.CaseV1{
	Title: "Four+ parts — middle=second, right=last",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "d",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "Empty input",
	ExpectedInput: args.Map{
		"left":    "",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitEdgesTestCase = coretestcases.CaseV1{
	Title: "Separator at edges",
	ExpectedInput: args.Map{
		"left":    "",
		"middle":  "",
		"right":   "",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftMiddleRightFromSplitTrimmedAllTestCase = coretestcases.CaseV1{
	Title: "Trims whitespace from all three parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "Trims with two parts only",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

var leftMiddleRightFromSplitNRemainderTestCase = coretestcases.CaseV1{
	Title: "SplitN keeps remainder in right",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c:d:e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNExact3TestCase = coretestcases.CaseV1{
	Title: "SplitN with exactly 3 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTwoOnlyTestCase = coretestcases.CaseV1{
	Title: "SplitN with 2 parts only",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitNMissingSepTestCase = coretestcases.CaseV1{
	Title: "SplitN missing separator",
	ExpectedInput: args.Map{
		"left":    "nosep",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

var leftMiddleRightFromSplitNTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title: "SplitN trimmed keeps remainder trimmed",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c : d : e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "SplitN trimmed with 2 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}
