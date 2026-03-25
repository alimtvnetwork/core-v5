package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
)

func Test_Cov_Compare_IsCompareEqualLogically_LeftLessEqual(t *testing.T) {
	result := corecomparator.LeftLess.IsCompareEqualLogically(corecomparator.LeftLessEqual)
	if !result {
		t.Error("LeftLess should match LeftLessEqual logically")
	}
}

func Test_Cov_Compare_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Inconclusive compared with LeftGreater should return false
	result := corecomparator.Inconclusive.IsCompareEqualLogically(corecomparator.LeftGreater)
	if result {
		t.Error("expected false")
	}
}
