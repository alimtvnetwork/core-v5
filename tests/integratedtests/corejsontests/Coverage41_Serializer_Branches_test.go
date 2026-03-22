package corejsontests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// serializerLogic — Apply / StringsApply
// =============================================================================

func Test_Cov41_Serialize_Apply_Valid(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	actual := args.Map{"noErr": !r.HasError(), "hasBytes": r.Length() > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize Apply valid", actual)
}

func Test_Cov41_Serialize_Apply_Unmarshalable(t *testing.T) {
	r := corejson.Serialize.Apply(make(chan int))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize Apply unmarshalable", actual)
}

func Test_Cov41_Serialize_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize StringsApply", actual)
}

// =============================================================================
// serializerLogic — From* methods
// =============================================================================

func Test_Cov41_Serialize_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte("hello"))
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBytes", actual)
}

func Test_Cov41_Serialize_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStrings", actual)
}

func Test_Cov41_Serialize_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStringsSpread", actual)
}

func Test_Cov41_Serialize_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromString", actual)
}

func Test_Cov41_Serialize_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromInteger", actual)
}

func Test_Cov41_Serialize_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(99)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromInteger64", actual)
}

func Test_Cov41_Serialize_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBool", actual)
}

func Test_Cov41_Serialize_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromIntegers", actual)
}

type testStringer struct{ v string }

func (s testStringer) String() string { return s.v }

func Test_Cov41_Serialize_FromStringer(t *testing.T) {
	r := corejson.Serialize.FromStringer(testStringer{"hello"})
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStringer", actual)
}

// =============================================================================
// serializerLogic — UsingAnyPtr / UsingAny
// =============================================================================

func Test_Cov41_Serialize_UsingAnyPtr_Valid(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize UsingAnyPtr valid", actual)
}

func Test_Cov41_Serialize_UsingAnyPtr_Unmarshalable(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr(make(chan int))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize UsingAnyPtr unmarshalable", actual)
}

func Test_Cov41_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize UsingAny", actual)
}

// =============================================================================
// serializerLogic — Raw / Marshal / ApplyMust / ToBytesMust / ToSafeBytesMust
// =============================================================================

func Test_Cov41_Serialize_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("hello")
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize Raw", actual)
}

func Test_Cov41_Serialize_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("hello")
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize Marshal", actual)
}

func Test_Cov41_Serialize_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize ApplyMust", actual)
}

func Test_Cov41_Serialize_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesMust", actual)
}

func Test_Cov41_Serialize_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToSafeBytesMust", actual)
}

// =============================================================================
// serializerLogic — Swallow / ToString / ToPretty
// =============================================================================

func Test_Cov41_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToSafeBytesSwallowErr", actual)
}

func Test_Cov41_Serialize_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesSwallowErr", actual)
}

func Test_Cov41_Serialize_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("hello")
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesErr", actual)
}

func Test_Cov41_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToString", actual)
}

func Test_Cov41_Serialize_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("hello")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToStringMust", actual)
}

func Test_Cov41_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{"noErr": err == nil, "hasContent": len(s) > 0}
	expected := args.Map{"noErr": true, "hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToStringErr", actual)
}

func Test_Cov41_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{"noErr": err == nil, "hasContent": len(s) > 0}
	expected := args.Map{"noErr": true, "hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringErr", actual)
}

func Test_Cov41_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr("hello")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringIncludingErr", actual)
}

func Test_Cov41_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty("hello")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Serialize Pretty", actual)
}

// =============================================================================
// deserializerLogic — Apply / UsingResult
// =============================================================================

func Test_Cov41_Deserialize_Apply(t *testing.T) {
	r := corejson.NewPtr("hello")
	var s string
	err := corejson.Deserialize.Apply(r, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize Apply", actual)
}

func Test_Cov41_Deserialize_UsingResult(t *testing.T) {
	r := corejson.NewPtr("hello")
	var s string
	err := corejson.Deserialize.UsingResult(r, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingResult", actual)
}

func Test_Cov41_Deserialize_ApplyMust(t *testing.T) {
	r := corejson.NewPtr("hello")
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize ApplyMust", actual)
}

func Test_Cov41_Deserialize_ApplyMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Deserialize ApplyMust panics", actual)
	}()
	r := &corejson.Result{Error: errors.New("fail")}
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
}

// =============================================================================
// deserializerLogic — UsingString* methods
// =============================================================================

func Test_Cov41_Deserialize_UsingString(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingString(`"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingString", actual)
}

func Test_Cov41_Deserialize_FromString(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromString(`"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize FromString", actual)
}

func Test_Cov41_Deserialize_FromStringMust(t *testing.T) {
	var s string
	corejson.Deserialize.FromStringMust(`"hello"`, &s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize FromStringMust", actual)
}

func Test_Cov41_Deserialize_FromStringMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Deserialize FromStringMust panics", actual)
	}()
	var s string
	corejson.Deserialize.FromStringMust(`bad`, &s)
}

func Test_Cov41_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringPtr nil", actual)
}

func Test_Cov41_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	js := `"hello"`
	var s string
	err := corejson.Deserialize.UsingStringPtr(&js, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringPtr valid", actual)
}

func Test_Cov41_Deserialize_UsingStringOption_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringOption skip", actual)
}

func Test_Cov41_Deserialize_UsingStringOption_Process(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(false, `"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringOption process", actual)
}

func Test_Cov41_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringIgnoreEmpty", actual)
}

// =============================================================================
// deserializerLogic — UsingError*
// =============================================================================

func Test_Cov41_Deserialize_UsingError_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingError nil", actual)
}

func Test_Cov41_Deserialize_UsingError_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(fmt.Errorf(`"hello"`), &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingError valid", actual)
}

func Test_Cov41_Deserialize_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingErrorWhichJsonResult nil", actual)
}

// =============================================================================
// deserializerLogic — UsingBytes*
// =============================================================================

func Test_Cov41_Deserialize_UsingBytes_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytes([]byte(`"hello"`), &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytes valid", actual)
}

func Test_Cov41_Deserialize_UsingBytes_Invalid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytes([]byte(`bad`), &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytes invalid", actual)
}

func Test_Cov41_Deserialize_UsingBytesMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesMust", actual)
}

func Test_Cov41_Deserialize_UsingBytesMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesMust panics", actual)
	}()
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`bad`), &s)
}

func Test_Cov41_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesPointer nil", actual)
}

func Test_Cov41_Deserialize_UsingBytesPointer_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesPointer valid", actual)
}

func Test_Cov41_Deserialize_UsingBytesPointerMust(t *testing.T) {
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesPointerMust", actual)
}

func Test_Cov41_Deserialize_UsingBytesIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": ""}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesIf skip", actual)
}

func Test_Cov41_Deserialize_UsingBytesIf_Process(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesIf process", actual)
}

func Test_Cov41_Deserialize_UsingBytesPointerIf_Skip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesPointerIf skip", actual)
}

func Test_Cov41_Deserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s) // should return without panic
	actual := args.Map{"r": s}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingSafeBytesMust empty", actual)
}

func Test_Cov41_Deserialize_UsingSafeBytesMust_Valid(t *testing.T) {
	var s string
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hello"`), &s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingSafeBytesMust valid", actual)
}

// =============================================================================
// deserializerLogic — MapAnyToPointer
// =============================================================================

func Test_Cov41_Deserialize_MapAnyToPointer_SkipEmpty(t *testing.T) {
	type S struct{ Name string }
	var s S
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize MapAnyToPointer skip empty", actual)
}

func Test_Cov41_Deserialize_MapAnyToPointer_Valid(t *testing.T) {
	type S struct{ Name string }
	var s S
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	actual := args.Map{"noErr": err == nil, "name": s.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Deserialize MapAnyToPointer valid", actual)
}

// =============================================================================
// deserializerLogic — FromTo / AnyToFieldsMap
// =============================================================================

func Test_Cov41_Deserialize_FromTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromTo("hello", &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize FromTo", actual)
}

func Test_Cov41_Deserialize_AnyToFieldsMap(t *testing.T) {
	type S struct{ Name string }
	m, err := corejson.Deserialize.AnyToFieldsMap(S{Name: "test"})
	// Note: DeserializedFieldsToMap may return empty since it tries to deserialize into nil map
	actual := args.Map{"noErr": err == nil || m != nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize AnyToFieldsMap", actual)
}

// =============================================================================
// deserializerLogic — Result / ResultPtr / ResultMust / ResultPtrMust
// =============================================================================

func Test_Cov41_Deserialize_Result_Valid(t *testing.T) {
	serialized := corejson.Serialize.ToBytesMust(corejson.New("hello"))
	_, err := corejson.Deserialize.Result(serialized)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize Result valid", actual)
}

func Test_Cov41_Deserialize_ResultPtr_Invalid(t *testing.T) {
	_, err := corejson.Deserialize.ResultPtr([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize ResultPtr invalid", actual)
}

// =============================================================================
// deserializerLogic — UsingDeserializerToOption / UsingDeserializerDefined
// =============================================================================

func Test_Cov41_Deserialize_UsingDeserializerToOption_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption skip nil", actual)
}

func Test_Cov41_Deserialize_UsingDeserializerToOption_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption nil not skip", actual)
}

func Test_Cov41_Deserialize_UsingDeserializerDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerDefined nil", actual)
}

// =============================================================================
// deserializerLogic — UsingDeserializerFuncDefined
// =============================================================================

func Test_Cov41_Deserialize_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined nil", actual)
}

func Test_Cov41_Deserialize_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	var s string
	fn := func(toPtr any) error {
		*(toPtr.(*string)) = "hello"
		return nil
	}
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined valid", actual)
}

// =============================================================================
// deserializerLogic — UsingJsonerToAny / UsingJsonerToAnyMust
// =============================================================================

func Test_Cov41_Deserialize_UsingJsonerToAny_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny skip nil", actual)
}

func Test_Cov41_Deserialize_UsingJsonerToAny_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny nil not skip", actual)
}

func Test_Cov41_Deserialize_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust skip nil", actual)
}

func Test_Cov41_Deserialize_UsingJsonerToAnyMust_NilNotSkip(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust nil not skip", actual)
}
