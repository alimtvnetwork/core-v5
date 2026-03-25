package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_WithStartEnd_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithStartEnd("[", "hello")

	// Assert
	if !strings.Contains(result, "[") {
		t.Error("should contain bracket wrapper")
	}
}

func Test_WithBracketsQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithBracketsQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_WithCurlyQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithCurlyQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_WithParenthesisQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithParenthesisQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleSquareCsvMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleSquareCsvMeta("title", "a", "b")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleQuotationMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_DoubleQuoteWrapElements_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2 elements, got %d", len(result))
	}
}
