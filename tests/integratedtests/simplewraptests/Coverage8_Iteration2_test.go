package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/simplewrap"
)

// Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice tests the length==0 branch
// with a non-nil empty slice (distinct from nil).
func Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElements(false, input...)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"DoubleQuoteWrapElements with empty non-nil slice should return empty",
		len(actual), 0,
	)
}

// Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice tests length==0 branch.
func Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElementsWithIndexes(input...)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"DoubleQuoteWrapElementsWithIndexes with empty non-nil slice should return empty",
		len(actual), 0,
	)
}
