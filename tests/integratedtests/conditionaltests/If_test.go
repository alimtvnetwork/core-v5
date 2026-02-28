package conditionaltests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_If_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ifStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsString("trueValue")
		falseValue, _ := input.GetAsString("falseValue")

		// Act
		result := conditional.If[string](isTrue, trueValue, falseValue)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result,
		)
	}
}

func Test_If_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range ifIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsInt("trueValue")
		falseValue, _ := input.GetAsInt("falseValue")

		// Act
		result := conditional.If[int](isTrue, trueValue, falseValue)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}

func Test_NilDef_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilDef[string](ptr, defVal)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result,
		)
	}
}
