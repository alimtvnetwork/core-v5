package integratedtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_GetAssert_SortedMessage_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedMessageTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.SortedMessage

		// Act
		outputs := quickFunc(
			input["isPrint"].(bool),
			input["message"].(string),
			input["joiner"].(string),
		)

		actualSlice.Add(outputs)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
