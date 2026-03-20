package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
)

// Test_Cov12_MergeSlicesOfSlices_AllEmpty tests MergeSlicesOfSlices with all-empty slices.
func Test_Cov12_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	// Arrange / Act
	actual := stringslice.MergeSlicesOfSlices([]string{}, []string{})

	// Assert
	if len(actual) != 0 {
		t.Fatalf("MergeSlicesOfSlices with all empty: got len %d, want 0", len(actual))
	}
}

// Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult tests empty content with regex split.
func Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`\s+`)

	// Act
	actual := stringslice.RegexTrimmedSplitNonEmptyAll(re, "   ")

	// Assert
	if len(actual) != 0 {
		t.Fatalf("RegexTrimmedSplitNonEmptyAll with whitespace-only: got len %d, want 0", len(actual))
	}
}

// Test_Cov12_SplitTrimmedNonEmpty_ZeroCount tests SplitTrimmedNonEmpty with n=0.
func Test_Cov12_SplitTrimmedNonEmpty_ZeroCount(t *testing.T) {
	// Arrange / Act
	actual := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 0)

	// Assert
	if len(actual) != 0 {
		t.Fatalf("SplitTrimmedNonEmpty with n=0: got len %d, want 0", len(actual))
	}
}

// Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent tests with empty splitter result.
func Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent(t *testing.T) {
	// Arrange / Act
	actual := stringslice.SplitTrimmedNonEmptyAll("", "")

	// Assert
	if len(actual) != 0 {
		t.Fatalf("SplitTrimmedNonEmptyAll empty/empty: got len %d, want 0", len(actual))
	}
}
