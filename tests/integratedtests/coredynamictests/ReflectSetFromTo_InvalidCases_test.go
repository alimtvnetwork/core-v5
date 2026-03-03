package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/trydo"
	"gitlab.com/auk-go/core/tests/testwrappers/coredynamictestwrappers"
)

// Test_ReflectSetFromTo_ValidCases
//
// Valid Inputs:
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
func Test_ReflectSetFromTo_Invalid_Cases_With_Error_Verifications(t *testing.T) {
	for caseIndex, testCase := range coredynamictestwrappers.ReflectSetFromToInvalidTestCases {
		// Act
		wrappedResult := trydo.ErrorFuncWrapPanic(
			func() error {
				return coredynamic.ReflectSetFromTo(
					testCase.From,
					testCase.To,
				)
			},
		)

		err := wrappedResult.Error
		testCase.SetActual(wrappedResult)

		// Assert - error expectation
		hasErr := fmt.Sprintf("%v", err != nil)
		expectedHasErr := fmt.Sprintf("%v", testCase.IsExpectingError)

		actLines := []string{hasErr}
		expected := []string{expectedHasErr}

		// Assert - validator verification
		parameter := &corevalidator.Parameter{
			CaseIndex:                  caseIndex,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            true,
		}

		finalErr := getFinalVerificationError(
			testCase,
			testCase.Validator,
			parameter,
			wrappedResult,
		)

		actLines = append(actLines, fmt.Sprintf("%v", finalErr == nil))
		expected = append(expected, "true")

		errcore.PrintDiffOnMismatch(caseIndex, testCase.Header, actLines, expected)

		for i, act := range actLines {
			if act != expected[i] {
				t.Errorf("[case %d] %s: line %d got %q, want %q",
					caseIndex, testCase.Header, i, act, expected[i])
			}
		}
	}
}

func getFinalVerificationError(
	testCase coredynamictestwrappers.FromToTestWrapper,
	validator corevalidator.TextValidator,
	parameter *corevalidator.Parameter,
	wrappedResult trydo.WrappedErr,
) error {
	if testCase.HasPanic {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ExceptionString(),
		)
	}

	if testCase.IsExpectingError {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ErrorString(),
		)
	}

	return validator.VerifyDetailError(
		parameter,
		converters.AnyTo.ValueString(testCase.ExpectedValue),
	)
}
