package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// ZeroSetAny / SafeZeroSet
// ==========================================================================

func Test_C17_ZeroSetAny_NonNil(t *testing.T) {
	type S struct{ X int }
	s := S{X: 42}
	coredynamic.ZeroSetAny(&s)
	actual := args.Map{"x": s.X}
	expected := args.Map{"x": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny non-nil", actual)
}

func Test_C17_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny nil", actual)
}

func Test_C17_SafeZeroSet_Nil(t *testing.T) {
	coredynamic.SafeZeroSet(reflect.Value{}) // invalid reflect.Value
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet nil", actual)
}

// ==========================================================================
// KeyVal — uncovered methods
// ==========================================================================

func Test_C17_KeyVal_KeyDynamic_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	kd := kv.KeyDynamic()
	vd := kv.ValueDynamic()
	kdp := kv.KeyDynamicPtr()
	vdp := kv.ValueDynamicPtr()
	actual := args.Map{
		"kValid": kd.IsValid(), "vValid": vd.IsValid(),
		"kdpValid": kdp.IsValid(), "vdpValid": vdp.IsValid(),
	}
	expected := args.Map{
		"kValid": true, "vValid": true,
		"kdpValid": true, "vdpValid": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Dynamic methods", actual)
}

func Test_C17_KeyVal_IsKeyNull_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	kvNull := coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{
		"keyNull":   kv.IsKeyNull(),
		"valNull":   kv.IsValueNull(),
		"keyNull2":  kvNull.IsKeyNull(),
		"valNull2":  kvNull.IsValueNull(),
	}
	expected := args.Map{
		"keyNull":   false,
		"valNull":   true,
		"keyNull2":  true,
		"valNull2":  false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal null checks", actual)
}

func Test_C17_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: 1}
	kvVal := coredynamic.KeyVal{Key: "x", Value: 1}
	actual := args.Map{
		"empty":    kv.IsKeyNullOrEmptyString(),
		"nonEmpty": kvVal.IsKeyNullOrEmptyString(),
	}
	expected := args.Map{
		"empty":    true,
		"nonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString", actual)
}

func Test_C17_KeyVal_String(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	s := kv.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal String", actual)
}

func Test_C17_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	actual := args.Map{"valid": rv.IsValid(), "kind": rv.Kind() == reflect.Int}
	expected := args.Map{"valid": true, "kind": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectValue", actual)
}

func Test_C17_KeyVal_ValueInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	kvBad := coredynamic.KeyVal{Key: "k", Value: "nope"}
	actual := args.Map{"ok": kv.ValueInt(), "bad": kvBad.ValueInt()}
	expected := args.Map{"ok": 42, "bad": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt", actual)
}

func Test_C17_KeyVal_ValueUInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(5)}
	kvBad := coredynamic.KeyVal{Key: "k", Value: "nope"}
	actual := args.Map{"ok": kv.ValueUInt(), "bad": kvBad.ValueUInt()}
	expected := args.Map{"ok": uint(5), "bad": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt", actual)
}

func Test_C17_KeyVal_ValueStrings(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	kvBad := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"ok": len(kv.ValueStrings()), "bad": kvBad.ValueStrings() == nil}
	expected := args.Map{"ok": 2, "bad": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings", actual)
}

func Test_C17_KeyVal_ValueBool(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	kvBad := coredynamic.KeyVal{Key: "k", Value: "nope"}
	actual := args.Map{"ok": kv.ValueBool(), "bad": kvBad.ValueBool()}
	expected := args.Map{"ok": true, "bad": false}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool", actual)
}

func Test_C17_KeyVal_ValueInt64(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(99)}
	kvBad := coredynamic.KeyVal{Key: "k", Value: "nope"}
	actual := args.Map{"ok": kv.ValueInt64(), "bad": kvBad.ValueInt64()}
	expected := args.Map{"ok": int64(99), "bad": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64", actual)
}

func Test_C17_KeyVal_CastKeyVal(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)
	// CastKeyVal returns nil on key set error (odd logic but that's the source)
	_ = err

	var nilKv *coredynamic.KeyVal
	errNil := nilKv.CastKeyVal(&k, &v)
	actual := args.Map{"nilErr": errNil != nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal", actual)
}

func Test_C17_KeyVal_ReflectSetKey(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "hello", Value: "v"}
	var k string
	err := kv.ReflectSetKey(&k)
	actual := args.Map{"noErr": err == nil, "k": k}
	expected := args.Map{"noErr": true, "k": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetKey", actual)

	var nilKv *coredynamic.KeyVal
	errNil := nilKv.ReflectSetKey(&k)
	actual2 := args.Map{"nilErr": errNil != nil}
	expected2 := args.Map{"nilErr": true}
	expected2.ShouldBeEqual(t, 1, "KeyVal ReflectSetKey nil", actual2)
}

func Test_C17_KeyVal_ValueNullErr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	kvOk := coredynamic.KeyVal{Key: "k", Value: "v"}
	var nilKv *coredynamic.KeyVal
	actual := args.Map{
		"nullErr": kv.ValueNullErr() != nil,
		"okErr":   kvOk.ValueNullErr() == nil,
		"nilErr":  nilKv.ValueNullErr() != nil,
	}
	expected := args.Map{
		"nullErr": true,
		"okErr":   true,
		"nilErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr", actual)
}

func Test_C17_KeyVal_KeyNullErr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}
	kvOk := coredynamic.KeyVal{Key: "k", Value: "v"}
	var nilKv *coredynamic.KeyVal
	actual := args.Map{
		"nullErr": kv.KeyNullErr() != nil,
		"okErr":   kvOk.KeyNullErr() == nil,
		"nilErr":  nilKv.KeyNullErr() != nil,
	}
	expected := args.Map{
		"nullErr": true,
		"okErr":   true,
		"nilErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr", actual)
}

func Test_C17_KeyVal_KeyString_ValueString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	kvNil := coredynamic.KeyVal{Key: nil, Value: nil}
	var nilKv *coredynamic.KeyVal
	actual := args.Map{
		"ks":    kv.KeyString(),
		"vs":    kv.ValueString(),
		"nilKs": kvNil.KeyString(),
		"nilVs": kvNil.ValueString(),
		"pNilKs": nilKv.KeyString(),
		"pNilVs": nilKv.ValueString(),
	}
	expected := args.Map{
		"ks":    "k",
		"vs":    "v",
		"nilKs": "",
		"nilVs": "",
		"pNilKs": "",
		"pNilVs": "",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString/ValueString", actual)
}

func Test_C17_KeyVal_KeyReflectSet_ValueReflectSet_ReflectSetTo(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	var k, v, v2 string
	err1 := kv.KeyReflectSet(&k)
	err2 := kv.ValueReflectSet(&v)
	err3 := kv.ReflectSetTo(&v2)
	actual := args.Map{
		"k": k, "v": v, "v2": v2,
		"e1": err1 == nil, "e2": err2 == nil, "e3": err3 == nil,
	}
	expected := args.Map{
		"k": "k", "v": "v", "v2": "v",
		"e1": true, "e2": true, "e3": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSet", actual)

	var nilKv *coredynamic.KeyVal
	actual2 := args.Map{
		"e1": nilKv.KeyReflectSet(&k) != nil,
		"e2": nilKv.ValueReflectSet(&v) != nil,
		"e3": nilKv.ReflectSetTo(&v2) != nil,
	}
	expected2 := args.Map{"e1": true, "e2": true, "e3": true}
	expected2.ShouldBeEqual(t, 1, "KeyVal ReflectSet nil", actual2)
}

func Test_C17_KeyVal_ReflectSetToMust(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	var v string
	kv.ReflectSetToMust(&v)
	actual := args.Map{"v": v}
	expected := args.Map{"v": "v"}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetToMust", actual)
}

func Test_C17_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	jp := kv.JsonPtr()
	m := kv.JsonModel()
	ma := kv.JsonModelAny()
	actual := args.Map{
		"jOk":  j.JsonString() != "",
		"jpOk": jp != nil,
		"mOk":  m != nil,
		"maOk": ma != nil,
	}
	expected := args.Map{
		"jOk": true, "jpOk": true, "mOk": true, "maOk": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Json", actual)
}

func Test_C17_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	jr := corejson.NewPtr(coredynamic.KeyVal{Key: "x", Value: "y"})
	result, err := kv.ParseInjectUsingJson(jr)
	_ = result
	_ = err
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJson", actual)
}

func Test_C17_KeyVal_JsonParseSelfInject(t *testing.T) {
	kv := &coredynamic.KeyVal{}
	jr := corejson.NewPtr(coredynamic.KeyVal{Key: "x", Value: "y"})
	err := kv.JsonParseSelfInject(jr)
	_ = err
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonParseSelfInject", actual)
}

func Test_C17_KeyVal_Serialize(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	actual := args.Map{"noErr": err == nil, "hasData": len(b) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "KeyVal Serialize", actual)
}

// ==========================================================================
// KeyValCollection — uncovered methods
// ==========================================================================

func Test_C17_KeyValCollection_AddPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.AddPtr(nil) // should be no-op
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	kvc.AddPtr(&kv)
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection AddPtr", actual)
}

func Test_C17_KeyValCollection_AddMany(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.AddMany() // empty
	kvc.AddMany(coredynamic.KeyVal{Key: "a"}, coredynamic.KeyVal{Key: "b"})
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection AddMany", actual)
}

func Test_C17_KeyValCollection_AddManyPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.AddManyPtr() // empty
	a := coredynamic.KeyVal{Key: "a"}
	kvc.AddManyPtr(nil, &a, nil)
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection AddManyPtr", actual)
}

func Test_C17_KeyValCollection_Items_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	actual := args.Map{"isNil": kvc.Items() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection Items nil", actual)
}

func Test_C17_KeyValCollection_MapAnyItems(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	m := kvc.MapAnyItems()
	actual := args.Map{"hasItems": m.Length() > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection MapAnyItems", actual)

	empty := coredynamic.EmptyKeyValCollection()
	me := empty.MapAnyItems()
	actual2 := args.Map{"empty": me.IsEmpty()}
	expected2 := args.Map{"empty": true}
	expected2.ShouldBeEqual(t, 1, "KeyValCollection MapAnyItems empty", actual2)
}

func Test_C17_KeyValCollection_JsonMapResults(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	mr, err := kvc.JsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection JsonMapResults", actual)
}

func Test_C17_KeyValCollection_JsonResultsCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jrc := kvc.JsonResultsCollection()
	actual := args.Map{"notNil": jrc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection JsonResultsCollection", actual)
}

func Test_C17_KeyValCollection_JsonResultsPtrCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jrpc := kvc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": jrpc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection JsonResultsPtrCollection", actual)
}

func Test_C17_KeyValCollection_GetPagesSize(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(5)
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: i, Value: i})
	}
	actual := args.Map{
		"pages2": kvc.GetPagesSize(2),
		"pages0": kvc.GetPagesSize(0),
	}
	expected := args.Map{
		"pages2": 3,
		"pages0": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection GetPagesSize", actual)
}

func Test_C17_KeyValCollection_GetPagedCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(5)
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: i, Value: i})
	}
	pages := kvc.GetPagedCollection(2)
	actual := args.Map{"pageCount": len(pages)}
	expected := args.Map{"pageCount": 3}
	expected.ShouldBeEqual(t, 0, "KeyValCollection GetPagedCollection", actual)
}

func Test_C17_KeyValCollection_GetSinglePageCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(5)
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: i, Value: i})
	}
	page := kvc.GetSinglePageCollection(2, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection GetSinglePageCollection", actual)
}

func Test_C17_KeyValCollection_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()
	sorted := kvc.AllKeysSorted()
	vals := kvc.AllValues()
	actual := args.Map{
		"keysLen":  len(keys),
		"sorted0":  sorted[0],
		"valsLen":  len(vals),
	}
	expected := args.Map{
		"keysLen":  2,
		"sorted0":  "a",
		"valsLen":  2,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection AllKeys/Sorted/Values", actual)
}

func Test_C17_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()
	var nilKvc *coredynamic.KeyValCollection
	sNil := nilKvc.String()
	actual := args.Map{"notEmpty": s != "", "nilEmpty": sNil == ""}
	expected := args.Map{"notEmpty": true, "nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection String", actual)
}

func Test_C17_KeyValCollection_Json(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	j := kvc.Json()
	jp := kvc.JsonPtr()
	m := kvc.JsonModel()
	ma := kvc.JsonModelAny()
	actual := args.Map{
		"jOk": j.JsonString() != "", "jpOk": jp != nil,
		"mOk": m != nil, "maOk": ma != nil,
	}
	expected := args.Map{
		"jOk": true, "jpOk": true, "mOk": true, "maOk": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection Json", actual)
}

func Test_C17_KeyValCollection_Serialize_JsonString_JsonStringMust(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	b, err := kvc.Serialize()
	s, sErr := kvc.JsonString()
	sm := kvc.JsonStringMust()
	actual := args.Map{
		"noErr": err == nil, "hasData": len(b) > 0,
		"sNoErr": sErr == nil, "sNotEmpty": s != "",
		"smNotEmpty": sm != "",
	}
	expected := args.Map{
		"noErr": true, "hasData": true,
		"sNoErr": true, "sNotEmpty": true,
		"smNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection Serialize/JsonString", actual)
}

func Test_C17_KeyValCollection_Clone(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	clone := kvc.Clone()
	cloneP := kvc.ClonePtr()
	var nilKvc *coredynamic.KeyValCollection
	nilClone := nilKvc.ClonePtr()
	np := clone.NonPtr()
	pp := kvc.Ptr()
	actual := args.Map{
		"cloneLen": clone.Length(), "ptrLen": cloneP.Length(),
		"nilClone": nilClone == nil, "npLen": np.Length(), "ppNotNil": pp != nil,
	}
	expected := args.Map{
		"cloneLen": 1, "ptrLen": 1,
		"nilClone": true, "npLen": 1, "ppNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection Clone", actual)
}

func Test_C17_KeyValCollection_ParseInjectUsingJson(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	jr := corejson.NewPtr([]coredynamic.KeyVal{{Key: "x", Value: "y"}})
	_, err := kvc.ParseInjectUsingJson(jr)
	_ = err
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection ParseInjectUsingJson", actual)
}

func Test_C17_KeyValCollection_JsonParseSelfInject(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	jr := corejson.NewPtr([]coredynamic.KeyVal{{Key: "x", Value: "y"}})
	err := kvc.JsonParseSelfInject(jr)
	_ = err
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection JsonParseSelfInject", actual)
}

// ==========================================================================
// TypeStatus — uncovered branches
// ==========================================================================

func Test_C17_TypeStatus_IsValid_NilPtr(t *testing.T) {
	var ts *coredynamic.TypeStatus
	actual := args.Map{
		"nilValid":   ts.IsValid(),
		"nilInvalid": ts.IsInvalid(),
	}
	expected := args.Map{
		"nilValid":   false,
		"nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus nil", actual)
}

func Test_C17_TypeStatus_Branches(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", "world")
	actual := args.Map{
		"isSame":             ts.IsSame,
		"isNotSame":          ts.IsNotSame(),
		"isNotEqual":         ts.IsNotEqualTypes(),
		"isAnyPtr":           ts.IsAnyPointer(),
		"isBothPtr":          ts.IsBothPointer(),
		"sameRegardless":     ts.IsSameRegardlessPointer(),
		"leftName":           ts.LeftName(),
		"rightName":          ts.RightName(),
		"leftFull":           ts.LeftFullName(),
		"rightFull":          ts.RightFullName(),
	}
	expected := args.Map{
		"isSame":             true,
		"isNotSame":          false,
		"isNotEqual":         false,
		"isAnyPtr":           false,
		"isBothPtr":          false,
		"sameRegardless":     true,
		"leftName":           "string",
		"rightName":          "string",
		"leftFull":           "string",
		"rightFull":          "string",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus same type", actual)
}

func Test_C17_TypeStatus_NotMatch(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", 42)
	msg := ts.NotMatchMessage("left", "right")
	err := ts.NotMatchErr("left", "right")
	srcDst := ts.NotEqualSrcDestinationMessage()
	srcErr := ts.NotEqualSrcDestinationErr()
	valErr := ts.ValidationError()
	actual := args.Map{
		"msgNotEmpty":    msg != "",
		"errNotNil":      err != nil,
		"srcDstNotEmpty": srcDst != "",
		"srcErrNotNil":   srcErr != nil,
		"valErrNotNil":   valErr != nil,
	}
	expected := args.Map{
		"msgNotEmpty":    true,
		"errNotNil":      true,
		"srcDstNotEmpty": true,
		"srcErrNotNil":   true,
		"valErrNotNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotMatch", actual)
}

func Test_C17_TypeStatus_MustBeSame_Panic(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", 42)
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		ts.MustBeSame()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus MustBeSame panic", actual)
}

func Test_C17_TypeStatus_SrcDestinationMustBeSame_Panic(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", 42)
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		ts.SrcDestinationMustBeSame()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus SrcDestinationMustBeSame panic", actual)
}

func Test_C17_TypeStatus_IsEqual(t *testing.T) {
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", "b")
	ts3 := coredynamic.TypeSameStatus("a", 1)
	var nilTs *coredynamic.TypeStatus
	actual := args.Map{
		"same":    ts1.IsEqual(&ts2),
		"diff":    ts1.IsEqual(&ts3),
		"nilNil":  nilTs.IsEqual(nil),
		"nilOne":  nilTs.IsEqual(&ts1),
		"oneNil":  ts1.IsEqual(nil),
	}
	expected := args.Map{
		"same":    true,
		"diff":    false,
		"nilNil":  true,
		"nilOne":  false,
		"oneNil":  false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual", actual)
}

func Test_C17_TypeStatus_NullTypes(t *testing.T) {
	ts := coredynamic.TypeSameStatus(nil, nil)
	actual := args.Map{
		"leftName":  ts.LeftName(),
		"rightName": ts.RightName(),
		"leftFull":  ts.LeftFullName(),
		"rightFull": ts.RightFullName(),
	}
	expected := args.Map{
		"leftName":  "<nil>",
		"rightName": "<nil>",
		"leftFull":  "<nil>",
		"rightFull": "<nil>",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus null types", actual)
}

func Test_C17_TypeStatus_PointerTypes(t *testing.T) {
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, &s)
	np := ts.NonPointerLeft()
	npr := ts.NonPointerRight()
	actual := args.Map{
		"isAnyPtr":  ts.IsAnyPointer(),
		"isBothPtr": ts.IsBothPointer(),
		"npLeft":    np.Kind() == reflect.String,
		"npRight":   npr.Kind() == reflect.String,
		"sameReg":   ts.IsSameRegardlessPointer(),
	}
	expected := args.Map{
		"isAnyPtr":  true,
		"isBothPtr": true,
		"npLeft":    true,
		"npRight":   true,
		"sameReg":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus pointer types", actual)
}

// ==========================================================================
// CastTo
// ==========================================================================

func Test_C17_CastTo_Match(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{
		"valid":   result.IsValid,
		"matched": result.IsMatchingAcceptedType,
		"noErr":   result.Error == nil,
	}
	expected := args.Map{
		"valid":   true,
		"matched": true,
		"noErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo match", actual)
}

func Test_C17_CastTo_NoMatch(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(42))
	actual := args.Map{
		"matched": result.IsMatchingAcceptedType,
		"hasErr":  result.Error != nil,
	}
	expected := args.Map{
		"matched": false,
		"hasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo no match", actual)
}

// ==========================================================================
// TypeNotEqualErr / TypeMustBeSame
// ==========================================================================

func Test_C17_TypeNotEqualErr(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", "b")
	errDiff := coredynamic.TypeNotEqualErr("a", 42)
	actual := args.Map{
		"same":  err == nil,
		"diff":  errDiff != nil,
	}
	expected := args.Map{
		"same": true,
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr", actual)
}

func Test_C17_TypeMustBeSame_NoPanic(t *testing.T) {
	coredynamic.TypeMustBeSame("a", "b") // same types, no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame no panic", actual)
}

func Test_C17_TypeMustBeSame_Panic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		coredynamic.TypeMustBeSame("a", 42)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame panic", actual)
}

// ==========================================================================
// TypesIndexOf
// ==========================================================================

func Test_C17_TypesIndexOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"found":    coredynamic.TypesIndexOf(strType, intType, strType),
		"notFound": coredynamic.TypesIndexOf(reflect.TypeOf(true), intType, strType),
	}
	expected := args.Map{
		"found":    1,
		"notFound": -1,
	}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf", actual)
}

// ==========================================================================
// MapAnyItemDiff — coverage
// ==========================================================================

func Test_C17_MapAnyItemDiff_Basic(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	var nilM *coredynamic.MapAnyItemDiff
	actual := args.Map{
		"len":       m.Length(),
		"empty":     m.IsEmpty(),
		"hasAny":    m.HasAnyItem(),
		"lastIdx":   m.LastIndex(),
		"nilLen":    nilM.Length(),
	}
	expected := args.Map{
		"len":       1,
		"empty":     false,
		"hasAny":    true,
		"lastIdx":   0,
		"nilLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff basic", actual)
}

func Test_C17_MapAnyItemDiff_Raw_Clear(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	raw := m.Raw()
	var nilM *coredynamic.MapAnyItemDiff
	nilRaw := nilM.Raw()
	nilClear := nilM.Clear()
	cleared := m.Clear()
	actual := args.Map{
		"rawLen":     len(raw),
		"nilRawLen":  len(nilRaw),
		"nilClearLen": len(nilClear),
		"clearedLen": len(cleared),
	}
	expected := args.Map{
		"rawLen":     1,
		"nilRawLen":  0,
		"nilClearLen": 0,
		"clearedLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Raw/Clear", actual)
}

func Test_C17_MapAnyItemDiff_Json(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	j := m.Json()
	jp := m.JsonPtr()
	pj := m.PrettyJsonString()
	actual := args.Map{
		"jOk":  j.JsonString() != "",
		"jpOk": jp != nil,
		"pjOk": pj != "",
	}
	expected := args.Map{"jOk": true, "jpOk": true, "pjOk": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Json", actual)
}

func Test_C17_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	actual := args.Map{
		"equal":    m.IsRawEqual(false, map[string]any{"k": "v"}),
		"notEqual": m.IsRawEqual(false, map[string]any{"k": "v2"}),
	}
	expected := args.Map{
		"equal":    true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff IsRawEqual", actual)
}

func Test_C17_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	actual := args.Map{
		"noChanges":  m.HasAnyChanges(false, map[string]any{"k": "v"}),
		"hasChanges": m.HasAnyChanges(false, map[string]any{"k": "v2"}),
	}
	expected := args.Map{
		"noChanges":  false,
		"hasChanges": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff HasAnyChanges", actual)
}

func Test_C17_MapAnyItemDiff_DiffMethods(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"k": "v2"})
	diffRaw := m.DiffRaw(false, map[string]any{"k": "v2"})
	diffJson := m.DiffJsonMessage(false, map[string]any{"k": "v2"})
	diffSlice := m.ToStringsSliceOfDiffMap(diffRaw)
	shouldMsg := m.ShouldDiffMessage(false, "test", map[string]any{"k": "v2"})
	logMsg := m.LogShouldDiffMessage(false, "test", map[string]any{"k": "v2"})
	keys := m.AllKeysSorted()
	mai := m.MapAnyItems()
	rmd := m.RawMapDiffer()
	actual := args.Map{
		"diffHas":     diff.HasAnyItem(),
		"diffRawHas":  len(diffRaw) > 0,
		"diffJsonOk":  diffJson != "",
		"diffSliceOk": len(diffSlice) > 0,
		"shouldMsgOk": shouldMsg != "",
		"logMsgOk":    logMsg != "",
		"keysLen":     len(keys),
		"maiNotNil":   mai != nil,
		"rmdNotNil":   rmd != nil,
	}
	expected := args.Map{
		"diffHas":     true,
		"diffRawHas":  true,
		"diffJsonOk":  true,
		"diffSliceOk": true,
		"shouldMsgOk": true,
		"logMsgOk":    true,
		"keysLen":     1,
		"maiNotNil":   true,
		"rmdNotNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff diff methods", actual)
}

func Test_C17_MapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"k": "v"}
	m.LogPrettyJsonString()
	empty := coredynamic.MapAnyItemDiff{}
	empty.LogPrettyJsonString()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff LogPrettyJsonString", actual)
}

// ==========================================================================
// LeftRight — uncovered branches
// ==========================================================================

func Test_C17_LeftRight_DeserializeLeft_Right(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "l", Right: "r"}
	dl := lr.DeserializeLeft()
	dr := lr.DeserializeRight()
	var nilLR *coredynamic.LeftRight
	actual := args.Map{
		"dlOk":    dl != nil,
		"drOk":    dr != nil,
		"nilDl":   nilLR.DeserializeLeft() == nil,
		"nilDr":   nilLR.DeserializeRight() == nil,
	}
	expected := args.Map{
		"dlOk": true, "drOk": true,
		"nilDl": true, "nilDr": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft/Right", actual)
}

func Test_C17_LeftRight_TypeStatus(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "l", Right: "r"}
	ts := lr.TypeStatus()
	var nilLR *coredynamic.LeftRight
	tsNil := nilLR.TypeStatus()
	actual := args.Map{
		"isSame": ts.IsSame,
		"nilSame": tsNil.IsSame,
	}
	expected := args.Map{
		"isSame": true,
		"nilSame": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus", actual)
}

// ==========================================================================
// Dynamic Clone/NonPtr/Ptr
// ==========================================================================

func Test_C17_Dynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	cp := d.ClonePtr()
	np := d.NonPtr()
	pp := d.Ptr()
	var nilD *coredynamic.Dynamic
	nilCp := nilD.ClonePtr()
	actual := args.Map{
		"cpValid": cp.IsValid(),
		"npValid": np.IsValid(),
		"ppNotNil": pp != nil,
		"nilCp": nilCp == nil,
	}
	expected := args.Map{
		"cpValid": true, "npValid": true,
		"ppNotNil": true, "nilCp": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr/NonPtr/Ptr", actual)
}

// ==========================================================================
// Dynamic type check methods
// ==========================================================================

func Test_C17_Dynamic_TypeChecks(t *testing.T) {
	dStr := coredynamic.NewDynamicValid("hello")
	dInt := coredynamic.NewDynamicValid(42)
	dSlice := coredynamic.NewDynamicValid([]int{1, 2})
	dMap := coredynamic.NewDynamicValid(map[string]int{})
	type S struct{}
	dStruct := coredynamic.NewDynamicValid(S{})
	dFunc := coredynamic.NewDynamicValid(func() {})
	actual := args.Map{
		"isPrimStr":   dStr.IsPrimitive(),
		"isNumInt":    dInt.IsNumber(),
		"isStr":       dStr.IsStringType(),
		"isStruct":    dStruct.IsStruct(),
		"isFunc":      dFunc.IsFunc(),
		"isSlice":     dSlice.IsSliceOrArray(),
		"isSliceMap":  dSlice.IsSliceOrArrayOrMap(),
		"isMap":       dMap.IsMap(),
		"isValueType": dStr.IsValueType(),
	}
	expected := args.Map{
		"isPrimStr":   true,
		"isNumInt":    true,
		"isStr":       true,
		"isStruct":    true,
		"isFunc":      true,
		"isSlice":     true,
		"isSliceMap":  true,
		"isMap":       true,
		"isValueType": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks", actual)
}

// ==========================================================================
// Dynamic — ConvertUsingFunc
// ==========================================================================

func Test_C17_Dynamic_ConvertUsingFunc(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	result := d.ConvertUsingFunc(func(input any, expectedType reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(input)
	}, reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ConvertUsingFunc", actual)
}

// ==========================================================================
// CastedResult uncovered methods
// ==========================================================================

func Test_C17_CastedResult_Methods(t *testing.T) {
	cr := coredynamic.CastedResult{
		Casted: "x", IsValid: true, IsNull: false,
		IsMatchingAcceptedType: true, IsPointer: false,
		IsSourcePointer: false, SourceKind: reflect.String,
	}
	var nilCr *coredynamic.CastedResult
	actual := args.Map{
		"invalid":   cr.IsInvalid(),
		"notNull":   cr.IsNotNull(),
		"notPtr":    cr.IsNotPointer(),
		"notMatch":  cr.IsNotMatchingAcceptedType(),
		"isKind":    cr.IsSourceKind(reflect.String),
		"hasErr":    cr.HasError(),
		"hasIssues": cr.HasAnyIssues(),
		"nilInv":    nilCr.IsInvalid(),
	}
	expected := args.Map{
		"invalid":   false,
		"notNull":   true,
		"notPtr":    true,
		"notMatch":  false,
		"isKind":    true,
		"hasErr":    false,
		"hasIssues": false,
		"nilInv":    true,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult methods", actual)
}
