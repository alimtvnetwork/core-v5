package corerangestests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/errcore"
)

func Test_IntRanges_ValidCases(t *testing.T) {
	for _, testCase := range validIntRangeTestCases {
		// Arrange
		arrangeInputs := testCase.Arrange()
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(
				actualRanges,
				should.Equal,
				testCase.ExpectedInput)
		})

		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

func Test_Int_ExceptRanges_Verify(t *testing.T) {
	// Arrange
	arrangeInput := corerange.MinMaxInt{
		Min: 1,
		Max: 15,
	}

	// Act
	actualRanges := arrangeInput.RangesExcept(
		3, 4, 5)

	// Assert
	convey.Convey("Ranges 1-15, RangesExcept(3, 4, 5), should not contain 3,4,5", t, func() {
		convey.So(actualRanges, should.Equal, []int{
			1, 2, 6,
			7, 8, 9,
			10, 11,
			12, 13,
			14, 15,
		})
	})
}

func Test_Int8Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validInt8RangeTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.MinMaxInt8)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(
				actualRanges,
				should.Equal,
				testCase.ExpectedInput)
		})

		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

func Test_StartEndRanges_ValidCases(t *testing.T) {
	for _, testCase := range validStartEndRangesTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(
				actualRanges,
				should.Equal,
				testCase.ExpectedInput)
		})

		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

func Test_StartEndString_Functions_Result_Verification(t *testing.T) {
	type inputWrapper struct {
		Name string
		corerange.StartEndInt
		functions map[string]func() string
	}

	for caseIndex, testCase := range startEndRangesStringFunctionsVerificationTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		var wrappers []inputWrapper
		sliceValidator := testCase.Validator

		for _, input := range arrangeInputs {
			w := inputWrapper{
				Name:        input.String(),
				StartEndInt: input,
				functions:   map[string]func() string{},
			}

			functions := w.functions
			functions["String"] = input.String
			functions["StringColon"] = input.StringColon
			functions["StringHyphen"] = input.StringHyphen
			functions["StringSpace"] = input.StringSpace

			wrappers = append(wrappers, w)
		}

		slice := corestr.New.SimpleSlice.Cap(len(wrappers) * 10)

		// Act
		for i, wrapper := range wrappers {
			slice.AppendFmt("StartEnd : %s", wrapper.Name)

			for funcName, f := range wrapper.functions {
				slice.AppendFmt(
					"    [%d] func : %s | result : %s",
					i,
					funcName,
					f())
			}
		}

		actual := slice.Strings()
		testCase.SetActual(actual)
		sliceValidator.SetActual(actual)

		nextBaseParam := corevalidator.ValidatorParamsBase{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		// Act
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
