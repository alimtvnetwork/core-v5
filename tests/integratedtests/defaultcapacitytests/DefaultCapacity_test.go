package defaultcapacitytests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/defaultcapacity"
)

func Test_Predictive_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveTestCases {
		// Arrange
		input := testCase.ArrangeInput.(predictiveInput)

		// Act
		result := defaultcapacity.Predictive(input.possibleLen, input.multiplier, input.additionalCap)
		resultStr := fmt.Sprintf("%d", result)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

func Test_MaxLimit_Verification(t *testing.T) {
	for caseIndex, testCase := range maxLimitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(maxLimitInput)

		// Act
		result := defaultcapacity.MaxLimit(input.wholeLength, input.limit)
		resultStr := fmt.Sprintf("%d", result)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

func Test_OfSearch_Verification(t *testing.T) {
	for caseIndex, testCase := range ofSearchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(int)

		// Act
		result := defaultcapacity.OfSearch(input)
		resultStr := fmt.Sprintf("%d", result)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

func Test_PredictiveDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(int)

		// Act
		result := defaultcapacity.PredictiveDefault(input)
		isPositive := fmt.Sprintf("%v", result > 0)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isPositive)
	}
}
