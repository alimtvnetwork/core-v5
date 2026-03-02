package corestrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
)

// ==========================================================================
// Test: LeftMiddleRightFromSplit — edge cases
// ==========================================================================

func Test_LeftMiddleRightFromSplit(t *testing.T) {
	// Case 0: Normal three-part split
	{
		tc := leftMiddleRightFromSplitTestCases[0]
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 1: Two parts only
	{
		tc := leftMiddleRightFromSplitTestCases[1]
		lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")
		tc.ShouldBeEqual(t, 1, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 2: Single part
	{
		tc := leftMiddleRightFromSplitTestCases[2]
		lmr := corestr.LeftMiddleRightFromSplit("hello", ".")
		tc.ShouldBeEqual(t, 2, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 3: Four+ parts
	{
		tc := leftMiddleRightFromSplitTestCases[3]
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c.d", ".")
		tc.ShouldBeEqual(t, 3, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 4: Empty input
	{
		tc := leftMiddleRightFromSplitTestCases[4]
		lmr := corestr.LeftMiddleRightFromSplit("", ".")
		tc.ShouldBeEqual(t, 4, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 5: Separator at edges
	{
		tc := leftMiddleRightFromSplitTestCases[5]
		lmr := corestr.LeftMiddleRightFromSplit("..", ".")
		tc.ShouldBeEqual(t, 5, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitTrimmed — trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	// Case 0: Trims all parts
	{
		tc := leftMiddleRightFromSplitTrimmedTestCases[0]
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")
		tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 1: Trims with two parts
	{
		tc := leftMiddleRightFromSplitTrimmedTestCases[1]
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  ", ".")
		tc.ShouldBeEqual(t, 1, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

func Test_LeftMiddleRightFromSplitN(t *testing.T) {
	// Case 0: Remainder in right
	{
		tc := leftMiddleRightFromSplitNTestCases[0]
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 1: Exactly 3 parts
	{
		tc := leftMiddleRightFromSplitNTestCases[1]
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c", ":")
		tc.ShouldBeEqual(t, 1, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 2: Two parts only
	{
		tc := leftMiddleRightFromSplitNTestCases[2]
		lmr := corestr.LeftMiddleRightFromSplitN("a:b", ":")
		tc.ShouldBeEqual(t, 2, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 3: Missing separator
	{
		tc := leftMiddleRightFromSplitNTestCases[3]
		lmr := corestr.LeftMiddleRightFromSplitN("nosep", ":")
		tc.ShouldBeEqual(t, 3, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	// Case 0: Remainder trimmed
	{
		tc := leftMiddleRightFromSplitNTrimmedTestCases[0]
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d : e ", ":")
		tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}

	// Case 1: Two parts trimmed
	{
		tc := leftMiddleRightFromSplitNTrimmedTestCases[1]
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b ", ":")
		tc.ShouldBeEqual(t, 1, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
	}
}
