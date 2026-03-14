package convertinternaltests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// ── AnyTo extended ──

func Test_Cov3_AnyTo_FullPropertyString(t *testing.T) {
	actual := args.Map{
		"int":    convertinternal.AnyTo.FullPropertyString(42) != "",
		"nil":    convertinternal.AnyTo.FullPropertyString(nil),
		"string": convertinternal.AnyTo.FullPropertyString("hi") != "",
	}
	expected := args.Map{"int": true, "nil": "", "string": true}
	expected.ShouldBeEqual(t, 0, "AnyTo FullPropertyString", actual)
}

func Test_Cov3_AnyTo_TypeName(t *testing.T) {
	actual := args.Map{
		"int":    convertinternal.AnyTo.TypeName(42) != "",
		"nil":    convertinternal.AnyTo.TypeName(nil),
		"string": convertinternal.AnyTo.TypeName("hi") != "",
	}
	expected := args.Map{"int": true, "nil": "", "string": true}
	expected.ShouldBeEqual(t, 0, "AnyTo TypeName", actual)
}

func Test_Cov3_AnyTo_SmartString_Error(t *testing.T) {
	err := errors.New("test error")
	var nilErr error
	actual := args.Map{
		"error":   convertinternal.AnyTo.SmartString(err),
		"nilErr":  convertinternal.AnyTo.SmartString(nilErr),
	}
	expected := args.Map{"error": "test error", "nilErr": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartString error", actual)
}

func Test_Cov3_AnyTo_SmartString_Stringer(t *testing.T) {
	type myStringer struct{ val string }
	// Use a type that implements fmt.Stringer
	result := convertinternal.AnyTo.SmartString(fmt.Errorf("stringer"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartString stringer", actual)
}

func Test_Cov3_AnyTo_SmartString_StringSlice(t *testing.T) {
	result := convertinternal.AnyTo.SmartString([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartString string slice", actual)
}

func Test_Cov3_AnyTo_SmartString_AnySlice(t *testing.T) {
	result := convertinternal.AnyTo.SmartString([]any{"a", 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartString any slice", actual)
}

func Test_Cov3_AnyTo_SmartString_EmptyAnySlice(t *testing.T) {
	result := convertinternal.AnyTo.SmartString([]any{})
	actual := args.Map{"empty": result}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartString empty any slice", actual)
}

func Test_Cov3_AnyTo_SmartJson_Error(t *testing.T) {
	err := errors.New("json err")
	var nilErr error
	actual := args.Map{
		"error":  convertinternal.AnyTo.SmartJson(err),
		"nilErr": convertinternal.AnyTo.SmartJson(nilErr),
	}
	expected := args.Map{"error": "json err", "nilErr": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartJson error", actual)
}

func Test_Cov3_AnyTo_SmartJson_StringSlice(t *testing.T) {
	result := convertinternal.AnyTo.SmartJson([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartJson string slice", actual)
}

func Test_Cov3_AnyTo_SmartJson_Bool(t *testing.T) {
	result := convertinternal.AnyTo.SmartJson(true)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "true"}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartJson bool", actual)
}

func Test_Cov3_AnyTo_SmartJson_Struct(t *testing.T) {
	type s struct{ A int }
	result := convertinternal.AnyTo.SmartJson(s{A: 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartJson struct", actual)
}

func Test_Cov3_AnyTo_SmartPrettyJsonLines(t *testing.T) {
	strResult := convertinternal.AnyTo.SmartPrettyJsonLines("a\nb")
	sliceResult := convertinternal.AnyTo.SmartPrettyJsonLines([]string{"a", "b"})
	nilResult := convertinternal.AnyTo.SmartPrettyJsonLines(nil)
	structResult := convertinternal.AnyTo.SmartPrettyJsonLines(map[string]int{"a": 1})
	actual := args.Map{
		"strLen":    len(strResult),
		"sliceLen":  len(sliceResult),
		"nilLen":    len(nilResult),
		"structGt0": len(structResult) > 0,
	}
	expected := args.Map{"strLen": 2, "sliceLen": 2, "nilLen": 0, "structGt0": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SmartPrettyJsonLines", actual)
}

func Test_Cov3_AnyTo_PrettyJsonLines(t *testing.T) {
	result := convertinternal.AnyTo.PrettyJsonLines(map[string]int{"a": 1})
	nilResult := convertinternal.AnyTo.PrettyJsonLines(nil)
	actual := args.Map{"gt0": len(result) > 0, "nilLen": len(nilResult)}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyJsonLines", actual)
}

func Test_Cov3_AnyTo_Strings_MapAny(t *testing.T) {
	result := convertinternal.AnyTo.Strings(map[string]any{"a": 1})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings map[string]any", actual)
}

func Test_Cov3_AnyTo_Strings_MapAnyAny(t *testing.T) {
	result := convertinternal.AnyTo.Strings(map[any]any{"a": 1})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings map[any]any", actual)
}

func Test_Cov3_AnyTo_Strings_MapIntString(t *testing.T) {
	result := convertinternal.AnyTo.Strings(map[int]string{1: "a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings map[int]string", actual)
}

func Test_Cov3_AnyTo_Strings_MapStringInt(t *testing.T) {
	result := convertinternal.AnyTo.Strings(map[string]int{"a": 1})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings map[string]int", actual)
}

func Test_Cov3_AnyTo_Strings_AnySlice(t *testing.T) {
	result := convertinternal.AnyTo.Strings([]any{"a", 1})
	emptyResult := convertinternal.AnyTo.Strings([]any{})
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings any slice", actual)
}

func Test_Cov3_AnyTo_Strings_Error(t *testing.T) {
	err := errors.New("line1\nline2")
	var nilErr error
	result := convertinternal.AnyTo.Strings(err)
	nilResult := convertinternal.AnyTo.Strings(nilErr)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings error", actual)
}

func Test_Cov3_AnyTo_Strings_EmptyString(t *testing.T) {
	result := convertinternal.AnyTo.Strings("")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings empty string", actual)
}

func Test_Cov3_AnyTo_Strings_Int64Slice(t *testing.T) {
	result := convertinternal.AnyTo.Strings([]int64{1, 2})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings int64 slice", actual)
}

func Test_Cov3_AnyTo_Strings_Float64Slice(t *testing.T) {
	result := convertinternal.AnyTo.Strings([]float64{1.1, 2.2})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings float64 slice", actual)
}

func Test_Cov3_AnyTo_Strings_ByteSlice(t *testing.T) {
	result := convertinternal.AnyTo.Strings([]byte{1, 2, 3})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings byte slice", actual)
}

func Test_Cov3_AnyTo_Strings_Stringer(t *testing.T) {
	result := convertinternal.AnyTo.Strings(fmt.Errorf("a\nb"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings stringer", actual)
}

func Test_Cov3_AnyTo_Strings_Bool(t *testing.T) {
	result := convertinternal.AnyTo.Strings(true)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "true"}
	expected.ShouldBeEqual(t, 0, "AnyTo Strings bool", actual)
}

func Test_Cov3_AnyTo_String_NilPtr(t *testing.T) {
	var p *string
	result := convertinternal.AnyTo.String(p)
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo String nil ptr", actual)
}

func Test_Cov3_AnyTo_String_Error(t *testing.T) {
	err := errors.New("test")
	var nilErr error
	actual := args.Map{
		"error":  convertinternal.AnyTo.String(err),
		"nilErr": convertinternal.AnyTo.String(nilErr),
	}
	expected := args.Map{"error": "test", "nilErr": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo String error", actual)
}

// ── CodeFormatter ──

func Test_Cov3_CodeFormatter_GolangRaw(t *testing.T) {
	code := []byte("package main\nfunc main(){}")
	result, err := convertinternal.CodeFormat.GolangRaw(code)
	emptyResult, emptyErr := convertinternal.CodeFormat.GolangRaw([]byte{})
	actual := args.Map{
		"notEmpty": len(result) > 0, "noErr": err == nil,
		"emptyLen": len(emptyResult), "emptyNoErr": emptyErr == nil,
	}
	expected := args.Map{"notEmpty": true, "noErr": true, "emptyLen": 0, "emptyNoErr": true}
	expected.ShouldBeEqual(t, 0, "CodeFormatter GolangRaw", actual)
}

func Test_Cov3_CodeFormatter_Golang(t *testing.T) {
	result, err := convertinternal.CodeFormat.Golang("package main\nfunc main(){}")
	emptyResult, emptyErr := convertinternal.CodeFormat.Golang("")
	actual := args.Map{
		"notEmpty": result != "", "noErr": err == nil,
		"emptyResult": emptyResult, "emptyNoErr": emptyErr == nil,
	}
	expected := args.Map{"notEmpty": true, "noErr": true, "emptyResult": "", "emptyNoErr": true}
	expected.ShouldBeEqual(t, 0, "CodeFormatter Golang", actual)
}

func Test_Cov3_CodeFormatter_Golang_Invalid(t *testing.T) {
	_, err := convertinternal.CodeFormat.Golang("not valid go code {{{")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CodeFormatter Golang invalid", actual)
}

// ── Integers ──

func Test_Cov3_Integers_ToMapBool(t *testing.T) {
	result := convertinternal.Integers.ToMapBool(1, 2, 3)
	emptyResult := convertinternal.Integers.ToMapBool()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult), "has1": result[1]}
	expected := args.Map{"len": 3, "emptyLen": 0, "has1": true}
	expected.ShouldBeEqual(t, 0, "Integers ToMapBool", actual)
}

func Test_Cov3_Integers_Int8ToMapBool(t *testing.T) {
	result := convertinternal.Integers.Int8ToMapBool(1, 2)
	emptyResult := convertinternal.Integers.Int8ToMapBool()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "Integers Int8ToMapBool", actual)
}

func Test_Cov3_Integers_FromIntegersToMap(t *testing.T) {
	result := convertinternal.Integers.FromIntegersToMap(1, 2)
	emptyResult := convertinternal.Integers.FromIntegersToMap()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "Integers FromIntegersToMap", actual)
}

func Test_Cov3_Integers_IntegersToStrings(t *testing.T) {
	result := convertinternal.Integers.IntegersToStrings([]int{1, 2})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "1"}
	expected.ShouldBeEqual(t, 0, "Integers IntegersToStrings", actual)
}

// ── KeyValues ──

func Test_Cov3_KeyValues_ToMap(t *testing.T) {
	result := convertinternal.KeyValues.ToMap([]string{"a", "b"}, []string{"1", "2"})
	nilResult := convertinternal.KeyValues.ToMap(nil, nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult), "a": result["a"]}
	expected := args.Map{"len": 2, "nilLen": 0, "a": "1"}
	expected.ShouldBeEqual(t, 0, "KeyValues ToMap", actual)
}

func Test_Cov3_KeyValues_ToMapPtr(t *testing.T) {
	keys := []string{"a"}
	vals := []string{"1"}
	result := convertinternal.KeyValues.ToMapPtr(&keys, &vals)
	nilResult := convertinternal.KeyValues.ToMapPtr(nil, nil)
	actual := args.Map{"notNil": result != nil, "nilNotNil": nilResult != nil}
	expected := args.Map{"notNil": true, "nilNotNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValues ToMapPtr", actual)
}

// ── Map ──

func Test_Cov3_Map_Keys_StringString(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[string]string{"a": "1", "b": "2"})
	actual := args.Map{"len": len(keys), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Map Keys string string", actual)
}

func Test_Cov3_Map_Keys_StringAny(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[string]any{"a": 1})
	actual := args.Map{"len": len(keys), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Map Keys string any", actual)
}

func Test_Cov3_Map_Keys_IntAny(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[int]any{1: "a"})
	actual := args.Map{"len": len(keys), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Map Keys int any", actual)
}

func Test_Cov3_Map_Keys_Unsupported(t *testing.T) {
	_, err := convertinternal.Map.Keys("not a map")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Map Keys unsupported", actual)
}

func Test_Cov3_Map_KeysValues_StringString(t *testing.T) {
	keys, vals, err := convertinternal.Map.KeysValues(map[string]string{"a": "1"})
	actual := args.Map{"kLen": len(keys), "vLen": len(vals), "noErr": err == nil}
	expected := args.Map{"kLen": 1, "vLen": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Map KeysValues string string", actual)
}

func Test_Cov3_Map_KeysValues_StringAny(t *testing.T) {
	keys, vals, err := convertinternal.Map.KeysValues(map[string]any{"a": 1})
	actual := args.Map{"kLen": len(keys), "vLen": len(vals), "noErr": err == nil}
	expected := args.Map{"kLen": 1, "vLen": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Map KeysValues string any", actual)
}

func Test_Cov3_Map_KeysValues_Unsupported(t *testing.T) {
	_, _, err := convertinternal.Map.KeysValues("not a map")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Map KeysValues unsupported", actual)
}
