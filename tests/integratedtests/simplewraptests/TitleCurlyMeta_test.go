package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_TitleCurlyMeta_Wraps_Verification(t *testing.T) {
	for caseIndex, testCase := range titleCurlyMetaTestCases {
		// Arrange
		sliceValidator := corevalidator.SliceValidator{
			Condition: corevalidator.DefaultTrimCoreCondition,
		}

		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		title := inputs[0]
		value := inputs[1]
		meta := inputs[2]
		realMeta := &map[string]string{}
		corejson.Deserialize.FromStringMust(meta, realMeta)

		// Act
		actualSlice.Add(
			simplewrap.TitleCurlyMeta(
				title,
				value,
				meta,
			),
		)

		finalActLines := actualSlice.Strings()
		testCase.SetActual(finalActLines)
		sliceValidator.SetActual(finalActLines)
		sliceValidator.ExpectedLines = testCase.ExpectedInput.([]string)

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam,
		)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				errcore.PrintErrorWithTestIndex(
					caseIndex,
					testCase.Title,
					validationFinalError,
				)

				convey.So(
					validationFinalError,
					should.BeNil,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}

func Test_TitleCurly_Wraps_Verification(t *testing.T) {
	for caseIndex, testCase := range titleCurlyTestCases {
		// Arrange
		sliceValidator := corevalidator.SliceValidator{
			Condition: corevalidator.DefaultTrimCoreCondition,
		}

		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		title := inputs[0]
		value := inputs[1]

		// Act
		actualSlice.Add(
			simplewrap.TitleCurlyWrap(
				title,
				value,
			),
		)

		finalActLines := actualSlice.Strings()
		testCase.SetActual(finalActLines)
		sliceValidator.SetActual(finalActLines)
		sliceValidator.ExpectedLines = testCase.ExpectedInput.([]string)

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam,
		)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				errcore.PrintErrorWithTestIndex(
					caseIndex,
					testCase.Title,
					validationFinalError,
				)

				convey.So(
					validationFinalError,
					should.BeNil,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}
