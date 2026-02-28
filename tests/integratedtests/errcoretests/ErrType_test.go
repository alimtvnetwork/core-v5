package errcorretests

import (
	"errors"
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_ErrType_Combine_Verification(t *testing.T) {
	for caseIndex, testCase := range errTypeCombineTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")
		ref, _ := input.GetAsString("ref")
		errType := errcore.BytesAreNilOrEmptyType

		// Act
		result := errType.Combine(message, ref)

		// Assert
		testCase.ShouldBeRegex(
			t,
			caseIndex,
			result,
		)
	}
}

func Test_ErrCore_MergeErrors_Verification(t *testing.T) {
	for caseIndex, testCase := range errMergeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasError := input.GetDirectLower("hasError")

		var err1, err2 error
		if hasError == true {
			err1 = errors.New("test error")
		}

		// Act
		merged := errcore.MergeErrors(err1, err2)
		isNil := fmt.Sprintf("%v", merged == nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isNil,
		)
	}
}
