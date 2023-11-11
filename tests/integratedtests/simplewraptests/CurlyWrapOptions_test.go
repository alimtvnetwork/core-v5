package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_CurlyWrapOptions_Wraps_All_CheckConditionally_NoDuplicateCurly(t *testing.T) {
	for caseIndex, testCase := range curlyWrapOptionsValidTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.CurlyWrapOption(
					true, input))
		}

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
