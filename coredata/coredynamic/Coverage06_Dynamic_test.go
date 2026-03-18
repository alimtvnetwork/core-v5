package coredynamic

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestDynamic_Constructors(t *testing.T) {
	d := NewDynamic("hello", true)
	if d.IsInvalid() { t.Fatal("expected valid") }
	if d.IsNull() { t.Fatal("expected not null") }
	dp := NewDynamicPtr("hello", true)
	if dp == nil { t.Fatal("expected non-nil") }
	dv := NewDynamicValid("hello")
	if dv.IsInvalid() { t.Fatal("expected valid") }
	id := InvalidDynamic()
	if !id.IsInvalid() { t.Fatal("expected invalid") }
	idp := InvalidDynamicPtr()
	if idp == nil { t.Fatal("expected non-nil") }
}

func TestDynamic_Clone(t *testing.T) {
	d := NewDynamic("hello", true)
	c := d.Clone()
	if c.IsInvalid() { t.Fatal("expected valid") }
	_ = d.NonPtr()
	_ = d.Ptr()
	dp := NewDynamicPtr("hello", true)
	cp := dp.ClonePtr()
	if cp == nil { t.Fatal("expected non-nil") }
	var nilD *Dynamic
	if nilD.ClonePtr() != nil { t.Fatal("expected nil") }
}

func TestDynamic_Getters(t *testing.T) {
	d := NewDynamic("hello", true)
	if d.Data() != "hello" { t.Fatal("expected hello") }
	if d.Value() != "hello" { t.Fatal("expected hello") }
	_ = d.String()
	_ = d.StructString()
	_ = d.StructStringPtr()
	// Call twice for caching
	_ = d.StructStringPtr()
}

func TestDynamic_TypeChecks(t *testing.T) {
	d := NewDynamic("hello", true)
	if !d.IsValid() { t.Fatal("expected valid") }
	if !d.IsStringType() { t.Fatal("expected string") }
	if d.IsPointer() { t.Fatal("expected non-pointer") }
	if !d.IsValueType() { t.Fatal("expected value type") }
	if d.IsNull() { t.Fatal("expected not null") }
	if d.IsStruct() { t.Fatal("expected not struct") }
	if d.IsFunc() { t.Fatal("expected not func") }
	if d.IsSliceOrArray() { t.Fatal("expected not slice") }
	if d.IsSliceOrArrayOrMap() { t.Fatal("expected not map") }
	if d.IsMap() { t.Fatal("expected not map") }
	if d.IsPrimitive() != true { t.Fatal("expected primitive") }
	if d.IsNumber() { t.Fatal("expected not number") }
}

func TestDynamic_TypeChecks_Struct(t *testing.T) {
	type testStruct struct{ Name string }
	d := NewDynamic(testStruct{Name: "x"}, true)
	if !d.IsStruct() { t.Fatal("expected struct") }
}

func TestDynamic_TypeChecks_Slice(t *testing.T) {
	d := NewDynamic([]string{"a"}, true)
	if !d.IsSliceOrArray() { t.Fatal("expected slice") }
	if !d.IsSliceOrArrayOrMap() { t.Fatal("expected slice/array/map") }
}

func TestDynamic_TypeChecks_Map(t *testing.T) {
	d := NewDynamic(map[string]int{"a": 1}, true)
	if !d.IsMap() { t.Fatal("expected map") }
}

func TestDynamic_TypeChecks_Number(t *testing.T) {
	d := NewDynamic(42, true)
	if !d.IsNumber() { t.Fatal("expected number") }
}

func TestDynamic_TypeChecks_Pointer(t *testing.T) {
	s := "hello"
	d := NewDynamic(&s, true)
	if !d.IsPointer() { t.Fatal("expected pointer") }
}

func TestDynamic_IsStructStringNullOrEmpty(t *testing.T) {
	d := NewDynamic(nil, false)
	if !d.IsStructStringNullOrEmpty() { t.Fatal("expected true") }
	d2 := NewDynamic("hello", true)
	if d2.IsStructStringNullOrEmpty() { t.Fatal("expected false") }
}

func TestDynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d := NewDynamic(nil, false)
	if !d.IsStructStringNullOrEmptyOrWhitespace() { t.Fatal("expected true") }
}

func TestDynamic_ValueExtraction(t *testing.T) {
	di := NewDynamic(42, true)
	if di.ValueInt() != 42 { t.Fatal("expected 42") }
	ds := NewDynamic("hello", true)
	if ds.ValueString() != "hello" { t.Fatal("expected hello") }
	db := NewDynamic(true, true)
	if !db.ValueBool() { t.Fatal("expected true") }
	d64 := NewDynamic(int64(99), true)
	if d64.ValueInt64() != 99 { t.Fatal("expected 99") }
	du := NewDynamic(uint(5), true)
	if du.ValueUInt() != 5 { t.Fatal("expected 5") }
	dstrs := NewDynamic([]string{"a"}, true)
	if len(dstrs.ValueStrings()) != 1 { t.Fatal("expected 1") }
}

func TestDynamic_ValueExtraction_Wrong(t *testing.T) {
	d := NewDynamic("hello", true)
	if d.ValueInt() != -1 { t.Fatal("expected -1") }
	if d.ValueBool() { t.Fatal("expected false") }
	if d.ValueInt64() != -1 { t.Fatal("expected -1") }
	if d.ValueUInt() != 0 { t.Fatal("expected 0") }
	if d.ValueStrings() != nil { t.Fatal("expected nil") }
}

func TestDynamic_ValueString_Nil(t *testing.T) {
	var d *Dynamic
	if d.ValueString() != "" { t.Fatal("expected empty") }
	d2 := NewDynamicPtr(nil, false)
	if d2.ValueString() != "" { t.Fatal("expected empty") }
	d3 := NewDynamicPtr(42, true)
	s := d3.ValueString()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestDynamic_ValueNullErr(t *testing.T) {
	var d *Dynamic
	if d.ValueNullErr() == nil { t.Fatal("expected error") }
	d2 := NewDynamicPtr(nil, false)
	if d2.ValueNullErr() == nil { t.Fatal("expected error") }
	d3 := NewDynamicPtr("x", true)
	if d3.ValueNullErr() != nil { t.Fatal("expected nil") }
}

func TestDynamic_Bytes(t *testing.T) {
	var d *Dynamic
	b, ok := d.Bytes()
	if b != nil || ok { t.Fatal("expected nil, false") }
	d2 := NewDynamicPtr([]byte("hello"), true)
	b2, ok2 := d2.Bytes()
	if !ok2 || len(b2) == 0 { t.Fatal("expected bytes") }
	d3 := NewDynamicPtr("hello", true)
	_, ok3 := d3.Bytes()
	if ok3 { t.Fatal("expected false") }
}

func TestDynamic_IntDefault(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	val, ok := d.IntDefault(99)
	if ok || val != 99 { t.Fatal("expected 99, false") }
	d2 := NewDynamicPtr("42", true)
	val2, ok2 := d2.IntDefault(0)
	if !ok2 || val2 != 42 { t.Fatal("expected 42, true") }
	d3 := NewDynamicPtr("abc", true)
	val3, ok3 := d3.IntDefault(10)
	if ok3 || val3 != 10 { t.Fatal("expected 10, false") }
}

func TestDynamic_Float64(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	_, err := d.Float64()
	if err == nil { t.Fatal("expected error") }
	d2 := NewDynamicPtr("3.14", true)
	val, err2 := d2.Float64()
	if err2 != nil || val != 3.14 { t.Fatal("unexpected") }
	d3 := NewDynamicPtr("abc", true)
	_, err3 := d3.Float64()
	if err3 == nil { t.Fatal("expected error") }
}

func TestDynamic_Length(t *testing.T) {
	d := NewDynamic([]int{1, 2, 3}, true)
	if d.Length() != 3 { t.Fatal("expected 3") }
	d2 := NewDynamic(nil, false)
	if d2.Length() != 0 { t.Fatal("expected 0") }
}

// === DynamicReflect ===

func TestDynamic_ReflectValue(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	rv := d.ReflectValue()
	if rv == nil { t.Fatal("expected non-nil") }
	// Call again for caching
	rv2 := d.ReflectValue()
	if rv != rv2 { t.Fatal("expected same ptr") }
}

func TestDynamic_ReflectKind(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	if d.ReflectKind() != reflect.String { t.Fatal("expected string") }
}

func TestDynamic_ReflectType(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	rt := d.ReflectType()
	if rt == nil { t.Fatal("expected non-nil") }
	// Call again for caching
	rt2 := d.ReflectType()
	if rt != rt2 { t.Fatal("expected same") }
}

func TestDynamic_ReflectTypeName(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	s := d.ReflectTypeName()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestDynamic_IsReflectTypeOf(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	if !d.IsReflectTypeOf(reflect.TypeOf("")) { t.Fatal("expected true") }
}

func TestDynamic_IsReflectKind(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	if !d.IsReflectKind(reflect.String) { t.Fatal("expected true") }
}

func TestDynamic_ItemUsingIndex(t *testing.T) {
	d := NewDynamicPtr([]string{"a", "b"}, true)
	if d.ItemUsingIndex(0) != "a" { t.Fatal("expected a") }
	rv := d.ItemReflectValueUsingIndex(0)
	if !rv.IsValid() { t.Fatal("expected valid") }
}

func TestDynamic_ItemUsingKey(t *testing.T) {
	d := NewDynamicPtr(map[string]int{"a": 1}, true)
	if d.ItemUsingKey("a") != 1 { t.Fatal("expected 1") }
	rv := d.ItemReflectValueUsingKey("a")
	if !rv.IsValid() { t.Fatal("expected valid") }
}

func TestDynamic_ReflectSetTo(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	var s string
	err := d.ReflectSetTo(&s)
	if err != nil { t.Fatal("unexpected error") }
	var nilD *Dynamic
	err2 := nilD.ReflectSetTo(&s)
	if err2 == nil { t.Fatal("expected error") }
}

func TestDynamic_MapToKeyVal(t *testing.T) {
	d := NewDynamicPtr(map[string]int{"a": 1}, true)
	kvc, err := d.MapToKeyVal()
	if err != nil { t.Fatal("unexpected error") }
	if kvc.Length() != 1 { t.Fatal("expected 1") }
}

func TestDynamic_Loop(t *testing.T) {
	d := NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return false
	})
	if count != 3 { t.Fatal("expected 3") }
}

func TestDynamic_Loop_Break(t *testing.T) {
	d := NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return index == 0
	})
	if count != 1 { t.Fatal("expected 1") }
}

func TestDynamic_Loop_Invalid(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	called := d.Loop(func(index int, item any) bool { return false })
	if called { t.Fatal("expected false") }
}

func TestDynamic_FilterAsDynamicCollection(t *testing.T) {
	d := NewDynamicPtr([]int{1, 2, 3, 4}, true)
	dc := d.FilterAsDynamicCollection(func(index int, item Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})
	if dc.Length() != 2 { t.Fatal("expected 2") }
}

func TestDynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := NewDynamicPtr([]int{1, 2, 3}, true)
	dc := d.FilterAsDynamicCollection(func(index int, item Dynamic) (bool, bool) {
		return true, index == 1
	})
	if dc.Length() != 2 { t.Fatal("expected 2") }
}

func TestDynamic_FilterAsDynamicCollection_Invalid(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	dc := d.FilterAsDynamicCollection(func(index int, item Dynamic) (bool, bool) {
		return true, false
	})
	if dc.Length() != 0 { t.Fatal("expected 0") }
}

func TestDynamic_LoopMap(t *testing.T) {
	d := NewDynamicPtr(map[string]int{"a": 1}, true)
	called := d.LoopMap(func(index int, key, value any) bool { return false })
	if !called { t.Fatal("expected true") }
}

func TestDynamic_LoopMap_Break(t *testing.T) {
	d := NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	d.LoopMap(func(index int, key, value any) bool { return true })
}

func TestDynamic_LoopMap_Invalid(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	called := d.LoopMap(func(index int, key, value any) bool { return false })
	if called { t.Fatal("expected false") }
}

// === DynamicJson ===

func TestDynamic_JsonString_Method(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	s, err := d.JsonString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestDynamic_JsonStringMust(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	s := d.JsonStringMust()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestDynamic_JsonBytes(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	b, err := d.JsonBytes()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestDynamic_JsonBytesPtr(t *testing.T) {
	d := NewDynamicPtr(nil, false)
	b, err := d.JsonBytesPtr()
	if err != nil || len(b) != 0 { t.Fatal("unexpected") }
	d2 := NewDynamicPtr("hello", true)
	b2, err2 := d2.JsonBytesPtr()
	if err2 != nil || len(b2) == 0 { t.Fatal("unexpected") }
}

func TestDynamic_ValueMarshal(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	b, err := d.ValueMarshal()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
	var nilD *Dynamic
	_, err2 := nilD.ValueMarshal()
	if err2 == nil { t.Fatal("expected error") }
}

func TestDynamic_JsonPayloadMust(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	b := d.JsonPayloadMust()
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func TestDynamic_MarshalUnmarshalJSON(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	b, err := d.MarshalJSON()
	if err != nil { t.Fatal("unexpected") }
	d2 := NewDynamicPtr("", true)
	_ = d2.UnmarshalJSON(b)
}

func TestDynamic_UnmarshalJSON_Nil(t *testing.T) {
	var d *Dynamic
	err := d.UnmarshalJSON([]byte(`"x"`))
	if err == nil { t.Fatal("expected error") }
}

func TestDynamic_Deserialize(t *testing.T) {
	d := NewDynamicPtr("", true)
	_, err := d.Deserialize([]byte(`"hello"`))
	if err != nil { t.Fatal("unexpected error") }
	var nilD *Dynamic
	_, err2 := nilD.Deserialize([]byte(`"x"`))
	if err2 == nil { t.Fatal("expected error") }
}

func TestDynamic_JsonModel(t *testing.T) {
	d := NewDynamic("hello", true)
	_ = d.JsonModel()
	_ = d.JsonModelAny()
	_ = d.Json()
	_ = d.JsonPtr()
}

func TestDynamic_ParseInjectUsingJson(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	jr := corejson.NewResult.Any("world")
	_, _ = d.ParseInjectUsingJson(&jr)
}

func TestDynamic_ParseInjectUsingJsonMust(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	jr := corejson.NewResult.Any("world")
	_ = d.ParseInjectUsingJsonMust(&jr)
}

func TestDynamic_JsonParseSelfInject(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	jr := corejson.NewResult.Any("world")
	_ = d.JsonParseSelfInject(&jr)
}

func TestDynamic_ConvertUsingFunc(t *testing.T) {
	d := NewDynamicPtr("hello", true)
	converter := func(input any, expected reflect.Type) *SimpleResult {
		return NewSimpleResultValid(input)
	}
	r := d.ConvertUsingFunc(converter, reflect.TypeOf(""))
	if r == nil { t.Fatal("expected non-nil") }
}
