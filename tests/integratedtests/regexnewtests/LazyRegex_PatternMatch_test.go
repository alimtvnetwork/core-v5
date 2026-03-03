package regexnewtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

// ==========================================================================
// Test: IsMatch
// ==========================================================================

func Test_LazyRegex_IsMatch_FullDigit(t *testing.T) {
	tc := lazyRegexPatternMatchTestCases[0]
	lazyRegex := regexnew.New.LazyLock("^\\d+$")

	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsMatch("12345"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_LazyRegex_IsMatch_PartialMismatch(t *testing.T) {
	tc := lazyRegexPatternMatchTestCases[1]
	lazyRegex := regexnew.New.LazyLock("^\\d+$")

	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsMatch("123abc"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: IsFailedMatch
// ==========================================================================

func Test_LazyRegex_IsFailedMatch(t *testing.T) {
	tc := lazyRegexPatternMatchTestCases[2]
	lazyRegex := regexnew.New.LazyLock("^hello$")

	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsFailedMatch("hello"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: FirstMatchLine
// ==========================================================================

func Test_LazyRegex_FirstMatchLine_Found(t *testing.T) {
	tc := lazyRegexPatternMatchTestCases[3]
	lazyRegex := regexnew.New.LazyLock("(\\d+)")

	firstMatch, isInvalid := lazyRegex.FirstMatchLine("abc 123 def 456")

	tc.ShouldBeEqual(t, 0, firstMatch, fmt.Sprintf("%v", isInvalid))
}

func Test_LazyRegex_FirstMatchLine_NotFound(t *testing.T) {
	tc := lazyRegexPatternMatchTestCases[4]
	lazyRegex := regexnew.New.LazyLock("(\\d+)")

	firstMatch, isInvalid := lazyRegex.FirstMatchLine("no digits here")

	tc.ShouldBeEqual(t, 0, firstMatch, fmt.Sprintf("%v", isInvalid))
}
