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

func Test_MaxByte_Verification(t *testing.T) {
	for caseIndex, testCase := range maxByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxByte(byte(a), byte(b))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_MinByte_Verification(t *testing.T) {
	for caseIndex, testCase := range minByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinByte(byte(a), byte(b))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IntegerWithin_ToByte_Verification(t *testing.T) {
	for caseIndex, testCase := range integerWithinToByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToByte(value)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IntegerWithin_ToInt8_Verification(t *testing.T) {
	for caseIndex, testCase := range integerWithinToInt8TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt8(value)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IntegerOutOfRange_ToByte_Verification(t *testing.T) {
	for caseIndex, testCase := range integerOutOfRangeToByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToByte(value)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IntegerWithin_ToInt16_Verification(t *testing.T) {
	for caseIndex, testCase := range integerWithinToInt16TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt16(value)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
