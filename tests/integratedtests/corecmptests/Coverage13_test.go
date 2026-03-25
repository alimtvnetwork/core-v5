package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyItem
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AnyItem_BothNil(t *testing.T) {
	result := corecmp.AnyItem(nil, nil)
	actual := args.Map{"equal": result.IsEqual()}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_Cov13_AnyItem_LeftNil(t *testing.T) {
	result := corecmp.AnyItem(nil, "hello")
	actual := args.Map{"notEqual": result.IsNotEqual()}
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_Cov13_AnyItem_RightNil(t *testing.T) {
	result := corecmp.AnyItem("hello", nil)
	actual := args.Map{"notEqual": result.IsNotEqual()}
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_Cov13_AnyItem_Equal(t *testing.T) {
	result := corecmp.AnyItem(42, 42)
	actual := args.Map{"equal": result.IsEqual()}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

func Test_Cov13_AnyItem_Inconclusive(t *testing.T) {
	result := corecmp.AnyItem(42, 99)
	actual := args.Map{"inconclusive": result.IsInconclusive()}
	expected := args.Map{"inconclusive": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

func Test_Cov13_AnyItem_EqualStrings(t *testing.T) {
	result := corecmp.AnyItem("abc", "abc")
	actual := args.Map{"equal": result.IsEqual()}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal strings", actual)
}

func Test_Cov13_AnyItem_DifferentStrings(t *testing.T) {
	result := corecmp.AnyItem("abc", "xyz")
	actual := args.Map{"inconclusive": result.IsInconclusive()}
	expected := args.Map{"inconclusive": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- different strings", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Byte
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Byte_Equal(t *testing.T) {
	result := corecmp.Byte(5, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Cov13_Byte_LeftLess(t *testing.T) {
	result := corecmp.Byte(1, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left less", actual)
}

func Test_Cov13_Byte_LeftGreater(t *testing.T) {
	result := corecmp.Byte(10, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left greater", actual)
}

func Test_Cov13_Byte_Zero(t *testing.T) {
	result := corecmp.Byte(0, 0)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- zero", actual)
}

func Test_Cov13_Byte_Max(t *testing.T) {
	result := corecmp.Byte(255, 0)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- max", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_BytePtr_BothNil(t *testing.T) {
	result := corecmp.BytePtr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_Cov13_BytePtr_LeftNil(t *testing.T) {
	b := byte(5)
	result := corecmp.BytePtr(nil, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

func Test_Cov13_BytePtr_RightNil(t *testing.T) {
	b := byte(5)
	result := corecmp.BytePtr(&b, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

func Test_Cov13_BytePtr_Equal(t *testing.T) {
	a, b := byte(5), byte(5)
	result := corecmp.BytePtr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- equal", actual)
}

func Test_Cov13_BytePtr_LeftLess(t *testing.T) {
	a, b := byte(1), byte(9)
	result := corecmp.BytePtr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- left less", actual)
}

func Test_Cov13_BytePtr_LeftGreater(t *testing.T) {
	a, b := byte(9), byte(1)
	result := corecmp.BytePtr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- left greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Integer_Equal(t *testing.T) {
	result := corecmp.Integer(42, 42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Cov13_Integer_LeftLess(t *testing.T) {
	result := corecmp.Integer(1, 100)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left less", actual)
}

func Test_Cov13_Integer_LeftGreater(t *testing.T) {
	result := corecmp.Integer(100, 1)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left greater", actual)
}

func Test_Cov13_Integer_Negative(t *testing.T) {
	result := corecmp.Integer(-5, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- negative", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IntegerPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_IntegerPtr_BothNil(t *testing.T) {
	result := corecmp.IntegerPtr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_Cov13_IntegerPtr_LeftNil(t *testing.T) {
	v := 5
	result := corecmp.IntegerPtr(nil, &v)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_Cov13_IntegerPtr_RightNil(t *testing.T) {
	v := 5
	result := corecmp.IntegerPtr(&v, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

func Test_Cov13_IntegerPtr_Equal(t *testing.T) {
	a, b := 42, 42
	result := corecmp.IntegerPtr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- equal", actual)
}

func Test_Cov13_IntegerPtr_LeftLess(t *testing.T) {
	a, b := 1, 99
	result := corecmp.IntegerPtr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer8 / Integer8Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Integer8_Equal(t *testing.T) {
	result := corecmp.Integer8(5, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Cov13_Integer8_LeftLess(t *testing.T) {
	result := corecmp.Integer8(-10, 10)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left less", actual)
}

func Test_Cov13_Integer8_LeftGreater(t *testing.T) {
	result := corecmp.Integer8(10, -10)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left greater", actual)
}

func Test_Cov13_Integer8Ptr_BothNil(t *testing.T) {
	result := corecmp.Integer8Ptr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Cov13_Integer8Ptr_LeftNil(t *testing.T) {
	v := int8(5)
	result := corecmp.Integer8Ptr(nil, &v)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Cov13_Integer8Ptr_RightNil(t *testing.T) {
	v := int8(5)
	result := corecmp.Integer8Ptr(&v, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- right nil", actual)
}

func Test_Cov13_Integer8Ptr_Equal(t *testing.T) {
	a, b := int8(3), int8(3)
	result := corecmp.Integer8Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer16 / Integer16Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Integer16_Equal(t *testing.T) {
	result := corecmp.Integer16(100, 100)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- equal", actual)
}

func Test_Cov13_Integer16_LeftLess(t *testing.T) {
	result := corecmp.Integer16(-100, 100)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left less", actual)
}

func Test_Cov13_Integer16_LeftGreater(t *testing.T) {
	result := corecmp.Integer16(100, -100)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left greater", actual)
}

func Test_Cov13_Integer16Ptr_BothNil(t *testing.T) {
	result := corecmp.Integer16Ptr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Cov13_Integer16Ptr_LeftNil(t *testing.T) {
	v := int16(5)
	result := corecmp.Integer16Ptr(nil, &v)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Cov13_Integer16Ptr_RightNil(t *testing.T) {
	v := int16(5)
	result := corecmp.Integer16Ptr(&v, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- right nil", actual)
}

func Test_Cov13_Integer16Ptr_Equal(t *testing.T) {
	a, b := int16(7), int16(7)
	result := corecmp.Integer16Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer32 / Integer32Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Integer32_Equal(t *testing.T) {
	result := corecmp.Integer32(1000, 1000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Cov13_Integer32_LeftLess(t *testing.T) {
	result := corecmp.Integer32(-1000, 1000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left less", actual)
}

func Test_Cov13_Integer32_LeftGreater(t *testing.T) {
	result := corecmp.Integer32(1000, -1000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left greater", actual)
}

func Test_Cov13_Integer32Ptr_BothNil(t *testing.T) {
	result := corecmp.Integer32Ptr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Cov13_Integer32Ptr_LeftNil(t *testing.T) {
	v := int32(5)
	result := corecmp.Integer32Ptr(nil, &v)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Cov13_Integer32Ptr_RightNil(t *testing.T) {
	v := int32(5)
	result := corecmp.Integer32Ptr(&v, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- right nil", actual)
}

func Test_Cov13_Integer32Ptr_Equal(t *testing.T) {
	a, b := int32(7), int32(7)
	result := corecmp.Integer32Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- equal", actual)
}

func Test_Cov13_Integer32Ptr_LeftLess(t *testing.T) {
	a, b := int32(1), int32(99)
	result := corecmp.Integer32Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer64 / Integer64Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Integer64_Equal(t *testing.T) {
	result := corecmp.Integer64(100000, 100000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Cov13_Integer64_LeftLess(t *testing.T) {
	result := corecmp.Integer64(-100000, 100000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left less", actual)
}

func Test_Cov13_Integer64_LeftGreater(t *testing.T) {
	result := corecmp.Integer64(100000, -100000)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left greater", actual)
}

func Test_Cov13_Integer64Ptr_BothNil(t *testing.T) {
	result := corecmp.Integer64Ptr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Cov13_Integer64Ptr_LeftNil(t *testing.T) {
	v := int64(5)
	result := corecmp.Integer64Ptr(nil, &v)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

func Test_Cov13_Integer64Ptr_RightNil(t *testing.T) {
	v := int64(5)
	result := corecmp.Integer64Ptr(&v, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

func Test_Cov13_Integer64Ptr_Equal(t *testing.T) {
	a, b := int64(7), int64(7)
	result := corecmp.Integer64Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- equal", actual)
}

func Test_Cov13_Integer64Ptr_LeftGreater(t *testing.T) {
	a, b := int64(99), int64(1)
	result := corecmp.Integer64Ptr(&a, &b)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- left greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqual / IsIntegersEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_Cov13_IsIntegersEqual_LeftNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual(nil, []int{1})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- left nil", actual)
}

func Test_Cov13_IsIntegersEqual_RightNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1}, nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- right nil", actual)
}

func Test_Cov13_IsIntegersEqual_Same(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1, 2, 3}, []int{1, 2, 3})}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- same", actual)
}

func Test_Cov13_IsIntegersEqual_Different(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- different", actual)
}

func Test_Cov13_IsIntegersEqual_DifferentLength(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1}, []int{1, 2})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- different length", actual)
}

func Test_Cov13_IsIntegersEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_Cov13_IsIntegersEqualPtr_LeftNil(t *testing.T) {
	right := []int{1}
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(nil, &right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_Cov13_IsIntegersEqualPtr_RightNil(t *testing.T) {
	left := []int{1}
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- right nil", actual)
}

func Test_Cov13_IsIntegersEqualPtr_DifferentLength(t *testing.T) {
	left := []int{1}
	right := []int{1, 2}
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, &right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- different length", actual)
}

func Test_Cov13_IsIntegersEqualPtr_Same(t *testing.T) {
	left := []int{1, 2, 3}
	right := []int{1, 2, 3}
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, &right)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqual / IsStringsEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_IsStringsEqual_BothNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual(nil, nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_Cov13_IsStringsEqual_LeftNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual(nil, []string{"a"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

func Test_Cov13_IsStringsEqual_RightNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- right nil", actual)
}

func Test_Cov13_IsStringsEqual_Same(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- same", actual)
}

func Test_Cov13_IsStringsEqual_Different(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, []string{"b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different", actual)
}

func Test_Cov13_IsStringsEqual_DifferentLength(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different length", actual)
}

func Test_Cov13_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_Cov13_IsStringsEqualPtr_LeftNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- left nil", actual)
}

func Test_Cov13_IsStringsEqualPtr_RightNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- right nil", actual)
}

func Test_Cov13_IsStringsEqualPtr_DifferentLength(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- different length", actual)
}

func Test_Cov13_IsStringsEqualPtr_Same(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"x", "y"}, []string{"x", "y"})}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualWithoutOrder
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_IsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder(nil, nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- both nil", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_LeftNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- left nil", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_RightNil(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- right nil", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_DifferentLength(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- different length", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_SameOrder(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- same order", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_DifferentOrder(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- different order", actual)
}

func Test_Cov13_IsStringsEqualWithoutOrder_Mismatch(t *testing.T) {
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Time / TimePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Time_Equal(t *testing.T) {
	now := time.Now()
	result := corecmp.Time(now, now)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Cov13_Time_LeftLess(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	result := corecmp.Time(now, later)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left less", actual)
}

func Test_Cov13_Time_LeftGreater(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	result := corecmp.Time(now, earlier)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left greater", actual)
}

func Test_Cov13_TimePtr_BothNil(t *testing.T) {
	result := corecmp.TimePtr(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_Cov13_TimePtr_LeftNil(t *testing.T) {
	now := time.Now()
	result := corecmp.TimePtr(nil, &now)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_Cov13_TimePtr_RightNil(t *testing.T) {
	now := time.Now()
	result := corecmp.TimePtr(&now, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_Cov13_TimePtr_Equal(t *testing.T) {
	now := time.Now()
	same := now
	result := corecmp.TimePtr(&now, &same)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

func Test_Cov13_TimePtr_LeftLess(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	result := corecmp.TimePtr(&now, &later)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceByte
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_VersionSliceByte_BothNil(t *testing.T) {
	result := corecmp.VersionSliceByte(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_Cov13_VersionSliceByte_LeftNil(t *testing.T) {
	result := corecmp.VersionSliceByte(nil, []byte{1, 0, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- left nil", actual)
}

func Test_Cov13_VersionSliceByte_RightNil(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- right nil", actual)
}

func Test_Cov13_VersionSliceByte_Equal(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_Cov13_VersionSliceByte_LeftLess(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0, 1})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less", actual)
}

func Test_Cov13_VersionSliceByte_LeftGreater(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{2, 0, 0}, []byte{1, 9, 9})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater", actual)
}

func Test_Cov13_VersionSliceByte_ShorterLeft(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{1, 0}, []byte{1, 0, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter left", actual)
}

func Test_Cov13_VersionSliceByte_ShorterRight(t *testing.T) {
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceInteger
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_VersionSliceInteger_BothNil(t *testing.T) {
	result := corecmp.VersionSliceInteger(nil, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_Cov13_VersionSliceInteger_LeftNil(t *testing.T) {
	result := corecmp.VersionSliceInteger(nil, []int{1, 0, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- left nil", actual)
}

func Test_Cov13_VersionSliceInteger_RightNil(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- right nil", actual)
}

func Test_Cov13_VersionSliceInteger_Equal(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_Cov13_VersionSliceInteger_LeftLess(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0, 1})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less", actual)
}

func Test_Cov13_VersionSliceInteger_LeftGreater(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{2, 0, 0}, []int{1, 9, 9})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater", actual)
}

func Test_Cov13_VersionSliceInteger_ShorterLeft(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{1, 0}, []int{1, 0, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter left", actual)
}

func Test_Cov13_VersionSliceInteger_ShorterRight(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter right", actual)
}

func Test_Cov13_VersionSliceInteger_Empty(t *testing.T) {
	result := corecmp.VersionSliceInteger([]int{}, []int{})
	actual := args.Map{"val": result}
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns empty -- empty", actual)
}
