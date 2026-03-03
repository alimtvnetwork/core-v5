package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

func Test_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range createTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		regex, err := regexnew.New.DefaultLock(pattern)
		isCompiled := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)

		actLines := []string{isCompiled, hasError}

		// Assert

		testCase.ShouldBeEqualUsingExpected(
			t,
			caseIndex,
			actLines,
		)
	}
}

func Test_Create_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		isMatch := fmt.Sprintf("%v", regexnew.IsMatchLock(pattern, compareInput))

		actLines := []string{isMatch}

		// Assert

		testCase.ShouldBeEqualUsingExpected(
			t,
			caseIndex,
			actLines,
		)
	}
}

func Test_Create_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchFailedTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		isFailed := fmt.Sprintf("%v", regexnew.IsMatchFailed(pattern, compareInput))

		actLines := []string{isFailed}

		// Assert

		testCase.ShouldBeEqualUsingExpected(
			t,
			caseIndex,
			actLines,
		)
	}
}

// ==========================================================================
// Test: MatchError
// ==========================================================================

func Test_MatchError_Match(t *testing.T) {
	tc := matchErrorMatchTestCase
	err := regexnew.MatchError(tc.Input, "hello")

	actLines := []string{fmt.Sprintf("%v", err == nil)}

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_MatchError_Mismatch(t *testing.T) {
	tc := matchErrorMismatchTestCase
	err := regexnew.MatchError(tc.Input, "abc")

	actLines := []string{fmt.Sprintf("%v", err == nil)}

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

// ==========================================================================
// Test: MatchErrorLock
// ==========================================================================

func Test_MatchErrorLock_Match(t *testing.T) {
	tc := matchErrorLockMatchTestCase
	err := regexnew.MatchErrorLock(tc.Input, "hello world")

	actLines := []string{fmt.Sprintf("%v", err == nil)}

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_MatchErrorLock_Mismatch(t *testing.T) {
	tc := matchErrorLockMismatchTestCase
	err := regexnew.MatchErrorLock(tc.Input, "abc")

	actLines := []string{fmt.Sprintf("%v", err == nil)}

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}
