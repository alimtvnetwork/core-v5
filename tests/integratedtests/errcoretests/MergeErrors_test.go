package errcoretests

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// buildErrorSlice creates []error from string messages interleaved with nils.
func buildErrorSlice(input args.Map) []error {
	errorMsgs, _ := input.GetAsStrings("errors")
	nilCount, _ := input.GetAsInt("nils")

	var errs []error
	// Interleave: nil first, then errors, then remaining nils
	halfNils := nilCount / 2
	for i := 0; i < halfNils; i++ {
		errs = append(errs, nil)
	}
	for _, msg := range errorMsgs {
		errs = append(errs, errors.New(msg))
	}
	for i := halfNils; i < nilCount; i++ {
		errs = append(errs, nil)
	}

	return errs
}

func Test_SliceToError_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceToErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil, _ := input.GetAsBool("isNil")

		var slice []string
		if !isNil {
			slice, _ = input.GetAsStrings("input")
		}

		// Act
		err := errcore.SliceToError(slice)

		// Assert
		results := []string{fmt.Sprintf("%v", err != nil)}
		if err != nil {
			contain, _ := input.GetAsString("contain")
			results = append(results, fmt.Sprintf("%v", strings.Contains(err.Error(), contain)))
		}

		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_SliceToErrorPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceToErrorPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil, _ := input.GetAsBool("isNil")

		var slice []string
		if !isNil {
			slice, _ = input.GetAsStrings("input")
		}

		// Act
		err := errcore.SliceToErrorPtr(slice)

		// Assert
		results := []string{fmt.Sprintf("%v", err != nil)}
		if err != nil {
			contain, _ := input.GetAsString("contain")
			results = append(results, fmt.Sprintf("%v", strings.Contains(err.Error(), contain)))
		}

		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_MergeErrors_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		merged := errcore.MergeErrors(errs...)

		// Assert
		results := []string{fmt.Sprintf("%v", merged != nil)}
		if merged != nil {
			contain, _ := input.GetAsString("contain")
			results = append(results, fmt.Sprintf("%v", strings.Contains(merged.Error(), contain)))
		}

		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_SliceErrorsToStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceErrorsToStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		result := errcore.SliceErrorsToStrings(errs...)

		// Assert
		results := []string{fmt.Sprintf("%v", len(result))}
		results = append(results, result...)

		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_MergeErrorsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")
		errs := buildErrorSlice(input)

		// Act
		result := errcore.MergeErrorsToString(joiner, errs...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_MergeErrorsToStringDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsToStringDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		result := errcore.MergeErrorsToStringDefault(errs...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
