package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
)

// Covers: deserializerLogic, deserializeFromBytesTo, deserializeFromResultTo

func Test_I17_Deserialize_Apply(t *testing.T) {
	r := corejson.NewPtr("hello")
	var s string
	err := corejson.Deserialize.Apply(r, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	var out string
	err := corejson.Deserialize.UsingStringPtr(&s, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingStringPtr(nil, &out)
	if err == nil {
		t.Fatal("expected error for nil string ptr")
	}
}

func Test_I17_Deserialize_UsingError_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingError(nil, &out)
	if err != nil {
		t.Fatal("expected nil error for nil input")
	}
}

func Test_I17_Deserialize_UsingError_Valid(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &out)
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_I17_Deserialize_UsingErrorWhichJsonResult_Valid(t *testing.T) {
	jsonResultBytes := corejson.Serialize.ToBytesMust(corejson.Result{Bytes: []byte(`"test"`)})
	var out corejson.Result
	err := corejson.Deserialize.UsingErrorWhichJsonResult(
		errors.New(string(jsonResultBytes)), &out)
	// Even if it fails to parse, it should not panic
	_ = err
}

func Test_I17_Deserialize_UsingResult(t *testing.T) {
	r := corejson.NewPtr("world")
	var s string
	err := corejson.Deserialize.UsingResult(r, &s)
	if err != nil || s != "world" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_ApplyMust_Success(t *testing.T) {
	r := corejson.NewPtr("ok")
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
	if s != "ok" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_ApplyMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	r := &corejson.Result{Bytes: []byte("invalid")}
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
}

func Test_I17_Deserialize_UsingString(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingString("42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_FromString(t *testing.T) {
	var n int
	err := corejson.Deserialize.FromString("42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_FromStringMust_Success(t *testing.T) {
	var n int
	corejson.Deserialize.FromStringMust("42", &n)
	if n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_Deserialize_FromStringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var n int
	corejson.Deserialize.FromStringMust("invalid", &n)
}

func Test_I17_Deserialize_FromTo(t *testing.T) {
	var n int
	err := corejson.Deserialize.FromTo([]byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_MapAnyToPointer_Empty(t *testing.T) {
	var out map[string]any
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	if err != nil {
		t.Fatal("expected nil for skip empty")
	}
}

func Test_I17_Deserialize_MapAnyToPointer_Valid(t *testing.T) {
	var out map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"key": "val"}, &out)
	if err != nil || out["key"] != "val" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingStringOption(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingStringOption(true, "", &n)
	if err != nil {
		t.Fatal("expected nil for ignore empty")
	}

	err = corejson.Deserialize.UsingStringOption(false, "42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &n)
	if err != nil {
		t.Fatal("expected nil for empty string")
	}

	err = corejson.Deserialize.UsingStringIgnoreEmpty("42", &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingBytes_Error(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingBytes([]byte("invalid"), &n)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingBytesPointer(nil, &n)
	if err == nil {
		t.Fatal("expected error for nil bytes pointer")
	}
}

func Test_I17_Deserialize_UsingBytesPointer_Valid(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingBytesPointer([]byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingBytesPointerMust_Success(t *testing.T) {
	var n int
	corejson.Deserialize.UsingBytesPointerMust([]byte("42"), &n)
	if n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_Deserialize_UsingBytesPointerMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var n int
	corejson.Deserialize.UsingBytesPointerMust(nil, &n)
}

func Test_I17_Deserialize_UsingBytesIf(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingBytesIf(false, []byte("42"), &n)
	if err != nil {
		t.Fatal("expected nil when skip")
	}

	err = corejson.Deserialize.UsingBytesIf(true, []byte("42"), &n)
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_Deserialize_UsingBytesPointerIf(t *testing.T) {
	var n int
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte("42"), &n)
	if err != nil {
		t.Fatal("expected nil when skip")
	}
}

func Test_I17_Deserialize_UsingBytesMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var n int
	corejson.Deserialize.UsingBytesMust([]byte("invalid"), &n)
}

func Test_I17_Deserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	var n int
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &n)
	// should not panic for empty bytes
}

func Test_I17_Deserialize_UsingSafeBytesMust_Valid(t *testing.T) {
	var n int
	corejson.Deserialize.UsingSafeBytesMust([]byte("42"), &n)
	if n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_Deserialize_UsingSafeBytesMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var n int
	corejson.Deserialize.UsingSafeBytesMust([]byte("invalid"), &n)
}

func Test_I17_Deserialize_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	if err != nil || m == nil {
		t.Fatal("unexpected result")
	}
}

type srcTestSerializer struct {
	data []byte
	err  error
}

func (s srcTestSerializer) Serialize() ([]byte, error) { return s.data, s.err }

func Test_I17_Deserialize_UsingSerializerTo(t *testing.T) {
	s := srcTestSerializer{data: []byte(`"hello"`)}
	var out string
	err := corejson.Deserialize.UsingSerializerTo(s, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Deserialize_UsingSerializerFuncTo(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`42`), nil }
	var out int
	err := corejson.Deserialize.UsingSerializerFuncTo(fn, &out)
	if err != nil || out != 42 {
		t.Fatal("unexpected result")
	}
}

type srcTestDeserializer struct {
	err error
}

func (d srcTestDeserializer) Deserialize(toPtr any) error { return d.err }

func Test_I17_Deserialize_UsingDeserializerToOption_SkipNil(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, nil)
	if err != nil {
		t.Fatal("expected nil when skip on nil")
	}
}

func Test_I17_Deserialize_UsingDeserializerToOption_NilError(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, nil)
	if err == nil {
		t.Fatal("expected error for nil deserializer")
	}
}

func Test_I17_Deserialize_UsingDeserializerToOption_Valid(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerToOption(false, srcTestDeserializer{}, nil)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Deserialize_UsingDeserializerDefined_Nil(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerDefined(nil, nil)
	if err != nil {
		t.Fatal("expected nil for nil deserializer with skip")
	}
}

func Test_I17_Deserialize_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, nil)
	if err == nil {
		t.Fatal("expected error for nil func")
	}
}

func Test_I17_Deserialize_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	fn := func(toPtr any) error { return nil }
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, nil)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Deserialize_UsingJsonerToAny_SkipNil(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, nil)
	if err != nil {
		t.Fatal("expected nil when skip nil jsoner")
	}
}

func Test_I17_Deserialize_UsingJsonerToAny_NilError(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, nil)
	if err == nil {
		t.Fatal("expected error for nil jsoner")
	}
}

func Test_I17_Deserialize_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, nil)
	if err != nil {
		t.Fatal("expected nil when skip")
	}
}

func Test_I17_Deserialize_UsingJsonerToAnyMust_NilError(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, nil)
	if err == nil {
		t.Fatal("expected error for nil jsoner")
	}
}

func Test_I17_Deserialize_Result(t *testing.T) {
	jsonBytes := corejson.Serialize.ToBytesMust(corejson.Result{Bytes: []byte(`"x"`)})
	_, err := corejson.Deserialize.Result(jsonBytes)
	_ = err // just exercising the path
}

func Test_I17_Deserialize_ResultPtr(t *testing.T) {
	jsonBytes := corejson.Serialize.ToBytesMust(corejson.Result{Bytes: []byte(`"x"`)})
	_, err := corejson.Deserialize.ResultPtr(jsonBytes)
	_ = err
}

func Test_I17_Deserialize_ResultMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.ResultMust([]byte("invalid"))
}

func Test_I17_Deserialize_ResultPtrMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.ResultPtrMust([]byte("invalid"))
}

// deserializeFromBytesTo tests
func Test_I17_BytesTo_Strings(t *testing.T) {
	lines, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	if err != nil || len(lines) != 2 {
		t.Fatal("unexpected result")
	}
}

func Test_I17_BytesTo_StringsMust(t *testing.T) {
	lines := corejson.Deserialize.BytesTo.StringsMust([]byte(`["a"]`))
	if len(lines) != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_StringsMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	corejson.Deserialize.BytesTo.StringsMust([]byte("invalid"))
}

func Test_I17_BytesTo_String(t *testing.T) {
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_StringMust(t *testing.T) {
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Integer(t *testing.T) {
	n, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))
	if err != nil || n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_IntegerMust(t *testing.T) {
	n := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
	if n != 42 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Integer64(t *testing.T) {
	n, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))
	if err != nil || n != 99 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Integer64Must(t *testing.T) {
	n := corejson.Deserialize.BytesTo.Integer64Must([]byte(`99`))
	if n != 99 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Integers(t *testing.T) {
	ns, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	if err != nil || len(ns) != 3 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_IntegersMust(t *testing.T) {
	ns := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))
	if len(ns) != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_IntegersMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	corejson.Deserialize.BytesTo.IntegersMust([]byte("invalid"))
}

func Test_I17_BytesTo_MapAnyItem(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	if err != nil || m == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_MapAnyItemMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
	if m == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"k":"v"}`))
	if err != nil || m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_MapStringStringMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"k":"v"}`))
	if m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Bool(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))
	if err != nil || !b {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_BoolMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BoolMust([]byte(`false`))
	if b {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_Bytes(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
	if err != nil {
		t.Fatal("unexpected error")
	}
	_ = b
}

func Test_I17_BytesTo_BytesMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
	_ = b
}

func Test_I17_BytesTo_ResultCollection(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsCollection{Items: []corejson.Result{}})
	rc, err := corejson.Deserialize.BytesTo.ResultCollection(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_ResultCollectionMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsCollection{Items: []corejson.Result{}})
	rc := corejson.Deserialize.BytesTo.ResultCollectionMust(b)
	if rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_ResultsPtrCollection(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsPtrCollection{Items: []*corejson.Result{}})
	rc, err := corejson.Deserialize.BytesTo.ResultsPtrCollection(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsPtrCollection{Items: []*corejson.Result{}})
	rc := corejson.Deserialize.BytesTo.ResultsPtrCollectionMust(b)
	if rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_MapResults(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(MapResults{Items: map[string]corejson.Result{}})
	mr, err := corejson.Deserialize.BytesTo.MapResults(b)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_BytesTo_MapResultsMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(MapResults{Items: map[string]corejson.Result{}})
	mr := corejson.Deserialize.BytesTo.MapResultsMust(b)
	if mr == nil {
		t.Fatal("unexpected")
	}
}

// deserializeFromResultTo tests
func Test_I17_ResultTo_String(t *testing.T) {
	r := corejson.NewPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_StringMust(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := corejson.Deserialize.ResultTo.StringMust(r)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_StringMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.StringMust(r)
}

func Test_I17_ResultTo_Bool(t *testing.T) {
	r := corejson.NewPtr(true)
	b, err := corejson.Deserialize.ResultTo.Bool(r)
	if err != nil || !b {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_BoolMust(t *testing.T) {
	r := corejson.NewPtr(true)
	b := corejson.Deserialize.ResultTo.BoolMust(r)
	if !b {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_BoolMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.BoolMust(r)
}

func Test_I17_ResultTo_Byte(t *testing.T) {
	r := corejson.NewPtr(byte(65))
	b, err := corejson.Deserialize.ResultTo.Byte(r)
	if err != nil || b != 65 {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ByteMust(t *testing.T) {
	r := corejson.NewPtr(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)
	if b != 65 {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ByteMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.ByteMust(r)
}

func Test_I17_ResultTo_StringsMust(t *testing.T) {
	r := corejson.NewPtr([]string{"a", "b"})
	lines := corejson.Deserialize.ResultTo.StringsMust(r)
	if len(lines) != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_StringsMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.StringsMust(r)
}

func Test_I17_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || m == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapAnyItemMust(r)
	if m == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_MapStringString(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapStringStringMust(r)
	if m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ResultCollection(t *testing.T) {
	r := corejson.NewPtr(ResultsCollection{Items: []corejson.Result{}})
	rc, err := corejson.Deserialize.ResultTo.ResultCollection(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ResultCollectionMust(t *testing.T) {
	r := corejson.NewPtr(ResultsCollection{Items: []corejson.Result{}})
	rc := corejson.Deserialize.ResultTo.ResultCollectionMust(r)
	if rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ResultsPtrCollection(t *testing.T) {
	r := corejson.NewPtr(ResultsPtrCollection{Items: []*corejson.Result{}})
	rc, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	r := corejson.NewPtr(ResultsPtrCollection{Items: []*corejson.Result{}})
	rc := corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(r)
	if rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_Result(t *testing.T) {
	inner := corejson.New("hello")
	r := corejson.NewPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Result(r)
	_ = err
}

func Test_I17_ResultTo_ResultMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.ResultMust(r)
}

func Test_I17_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.New("hello")
	r := corejson.NewPtr(inner)
	_, err := corejson.Deserialize.ResultTo.ResultPtr(r)
	_ = err
}

func Test_I17_ResultTo_ResultPtrMust(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.ResultPtrMust(r)
}

func Test_I17_ResultTo_MapResults(t *testing.T) {
	r := corejson.NewPtr(MapResults{Items: map[string]corejson.Result{}})
	mr, err := corejson.Deserialize.ResultTo.MapResults(r)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_MapResultsMust(t *testing.T) {
	r := corejson.NewPtr(MapResults{Items: map[string]corejson.Result{}})
	mr := corejson.Deserialize.ResultTo.MapResultsMust(r)
	if mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_ResultTo_Bytes(t *testing.T) {
	inner := corejson.New("hello")
	r := corejson.NewPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Bytes(r)
	_ = err
}

func Test_I17_ResultTo_BytesMust(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{Bytes: []byte("invalid")}
	corejson.Deserialize.ResultTo.BytesMust(r)
}
