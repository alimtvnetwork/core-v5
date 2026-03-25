package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/isany"
)

func Test_DefinedAnyOf_Verification(t *testing.T) {
	for caseIndex, testCase := range definedAnyOfTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([][]any)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringFmt,
				i,
				isany.DefinedAnyOf(input...),
				corecsv.AnyToTypesCsvDefault(input...),
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
