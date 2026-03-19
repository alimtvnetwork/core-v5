package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
)

func Test_QW_IsExpectedCompareResult_Fallthrough(t *testing.T) {
	// Cover the final `return false` branch
	// We need a Compare value that doesn't match Equal, NotEqual,
	// LeftGreater, LeftGreaterEqual, or LeftLessEqual
	// Inconclusive should fall through all checks
	c := corecomparator.Equal
	result := c.IsExpectedCompareResult(corecomparator.Inconclusive)
	if result {
		t.Fatal("expected false for inconclusive")
	}
}
