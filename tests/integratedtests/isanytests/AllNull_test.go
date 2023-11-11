package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_AllNull_Verification(t *testing.T) {
	for caseIndex, testCase := range allNullTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]interface{})
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.AppendFmt(
			"%d : %t (%s)",
			caseIndex,
			isany.AllNull(inputs...),
			converters.Any.s)

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
