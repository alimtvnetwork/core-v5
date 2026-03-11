package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
)

func Test_Byte_Direct_Comparison(t *testing.T) {
	if corecmp.Byte(5, 5) != corecomparator.Equal {
		t.Error("same bytes should be Equal")
	}
	if corecmp.Byte(3, 7) != corecomparator.LeftLess {
		t.Error("3 < 7 should be LeftLess")
	}
	if corecmp.Byte(7, 3) != corecomparator.LeftGreater {
		t.Error("7 > 3 should be LeftGreater")
	}
}

func Test_BytePtr_Direct_Comparison(t *testing.T) {
	if corecmp.BytePtr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	b := byte(5)
	if corecmp.BytePtr(nil, &b) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
	if corecmp.BytePtr(&b, nil) != corecomparator.NotEqual {
		t.Error("right nil should be NotEqual")
	}
	b2 := byte(5)
	if corecmp.BytePtr(&b, &b2) != corecomparator.Equal {
		t.Error("same values should be Equal")
	}
}

func Test_Integer64_Direct_Comparison(t *testing.T) {
	if corecmp.Integer64(10, 10) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.Integer64(5, 10) != corecomparator.LeftLess {
		t.Error("5 < 10 should be LeftLess")
	}
	if corecmp.Integer64(10, 5) != corecomparator.LeftGreater {
		t.Error("10 > 5 should be LeftGreater")
	}
}

func Test_IntegerPtr_Direct_Comparison(t *testing.T) {
	if corecmp.IntegerPtr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	val := 5
	if corecmp.IntegerPtr(nil, &val) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
	val2 := 5
	if corecmp.IntegerPtr(&val, &val2) != corecomparator.Equal {
		t.Error("same values should be Equal")
	}
}

func Test_IsIntegersEqual_Direct_Verification(t *testing.T) {
	if !corecmp.IsIntegersEqual(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsIntegersEqual(nil, []int{1}) {
		t.Error("nil vs non-nil should not be equal")
	}
	if !corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2}) {
		t.Error("same slices should be equal")
	}
	if corecmp.IsIntegersEqual([]int{1}, []int{2}) {
		t.Error("different values should not be equal")
	}
}

func Test_VersionSliceByte_Direct_Comparison(t *testing.T) {
	if corecmp.VersionSliceByte(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	if corecmp.VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual {
		t.Error("nil vs non-nil should be NotEqual")
	}
	if corecmp.VersionSliceByte([]byte{1}, nil) != corecomparator.NotEqual {
		t.Error("non-nil vs nil should be NotEqual")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess {
		t.Error("1.2 < 1.3 should be LeftLess")
	}
	if corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater {
		t.Error("1.3 > 1.2 should be LeftGreater")
	}
	if corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Error("shorter left should be LeftLess")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater {
		t.Error("longer left should be LeftGreater")
	}
}

func Test_IsStringsEqualWithoutOrder_Direct_Verification(t *testing.T) {
	if !corecmp.IsStringsEqualWithoutOrder(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"}) {
		t.Error("nil vs non-nil should not be equal")
	}
	if !corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"}) {
		t.Error("same items different order should be equal")
	}
	if corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"b"}) {
		t.Error("different items should not be equal")
	}
	if corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}) {
		t.Error("different lengths should not be equal")
	}
}

func Test_Integer8_Direct_Comparison(t *testing.T) {
	if corecmp.Integer8(5, 5) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.Integer8(3, 7) != corecomparator.LeftLess {
		t.Error("3 < 7 should be LeftLess")
	}
	if corecmp.Integer8(7, 3) != corecomparator.LeftGreater {
		t.Error("7 > 3 should be LeftGreater")
	}
}

func Test_Integer16_Direct_Comparison(t *testing.T) {
	if corecmp.Integer16(5, 5) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.Integer16(3, 7) != corecomparator.LeftLess {
		t.Error("3 < 7 should be LeftLess")
	}
}

func Test_Integer32_Direct_Comparison(t *testing.T) {
	if corecmp.Integer32(5, 5) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.Integer32(3, 7) != corecomparator.LeftLess {
		t.Error("3 < 7 should be LeftLess")
	}
}

func Test_Integer8Ptr_Direct_Comparison(t *testing.T) {
	if corecmp.Integer8Ptr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	v := int8(5)
	if corecmp.Integer8Ptr(nil, &v) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
}

func Test_Integer16Ptr_Direct_Comparison(t *testing.T) {
	if corecmp.Integer16Ptr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	v := int16(5)
	if corecmp.Integer16Ptr(nil, &v) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
}

func Test_Integer32Ptr_Direct_Comparison(t *testing.T) {
	if corecmp.Integer32Ptr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	v := int32(5)
	if corecmp.Integer32Ptr(nil, &v) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
}

func Test_Integer64Ptr_Direct_Comparison(t *testing.T) {
	if corecmp.Integer64Ptr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	v := int64(5)
	if corecmp.Integer64Ptr(nil, &v) != corecomparator.NotEqual {
		t.Error("left nil should be NotEqual")
	}
}

func Test_IsIntegersEqualPtr_Verification(t *testing.T) {
	if !corecmp.IsIntegersEqualPtr(nil, nil) {
		t.Error("both nil should be equal")
	}
	a := []int{1, 2}
	if corecmp.IsIntegersEqualPtr(&a, nil) {
		t.Error("one nil should not be equal")
	}
	b := []int{1, 2}
	if !corecmp.IsIntegersEqualPtr(&a, &b) {
		t.Error("same slices should be equal")
	}
}

func Test_AnyItem_Verification(t *testing.T) {
	result := corecmp.AnyItem("hello", "hello")
	if result != corecomparator.Equal {
		t.Error("same strings should be Equal")
	}
	result2 := corecmp.AnyItem(1, 2)
	if result2 == corecomparator.Equal {
		t.Error("different ints should not be Equal")
	}
}

func Test_VersionSliceInteger_Direct_Comparison(t *testing.T) {
	if corecmp.VersionSliceInteger(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be Equal")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal {
		t.Error("same should be Equal")
	}
	if corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Error("shorter should be LeftLess")
	}
}
