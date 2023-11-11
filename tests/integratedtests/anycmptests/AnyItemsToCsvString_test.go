package anycmptests

import (
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_Cmp_All_True_SingleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringSingleQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.AnyItemsToCsvString(
				constants.CommaSpace,
				true, true,
				inputs...))

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
