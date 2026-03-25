package corejson

import (
	"errors"
	"testing"
)

// ── deserializerLogic ──

func TestDeserializer_Apply(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	var s string
	err := Deserialize.Apply(r, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingStringPtr(nil, &s)
	if err == nil { t.Fatal("expected error for nil bytes") }
}

func TestDeserializer_UsingStringPtr_Valid(t *testing.T) {
	str := `"hello"`
	var s string
	err := Deserialize.UsingStringPtr(&str, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingError_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingError(nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingError_Valid(t *testing.T) {
	e := errors.New(`"hello"`)
	var s string
	err := Deserialize.UsingError(e, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingErrorWhichJsonResult(nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingResult(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	var s string
	err := Deserialize.UsingResult(r, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingString(t *testing.T) {
	var s string
	err := Deserialize.UsingString(`"hello"`, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_FromString(t *testing.T) {
	var s string
	err := Deserialize.FromString(`"hello"`, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_FromTo(t *testing.T) {
	var out string
	err := Deserialize.FromTo([]byte(`"hi"`), &out)
	if err != nil || out != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_MapAnyToPointer_SkipEmpty(t *testing.T) {
	var out map[string]any
	err := Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_MapAnyToPointer_Valid(t *testing.T) {
	type simple struct {
		Name string `json:"name"`
	}
	var out simple
	err := Deserialize.MapAnyToPointer(false, map[string]any{"name": "test"}, &out)
	if err != nil || out.Name != "test" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingStringOption_IgnoreEmpty(t *testing.T) {
	var s string
	err := Deserialize.UsingStringOption(true, "", &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingStringOption_NotIgnore(t *testing.T) {
	var s string
	err := Deserialize.UsingStringOption(false, `"hi"`, &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingStringIgnoreEmpty_Empty(t *testing.T) {
	var s string
	err := Deserialize.UsingStringIgnoreEmpty("", &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingStringIgnoreEmpty_Valid(t *testing.T) {
	var s string
	err := Deserialize.UsingStringIgnoreEmpty(`"hi"`, &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingBytesIf_Skip(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesIf(false, nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingBytesIf_Do(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesIf(true, []byte(`"hi"`), &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingBytesPointer_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesPointer(nil, &s)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserializer_UsingBytesPointer_Valid(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesPointer([]byte(`"hi"`), &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingBytesPointerIf_Skip(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesPointerIf(false, nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingBytesPointerIf_Do(t *testing.T) {
	var s string
	err := Deserialize.UsingBytesPointerIf(true, []byte(`"hi"`), &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestDeserializer_AnyToFieldsMap(t *testing.T) {
	m, err := Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func TestDeserializer_UsingSerializerFuncTo(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var s string
	err := Deserialize.UsingSerializerFuncTo(fn, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestDeserializer_UsingDeserializerToOption_NilSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingDeserializerToOption(true, nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingDeserializerToOption_NilNoSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingDeserializerToOption(false, nil, &s)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserializer_UsingDeserializerDefined_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingDeserializerDefined(nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	var s string
	err := Deserialize.UsingDeserializerFuncDefined(nil, &s)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserializer_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	fn := func(toPtr any) error { return nil }
	var s string
	err := Deserialize.UsingDeserializerFuncDefined(fn, &s)
	if err != nil { t.Fatal("unexpected") }
}

func TestDeserializer_UsingJsonerToAny_NilSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingJsonerToAny(true, nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingJsonerToAny_NilNoSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingJsonerToAny(false, nil, &s)
	if err == nil { t.Fatal("expected error") }
}

func TestDeserializer_UsingJsonerToAnyMust_NilSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	if err != nil { t.Fatal("expected nil") }
}

func TestDeserializer_UsingJsonerToAnyMust_NilNoSkip(t *testing.T) {
	var s string
	err := Deserialize.UsingJsonerToAnyMust(false, nil, &s)
	if err == nil { t.Fatal("expected error") }
}

// ── deserializeFromBytesTo ──

func TestBytesTo_Strings(t *testing.T) {
	lines, err := Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	if err != nil || len(lines) != 2 { t.Fatal("unexpected") }
}

func TestBytesTo_String(t *testing.T) {
	s, err := Deserialize.BytesTo.String([]byte(`"hello"`))
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestBytesTo_Integer(t *testing.T) {
	v, err := Deserialize.BytesTo.Integer([]byte(`42`))
	if err != nil || v != 42 { t.Fatal("unexpected") }
}

func TestBytesTo_Integer64(t *testing.T) {
	v, err := Deserialize.BytesTo.Integer64([]byte(`64`))
	if err != nil || v != 64 { t.Fatal("unexpected") }
}

func TestBytesTo_Integers(t *testing.T) {
	v, err := Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	if err != nil || len(v) != 3 { t.Fatal("unexpected") }
}

func TestBytesTo_Bool(t *testing.T) {
	v, err := Deserialize.BytesTo.Bool([]byte(`true`))
	if err != nil || !v { t.Fatal("unexpected") }
}

func TestBytesTo_MapAnyItem(t *testing.T) {
	m, err := Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestBytesTo_MapStringString(t *testing.T) {
	m, err := Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

func TestBytesTo_Bytes(t *testing.T) {
	_, err := Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
	if err != nil { t.Fatal("unexpected") }
}

// ── deserializeFromResultTo ──

func TestResultTo_String(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	s, err := Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestResultTo_Bool(t *testing.T) {
	r := NewResult.AnyPtr(true)
	v, err := Deserialize.ResultTo.Bool(r)
	if err != nil || !v { t.Fatal("unexpected") }
}

func TestResultTo_Byte(t *testing.T) {
	r := NewResult.AnyPtr(byte(5))
	v, err := Deserialize.ResultTo.Byte(r)
	if err != nil || v != 5 { t.Fatal("unexpected") }
}

func TestResultTo_MapAnyItem(t *testing.T) {
	r := NewResult.AnyPtr(map[string]any{"a": 1})
	m, err := Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || m == nil { t.Fatal("unexpected") }
}

func TestResultTo_MapStringString(t *testing.T) {
	r := NewResult.AnyPtr(map[string]string{"a": "b"})
	m, err := Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["a"] != "b" { t.Fatal("unexpected") }
}

// ── Unmarshal on nil Result ──

func TestResult_Unmarshal_Nil(t *testing.T) {
	var r *Result
	var s string
	err := r.Unmarshal(&s)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Unmarshal_WithExistingError(t *testing.T) {
	r := &Result{Error: errors.New("existing")}
	var s string
	err := r.Unmarshal(&s)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Unmarshal_BadPayload(t *testing.T) {
	r := &Result{Bytes: []byte("bad")}
	var s string
	err := r.Unmarshal(&s)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Unmarshal_Valid(t *testing.T) {
	r := NewResult.AnyPtr("test")
	var s string
	err := r.Unmarshal(&s)
	if err != nil || s != "test" { t.Fatal("unexpected") }
}

// ── AnyTo.UsingSerializer ──

func TestAnyTo_UsingSerializer(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	r := AnyTo.UsingSerializer(fn)
	if r == nil || r.HasError() { t.Fatal("unexpected") }
}

// ── CastAny.FromToReflection ──

func TestCastAny_FromToReflection(t *testing.T) {
	var out string
	err := CastAny.FromToReflection([]byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

// ── AnyTo.SerializedJsonResult — error case ──

func TestAnyTo_SerializedJsonResult_Error(t *testing.T) {
	e := errors.New("test error message")
	r := AnyTo.SerializedJsonResult(e)
	if r == nil { t.Fatal("expected non-nil") }
}

// ── AnyTo.SerializedJsonResult — empty error ──

func TestAnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	e := errors.New("")
	r := AnyTo.SerializedJsonResult(e)
	if r == nil { t.Fatal("expected non-nil") }
}
