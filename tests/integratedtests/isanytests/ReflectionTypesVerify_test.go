package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_Reflection_Types_Verification(t *testing.T) {
	for caseIndex, testCase := range definedTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]interface{})
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				booleanPrintFormatWithType,
				i,
				isany.Defined(input),
				input,
				input)
		}

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.
			TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
