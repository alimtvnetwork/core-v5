package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_Cov_DoubleQuoteWrapElements_SkipOnExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, "hello", "\"already\"")
	if len(result) != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_DoubleQuoteWrapElements_NoSkip(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, "hello")
	if len(result) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_DoubleQuoteWrapElements_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false)
	if len(result) != 0 {
		t.Error("expected 0")
	}
}

func Test_Cov_DoubleQuoteWrapElementsWithIndexes(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")
	if len(result) != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	if len(result) != 0 {
		t.Error("expected 0")
	}
}
