package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
)

func Test_GetAssert_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range quickTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]interface{})
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))
		asserter := coretests.GetAssert
		quickFunc := asserter.QuickGherkins

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				"",
				i,
				(input),
				input,
				input,
			)
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
