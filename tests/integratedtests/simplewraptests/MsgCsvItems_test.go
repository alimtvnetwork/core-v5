package simplewraptests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_MsgCsvItems_Verification(t *testing.T) {
	for caseIndex, testCase := range msgCsvItemsTestCases {
		// Arrange
		sliceValidator := corevalidator.SliceValidator{
			ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
		}

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
		testCase.SetActual(finalActual)
		sliceValidator.SetActual(finalActual)
		sliceValidator.ExpectedLines = testCase.ExpectedInput.([]string)

		nextBaseParam := corevalidator.ValidatorParamsBase{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam)

		// Assert
		convey.Convey(testCase.Title, t, func() {
			errcore.PrintErrorWithTestIndex(
				caseIndex,
				testCase.Title,
				validationFinalError)

			convey.So(
				validationFinalError,
				should.BeNil)
		})

		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}
