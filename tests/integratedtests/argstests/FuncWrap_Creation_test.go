package argstests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_FuncWrap_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range funWrapCreationTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.OneFunc)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		actFunc := args.NewFuncWrap
		toStringsConv := coretests.GetAssert.ToStrings

		// Act
		output := actFunc(
			input.First,
		)

		actualSlice.Adds(toStringsConv(output)...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
