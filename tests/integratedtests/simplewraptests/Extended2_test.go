package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_WrapWithStartEnd_Verification(t *testing.T) {
	// Act
	result := simplewrap.WrapWithStartEnd('[', "hello", ']')

	// Assert
	if !strings.Contains(result, "[") || !strings.Contains(result, "]") {
		t.Error("should contain brackets")
	}
}

func Test_WithBracketsQuotation_Verification(t *testing.T) {
	// Act
	result := simplewrap.WithBracketsQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_WithCurlyQuotation_Verification(t *testing.T) {
	// Act
	result := simplewrap.WithCurlyQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_WithParenthesisQuotation_Verification(t *testing.T) {
	// Act
	result := simplewrap.WithParenthesisQuotation("hello")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleSquareCsvMeta_Verification(t *testing.T) {
	// Act
	result := simplewrap.TitleSquareCsvMeta("title", "a", "b")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleSquareMetaUsingFmt_Verification(t *testing.T) {
	// Act
	result := simplewrap.TitleSquareMetaUsingFmt("title", "value", "meta %s", "formatted")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleQuotationMeta_Verification(t *testing.T) {
	// Act
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_DoubleQuoteWrapElements_Verification(t *testing.T) {
	// Act
	result := simplewrap.DoubleQuoteWrapElements([]string{"a", "b"})

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2 elements, got %d", len(result))
	}
}

func Test_WrapDoubleQuoteOnNonExist_Verification(t *testing.T) {
	// Act - already quoted
	result1 := simplewrap.WrapDoubleQuoteOnNonExist("\"hello\"")
	if !strings.HasPrefix(result1, "\"") {
		t.Error("already quoted should stay quoted")
	}

	// Act - not quoted
	result2 := simplewrap.WrapDoubleQuoteOnNonExist("hello")
	if !strings.Contains(result2, "\"") {
		t.Error("unquoted should get quotes")
	}
}
