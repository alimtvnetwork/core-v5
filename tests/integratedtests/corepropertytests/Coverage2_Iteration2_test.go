package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
)

// Test_Cov2_Write_PrimitiveType tests Write with a primitive (int) to hit
// the default fallthrough in WritePropertyOptions (line 62).
func Test_Cov2_Write_PrimitiveType(t *testing.T) {
	// Arrange / Act
	actual := coreproperty.Writer.Write(42)

	// Assert
	if actual == "" {
		t.Fatalf("Write(42) should return non-empty string, got empty")
	}
}

// Test_Cov2_Write_NilPointer tests WritePointerRv with a nil pointer
// to hit the isany.Null branch (line 91-93).
func Test_Cov2_Write_NilPointer(t *testing.T) {
	// Arrange
	var nilPtr *int

	// Act
	actual := coreproperty.Writer.Write(nilPtr)

	// Assert
	if actual != "nil" {
		t.Fatalf("Write(nil pointer) should return 'nil', got %q", actual)
	}
}
