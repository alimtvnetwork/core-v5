package coredynamictests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// DynamicReflect — ReflectValue / ReflectType / ReflectKind
// =============================================================================

func Test_Cov41_DynReflect_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	rv := d.ReflectValue()
	actual := args.Map{"notNil": rv != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectValue", actual)
}

func Test_Cov41_DynReflect_ReflectValue_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()
	actual := args.Map{"same": rv1 == rv2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectValue cached", actual)
}

func Test_Cov41_DynReflect_ReflectKind(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"r": d.ReflectKind()}
	expected := args.Map{"r": reflect.String}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectKind", actual)
}

func Test_Cov41_DynReflect_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"nonEmpty": len(d.ReflectTypeName()) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectTypeName", actual)
}

func Test_Cov41_DynReflect_ReflectType(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	rt := d.ReflectType()
	actual := args.Map{"name": rt.Name()}
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectType", actual)
}

func Test_Cov41_DynReflect_ReflectType_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	rt1 := d.ReflectType()
	rt2 := d.ReflectType()
	actual := args.Map{"same": rt1 == rt2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectType cached", actual)
}

func Test_Cov41_DynReflect_IsReflectTypeOf_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"r": d.IsReflectTypeOf(reflect.TypeOf(0))}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsReflectTypeOf true", actual)
}

func Test_Cov41_DynReflect_IsReflectTypeOf_False(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"r": d.IsReflectTypeOf(reflect.TypeOf(""))}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsReflectTypeOf false", actual)
}

func Test_Cov41_DynReflect_IsReflectKind_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"r": d.IsReflectKind(reflect.String)}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IsReflectKind true", actual)
}

func Test_Cov41_DynReflect_IsReflectKind_False(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"r": d.IsReflectKind(reflect.Int)}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IsReflectKind false", actual)
}

// =============================================================================
// DynamicReflect — Index/Key access
// =============================================================================

func Test_Cov41_DynReflect_ItemReflectValueUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{10, 20, 30}, true)
	rv := d.ItemReflectValueUsingIndex(1)
	actual := args.Map{"val": int(rv.Int())}
	expected := args.Map{"val": 20}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingIndex", actual)
}

func Test_Cov41_DynReflect_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]string{"a", "b"}, true)
	actual := args.Map{"val": d.ItemUsingIndex(0)}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingIndex", actual)
}

func Test_Cov41_DynReflect_ItemReflectValueUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"x": 42}, true)
	rv := d.ItemReflectValueUsingKey("x")
	actual := args.Map{"val": int(rv.Int())}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingKey", actual)
}

func Test_Cov41_DynReflect_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]string{"k": "v"}, true)
	actual := args.Map{"val": d.ItemUsingKey("k")}
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingKey", actual)
}

// =============================================================================
// DynamicReflect — ReflectSetTo
// =============================================================================

func Test_Cov41_DynReflect_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectSetTo nil", actual)
}

// =============================================================================
// DynamicReflect — Loop
// =============================================================================

func Test_Cov41_DynReflect_Loop_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	called := d.Loop(func(i int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop invalid", actual)
}

func Test_Cov41_DynReflect_Loop_Nil(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.Loop(func(i int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop nil", actual)
}

func Test_Cov41_DynReflect_Loop_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	called := d.Loop(func(i int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop valid", actual)
}

func Test_Cov41_DynReflect_Loop_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	called := d.Loop(func(i int, item any) bool {
		count++
		return true
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop break", actual)
}

// =============================================================================
// DynamicReflect — FilterAsDynamicCollection
// =============================================================================

func Test_Cov41_DynReflect_FilterAsDynamicCollection_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	r := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) { return true, false })
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection invalid", actual)
}

func Test_Cov41_DynReflect_FilterAsDynamicCollection_TakeAll(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	r := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) { return true, false })
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection take all", actual)
}

func Test_Cov41_DynReflect_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	r := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 0
	})
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection break", actual)
}

func Test_Cov41_DynReflect_FilterAsDynamicCollection_Skip(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	r := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return i != 1, false
	})
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection skip", actual)
}

// =============================================================================
// DynamicReflect — LoopMap
// =============================================================================

func Test_Cov41_DynReflect_LoopMap_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap invalid", actual)
}

func Test_Cov41_DynReflect_LoopMap_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	count := 0
	called := d.LoopMap(func(i int, k, v any) bool {
		count++
		return false
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap valid", actual)
}

func Test_Cov41_DynReflect_LoopMap_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	count := 0
	called := d.LoopMap(func(i int, k, v any) bool {
		count++
		return true
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap break", actual)
}

// =============================================================================
// DynamicReflect — MapToKeyVal
// =============================================================================

func Test_Cov41_DynReflect_MapToKeyVal_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	kvc, err := d.MapToKeyVal()
	actual := args.Map{"noErr": err == nil, "notNil": kvc != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic MapToKeyVal valid", actual)
}

// =============================================================================
// DynamicJson — nil receiver branches
// =============================================================================

func Test_Cov41_DynJson_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`"test"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Deserialize nil", actual)
}

func Test_Cov41_DynJson_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueMarshal nil", actual)
}

func Test_Cov41_DynJson_ValueMarshal_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.ValueMarshal()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueMarshal valid", actual)
}

func Test_Cov41_DynJson_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b := d.JsonPayloadMust()
	actual := args.Map{"nonEmpty": len(b) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonPayloadMust", actual)
}

func Test_Cov41_DynJson_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"noErr": err == nil, "emptyBytes": len(b) == 0}
	expected := args.Map{"noErr": true, "emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonBytesPtr null", actual)
}

func Test_Cov41_DynJson_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonBytesPtr valid", actual)
}

func Test_Cov41_DynJson_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	b, err := d.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic MarshalJSON", actual)
}

func Test_Cov41_DynJson_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`"test"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic UnmarshalJSON nil", actual)
}

func Test_Cov41_DynJson_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	actual := args.Map{"notNil": d.JsonModel() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonModel", actual)
}

func Test_Cov41_DynJson_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	actual := args.Map{"notNil": d.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonModelAny", actual)
}

func Test_Cov41_DynJson_Json(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	r := d.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Json", actual)
}

func Test_Cov41_DynJson_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	r := d.JsonPtr()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonPtr", actual)
}

func Test_Cov41_DynJson_ParseInjectUsingJson_Error(t *testing.T) {
	d := coredynamic.NewDynamicPtr("x", true)
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := d.ParseInjectUsingJson(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ParseInjectUsingJson error", actual)
}

func Test_Cov41_DynJson_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	d := coredynamic.NewDynamicPtr("x", true)
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		d.ParseInjectUsingJsonMust(jr)
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ParseInjectUsingJsonMust panics", actual)
}

func Test_Cov41_DynJson_JsonParseSelfInject_Error(t *testing.T) {
	d := coredynamic.NewDynamicPtr("x", true)
	jr := &corejson.Result{Error: errors.New("fail")}
	err := d.JsonParseSelfInject(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonParseSelfInject error", actual)
}

func Test_Cov41_DynJson_JsonBytes_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.JsonBytes()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonBytes valid", actual)
}

func Test_Cov41_DynJson_JsonString_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	s, err := d.JsonString()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(s) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonString valid", actual)
}

func Test_Cov41_DynJson_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	s := d.JsonStringMust()
	actual := args.Map{"nonEmpty": len(s) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonStringMust", actual)
}
