package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// KeyVal — Dynamic accessors
// =============================================================================

func Test_Cov43_KeyVal_KeyDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "mykey", Value: 42}
	d := kv.KeyDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamic", actual)
}

func Test_Cov43_KeyVal_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.ValueDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamic", actual)
}

func Test_Cov43_KeyVal_KeyDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"isNil": kv.KeyDynamicPtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamicPtr nil", actual)
}

func Test_Cov43_KeyVal_KeyDynamicPtr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamicPtr()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamicPtr valid", actual)
}

func Test_Cov43_KeyVal_ValueDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"isNil": kv.ValueDynamicPtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamicPtr nil", actual)
}

func Test_Cov43_KeyVal_ValueDynamicPtr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 99}
	d := kv.ValueDynamicPtr()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamicPtr valid", actual)
}

// =============================================================================
// KeyVal — Null checks
// =============================================================================

func Test_Cov43_KeyVal_IsKeyNull_True(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"r": kv.IsKeyNull()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNull true", actual)
}

func Test_Cov43_KeyVal_IsKeyNull_False(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"r": kv.IsKeyNull()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNull false", actual)
}

func Test_Cov43_KeyVal_IsKeyNullOrEmptyString_Null(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString null", actual)
}

func Test_Cov43_KeyVal_IsKeyNullOrEmptyString_Empty(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: "v"}
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString empty", actual)
}

func Test_Cov43_KeyVal_IsKeyNullOrEmptyString_NonEmpty(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString non-empty", actual)
}

func Test_Cov43_KeyVal_IsValueNull_True(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"r": kv.IsValueNull()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsValueNull true", actual)
}

func Test_Cov43_KeyVal_IsValueNull_False(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"r": kv.IsValueNull()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsValueNull false", actual)
}

// =============================================================================
// KeyVal — String, KeyString, ValueString
// =============================================================================

func Test_Cov43_KeyVal_String_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.String()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal String nil", actual)
}

func Test_Cov43_KeyVal_String_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	s := kv.String()
	actual := args.Map{"nonEmpty": s != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal String valid", actual)
}

func Test_Cov43_KeyVal_KeyString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.KeyString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString nil receiver", actual)
}

func Test_Cov43_KeyVal_KeyString_NilKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"r": kv.KeyString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString nil key", actual)
}

func Test_Cov43_KeyVal_KeyString_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "myKey", Value: "v"}
	actual := args.Map{"r": kv.KeyString()}
	expected := args.Map{"r": "myKey"}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString valid", actual)
}

func Test_Cov43_KeyVal_ValueString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"r": kv.ValueString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString nil", actual)
}

func Test_Cov43_KeyVal_ValueString_NilValue(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"r": kv.ValueString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString nil value", actual)
}

func Test_Cov43_KeyVal_ValueString_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}
	actual := args.Map{"r": kv.ValueString()}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString valid", actual)
}

// =============================================================================
// KeyVal — Value typed accessors
// =============================================================================

func Test_Cov43_KeyVal_ValueInt_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"r": kv.ValueInt()}
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt valid", actual)
}

func Test_Cov43_KeyVal_ValueInt_Invalid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notint"}
	actual := args.Map{"r": kv.ValueInt()}
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt invalid", actual)
}

func Test_Cov43_KeyVal_ValueUInt_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(10)}
	actual := args.Map{"r": kv.ValueUInt()}
	expected := args.Map{"r": uint(10)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt valid", actual)
}

func Test_Cov43_KeyVal_ValueUInt_Invalid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notuint"}
	actual := args.Map{"r": kv.ValueUInt()}
	expected := args.Map{"r": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt invalid", actual)
}

func Test_Cov43_KeyVal_ValueBool_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	actual := args.Map{"r": kv.ValueBool()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool valid", actual)
}

func Test_Cov43_KeyVal_ValueBool_Invalid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notbool"}
	actual := args.Map{"r": kv.ValueBool()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool invalid", actual)
}

func Test_Cov43_KeyVal_ValueInt64_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(999)}
	actual := args.Map{"r": kv.ValueInt64()}
	expected := args.Map{"r": int64(999)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 valid", actual)
}

func Test_Cov43_KeyVal_ValueInt64_Invalid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notint64"}
	actual := args.Map{"r": kv.ValueInt64()}
	expected := args.Map{"r": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 invalid", actual)
}

func Test_Cov43_KeyVal_ValueStrings_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	actual := args.Map{"len": len(kv.ValueStrings())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings valid", actual)
}

func Test_Cov43_KeyVal_ValueStrings_Invalid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "notslice"}
	actual := args.Map{"isNil": kv.ValueStrings() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings invalid", actual)
}

// =============================================================================
// KeyVal — Null error methods
// =============================================================================

func Test_Cov43_KeyVal_ValueNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr nil", actual)
}

func Test_Cov43_KeyVal_ValueNullErr_NullValue(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}
	err := kv.ValueNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr null value", actual)
}

func Test_Cov43_KeyVal_ValueNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"noErr": kv.ValueNullErr() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr valid", actual)
}

func Test_Cov43_KeyVal_KeyNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr nil", actual)
}

func Test_Cov43_KeyVal_KeyNullErr_NullKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: nil, Value: 42}
	err := kv.KeyNullErr()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr null key", actual)
}

func Test_Cov43_KeyVal_KeyNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"noErr": kv.KeyNullErr() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr valid", actual)
}

// =============================================================================
// KeyVal — CastKeyVal
// =============================================================================

func Test_Cov43_KeyVal_CastKeyVal_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	var k, v string
	err := kv.CastKeyVal(&k, &v)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal nil", actual)
}

func Test_Cov43_KeyVal_CastKeyVal_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "mykey", Value: "myval"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)
	actual := args.Map{"noErr": err == nil, "val": v}
	expected := args.Map{"noErr": true, "val": "myval"}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal valid", actual)
}

// =============================================================================
// KeyVal — ReflectSet methods
// =============================================================================

func Test_Cov43_KeyVal_ReflectSetKey_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	var k string
	err := kv.ReflectSetKey(&k)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetKey nil", actual)
}

func Test_Cov43_KeyVal_KeyReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	var k string
	err := kv.KeyReflectSet(&k)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyReflectSet nil", actual)
}

func Test_Cov43_KeyVal_ValueReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	var v string
	err := kv.ValueReflectSet(&v)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectSet nil", actual)
}

func Test_Cov43_KeyVal_ReflectSetTo_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	var v string
	err := kv.ReflectSetTo(&v)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetTo nil", actual)
}

func Test_Cov43_KeyVal_ReflectSetTo_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}
	var v string
	err := kv.ReflectSetTo(&v)
	actual := args.Map{"noErr": err == nil, "v": v}
	expected := args.Map{"noErr": true, "v": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetTo valid", actual)
}

func Test_Cov43_KeyVal_ReflectSetToMust_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "world"}
	var v string
	kv.ReflectSetToMust(&v)
	actual := args.Map{"v": v}
	expected := args.Map{"v": "world"}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetToMust valid", actual)
}

func Test_Cov43_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	actual := args.Map{"valid": rv.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectValue", actual)
}

// =============================================================================
// KeyVal — JSON methods
// =============================================================================

func Test_Cov43_KeyVal_JsonModel(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"notNil": kv.JsonModel() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonModel", actual)
}

func Test_Cov43_KeyVal_JsonModelAny(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"notNil": kv.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonModelAny", actual)
}

func Test_Cov43_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	r := kv.Json()
	actual := args.Map{"hasBytes": r.HasAnyItem()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "KeyVal Json", actual)
}

func Test_Cov43_KeyVal_JsonPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"notNil": kv.JsonPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonPtr", actual)
}

func Test_Cov43_KeyVal_Serialize(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "KeyVal Serialize", actual)
}

func Test_Cov43_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	result, err := kv2.ParseInjectUsingJson(&jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJson", actual)
}

func Test_Cov43_KeyVal_ParseInjectUsingJsonMust(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	result := kv2.ParseInjectUsingJsonMust(&jr)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJsonMust", actual)
}

func Test_Cov43_KeyVal_JsonParseSelfInject(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	err := kv2.JsonParseSelfInject(&jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonParseSelfInject", actual)
}
