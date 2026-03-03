package corestrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
)

// ==========================================================================
// Test: LeftMiddleRightFromSplit — edge cases
// ==========================================================================

func Test_LeftMiddleRightFromSplit_Normal(t *testing.T) {
	tc := leftMiddleRightFromSplitNormalTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplit_TwoParts(t *testing.T) {
	tc := leftMiddleRightFromSplitTwoPartsTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplit_SinglePart(t *testing.T) {
	tc := leftMiddleRightFromSplitSinglePartTestCase
	lmr := corestr.LeftMiddleRightFromSplit("hello", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplit_FourPlus(t *testing.T) {
	tc := leftMiddleRightFromSplitFourPlusTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c.d", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplit_Empty(t *testing.T) {
	tc := leftMiddleRightFromSplitEmptyTestCase
	lmr := corestr.LeftMiddleRightFromSplit("", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplit_Edges(t *testing.T) {
	tc := leftMiddleRightFromSplitEdgesTestCase
	lmr := corestr.LeftMiddleRightFromSplit("..", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitTrimmed — trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitTrimmed_All(t *testing.T) {
	tc := leftMiddleRightFromSplitTrimmedAllTestCase
	lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplitTrimmed_Two(t *testing.T) {
	tc := leftMiddleRightFromSplitTrimmedTwoTestCase
	lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  ", ".")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

func Test_LeftMiddleRightFromSplitN_Remainder(t *testing.T) {
	tc := leftMiddleRightFromSplitNRemainderTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplitN_Exact3(t *testing.T) {
	tc := leftMiddleRightFromSplitNExact3TestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplitN_TwoOnly(t *testing.T) {
	tc := leftMiddleRightFromSplitNTwoOnlyTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplitN_MissingSep(t *testing.T) {
	tc := leftMiddleRightFromSplitNMissingSepTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("nosep", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitNTrimmed_Remainder(t *testing.T) {
	tc := leftMiddleRightFromSplitNTrimmedRemainderTestCase
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d : e ", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}

func Test_LeftMiddleRightFromSplitNTrimmed_Two(t *testing.T) {
	tc := leftMiddleRightFromSplitNTrimmedTwoTestCase
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b ", ":")
	tc.ShouldBeEqual(t, 0, lmr.Left, lmr.Middle, lmr.Right, fmt.Sprintf("%v", lmr.IsValid))
}
