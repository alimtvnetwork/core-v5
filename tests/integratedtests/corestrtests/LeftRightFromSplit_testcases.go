package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title:         "Normal key=value split",
	ExpectedInput: []string{"key", "value", "true"},
}

var leftRightFromSplitMissingSepTestCase = coretestcases.CaseV1{
	Title:         "Missing separator returns left only, invalid",
	ExpectedInput: []string{"no-separator-here", "", "false"},
}

var leftRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty input returns empty left, invalid",
	ExpectedInput: []string{"", "", "false"},
}

var leftRightFromSplitSepAtStartTestCase = coretestcases.CaseV1{
	Title:         "Separator at start gives empty left",
	ExpectedInput: []string{"", "value", "true"},
}

var leftRightFromSplitSepAtEndTestCase = coretestcases.CaseV1{
	Title:         "Separator at end gives empty right",
	ExpectedInput: []string{"key", "", "true"},
}

var leftRightFromSplitMultipleSepTestCase = coretestcases.CaseV1{
	Title:         "Multiple separators keeps first left and last right",
	ExpectedInput: []string{"a", "c", "true"},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTrimsTestCase = coretestcases.CaseV1{
	Title:         "Trims whitespace from both parts",
	ExpectedInput: []string{"key", "value", "true"},
}

var leftRightFromSplitTrimmedNoSepTestCase = coretestcases.CaseV1{
	Title:         "Trims with no separator returns trimmed left, invalid",
	ExpectedInput: []string{"hello", "", "false"},
}

var leftRightFromSplitTrimmedWhitespaceTestCase = coretestcases.CaseV1{
	Title:         "Trims whitespace-only parts to empty",
	ExpectedInput: []string{"", "", "true"},
}

// ==========================================================================
// LeftRightFromSplitFull — remainder handling
// ==========================================================================

var leftRightFromSplitFullRemainderTestCase = coretestcases.CaseV1{
	Title:         "Full split keeps remainder in right",
	ExpectedInput: []string{"a", "b:c:d", "true"},
}

var leftRightFromSplitFullSingleSepTestCase = coretestcases.CaseV1{
	Title:         "Full split with single separator same as normal",
	ExpectedInput: []string{"key", "value", "true"},
}

var leftRightFromSplitFullMissingSepTestCase = coretestcases.CaseV1{
	Title:         "Full split missing separator returns invalid",
	ExpectedInput: []string{"nosep", "", "false"},
}

// ==========================================================================
// LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

var leftRightFromSplitFullTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title:         "Full trimmed split keeps remainder trimmed",
	ExpectedInput: []string{"a", "b : c : d", "true"},
}

var leftRightFromSplitFullTrimmedMissingSepTestCase = coretestcases.CaseV1{
	Title:         "Full trimmed missing separator returns trimmed invalid",
	ExpectedInput: []string{"hello", "", "false"},
}
