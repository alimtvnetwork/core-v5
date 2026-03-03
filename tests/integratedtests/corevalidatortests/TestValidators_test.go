package corevalidatortests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

func Test_TestValidators_Verification(t *testing.T) {
	for caseIndex, testCase := range textValidatorsTestCases {
		// Arrange
		parameter := corevalidator.Parameter{
			CaseIndex:                  constants.Zero,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: testCase.IsSkipOnContentsEmpty,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		err := testCase.Validators.AllVerifyErrorMany(
			&parameter,
			testCase.ComparingLines...,
		)

		errorLines := errcore.ErrorToSplitLines(err)

		sliceValidator := corevalidator.SliceValidator{
			Condition:     corevalidator.DefaultDisabledCoreCondition,
			CompareAs:     stringcompareas.Equal,
			ActualLines:   errorLines,
			ExpectedLines: testCase.ExpectationLines,
		}

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:                  caseIndex,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(&nextBaseParam)
		isValid := validationFinalError == nil

		// Assert
		actLines := []string{fmt.Sprintf("%v", isValid)}
		expected := []string{"true"}

		errcore.PrintDiffOnMismatch(caseIndex, testCase.Header, actLines, expected)

		if !isValid {
			errcore.PrintErrorWithTestIndex(
				caseIndex,
				testCase.Header,
				validationFinalError,
			)

			t.Errorf("[case %d] %s: validation failed: %v",
				caseIndex, testCase.Header, validationFinalError)
		}
	}
}
