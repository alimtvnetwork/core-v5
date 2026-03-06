package corestrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================================================
// Test: LeftMiddleRightFromSplit — edge cases
// ==========================================================================

func Test_LeftMiddleRightFromSplit_Normal(t *testing.T) {
	tc := leftMiddleRightFromSplitNormalTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplit_TwoParts(t *testing.T) {
	tc := leftMiddleRightFromSplitTwoPartsTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplit_SinglePart(t *testing.T) {
	tc := leftMiddleRightFromSplitSinglePartTestCase
	lmr := corestr.LeftMiddleRightFromSplit("hello", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplit_FourPlus(t *testing.T) {
	tc := leftMiddleRightFromSplitFourPlusTestCase
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c.d", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplit_Empty(t *testing.T) {
	tc := leftMiddleRightFromSplitEmptyTestCase
	lmr := corestr.LeftMiddleRightFromSplit("", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplit_Edges(t *testing.T) {
	tc := leftMiddleRightFromSplitEdgesTestCase
	lmr := corestr.LeftMiddleRightFromSplit("..", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitTrimmed — trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitTrimmed_All(t *testing.T) {
	tc := leftMiddleRightFromSplitTrimmedAllTestCase
	lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplitTrimmed_Two(t *testing.T) {
	tc := leftMiddleRightFromSplitTrimmedTwoTestCase
	lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  ", ".")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

func Test_LeftMiddleRightFromSplitN_Remainder(t *testing.T) {
	tc := leftMiddleRightFromSplitNRemainderTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplitN_Exact3(t *testing.T) {
	tc := leftMiddleRightFromSplitNExact3TestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplitN_TwoOnly(t *testing.T) {
	tc := leftMiddleRightFromSplitNTwoOnlyTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("a:b", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplitN_MissingSep(t *testing.T) {
	tc := leftMiddleRightFromSplitNMissingSepTestCase
	lmr := corestr.LeftMiddleRightFromSplitN("nosep", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitNTrimmed_Remainder(t *testing.T) {
	tc := leftMiddleRightFromSplitNTrimmedRemainderTestCase
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d : e ", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LeftMiddleRightFromSplitNTrimmed_Two(t *testing.T) {
	tc := leftMiddleRightFromSplitNTrimmedTwoTestCase
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b ", ":")
	actual := args.Map{
		"left":    lmr.Left,
		"middle":  lmr.Middle,
		"right":   lmr.Right,
		"isValid": fmt.Sprintf("%v", lmr.IsValid),
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}
