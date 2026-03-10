package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/isany"
)

func Test_AnyNull_Verification(t *testing.T) {
	for caseIndex, testCase := range anyNullTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]any)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.AppendFmt(
			defaultCaseIndexBoolStringFmt,
			caseIndex,
			isany.AnyNull(inputs...),
			corecsv.AnyToTypesCsvDefault(inputs...),
		)

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
