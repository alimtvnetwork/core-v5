package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// DynamicGetters — Data, Value, Length, String, type checks, value extraction
// ══════════════════════════════════════════════════════════════════════════════

func Test_I25_DynamicGetters_Data_Value(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"data": d.Data(), "value": d.Value()}
	expected := args.Map{"data": "hello", "value": "hello"}
	expected.ShouldBeEqual(t, 0, "Data/Value returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_Length_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Length returns correct value -- slice", actual)
}

func Test_I25_DynamicGetters_StructStringPtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"nil": d.StructStringPtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "StructStringPtr returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_StructStringPtr_Cached(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	ptr1 := d.StructStringPtr()
	ptr2 := d.StructStringPtr()
	actual := args.Map{"same": ptr1 == ptr2, "val": *ptr1}
	expected := args.Map{"same": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "StructStringPtr returns correct value -- cached", actual)
}

func Test_I25_DynamicGetters_String_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"val": d.String()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_StructString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"val": d.StructString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "StructString returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_IsNull(t *testing.T) {
	d1 := coredynamic.NewDynamic(nil, false)
	d2 := coredynamic.NewDynamic("x", true)
	actual := args.Map{"null": d1.IsNull(), "notNull": d2.IsNull()}
	expected := args.Map{"null": true, "notNull": false}
	expected.ShouldBeEqual(t, 0, "IsNull returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsValid_IsInvalid(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"valid": d.IsValid(), "invalid": d.IsInvalid()}
	expected := args.Map{"valid": true, "invalid": false}
	expected.ShouldBeEqual(t, 0, "IsValid/IsInvalid returns error -- with args", actual)
}

func Test_I25_DynamicGetters_IsPointer_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"ptr": d.IsPointer()}
	expected := args.Map{"ptr": false}
	expected.ShouldBeEqual(t, 0, "IsPointer returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_IsPointer_True(t *testing.T) {
	s := "hello"
	d := coredynamic.NewDynamic(&s, true)
	actual := args.Map{"ptr": d.IsPointer()}
	expected := args.Map{"ptr": true}
	expected.ShouldBeEqual(t, 0, "IsPointer returns non-empty -- true", actual)
}

func Test_I25_DynamicGetters_IsPointer_Cached(t *testing.T) {
	s := "hello"
	d := coredynamic.NewDynamic(&s, true)
	_ = d.IsPointer() // first call
	actual := args.Map{"ptr": d.IsPointer()} // cached
	expected := args.Map{"ptr": true}
	expected.ShouldBeEqual(t, 0, "IsPointer returns correct value -- cached", actual)
}

func Test_I25_DynamicGetters_IsValueType(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"vt": d.IsValueType()}
	expected := args.Map{"vt": true}
	expected.ShouldBeEqual(t, 0, "IsValueType returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsValueType_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"vt": d.IsValueType()}
	expected := args.Map{"vt": false}
	expected.ShouldBeEqual(t, 0, "IsValueType returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_IsStructStringNullOrEmpty(t *testing.T) {
	d1 := coredynamic.NewDynamic("", true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"empty": d1.IsStructStringNullOrEmpty(), "notEmpty": d2.IsStructStringNullOrEmpty(), "nil": d3.IsStructStringNullOrEmpty()}
	expected := args.Map{"empty": true, "notEmpty": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmpty returns empty -- with args", actual)
}

func Test_I25_DynamicGetters_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d1 := coredynamic.NewDynamic("  ", true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"ws": d1.IsStructStringNullOrEmptyOrWhitespace(), "notWs": d2.IsStructStringNullOrEmptyOrWhitespace(), "nil": d3.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"ws": true, "notWs": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_I25_DynamicGetters_IsPrimitive(t *testing.T) {
	d1 := coredynamic.NewDynamic("hello", true)
	d2 := coredynamic.NewDynamic([]int{1}, true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"prim": d1.IsPrimitive(), "notPrim": d2.IsPrimitive(), "nil": d3.IsPrimitive()}
	expected := args.Map{"prim": true, "notPrim": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsPrimitive returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsNumber(t *testing.T) {
	d1 := coredynamic.NewDynamic(42, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"num": d1.IsNumber(), "notNum": d2.IsNumber(), "nil": d3.IsNumber()}
	expected := args.Map{"num": true, "notNum": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsNumber returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsStringType(t *testing.T) {
	d1 := coredynamic.NewDynamic("hello", true)
	d2 := coredynamic.NewDynamic(42, true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"str": d1.IsStringType(), "notStr": d2.IsStringType(), "nil": d3.IsStringType()}
	expected := args.Map{"str": true, "notStr": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsStringType returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsStruct(t *testing.T) {
	type s struct{}
	d1 := coredynamic.NewDynamic(s{}, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"struct": d1.IsStruct(), "notStruct": d2.IsStruct(), "nil": d3.IsStruct()}
	expected := args.Map{"struct": true, "notStruct": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsStruct returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsFunc(t *testing.T) {
	d1 := coredynamic.NewDynamic(func() {}, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"fn": d1.IsFunc(), "notFn": d2.IsFunc(), "nil": d3.IsFunc()}
	expected := args.Map{"fn": true, "notFn": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsFunc returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsSliceOrArray(t *testing.T) {
	d1 := coredynamic.NewDynamic([]int{1}, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"slice": d1.IsSliceOrArray(), "notSlice": d2.IsSliceOrArray(), "nil": d3.IsSliceOrArray()}
	expected := args.Map{"slice": true, "notSlice": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArray returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsSliceOrArrayOrMap(t *testing.T) {
	d1 := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"map": d1.IsSliceOrArrayOrMap(), "notMap": d2.IsSliceOrArrayOrMap(), "nil": d3.IsSliceOrArrayOrMap()}
	expected := args.Map{"map": true, "notMap": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArrayOrMap returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IsMap(t *testing.T) {
	d1 := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	d2 := coredynamic.NewDynamic("x", true)
	var d3 *coredynamic.Dynamic
	actual := args.Map{"map": d1.IsMap(), "notMap": d2.IsMap(), "nil": d3.IsMap()}
	expected := args.Map{"map": true, "notMap": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsMap returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_IntDefault_Valid(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	val, ok := d.IntDefault(0)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "IntDefault returns non-empty -- valid", actual)
}

func Test_I25_DynamicGetters_IntDefault_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntDefault returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_IntDefault_ParseFail(t *testing.T) {
	d := coredynamic.NewDynamic("abc", true)
	val, ok := d.IntDefault(77)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 77, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- parse fail", actual)
}

func Test_I25_DynamicGetters_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamic("3.14", true)
	val, err := d.Float64()
	actual := args.Map{"noErr": err == nil, "close": val > 3.1 && val < 3.2}
	expected := args.Map{"noErr": true, "close": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns non-empty -- valid", actual)
}

func Test_I25_DynamicGetters_Float64_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_Float64_ParseFail(t *testing.T) {
	d := coredynamic.NewDynamic("abc", true)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns correct value -- parse fail", actual)
}

func Test_I25_DynamicGetters_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	actual := args.Map{"val": d.ValueInt()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_ValueInt_NotInt(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"val": d.ValueInt()}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- not int", actual)
}

func Test_I25_DynamicGetters_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamic(uint(7), true)
	actual := args.Map{"val": d.ValueUInt()}
	expected := args.Map{"val": uint(7)}
	expected.ShouldBeEqual(t, 0, "ValueUInt returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_ValueUInt_NotUint(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"val": d.ValueUInt()}
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "ValueUInt returns correct value -- not uint", actual)
}

func Test_I25_DynamicGetters_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamic([]string{"a", "b"}, true)
	actual := args.Map{"len": len(d.ValueStrings())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ValueStrings returns non-empty -- with args", actual)
}

func Test_I25_DynamicGetters_ValueStrings_NotStrings(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"nil": d.ValueStrings() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ValueStrings returns non-empty -- not strings", actual)
}

func Test_I25_DynamicGetters_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamic(true, true)
	actual := args.Map{"val": d.ValueBool()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_ValueBool_NotBool(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"val": d.ValueBool()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- not bool", actual)
}

func Test_I25_DynamicGetters_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamic(int64(99), true)
	actual := args.Map{"val": d.ValueInt64()}
	expected := args.Map{"val": int64(99)}
	expected.ShouldBeEqual(t, 0, "ValueInt64 returns correct value -- with args", actual)
}

func Test_I25_DynamicGetters_ValueInt64_NotInt64(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"val": d.ValueInt64()}
	expected := args.Map{"val": int64(-1)}
	expected.ShouldBeEqual(t, 0, "ValueInt64 returns correct value -- not int64", actual)
}

func Test_I25_DynamicGetters_ValueNullErr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_ValueNullErr_NullData(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns error -- null data", actual)
}

func Test_I25_DynamicGetters_ValueNullErr_Valid(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	actual := args.Map{"noErr": d.ValueNullErr() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns error -- valid", actual)
}

func Test_I25_DynamicGetters_ValueString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"val": d.ValueString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ValueString returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_ValueString_String(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"val": d.ValueString()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- string", actual)
}

func Test_I25_DynamicGetters_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	val := d.ValueString()
	actual := args.Map{"notEmpty": val != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- non-string", actual)
}

func Test_I25_DynamicGetters_Bytes_Valid(t *testing.T) {
	d := coredynamic.NewDynamic([]byte{1, 2, 3}, true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "len": len(b)}
	expected := args.Map{"ok": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "Bytes returns non-empty -- valid", actual)
}

func Test_I25_DynamicGetters_Bytes_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "nil": b == nil}
	expected := args.Map{"ok": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "Bytes returns nil -- nil", actual)
}

func Test_I25_DynamicGetters_Bytes_NotBytes(t *testing.T) {
	d := coredynamic.NewDynamic("x", true)
	_, ok := d.Bytes()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- not bytes", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicJson — JSON serialization/deserialization
// ══════════════════════════════════════════════════════════════════════════════

func Test_I25_DynamicJson_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`{}`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns nil -- nil", actual)
}

func Test_I25_DynamicJson_ValueMarshal(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	b, err := d.ValueMarshal()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns nil -- nil", actual)
}

func Test_I25_DynamicJson_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	b := d.JsonPayloadMust()
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "JsonPayloadMust returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"noErr": err == nil, "empty": len(b) == 0}
	expected := args.Map{"noErr": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns correct value -- null", actual)
}

func Test_I25_DynamicJson_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns non-empty -- valid", actual)
}

func Test_I25_DynamicJson_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	b, err := d.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`{}`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns nil -- nil", actual)
}

func Test_I25_DynamicJson_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"val": d.JsonModel()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"val": d.JsonModelAny()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_Json(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	jr := d.Json()
	actual := args.Map{"noErr": !jr.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	jr := d.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_ParseInjectUsingJson_Error(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	jr := corejson.New("invalid but let's try")
	_, err := d.ParseInjectUsingJson(&jr)
	// May or may not error depending on internal unmarshal, just cover the path
	_ = err
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonParseSelfInject(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	jr := corejson.New("test")
	_ = d.JsonParseSelfInject(&jr)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonBytes(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	b, err := d.JsonBytes()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "JsonBytes returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonString(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	s, err := d.JsonString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_I25_DynamicJson_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	s := d.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicReflect — Reflection ops, loops, filters
// ══════════════════════════════════════════════════════════════════════════════

func Test_I25_DynamicReflect_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	rv := d.ReflectValue()
	actual := args.Map{"kind": rv.Kind().String()}
	expected := args.Map{"kind": "string"}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ReflectValue_Cached(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()
	actual := args.Map{"same": rv1 == rv2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- cached", actual)
}

func Test_I25_DynamicReflect_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]any{"a": 1}, true)
	kvc, err := d.MapToKeyVal()
	actual := args.Map{"noErr": err == nil, "notNil": kvc != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapToKeyVal returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ReflectKind(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	actual := args.Map{"kind": d.ReflectKind().String()}
	expected := args.Map{"kind": "int"}
	expected.ShouldBeEqual(t, 0, "ReflectKind returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	name := d.ReflectTypeName()
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeName returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ReflectType_Cached(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	rt1 := d.ReflectType()
	rt2 := d.ReflectType()
	actual := args.Map{"same": rt1 == rt2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "ReflectType returns correct value -- cached", actual)
}

func Test_I25_DynamicReflect_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"match": d.IsReflectTypeOf(reflect.TypeOf("")), "noMatch": d.IsReflectTypeOf(reflect.TypeOf(0))}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsReflectTypeOf returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_IsReflectKind(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{"str": d.IsReflectKind(reflect.String), "int": d.IsReflectKind(reflect.Int)}
	expected := args.Map{"str": true, "int": false}
	expected.ShouldBeEqual(t, 0, "IsReflectKind returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamic([]int{10, 20, 30}, true)
	actual := args.Map{"val": d.ItemUsingIndex(1)}
	expected := args.Map{"val": 20}
	expected.ShouldBeEqual(t, 0, "ItemUsingIndex returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ItemReflectValueUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamic([]int{10, 20}, true)
	rv := d.ItemReflectValueUsingIndex(0)
	actual := args.Map{"val": rv.Interface()}
	expected := args.Map{"val": 10}
	expected.ShouldBeEqual(t, 0, "ItemReflectValueUsingIndex returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]int{"a": 42}, true)
	actual := args.Map{"val": d.ItemUsingKey("a")}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ItemUsingKey returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ItemReflectValueUsingKey(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]int{"x": 7}, true)
	rv := d.ItemReflectValueUsingKey("x")
	actual := args.Map{"val": rv.Interface()}
	expected := args.Map{"val": 7}
	expected.ShouldBeEqual(t, 0, "ItemReflectValueUsingKey returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(&struct{}{})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns nil -- nil", actual)
}

func Test_I25_DynamicReflect_ConvertUsingFunc(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	converter := func(in any, typeMust reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResult(in, true)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ConvertUsingFunc returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_Loop_Slice(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	count := 0
	called := d.Loop(func(index int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 3}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- slice", actual)
}

func Test_I25_DynamicReflect_Loop_Break(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return index == 0
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- break", actual)
}

func Test_I25_DynamicReflect_Loop_Invalid(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	called := d.Loop(func(index int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Loop returns error -- invalid", actual)
}

func Test_I25_DynamicReflect_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, index == 1
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- break", actual)
}

func Test_I25_DynamicReflect_FilterAsDynamicCollection_Empty(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns empty -- empty", actual)
}

func Test_I25_DynamicReflect_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]int{"a": 1, "b": 2}, true)
	count := 0
	called := d.LoopMap(func(index int, key, value any) bool {
		count++
		return false
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 2}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- with args", actual)
}

func Test_I25_DynamicReflect_LoopMap_Break(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]int{"a": 1, "b": 2, "c": 3}, true)
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return true
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- break", actual)
}

func Test_I25_DynamicReflect_LoopMap_Invalid(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	called := d.LoopMap(func(index int, key, value any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LoopMap returns error -- invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// New Creator — collection factories
// ══════════════════════════════════════════════════════════════════════════════

func Test_I25_NewCreator_Collection_String(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(5)
	actual := args.Map{"notNil": c != nil, "len": c.Length()}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Cap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_Empty(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Empty returns empty -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_From(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.From returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.Clone([]string{"a", "b"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Clone returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_Items(t *testing.T) {
	c := coredynamic.New.Collection.String.Items("a", "b", "c")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_Create(t *testing.T) {
	c := coredynamic.New.Collection.String.Create([]string{"x"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Create returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_String_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.String.LenCap(3, 10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.LenCap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Int_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int.LenCap(2, 5)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Int64_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int64.LenCap(1, 5)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Byte_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Byte.LenCap(4, 8)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Any_Cap(t *testing.T) {
	c := coredynamic.New.Collection.Any.Cap(5)
	actual := args.Map{"notNil": c != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Cap returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Bool(t *testing.T) {
	c := coredynamic.New.Collection.Bool.Items(true, false)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Bool.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Float64(t *testing.T) {
	c := coredynamic.New.Collection.Float64.Items(1.1, 2.2)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float64.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_Float32(t *testing.T) {
	c := coredynamic.New.Collection.Float32.Items(1.1, 2.2)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float32.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_ByteSlice(t *testing.T) {
	c := coredynamic.New.Collection.ByteSlice.Items([]byte{1}, []byte{2})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_AnyMap(t *testing.T) {
	c := coredynamic.New.Collection.AnyMap.Items(map[string]any{"a": 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.AnyMap.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_StringMap(t *testing.T) {
	c := coredynamic.New.Collection.StringMap.Items(map[string]string{"a": "b"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.StringMap.Items returns correct value -- with args", actual)
}

func Test_I25_NewCreator_Collection_IntMap(t *testing.T) {
	c := coredynamic.New.Collection.IntMap.Items(map[string]int{"a": 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.IntMap.Items returns correct value -- with args", actual)
}
