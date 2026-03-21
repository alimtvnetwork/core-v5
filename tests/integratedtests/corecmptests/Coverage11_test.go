package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── IntegerPtr — all branches ──

func Test_Cov11_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_Cov11_IntegerPtr_LeftNil(t *testing.T) {
	r := 5
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_Cov11_IntegerPtr_RightNil(t *testing.T) {
	l := 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

func Test_Cov11_IntegerPtr_Equal(t *testing.T) {
	l, r := 5, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- equal", actual)
}

func Test_Cov11_IntegerPtr_Less(t *testing.T) {
	l, r := 3, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- less", actual)
}

func Test_Cov11_IntegerPtr_Greater(t *testing.T) {
	l, r := 10, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- greater", actual)
}

// ── Integer — remaining branches ──

func Test_Cov11_Integer_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Cov11_Integer_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- less", actual)
}

// ── Integer8 — remaining branches ──

func Test_Cov11_Integer8_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Cov11_Integer8_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- less", actual)
}

// ── Integer64 — remaining branches ──

func Test_Cov11_Integer64_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Cov11_Integer64_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- less", actual)
}

// ── Byte — all branches ──

func Test_Cov11_Byte_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Cov11_Byte_Less(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(3, 5)}
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- less", actual)
}

func Test_Cov11_Byte_Greater(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(10, 5)}
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- greater", actual)
}

// ── BytePtr — BothNil ──

func Test_Cov11_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_Cov11_BytePtr_LeftNil(t *testing.T) {
	r := byte(5)
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

// ── Integer8Ptr — BothNil ──

func Test_Cov11_Integer8Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Cov11_Integer8Ptr_RightNil(t *testing.T) {
	l := int8(5)
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- right nil", actual)
}

// ── Integer16Ptr — BothNil ──

func Test_Cov11_Integer16Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Cov11_Integer16Ptr_RightNil(t *testing.T) {
	l := int16(5)
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- right nil", actual)
}

// ── Integer32Ptr — BothNil ──

func Test_Cov11_Integer32Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Cov11_Integer32Ptr_RightNil(t *testing.T) {
	l := int32(5)
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- right nil", actual)
}

// ── Integer64Ptr — BothNil ──

func Test_Cov11_Integer64Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Cov11_Integer64Ptr_RightNil(t *testing.T) {
	l := int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

// ── AnyItem — all branches ──

func Test_Cov11_AnyItem_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_Cov11_AnyItem_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, 5)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_Cov11_AnyItem_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(5, 5)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

// ── IsStringsEqual — all equal ──

func Test_Cov11_IsStringsEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_Cov11_IsStringsEqual_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- diff len", actual)
}

func Test_Cov11_IsStringsEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- equal", actual)
}

func Test_Cov11_IsStringsEqual_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

// ── IsStringsEqualPtr — all equal ──

func Test_Cov11_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_Cov11_IsStringsEqualPtr_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- equal", actual)
}

func Test_Cov11_IsStringsEqualPtr_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- left nil", actual)
}

// ── IsIntegersEqual — both nil, equal ──

func Test_Cov11_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_Cov11_IsIntegersEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- equal", actual)
}

// ── IsIntegersEqualPtr — both nil, equal ──

func Test_Cov11_IsIntegersEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_Cov11_IsIntegersEqualPtr_LeftNil(t *testing.T) {
	r := []int{1}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, &r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_Cov11_IsIntegersEqualPtr_Equal(t *testing.T) {
	l := []int{1, 2}
	r := []int{1, 2}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- equal", actual)
}

// ── VersionSliceByte — RightNil ──

func Test_Cov11_VersionSliceByte_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- right nil", actual)
}

// ── VersionSliceInteger — RightNil ──

func Test_Cov11_VersionSliceInteger_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- right nil", actual)
}
