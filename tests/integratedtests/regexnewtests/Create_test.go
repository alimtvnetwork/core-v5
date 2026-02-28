package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

func Test_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range createTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		regex, err := regexnew.New.DefaultLock(pattern)
		isCompiled := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isCompiled,
			hasError,
		)
	}
}

func Test_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		isMatch := fmt.Sprintf("%v", regexnew.IsMatchLock(pattern, compareInput))

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isMatch,
		)
	}
}

func Test_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchFailedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		isFailed := fmt.Sprintf("%v", regexnew.IsMatchFailed(pattern, compareInput))

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isFailed,
		)
	}
}

func Test_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range matchErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")
		when := input.When()

		// Act
		var err error
		switch {
		case fmt.Sprintf("%v", when) == "given matching input to MatchErrorLock",
			fmt.Sprintf("%v", when) == "given non-matching input to MatchErrorLock":
			err = regexnew.MatchErrorLock(pattern, compareInput)
		default:
			err = regexnew.MatchError(pattern, compareInput)
		}

		isNoError := fmt.Sprintf("%v", err == nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isNoError,
		)
	}
}
