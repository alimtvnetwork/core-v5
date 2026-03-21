package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — value accessor methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_KeyVal_KeyDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "name", Value: 42}
	d := kv.KeyDynamic()
	actual := args.Map{"valid": d.IsValid(), "val": d.ValueString()}
	expected := args.Map{"valid": true, "val": "name"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyDynamic", actual)
}

func Test_I15_KeyVal_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "hello"}
	d := kv.ValueDynamic()
	actual := args.Map{"valid": d.IsValid(), "val": d.ValueString()}
	expected := args.Map{"valid": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueDynamic", actual)
}

func Test_I15_KeyVal_KeyDynamicPtr(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamicPtr()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyDynamicPtr", actual)
}

func Test_I15_KeyVal_KeyDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"nil": kv.KeyDynamicPtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyDynamicPtr nil", actual)
}

func Test_I15_KeyVal_ValueDynamicPtr(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 99}
	d := kv.ValueDynamicPtr()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueDynamicPtr", actual)
}

func Test_I15_KeyVal_ValueDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"nil": kv.ValueDynamicPtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueDynamicPtr nil", actual)
}

func Test_I15_KeyVal_IsKeyNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"isNull": kv.IsKeyNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsKeyNull", actual)
}

func Test_I15_KeyVal_IsKeyNull_NotNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"isNull": kv.IsKeyNull()}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsKeyNull not null", actual)
}

func Test_I15_KeyVal_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"isNull": kv.IsValueNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsValueNull", actual)
}

func Test_I15_KeyVal_IsValueNull_NotNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"isNull": kv.IsValueNull()}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsValueNull not null", actual)
}

func Test_I15_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: "v"}
	actual := args.Map{"empty": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns empty -- IsKeyNullOrEmptyString", actual)
}

func Test_I15_KeyVal_String(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "name", Value: 42}
	s := kv.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- String", actual)
}

func Test_I15_KeyVal_String_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"empty": kv.String() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- String nil", actual)
}

func Test_I15_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	actual := args.Map{"valid": rv.IsValid(), "val": rv.Interface()}
	expected := args.Map{"valid": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueReflectValue", actual)
}

func Test_I15_KeyVal_ValueInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"val": kv.ValueInt()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt", actual)
}

func Test_I15_KeyVal_ValueInt_NotInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}
	actual := args.Map{"val": kv.ValueInt()}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt not int", actual)
}

func Test_I15_KeyVal_ValueUInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(7)}
	actual := args.Map{"val": kv.ValueUInt()}
	expected := args.Map{"val": uint(7)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueUInt", actual)
}

func Test_I15_KeyVal_ValueUInt_NotUint(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}
	actual := args.Map{"val": kv.ValueUInt()}
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueUInt not uint", actual)
}

func Test_I15_KeyVal_ValueStrings(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	actual := args.Map{"len": len(kv.ValueStrings())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueStrings", actual)
}

func Test_I15_KeyVal_ValueStrings_NotSlice(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"nil": kv.ValueStrings() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueStrings not slice", actual)
}

func Test_I15_KeyVal_ValueBool(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	actual := args.Map{"val": kv.ValueBool()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueBool", actual)
}

func Test_I15_KeyVal_ValueBool_NotBool(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"val": kv.ValueBool()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueBool not bool", actual)
}

func Test_I15_KeyVal_ValueInt64(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(999)}
	actual := args.Map{"val": kv.ValueInt64()}
	expected := args.Map{"val": int64(999)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt64", actual)
}

func Test_I15_KeyVal_ValueInt64_NotInt64(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}
	actual := args.Map{"val": kv.ValueInt64()}
	expected := args.Map{"val": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt64 not int64", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — nil receiver error methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_KeyVal_ValueNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.ValueNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueNullErr nil", actual)
}

func Test_I15_KeyVal_ValueNullErr_NullValue(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"hasErr": kv.ValueNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- ValueNullErr null value", actual)
}

func Test_I15_KeyVal_ValueNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"noErr": kv.ValueNullErr() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- ValueNullErr valid", actual)
}

func Test_I15_KeyVal_KeyNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.KeyNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyNullErr nil", actual)
}

func Test_I15_KeyVal_KeyNullErr_NullKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: nil, Value: 42}
	actual := args.Map{"hasErr": kv.KeyNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- KeyNullErr null key", actual)
}

func Test_I15_KeyVal_KeyNullErr_Valid(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"noErr": kv.KeyNullErr() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- KeyNullErr valid", actual)
}

func Test_I15_KeyVal_KeyString(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 1}
	actual := args.Map{"notEmpty": kv.KeyString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyString", actual)
}

func Test_I15_KeyVal_KeyString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"empty": kv.KeyString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyString nil", actual)
}

func Test_I15_KeyVal_ValueString(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}
	actual := args.Map{"notEmpty": kv.ValueString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueString", actual)
}

func Test_I15_KeyVal_ValueString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"empty": kv.ValueString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueString nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — Reflect set methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_KeyVal_ReflectSetKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 42}
	var target string
	err := kv.ReflectSetKey(&target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "mykey"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetKey", actual)
}

func Test_I15_KeyVal_ReflectSetKey_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.ReflectSetKey(nil) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ReflectSetKey nil", actual)
}

func Test_I15_KeyVal_KeyReflectSet(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "abc", Value: 1}
	var target string
	err := kv.KeyReflectSet(&target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "abc"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyReflectSet", actual)
}

func Test_I15_KeyVal_KeyReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.KeyReflectSet(nil) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyReflectSet nil", actual)
}

func Test_I15_KeyVal_ValueReflectSet(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	var target int
	err := kv.ValueReflectSet(&target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueReflectSet", actual)
}

func Test_I15_KeyVal_ValueReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.ValueReflectSet(nil) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueReflectSet nil", actual)
}

func Test_I15_KeyVal_ReflectSetTo(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "world"}
	var target string
	err := kv.ReflectSetTo(&target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "world"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetTo", actual)
}

func Test_I15_KeyVal_ReflectSetTo_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"hasErr": kv.ReflectSetTo(nil) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ReflectSetTo nil", actual)
}

func Test_I15_KeyVal_ReflectSetToMust_Success(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	var target int
	kv.ReflectSetToMust(&target)
	actual := args.Map{"val": target}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetToMust success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — JSON methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_KeyVal_JsonModel(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"notNil": kv.JsonModel() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonModel", actual)
}

func Test_I15_KeyVal_JsonModelAny(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"notNil": kv.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonModelAny", actual)
}

func Test_I15_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	actual := args.Map{"noErr": !jr.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- Json", actual)
}

func Test_I15_KeyVal_JsonPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.JsonPtr()
	actual := args.Map{"notNil": jr != nil, "noErr": !jr.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonPtr", actual)
}

func Test_I15_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	result, err := kv.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ParseInjectUsingJson", actual)
}

func Test_I15_KeyVal_ParseInjectUsingJsonMust(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	result := kv.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_I15_KeyVal_JsonParseSelfInject(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	err := kv.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonParseSelfInject", actual)
}

func Test_I15_KeyVal_Serialize(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	bytes, err := kv.Serialize()
	actual := args.Map{"noErr": err == nil, "notEmpty": len(bytes) > 0}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- Serialize", actual)
}

func Test_I15_KeyVal_CastKeyVal_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- CastKeyVal nil", actual)
}

func Test_I15_KeyVal_CastKeyVal_Success(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 42}
	var k string
	var v int
	err := kv.CastKeyVal(&k, &v)
	actual := args.Map{"noErr": err == nil, "k": k, "v": v}
	expected := args.Map{"noErr": true, "k": "mykey", "v": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- CastKeyVal success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// New.Collection creators — Generic, String, Int, Int64, Byte, Any
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_NewCollection_String_Empty(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	actual := args.Map{"notNil": c != nil, "len": c.Length()}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_String_Cap(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(10)
	actual := args.Map{"notNil": c != nil, "len": c.Length()}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Cap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_String_From(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.From returns correct value -- with args", actual)
}

func Test_I15_NewCollection_String_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.Clone([]string{"x", "y"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Clone returns correct value -- with args", actual)
}

func Test_I15_NewCollection_String_Items(t *testing.T) {
	c := coredynamic.New.Collection.String.Items("a", "b")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Items returns correct value -- with args", actual)
}

func Test_I15_NewCollection_String_Create(t *testing.T) {
	c := coredynamic.New.Collection.String.Create([]string{"x"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Create returns correct value -- with args", actual)
}

func Test_I15_NewCollection_String_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.String.LenCap(3, 10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.LenCap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Int_Cap(t *testing.T) {
	c := coredynamic.New.Collection.Int.Cap(5)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Cap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int_From(t *testing.T) {
	c := coredynamic.New.Collection.Int.From([]int{1, 2, 3})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.From returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int.LenCap(2, 8)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int64_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int64.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Int64_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int64.LenCap(1, 4)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Byte_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Byte.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Byte_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Byte.LenCap(5, 10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Any_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Any.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Any_From(t *testing.T) {
	c := coredynamic.New.Collection.Any.From([]any{"a", 1, true})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.From returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Bool_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Bool.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Bool.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Float64_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Float64.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float64.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_Float32_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Float32.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float32.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_AnyMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.AnyMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.AnyMap.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_StringMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.StringMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.StringMap.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_IntMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.IntMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.IntMap.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_ByteSlice_Empty(t *testing.T) {
	c := coredynamic.New.Collection.ByteSlice.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.Empty returns empty -- with args", actual)
}

func Test_I15_NewCollection_ByteSlice_From(t *testing.T) {
	c := coredynamic.New.Collection.ByteSlice.From([][]byte{{1, 2}, {3, 4}})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.From returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int_Clone(t *testing.T) {
	c := coredynamic.New.Collection.Int.Clone([]int{10, 20, 30})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Clone returns correct value -- with args", actual)
}

func Test_I15_NewCollection_Int_Items(t *testing.T) {
	c := coredynamic.New.Collection.Int.Items(1, 2, 3, 4)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Items returns correct value -- with args", actual)
}
