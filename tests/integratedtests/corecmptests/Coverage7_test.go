package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Integer ──

func Test_Cov7_Integer(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.Integer(5, 5),
		"less":    corecmp.Integer(3, 5),
		"greater": corecmp.Integer(7, 5),
	}
	expected := args.Map{
		"equal": corecomparator.Equal, "less": corecomparator.LeftLess, "greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer", actual)
}

// ── IntegerPtr ──

func Test_Cov7_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr both nil", actual)
}

func Test_Cov7_IntegerPtr_LeftNil(t *testing.T) {
	r := 5
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr left nil", actual)
}

func Test_Cov7_IntegerPtr_Equal(t *testing.T) {
	l, r := 5, 5
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr equal", actual)
}

// ── Integer8 / Integer8Ptr ──

func Test_Cov7_Integer8(t *testing.T) {
	actual := args.Map{
		"equal": corecmp.Integer8(5, 5),
		"less":  corecmp.Integer8(3, 5),
	}
	expected := args.Map{"equal": corecomparator.Equal, "less": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8", actual)
}

func Test_Cov7_Integer8Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr both nil", actual)
}

// ── Integer16 / Integer16Ptr ──

func Test_Cov7_Integer16(t *testing.T) {
	actual := args.Map{"equal": corecmp.Integer16(5, 5)}
	expected := args.Map{"equal": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16", actual)
}

func Test_Cov7_Integer16Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr both nil", actual)
}

// ── Integer32 / Integer32Ptr ──

func Test_Cov7_Integer32(t *testing.T) {
	actual := args.Map{"equal": corecmp.Integer32(5, 5)}
	expected := args.Map{"equal": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32", actual)
}

func Test_Cov7_Integer32Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr both nil", actual)
}

// ── Integer64 / Integer64Ptr ──

func Test_Cov7_Integer64(t *testing.T) {
	actual := args.Map{"equal": corecmp.Integer64(5, 5), "less": corecmp.Integer64(3, 5)}
	expected := args.Map{"equal": corecomparator.Equal, "less": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64", actual)
}

func Test_Cov7_Integer64Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr both nil", actual)
}

func Test_Cov7_Integer64Ptr_RightNil(t *testing.T) {
	l := int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr right nil", actual)
}

// ── Byte / BytePtr ──

func Test_Cov7_Byte(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.Byte(5, 5),
		"less":    corecmp.Byte(3, 5),
		"greater": corecmp.Byte(7, 5),
	}
	expected := args.Map{"equal": corecomparator.Equal, "less": corecomparator.LeftLess, "greater": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte", actual)
}

func Test_Cov7_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr both nil", actual)
}

func Test_Cov7_BytePtr_LeftNil(t *testing.T) {
	r := byte(5)
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr left nil", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_Cov7_IsIntegersEqual(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2}),
		"notEq":   corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3}),
		"diffLen": corecmp.IsIntegersEqual([]int{1}, []int{1, 2}),
		"bothNil": corecmp.IsIntegersEqual(nil, nil),
	}
	expected := args.Map{"equal": true, "notEq": false, "diffLen": false, "bothNil": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual", actual)
}

func Test_Cov7_IsIntegersEqualPtr(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.IsIntegersEqualPtr([]int{1, 2}, []int{1, 2}),
		"bothNil": corecmp.IsIntegersEqualPtr(nil, nil),
		"leftNil": corecmp.IsIntegersEqualPtr(nil, []int{1}),
	}
	expected := args.Map{"equal": true, "bothNil": true, "leftNil": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr ──

func Test_Cov7_IsStringsEqual(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"}),
		"bothNil": corecmp.IsStringsEqual(nil, nil),
		"leftNil": corecmp.IsStringsEqual(nil, []string{"a"}),
	}
	expected := args.Map{"equal": true, "bothNil": true, "leftNil": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual", actual)
}

func Test_Cov7_IsStringsEqualPtr(t *testing.T) {
	actual := args.Map{
		"equal":   corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a"}),
		"bothNil": corecmp.IsStringsEqualPtr(nil, nil),
	}
	expected := args.Map{"equal": true, "bothNil": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr", actual)
}

// ── AnyItem ──

func Test_Cov7_AnyItem(t *testing.T) {
	actual := args.Map{
		"bothNil": corecmp.AnyItem(nil, nil),
		"leftNil": corecmp.AnyItem(nil, 5),
		"equal":   corecmp.AnyItem(5, 5),
		"notEq":   corecmp.AnyItem(5, 6),
	}
	expected := args.Map{
		"bothNil": corecomparator.Equal,
		"leftNil": corecomparator.NotEqual,
		"equal":   corecomparator.Equal,
		"notEq":   corecomparator.NotEqual,
	}
	expected.ShouldBeEqual(t, 0, "AnyItem", actual)
}

// ── VersionSliceInteger — RightNil ──

func Test_Cov7_VersionSliceInteger_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger right nil", actual)
}
