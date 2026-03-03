package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Name:      "Normal key=value split",
	WantLines: []string{"key", "value", "true"},
}

var leftRightFromSplitMissingSepTestCase = coretestcases.CaseV1{
	Name:      "Missing separator returns left only, invalid",
	WantLines: []string{"no-separator-here", "", "false"},
}

var leftRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Name:      "Empty input returns empty left, invalid",
	WantLines: []string{"", "", "false"},
}

var leftRightFromSplitSepAtStartTestCase = coretestcases.CaseV1{
	Name:      "Separator at start gives empty left",
	WantLines: []string{"", "value", "true"},
}

var leftRightFromSplitSepAtEndTestCase = coretestcases.CaseV1{
	Name:      "Separator at end gives empty right",
	WantLines: []string{"key", "", "true"},
}

var leftRightFromSplitMultipleSepTestCase = coretestcases.CaseV1{
	Name:      "Multiple separators keeps first left and last right",
	WantLines: []string{"a", "c", "true"},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTrimsTestCase = coretestcases.CaseV1{
	Name:      "Trims whitespace from both parts",
	WantLines: []string{"key", "value", "true"},
}

var leftRightFromSplitTrimmedNoSepTestCase = coretestcases.CaseV1{
	Name:      "Trims with no separator returns trimmed left, invalid",
	WantLines: []string{"hello", "", "false"},
}

var leftRightFromSplitTrimmedWhitespaceTestCase = coretestcases.CaseV1{
	Name:      "Trims whitespace-only parts to empty",
	WantLines: []string{"", "", "true"},
}

// ==========================================================================
// LeftRightFromSplitFull — remainder handling
// ==========================================================================

var leftRightFromSplitFullRemainderTestCase = coretestcases.CaseV1{
	Name:      "Full split keeps remainder in right",
	WantLines: []string{"a", "b:c:d", "true"},
}

var leftRightFromSplitFullSingleSepTestCase = coretestcases.CaseV1{
	Name:      "Full split with single separator same as normal",
	WantLines: []string{"key", "value", "true"},
}

var leftRightFromSplitFullMissingSepTestCase = coretestcases.CaseV1{
	Name:      "Full split missing separator returns invalid",
	WantLines: []string{"nosep", "", "false"},
}

// ==========================================================================
// LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

var leftRightFromSplitFullTrimmedRemainderTestCase = coretestcases.CaseV1{
	Name:      "Full trimmed split keeps remainder trimmed",
	WantLines: []string{"a", "b : c : d", "true"},
}

var leftRightFromSplitFullTrimmedMissingSepTestCase = coretestcases.CaseV1{
	Name:      "Full trimmed missing separator returns trimmed invalid",
	WantLines: []string{"hello", "", "false"},
}
