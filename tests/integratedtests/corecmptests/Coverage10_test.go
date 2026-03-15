package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Integer — Greater branch
// ═══════════════════════════════════════════

func Test_Cov10_Integer_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer greater", actual)
}

// ═══════════════════════════════════════════
// Integer8 — Greater branch
// ═══════════════════════════════════════════

func Test_Cov10_Integer8_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 greater", actual)
}

func Test_Cov10_Integer8Ptr_LeftNil(t *testing.T) {
	r := int8(5)
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr left nil", actual)
}

func Test_Cov10_Integer8Ptr_Equal(t *testing.T) {
	l, r := int8(5), int8(5)
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr equal", actual)
}

// ═══════════════════════════════════════════
// Integer16 — all branches
// ═══════════════════════════════════════════

func Test_Cov10_Integer16_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 equal", actual)
}

func Test_Cov10_Integer16_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 less", actual)
}

func Test_Cov10_Integer16_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 greater", actual)
}

func Test_Cov10_Integer16Ptr_LeftNil(t *testing.T) {
	r := int16(5)
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr left nil", actual)
}

func Test_Cov10_Integer16Ptr_Equal(t *testing.T) {
	l, r := int16(5), int16(5)
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr equal", actual)
}

// ═══════════════════════════════════════════
// Integer32 — all branches
// ═══════════════════════════════════════════

func Test_Cov10_Integer32_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 equal", actual)
}

func Test_Cov10_Integer32_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 less", actual)
}

func Test_Cov10_Integer32_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 greater", actual)
}

func Test_Cov10_Integer32Ptr_LeftNil(t *testing.T) {
	r := int32(5)
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr left nil", actual)
}

func Test_Cov10_Integer32Ptr_Equal(t *testing.T) {
	l, r := int32(5), int32(5)
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr equal", actual)
}

// ═══════════════════════════════════════════
// Integer64 — Greater branch
// ═══════════════════════════════════════════

func Test_Cov10_Integer64_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 greater", actual)
}

func Test_Cov10_Integer64Ptr_LeftNil(t *testing.T) {
	r := int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr left nil", actual)
}

func Test_Cov10_Integer64Ptr_Equal(t *testing.T) {
	l, r := int64(5), int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr equal", actual)
}

// ═══════════════════════════════════════════
// BytePtr — remaining branches
// ═══════════════════════════════════════════

func Test_Cov10_BytePtr_RightNil(t *testing.T) {
	l := byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr right nil", actual)
}

func Test_Cov10_BytePtr_Equal(t *testing.T) {
	l, r := byte(5), byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr equal", actual)
}

func Test_Cov10_BytePtr_Less(t *testing.T) {
	l, r := byte(3), byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr less", actual)
}

func Test_Cov10_BytePtr_Greater(t *testing.T) {
	l, r := byte(10), byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "BytePtr greater", actual)
}

// ═══════════════════════════════════════════
// AnyItem — Inconclusive
// ═══════════════════════════════════════════

func Test_Cov10_AnyItem_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(5, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem right nil", actual)
}

func Test_Cov10_AnyItem_Inconclusive(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(5, 10)}
	expected := args.Map{"result": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem inconclusive", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqual — NotEqual items
// ═══════════════════════════════════════════

func Test_Cov10_IsStringsEqual_NotEqualItems(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual not equal items", actual)
}

func Test_Cov10_IsStringsEqual_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual right nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualPtr — DiffLen, RightNil
// ═══════════════════════════════════════════

func Test_Cov10_IsStringsEqualPtr_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr right nil", actual)
}

func Test_Cov10_IsStringsEqualPtr_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr diff len", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualWithoutOrder — all branches
// ═══════════════════════════════════════════

func Test_Cov10_IsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder both nil", actual)
}

func Test_Cov10_IsStringsEqualWithoutOrder_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder left nil", actual)
}

func Test_Cov10_IsStringsEqualWithoutOrder_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder right nil", actual)
}

func Test_Cov10_IsStringsEqualWithoutOrder_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder diff len", actual)
}

func Test_Cov10_IsStringsEqualWithoutOrder_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder equal", actual)
}

func Test_Cov10_IsStringsEqualWithoutOrder_NotEqual(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder not equal", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqual — LeftNil
// ═══════════════════════════════════════════

func Test_Cov10_IsIntegersEqual_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual left nil", actual)
}

func Test_Cov10_IsIntegersEqual_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1}, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual right nil", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqualPtr — RightNil, DiffLen
// ═══════════════════════════════════════════

func Test_Cov10_IsIntegersEqualPtr_RightNil(t *testing.T) {
	l := []int{1}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr right nil", actual)
}

func Test_Cov10_IsIntegersEqualPtr_DiffLen(t *testing.T) {
	l := []int{1}
	r := []int{1, 2}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr diff len", actual)
}

// ═══════════════════════════════════════════
// VersionSliceByte — all branches
// ═══════════════════════════════════════════

func Test_Cov10_VersionSliceByte_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte both nil", actual)
}

func Test_Cov10_VersionSliceByte_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1})}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left nil", actual)
}

func Test_Cov10_VersionSliceByte_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte equal", actual)
}

func Test_Cov10_VersionSliceByte_LeftLess_SameLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 4})}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left less same len", actual)
}

func Test_Cov10_VersionSliceByte_LeftGreater_SameLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 4}, []byte{1, 2, 3})}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left greater same len", actual)
}

func Test_Cov10_VersionSliceByte_LeftLess_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2, 3})}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left less diff len", actual)
}

func Test_Cov10_VersionSliceByte_LeftGreater_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2})}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left greater diff len", actual)
}

// ═══════════════════════════════════════════
// VersionSliceInteger — all branches
// ═══════════════════════════════════════════

func Test_Cov10_VersionSliceInteger_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger both nil", actual)
}

func Test_Cov10_VersionSliceInteger_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1})}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left nil", actual)
}

func Test_Cov10_VersionSliceInteger_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger equal", actual)
}

func Test_Cov10_VersionSliceInteger_LeftLess_SameLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 4})}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left less same len", actual)
}

func Test_Cov10_VersionSliceInteger_LeftGreater_SameLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 4}, []int{1, 2, 3})}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left greater same len", actual)
}

func Test_Cov10_VersionSliceInteger_LeftLess_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2, 3})}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left less diff len", actual)
}

func Test_Cov10_VersionSliceInteger_LeftGreater_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2})}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left greater diff len", actual)
}

// ═══════════════════════════════════════════
// Time — all branches
// ═══════════════════════════════════════════

func Test_Cov10_Time_Equal(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.Time(now, now)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time equal", actual)
}

func Test_Cov10_Time_Less(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	actual := args.Map{"result": corecmp.Time(now, later)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time less", actual)
}

func Test_Cov10_Time_Greater(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	actual := args.Map{"result": corecmp.Time(now, earlier)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time greater", actual)
}

// ═══════════════════════════════════════════
// TimePtr — all branches
// ═══════════════════════════════════════════

func Test_Cov10_TimePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.TimePtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr both nil", actual)
}

func Test_Cov10_TimePtr_LeftNil(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.TimePtr(nil, &now)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr left nil", actual)
}

func Test_Cov10_TimePtr_RightNil(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.TimePtr(&now, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr right nil", actual)
}

func Test_Cov10_TimePtr_Equal(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.TimePtr(&now, &now)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr equal", actual)
}

func Test_Cov10_TimePtr_Less(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	actual := args.Map{"result": corecmp.TimePtr(&now, &later)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr less", actual)
}

func Test_Cov10_TimePtr_Greater(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	actual := args.Map{"result": corecmp.TimePtr(&now, &earlier)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "TimePtr greater", actual)
}
