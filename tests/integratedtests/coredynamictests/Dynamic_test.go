package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// =============================================================================
// Dynamic — Constructors
// =============================================================================

func Test_Dynamic_NewDynamicValid(t *testing.T) {
	convey.Convey("NewDynamicValid creates valid Dynamic", t, func() {
		d := coredynamic.NewDynamicValid("hello")
		convey.So(d.IsValid(), should.BeTrue)
		convey.So(d.Value(), should.Equal, "hello")
	})
}

func Test_Dynamic_NewDynamic_Invalid(t *testing.T) {
	convey.Convey("NewDynamic with isValid=false creates invalid Dynamic", t, func() {
		d := coredynamic.NewDynamic(nil, false)
		convey.So(d.IsValid(), should.BeFalse)
		convey.So(d.IsInvalid(), should.BeTrue)
	})
}

func Test_Dynamic_InvalidDynamic(t *testing.T) {
	convey.Convey("InvalidDynamic creates invalid nil Dynamic", t, func() {
		d := coredynamic.InvalidDynamic()
		convey.So(d.IsValid(), should.BeFalse)
		convey.So(d.IsNull(), should.BeTrue)
	})
}

func Test_Dynamic_InvalidDynamicPtr(t *testing.T) {
	convey.Convey("InvalidDynamicPtr creates invalid nil Dynamic pointer", t, func() {
		d := coredynamic.InvalidDynamicPtr()
		convey.So(d, should.NotBeNil)
		convey.So(d.IsValid(), should.BeFalse)
		convey.So(d.IsNull(), should.BeTrue)
	})
}

func Test_Dynamic_NewDynamicPtr(t *testing.T) {
	convey.Convey("NewDynamicPtr creates pointer Dynamic", t, func() {
		d := coredynamic.NewDynamicPtr(42, true)
		convey.So(d, should.NotBeNil)
		convey.So(d.IsValid(), should.BeTrue)
		convey.So(d.Value(), should.Equal, 42)
	})
}

// =============================================================================
// Dynamic — Clone
// =============================================================================

func Test_Dynamic_Clone(t *testing.T) {
	convey.Convey("Dynamic.Clone creates independent copy", t, func() {
		original := coredynamic.NewDynamicValid("data")
		cloned := original.Clone()
		convey.So(cloned.Value(), should.Equal, "data")
		convey.So(cloned.IsValid(), should.BeTrue)
	})
}

func Test_Dynamic_ClonePtr(t *testing.T) {
	convey.Convey("Dynamic.ClonePtr creates independent pointer copy", t, func() {
		original := coredynamic.NewDynamicPtr("data", true)
		cloned := original.ClonePtr()
		convey.So(cloned, should.NotBeNil)
		convey.So(cloned.Value(), should.Equal, "data")
	})
}

func Test_Dynamic_ClonePtr_NilReceiver(t *testing.T) {
	convey.Convey("Dynamic.ClonePtr returns nil on nil receiver", t, func() {
		var d *coredynamic.Dynamic
		convey.So(d.ClonePtr(), should.BeNil)
	})
}

func Test_Dynamic_NonPtr(t *testing.T) {
	convey.Convey("Dynamic.NonPtr returns value copy", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.NonPtr().Value(), should.Equal, "x")
	})
}

func Test_Dynamic_Ptr(t *testing.T) {
	convey.Convey("Dynamic.Ptr returns pointer to self", t, func() {
		d := coredynamic.NewDynamicPtr("x", true)
		convey.So(d.Ptr(), should.Equal, d)
	})
}

// =============================================================================
// Dynamic — Getters & Type Checks
// =============================================================================

func Test_Dynamic_Data_And_Value(t *testing.T) {
	convey.Convey("Dynamic.Data and Value return same inner data", t, func() {
		d := coredynamic.NewDynamicValid(99)
		convey.So(d.Data(), should.Equal, 99)
		convey.So(d.Value(), should.Equal, d.Data())
	})
}

func Test_Dynamic_IsNull_True(t *testing.T) {
	convey.Convey("Dynamic.IsNull true for nil data", t, func() {
		d := coredynamic.NewDynamic(nil, true)
		convey.So(d.IsNull(), should.BeTrue)
	})
}

func Test_Dynamic_IsNull_False(t *testing.T) {
	convey.Convey("Dynamic.IsNull false for non-nil data", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.IsNull(), should.BeFalse)
	})
}

func Test_Dynamic_String(t *testing.T) {
	convey.Convey("Dynamic.String returns string representation", t, func() {
		d := coredynamic.NewDynamicValid("hello")
		convey.So(d.String(), should.NotBeEmpty)
	})
}

func Test_Dynamic_IsStringType_True(t *testing.T) {
	convey.Convey("Dynamic.IsStringType true for string data", t, func() {
		d := coredynamic.NewDynamicValid("text")
		convey.So(d.IsStringType(), should.BeTrue)
	})
}

func Test_Dynamic_IsStringType_False(t *testing.T) {
	convey.Convey("Dynamic.IsStringType false for non-string data", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.IsStringType(), should.BeFalse)
	})
}

func Test_Dynamic_IsNumber_True(t *testing.T) {
	convey.Convey("Dynamic.IsNumber true for int", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.IsNumber(), should.BeTrue)
	})
}

func Test_Dynamic_IsNumber_False(t *testing.T) {
	convey.Convey("Dynamic.IsNumber false for string", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.IsNumber(), should.BeFalse)
	})
}

func Test_Dynamic_IsPrimitive_True(t *testing.T) {
	convey.Convey("Dynamic.IsPrimitive true for int", t, func() {
		d := coredynamic.NewDynamicValid(10)
		convey.So(d.IsPrimitive(), should.BeTrue)
	})
}

func Test_Dynamic_IsStruct_True(t *testing.T) {
	convey.Convey("Dynamic.IsStruct true for struct", t, func() {
		type sample struct{ Name string }
		d := coredynamic.NewDynamicValid(sample{Name: "test"})
		convey.So(d.IsStruct(), should.BeTrue)
	})
}

func Test_Dynamic_IsStruct_False(t *testing.T) {
	convey.Convey("Dynamic.IsStruct false for int", t, func() {
		d := coredynamic.NewDynamicValid(5)
		convey.So(d.IsStruct(), should.BeFalse)
	})
}

func Test_Dynamic_IsFunc_True(t *testing.T) {
	convey.Convey("Dynamic.IsFunc true for function", t, func() {
		d := coredynamic.NewDynamicValid(func() {})
		convey.So(d.IsFunc(), should.BeTrue)
	})
}

func Test_Dynamic_IsSliceOrArray_True(t *testing.T) {
	convey.Convey("Dynamic.IsSliceOrArray true for slice", t, func() {
		d := coredynamic.NewDynamicValid([]int{1, 2, 3})
		convey.So(d.IsSliceOrArray(), should.BeTrue)
	})
}

func Test_Dynamic_IsSliceOrArray_False(t *testing.T) {
	convey.Convey("Dynamic.IsSliceOrArray false for string", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.IsSliceOrArray(), should.BeFalse)
	})
}

func Test_Dynamic_IsSliceOrArrayOrMap_Map(t *testing.T) {
	convey.Convey("Dynamic.IsSliceOrArrayOrMap true for map", t, func() {
		d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
		convey.So(d.IsSliceOrArrayOrMap(), should.BeTrue)
	})
}

func Test_Dynamic_IsMap_True(t *testing.T) {
	convey.Convey("Dynamic.IsMap true for map", t, func() {
		d := coredynamic.NewDynamicValid(map[string]int{"x": 1})
		convey.So(d.IsMap(), should.BeTrue)
	})
}

func Test_Dynamic_IsPointer_True(t *testing.T) {
	convey.Convey("Dynamic.IsPointer true for pointer data", t, func() {
		val := 42
		d := coredynamic.NewDynamicValid(&val)
		convey.So(d.IsPointer(), should.BeTrue)
	})
}

func Test_Dynamic_IsValueType_True(t *testing.T) {
	convey.Convey("Dynamic.IsValueType true for non-pointer data", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.IsValueType(), should.BeTrue)
	})
}

// =============================================================================
// Dynamic — Length
// =============================================================================

func Test_Dynamic_Length_Slice(t *testing.T) {
	convey.Convey("Dynamic.Length returns slice length", t, func() {
		d := coredynamic.NewDynamicValid([]int{1, 2, 3})
		convey.So(d.Length(), should.Equal, 3)
	})
}

func Test_Dynamic_Length_Nil(t *testing.T) {
	convey.Convey("Dynamic.Length returns 0 for nil data", t, func() {
		d := coredynamic.NewDynamic(nil, false)
		convey.So(d.Length(), should.Equal, 0)
	})
}

func Test_Dynamic_Length_Map(t *testing.T) {
	convey.Convey("Dynamic.Length returns map length", t, func() {
		d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
		convey.So(d.Length(), should.Equal, 2)
	})
}

// =============================================================================
// Dynamic — Value Extraction
// =============================================================================

func Test_Dynamic_ValueInt_Success(t *testing.T) {
	convey.Convey("Dynamic.ValueInt returns int value", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.ValueInt(), should.Equal, 42)
	})
}

func Test_Dynamic_ValueInt_NonInt(t *testing.T) {
	convey.Convey("Dynamic.ValueInt returns InvalidValue for non-int", t, func() {
		d := coredynamic.NewDynamicValid("not-int")
		convey.So(d.ValueInt(), should.Equal, -1)
	})
}

func Test_Dynamic_ValueBool_True(t *testing.T) {
	convey.Convey("Dynamic.ValueBool returns true", t, func() {
		d := coredynamic.NewDynamicValid(true)
		convey.So(d.ValueBool(), should.BeTrue)
	})
}

func Test_Dynamic_ValueBool_NonBool(t *testing.T) {
	convey.Convey("Dynamic.ValueBool returns false for non-bool", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.ValueBool(), should.BeFalse)
	})
}

func Test_Dynamic_ValueString_String(t *testing.T) {
	convey.Convey("Dynamic.ValueString returns string directly", t, func() {
		d := coredynamic.NewDynamicValid("hello")
		convey.So(d.ValueString(), should.Equal, "hello")
	})
}

func Test_Dynamic_ValueString_NonString(t *testing.T) {
	convey.Convey("Dynamic.ValueString formats non-string", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.ValueString(), should.NotBeEmpty)
	})
}

func Test_Dynamic_ValueString_Nil(t *testing.T) {
	convey.Convey("Dynamic.ValueString returns empty for nil data", t, func() {
		d := coredynamic.NewDynamic(nil, true)
		convey.So(d.ValueString(), should.BeEmpty)
	})
}

func Test_Dynamic_ValueStrings_Success(t *testing.T) {
	convey.Convey("Dynamic.ValueStrings returns []string", t, func() {
		d := coredynamic.NewDynamicValid([]string{"a", "b"})
		convey.So(d.ValueStrings(), should.Resemble, []string{"a", "b"})
	})
}

func Test_Dynamic_ValueStrings_NonSlice(t *testing.T) {
	convey.Convey("Dynamic.ValueStrings returns nil for non-[]string", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.ValueStrings(), should.BeNil)
	})
}

func Test_Dynamic_ValueUInt_Success(t *testing.T) {
	convey.Convey("Dynamic.ValueUInt returns uint value", t, func() {
		d := coredynamic.NewDynamicValid(uint(10))
		convey.So(d.ValueUInt(), should.Equal, uint(10))
	})
}

func Test_Dynamic_ValueInt64_Success(t *testing.T) {
	convey.Convey("Dynamic.ValueInt64 returns int64 value", t, func() {
		d := coredynamic.NewDynamicValid(int64(999))
		convey.So(d.ValueInt64(), should.Equal, int64(999))
	})
}

func Test_Dynamic_Bytes_Success(t *testing.T) {
	convey.Convey("Dynamic.Bytes returns []byte", t, func() {
		d := coredynamic.NewDynamicValid([]byte("raw"))
		raw, ok := d.Bytes()
		convey.So(ok, should.BeTrue)
		convey.So(string(raw), should.Equal, "raw")
	})
}

func Test_Dynamic_Bytes_NonBytes(t *testing.T) {
	convey.Convey("Dynamic.Bytes returns false for non-bytes", t, func() {
		d := coredynamic.NewDynamicValid("str")
		_, ok := d.Bytes()
		convey.So(ok, should.BeFalse)
	})
}

func Test_Dynamic_Bytes_NilReceiver(t *testing.T) {
	convey.Convey("Dynamic.Bytes returns nil,false on nil receiver", t, func() {
		var d *coredynamic.Dynamic
		raw, ok := d.Bytes()
		convey.So(raw, should.BeNil)
		convey.So(ok, should.BeFalse)
	})
}

func Test_Dynamic_IntDefault_Success(t *testing.T) {
	convey.Convey("Dynamic.IntDefault parses int from string repr", t, func() {
		d := coredynamic.NewDynamicValid(42)
		val, ok := d.IntDefault(0)
		convey.So(ok, should.BeTrue)
		convey.So(val, should.Equal, 42)
	})
}

func Test_Dynamic_IntDefault_Null(t *testing.T) {
	convey.Convey("Dynamic.IntDefault returns default on nil data", t, func() {
		d := coredynamic.NewDynamic(nil, true)
		val, ok := d.IntDefault(99)
		convey.So(ok, should.BeFalse)
		convey.So(val, should.Equal, 99)
	})
}

func Test_Dynamic_ValueNullErr_Nil(t *testing.T) {
	convey.Convey("Dynamic.ValueNullErr returns error on nil receiver", t, func() {
		var d *coredynamic.Dynamic
		convey.So(d.ValueNullErr(), should.NotBeNil)
	})
}

func Test_Dynamic_ValueNullErr_NullData(t *testing.T) {
	convey.Convey("Dynamic.ValueNullErr returns error on null data", t, func() {
		d := coredynamic.NewDynamic(nil, true)
		convey.So(d.ValueNullErr(), should.NotBeNil)
	})
}

func Test_Dynamic_ValueNullErr_Valid(t *testing.T) {
	convey.Convey("Dynamic.ValueNullErr returns nil for valid data", t, func() {
		d := coredynamic.NewDynamicValid("ok")
		convey.So(d.ValueNullErr(), should.BeNil)
	})
}

// =============================================================================
// Dynamic — Reflect
// =============================================================================

func Test_Dynamic_ReflectKind(t *testing.T) {
	convey.Convey("Dynamic.ReflectKind returns correct kind", t, func() {
		d := coredynamic.NewDynamicValid("text")
		convey.So(d.ReflectKind(), should.Equal, reflect.String)
	})
}

func Test_Dynamic_ReflectKind_Int(t *testing.T) {
	convey.Convey("Dynamic.ReflectKind returns Int for int", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.ReflectKind(), should.Equal, reflect.Int)
	})
}

func Test_Dynamic_IsReflectKind_True(t *testing.T) {
	convey.Convey("Dynamic.IsReflectKind matches correctly", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.IsReflectKind(reflect.String), should.BeTrue)
	})
}

func Test_Dynamic_IsReflectKind_False(t *testing.T) {
	convey.Convey("Dynamic.IsReflectKind returns false on mismatch", t, func() {
		d := coredynamic.NewDynamicValid("x")
		convey.So(d.IsReflectKind(reflect.Int), should.BeFalse)
	})
}

func Test_Dynamic_ReflectTypeName(t *testing.T) {
	convey.Convey("Dynamic.ReflectTypeName returns type string", t, func() {
		d := coredynamic.NewDynamicValid("text")
		convey.So(d.ReflectTypeName(), should.NotBeEmpty)
	})
}

func Test_Dynamic_ReflectType(t *testing.T) {
	convey.Convey("Dynamic.ReflectType returns reflect.Type", t, func() {
		d := coredynamic.NewDynamicValid(42)
		convey.So(d.ReflectType(), should.Equal, reflect.TypeOf(42))
	})
}

func Test_Dynamic_IsReflectTypeOf_True(t *testing.T) {
	convey.Convey("Dynamic.IsReflectTypeOf matches type", t, func() {
		d := coredynamic.NewDynamicValid("hello")
		convey.So(d.IsReflectTypeOf(reflect.TypeOf("")), should.BeTrue)
	})
}

func Test_Dynamic_ReflectValue(t *testing.T) {
	convey.Convey("Dynamic.ReflectValue returns cached reflect.Value", t, func() {
		d := coredynamic.NewDynamicPtr(42, true)
		rv1 := d.ReflectValue()
		rv2 := d.ReflectValue()
		convey.So(rv1, should.Equal, rv2) // same pointer
		convey.So(rv1.Int(), should.Equal, 42)
	})
}

// =============================================================================
// Dynamic — Loop
// =============================================================================

func Test_Dynamic_Loop_Slice(t *testing.T) {
	convey.Convey("Dynamic.Loop iterates slice items", t, func() {
		d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
		collected := make([]string, 0, 3)
		called := d.Loop(func(index int, item any) bool {
			collected = append(collected, item.(string))
			return false
		})
		convey.So(called, should.BeTrue)
		convey.So(collected, should.Resemble, []string{"a", "b", "c"})
	})
}

func Test_Dynamic_Loop_Invalid(t *testing.T) {
	convey.Convey("Dynamic.Loop returns false for invalid", t, func() {
		d := coredynamic.InvalidDynamicPtr()
		called := d.Loop(func(index int, item any) bool { return false })
		convey.So(called, should.BeFalse)
	})
}

func Test_Dynamic_Loop_Break(t *testing.T) {
	convey.Convey("Dynamic.Loop respects break", t, func() {
		d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4})
		count := 0
		d.Loop(func(index int, item any) bool {
			count++
			return index == 1
		})
		convey.So(count, should.Equal, 2)
	})
}

// =============================================================================
// Dynamic — ItemUsingIndex / ItemUsingKey
// =============================================================================

func Test_Dynamic_ItemUsingIndex(t *testing.T) {
	convey.Convey("Dynamic.ItemUsingIndex returns correct element", t, func() {
		d := coredynamic.NewDynamicValid([]string{"a", "b"})
		convey.So(d.ItemUsingIndex(0), should.Equal, "a")
		convey.So(d.ItemUsingIndex(1), should.Equal, "b")
	})
}

func Test_Dynamic_ItemUsingKey(t *testing.T) {
	convey.Convey("Dynamic.ItemUsingKey returns map value", t, func() {
		d := coredynamic.NewDynamicValid(map[string]int{"k": 42})
		convey.So(d.ItemUsingKey("k"), should.Equal, 42)
	})
}

// =============================================================================
// Dynamic — IsStructStringNullOrEmpty
// =============================================================================

func Test_Dynamic_IsStructStringNullOrEmpty_Null(t *testing.T) {
	convey.Convey("Dynamic.IsStructStringNullOrEmpty true on nil data", t, func() {
		d := coredynamic.NewDynamic(nil, true)
		convey.So(d.IsStructStringNullOrEmpty(), should.BeTrue)
	})
}

func Test_Dynamic_IsStructStringNullOrEmpty_NonEmpty(t *testing.T) {
	convey.Convey("Dynamic.IsStructStringNullOrEmpty false for non-empty", t, func() {
		d := coredynamic.NewDynamicValid("text")
		convey.So(d.IsStructStringNullOrEmpty(), should.BeFalse)
	})
}
