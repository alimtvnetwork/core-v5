package anycmptests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_RangeNamesWithValuesIndexes_Verification(t *testing.T) {
	for caseIndex, testCase := range rangeNamesWithValuesIndexesTestCases {
		// Arrange
		inputs := testCase.Arrange()

		// Act
		finalActual := corecsv.RangeNamesWithValuesIndexes(
			inputs...)

		finalTestCase := coretestcases.
			TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}

func Test_RangeNamesWithValuesIndexesCsvString_Verification(t *testing.T) {
	for caseIndex, testCase := range rangeNamesWithValuesIndexesTestCases {
		// Arrange
		inputs := testCase.Arrange()
		expectation := "some val at 0[0], " +
			"some val at 1[1], " +
			"some val at 2[2], " +
			"Alim Ul Karim[3], " +
			"Where It is[4], " +
			"[5]"

		// Act
		finalActual := corecsv.RangeNamesWithValuesIndexesCsvString(
			inputs...)

		finalTestCase := coretestcases.
			TestCaseV1(testCase.BaseTestCase)
		finalTestCase.SetExpected([]string{
			expectation,
		})

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual)
	}
}
