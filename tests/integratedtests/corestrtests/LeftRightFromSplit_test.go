package corestrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
)

// ==========================================================================
// Test: LeftRightFromSplit — edge cases
// ==========================================================================

func Test_LeftRightFromSplit(t *testing.T) {
	// Case 0: Normal key=value split
	{
		tc := leftRightFromSplitTestCases[0]
		lr := corestr.LeftRightFromSplit("key=value", "=")
		tc.ShouldBeEqual(t, 0, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 1: Missing separator
	{
		tc := leftRightFromSplitTestCases[1]
		lr := corestr.LeftRightFromSplit("no-separator-here", "=")
		tc.ShouldBeEqual(t, 1, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 2: Empty input
	{
		tc := leftRightFromSplitTestCases[2]
		lr := corestr.LeftRightFromSplit("", "=")
		tc.ShouldBeEqual(t, 2, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 3: Separator at start
	{
		tc := leftRightFromSplitTestCases[3]
		lr := corestr.LeftRightFromSplit("=value", "=")
		tc.ShouldBeEqual(t, 3, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 4: Separator at end
	{
		tc := leftRightFromSplitTestCases[4]
		lr := corestr.LeftRightFromSplit("key=", "=")
		tc.ShouldBeEqual(t, 4, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 5: Multiple separators
	{
		tc := leftRightFromSplitTestCases[5]
		lr := corestr.LeftRightFromSplit("a=b=c", "=")
		tc.ShouldBeEqual(t, 5, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}
}

// ==========================================================================
// Test: LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

func Test_LeftRightFromSplitTrimmed(t *testing.T) {
	// Case 0: Trims whitespace
	{
		tc := leftRightFromSplitTrimmedTestCases[0]
		lr := corestr.LeftRightFromSplitTrimmed("  key  =  value  ", "=")
		tc.ShouldBeEqual(t, 0, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 1: No separator
	{
		tc := leftRightFromSplitTrimmedTestCases[1]
		lr := corestr.LeftRightFromSplitTrimmed("  hello  ", "=")
		tc.ShouldBeEqual(t, 1, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 2: Whitespace-only parts
	{
		tc := leftRightFromSplitTrimmedTestCases[2]
		lr := corestr.LeftRightFromSplitTrimmed("   =   ", "=")
		tc.ShouldBeEqual(t, 2, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}
}

// ==========================================================================
// Test: LeftRightFromSplitFull — remainder handling
// ==========================================================================

func Test_LeftRightFromSplitFull(t *testing.T) {
	// Case 0: Remainder in right
	{
		tc := leftRightFromSplitFullTestCases[0]
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		tc.ShouldBeEqual(t, 0, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 1: Single separator
	{
		tc := leftRightFromSplitFullTestCases[1]
		lr := corestr.LeftRightFromSplitFull("key:value", ":")
		tc.ShouldBeEqual(t, 1, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 2: Missing separator
	{
		tc := leftRightFromSplitFullTestCases[2]
		lr := corestr.LeftRightFromSplitFull("nosep", ":")
		tc.ShouldBeEqual(t, 2, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}
}

// ==========================================================================
// Test: LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftRightFromSplitFullTrimmed(t *testing.T) {
	// Case 0: Remainder trimmed
	{
		tc := leftRightFromSplitFullTrimmedTestCases[0]
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c : d ", ":")
		tc.ShouldBeEqual(t, 0, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}

	// Case 1: Missing separator trimmed
	{
		tc := leftRightFromSplitFullTrimmedTestCases[1]
		lr := corestr.LeftRightFromSplitFullTrimmed("  hello  ", ":")
		tc.ShouldBeEqual(t, 1, lr.Left, lr.Right, fmt.Sprintf("%v", lr.IsValid))
	}
}
