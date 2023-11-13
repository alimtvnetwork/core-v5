package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_StringsToCsvString_All_True_SingleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToCsvStringSingleQuoteTestCases {
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

		finalActuals := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActuals...)
	}
}

func Test_DefaultCsvStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToCsvStringDoubleQuoteTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.DefaultCsv(
				inputs...))

		finalActuals := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActuals...)
	}
}

func Test_StringsToCsvString_DoubleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToCsvStringDoubleQuoteTestCases {
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

		finalActuals := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActuals...)
	}
}

func Test_StringsToCsvString_NoQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToCsvStringNoQuoteTestCases {
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

		finalActuals := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActuals...)
	}
}
