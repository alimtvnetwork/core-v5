package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
)

func Test_QW_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Cover the final `return false` branch in IsCompareEqualLogically
	c := corecomparator.Equal
	result := c.IsCompareEqualLogically(corecomparator.Inconclusive)
	if result {
		t.Fatal("expected false for inconclusive")
	}
}
