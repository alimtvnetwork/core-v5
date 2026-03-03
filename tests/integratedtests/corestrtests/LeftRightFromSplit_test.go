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
		tc := leftRightFromSplitNormalTestCase
		lr := corestr.LeftRightFromSplit("key=value", "=")
		tc.ShouldBeEqual(t, 0,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 1: Missing separator
	{
		tc := leftRightFromSplitMissingSepTestCase
		lr := corestr.LeftRightFromSplit("no-separator-here", "=")
		tc.ShouldBeEqual(t, 1,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 2: Empty input
	{
		tc := leftRightFromSplitEmptyTestCase
		lr := corestr.LeftRightFromSplit("", "=")
		tc.ShouldBeEqual(t, 2,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 3: Separator at start
	{
		tc := leftRightFromSplitSepAtStartTestCase
		lr := corestr.LeftRightFromSplit("=value", "=")
		tc.ShouldBeEqual(t, 3,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 4: Separator at end
	{
		tc := leftRightFromSplitSepAtEndTestCase
		lr := corestr.LeftRightFromSplit("key=", "=")
		tc.ShouldBeEqual(t, 4,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 5: Multiple separators
	{
		tc := leftRightFromSplitMultipleSepTestCase
		lr := corestr.LeftRightFromSplit("a=b=c", "=")
		tc.ShouldBeEqual(t, 5,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}
}

// ==========================================================================
// Test: LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

func Test_LeftRightFromSplitTrimmed(t *testing.T) {
	// Case 0: Trims whitespace
	{
		tc := leftRightFromSplitTrimmedTrimsTestCase
		lr := corestr.LeftRightFromSplitTrimmed("  key  =  value  ", "=")
		tc.ShouldBeEqual(t, 0,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 1: No separator
	{
		tc := leftRightFromSplitTrimmedNoSepTestCase
		lr := corestr.LeftRightFromSplitTrimmed("  hello  ", "=")
		tc.ShouldBeEqual(t, 1,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 2: Whitespace-only parts
	{
		tc := leftRightFromSplitTrimmedWhitespaceTestCase
		lr := corestr.LeftRightFromSplitTrimmed("   =   ", "=")
		tc.ShouldBeEqual(t, 2,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}
}

// ==========================================================================
// Test: LeftRightFromSplitFull — remainder handling
// ==========================================================================

func Test_LeftRightFromSplitFull(t *testing.T) {
	// Case 0: Remainder in right
	{
		tc := leftRightFromSplitFullRemainderTestCase
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		tc.ShouldBeEqual(t, 0,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 1: Single separator
	{
		tc := leftRightFromSplitFullSingleSepTestCase
		lr := corestr.LeftRightFromSplitFull("key:value", ":")
		tc.ShouldBeEqual(t, 1,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 2: Missing separator
	{
		tc := leftRightFromSplitFullMissingSepTestCase
		lr := corestr.LeftRightFromSplitFull("nosep", ":")
		tc.ShouldBeEqual(t, 2,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}
}

// ==========================================================================
// Test: LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftRightFromSplitFullTrimmed(t *testing.T) {
	// Case 0: Remainder trimmed
	{
		tc := leftRightFromSplitFullTrimmedRemainderTestCase
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c : d ", ":")
		tc.ShouldBeEqual(t, 0,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}

	// Case 1: Missing separator trimmed
	{
		tc := leftRightFromSplitFullTrimmedMissingSepTestCase
		lr := corestr.LeftRightFromSplitFullTrimmed("  hello  ", ":")
		tc.ShouldBeEqual(t, 1,
			lr.Left,
			lr.Right,
			fmt.Sprintf("%v", lr.IsValid),
		)
	}
}
