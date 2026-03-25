package coredynamictests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Dynamic — nil receiver branches (DynamicGetters.go)
// =============================================================================

func Test_Cov37_Dynamic_Length_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Dynamic Length nil receiver", actual)
}

func Test_Cov37_Dynamic_StructStringPtr_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"isNil": d.StructStringPtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic StructStringPtr nil receiver", actual)
}

func Test_Cov37_Dynamic_String_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.String()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic String nil receiver", actual)
}

func Test_Cov37_Dynamic_StructString_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.StructString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic StructString nil receiver", actual)
}

func Test_Cov37_Dynamic_IsPointer_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsPointer()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsPointer nil receiver", actual)
}

func Test_Cov37_Dynamic_IsValueType_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsValueType()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsValueType nil receiver", actual)
}

func Test_Cov37_Dynamic_IsStructStringNullOrEmpty_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsStructStringNullOrEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStructStringNullOrEmpty nil receiver", actual)
}

func Test_Cov37_Dynamic_IsStructStringNullOrEmptyOrWhitespace_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStructStringNullOrEmptyOrWhitespace nil receiver", actual)
}

func Test_Cov37_Dynamic_IsPrimitive_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsPrimitive()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsPrimitive nil receiver", actual)
}

func Test_Cov37_Dynamic_IsNumber_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsNumber()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsNumber nil receiver", actual)
}

func Test_Cov37_Dynamic_IsStringType_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsStringType()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStringType nil receiver", actual)
}

func Test_Cov37_Dynamic_IsStruct_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsStruct()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStruct nil receiver", actual)
}

func Test_Cov37_Dynamic_IsFunc_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsFunc()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsFunc nil receiver", actual)
}

func Test_Cov37_Dynamic_IsSliceOrArray_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsSliceOrArray()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsSliceOrArray nil receiver", actual)
}

func Test_Cov37_Dynamic_IsSliceOrArrayOrMap_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsSliceOrArrayOrMap()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsSliceOrArrayOrMap nil receiver", actual)
}

func Test_Cov37_Dynamic_IsMap_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.IsMap()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsMap nil receiver", actual)
}

func Test_Cov37_Dynamic_Bytes_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	actual := args.Map{"nil": b == nil, "ok": ok}
	expected := args.Map{"nil": true, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Bytes nil receiver", actual)
}

func Test_Cov37_Dynamic_ValueNullErr_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr nil receiver", actual)
}

func Test_Cov37_Dynamic_ValueString_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"r": d.ValueString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueString nil receiver", actual)
}

// =============================================================================
// Dynamic — Value extraction: type mismatch branches
// =============================================================================

func Test_Cov37_Dynamic_IntDefault_NilData(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	val, ok := d.IntDefault(42)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault nil data", actual)
}

func Test_Cov37_Dynamic_IntDefault_NonNumeric(t *testing.T) {
	d := coredynamic.NewDynamicPtr("abc", true)
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault non-numeric", actual)
}

func Test_Cov37_Dynamic_IntDefault_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("123", true)
	val, ok := d.IntDefault(0)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 123, "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault valid", actual)
}

func Test_Cov37_Dynamic_Float64_NilData(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 nil data", actual)
}

func Test_Cov37_Dynamic_Float64_NonNumeric(t *testing.T) {
	d := coredynamic.NewDynamicPtr("abc", true)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 non-numeric", actual)
}

func Test_Cov37_Dynamic_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("3.14", true)
	val, err := d.Float64()
	actual := args.Map{"noErr": err == nil, "gt3": val > 3.0}
	expected := args.Map{"noErr": true, "gt3": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 valid", actual)
}

func Test_Cov37_Dynamic_ValueInt_Mismatch(t *testing.T) {
	d := coredynamic.NewDynamic("notint", true)
	actual := args.Map{"r": d.ValueInt()}
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueInt mismatch", actual)
}

func Test_Cov37_Dynamic_ValueUInt_Mismatch(t *testing.T) {
	d := coredynamic.NewDynamic("notuint", true)
	actual := args.Map{"r": d.ValueUInt()}
	expected := args.Map{"r": uint(0)}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueUInt mismatch", actual)
}

func Test_Cov37_Dynamic_ValueStrings_Mismatch(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	actual := args.Map{"isNil": d.ValueStrings() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueStrings mismatch", actual)
}

func Test_Cov37_Dynamic_ValueBool_Mismatch(t *testing.T) {
	d := coredynamic.NewDynamic("notbool", true)
	actual := args.Map{"r": d.ValueBool()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueBool mismatch", actual)
}

func Test_Cov37_Dynamic_ValueInt64_Mismatch(t *testing.T) {
	d := coredynamic.NewDynamic("notint64", true)
	actual := args.Map{"r": d.ValueInt64()}
	expected := args.Map{"r": int64(-1)}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueInt64 mismatch", actual)
}

func Test_Cov37_Dynamic_ValueNullErr_NullData(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	err := d.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr null data", actual)
}

func Test_Cov37_Dynamic_ValueNullErr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("data", true)
	err := d.ValueNullErr()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr valid", actual)
}

func Test_Cov37_Dynamic_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	r := d.ValueString()
	actual := args.Map{"nonEmpty": len(r) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueString non-string", actual)
}

func Test_Cov37_Dynamic_Bytes_NotBytes(t *testing.T) {
	d := coredynamic.NewDynamicPtr("notbytes", true)
	_, ok := d.Bytes()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Bytes not bytes", actual)
}

func Test_Cov37_Dynamic_Bytes_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]byte{1, 2, 3}, true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "len": len(b)}
	expected := args.Map{"ok": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic Bytes valid", actual)
}

// =============================================================================
// Dynamic — Type check branches (positive cases)
// =============================================================================

func Test_Cov37_Dynamic_IsPointer_True(t *testing.T) {
	s := "hello"
	d := coredynamic.NewDynamicPtr(&s, true)
	actual := args.Map{"r": d.IsPointer()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsPointer true", actual)
}

func Test_Cov37_Dynamic_IsValueType_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"r": d.IsValueType()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsValueType true", actual)
}

func Test_Cov37_Dynamic_IsStruct_True(t *testing.T) {
	type s struct{ A int }
	d := coredynamic.NewDynamicPtr(s{A: 1}, true)
	actual := args.Map{"r": d.IsStruct()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStruct true", actual)
}

func Test_Cov37_Dynamic_IsFunc_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(func() {}, true)
	actual := args.Map{"r": d.IsFunc()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsFunc true", actual)
}

func Test_Cov37_Dynamic_IsSliceOrArray_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2}, true)
	actual := args.Map{"r": d.IsSliceOrArray()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsSliceOrArray true", actual)
}

func Test_Cov37_Dynamic_IsMap_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"r": d.IsMap()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsMap true", actual)
}

func Test_Cov37_Dynamic_IsSliceOrArrayOrMap_Map(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"r": d.IsSliceOrArrayOrMap()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsSliceOrArrayOrMap map", actual)
}

func Test_Cov37_Dynamic_IsNumber_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"r": d.IsNumber()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsNumber true", actual)
}

func Test_Cov37_Dynamic_IsPrimitive_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"r": d.IsPrimitive()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsPrimitive true", actual)
}

func Test_Cov37_Dynamic_IsStringType_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"r": d.IsStringType()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsStringType true", actual)
}

// =============================================================================
// KeyVal — nil receiver branches
// =============================================================================

func Test_Cov37_KeyVal_KeyDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"isNil": kv.KeyDynamicPtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamicPtr nil", actual)
}

func Test_Cov37_KeyVal_ValueDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"isNil": kv.ValueDynamicPtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamicPtr nil", actual)
}

func Test_Cov37_KeyVal_String_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.String()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal String nil", actual)
}

func Test_Cov37_KeyVal_CastKeyVal_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal nil", actual)
}

func Test_Cov37_KeyVal_ReflectSetKey_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetKey(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetKey nil", actual)
}

func Test_Cov37_KeyVal_ValueNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr nil", actual)
}

func Test_Cov37_KeyVal_KeyNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr nil", actual)
}

func Test_Cov37_KeyVal_KeyString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.KeyString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString nil", actual)
}

func Test_Cov37_KeyVal_ValueString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.ValueString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString nil", actual)
}

func Test_Cov37_KeyVal_KeyReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyReflectSet(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyReflectSet nil", actual)
}

func Test_Cov37_KeyVal_ValueReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueReflectSet(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectSet nil", actual)
}

func Test_Cov37_KeyVal_ReflectSetTo_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetTo(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetTo nil", actual)
}

// =============================================================================
// KeyVal — Value extraction: type mismatch branches
// =============================================================================

func Test_Cov37_KeyVal_ValueInt_Mismatch(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notint"}
	actual := args.Map{"r": kv.ValueInt()}
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt mismatch", actual)
}

func Test_Cov37_KeyVal_ValueUInt_Mismatch(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notuint"}
	actual := args.Map{"r": kv.ValueUInt()}
	expected := args.Map{"r": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt mismatch", actual)
}

func Test_Cov37_KeyVal_ValueStrings_Mismatch(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"isNil": kv.ValueStrings() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings mismatch", actual)
}

func Test_Cov37_KeyVal_ValueBool_Mismatch(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notbool"}
	actual := args.Map{"r": kv.ValueBool()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool mismatch", actual)
}

func Test_Cov37_KeyVal_ValueInt64_Mismatch(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notint64"}
	actual := args.Map{"r": kv.ValueInt64()}
	expected := args.Map{"r": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 mismatch", actual)
}

func Test_Cov37_KeyVal_ValueNullErr_NullValue(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}
	err := kv.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr null value", actual)
}

func Test_Cov37_KeyVal_KeyNullErr_NullKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: nil, Value: "v"}
	err := kv.KeyNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr null key", actual)
}

func Test_Cov37_KeyVal_KeyNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	err := kv.KeyNullErr()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr valid", actual)
}

func Test_Cov37_KeyVal_ValueNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	err := kv.ValueNullErr()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr valid", actual)
}

// =============================================================================
// KeyVal — JSON branches
// =============================================================================

func Test_Cov37_KeyVal_ParseInjectUsingJson_Error(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := kv.ParseInjectUsingJson(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJson error", actual)
}

func Test_Cov37_KeyVal_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		kv.ParseInjectUsingJsonMust(jr)
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJsonMust panics", actual)
}

func Test_Cov37_KeyVal_Serialize_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal Serialize valid", actual)
}

func Test_Cov37_KeyVal_ReflectSetToMust_Panics(t *testing.T) {
	var kv *coredynamic.KeyVal
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		kv.ReflectSetToMust(nil)
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetToMust panics", actual)
}

// =============================================================================
// KeyVal — Positive value extraction
// =============================================================================

func Test_Cov37_KeyVal_ValueInt_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"r": kv.ValueInt()}
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt valid", actual)
}

func Test_Cov37_KeyVal_ValueUInt_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(7)}
	actual := args.Map{"r": kv.ValueUInt()}
	expected := args.Map{"r": uint(7)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt valid", actual)
}

func Test_Cov37_KeyVal_ValueBool_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	actual := args.Map{"r": kv.ValueBool()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool valid", actual)
}

func Test_Cov37_KeyVal_ValueInt64_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(99)}
	actual := args.Map{"r": kv.ValueInt64()}
	expected := args.Map{"r": int64(99)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 valid", actual)
}

func Test_Cov37_KeyVal_ValueStrings_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	actual := args.Map{"len": len(kv.ValueStrings())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings valid", actual)
}

func Test_Cov37_KeyVal_KeyDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "mykey", Value: "myval"}
	d := kv.KeyDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamic", actual)
}

func Test_Cov37_KeyVal_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.ValueDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamic", actual)
}

func Test_Cov37_KeyVal_IsKeyNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"r": kv.IsKeyNull()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNull", actual)
}

func Test_Cov37_KeyVal_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"r": kv.IsValueNull()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsValueNull", actual)
}
