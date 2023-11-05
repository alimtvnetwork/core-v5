package corevalidatortests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/corevalidatortestwrappers"
)

func Test_TestValidators(t *testing.T) {
	for caseIndex, testCase := range corevalidatortestwrappers.TextValidatorsTestCases {
		// Arrange
		paramsBase := corevalidator.ValidatorParamsBase{
			CaseIndex:                         constants.Zero, // fixing test case number here as it is fixed data
			IsIgnoreCompareOnActualInputEmpty: testCase.IsSkipOnContentsEmpty,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   testCase.IsCaseSensitive,
		}

		err := testCase.Validators.AllVerifyErrorMany(
			&paramsBase,
			testCase.ComparingLines...)

		errorLines := errcore.ErrorToSplitLines(
			err)

		sliceValidator := corevalidator.SliceValidator{
			ActualLines:   errorLines,
			ExpectedLines: testCase.ExpectationLines,
			ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
				IsTrimCompare:        false,
				IsNonEmptyWhitespace: false,
				IsSortStringsBySpace: false,
			},
			CompareAs: stringcompareas.Equal,
		}

		nextBaseParam := corevalidator.ValidatorParamsBase{
			CaseIndex:                         caseIndex,
			Header:                            testCase.Header,
			IsIgnoreCompareOnActualInputEmpty: false,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   testCase.IsCaseSensitive,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam)

		isValid := validationFinalError == nil

		// Assert
		convey.Convey(testCase.Header, t, func() {
			errcore.ErrPrintWithTestIndex(
				caseIndex,
				testCase.Header,
				validationFinalError)

			convey.So(isValid, convey.ShouldBeTrue)
		})
	}
}
