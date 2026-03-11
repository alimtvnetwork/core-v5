package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns valid split -- 'key=value' normal input",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "no-separator-here",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty invalid -- empty input string",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitSepAtStartTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty left -- separator at start",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitSepAtEndTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty right -- separator at end",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "",
		"isValid": "true",
	},
}

var leftRightFromSplitMultipleSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns first-left and last-right -- multiple separators",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b=c",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTrimsTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns trimmed parts -- whitespace around both",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitTrimmedNoSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns trimmed invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitTrimmedWhitespaceTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns empty parts -- whitespace-only values",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftRightFromSplitFull — remainder handling
// ==========================================================================

var leftRightFromSplitFullRemainderTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns remainder in right -- 'a:b:c:d' multi-separator",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b:c:d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullSingleSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns same as normal -- single separator",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitFullMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "nosep",
		"right":   "",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

var leftRightFromSplitFullTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFullTrimmed returns trimmed remainder -- ' a : b : c : d ' with spaces",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b : c : d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullTrimmedMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFullTrimmed returns trimmed invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}
