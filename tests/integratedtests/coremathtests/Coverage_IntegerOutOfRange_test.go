package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
)

func Test_Cov_IntegerOutOfRange_ToInt(t *testing.T) {
	if coremath.IntegerOutOfRange.ToInt(0) {
		t.Error("0 should be in range")
	}
}
