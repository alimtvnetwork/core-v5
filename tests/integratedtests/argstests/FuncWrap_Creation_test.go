package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_FuncWrap_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range funWrapCreationTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.ThreeFuncAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(30)
		toStringsConv := coretests.GetAssert.ToStrings

		// Act
		output, err := input.InvokeWithValidArgs()

		actualSlice.Adds(toStringsConv(output)...)

		if err != nil {
			errLines := coretests.
				GetAssert.
				ErrorToLinesWithSpacesDefault(err)

			actualSlice.Add(
				"error : ",
			)

			actualSlice.Adds(
				errLines...,
			)
		}

		finalActLines := actualSlice.Strings()
		actualSlice.Dispose()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
