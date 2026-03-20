package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

// Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice tests the length==0 branch
// with a non-nil empty slice.
func Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElements(false, input...)

	// Assert
	if len(actual) != 0 {
		t.Fatalf("DoubleQuoteWrapElements with empty slice: got len %d, want 0", len(actual))
	}
}

// Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice tests length==0 branch.
func Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElementsWithIndexes(input...)

	// Assert
	if len(actual) != 0 {
		t.Fatalf("DoubleQuoteWrapElementsWithIndexes with empty slice: got len %d, want 0", len(actual))
	}
}
