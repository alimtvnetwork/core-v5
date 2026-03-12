package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
)

// ── VersionSliceInteger extended ──

func Test_VersionSliceInteger_LenDiff_Cov2(t *testing.T) {
	if corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Error("shorter < longer")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1}) != corecomparator.LeftGreater {
		t.Error("longer > shorter")
	}
}

func Test_VersionSliceInteger_LeftNil_Cov2(t *testing.T) {
	if corecmp.VersionSliceInteger([]int{1}, nil) != corecomparator.NotEqual {
		t.Error("non-nil vs nil should be not equal")
	}
}

// ── VersionSliceByte left nil ──

func Test_VersionSliceByte_LeftNil_Cov2(t *testing.T) {
	if corecmp.VersionSliceByte([]byte{1}, nil) != corecomparator.NotEqual {
		t.Error("non-nil vs nil should be not equal")
	}
}

// ── IsStringsEqualPtr same len different content ──

func Test_IsStringsEqualPtr_SameLenDiffContent_Cov2(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"a", "c"}
	if corecmp.IsStringsEqualPtr(a, b) {
		t.Error("different content should not be equal")
	}
}

func Test_IsStringsEqualPtr_DiffLen_Cov2(t *testing.T) {
	a := []string{"a"}
	b := []string{"a", "b"}
	if corecmp.IsStringsEqualPtr(a, b) {
		t.Error("different lengths should not be equal")
	}
}

func Test_IsStringsEqualPtr_LeftNil_Cov2(t *testing.T) {
	b := []string{"a"}
	if corecmp.IsStringsEqualPtr(nil, b) {
		t.Error("nil vs non-nil should not be equal")
	}
}

// ── IsIntegersEqualPtr same len diff content ──

func Test_IsIntegersEqualPtr_SameLenDiff_Cov2(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 3}
	if corecmp.IsIntegersEqualPtr(&a, &b) {
		t.Error("different content should not be equal")
	}
}

func Test_IsIntegersEqualPtr_DiffLen_Cov2(t *testing.T) {
	a := []int{1}
	b := []int{1, 2}
	if corecmp.IsIntegersEqualPtr(&a, &b) {
		t.Error("different lengths should not be equal")
	}
}

func Test_IsIntegersEqualPtr_LeftNil_Cov2(t *testing.T) {
	b := []int{1}
	if corecmp.IsIntegersEqualPtr(nil, &b) {
		t.Error("nil vs non-nil should not be equal")
	}
}

// ── IsStringsEqualWithoutOrder diff len ──

func Test_IsStringsEqualWithoutOrder_DiffLen_Cov2(t *testing.T) {
	if corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}) {
		t.Error("different lengths should not be equal")
	}
}

func Test_IsStringsEqualWithoutOrder_LeftNil_Cov2(t *testing.T) {
	if corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"}) {
		t.Error("nil vs non-nil should not be equal")
	}
}

func Test_IsStringsEqualWithoutOrder_RightNil_Cov2(t *testing.T) {
	if corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil) {
		t.Error("non-nil vs nil should not be equal")
	}
}

// ── BytePtr left nil right nil ──

func Test_BytePtr_LeftNilRightNotNil_Cov2(t *testing.T) {
	b := byte(5)
	if corecmp.BytePtr(nil, &b) != corecomparator.NotEqual {
		t.Error("nil vs non-nil should be not equal")
	}
}

func Test_BytePtr_LeftNotNilRightNil_Cov2(t *testing.T) {
	a := byte(5)
	if corecmp.BytePtr(&a, nil) != corecomparator.NotEqual {
		t.Error("non-nil vs nil should be not equal")
	}
}

func Test_BytePtr_LeftLess_Cov2(t *testing.T) {
	a, b := byte(3), byte(5)
	if corecmp.BytePtr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── IntegerPtr extended ──

func Test_IntegerPtr_LeftNil_Cov2(t *testing.T) {
	b := 5
	if corecmp.IntegerPtr(nil, &b) != corecomparator.NotEqual {
		t.Error("nil vs non-nil")
	}
}

func Test_IntegerPtr_LeftLess_Cov2(t *testing.T) {
	a, b := 3, 5
	if corecmp.IntegerPtr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── TimePtr left nil ──

func Test_TimePtr_LeftNil_Cov2(t *testing.T) {
	if corecmp.TimePtr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be equal")
	}
}

// ── Integer16Ptr left nil ──

func Test_Integer16Ptr_LeftNil_Cov2(t *testing.T) {
	b := int16(5)
	if corecmp.Integer16Ptr(nil, &b) != corecomparator.NotEqual {
		t.Error("nil vs non-nil")
	}
}

func Test_Integer16Ptr_LeftLess_Cov2(t *testing.T) {
	a, b := int16(3), int16(5)
	if corecmp.Integer16Ptr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── Integer32Ptr left nil ──

func Test_Integer32Ptr_LeftNil_Cov2(t *testing.T) {
	b := int32(5)
	if corecmp.Integer32Ptr(nil, &b) != corecomparator.NotEqual {
		t.Error("nil vs non-nil")
	}
}

func Test_Integer32Ptr_LeftLess_Cov2(t *testing.T) {
	a, b := int32(3), int32(5)
	if corecmp.Integer32Ptr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── Integer64Ptr extended ──

func Test_Integer64Ptr_LeftLess_Cov2(t *testing.T) {
	a, b := int64(3), int64(5)
	if corecmp.Integer64Ptr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── Integer8Ptr extended ──

func Test_Integer8Ptr_LeftLess_Cov2(t *testing.T) {
	a, b := int8(3), int8(5)
	if corecmp.Integer8Ptr(&a, &b) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
}

// ── IsIntegersEqual left nil ──

func Test_IsIntegersEqual_LeftNil_Cov2(t *testing.T) {
	if corecmp.IsIntegersEqual(nil, []int{1}) {
		t.Error("nil vs non-nil should not be equal")
	}
}

func Test_IsIntegersEqual_RightNil_Cov2(t *testing.T) {
	if corecmp.IsIntegersEqual([]int{1}, nil) {
		t.Error("non-nil vs nil should not be equal")
	}
}

// ── IsStringsEqual left nil ──

func Test_IsStringsEqual_LeftNil_Cov2(t *testing.T) {
	if corecmp.IsStringsEqual(nil, []string{"a"}) {
		t.Error("nil vs non-nil should not be equal")
	}
}

func Test_IsStringsEqual_RightNil_Cov2(t *testing.T) {
	if corecmp.IsStringsEqual([]string{"a"}, nil) {
		t.Error("non-nil vs nil should not be equal")
	}
}
