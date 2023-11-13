package corevalidatortests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

func Test_SliceValidator(t *testing.T) {
	for caseIndex, testCase := range textValidatorsTestCases {
		// Arrange
		paramsBase := corevalidator.Parameter{
			CaseIndex:                  caseIndex, // fixing test case number here as it is fixed data
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: testCase.IsSkipOnContentsEmpty,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		err := testCase.Validators.AllVerifyErrorMany(
			&paramsBase,
			testCase.ComparingLines...)

		errorLines := errcore.ErrorToSplitLines(
			err)

		sliceValidator := corevalidator.SliceValidator{
			ValidatorCoreCondition: corevalidator.DefaultDisabledCoreCondition,
			CompareAs:              stringcompareas.Equal,
			ActualLines:            errorLines,
			ExpectedLines:          testCase.ExpectationLines,
		}

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:                  caseIndex,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam)

		isValid := validationFinalError == nil

		// Assert
		convey.Convey(testCase.Header, t, func() {
			errcore.PrintErrorWithTestIndex(
				caseIndex,
				testCase.Header,
				validationFinalError)

			convey.So(isValid, convey.ShouldBeTrue)
		})
	}
}
