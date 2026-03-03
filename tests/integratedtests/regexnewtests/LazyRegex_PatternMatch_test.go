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
	tc := lazyRegexIsMatchFullDigitTestCase

	// Arrange
	compareInput, _ := tc.GetExtraAsString("compareInput")
	lazyRegex := regexnew.New.LazyLock(tc.Input)

	// Act
	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_LazyRegex_IsMatch_PartialMismatch(t *testing.T) {
	tc := lazyRegexIsMatchPartialMismatchTestCase

	// Arrange
	compareInput, _ := tc.GetExtraAsString("compareInput")
	lazyRegex := regexnew.New.LazyLock(tc.Input)

	// Act
	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsMatch(compareInput))}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

// ==========================================================================
// Test: IsFailedMatch
// ==========================================================================

func Test_LazyRegex_IsFailedMatch(t *testing.T) {
	tc := lazyRegexIsFailedMatchTestCase

	// Arrange
	compareInput, _ := tc.GetExtraAsString("compareInput")
	lazyRegex := regexnew.New.LazyLock(tc.Input)

	// Act
	actLines := []string{fmt.Sprintf("%v", lazyRegex.IsFailedMatch(compareInput))}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

// ==========================================================================
// Test: FirstMatchLine
// ==========================================================================

func Test_LazyRegex_FirstMatchLine_Found(t *testing.T) {
	tc := lazyRegexFirstMatchLineFoundTestCase

	// Arrange
	compareInput, _ := tc.GetExtraAsString("compareInput")
	lazyRegex := regexnew.New.LazyLock(tc.Input)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)
	actLines := []string{firstMatch, fmt.Sprintf("%v", isInvalid)}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_LazyRegex_FirstMatchLine_NotFound(t *testing.T) {
	tc := lazyRegexFirstMatchLineNotFoundTestCase

	// Arrange
	compareInput, _ := tc.GetExtraAsString("compareInput")
	lazyRegex := regexnew.New.LazyLock(tc.Input)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)
	actLines := []string{firstMatch, fmt.Sprintf("%v", isInvalid)}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}
