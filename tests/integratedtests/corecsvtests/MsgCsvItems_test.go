package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_MsgCsvItems_Verification(t *testing.T) {
	for caseIndex, testCase := range msgCsvItemsTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		title := inputs[0].(string)
		csvItems := inputs[1].([]interface{})

		// Act
		actualSlice.Add(
			simplewrap.MsgCsvItems(
				title,
				csvItems...))

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.
			TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
