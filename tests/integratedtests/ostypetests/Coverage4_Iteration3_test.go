package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/ostype"
)

// Test_Cov4_Group_UnmarshallEnumToValue covers
// ostype/Group.go L191-193: UnmarshallEnumToValue.
func Test_Cov4_Group_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup
	validBytes := []byte("1")

	// Act
	val, err := g.UnmarshallEnumToValue(validBytes)

	// Assert
	coretests.ShouldBeNil(t, err)
	coretests.ShouldBeEqual(t, byte(ostype.UnixGroup), val)
}
