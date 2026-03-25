package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isMatch: regexnew.IsMatchLock(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchFailedTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isFailed: regexnew.IsMatchFailed(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockLazyIsMatchTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazy := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.isMatch: lazy.IsMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_Compile_Lock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockCompileTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		regex, err := lazy.Compile()

		actual := args.Map{
			params.regexNotNil:  regex != nil,
			params.hasError:     err != nil,
			params.isApplicable: lazy.IsApplicable(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_IsFailedMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockIsFailedMatchTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazy := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.isFailedMatch: lazy.IsFailedMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_PatternString_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockPatternStringTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazy := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.patternResult: lazy.Pattern(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_MatchError_Lock_Verification(t *testing.T) {
	for caseIndex, testCase := range isMatchLockMatchErrorTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazy := regexnew.New.LazyLock(pattern)
		err := lazy.MatchError(compareInput)

		actual := args.Map{
			params.isNoError: err == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
