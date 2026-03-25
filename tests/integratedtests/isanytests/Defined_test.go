package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/isany"
)

func Test_Defined_Verification(t *testing.T) {
	for caseIndex, testCase := range definedTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]any)
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
				input,
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
