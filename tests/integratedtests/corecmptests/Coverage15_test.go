package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyItem ──

func Test_Cov15_AnyItem_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem both nil", actual)
}

func Test_Cov15_AnyItem_LeftNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem(nil, "a") == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem left nil", actual)
}

func Test_Cov15_AnyItem_RightNil(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem("a", nil) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem right nil", actual)
}

func Test_Cov15_AnyItem_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem("a", "a") == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem equal", actual)
}

func Test_Cov15_AnyItem_Inconclusive(t *testing.T) {
	actual := args.Map{"result": corecmp.AnyItem("a", "b") == corecomparator.Inconclusive}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem inconclusive", actual)
}

// ── Byte / BytePtr ──

func Test_Cov15_Byte_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(1, 1) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte equal", actual)
}

func Test_Cov15_Byte_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(1, 2) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte left less", actual)
}

func Test_Cov15_Byte_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Byte(2, 1) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte left greater", actual)
}

func Test_Cov15_BytePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.BytePtr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr both nil", actual)
}

func Test_Cov15_BytePtr_LeftNil(t *testing.T) {
	b := byte(1)
	actual := args.Map{"result": corecmp.BytePtr(nil, &b) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr left nil", actual)
}

func Test_Cov15_BytePtr_RightNil(t *testing.T) {
	b := byte(1)
	actual := args.Map{"result": corecmp.BytePtr(&b, nil) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr right nil", actual)
}

func Test_Cov15_BytePtr_Equal(t *testing.T) {
	a, b := byte(5), byte(5)
	actual := args.Map{"result": corecmp.BytePtr(&a, &b) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr equal", actual)
}

// ── Integer / IntegerPtr ──

func Test_Cov15_Integer_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer equal", actual)
}

func Test_Cov15_Integer_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer left less", actual)
}

func Test_Cov15_Integer_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer(5, 3) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer left greater", actual)
}

func Test_Cov15_IntegerPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr both nil", actual)
}

func Test_Cov15_IntegerPtr_LeftNil(t *testing.T) {
	v := 5
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr left nil", actual)
}

func Test_Cov15_IntegerPtr_RightNil(t *testing.T) {
	v := 5
	actual := args.Map{"result": corecmp.IntegerPtr(&v, nil) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr right nil", actual)
}

// ── Integer8 / Integer8Ptr ──

func Test_Cov15_Integer8_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 equal", actual)
}

func Test_Cov15_Integer8_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 left less", actual)
}

func Test_Cov15_Integer8_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8(5, 3) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 left greater", actual)
}

func Test_Cov15_Integer8Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr both nil", actual)
}

func Test_Cov15_Integer8Ptr_OneNil(t *testing.T) {
	v := int8(5)
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr one nil", actual)
}

// ── Integer16 / Integer16Ptr ──

func Test_Cov15_Integer16_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 equal", actual)
}

func Test_Cov15_Integer16_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 left less", actual)
}

func Test_Cov15_Integer16_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16(5, 3) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 left greater", actual)
}

func Test_Cov15_Integer16Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr both nil", actual)
}

func Test_Cov15_Integer16Ptr_OneNil(t *testing.T) {
	v := int16(5)
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr one nil", actual)
}

// ── Integer32 / Integer32Ptr ──

func Test_Cov15_Integer32_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 equal", actual)
}

func Test_Cov15_Integer32_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 left less", actual)
}

func Test_Cov15_Integer32_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32(5, 3) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 left greater", actual)
}

func Test_Cov15_Integer32Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr both nil", actual)
}

func Test_Cov15_Integer32Ptr_OneNil(t *testing.T) {
	v := int32(5)
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr one nil", actual)
}

// ── Integer64 / Integer64Ptr ──

func Test_Cov15_Integer64_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(5, 5) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 equal", actual)
}

func Test_Cov15_Integer64_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(3, 5) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 left less", actual)
}

func Test_Cov15_Integer64_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64(5, 3) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 left greater", actual)
}

func Test_Cov15_Integer64Ptr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr both nil", actual)
}

func Test_Cov15_Integer64Ptr_OneNil(t *testing.T) {
	v := int64(5)
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &v) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr one nil", actual)
}

// ── Time / TimePtr ──

func Test_Cov15_Time_Equal(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.Time(now, now) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time equal", actual)
}

func Test_Cov15_Time_LeftLess(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)
	actual := args.Map{"result": corecmp.Time(now, later) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time left less", actual)
}

func Test_Cov15_Time_LeftGreater(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-time.Hour)
	actual := args.Map{"result": corecmp.Time(now, earlier) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time left greater", actual)
}

func Test_Cov15_TimePtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.TimePtr(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr both nil", actual)
}

func Test_Cov15_TimePtr_OneNil(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.TimePtr(nil, &now) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr one nil", actual)
}

func Test_Cov15_TimePtr_Equal(t *testing.T) {
	now := time.Now()
	actual := args.Map{"result": corecmp.TimePtr(&now, &now) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr equal", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr / IsStringsEqualWithoutOrder ──

func Test_Cov15_IsStringsEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual both nil", actual)
}

func Test_Cov15_IsStringsEqual_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual one nil", actual)
}

func Test_Cov15_IsStringsEqual_DifferentLength(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual different length", actual)
}

func Test_Cov15_IsStringsEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual equal", actual)
}

func Test_Cov15_IsStringsEqual_NotEqual(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual not equal", actual)
}

func Test_Cov15_IsStringsEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr both nil", actual)
}

func Test_Cov15_IsStringsEqualPtr_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr one nil", actual)
}

func Test_Cov15_IsStringsEqualPtr_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr diff len", actual)
}

func Test_Cov15_IsStringsEqualWithoutOrder_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder equal", actual)
}

func Test_Cov15_IsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder both nil", actual)
}

func Test_Cov15_IsStringsEqualWithoutOrder_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder one nil", actual)
}

func Test_Cov15_IsStringsEqualWithoutOrder_DiffLen(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder diff len", actual)
}

func Test_Cov15_IsStringsEqualWithoutOrder_NotEqual(t *testing.T) {
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder not equal", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_Cov15_IsIntegersEqual_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual both nil", actual)
}

func Test_Cov15_IsIntegersEqual_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual one nil", actual)
}

func Test_Cov15_IsIntegersEqual_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual equal", actual)
}

func Test_Cov15_IsIntegersEqualPtr_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr both nil", actual)
}

func Test_Cov15_IsIntegersEqualPtr_OneNil(t *testing.T) {
	a := []int{1}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&a, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr one nil", actual)
}

func Test_Cov15_IsIntegersEqualPtr_DiffLen(t *testing.T) {
	a := []int{1}
	b := []int{1, 2}
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&a, &b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr diff len", actual)
}

// ── VersionSliceByte ──

func Test_Cov15_VersionSliceByte_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte equal", actual)
}

func Test_Cov15_VersionSliceByte_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left less", actual)
}

func Test_Cov15_VersionSliceByte_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte left greater", actual)
}

func Test_Cov15_VersionSliceByte_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte both nil", actual)
}

func Test_Cov15_VersionSliceByte_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1}) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte one nil", actual)
}

func Test_Cov15_VersionSliceByte_ShorterLeft(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte shorter left", actual)
}

func Test_Cov15_VersionSliceByte_LongerLeft(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte longer left", actual)
}

// ── VersionSliceInteger ──

func Test_Cov15_VersionSliceInteger_Equal(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger equal", actual)
}

func Test_Cov15_VersionSliceInteger_LeftLess(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left less", actual)
}

func Test_Cov15_VersionSliceInteger_LeftGreater(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger left greater", actual)
}

func Test_Cov15_VersionSliceInteger_BothNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil) == corecomparator.Equal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger both nil", actual)
}

func Test_Cov15_VersionSliceInteger_OneNil(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1}) == corecomparator.NotEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger one nil", actual)
}

func Test_Cov15_VersionSliceInteger_ShorterLeft(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) == corecomparator.LeftLess}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger shorter left", actual)
}

func Test_Cov15_VersionSliceInteger_LongerLeft(t *testing.T) {
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1}) == corecomparator.LeftGreater}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger longer left", actual)
}
