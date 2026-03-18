package corejson

import (
	"errors"
	"testing"
)

// ===================== anyTo =====================

func TestAnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := AnyTo.SerializedJsonResult(nil)
	if r == nil || r.Error == nil { t.Fatal("expected error") }
}

func TestAnyTo_SerializedJsonResult_Result(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	out := AnyTo.SerializedJsonResult(r)
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	out := AnyTo.SerializedJsonResult(r)
	if out != r { t.Fatal("expected same ptr") }
}

func TestAnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	out := AnyTo.SerializedJsonResult([]byte(`"hello"`))
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_String(t *testing.T) {
	out := AnyTo.SerializedJsonResult(`"hello"`)
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_Error(t *testing.T) {
	out := AnyTo.SerializedJsonResult(errors.New("oops"))
	if out == nil { t.Fatal("expected non-nil") }
}

func TestAnyTo_SerializedJsonResult_AnyItem(t *testing.T) {
	out := AnyTo.SerializedJsonResult(42)
	if out == nil || out.HasError() { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedRaw(t *testing.T) {
	b, err := AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedString(t *testing.T) {
	s, err := AnyTo.SerializedString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
	_, err2 := AnyTo.SerializedString(nil)
	if err2 == nil { t.Fatal("expected error") }
}

func TestAnyTo_SerializedSafeString(t *testing.T) {
	s := AnyTo.SerializedSafeString("hello")
	if s == "" { t.Fatal("expected non-empty") }
	s2 := AnyTo.SerializedSafeString(nil)
	if s2 != "" { t.Fatal("expected empty") }
}

func TestAnyTo_SerializedStringMust(t *testing.T) {
	s := AnyTo.SerializedStringMust("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyTo_SafeJsonString(t *testing.T) {
	s := AnyTo.SafeJsonString("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyTo_PrettyStringWithError(t *testing.T) {
	s, err := AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	s2, err2 := AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" { t.Fatal("unexpected") }
	r := Result{Bytes: []byte(`{"a":1}`)}
	s3, err3 := AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" { t.Fatal("unexpected") }
	rp := &Result{Bytes: []byte(`{"a":1}`)}
	s4, err4 := AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" { t.Fatal("unexpected") }
	re := Result{Error: errors.New("e")}
	_, err5 := AnyTo.PrettyStringWithError(re)
	if err5 == nil { t.Fatal("expected error") }
	rep := &Result{Error: errors.New("e")}
	_, err6 := AnyTo.PrettyStringWithError(rep)
	if err6 == nil { t.Fatal("expected error") }
	s7, err7 := AnyTo.PrettyStringWithError(42)
	if err7 != nil || s7 == "" { t.Fatal("unexpected") }
}

func TestAnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = AnyTo.SafeJsonPrettyString("hello")
	_ = AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = AnyTo.SafeJsonPrettyString(Result{Bytes: []byte(`{"a":1}`)})
	_ = AnyTo.SafeJsonPrettyString(&Result{Bytes: []byte(`{"a":1}`)})
	_ = AnyTo.SafeJsonPrettyString(42)
}

func TestAnyTo_JsonString(t *testing.T) {
	_ = AnyTo.JsonString("hello")
	_ = AnyTo.JsonString([]byte(`"x"`))
	_ = AnyTo.JsonString(Result{Bytes: []byte(`"x"`)})
	_ = AnyTo.JsonString(&Result{Bytes: []byte(`"x"`)})
	_ = AnyTo.JsonString(42)
}

func TestAnyTo_JsonStringWithErr(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	_, _ = AnyTo.JsonStringWithErr([]byte(`"x"`))
	_, _ = AnyTo.JsonStringWithErr(Result{Bytes: []byte(`"x"`)})
	_, _ = AnyTo.JsonStringWithErr(&Result{Bytes: []byte(`"x"`)})
	_, err2 := AnyTo.JsonStringWithErr(Result{Error: errors.New("e")})
	if err2 == nil { t.Fatal("expected error") }
	_, err3 := AnyTo.JsonStringWithErr(&Result{Error: errors.New("e")})
	if err3 == nil { t.Fatal("expected error") }
	_, _ = AnyTo.JsonStringWithErr(42)
}

func TestAnyTo_JsonStringMust(t *testing.T) {
	s := AnyTo.JsonStringMust("hello")
	if s != "hello" { t.Fatal("unexpected") }
}

func TestAnyTo_PrettyStringMust(t *testing.T) {
	s := AnyTo.PrettyStringMust("hello")
	if s != "hello" { t.Fatal("unexpected") }
}

func TestAnyTo_SerializedFieldsMap(t *testing.T) {
	type s struct{ Name string }
	fm, err := AnyTo.SerializedFieldsMap(s{Name: "x"})
	if err != nil { t.Fatal("unexpected error") }
	_ = fm
}

// ===================== castingAny =====================

func TestCastAny_FromToDefault(t *testing.T) {
	src := `"hello"`
	var dst string
	err := CastAny.FromToDefault([]byte(src), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_Bytes(t *testing.T) {
	var dst string
	err := CastAny.FromToOption(false, []byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_String(t *testing.T) {
	var dst int
	err := CastAny.FromToOption(false, `42`, &dst)
	if err != nil || dst != 42 { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_Result(t *testing.T) {
	r := Result{Bytes: []byte(`"hello"`)}
	var dst string
	err := CastAny.FromToOption(false, r, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_ResultPtr(t *testing.T) {
	r := &Result{Bytes: []byte(`42`)}
	var dst int
	err := CastAny.FromToOption(false, r, &dst)
	if err != nil || dst != 42 { t.Fatal("unexpected") }
}

func TestCastAny_FromToOption_AnyItem(t *testing.T) {
	type s struct{ N int }
	src := s{N: 42}
	var dst s
	err := CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.N != 42 { t.Fatal("unexpected") }
}

func TestCastAny_OrDeserializeTo(t *testing.T) {
	var dst string
	err := CastAny.OrDeserializeTo([]byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

// ===================== deserializerLogic extras =====================

func TestDeserialize_UsingStringPtr_Nil(t *testing.T) {
	var dst string
	err := Deserialize.UsingStringPtr(nil, &dst)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingStringPtr_Normal(t *testing.T) {
	s := `"hello"`
	var dst string
	err := Deserialize.UsingStringPtr(&s, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingError_Nil(t *testing.T) {
	var dst string
	err := Deserialize.UsingError(nil, &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingErrorWhichJsonResult(t *testing.T) {
	err := Deserialize.UsingErrorWhichJsonResult(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_ApplyMust(t *testing.T) {
	r := Serialize.Apply("hello")
	var s string
	Deserialize.ApplyMust(r, &s)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_FromString(t *testing.T) {
	var dst string
	err := Deserialize.FromString(`"hello"`, &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_FromTo(t *testing.T) {
	var dst string
	err := Deserialize.FromTo([]byte(`"hello"`), &dst)
	if err != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_MapAnyToPointer(t *testing.T) {
	type s struct{ Name string }
	err := Deserialize.MapAnyToPointer(true, nil, &s{})
	if err != nil { t.Fatal("expected nil for empty skip") }
	err2 := Deserialize.MapAnyToPointer(false, map[string]any{"Name": "x"}, &s{})
	if err2 != nil { t.Fatal("unexpected error") }
}

func TestDeserialize_UsingStringOption(t *testing.T) {
	var dst string
	err := Deserialize.UsingStringOption(true, "", &dst)
	if err != nil { t.Fatal("expected nil") }
	err2 := Deserialize.UsingStringOption(false, `"x"`, &dst)
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var dst string
	err := Deserialize.UsingStringIgnoreEmpty("", &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingBytesPointer(t *testing.T) {
	var dst string
	err := Deserialize.UsingBytesPointer(nil, &dst)
	if err == nil { t.Fatal("expected error") }
	err2 := Deserialize.UsingBytesPointer([]byte(`"hello"`), &dst)
	if err2 != nil || dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingBytesIf(t *testing.T) {
	var dst string
	err := Deserialize.UsingBytesIf(false, []byte(`"x"`), &dst)
	if err != nil { t.Fatal("expected nil") }
	err2 := Deserialize.UsingBytesIf(true, []byte(`"x"`), &dst)
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingBytesPointerIf(t *testing.T) {
	var dst string
	err := Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &dst)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingBytesMust(t *testing.T) {
	var dst string
	Deserialize.UsingBytesMust([]byte(`"hello"`), &dst)
	if dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	var dst string
	Deserialize.UsingSafeBytesMust(nil, &dst) // should not panic
}

func TestDeserialize_UsingBytesPointerMust(t *testing.T) {
	var dst string
	Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &dst)
	if dst != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_AnyToFieldsMap(t *testing.T) {
	type s struct{ Name string }
	fm, err := Deserialize.AnyToFieldsMap(s{Name: "x"})
	if err != nil { t.Fatal("unexpected") }
	_ = fm
}

func TestDeserialize_UsingDeserializerToOption(t *testing.T) {
	err := Deserialize.UsingDeserializerToOption(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := Deserialize.UsingDeserializerToOption(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingDeserializerDefined(t *testing.T) {
	err := Deserialize.UsingDeserializerDefined(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserialize_UsingDeserializerFuncDefined(t *testing.T) {
	err := Deserialize.UsingDeserializerFuncDefined(nil, &struct{}{})
	if err == nil { t.Fatal("expected error") }
	err2 := Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error { return nil }, &struct{}{})
	if err2 != nil { t.Fatal("unexpected") }
}

func TestDeserialize_UsingJsonerToAny(t *testing.T) {
	err := Deserialize.UsingJsonerToAny(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := Deserialize.UsingJsonerToAny(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func TestDeserialize_UsingJsonerToAnyMust(t *testing.T) {
	err := Deserialize.UsingJsonerToAnyMust(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

// ===================== deserializeFromBytesTo =====================

func TestDeserializeBytesTo_Strings(t *testing.T) {
	lines, err := Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	if err != nil || len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_StringsMust(t *testing.T) {
	lines := Deserialize.BytesTo.StringsMust([]byte(`["a","b"]`))
	if len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_String(t *testing.T) {
	s, err := Deserialize.BytesTo.String([]byte(`"hello"`))
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_StringMust(t *testing.T) {
	s := Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer(t *testing.T) {
	i, err := Deserialize.BytesTo.Integer([]byte(`42`))
	if err != nil || i != 42 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_IntegerMust(t *testing.T) {
	i := Deserialize.BytesTo.IntegerMust([]byte(`42`))
	if i != 42 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer64(t *testing.T) {
	i, err := Deserialize.BytesTo.Integer64([]byte(`99`))
	if err != nil || i != 99 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integer64Must(t *testing.T) {
	i := Deserialize.BytesTo.Integer64Must([]byte(`99`))
	if i != 99 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Integers(t *testing.T) {
	ints, err := Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	if err != nil || len(ints) != 3 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_IntegersMust(t *testing.T) {
	ints := Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))
	if len(ints) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapAnyItem(t *testing.T) {
	m, err := Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapAnyItemMust(t *testing.T) {
	m := Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
	if m == nil { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapStringString(t *testing.T) {
	m, err := Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_MapStringStringMust(t *testing.T) {
	m := Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
	if m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Bool(t *testing.T) {
	b, err := Deserialize.BytesTo.Bool([]byte(`true`))
	if err != nil || !b { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_BoolMust(t *testing.T) {
	b := Deserialize.BytesTo.BoolMust([]byte(`true`))
	if !b { t.Fatal("unexpected") }
}

func TestDeserializeBytesTo_Bytes(t *testing.T) {
	_, _ = Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
}

func TestDeserializeBytesTo_BytesMust(t *testing.T) {
	_ = Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
}

// ===================== deserializeFromResultTo =====================

func TestDeserializeResultTo_String(t *testing.T) {
	r := Serialize.Apply("hello")
	s, err := Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_StringMust(t *testing.T) {
	r := Serialize.Apply("hello")
	s := Deserialize.ResultTo.StringMust(r)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_StringsMust(t *testing.T) {
	r := Serialize.Apply([]string{"a", "b"})
	lines := Deserialize.ResultTo.StringsMust(r)
	if len(lines) != 2 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_Bool(t *testing.T) {
	r := Serialize.Apply(true)
	b, err := Deserialize.ResultTo.Bool(r)
	if err != nil || !b { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_BoolMust(t *testing.T) {
	r := Serialize.Apply(true)
	if !Deserialize.ResultTo.BoolMust(r) { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_Byte(t *testing.T) {
	r := Serialize.Apply(byte(65))
	b, err := Deserialize.ResultTo.Byte(r)
	if err != nil || b != 65 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_ByteMust(t *testing.T) {
	r := Serialize.Apply(byte(65))
	b := Deserialize.ResultTo.ByteMust(r)
	if b != 65 { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapAnyItem(t *testing.T) {
	r := Serialize.Apply(map[string]any{"a": 1})
	m, err := Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapAnyItemMust(t *testing.T) {
	r := Serialize.Apply(map[string]any{"a": 1})
	_ = Deserialize.ResultTo.MapAnyItemMust(r)
}

func TestDeserializeResultTo_MapStringString(t *testing.T) {
	r := Serialize.Apply(map[string]string{"a": "b"})
	m, err := Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

func TestDeserializeResultTo_MapStringStringMust(t *testing.T) {
	r := Serialize.Apply(map[string]string{"a": "b"})
	_ = Deserialize.ResultTo.MapStringStringMust(r)
}

// ===================== global funcs =====================

func TestJsonString_Global(t *testing.T) {
	s, err := JsonString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestJsonStringOrErrMsg_OK(t *testing.T) {
	s := JsonStringOrErrMsg("hello")
	if s == "" { t.Fatal("unexpected") }
}

func TestJsonStringOrErrMsg_Error(t *testing.T) {
	s := JsonStringOrErrMsg(make(chan int))
	if s == "" { t.Fatal("expected error msg") }
}

// ===================== Bytes helpers =====================

func TestBytesToString(t *testing.T) {
	if BytesToString(nil) != "" { t.Fatal("expected empty") }
	if BytesToString([]byte(`"x"`)) == "" { t.Fatal("expected non-empty") }
}

func TestBytesToPrettyString(t *testing.T) {
	if BytesToPrettyString(nil) != "" { t.Fatal("expected empty") }
	s := BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" { t.Fatal("expected non-empty") }
}

func TestBytesCloneIf_NoClone(t *testing.T) {
	b := BytesCloneIf(false, []byte(`"x"`))
	if len(b) != 0 { t.Fatal("expected empty") }
}

func TestBytesCloneIf_DeepClone(t *testing.T) {
	b := BytesCloneIf(true, []byte(`"x"`))
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

func TestBytesCloneIf_Empty(t *testing.T) {
	b := BytesCloneIf(true, nil)
	if len(b) != 0 { t.Fatal("expected empty") }
}

func TestBytesDeepClone(t *testing.T) {
	if len(BytesDeepClone(nil)) != 0 { t.Fatal("expected empty") }
	b := BytesDeepClone([]byte(`"x"`))
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

// ===================== newResultCreator extras =====================

func TestNewResult_Various(t *testing.T) {
	_ = NewResult.UsingBytes([]byte(`"x"`))
	_ = NewResult.UsingBytesType([]byte(`"x"`), "t")
	_ = NewResult.UsingBytesTypePtr([]byte(`"x"`), "t")
	_ = NewResult.UsingTypeBytesPtr("t", []byte(`"x"`))
	_ = NewResult.UsingBytesPtr(nil)
	_ = NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	_ = NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingBytesErrPtr(nil, errors.New("e"), "t")
	_ = NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	s := `"hello"`
	_ = NewResult.PtrUsingStringPtr(&s, "t")
	_ = NewResult.PtrUsingStringPtr(nil, "t")
	_ = NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
	_ = NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = NewResult.UsingTypePlusString("t", `"x"`)
	sp := `"x"`
	_ = NewResult.UsingTypePlusStringPtr("t", &sp)
	_ = NewResult.UsingTypePlusStringPtr("t", nil)
	_ = NewResult.UsingStringWithType(`"x"`, "t")
	_ = NewResult.UsingString(`"x"`)
	_ = NewResult.UsingStringPtr(&sp)
	_ = NewResult.UsingStringPtr(nil)
	_ = NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.Create([]byte(`"x"`), nil, "t")
	_ = NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.CastingAny("hello")
	_ = NewResult.Error(errors.New("e"))
	_ = NewResult.ErrorPtr(errors.New("e"))
	_ = NewResult.Empty()
	_ = NewResult.EmptyPtr()
	_ = NewResult.TypeName("t")
	_ = NewResult.TypeNameBytes("t")
	_ = NewResult.Many("a", "b")
	_ = NewResult.Serialize("hello")
	_ = NewResult.Marshal("hello")
	_ = NewResult.UsingSerializer(nil)
	_ = NewResult.UsingSerializerFunc(nil)
	_ = NewResult.UsingJsoner(nil)
	_ = NewResult.AnyToCastingResult("hello")
	_ = NewResult.UnmarshalUsingBytes([]byte(`{}`))
	_ = NewResult.DeserializeUsingBytes([]byte(`{}`))
}

// ===================== emptyCreator =====================

func TestEmpty_All(t *testing.T) {
	_ = Empty.Result()
	_ = Empty.ResultPtr()
	_ = Empty.ResultWithErr("t", errors.New("e"))
	_ = Empty.ResultPtrWithErr("t", errors.New("e"))
	_ = Empty.BytesCollection()
	_ = Empty.BytesCollectionPtr()
	_ = Empty.ResultsCollection()
	_ = Empty.ResultsPtrCollection()
	_ = Empty.MapResults()
}

// ===================== Serializer extras =====================

func TestSerializer_FromStringer(t *testing.T) {
	type myStringer struct{}
	// Can't test without a Stringer impl, but we can test with errors.New which implements String via Error
}

func TestStaticJsonError(t *testing.T) {
	if StaticJsonError == nil { t.Fatal("expected non-nil") }
}
