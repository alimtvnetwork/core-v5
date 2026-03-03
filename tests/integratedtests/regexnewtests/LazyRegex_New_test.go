package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

func Test_New_Lazy_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexNewTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazyRegex := regexnew.New.Lazy(pattern)

		isDefined := fmt.Sprintf("%v", lazyRegex.IsDefined())
		isApplicable := fmt.Sprintf("%v", lazyRegex.IsApplicable())
		isMatch := fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))
		isFailedMatch := fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))

		actLines := []string{
			lazyRegex.Pattern(),
			isDefined,
			isApplicable,
			isMatch,
			isFailedMatch,
		}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_New_LazyLock_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		isDefined := fmt.Sprintf("%v", lazyRegex.IsDefined())
		isApplicable := fmt.Sprintf("%v", lazyRegex.IsApplicable())
		isMatch := fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))
		isFailedMatch := fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))

		actLines := []string{
			lazyRegex.Pattern(),
			isDefined,
			isApplicable,
			isMatch,
			isFailedMatch,
		}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}
