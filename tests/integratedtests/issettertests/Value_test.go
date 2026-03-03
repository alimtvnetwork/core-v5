package issettertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/issetter"
)

func Test_Value_New_Verification(t *testing.T) {
	for caseIndex, testCase := range valueNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		val, err := issetter.New(input)
		hasErr := fmt.Sprintf("%v", err != nil)
		name := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			hasErr,
			name,
		)
	}
}

func Test_Value_GetBool_Verification(t *testing.T) {
	for caseIndex, testCase := range getBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(bool)

		// Act
		val := issetter.GetBool(input)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_NewBool_Verification(t *testing.T) {
	for caseIndex, testCase := range newBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(bool)

		// Act
		val := issetter.NewBool(input)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_BooleanLogic_Verification(t *testing.T) {
	for caseIndex, testCase := range booleanLogicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(issetter.Value)

		// Act
		isOn := fmt.Sprintf("%v", input.IsOn())
		isOff := fmt.Sprintf("%v", input.IsOff())
		isTrue := fmt.Sprintf("%v", input.IsTrue())
		isFalse := fmt.Sprintf("%v", input.IsFalse())
		isSet := fmt.Sprintf("%v", input.IsSet())
		isUnset := fmt.Sprintf("%v", input.IsUnset())
		isValid := fmt.Sprintf("%v", input.IsValid())
		isWildcard := fmt.Sprintf("%v", input.IsWildcard())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isOn,
			isOff,
			isTrue,
			isFalse,
			isSet,
			isUnset,
			isValid,
			isWildcard,
		)
	}
}

func Test_Value_CombinedBooleans_Verification(t *testing.T) {
	for caseIndex, testCase := range combinedBooleansTestCases {
		// Arrange
		input := testCase.ArrangeInput.([]bool)

		// Act
		val := issetter.CombinedBooleans(input...)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_Conversions_Verification(t *testing.T) {
	for caseIndex, testCase := range conversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(issetter.Value)

		// Act
		toBool := input.ToBooleanValue().Name()
		toSetUnset := input.ToSetUnsetValue().Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			toBool,
			toSetUnset,
		)
	}
}

func Test_Value_GetSet_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(getSetInput)

		// Act
		val := issetter.GetSet(input.condition, input.trueVal, input.falseVal)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_IsOutOfRange_Verification(t *testing.T) {
	for caseIndex, testCase := range isOutOfRangeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(byte)

		// Act
		result := fmt.Sprintf("%v", issetter.IsOutOfRange(input))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
