package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
)

// ===================== anyTo =====================

func TestAnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	if r == nil || r.Error == nil { t.Fatal("expected error") }
}

func TestAnyTo_SerializedJsonResult_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	out := corejson.AnyTo.SerializedJsonResult(r)
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	out := corejson.AnyTo.SerializedJsonResult(r)
	if out != r { t.Fatal("expected same ptr") }
}

func TestAnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_String(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_Error(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(errors.corejson.New("oops"))
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_AnyItem(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(42)
	if out == nil || out.HasError() { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
	_, err2 := corejson.AnyTo.SerializedString(nil)
	if err2 == nil { t.Fatal("expected error") }
}

func TestAnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	if s == "" { t.Fatal("expected non-empty") }
	s2 := corejson.AnyTo.SerializedSafeString(nil)
	if s2 != "" { t.Fatal("expected empty") }
}

func TestAnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" { t.Fatal("unexpected") }
	r := corejson.Result{Bytes: []byte(`{"a":1}`)}
	s3, err3 := corejson.AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" { t.Fatal("unexpected") }
	rp := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s4, err4 := corejson.AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" { t.Fatal("unexpected") }
	re := corejson.Result{Error: errors.corejson.New("e")}
	_, err5 := corejson.AnyTo.PrettyStringWithError(re)
	if err5 == nil { t.Fatal("expected error") }
	rep := &corejson.Result{Error: errors.corejson.New("e")}
	_, err6 := corejson.AnyTo.PrettyStringWithError(rep)
	if err6 == nil { t.Fatal("expected error") }
	s7, err7 := corejson.AnyTo.PrettyStringWithError(42)
	if err7 != nil || s7 == "" { t.Fatal("unexpected") }
}

func TestAnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = corejson.AnyTo.SafeJsonPrettyString("hello")
	_ = corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.Result{Bytes: []byte(`{"a":1}`)})
	_ = corejson.AnyTo.SafeJsonPrettyString(&corejson.Result{Bytes: []byte(`{"a":1}`)})
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func TestAnyTo_JsonString(t *testing.T) {
	_ = corejson.AnyTo.corejson.JsonString("hello")
	_ = corejson.AnyTo.corejson.JsonString([]byte(`"x"`))
	_ = corejson.AnyTo.corejson.JsonString(corejson.Result{Bytes: []byte(`"x"`)})
	_ = corejson.AnyTo.corejson.JsonString(&corejson.Result{Bytes: []byte(`"x"`)})
	_ = corejson.AnyTo.corejson.JsonString(42)
}

func TestAnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.Result{Bytes: []byte(`"x"`)})
	_, _ = corejson.AnyTo.JsonStringWithErr(&corejson.Result{Bytes: []byte(`"x"`)})
	_, err2 := corejson.AnyTo.JsonStringWithErr(corejson.Result{Error: errors.corejson.New("e")})
	if err2 == nil { t.Fatal("expected error") }
	_, err3 := corejson.AnyTo.JsonStringWithErr(&corejson.Result{Error: errors.corejson.New("e")})
	if err3 == nil { t.Fatal("expected error") }
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func TestAnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	if s != "hello" { t.Fatal("unexpected") }
}

func TestAnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	if s != "hello" { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedFieldsMap(t *testing.T) {
	type s struct{ Name string }
	fm, err := corejson.AnyTo.SerializedFieldsMap(s{Name: "x"})
	if err != nil { t.Fatal("unexpected error") }
	_ = fm
}

// ===================== castingAny =====================

func TestCastAny_FromToDefault(t *testing.T) {
	src := `"hello"`
	var dst string
	err := corejson.CastAny.FromToDefault([]byte(src), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_Bytes(t *testing.T) {
	var dst string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_String(t *testing.T) {
	var dst int
	err := corejson.CastAny.FromToOption(false, `42`, &dst)
	if err != nil || dst != 42 { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hello"`)}
	var dst string
	err := corejson.CastAny.FromToOption(false, r, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`42`)}
	var dst int
	err := corejson.CastAny.FromToOption(false, r, &dst)
	if err != nil || dst != 42 { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_AnyItem(t *testing.T) {
	type s struct{ N int }
	src := s{N: 42}
	var dst s
	err := corejson.CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.N != 42 { t.Fatal("unexpected") }
}

func TestCastAny_OrDeserializeTo(t *testing.T) {
	var dst string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

// ===================== deserializerLogic extras =====================

func TestDeserialize_UsingStringPtr_Nil(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingStringPtr(nil, &dst)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingStringPtr_Normal(t *testing.T) {
	s := `"hello"`
	var dst string
	err := corejson.Deserialize.UsingStringPtr(&s, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingError_Nil(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingError(nil, &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingErrorWhichJsonResult(t *testing.T) {
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_ApplyMust(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_FromString(t *testing.T) {
	var dst string
	err := corejson.Deserialize.FromString(`"hello"`, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_FromTo(t *testing.T) {
	var dst string
	err := corejson.Deserialize.FromTo([]byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_MapAnyToPointer(t *testing.T) {
	type s struct{ Name string }
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s{})
	if err != nil { t.Fatal("expected nil for empty skip") }
	err2 := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "x"}, &s{})
	if err2 != nil { t.Fatal("unexpected error") }
}

func TestDeserialize_UsingStringOption(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingStringOption(true, "", &dst)
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingStringOption(false, `"x"`, &dst)
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingBytesPointer(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingBytesPointer(nil, &dst)
	if err == nil { t.Fatal("expected error") }
	err2 := corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &dst)
	if err2 != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingBytesIf(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &dst)
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &dst)
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingBytesPointerIf(t *testing.T) {
	var dst string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingBytesMust(t *testing.T) {
	var dst string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &dst)
	if dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	var dst string
	corejson.Deserialize.UsingSafeBytesMust(nil, &dst) // should not panic
}

func TestDeserialize_UsingBytesPointerMust(t *testing.T) {
	var dst string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &dst)
	if dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_AnyToFieldsMap(t *testing.T) {
	type s struct{ Name string }
	fm, err := corejson.Deserialize.AnyToFieldsMap(s{Name: "x"})
	if err != nil { t.Fatal("unexpected") }
	_ = fm
}

func TestDeserialize_UsingDeserializerToOption(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingDeserializerToOption(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingDeserializerDefined(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingDeserializerFuncDefined(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &struct{}{})
	if err == nil { t.Fatal("expected error") }
	err2 := corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error { return nil }, &struct{}{})
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingJsonerToAny(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingJsonerToAny(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingJsonerToAnyMust(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

// ===================== deserializeFromBytesTo =====================

func TestDeserializeBytesTo_Strings(t *testing.T) {
	lines, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	if err != nil || len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_StringsMust(t *testing.T) {
	lines := corejson.Deserialize.BytesTo.StringsMust([]byte(`["a","b"]`))
	if len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_String(t *testing.T) {
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_StringMust(t *testing.T) {
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer(t *testing.T) {
	i, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))
	if err != nil || i != 42 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_IntegerMust(t *testing.T) {
	i := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
	if i != 42 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer64(t *testing.T) {
	i, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))
	if err != nil || i != 99 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer64Must(t *testing.T) {
	i := corejson.Deserialize.BytesTo.Integer64Must([]byte(`99`))
	if i != 99 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integers(t *testing.T) {
	ints, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	if err != nil || len(ints) != 3 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_IntegersMust(t *testing.T) {
	ints := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))
	if len(ints) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapAnyItem(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapAnyItemMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
	if m == nil { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapStringStringMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
	if m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Bool(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))
	if err != nil || !b { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_BoolMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BoolMust([]byte(`true`))
	if !b { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Bytes(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
}

func TestDeserializeBytesTo_BytesMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
}

// ===================== deserializeFromResultTo =====================

func TestDeserializeResultTo_String(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_StringMust(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	s := corejson.Deserialize.ResultTo.StringMust(r)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_StringsMust(t *testing.T) {
	r := corejson.Serialize.Apply([]string{"a", "b"})
	lines := corejson.Deserialize.ResultTo.StringsMust(r)
	if len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_Bool(t *testing.T) {
	r := corejson.Serialize.Apply(true)
	b, err := corejson.Deserialize.ResultTo.Bool(r)
	if err != nil || !b { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_BoolMust(t *testing.T) {
	r := corejson.Serialize.Apply(true)
	if !corejson.Deserialize.ResultTo.BoolMust(r) { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_Byte(t *testing.T) {
	r := corejson.Serialize.Apply(byte(65))
	b, err := corejson.Deserialize.ResultTo.Byte(r)
	if err != nil || b != 65 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_ByteMust(t *testing.T) {
	r := corejson.Serialize.Apply(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)
	if b != 65 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapAnyItem(t *testing.T) {
	r := corejson.Serialize.Apply(map[string]any{"a": 1})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.Serialize.Apply(map[string]any{"a": 1})
	_ = corejson.Deserialize.ResultTo.MapAnyItemMust(r)
}

func TestDeserializeResultTo_MapStringString(t *testing.T) {
	r := corejson.Serialize.Apply(map[string]string{"a": "b"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.Serialize.Apply(map[string]string{"a": "b"})
	_ = corejson.Deserialize.ResultTo.MapStringStringMust(r)
}

// ===================== global funcs =====================

func TestJsonString_Global(t *testing.T) {
	s, err := corejson.JsonString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestJsonStringOrErrMsg_OK(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	if s == "" { t.Fatal("unexpected") }
}

func TestJsonStringOrErrMsg_Error(t *testing.T) {
	s := corejson.JsonStringOrErrMsg(make(chan int))
	if s == "" { t.Fatal("expected error msg") }
}

// ===================== Bytes helpers =====================

func TestBytesToString(t *testing.T) {
	if corejson.BytesToString(nil) != "" { t.Fatal("expected empty") }
	if corejson.BytesToString([]byte(`"x"`)) == "" { t.Fatal("expected non-empty") }
}

func TestBytesToPrettyString(t *testing.T) {
	if corejson.BytesToPrettyString(nil) != "" { t.Fatal("expected empty") }
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" { t.Fatal("expected non-empty") }
}

func TestBytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte(`"x"`))
	if len(b) != 0 { t.Fatal("expected empty") }
}

func TestBytesCloneIf_DeepClone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte(`"x"`))
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

func TestBytesCloneIf_Empty(t *testing.T) {
	b := corejson.BytesCloneIf(true, nil)
	if len(b) != 0 { t.Fatal("expected empty") }
}

func TestBytesDeepClone(t *testing.T) {
	if len(BytesDeepClone(nil)) != 0 { t.Fatal("expected empty") }
	b := BytesDeepClone([]byte(`"x"`))
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

// ===================== newResultCreator extras =====================

func TestNewResult_Various(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesType([]byte(`"x"`), "t")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`"x"`), "t")
	_ = corejson.NewResult.UsingTypeBytesPtr("t", []byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.corejson.New("e"), "t")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr(nil, errors.corejson.New("e"), "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	s := `"hello"`
	_ = corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.corejson.New("e"), nil, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = corejson.NewResult.UsingTypePlusString("t", `"x"`)
	sp := `"x"`
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", &sp)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", nil)
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "t")
	_ = corejson.NewResult.UsingString(`"x"`)
	_ = corejson.NewResult.UsingStringPtr(&sp)
	_ = corejson.NewResult.UsingStringPtr(nil)
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.corejson.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.CastingAny("hello")
	_ = corejson.NewResult.Error(errors.corejson.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.corejson.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("t")
	_ = corejson.NewResult.TypeNameBytes("t")
	_ = corejson.NewResult.Many("a", "b")
	_ = corejson.NewResult.Serialize("hello")
	_ = corejson.NewResult.Marshal("hello")
	_ = corejson.NewResult.UsingSerializer(nil)
	_ = corejson.NewResult.UsingSerializerFunc(nil)
	_ = corejson.NewResult.UsingJsoner(nil)
	_ = corejson.NewResult.AnyToCastingResult("hello")
	_ = corejson.NewResult.UnmarshalUsingBytes([]byte(`{}`))
	_ = corejson.NewResult.DeserializeUsingBytes([]byte(`{}`))
}

// ===================== emptyCreator =====================

func TestEmpty_All(t *testing.T) {
	_ = corejson.Empty.corejson.Result()
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.ResultWithErr("t", errors.corejson.New("e"))
	_ = corejson.Empty.ResultPtrWithErr("t", errors.corejson.New("e"))
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}

// ===================== Serializer extras =====================

func TestSerializer_FromStringer(t *testing.T) {
	type myStringer struct{}
	// Can't test without a Stringer impl, but we can test with errors.New which implements String via Error
}

func TestStaticJsonError(t *testing.T) {
	if StaticJsonError == nil { t.Fatal("expected non-nil") }
}
