package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
)

// ── Byte ──

func Test_Byte_Coverage(t *testing.T) {
	if corecmp.Byte(5, 5) != corecomparator.Equal {
		t.Error("5 == 5")
	}
	if corecmp.Byte(3, 5) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
	if corecmp.Byte(5, 3) != corecomparator.LeftGreater {
		t.Error("5 > 3")
	}
}

func Test_BytePtr_Coverage(t *testing.T) {
	a, b := byte(5), byte(3)
	if corecmp.BytePtr(&a, &a) != corecomparator.Equal {
		t.Error("same should be equal")
	}
	if corecmp.BytePtr(&a, &b) != corecomparator.LeftGreater {
		t.Error("5 > 3")
	}
	if corecmp.BytePtr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be equal")
	}
	if corecmp.BytePtr(nil, &a) != corecomparator.NotEqual {
		t.Error("one nil should be not equal")
	}
	if corecmp.BytePtr(&a, nil) != corecomparator.NotEqual {
		t.Error("one nil should be not equal")
	}
}

// ── Integer ──

func Test_Integer_Coverage(t *testing.T) {
	if corecmp.Integer(5, 5) != corecomparator.Equal {
		t.Error("5 == 5")
	}
	if corecmp.Integer(3, 5) != corecomparator.LeftLess {
		t.Error("3 < 5")
	}
	if corecmp.Integer(5, 3) != corecomparator.LeftGreater {
		t.Error("5 > 3")
	}
}

func Test_IntegerPtr_Coverage(t *testing.T) {
	a, b := 5, 3
	if corecmp.IntegerPtr(&a, &a) != corecomparator.Equal {
		t.Error("same should be equal")
	}
	if corecmp.IntegerPtr(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be equal")
	}
	if corecmp.IntegerPtr(nil, &a) != corecomparator.NotEqual {
		t.Error("one nil should be not equal")
	}
	if corecmp.IntegerPtr(&a, &b) != corecomparator.LeftGreater {
		t.Error("5 > 3")
	}
}

// ── Integer64 ──

func Test_Integer64_Coverage(t *testing.T) {
	if corecmp.Integer64(5, 5) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer64(3, 5) != corecomparator.LeftLess {
		t.Error("less")
	}
	if corecmp.Integer64(5, 3) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

// ── Integer16, Integer32, Integer8 ──

func Test_Integer16_Coverage(t *testing.T) {
	if corecmp.Integer16(5, 5) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer16(3, 5) != corecomparator.LeftLess {
		t.Error("less")
	}
	if corecmp.Integer16(5, 3) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

func Test_Integer32_Coverage(t *testing.T) {
	if corecmp.Integer32(5, 5) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer32(3, 5) != corecomparator.LeftLess {
		t.Error("less")
	}
	if corecmp.Integer32(5, 3) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

func Test_Integer8_Coverage(t *testing.T) {
	if corecmp.Integer8(5, 5) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer8(3, 5) != corecomparator.LeftLess {
		t.Error("less")
	}
	if corecmp.Integer8(5, 3) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

// ── Ptr variants for 16/32/64/8 ──

func Test_Integer16Ptr_Coverage(t *testing.T) {
	a, b := int16(5), int16(3)
	if corecmp.Integer16Ptr(&a, &a) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer16Ptr(nil, nil) != corecomparator.Equal {
		t.Error("nil equal")
	}
	if corecmp.Integer16Ptr(nil, &a) != corecomparator.NotEqual {
		t.Error("nil not equal")
	}
	if corecmp.Integer16Ptr(&a, &b) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

func Test_Integer32Ptr_Coverage(t *testing.T) {
	a, b := int32(5), int32(3)
	if corecmp.Integer32Ptr(&a, &a) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer32Ptr(nil, nil) != corecomparator.Equal {
		t.Error("nil equal")
	}
	if corecmp.Integer32Ptr(nil, &a) != corecomparator.NotEqual {
		t.Error("nil not equal")
	}
	if corecmp.Integer32Ptr(&a, &b) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

func Test_Integer64Ptr_Coverage(t *testing.T) {
	a, b := int64(5), int64(3)
	if corecmp.Integer64Ptr(&a, &a) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer64Ptr(nil, nil) != corecomparator.Equal {
		t.Error("nil equal")
	}
	if corecmp.Integer64Ptr(nil, &a) != corecomparator.NotEqual {
		t.Error("nil not equal")
	}
	if corecmp.Integer64Ptr(&a, &b) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

func Test_Integer8Ptr_Coverage(t *testing.T) {
	a, b := int8(5), int8(3)
	if corecmp.Integer8Ptr(&a, &a) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.Integer8Ptr(nil, nil) != corecomparator.Equal {
		t.Error("nil equal")
	}
	if corecmp.Integer8Ptr(nil, &a) != corecomparator.NotEqual {
		t.Error("nil not equal")
	}
	if corecmp.Integer8Ptr(&a, &b) != corecomparator.LeftGreater {
		t.Error("greater")
	}
}

// ── Time / TimePtr ──

func Test_Time_Coverage(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	later := now.Add(time.Hour)

	if corecmp.Time(now, now) != corecomparator.Equal {
		t.Error("same time should be equal")
	}
	if corecmp.Time(earlier, later) != corecomparator.LeftLess {
		t.Error("earlier < later")
	}
	if corecmp.Time(later, earlier) != corecomparator.LeftGreater {
		t.Error("later > earlier")
	}
}

func Test_TimePtr_Coverage(t *testing.T) {
	now := time.Now()
	if corecmp.TimePtr(&now, &now) != corecomparator.Equal {
		t.Error("equal")
	}
	if corecmp.TimePtr(nil, nil) != corecomparator.Equal {
		t.Error("nil equal")
	}
	if corecmp.TimePtr(nil, &now) != corecomparator.NotEqual {
		t.Error("nil not equal")
	}
}

// ── IsStringsEqual / IsStringsEqualPtr / IsStringsEqualWithoutOrder ──

func Test_IsStringsEqual_Coverage(t *testing.T) {
	if !corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"}) {
		t.Error("same should be equal")
	}
	if corecmp.IsStringsEqual([]string{"a"}, []string{"b"}) {
		t.Error("different should not be equal")
	}
	if corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"}) {
		t.Error("different length should not be equal")
	}
	if !corecmp.IsStringsEqual(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsStringsEqual(nil, []string{}) {
		t.Error("nil vs empty should not be equal")
	}
}

func Test_IsStringsEqualPtr_Coverage(t *testing.T) {
	if !corecmp.IsStringsEqualPtr(nil, nil) {
		t.Error("both nil should be equal")
	}
	a := []string{"a"}
	if corecmp.IsStringsEqualPtr(nil, a) {
		t.Error("nil vs non-nil should not be equal")
	}
	if !corecmp.IsStringsEqualPtr(a, a) {
		t.Error("same should be equal")
	}
}

func Test_IsStringsEqualWithoutOrder_Coverage(t *testing.T) {
	if !corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"}) {
		t.Error("same items different order should be equal")
	}
	if corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"b"}) {
		t.Error("different items should not be equal")
	}
	if !corecmp.IsStringsEqualWithoutOrder(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsStringsEqualWithoutOrder(nil, []string{}) {
		t.Error("nil vs empty should not be equal")
	}
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_IsIntegersEqual_Coverage(t *testing.T) {
	if !corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2}) {
		t.Error("same should be equal")
	}
	if corecmp.IsIntegersEqual([]int{1}, []int{2}) {
		t.Error("different should not be equal")
	}
	if !corecmp.IsIntegersEqual(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsIntegersEqual(nil, []int{}) {
		t.Error("nil vs empty should not be equal")
	}
}

func Test_IsIntegersEqualPtr_Coverage(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 2}
	c := []int{3}

	if !corecmp.IsIntegersEqualPtr(&a, &b) {
		t.Error("same should be equal")
	}
	if corecmp.IsIntegersEqualPtr(&a, &c) {
		t.Error("different should not be equal")
	}
	if !corecmp.IsIntegersEqualPtr(nil, nil) {
		t.Error("both nil should be equal")
	}
	if corecmp.IsIntegersEqualPtr(nil, &a) {
		t.Error("nil vs non-nil should not be equal")
	}
}

// ── AnyItem ──

func Test_AnyItem_Coverage(t *testing.T) {
	if corecmp.AnyItem(nil, nil) != corecomparator.Equal {
		t.Error("both nil should be equal")
	}
	if corecmp.AnyItem(nil, 42) != corecomparator.NotEqual {
		t.Error("nil vs non-nil should be not equal")
	}
	if corecmp.AnyItem(42, nil) != corecomparator.NotEqual {
		t.Error("non-nil vs nil should be not equal")
	}
	if corecmp.AnyItem(42, 42) != corecomparator.Equal {
		t.Error("same should be equal")
	}
	if corecmp.AnyItem(42, 43) != corecomparator.Inconclusive {
		t.Error("different should be inconclusive")
	}
}

// ── VersionSliceByte ──

func Test_VersionSliceByte_Coverage(t *testing.T) {
	if corecmp.VersionSliceByte(nil, nil) != corecomparator.Equal {
		t.Error("both nil equal")
	}
	if corecmp.VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual {
		t.Error("nil vs non-nil not equal")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal {
		t.Error("same should be equal")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess {
		t.Error("1.2 < 1.3")
	}
	if corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater {
		t.Error("1.3 > 1.2")
	}
	if corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Error("shorter version less")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater {
		t.Error("longer version greater")
	}
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_Coverage(t *testing.T) {
	if corecmp.VersionSliceInteger(nil, nil) != corecomparator.Equal {
		t.Error("both nil equal")
	}
	if corecmp.VersionSliceInteger(nil, []int{1}) != corecomparator.NotEqual {
		t.Error("nil vs non-nil not equal")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal {
		t.Error("same should be equal")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) != corecomparator.LeftLess {
		t.Error("1.2 < 1.3")
	}
	if corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) != corecomparator.LeftGreater {
		t.Error("1.3 > 1.2")
	}
}
