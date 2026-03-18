package corecmp

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
)

// ── AnyItem ──

func TestAnyItem_BothNil(t *testing.T) {
	if AnyItem(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestAnyItem_LeftNil(t *testing.T) {
	if AnyItem(nil, "a") != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestAnyItem_RightNil(t *testing.T) {
	if AnyItem("a", nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestAnyItem_Equal(t *testing.T) {
	if AnyItem("a", "a") != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestAnyItem_Inconclusive(t *testing.T) {
	if AnyItem("a", "b") != corecomparator.Inconclusive {
		t.Fatal("expected Inconclusive")
	}
}

// ── Byte ──

func TestByte_Equal(t *testing.T) {
	if Byte(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestByte_LeftLess(t *testing.T) {
	if Byte(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestByte_LeftGreater(t *testing.T) {
	if Byte(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── BytePtr ──

func TestBytePtr_BothNil(t *testing.T) {
	if BytePtr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestBytePtr_LeftNil(t *testing.T) {
	b := byte(1)
	if BytePtr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestBytePtr_RightNil(t *testing.T) {
	b := byte(1)
	if BytePtr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestBytePtr_Equal(t *testing.T) {
	a, b := byte(5), byte(5)
	if BytePtr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── Integer ──

func TestInteger_Equal(t *testing.T) {
	if Integer(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger_LeftLess(t *testing.T) {
	if Integer(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestInteger_LeftGreater(t *testing.T) {
	if Integer(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── IntegerPtr ──

func TestIntegerPtr_BothNil(t *testing.T) {
	if IntegerPtr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestIntegerPtr_LeftNil(t *testing.T) {
	b := 1
	if IntegerPtr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestIntegerPtr_RightNil(t *testing.T) {
	b := 1
	if IntegerPtr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestIntegerPtr_Equal(t *testing.T) {
	a, b := 5, 5
	if IntegerPtr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── Integer8 ──

func TestInteger8_Equal(t *testing.T) {
	if Integer8(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger8_LeftLess(t *testing.T) {
	if Integer8(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestInteger8_LeftGreater(t *testing.T) {
	if Integer8(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── Integer8Ptr ──

func TestInteger8Ptr_BothNil(t *testing.T) {
	if Integer8Ptr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger8Ptr_LeftNil(t *testing.T) {
	b := int8(1)
	if Integer8Ptr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger8Ptr_RightNil(t *testing.T) {
	b := int8(1)
	if Integer8Ptr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger8Ptr_Equal(t *testing.T) {
	a, b := int8(5), int8(5)
	if Integer8Ptr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── Integer16 ──

func TestInteger16_Equal(t *testing.T) {
	if Integer16(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger16_LeftLess(t *testing.T) {
	if Integer16(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestInteger16_LeftGreater(t *testing.T) {
	if Integer16(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── Integer16Ptr ──

func TestInteger16Ptr_BothNil(t *testing.T) {
	if Integer16Ptr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger16Ptr_LeftNil(t *testing.T) {
	b := int16(1)
	if Integer16Ptr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger16Ptr_RightNil(t *testing.T) {
	b := int16(1)
	if Integer16Ptr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger16Ptr_Equal(t *testing.T) {
	a, b := int16(5), int16(5)
	if Integer16Ptr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── Integer32 ──

func TestInteger32_Equal(t *testing.T) {
	if Integer32(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger32_LeftLess(t *testing.T) {
	if Integer32(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestInteger32_LeftGreater(t *testing.T) {
	if Integer32(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── Integer32Ptr ──

func TestInteger32Ptr_BothNil(t *testing.T) {
	if Integer32Ptr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger32Ptr_LeftNil(t *testing.T) {
	b := int32(1)
	if Integer32Ptr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger32Ptr_RightNil(t *testing.T) {
	b := int32(1)
	if Integer32Ptr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger32Ptr_Equal(t *testing.T) {
	a, b := int32(5), int32(5)
	if Integer32Ptr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── Integer64 ──

func TestInteger64_Equal(t *testing.T) {
	if Integer64(1, 1) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger64_LeftLess(t *testing.T) {
	if Integer64(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestInteger64_LeftGreater(t *testing.T) {
	if Integer64(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── Integer64Ptr ──

func TestInteger64Ptr_BothNil(t *testing.T) {
	if Integer64Ptr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestInteger64Ptr_LeftNil(t *testing.T) {
	b := int64(1)
	if Integer64Ptr(nil, &b) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger64Ptr_RightNil(t *testing.T) {
	b := int64(1)
	if Integer64Ptr(&b, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestInteger64Ptr_Equal(t *testing.T) {
	a, b := int64(5), int64(5)
	if Integer64Ptr(&a, &b) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── IsStringsEqual ──

func TestIsStringsEqual_BothNil(t *testing.T) {
	if !IsStringsEqual(nil, nil) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqual_LeftNil(t *testing.T) {
	if IsStringsEqual(nil, []string{"a"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqual_RightNil(t *testing.T) {
	if IsStringsEqual([]string{"a"}, nil) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqual_DiffLen(t *testing.T) {
	if IsStringsEqual([]string{"a"}, []string{"a", "b"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqual_Equal(t *testing.T) {
	if !IsStringsEqual([]string{"a", "b"}, []string{"a", "b"}) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqual_NotEqual(t *testing.T) {
	if IsStringsEqual([]string{"a"}, []string{"b"}) {
		t.Fatal("expected false")
	}
}

// ── IsStringsEqualPtr ──

func TestIsStringsEqualPtr_BothNil(t *testing.T) {
	if !IsStringsEqualPtr(nil, nil) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqualPtr_LeftNil(t *testing.T) {
	if IsStringsEqualPtr(nil, []string{"a"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqualPtr_RightNil(t *testing.T) {
	if IsStringsEqualPtr([]string{"a"}, nil) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqualPtr_DiffLen(t *testing.T) {
	if IsStringsEqualPtr([]string{"a"}, []string{"a", "b"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqualPtr_Equal(t *testing.T) {
	if !IsStringsEqualPtr([]string{"a"}, []string{"a"}) {
		t.Fatal("expected true")
	}
}

// ── IsStringsEqualWithoutOrder ──

func TestIsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	if !IsStringsEqualWithoutOrder(nil, nil) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqualWithoutOrder_LeftNil(t *testing.T) {
	if IsStringsEqualWithoutOrder(nil, []string{"a"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqualWithoutOrder_DiffLen(t *testing.T) {
	if IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}) {
		t.Fatal("expected false")
	}
}

func TestIsStringsEqualWithoutOrder_SameOrder(t *testing.T) {
	if !IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "b"}) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqualWithoutOrder_DiffOrder(t *testing.T) {
	if !IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"}) {
		t.Fatal("expected true")
	}
}

func TestIsStringsEqualWithoutOrder_NotEqual(t *testing.T) {
	if IsStringsEqualWithoutOrder([]string{"a"}, []string{"b"}) {
		t.Fatal("expected false")
	}
}

// ── IsIntegersEqual ──

func TestIsIntegersEqual_BothNil(t *testing.T) {
	if !IsIntegersEqual(nil, nil) {
		t.Fatal("expected true")
	}
}

func TestIsIntegersEqual_LeftNil(t *testing.T) {
	if IsIntegersEqual(nil, []int{1}) {
		t.Fatal("expected false")
	}
}

func TestIsIntegersEqual_RightNil(t *testing.T) {
	if IsIntegersEqual([]int{1}, nil) {
		t.Fatal("expected false")
	}
}

func TestIsIntegersEqual_Equal(t *testing.T) {
	if !IsIntegersEqual([]int{1, 2}, []int{1, 2}) {
		t.Fatal("expected true")
	}
}

// ── IsIntegersEqualPtr ──

func TestIsIntegersEqualPtr_BothNil(t *testing.T) {
	if !IsIntegersEqualPtr(nil, nil) {
		t.Fatal("expected true")
	}
}

func TestIsIntegersEqualPtr_LeftNil(t *testing.T) {
	b := []int{1}
	if IsIntegersEqualPtr(nil, &b) {
		t.Fatal("expected false")
	}
}

func TestIsIntegersEqualPtr_RightNil(t *testing.T) {
	a := []int{1}
	if IsIntegersEqualPtr(&a, nil) {
		t.Fatal("expected false")
	}
}

func TestIsIntegersEqualPtr_DiffLen(t *testing.T) {
	a := []int{1}
	b := []int{1, 2}
	if IsIntegersEqualPtr(&a, &b) {
		t.Fatal("expected false")
	}
}

func TestIsIntegersEqualPtr_Equal(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 2}
	if !IsIntegersEqualPtr(&a, &b) {
		t.Fatal("expected true")
	}
}

// ── Time ──

func TestTime_Equal(t *testing.T) {
	now := time.Now()
	if Time(now, now) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestTime_LeftLess(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	if Time(now, later) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestTime_LeftGreater(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	if Time(now, earlier) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── TimePtr ──

func TestTimePtr_BothNil(t *testing.T) {
	if TimePtr(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestTimePtr_LeftNil(t *testing.T) {
	now := time.Now()
	if TimePtr(nil, &now) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestTimePtr_RightNil(t *testing.T) {
	now := time.Now()
	if TimePtr(&now, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestTimePtr_Equal(t *testing.T) {
	now := time.Now()
	now2 := now
	if TimePtr(&now, &now2) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

// ── VersionSliceByte ──

func TestVersionSliceByte_BothNil(t *testing.T) {
	if VersionSliceByte(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestVersionSliceByte_LeftNil(t *testing.T) {
	if VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestVersionSliceByte_RightNil(t *testing.T) {
	if VersionSliceByte([]byte{1}, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestVersionSliceByte_Equal(t *testing.T) {
	if VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestVersionSliceByte_LeftLess_SameLen(t *testing.T) {
	if VersionSliceByte([]byte{1, 1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestVersionSliceByte_LeftGreater_SameLen(t *testing.T) {
	if VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

func TestVersionSliceByte_LeftLess_ShorterLen(t *testing.T) {
	if VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestVersionSliceByte_LeftGreater_LongerLen(t *testing.T) {
	if VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

// ── VersionSliceInteger ──

func TestVersionSliceInteger_BothNil(t *testing.T) {
	if VersionSliceInteger(nil, nil) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestVersionSliceInteger_LeftNil(t *testing.T) {
	if VersionSliceInteger(nil, []int{1}) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestVersionSliceInteger_RightNil(t *testing.T) {
	if VersionSliceInteger([]int{1}, nil) != corecomparator.NotEqual {
		t.Fatal("expected NotEqual")
	}
}

func TestVersionSliceInteger_Equal(t *testing.T) {
	if VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal {
		t.Fatal("expected Equal")
	}
}

func TestVersionSliceInteger_LeftLess_SameLen(t *testing.T) {
	if VersionSliceInteger([]int{1, 1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestVersionSliceInteger_LeftGreater_SameLen(t *testing.T) {
	if VersionSliceInteger([]int{1, 3}, []int{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}

func TestVersionSliceInteger_LeftLess_ShorterLen(t *testing.T) {
	if VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected LeftLess")
	}
}

func TestVersionSliceInteger_LeftGreater_LongerLen(t *testing.T) {
	if VersionSliceInteger([]int{1, 2}, []int{1}) != corecomparator.LeftGreater {
		t.Fatal("expected LeftGreater")
	}
}
