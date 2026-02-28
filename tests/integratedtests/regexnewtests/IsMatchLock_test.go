package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

func Test_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		result := regexnew.IsMatchLock(pattern, compareInput)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchFailedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		result := regexnew.IsMatchFailed(pattern, compareInput)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_LazyRegex_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.IsMatch(compareInput)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		regex, err := lazy.Compile()
		isNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)
		isApplicable := fmt.Sprintf("%v", lazy.IsApplicable())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotNil, hasError, isApplicable)
	}
}

func Test_LazyRegex_IsFailedMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexIsFailedMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.IsFailedMatch(compareInput)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_LazyRegex_PatternString_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexPatternStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		result := lazy.Pattern()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		err := lazy.MatchError(compareInput)
		isNoError := fmt.Sprintf("%v", err == nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNoError)
	}
}
