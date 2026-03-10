package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "Normal key=value split",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitMissingSepTestCase = coretestcases.CaseV1{
	Title: "Missing separator returns left only, invalid",
	ExpectedInput: args.Map{
		"left":    "no-separator-here",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "Empty input returns empty left, invalid",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitSepAtStartTestCase = coretestcases.CaseV1{
	Title: "Separator at start gives empty left",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitSepAtEndTestCase = coretestcases.CaseV1{
	Title: "Separator at end gives empty right",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "",
		"isValid": "true",
	},
}

var leftRightFromSplitMultipleSepTestCase = coretestcases.CaseV1{
	Title: "Multiple separators keeps first left and last right",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "c",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTrimsTestCase = coretestcases.CaseV1{
	Title: "Trims whitespace from both parts",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitTrimmedNoSepTestCase = coretestcases.CaseV1{
	Title: "Trims with no separator returns trimmed left, invalid",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitTrimmedWhitespaceTestCase = coretestcases.CaseV1{
	Title: "Trims whitespace-only parts to empty",
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
	Title: "Full split keeps remainder in right",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b:c:d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullSingleSepTestCase = coretestcases.CaseV1{
	Title: "Full split with single separator same as normal",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitFullMissingSepTestCase = coretestcases.CaseV1{
	Title: "Full split missing separator returns invalid",
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
	Title: "Full trimmed split keeps remainder trimmed",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b : c : d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullTrimmedMissingSepTestCase = coretestcases.CaseV1{
	Title: "Full trimmed missing separator returns trimmed invalid",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}
