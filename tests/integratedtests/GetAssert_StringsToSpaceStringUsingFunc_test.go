package integratedtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_GetAssert_StringsToSpaceStringUsingFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToSpaceStringUsingFuncTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		actFunc := asserter.StringsToSpaceStringUsingFunc

		// Act
		outputs := actFunc(
			input["spaceCount"].(int),
			input["converterFunc"].(coretests.ToLineConverterFunc),
			input["lines"].([]string)...,
		)

		actualSlice.Adds(outputs...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
