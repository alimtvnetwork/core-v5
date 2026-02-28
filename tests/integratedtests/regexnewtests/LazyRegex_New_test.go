package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

func Test_New_Lazy_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazyRegex := regexnew.New.Lazy(pattern)

		isDefined := fmt.Sprintf("%v", lazyRegex.IsDefined())
		isApplicable := fmt.Sprintf("%v", lazyRegex.IsApplicable())
		isMatch := fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))
		isFailedMatch := fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			lazyRegex.Pattern(),
			isDefined,
			isApplicable,
			isMatch,
			isFailedMatch,
		)
	}
}

func Test_New_LazyLock_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		isDefined := fmt.Sprintf("%v", lazyRegex.IsDefined())
		isApplicable := fmt.Sprintf("%v", lazyRegex.IsApplicable())
		isMatch := fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))
		isFailedMatch := fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			lazyRegex.Pattern(),
			isDefined,
			isApplicable,
			isMatch,
			isFailedMatch,
		)
	}
}
