package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexCompileTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		regex, err := lazyRegex.Compile()

		isRegexNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)
		isCompiled := fmt.Sprintf("%v", lazyRegex.IsCompiled())

		actLines := []string{isRegexNotNil, hasError, isCompiled}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_HasError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexHasErrorTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		hasError := fmt.Sprintf("%v", lazyRegex.HasError())
		isInvalid := fmt.Sprintf("%v", lazyRegex.IsInvalid())

		actLines := []string{hasError, isInvalid}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_MatchBytes_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchBytesTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		isMatchBytes := fmt.Sprintf("%v", lazyRegex.IsMatchBytes([]byte(compareInput)))
		isFailedMatchBytes := fmt.Sprintf("%v", lazyRegex.IsFailedMatchBytes([]byte(compareInput)))

		actLines := []string{isMatchBytes, isFailedMatchBytes}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchErrorTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		matchErr := lazyRegex.MatchError(compareInput)
		isNoError := fmt.Sprintf("%v", matchErr == nil)

		actLines := []string{isNoError}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_LazyRegex_String_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexStringTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		result := lazyRegex.String()

		actLines := []string{result}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}
