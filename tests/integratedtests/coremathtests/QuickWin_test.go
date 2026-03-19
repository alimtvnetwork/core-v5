package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
)

func Test_QW_IsOutOfRange_Integer_ToUnsignedInt32(t *testing.T) {
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(-1)
	if !result {
		t.Fatal("expected true for negative value")
	}
	result2 := coremath.IsOutOfRange.Integer.ToUnsignedInt32(100)
	if result2 {
		t.Fatal("expected false for valid value")
	}
}
