package regexnewtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/regexnew"
)

func Test_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		result := regexnew.IsMatchLock(pattern, compareInput)

		actLines := []string{fmt.Sprintf("%v", result)}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchFailedTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		result := regexnew.IsMatchFailed(pattern, compareInput)

		actLines := []string{fmt.Sprintf("%v", result)}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockLazyIsMatchTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.IsMatch(compareInput)

		actLines := []string{fmt.Sprintf("%v", result)}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockCompileTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		regex, err := lazy.Compile()
		isNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)
		isApplicable := fmt.Sprintf("%v", lazy.IsApplicable())

		actLines := []string{isNotNil, hasError, isApplicable}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_IsFailedMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockIsFailedMatchTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.IsFailedMatch(compareInput)

		actLines := []string{fmt.Sprintf("%v", result)}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_PatternString_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockPatternStringTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.Pattern()

		actLines := []string{result}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockMatchErrorTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		err := lazy.MatchError(compareInput)
		isNoError := fmt.Sprintf("%v", err == nil)

		actLines := []string{isNoError}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}
