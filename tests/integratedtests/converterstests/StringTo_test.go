package converterstests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_StringTo_Integer_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Integer(inputStr)
		valueStr := fmt.Sprintf("%v", value)
		hasError := fmt.Sprintf("%v", err != nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			valueStr,
			hasError,
		)
	}
}

func Test_BytesTo_String_Verification(t *testing.T) {
	for caseIndex, testCase := range bytesToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		inputBytes := []byte(inputStr)

		// Act
		result := converters.BytesTo.String(inputBytes)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result,
		)
	}
}
