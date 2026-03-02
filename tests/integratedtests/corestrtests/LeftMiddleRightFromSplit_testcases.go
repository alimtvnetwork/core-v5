package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitTestCases = []coretestcases.CaseV1{
	{
		Name:      "Normal three-part split",
		WantLines: []string{"a", "b", "c", "true"},
	},
	{
		Name:      "Two parts only — middle empty, invalid",
		WantLines: []string{"a", "", "b", "false"},
	},
	{
		Name:      "Single part — no separator found",
		WantLines: []string{"hello", "", "", "false"},
	},
	{
		Name:      "Four+ parts — middle=second, right=last",
		WantLines: []string{"a", "b", "d", "true"},
	},
	{
		Name:      "Empty input",
		WantLines: []string{"", "", "", "false"},
	},
	{
		Name:      "Separator at edges",
		WantLines: []string{"", "", "", "true"},
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftMiddleRightFromSplitTrimmedTestCases = []coretestcases.CaseV1{
	{
		Name:      "Trims whitespace from all three parts",
		WantLines: []string{"a", "b", "c", "true"},
	},
	{
		Name:      "Trims with two parts only",
		WantLines: []string{"a", "", "b", "false"},
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

var leftMiddleRightFromSplitNTestCases = []coretestcases.CaseV1{
	{
		Name:      "SplitN keeps remainder in right",
		WantLines: []string{"a", "b", "c:d:e", "true"},
	},
	{
		Name:      "SplitN with exactly 3 parts",
		WantLines: []string{"a", "b", "c", "true"},
	},
	{
		Name:      "SplitN with 2 parts only",
		WantLines: []string{"a", "", "b", "false"},
	},
	{
		Name:      "SplitN missing separator",
		WantLines: []string{"nosep", "", "", "false"},
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

var leftMiddleRightFromSplitNTrimmedTestCases = []coretestcases.CaseV1{
	{
		Name:      "SplitN trimmed keeps remainder trimmed",
		WantLines: []string{"a", "b", "c : d : e", "true"},
	},
	{
		Name:      "SplitN trimmed with 2 parts",
		WantLines: []string{"a", "", "b", "false"},
	},
}
