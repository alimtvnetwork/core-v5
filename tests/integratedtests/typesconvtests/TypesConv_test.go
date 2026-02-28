package typesconvtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/typesconv"
)

func Test_StringToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := typesconv.StringToBool(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IntPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int
		if !isNil {
			v, _ := input.GetAsInt("value")
			ptr = &v
		}
		result := typesconv.IntPtrToSimple(ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
