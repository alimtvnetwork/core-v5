package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coredata/stringslice"
)

// Test_Cov12_MergeSlicesOfSlices_AllEmpty tests MergeSlicesOfSlices with all-empty slices.
func Test_Cov12_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	// Arrange
	expected := 0

	// Act
	actual := stringslice.MergeSlicesOfSlices([]string{}, []string{})

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"MergeSlicesOfSlices with all empty slices should return empty",
		len(actual),
		expected,
	)
}

// Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult tests empty content with regex split.
func Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`\s+`)

	// Act — empty string split by whitespace gives [""] which has len > 0,
	// so the empty branch (line 17-19) is hit only when regexp.Split returns empty.
	// Actually regexp.Split("", pattern) returns [""], not [].
	// The `len(items) == 0` branch may be dead code here too.
	// Let's try to trigger it anyway with a pattern that matches everything.
	actual := stringslice.RegexTrimmedSplitNonEmptyAll(re, "   ")

	// Assert — all whitespace trimmed to empty, TrimmedEachWords filters empties
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"RegexTrimmedSplitNonEmptyAll with whitespace-only returns trimmed result",
		len(actual),
		0,
	)
}

// Test_Cov12_SplitTrimmedNonEmpty_ZeroCount tests SplitTrimmedNonEmpty with n=0 (nil result).
func Test_Cov12_SplitTrimmedNonEmpty_ZeroCount(t *testing.T) {
	// Arrange
	// When n == 0, strings.SplitN returns nil (length 0), hitting the empty branch.

	// Act
	actual := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 0)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"SplitTrimmedNonEmpty with n=0 should return empty",
		len(actual),
		0,
	)
}

// Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent tests SplitTrimmedNonEmptyAll with empty splitter result.
func Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent(t *testing.T) {
	// Arrange
	// strings.Split("", ",") returns [""], not [], so len > 0.
	// The `len(items) == 0` branch might be dead code.
	// But splitting empty string by empty string: strings.Split("", "") returns [].
	// Actually no — strings.Split("", "") returns [].
	// Let's try that.

	// Act
	actual := stringslice.SplitTrimmedNonEmptyAll("", "")

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"SplitTrimmedNonEmptyAll empty content with empty splitter",
		len(actual),
		0,
	)
}
