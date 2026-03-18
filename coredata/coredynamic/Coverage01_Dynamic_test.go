package coredynamic

import (
	"reflect"
	"testing"
)

func TestDynamic_Constructors(t *testing.T) {
	d := InvalidDynamic()
	if d.IsValid() { t.Fatal("expected invalid") }
	dp := InvalidDynamicPtr()
	if dp.IsValid() { t.Fatal("expected invalid") }
	dv := NewDynamicValid("hello")
	if !dv.IsValid() || dv.Data() != "hello" { t.Fatal("unexpected") }
	d2 := NewDynamic(42, true)
	if d2.Value() != 42 { t.Fatal("unexpected") }
	d3 := NewDynamicPtr("test", true)
	if d3.Value() != "test" { t.Fatal("unexpected") }
}

func TestDynamic_Clone(t *testing.T) {
	d := NewDynamicValid("hello")
	c := d.Clone()
	if c.Value() != "hello" { t.Fatal("unexpected") }
	cp := d.ClonePtr()
	if cp.Value() != "hello" { t.Fatal("unexpected") }
	var nilD *Dynamic
	if nilD.ClonePtr() != nil { t.Fatal("expected nil") }
	_ = d.NonPtr()
	_ = d.Ptr()
}

func TestDynamic_Getters(t *testing.T) {
	d := NewDynamicValid("hello")
	if d.IsNull() { t.Fatal("expected not null") }
	if d.IsInvalid() { t.Fatal("expected valid") }
	if d.Length() != 0 { /* strings have 0 length in reflect */ }
	_ = d.String()
	_ = d.StructString()
	_ = d.StructStringPtr()
	if !d.IsStringType() { t.Fatal("expected string") }
	if d.IsPointer() { t.Fatal("expected not pointer") }
	if !d.IsValueType() { t.Fatal("expected value type") }
	if d.IsNumber() { t.Fatal("expected not number") }
	if !d.IsPrimitive() { t.Fatal("expected primitive") }
	if d.IsStruct() || d.IsFunc() || d.IsSliceOrArray() || d.IsMap() { t.Fatal("unexpected") }
	if d.IsSliceOrArrayOrMap() { t.Fatal("unexpected") }
	_ = d.IsStructStringNullOrEmpty()
	_ = d.IsStructStringNullOrEmptyOrWhitespace()
}

func TestDynamic_ValueExtraction(t *testing.T) {
	d := NewDynamicValid(42)
	if d.ValueInt() != 42 { t.Fatal("expected 42") }
	if d.ValueBool() { t.Fatal("expected false") }
	d2 := NewDynamicValid(true)
	if !d2.ValueBool() { t.Fatal("expected true") }
	d3 := NewDynamicValid(int64(99))
	if d3.ValueInt64() != 99 { t.Fatal("expected 99") }
	d4 := NewDynamicValid(uint(5))
	if d4.ValueUInt() != 5 { t.Fatal("expected 5") }
	d5 := NewDynamicValid([]string{"a", "b"})
	if len(d5.ValueStrings()) != 2 { t.Fatal("expected 2") }
	d6 := NewDynamicValid("hello")
	if d6.ValueString() != "hello" { t.Fatal("unexpected") }
	var nilD *Dynamic
	if nilD.ValueNullErr() == nil { t.Fatal("expected error") }
	nullD := NewDynamicPtr(nil, true)
	if nullD.ValueNullErr() == nil { t.Fatal("expected error") }
	if nullD.ValueString() != "" { t.Fatal("expected empty") }
	_, ok := nullD.Bytes()
	if ok { t.Fatal("expected false") }
	d7 := NewDynamicPtr([]byte("test"), true)
	b, ok := d7.Bytes()
	if !ok || len(b) != 4 { t.Fatal("unexpected") }
}

func TestDynamic_IntDefault(t *testing.T) {
	d := NewDynamicValid("42")
	v, ok := d.IntDefault(0)
	if !ok || v != 42 { t.Fatal("unexpected") }
	d2 := NewDynamicPtr(nil, true)
	v2, ok2 := d2.IntDefault(99)
	if ok2 || v2 != 99 { t.Fatal("unexpected") }
}

func TestDynamic_Float64(t *testing.T) {
	d := NewDynamicValid("3.14")
	v, err := d.Float64()
	if err != nil || v == 0 { t.Fatal("unexpected") }
	d2 := NewDynamicPtr(nil, true)
	_, err2 := d2.Float64()
	if err2 == nil { t.Fatal("expected error") }
}

func TestDynamic_Reflect(t *testing.T) {
	d := NewDynamicValid("hello")
	rv := d.ReflectValue()
	if rv.Kind() != reflect.String { t.Fatal("expected string") }
	if d.ReflectKind() != reflect.String { t.Fatal("expected string") }
	_ = d.ReflectTypeName()
	_ = d.ReflectType()
	if !d.IsReflectTypeOf(reflect.TypeOf("")) { t.Fatal("expected true") }
	if !d.IsReflectKind(reflect.String) { t.Fatal("expected true") }
}

func TestDynamic_Loop(t *testing.T) {
	d := NewDynamicValid([]int{1, 2, 3})
	called := d.Loop(func(i int, item any) bool { return false })
	if !called { t.Fatal("expected called") }
	invalid := NewDynamicPtr(nil, false)
	called2 := invalid.Loop(func(i int, item any) bool { return false })
	if called2 { t.Fatal("expected not called") }
}

func TestDynamic_LoopMap(t *testing.T) {
	d := NewDynamicValid(map[string]int{"a": 1})
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	if !called { t.Fatal("expected called") }
}

func TestDynamic_Json(t *testing.T) {
	d := NewDynamicValid("hello")
	_ = d.JsonModel()
	_ = d.JsonModelAny()
	_ = d.Json()
	_ = d.JsonPtr()
	_, _ = d.MarshalJSON()
	_, _ = d.ValueMarshal()
	_ = d.JsonPayloadMust()
	_, _ = d.JsonBytesPtr()
	_, _ = d.JsonBytes()
	_, _ = d.JsonString()
	_ = d.JsonStringMust()
	var nilD *Dynamic
	_, err := nilD.ValueMarshal()
	if err == nil { t.Fatal("expected error") }
	_, err2 := nilD.Deserialize(nil)
	if err2 == nil { t.Fatal("expected error") }
}

func TestDynamic_ReflectSetTo(t *testing.T) {
	d := NewDynamicValid("hello")
	var s string
	_ = d.ReflectSetTo(&s)
	var nilD *Dynamic
	if nilD.ReflectSetTo(&s) == nil { t.Fatal("expected error") }
}
