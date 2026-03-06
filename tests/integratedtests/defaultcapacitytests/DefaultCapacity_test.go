package defaultcapacitytests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/defaultcapacity"
)

func Test_Predictive_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		possibleLen := input.GetAsIntDefault("possibleLen", 0)
		multiplierRaw, _ := input.Get("multiplier")
		multiplier := multiplierRaw.(float64)
		additionalCap := input.GetAsIntDefault("additionalCap", 0)

		// Act
		result := defaultcapacity.Predictive(possibleLen, multiplier, additionalCap)
		resultStr := fmt.Sprintf("%d", result)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

func Test_MaxLimit_Verification(t *testing.T) {
	for caseIndex, testCase := range maxLimitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		wholeLength := input.GetAsIntDefault("wholeLength", 0)
		limit := input.GetAsIntDefault("limit", 0)

		// Act
		result := defaultcapacity.MaxLimit(wholeLength, limit)
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
