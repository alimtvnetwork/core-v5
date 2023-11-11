package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_ParenthesisWrapIf_Wraps_All_Without_Existing_Condition_Checking_Can_Have_DuplicateParenthesis(t *testing.T) {
	for caseIndex, testCase := range parenthesisValidTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.ParenthesisWrapIf(
					true,
					input))
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

func Test_ParenthesisWrapIf_Disabled_Wraps_All_Without_Existing_Condition_Checking_Can_Have_DuplicateParenthesis(t *testing.T) {

	for caseIndex, testCase := range parenthesisDisabledRemainsAsItIsTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		for _, input := range inputs {
			actualSlice.Add(simplewrap.ParenthesisWrapIf(false, input))
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
