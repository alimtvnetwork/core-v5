package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_SquareWrapIf_Enabled_Verification(
	t *testing.T,
) {
	for caseIndex, testCase := range squareBracketWrapTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.SquareWrapIf(
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

func Test_SquareWrapIf_Disabled_Wraps_Nothing(
	t *testing.T,
) {
	for caseIndex, testCase := range squareBracketWrapDisabledTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.SquareWrapIf(
					false,
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
