package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
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
			))

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
