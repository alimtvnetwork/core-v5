package corecmptests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Integer_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range integerCompareTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsInt("left")
		right, _ := input.GetAsInt("right")

		// Act
		result := corecmp.Integer(left, right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result.String(),
		)
	}
}

func Test_IsStringsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsStrings("left")
		right, _ := input.GetAsStrings("right")

		// Act
		result := corecmp.IsStringsEqual(left, right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}
