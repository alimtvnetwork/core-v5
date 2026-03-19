package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_QW_DoubleQuoteWrapElements_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)
	if len(result) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_QW_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)
	if len(result) != 0 {
		t.Fatal("expected empty for nil")
	}
}
