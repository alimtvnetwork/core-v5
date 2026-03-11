package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_New_Lazy_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexNewTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.Lazy(pattern)

		actual := args.Map{
			params.patternResult: lazyRegex.Pattern(),
			params.isDefined:     lazyRegex.IsDefined(),
			params.isApplicable:  lazyRegex.IsApplicable(),
			params.isMatch:       lazyRegex.IsMatch(compareInput),
			params.isFailedMatch: lazyRegex.IsFailedMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_New_LazyLock_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.patternResult: lazyRegex.Pattern(),
			params.isDefined:     lazyRegex.IsDefined(),
			params.isApplicable:  lazyRegex.IsApplicable(),
			params.isMatch:       lazyRegex.IsMatch(compareInput),
			params.isFailedMatch: lazyRegex.IsFailedMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
