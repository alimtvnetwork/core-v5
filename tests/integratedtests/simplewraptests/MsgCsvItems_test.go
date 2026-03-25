package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_MsgCsvItems_Verification(t *testing.T) {
	for caseIndex, testCase := range msgCsvItemsTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]any)
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		title := inputs[0].(string)
		csvItems := inputs[1].([]any)

		// Act
		actualSlice.Add(
			simplewrap.MsgCsvItems(
				title,
				csvItems...,
			),
		)

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
