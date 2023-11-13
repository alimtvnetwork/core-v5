package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_AnyItemsToCsvString_All_True_SingleQuotation_Verification(t *testing.T) {
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

func Test_AnyItemsToCsvString_DoubleQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringDoubleQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.AnyItemsToCsvString(
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

func Test_AnyItemsToCsvString_NoQuotation_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringNoQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.AnyItemsToCsvString(
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
