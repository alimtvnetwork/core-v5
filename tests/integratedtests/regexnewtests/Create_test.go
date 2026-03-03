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

// ==========================================================================
// Test: MatchError
// ==========================================================================

func Test_MatchError_Match(t *testing.T) {
	tc := matchErrorMatchTestCase
	err := regexnew.MatchError("^hello$", "hello")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err == nil))
}

func Test_MatchError_Mismatch(t *testing.T) {
	tc := matchErrorMismatchTestCase
	err := regexnew.MatchError("^\\d+$", "abc")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err == nil))
}

// ==========================================================================
// Test: MatchErrorLock
// ==========================================================================

func Test_MatchErrorLock_Match(t *testing.T) {
	tc := matchErrorLockMatchTestCase
	err := regexnew.MatchErrorLock("world", "hello world")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err == nil))
}

func Test_MatchErrorLock_Mismatch(t *testing.T) {
	tc := matchErrorLockMismatchTestCase
	err := regexnew.MatchErrorLock("^xyz$", "abc")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err == nil))
}
