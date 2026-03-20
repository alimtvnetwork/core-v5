package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
	"github.com/alimtvnetwork/core/coretests"
)

// Test_Cov2_Write_PrimitiveType tests Write with a primitive (int) to hit
// the default fallthrough in WritePropertyOptions (line 62).
func Test_Cov2_Write_PrimitiveType(t *testing.T) {
	// Arrange
	// An int is not struct/slice/ptr/map, so it falls through the switch.

	// Act
	actual := coreproperty.Writer.Write(42)

	// Assert — should produce some string representation
	if actual == "" {
		t.Errorf("Write(42) should return non-empty string, got empty")
	}
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"Write(int) should produce a string representation",
		actual != "", true,
	)
}

// Test_Cov2_Write_NilPointer tests WritePointerRv with a nil pointer
// to hit the isany.Null branch (line 91-93).
func Test_Cov2_Write_NilPointer(t *testing.T) {
	// Arrange
	var nilPtr *int

	// Act
	actual := coreproperty.Writer.Write(nilPtr)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"Write(nil pointer) should return 'nil'",
		actual, "nil",
	)
}
