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
			ArrangeInput.(args.ThreeFunc)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		actFunc := args.NewFuncWrap
		toStringsConv := coretests.GetAssert.ToStrings

		// Act
		funcWrap := actFunc(
			input.WorkFunc,
		)

		output, err := funcWrap.Invoke(
			input.First,
			input.Second,
		)

		actualSlice.Adds(toStringsConv(output)...)

		if err != nil {
			errLines := coretests.
				GetAssert.
				ErrorToLinesWithSpaces(4, err)

			actualSlice.Add(
				"error : ",
			)

			actualSlice.Adds(
				errLines...,
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
