package corestrtests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitTestCases = []coretestcases.CaseV1{
	{
		Name:      "Normal key=value split",
		WantLines: []string{"key", "value", "true"},
	},
	{
		Name:      "Missing separator returns left only, invalid",
		WantLines: []string{"no-separator-here", "", "false"},
	},
	{
		Name:      "Empty input returns empty left, invalid",
		WantLines: []string{"", "", "false"},
	},
	{
		Name:      "Separator at start gives empty left",
		WantLines: []string{"", "value", "true"},
	},
	{
		Name:      "Separator at end gives empty right",
		WantLines: []string{"key", "", "true"},
	},
	{
		Name:      "Multiple separators keeps first left and last right",
		WantLines: []string{"a", "c", "true"},
	},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTestCases = []coretestcases.CaseV1{
	{
		Name:      "Trims whitespace from both parts",
		WantLines: []string{"key", "value", "true"},
	},
	{
		Name:      "Trims with no separator returns trimmed left, invalid",
		WantLines: []string{"hello", "", "false"},
	},
	{
		Name:      "Trims whitespace-only parts to empty",
		WantLines: []string{"", "", "true"},
	},
}

// ==========================================================================
// LeftRightFromSplitFull — remainder handling
// ==========================================================================

var leftRightFromSplitFullTestCases = []coretestcases.CaseV1{
	{
		Name:      "Full split keeps remainder in right",
		WantLines: []string{"a", "b:c:d", "true"},
	},
	{
		Name:      "Full split with single separator same as normal",
		WantLines: []string{"key", "value", "true"},
	},
	{
		Name:      "Full split missing separator returns invalid",
		WantLines: []string{"nosep", "", "false"},
	},
}

// ==========================================================================
// LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

var leftRightFromSplitFullTrimmedTestCases = []coretestcases.CaseV1{
	{
		Name:      "Full trimmed split keeps remainder trimmed",
		WantLines: []string{"a", "b : c : d", "true"},
	},
	{
		Name:      "Full trimmed missing separator returns trimmed invalid",
		WantLines: []string{"hello", "", "false"},
	},
}
