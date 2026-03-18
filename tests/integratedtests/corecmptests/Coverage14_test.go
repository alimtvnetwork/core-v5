package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyItem — all 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_AnyItem_BothNil(t *testing.T) {
	r := corecmp.AnyItem(nil, nil)
	actual := args.Map{"v": r}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem both nil → Equal", actual)
}

func Test_Cov14_AnyItem_LeftNilOnly(t *testing.T) {
	r := corecmp.AnyItem(nil, 1)
	actual := args.Map{"v": r}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem left nil → NotEqual", actual)
}

func Test_Cov14_AnyItem_RightNilOnly(t *testing.T) {
	r := corecmp.AnyItem(1, nil)
	actual := args.Map{"v": r}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem right nil → NotEqual", actual)
}

func Test_Cov14_AnyItem_SameValue(t *testing.T) {
	r := corecmp.AnyItem(42, 42)
	actual := args.Map{"v": r}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem same → Equal", actual)
}

func Test_Cov14_AnyItem_DiffValue(t *testing.T) {
	r := corecmp.AnyItem(1, 2)
	actual := args.Map{"v": r}
	expected := args.Map{"v": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem diff → Inconclusive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Byte — 3 branches: Equal, LeftLess, LeftGreater
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Byte_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Byte(5, 5)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte equal", actual)
}

func Test_Cov14_Byte_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Byte(1, 9)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte less", actual)
}

func Test_Cov14_Byte_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Byte(9, 1)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytePtr — 4 branches: BothNil, LeftNil, RightNil, Delegate
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.BytePtr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr both nil", actual)
}

func Test_Cov14_BytePtr_LeftNil(t *testing.T) {
	b := byte(1)
	actual := args.Map{"v": corecmp.BytePtr(nil, &b)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr left nil", actual)
}

func Test_Cov14_BytePtr_RightNil(t *testing.T) {
	b := byte(1)
	actual := args.Map{"v": corecmp.BytePtr(&b, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr right nil", actual)
}

func Test_Cov14_BytePtr_Delegate(t *testing.T) {
	a, b := byte(3), byte(7)
	actual := args.Map{"v": corecmp.BytePtr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer — 3 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Integer_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer(10, 10)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer equal", actual)
}

func Test_Cov14_Integer_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer(-5, 5)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer less", actual)
}

func Test_Cov14_Integer_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer(5, -5)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IntegerPtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IntegerPtr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr both nil", actual)
}

func Test_Cov14_IntegerPtr_LeftNil(t *testing.T) {
	v := 1
	actual := args.Map{"v": corecmp.IntegerPtr(nil, &v)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr left nil", actual)
}

func Test_Cov14_IntegerPtr_RightNil(t *testing.T) {
	v := 1
	actual := args.Map{"v": corecmp.IntegerPtr(&v, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr right nil", actual)
}

func Test_Cov14_IntegerPtr_Delegate(t *testing.T) {
	a, b := 10, 20
	actual := args.Map{"v": corecmp.IntegerPtr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer8 — 3 branches + Integer8Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Integer8_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer8(5, 5)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 equal", actual)
}

func Test_Cov14_Integer8_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer8(-10, 10)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 less", actual)
}

func Test_Cov14_Integer8_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer8(10, -10)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 greater", actual)
}

func Test_Cov14_Integer8Ptr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer8Ptr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr both nil", actual)
}

func Test_Cov14_Integer8Ptr_LeftNil(t *testing.T) {
	v := int8(1)
	actual := args.Map{"v": corecmp.Integer8Ptr(nil, &v)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr left nil", actual)
}

func Test_Cov14_Integer8Ptr_RightNil(t *testing.T) {
	v := int8(1)
	actual := args.Map{"v": corecmp.Integer8Ptr(&v, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr right nil", actual)
}

func Test_Cov14_Integer8Ptr_Delegate(t *testing.T) {
	a, b := int8(3), int8(7)
	actual := args.Map{"v": corecmp.Integer8Ptr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer16 — 3 branches + Integer16Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Integer16_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer16(100, 100)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 equal", actual)
}

func Test_Cov14_Integer16_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer16(-100, 100)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 less", actual)
}

func Test_Cov14_Integer16_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer16(100, -100)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 greater", actual)
}

func Test_Cov14_Integer16Ptr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer16Ptr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr both nil", actual)
}

func Test_Cov14_Integer16Ptr_LeftNil(t *testing.T) {
	v := int16(1)
	actual := args.Map{"v": corecmp.Integer16Ptr(nil, &v)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr left nil", actual)
}

func Test_Cov14_Integer16Ptr_RightNil(t *testing.T) {
	v := int16(1)
	actual := args.Map{"v": corecmp.Integer16Ptr(&v, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr right nil", actual)
}

func Test_Cov14_Integer16Ptr_Delegate(t *testing.T) {
	a, b := int16(3), int16(7)
	actual := args.Map{"v": corecmp.Integer16Ptr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer32 — 3 branches + Integer32Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Integer32_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer32(1000, 1000)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 equal", actual)
}

func Test_Cov14_Integer32_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer32(-1000, 1000)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 less", actual)
}

func Test_Cov14_Integer32_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer32(1000, -1000)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 greater", actual)
}

func Test_Cov14_Integer32Ptr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer32Ptr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr both nil", actual)
}

func Test_Cov14_Integer32Ptr_LeftNil(t *testing.T) {
	v := int32(1)
	actual := args.Map{"v": corecmp.Integer32Ptr(nil, &v)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr left nil", actual)
}

func Test_Cov14_Integer32Ptr_RightNil(t *testing.T) {
	v := int32(1)
	actual := args.Map{"v": corecmp.Integer32Ptr(&v, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr right nil", actual)
}

func Test_Cov14_Integer32Ptr_Delegate(t *testing.T) {
	a, b := int32(3), int32(7)
	actual := args.Map{"v": corecmp.Integer32Ptr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer64 — 3 branches + Integer64Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Integer64_Equal(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer64(99999, 99999)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 equal", actual)
}

func Test_Cov14_Integer64_Less(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer64(-99999, 99999)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 less", actual)
}

func Test_Cov14_Integer64_Greater(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer64(99999, -99999)}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 greater", actual)
}

func Test_Cov14_Integer64Ptr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.Integer64Ptr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr both nil", actual)
}

func Test_Cov14_Integer64Ptr_LeftNil(t *testing.T) {
	v := int64(1)
	actual := args.Map{"v": corecmp.Integer64Ptr(nil, &v)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr left nil", actual)
}

func Test_Cov14_Integer64Ptr_RightNil(t *testing.T) {
	v := int64(1)
	actual := args.Map{"v": corecmp.Integer64Ptr(&v, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr right nil", actual)
}

func Test_Cov14_Integer64Ptr_Delegate(t *testing.T) {
	a, b := int64(3), int64(7)
	actual := args.Map{"v": corecmp.Integer64Ptr(&a, &b)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqual — 4 branches (nil/nil, nil/val, val/nil, delegate)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual both nil", actual)
}

func Test_Cov14_IsIntegersEqual_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsIntegersEqual(nil, []int{1})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual left nil", actual)
}

func Test_Cov14_IsIntegersEqual_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsIntegersEqual([]int{1}, nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual right nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqualPtr — 4 branches (nil/nil, nil/val, diffLen, delegate)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IsIntegersEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr both nil", actual)
}

func Test_Cov14_IsIntegersEqualPtr_LeftNil(t *testing.T) {
	r := []int{1}
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(nil, &r)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr left nil", actual)
}

func Test_Cov14_IsIntegersEqualPtr_RightNil(t *testing.T) {
	l := []int{1}
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(&l, nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr right nil", actual)
}

func Test_Cov14_IsIntegersEqualPtr_DiffLen(t *testing.T) {
	l := []int{1}
	r := []int{1, 2}
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(&l, &r)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr diff len", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqual — 5 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IsStringsEqual_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual(nil, nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual both nil", actual)
}

func Test_Cov14_IsStringsEqual_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual(nil, []string{"a"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual left nil", actual)
}

func Test_Cov14_IsStringsEqual_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a"}, nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual right nil", actual)
}

func Test_Cov14_IsStringsEqual_DiffLen(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual diff len", actual)
}

func Test_Cov14_IsStringsEqual_Same(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual same", actual)
}

func Test_Cov14_IsStringsEqual_Diff(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual diff", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualPtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr both nil", actual)
}

func Test_Cov14_IsStringsEqualPtr_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr left nil", actual)
}

func Test_Cov14_IsStringsEqualPtr_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr right nil", actual)
}

func Test_Cov14_IsStringsEqualPtr_DiffLen(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr diff len", actual)
}

func Test_Cov14_IsStringsEqualPtr_Same(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"x", "y"}, []string{"x", "y"})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualWithoutOrder — 5 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_IsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder(nil, nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "WithoutOrder both nil", actual)
}

func Test_Cov14_IsStringsEqualWithoutOrder_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder left nil", actual)
}

func Test_Cov14_IsStringsEqualWithoutOrder_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder right nil", actual)
}

func Test_Cov14_IsStringsEqualWithoutOrder_DiffLen(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder diff len", actual)
}

func Test_Cov14_IsStringsEqualWithoutOrder_Reordered(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "WithoutOrder reordered", actual)
}

func Test_Cov14_IsStringsEqualWithoutOrder_Mismatch(t *testing.T) {
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Time — 3 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Time_Equal(t *testing.T) {
	now := time.Now()
	actual := args.Map{"v": corecmp.Time(now, now)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time equal", actual)
}

func Test_Cov14_Time_Less(t *testing.T) {
	now := time.Now()
	actual := args.Map{"v": corecmp.Time(now, now.Add(time.Hour))}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time less", actual)
}

func Test_Cov14_Time_Greater(t *testing.T) {
	now := time.Now()
	actual := args.Map{"v": corecmp.Time(now, now.Add(-time.Hour))}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TimePtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_TimePtr_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.TimePtr(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr both nil", actual)
}

func Test_Cov14_TimePtr_LeftNil(t *testing.T) {
	now := time.Now()
	actual := args.Map{"v": corecmp.TimePtr(nil, &now)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr left nil", actual)
}

func Test_Cov14_TimePtr_RightNil(t *testing.T) {
	now := time.Now()
	actual := args.Map{"v": corecmp.TimePtr(&now, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr right nil", actual)
}

func Test_Cov14_TimePtr_Delegate(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	actual := args.Map{"v": corecmp.TimePtr(&now, &later)}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceByte — 7 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_VersionSliceByte_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSB both nil", actual)
}

func Test_Cov14_VersionSliceByte_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte(nil, []byte{1})}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSB left nil", actual)
}

func Test_Cov14_VersionSliceByte_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1}, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSB right nil", actual)
}

func Test_Cov14_VersionSliceByte_EqualSameLen(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSB equal same len", actual)
}

func Test_Cov14_VersionSliceByte_LoopLeftLess(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0, 1})}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSB loop left less", actual)
}

func Test_Cov14_VersionSliceByte_LoopLeftGreater(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{2, 0, 0}, []byte{1, 9, 9})}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSB loop left greater", actual)
}

func Test_Cov14_VersionSliceByte_ShorterLeft(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0}, []byte{1, 0, 0})}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSB shorter left", actual)
}

func Test_Cov14_VersionSliceByte_ShorterRight(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0})}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSB shorter right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceInteger — 7 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_VersionSliceInteger_BothNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger(nil, nil)}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSI both nil", actual)
}

func Test_Cov14_VersionSliceInteger_LeftNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger(nil, []int{1})}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSI left nil", actual)
}

func Test_Cov14_VersionSliceInteger_RightNil(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1}, nil)}
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSI right nil", actual)
}

func Test_Cov14_VersionSliceInteger_EqualSameLen(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})}
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSI equal same len", actual)
}

func Test_Cov14_VersionSliceInteger_LoopLeftLess(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0, 1})}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSI loop left less", actual)
}

func Test_Cov14_VersionSliceInteger_LoopLeftGreater(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{2, 0, 0}, []int{1, 9, 9})}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSI loop left greater", actual)
}

func Test_Cov14_VersionSliceInteger_ShorterLeft(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0}, []int{1, 0, 0})}
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSI shorter left", actual)
}

func Test_Cov14_VersionSliceInteger_ShorterRight(t *testing.T) {
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0})}
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSI shorter right", actual)
}
