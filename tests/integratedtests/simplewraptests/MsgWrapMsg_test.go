package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_MsgWrapMsg_Wraps_Verification(t *testing.T) {
	for caseIndex, testCase := range msgWrapsMsgTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		firstMsg := inputs[0]
		secondMsg := inputs[1]

		// Act
		actualSlice.Add(
			simplewrap.MsgWrapMsg(
				firstMsg,
				secondMsg,
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
