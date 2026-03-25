package corejson

import (
	"errors"
	"testing"
)

// Covers: anyTo and castingAny methods

func Test_I17_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := AnyTo.SerializedJsonResult(nil)
	if r.Error == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_Result(t *testing.T) {
	result := New("hello")
	r := AnyTo.SerializedJsonResult(result)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	result := NewPtr("hello")
	r := AnyTo.SerializedJsonResult(result)
	if r != result {
		t.Fatal("expected same pointer")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	r := AnyTo.SerializedJsonResult([]byte(`"hi"`))
	if r == nil || r.TypeName != "RawBytes" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_String(t *testing.T) {
	r := AnyTo.SerializedJsonResult("hello")
	if r == nil || r.TypeName != "RawString" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	r := AnyTo.SerializedJsonResult(errors.New("test error"))
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	r := AnyTo.SerializedJsonResult(errors.New(""))
	if r == nil {
		t.Fatal("expected non-nil for empty error")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_Serializer(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	r := AnyTo.SerializedJsonResult(s)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_AnyTo_SerializedJsonResult_AnyItem(t *testing.T) {
	r := AnyTo.SerializedJsonResult(42)
	if r == nil || r.HasError() {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SerializedRaw(t *testing.T) {
	b, err := AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SerializedString(t *testing.T) {
	s, err := AnyTo.SerializedString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SerializedString_Error(t *testing.T) {
	_, err := AnyTo.SerializedString(nil)
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_I17_AnyTo_SerializedSafeString(t *testing.T) {
	s := AnyTo.SerializedSafeString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_SerializedSafeString_Error(t *testing.T) {
	s := AnyTo.SerializedSafeString(nil)
	if s != "" {
		t.Fatal("expected empty for error")
	}
}

func Test_I17_AnyTo_SerializedStringMust(t *testing.T) {
	s := AnyTo.SerializedStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_SerializedStringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	AnyTo.SerializedStringMust(nil)
}

func Test_I17_AnyTo_SafeJsonString(t *testing.T) {
	s := AnyTo.SafeJsonString("hello")
	if s != `"hello"` {
		t.Fatalf("unexpected: %s", s)
	}
}

func Test_I17_AnyTo_PrettyStringWithError_String(t *testing.T) {
	s, err := AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := New("hello")
	s, err := AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := NewPtr("hello")
	s, err := AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_ResultWithErr(t *testing.T) {
	r := New("hello")
	r.Error = errors.New("test")
	s, err := AnyTo.PrettyStringWithError(r)
	if err == nil || s == "" {
		t.Fatal("expected error path")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_ResultPtrWithErr(t *testing.T) {
	r := NewPtr("hello")
	r.Error = errors.New("test")
	s, err := AnyTo.PrettyStringWithError(r)
	if err == nil || s == "" {
		t.Fatal("expected error path")
	}
}

func Test_I17_AnyTo_PrettyStringWithError_AnyItem(t *testing.T) {
	s, err := AnyTo.PrettyStringWithError(42)
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	s := AnyTo.SafeJsonPrettyString("hello")
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	s := AnyTo.SafeJsonPrettyString(New("hello"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	s := AnyTo.SafeJsonPrettyString(NewPtr("hello"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_SafeJsonPrettyString_AnyItem(t *testing.T) {
	s := AnyTo.SafeJsonPrettyString(42)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_JsonString_String(t *testing.T) {
	s := AnyTo.JsonString("hello")
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonString_Bytes(t *testing.T) {
	s := AnyTo.JsonString([]byte(`"hi"`))
	if s != `"hi"` {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonString_Result(t *testing.T) {
	s := AnyTo.JsonString(New("hello"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_JsonString_ResultPtr(t *testing.T) {
	s := AnyTo.JsonString(NewPtr("hello"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_JsonString_AnyItem(t *testing.T) {
	s := AnyTo.JsonString(42)
	if s != "42" {
		t.Fatalf("unexpected: %s", s)
	}
}

func Test_I17_AnyTo_JsonStringWithErr_String(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr([]byte(`"hi"`))
	if err != nil || s != `"hi"` {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_Result(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr(New("hello"))
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_ResultWithErr(t *testing.T) {
	r := New("hello")
	r.Error = errors.New("test")
	s, err := AnyTo.JsonStringWithErr(r)
	if err == nil || s == "" {
		t.Fatal("expected error path")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_ResultPtr(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr(NewPtr("hello"))
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_ResultPtrWithErr(t *testing.T) {
	r := NewPtr("hello")
	r.Error = errors.New("test")
	s, err := AnyTo.JsonStringWithErr(r)
	if err == nil || s == "" {
		t.Fatal("expected error path")
	}
}

func Test_I17_AnyTo_JsonStringWithErr_AnyItem(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr(42)
	if err != nil || s != "42" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringMust(t *testing.T) {
	s := AnyTo.JsonStringMust("hello")
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_AnyTo_JsonStringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	AnyTo.JsonStringMust(func() {})
}

func Test_I17_AnyTo_PrettyStringMust(t *testing.T) {
	s := AnyTo.PrettyStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_AnyTo_PrettyStringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	AnyTo.PrettyStringMust(func() {})
}

func Test_I17_AnyTo_UsingSerializer(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	r := AnyTo.UsingSerializer(s)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_AnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	if err != nil || m == nil {
		t.Fatal("unexpected")
	}
}

// castingAny tests
func Test_I17_CastAny_FromToDefault_Bytes(t *testing.T) {
	var n int
	err := CastAny.FromToDefault([]byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToDefault_String(t *testing.T) {
	var n int
	err := CastAny.FromToDefault("42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToReflection(t *testing.T) {
	var n int
	err := CastAny.FromToReflection([]byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_SkipReflection(t *testing.T) {
	var n int
	err := CastAny.FromToOption(false, []byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_NilFrom(t *testing.T) {
	var n int
	err := CastAny.FromToOption(true, nil, &n)
	if err == nil {
		t.Fatal("expected error for nil from")
	}
}

func Test_I17_CastAny_FromToOption_Result(t *testing.T) {
	r := New("hello")
	var s string
	err := CastAny.FromToOption(false, r, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_ResultPtr(t *testing.T) {
	r := NewPtr("hello")
	var s string
	err := CastAny.FromToOption(false, r, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_Serializer(t *testing.T) {
	s := testSerializer{data: []byte(`"hello"`)}
	var out string
	err := CastAny.FromToOption(false, s, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_SerializerError(t *testing.T) {
	s := testSerializer{err: errors.New("fail")}
	var out string
	err := CastAny.FromToOption(false, s, &out)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var out string
	err := CastAny.FromToOption(false, fn, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_Error(t *testing.T) {
	errInput := errors.New(`"hello"`)
	var out string
	err := CastAny.FromToOption(false, errInput, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_FromToOption_ErrorParseFailure(t *testing.T) {
	errInput := errors.New("not-json")
	var out int
	err := CastAny.FromToOption(false, errInput, &out)
	if err == nil {
		t.Fatal("expected error for bad json error")
	}
}

func Test_I17_CastAny_FromToOption_AnyFallback(t *testing.T) {
	type sample struct {
		Name string
	}
	from := sample{Name: "test"}
	var to sample
	err := CastAny.FromToOption(false, from, &to)
	if err != nil || to.Name != "test" {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_OrDeserializeTo(t *testing.T) {
	var n int
	err := CastAny.OrDeserializeTo([]byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_CastAny_ReflectionCasting_SameType(t *testing.T) {
	type sample struct{ N int }
	from := &sample{N: 42}
	to := &sample{}
	err := CastAny.FromToOption(true, from, to)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_CastAny_ReflectionCasting_DiffType(t *testing.T) {
	var n int
	err := CastAny.FromToOption(true, "42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}
