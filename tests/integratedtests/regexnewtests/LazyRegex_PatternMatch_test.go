package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

func Test_LazyRegex_PatternMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexPatternMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		compareInput, _ := input.GetAsString("input")
		when := input.When()

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		var actLines []string

		switch {
		case fmt.Sprintf("%v", when) == "given matching input to IsFailedMatch":
			// IsFailedMatch test
			isFailedMatch := fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))
			actLines = append(actLines, isFailedMatch)

		case fmt.Sprintf("%v", when) == "given a pattern with capture group",
			fmt.Sprintf("%v", when) == "given a pattern that does not match":
			// FirstMatchLine test
			firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)
			actLines = append(actLines, firstMatch)
			actLines = append(actLines, fmt.Sprintf("%v", isInvalid))

		default:
			// IsMatch test
			isMatch := fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))
			actLines = append(actLines, isMatch)
		}

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			actLines...,
		)
	}
}
