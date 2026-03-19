package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
)

func Test_QW_IsOutOfRange_ToUnsignedInt32(t *testing.T) {
	// ToUnsignedInt32 takes an int argument
	result := coremath.IsOutOfRange.ToUnsignedInt32(-1)
	if !result {
		t.Fatal("expected true for negative value")
	}
	result2 := coremath.IsOutOfRange.ToUnsignedInt32(100)
	if result2 {
		t.Fatal("expected false for valid value")
	}
}
