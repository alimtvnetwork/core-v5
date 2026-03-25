package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexCompileTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		regex, err := lazyRegex.Compile()

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
			params.isCompiled:  lazyRegex.IsCompiled(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_HasError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexHasErrorTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.hasError:  lazyRegex.HasError(),
			params.isInvalid: lazyRegex.IsInvalid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_MatchBytes_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchBytesTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.isMatchBytes:       lazyRegex.IsMatchBytes([]byte(compareInput)),
			params.isFailedMatchBytes: lazyRegex.IsFailedMatchBytes([]byte(compareInput)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchErrorTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		matchErr := lazyRegex.MatchError(compareInput)

		actual := args.Map{
			params.isNoError: matchErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_String_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexStringTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.stringResult: lazyRegex.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
