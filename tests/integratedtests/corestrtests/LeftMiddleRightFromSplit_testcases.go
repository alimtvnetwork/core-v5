package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title:         "Normal three-part split",
	ExpectedInput: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitTwoPartsTestCase = coretestcases.CaseV1{
	Title:         "Two parts only — middle empty, invalid",
	ExpectedInput: []string{"a", "", "b", "false"},
}

var leftMiddleRightFromSplitSinglePartTestCase = coretestcases.CaseV1{
	Title:         "Single part — no separator found",
	ExpectedInput: []string{"hello", "", "", "false"},
}

var leftMiddleRightFromSplitFourPlusTestCase = coretestcases.CaseV1{
	Title:         "Four+ parts — middle=second, right=last",
	ExpectedInput: []string{"a", "b", "d", "true"},
}

var leftMiddleRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty input",
	ExpectedInput: []string{"", "", "", "false"},
}

var leftMiddleRightFromSplitEdgesTestCase = coretestcases.CaseV1{
	Title:         "Separator at edges",
	ExpectedInput: []string{"", "", "", "true"},
}

// ==========================================================================
// LeftMiddleRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftMiddleRightFromSplitTrimmedAllTestCase = coretestcases.CaseV1{
	Title:         "Trims whitespace from all three parts",
	ExpectedInput: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitTrimmedTwoTestCase = coretestcases.CaseV1{
	Title:         "Trims with two parts only",
	ExpectedInput: []string{"a", "", "b", "false"},
}

// ==========================================================================
// LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

var leftMiddleRightFromSplitNRemainderTestCase = coretestcases.CaseV1{
	Title:         "SplitN keeps remainder in right",
	ExpectedInput: []string{"a", "b", "c:d:e", "true"},
}

var leftMiddleRightFromSplitNExact3TestCase = coretestcases.CaseV1{
	Title:         "SplitN with exactly 3 parts",
	ExpectedInput: []string{"a", "b", "c", "true"},
}

var leftMiddleRightFromSplitNTwoOnlyTestCase = coretestcases.CaseV1{
	Title:         "SplitN with 2 parts only",
	ExpectedInput: []string{"a", "", "b", "false"},
}

var leftMiddleRightFromSplitNMissingSepTestCase = coretestcases.CaseV1{
	Title:         "SplitN missing separator",
	ExpectedInput: []string{"nosep", "", "", "false"},
}

// ==========================================================================
// LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

var leftMiddleRightFromSplitNTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title:         "SplitN trimmed keeps remainder trimmed",
	ExpectedInput: []string{"a", "b", "c : d : e", "true"},
}

var leftMiddleRightFromSplitNTrimmedTwoTestCase = coretestcases.CaseV1{
	Title:         "SplitN trimmed with 2 parts",
	ExpectedInput: []string{"a", "", "b", "false"},
}
