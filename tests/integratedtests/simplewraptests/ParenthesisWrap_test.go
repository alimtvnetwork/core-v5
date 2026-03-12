package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_ParenthesisWrapIf_Enabled_Verification(t *testing.T) {
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
					input,
				),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_ParenthesisWrapIf_Disabled_Verification(t *testing.T) {

	for caseIndex, testCase := range parenthesisDisabledRemainsAsItIsTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		for _, input := range inputs {
			actualSlice.Add(simplewrap.ParenthesisWrapIf(false, input))
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
