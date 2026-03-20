package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests"
)

// Test_Cov6_ToStringsUsingProcessor_EmptyAnyItems covers
// converters/anyItemConverter.go L147-149: when ToAnyItems returns empty slice.
func Test_Cov6_ToStringsUsingProcessor_EmptyAnyItems(t *testing.T) {
	// Arrange
	processor := func(index int, in any) (string, bool, bool) {
		return "", false, false
	}

	// Act
	// Pass a non-nil but empty slice so anyVal != nil but ToAnyItems returns empty
	result := converters.AnyTo.ToStringsUsingProcessor(
		true,
		processor,
		[]int{},
	)

	// Assert
	coretests.ShouldBeEqual(t, 0, len(result))
}
