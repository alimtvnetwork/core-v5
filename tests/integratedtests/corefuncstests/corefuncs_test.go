package corefuncstests

import (
	"errors"
	"testing"

	"gitlab.com/auk-go/core/corefuncs"
	"gitlab.com/auk-go/core/coretests/args"
)

// sampleFunc is a helper for GetFuncName tests.
func sampleFunc() {}

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncNameTestCases {
		// Act
		shortName := corefuncs.GetFuncName(sampleFunc)
		fullName := corefuncs.GetFuncFullName(sampleFunc)

		// Assert
		actual := args.Map{
			"hasShortName":        len(shortName) > 0,
			"fullLongerThanShort": len(fullName) > len(shortName),
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ActionReturnsErrorFuncWrapper_Success_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrWrapperSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return nil
		})

		// Act
		err := wrapper.Exec()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
			"name":  wrapper.Name,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ActionReturnsErrorFuncWrapper_Failure_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrWrapperFailureTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return errors.New("cleanup failed")
		})

		// Act
		err := wrapper.Exec()

		// Assert
		actual := args.Map{
			"isNil":    err == nil,
			"hasError": err != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsSuccessFuncWrapper_Verification(t *testing.T) {
	results := []bool{true, false}

	for caseIndex, tc := range isSuccessWrapperTestCases {
		// Arrange
		expectedResult := results[caseIndex]
		wrapper := corefuncs.New.IsSuccess("checker", func() bool {
			return expectedResult
		})

		// Act & Assert
		actual := args.Map{
			"result": wrapper.Exec(),
			"name":   wrapper.Name,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_InOutErrFuncWrapperOf_Success_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrWrapperOfSuccessTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"strlen",
			func(s string) (int, error) {
				return len(s), nil
			},
		)

		// Act
		result, err := wrapper.Exec(inputStr)

		// Assert
		actual := args.Map{
			"result": result,
			"isNil":  err == nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_InOutErrFuncWrapperOf_Failure_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrWrapperOfFailureTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
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

		// Assert
		actual := args.Map{
			"result": result,
			"isNil":  err == nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_ActionErr_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorActionErrTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("my-action", func() error {
			return nil
		})

		// Assert
		actual := args.Map{
			"name":      wrapper.Name,
			"hasAction": wrapper.Action != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_IsSuccess_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorIsSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.IsSuccess("my-check", func() bool {
			return true
		})

		// Assert
		actual := args.Map{
			"name":      wrapper.Name,
			"hasAction": wrapper.Action != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
