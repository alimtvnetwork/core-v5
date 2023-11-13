package corevalidatortests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
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

		finalActLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyAllEqual(
			caseIndex,
			finalActLines...)
		validator := testCase.Validator

		// Assert
		convey.Convey(testCase.Case.Title, t, func() {
			validator.AllVerifyError()
			convey.So(
				actualError.Error(),
				should.Equal,
			)
		})
	}
}
