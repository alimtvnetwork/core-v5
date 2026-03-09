package regexnewtests

import (
	"regexp"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/regexnew"
)

// =============================================================================
// LazyRegex.FullString
// =============================================================================

func Test_LazyRegex_FullString(t *testing.T) {
	for caseIndex, tc := range lazyRegexFullStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		result := lazy.FullString()

		actual := args.Map{
			"isNotEmpty": result != "",
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.CompileMust
// =============================================================================

func Test_LazyRegex_CompileMust(t *testing.T) {
	for caseIndex, tc := range lazyRegexCompileMustTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool
		var result *regexp.Regexp

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			result = lazy.CompileMust()
		}()

		actual := args.Map{
			"regexNotNil": result != nil,
			"panicked":    panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.FirstMatchLine
// =============================================================================

func Test_LazyRegex_FirstMatchLine(t *testing.T) {
	for caseIndex, tc := range lazyRegexFirstMatchLineTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		content, _ := input.GetAsString("content")
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		firstMatch, isInvalidMatch := lazy.FirstMatchLine(content)

		actual := args.Map{
			"firstMatch":     firstMatch,
			"isInvalidMatch": isInvalidMatch,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.IsFailedMatchBytes
// =============================================================================

func Test_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	for caseIndex, tc := range lazyRegexIsFailedMatchBytesTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		inputStr, _ := input.GetAsString("input")
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		isFailed := lazy.IsFailedMatchBytes([]byte(inputStr))

		actual := args.Map{
			"isFailed": isFailed,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.MatchUsingFuncError
// =============================================================================

func Test_LazyRegex_MatchUsingFuncError(t *testing.T) {
	matchFunc := func(regex *regexp.Regexp, lookingTerm string) bool {
		return regex.MatchString(lookingTerm)
	}

	for caseIndex, tc := range lazyRegexMatchUsingFuncErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		comparing, _ := input.GetAsString("comparing")
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		err := lazy.MatchUsingFuncError(comparing, matchFunc)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.OnRequiredCompiledMust
// =============================================================================

func Test_LazyRegex_OnRequiredCompiledMust(t *testing.T) {
	for caseIndex, tc := range lazyRegexOnRequiredCompiledMustTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			lazy.OnRequiredCompiledMust()
		}()

		actual := args.Map{
			"panicked": panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.MustBeSafe
// =============================================================================

func Test_LazyRegex_MustBeSafe(t *testing.T) {
	for caseIndex, tc := range lazyRegexMustBeSafeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			lazy.MustBeSafe()
		}()

		actual := args.Map{
			"panicked": panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
