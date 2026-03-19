package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_QW_Value_IsSet_WithNoNames(t *testing.T) {
	// The internal toHashset is called with the names param.
	// To cover the empty branch, call IsSet with no valid names
	v := issetter.Set
	result := v.IsSet()
	if !result {
		t.Fatal("expected true for Set value")
	}
}
