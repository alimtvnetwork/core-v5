package anycmptests

import (
	"testing"

	"gitlab.com/auk-go/core/anycmp"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_Cmp_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToCsvStringSingleQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for i, input := range inputs {
			parameter := input.(coretests.DataHolder)
			actualSlice.AppendFmt(
				"%d : %s (%T, %T)",
				i,
				anycmp.Cmp(
					parameter.First,
					parameter.Second).
					String(),
				parameter.First,
				parameter.Second)
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
