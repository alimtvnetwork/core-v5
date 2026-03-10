package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns valid triple -- 'a:b:c' three-part split",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTwoPartsTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns invalid -- 'a:b' two parts only",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitSinglePartTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns invalid -- 'hello' no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitFourPlusTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns middle=second, right=last -- 'a:b:c:d' four+ parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "d",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns all empty invalid -- empty string input",
	ExpectedInput: args.Map{
		"left":    "",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitEdgesTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns valid empty parts -- separator at edges",
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
	Title: "LeftMiddleRightFromSplitTrimmed returns trimmed parts -- ' a : b : c ' with whitespace",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplitTrimmed returns invalid -- ' a : b ' two parts with whitespace",
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
	Title: "SplitN returns remainder in right -- 'a:b:c:d:e' keeps 'c:d:e'",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c:d:e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNExact3TestCase = coretestcases.CaseV1{
	Title: "SplitN returns exact parts -- 'a:b:c' exactly 3 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTwoOnlyTestCase = coretestcases.CaseV1{
	Title: "SplitN returns invalid -- 'a:b' only 2 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitNMissingSepTestCase = coretestcases.CaseV1{
	Title: "SplitN returns invalid -- 'nosep' missing separator",
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
	Title: "SplitNTrimmed returns trimmed remainder -- ' a : b : c : d : e ' with whitespace",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c : d : e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "SplitNTrimmed returns invalid -- ' a : b ' only 2 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}
