package corefuncstests

import (
	"errors"
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/corefuncs"
	"gitlab.com/auk-go/core/coretests/args"
)

// sampleFunc is a helper for GetFuncName tests.
func sampleFunc() {}

// ==========================================
// Test: GetFuncName / GetFuncFullName
// ==========================================

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, testCase := range getFuncNameTestCases {
		// Arrange — use sampleFunc

		// Act
		shortName := corefuncs.GetFuncName(sampleFunc)
		fullName := corefuncs.GetFuncFullName(sampleFunc)
		actLines := []string{
			fmt.Sprintf("%v", len(shortName) > 0),
			fmt.Sprintf("%v", len(fullName) > len(shortName)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ActionReturnsErrorFuncWrapper — success
// ==========================================

func Test_ActionReturnsErrorFuncWrapper_Success_Verification(t *testing.T) {
	for caseIndex, testCase := range actionErrWrapperSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return nil
		})

		// Act
		err := wrapper.Exec()
		actLines := []string{
			fmt.Sprintf("%v", err == nil),
			wrapper.Name,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ActionReturnsErrorFuncWrapper — failure
// ==========================================

func Test_ActionReturnsErrorFuncWrapper_Failure_Verification(t *testing.T) {
	for caseIndex, testCase := range actionErrWrapperFailureTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return errors.New("cleanup failed")
		})

		// Act
		err := wrapper.Exec()
		actLines := []string{
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%v", err != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsSuccessFuncWrapper
// ==========================================

func Test_IsSuccessFuncWrapper_Verification(t *testing.T) {
	results := []bool{true, false}

	for caseIndex, testCase := range isSuccessWrapperTestCases {
		// Arrange
		expectedResult := results[caseIndex]
		wrapper := corefuncs.New.IsSuccess("checker", func() bool {
			return expectedResult
		})

		// Act
		actLines := []string{
			fmt.Sprintf("%v", wrapper.Exec()),
			wrapper.Name,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InOutErrFuncWrapperOf — success
// ==========================================

func Test_InOutErrFuncWrapperOf_Success_Verification(t *testing.T) {
	for caseIndex, testCase := range inOutErrWrapperOfSuccessTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"strlen",
			func(s string) (int, error) {
				return len(s), nil
			},
		)

		// Act
		result, err := wrapper.Exec(inputStr)
		actLines := []string{
			fmt.Sprintf("%d", result),
			fmt.Sprintf("%v", err == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InOutErrFuncWrapperOf — failure
// ==========================================

func Test_InOutErrFuncWrapperOf_Failure_Verification(t *testing.T) {
	for caseIndex, testCase := range inOutErrWrapperOfFailureTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"strlen",
			func(s string) (int, error) {
				if s == "" {
					return 0, errors.New("empty input")
				}
				return len(s), nil
			},
		)

		// Act
		result, err := wrapper.Exec(inputStr)
		actLines := []string{
			fmt.Sprintf("%d", result),
			fmt.Sprintf("%v", err == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New creator factory — ActionErr
// ==========================================

func Test_NewCreator_ActionErr_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorActionErrTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("my-action", func() error {
			return nil
		})

		// Act
		actLines := []string{
			wrapper.Name,
			fmt.Sprintf("%v", wrapper.Action != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New creator factory — IsSuccess
// ==========================================

func Test_NewCreator_IsSuccess_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorIsSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.IsSuccess("my-check", func() bool {
			return true
		})

		// Act
		actLines := []string{
			wrapper.Name,
			fmt.Sprintf("%v", wrapper.Action != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
