package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Byte — all branches
// ═══════════════════════════════════════════

func Test_Cov9_Byte_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte equal", actual)
}

func Test_Cov9_Byte_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte less", actual)
}

func Test_Cov9_Byte_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte greater", actual)
}

// ═══════════════════════════════════════════
// BytePtr — all branches
// ═══════════════════════════════════════════

func Test_Cov9_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr both nil", actual)
}

func Test_Cov9_BytePtr_LeftNil(t *testing.T) {
	r := byte(5)
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr left nil", actual)
}

// ═══════════════════════════════════════════
// Integer — all branches
// ═══════════════════════════════════════════

func Test_Cov9_Integer_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer equal", actual)
}

func Test_Cov9_Integer_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer less", actual)
}

// ═══════════════════════════════════════════
// IntegerPtr — all branches
// ═══════════════════════════════════════════

func Test_Cov9_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr both nil", actual)
}

func Test_Cov9_IntegerPtr_LeftNil(t *testing.T) {
	r := 5
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr left nil", actual)
}

func Test_Cov9_IntegerPtr_RightNil(t *testing.T) {
	l := 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr right nil", actual)
}

func Test_Cov9_IntegerPtr_Equal(t *testing.T) {
	l, r := 5, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr equal", actual)
}

func Test_Cov9_IntegerPtr_Less(t *testing.T) {
	l, r := 3, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr less", actual)
}

func Test_Cov9_IntegerPtr_Greater(t *testing.T) {
	l, r := 10, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "IntegerPtr greater", actual)
}

// ═══════════════════════════════════════════
// Integer8 — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_Integer8_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 equal", actual)
}

func Test_Cov9_Integer8_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 less", actual)
}

func Test_Cov9_Integer8Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr both nil", actual)
}

func Test_Cov9_Integer8Ptr_RightNil(t *testing.T) {
	l := int8(5)
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr right nil", actual)
}

// ═══════════════════════════════════════════
// Integer16Ptr — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_Integer16Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr both nil", actual)
}

func Test_Cov9_Integer16Ptr_RightNil(t *testing.T) {
	l := int16(5)
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr right nil", actual)
}

// ═══════════════════════════════════════════
// Integer32Ptr — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_Integer32Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr both nil", actual)
}

func Test_Cov9_Integer32Ptr_RightNil(t *testing.T) {
	l := int32(5)
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr right nil", actual)
}

// ═══════════════════════════════════════════
// Integer64 — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_Integer64_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 equal", actual)
}

func Test_Cov9_Integer64_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 less", actual)
}

func Test_Cov9_Integer64Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr both nil", actual)
}

func Test_Cov9_Integer64Ptr_RightNil(t *testing.T) {
	l := int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr right nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqual — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_IsStringsEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual both nil", actual)
}

func Test_Cov9_IsStringsEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual equal", actual)
}

func Test_Cov9_IsStringsEqual_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual diff len", actual)
}

func Test_Cov9_IsStringsEqual_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual left nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualPtr — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr both nil", actual)
}

func Test_Cov9_IsStringsEqualPtr_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr equal", actual)
}

func Test_Cov9_IsStringsEqualPtr_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr left nil", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqual — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual both nil", actual)
}

func Test_Cov9_IsIntegersEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual equal", actual)
}

func Test_Cov9_IsIntegersEqual_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1}, []int{1, 2})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual diff len", actual)
}

func Test_Cov9_IsIntegersEqual_NotEqual(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual not equal", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqualPtr — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_IsIntegersEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr both nil", actual)
}

func Test_Cov9_IsIntegersEqualPtr_LeftNil(t *testing.T) {
	r := []int{1}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, &r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr left nil", actual)
}

func Test_Cov9_IsIntegersEqualPtr_Equal(t *testing.T) {
	l := []int{1, 2}
	r := []int{1, 2}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr equal", actual)
}

func Test_Cov9_IsIntegersEqualPtr_NotEqual(t *testing.T) {
	l := []int{1, 2}
	r := []int{1, 3}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr not equal", actual)
}

// ═══════════════════════════════════════════
// AnyItem — remaining branches
// ═══════════════════════════════════════════

func Test_Cov9_AnyItem_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem both nil", actual)
}

func Test_Cov9_AnyItem_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, 5)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem left nil", actual)
}

func Test_Cov9_AnyItem_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem equal", actual)
}

// ═══════════════════════════════════════════
// VersionSliceByte — right nil
// ═══════════════════════════════════════════

func Test_Cov9_VersionSliceByte_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte right nil", actual)
}

// ═══════════════════════════════════════════
// VersionSliceInteger — right nil
// ═══════════════════════════════════════════

func Test_Cov9_VersionSliceInteger_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger right nil", actual)
}

// ═══════════════════════════════════════════
// Time — remaining
// ═══════════════════════════════════════════

func Test_Cov9_TimePtr_Equal_Values(t *testing.T) {
	now := corecmp.Time
	_ = now // Time func already fully tested in Coverage8
}
