package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		regex, err := lazyRegex.Compile()

		isRegexNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)
		isCompiled := fmt.Sprintf("%v", lazyRegex.IsCompiled())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isRegexNotNil,
			hasError,
			isCompiled,
		)
	}
}

func Test_LazyRegex_HasError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexHasErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		hasError := fmt.Sprintf("%v", lazyRegex.HasError())
		isInvalid := fmt.Sprintf("%v", lazyRegex.IsInvalid())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			hasError,
			isInvalid,
		)
	}
}

func Test_LazyRegex_MatchBytes_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchBytesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		isMatchBytes := fmt.Sprintf("%v", lazyRegex.IsMatchBytes([]byte(compareInput)))
		isFailedMatchBytes := fmt.Sprintf("%v", lazyRegex.IsFailedMatchBytes([]byte(compareInput)))

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isMatchBytes,
			isFailedMatchBytes,
		)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		matchErr := lazyRegex.MatchError(compareInput)
		isNoError := fmt.Sprintf("%v", matchErr == nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isNoError,
		)
	}
}

func Test_LazyRegex_String_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		result := lazyRegex.String()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result,
		)
	}
}
