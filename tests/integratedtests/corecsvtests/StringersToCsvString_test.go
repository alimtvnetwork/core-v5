package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_StringersToCsvString_All_True_SingleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range stringersToCsvStringTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.StringsToCsvString(
				constants.CommaSpace,
				true,
				true,
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

func Test_StringersToCsvString_DoubleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringDoubleQuoteTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.StringsToCsvString(
				constants.CommaSpace,
				true,
				false,
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

func Test_StringersToCsvString_NoQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringNoQuoteTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.StringsToCsvString(
				constants.CommaSpace,
				false,
				false,
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
