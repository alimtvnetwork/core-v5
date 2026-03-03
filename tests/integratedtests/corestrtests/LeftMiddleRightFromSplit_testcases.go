package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Name:      "Normal three-part split",
	WantLines: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitTwoPartsTestCase = coretestcases.CaseV1{
	Name:      "Two parts only — middle empty, invalid",
	WantLines: []string{"a", "", "b", "false"},
}

var leftMiddleRightFromSplitSinglePartTestCase = coretestcases.CaseV1{
	Name:      "Single part — no separator found",
	WantLines: []string{"hello", "", "", "false"},
}

var leftMiddleRightFromSplitFourPlusTestCase = coretestcases.CaseV1{
	Name:      "Four+ parts — middle=second, right=last",
	WantLines: []string{"a", "b", "d", "true"},
}

var leftMiddleRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Name:      "Empty input",
	WantLines: []string{"", "", "", "false"},
}

var leftMiddleRightFromSplitEdgesTestCase = coretestcases.CaseV1{
	Name:      "Separator at edges",
	WantLines: []string{"", "", "", "true"},
}

// ==========================================================================
// LeftMiddleRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftMiddleRightFromSplitTrimmedAllTestCase = coretestcases.CaseV1{
	Name:      "Trims whitespace from all three parts",
	WantLines: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitTrimmedTwoTestCase = coretestcases.CaseV1{
	Name:      "Trims with two parts only",
	WantLines: []string{"a", "", "b", "false"},
}

// ==========================================================================
// LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

var leftMiddleRightFromSplitNRemainderTestCase = coretestcases.CaseV1{
	Name:      "SplitN keeps remainder in right",
	WantLines: []string{"a", "b", "c:d:e", "true"},
}

var leftMiddleRightFromSplitNExact3TestCase = coretestcases.CaseV1{
	Name:      "SplitN with exactly 3 parts",
	WantLines: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitNTwoOnlyTestCase = coretestcases.CaseV1{
	Name:      "SplitN with 2 parts only",
	WantLines: []string{"a", "", "b", "false"},
}

var leftMiddleRightFromSplitNMissingSepTestCase = coretestcases.CaseV1{
	Name:      "SplitN missing separator",
	WantLines: []string{"nosep", "", "", "false"},
}

// ==========================================================================
// LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

var leftMiddleRightFromSplitNTrimmedRemainderTestCase = coretestcases.CaseV1{
	Name:      "SplitN trimmed keeps remainder trimmed",
	WantLines: []string{"a", "b", "c : d : e", "true"},
}

var leftMiddleRightFromSplitNTrimmedTwoTestCase = coretestcases.CaseV1{
	Name:      "SplitN trimmed with 2 parts",
	WantLines: []string{"a", "", "b", "false"},
}
