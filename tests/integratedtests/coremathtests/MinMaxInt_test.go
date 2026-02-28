package coremathtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coremath"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_MaxInt_Verification(t *testing.T) {
	for caseIndex, testCase := range maxIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxInt(a, b)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_MinInt_Verification(t *testing.T) {
	for caseIndex, testCase := range minIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinInt(a, b)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
