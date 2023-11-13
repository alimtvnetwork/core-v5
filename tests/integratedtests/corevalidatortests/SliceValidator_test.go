package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/isany"
)

func Test_SliceValidator(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorsTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]coretests.ArgTwo)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s))
		}

		finalActuals := actualSlice.Strings()

		// Assert
		testCase.Case.AssertEqual(
			t,
			caseIndex,
			finalActuals...)
	}
}
