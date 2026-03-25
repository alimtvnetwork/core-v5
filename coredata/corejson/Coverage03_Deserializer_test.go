package corejson

import (
	"errors"
	"testing"
)

func TestDeserializer_UsingBytes(t *testing.T) {
	var out string
	err := Deserialize.UsingBytes([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}

	err2 := Deserialize.UsingBytes([]byte(`invalid`), &out)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestDeserializer_UsingString(t *testing.T) {
	var out int
	err := Deserialize.UsingString("42", &out)
	if err != nil || out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_FromString(t *testing.T) {
	var out int
	err := Deserialize.FromString("42", &out)
	if err != nil || out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_FromStringMust(t *testing.T) {
	var out int
	Deserialize.FromStringMust("42", &out)
	if out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_UsingStringPtr(t *testing.T) {
	s := `"hello"`
	var out string
	err := Deserialize.UsingStringPtr(&s, &out)
	if err != nil {
		t.Fatal(err)
	}

	err2 := Deserialize.UsingStringPtr(nil, &out)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func TestDeserializer_UsingError(t *testing.T) {
	err := Deserialize.UsingError(nil, nil)
	if err != nil {
		t.Fatal("expected nil for nil error")
	}

	var out string
	err2 := Deserialize.UsingError(errors.New(`"hello"`), &out)
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestDeserializer_UsingStringOption(t *testing.T) {
	var out string
	err := Deserialize.UsingStringOption(true, "", &out)
	if err != nil {
		t.Fatal("expected nil for empty string skip")
	}

	err2 := Deserialize.UsingStringOption(false, `"x"`, &out)
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestDeserializer_UsingStringIgnoreEmpty(t *testing.T) {
	var out string
	err := Deserialize.UsingStringIgnoreEmpty("", &out)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func TestDeserializer_UsingBytesPointer(t *testing.T) {
	var out string
	err := Deserialize.UsingBytesPointer(nil, &out)
	if err == nil {
		t.Fatal("expected error for nil")
	}

	err2 := Deserialize.UsingBytesPointer([]byte(`"x"`), &out)
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestDeserializer_UsingBytesPointerMust(t *testing.T) {
	var out string
	Deserialize.UsingBytesPointerMust([]byte(`"x"`), &out)
	if out != "x" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_UsingBytesMust(t *testing.T) {
	var out int
	Deserialize.UsingBytesMust([]byte("42"), &out)
	if out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_UsingSafeBytesMust(t *testing.T) {
	var out int
	Deserialize.UsingSafeBytesMust([]byte{}, &out)

	Deserialize.UsingSafeBytesMust([]byte("42"), &out)
	if out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_UsingBytesIf(t *testing.T) {
	var out string
	err := Deserialize.UsingBytesIf(false, []byte(`"x"`), &out)
	if err != nil {
		t.Fatal("expected nil when skip")
	}

	err2 := Deserialize.UsingBytesIf(true, []byte(`"x"`), &out)
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestDeserializer_UsingBytesPointerIf(t *testing.T) {
	var out string
	err := Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &out)
	if err != nil {
		t.Fatal("expected nil when skip")
	}
}

func TestDeserializer_Apply(t *testing.T) {
	r := NewResult.Any("hello")
	var out string
	err := Deserialize.Apply(r.Ptr(), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ApplyMust(t *testing.T) {
	r := NewResult.Any(42)
	var out int
	Deserialize.ApplyMust(r.Ptr(), &out)
	if out != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_UsingResult(t *testing.T) {
	r := NewResult.Any("hi")
	var out string
	err := Deserialize.UsingResult(r.Ptr(), &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeserializer_MapAnyToPointer(t *testing.T) {
	m := map[string]any{"key": "val"}
	var out map[string]any
	err := Deserialize.MapAnyToPointer(false, m, &out)
	if err != nil {
		t.Fatal(err)
	}

	err2 := Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	if err2 != nil {
		t.Fatal("expected nil for empty skip")
	}
}

func TestDeserializer_AnyToFieldsMap(t *testing.T) {
	m, err := Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func TestDeserializer_BytesTo_Strings(t *testing.T) {
	b, _ := Serialize.Raw([]string{"a", "b"})
	lines, err := Deserialize.BytesTo.Strings(b)
	if err != nil || len(lines) != 2 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_StringsMust(t *testing.T) {
	b, _ := Serialize.Raw([]string{"a"})
	lines := Deserialize.BytesTo.StringsMust(b)
	if len(lines) != 1 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_String(t *testing.T) {
	b, _ := Serialize.Raw("hello")
	s, err := Deserialize.BytesTo.String(b)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_StringMust(t *testing.T) {
	b, _ := Serialize.Raw("x")
	s := Deserialize.BytesTo.StringMust(b)
	if s != "x" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Integer(t *testing.T) {
	b, _ := Serialize.Raw(42)
	i, err := Deserialize.BytesTo.Integer(b)
	if err != nil || i != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_IntegerMust(t *testing.T) {
	b, _ := Serialize.Raw(42)
	i := Deserialize.BytesTo.IntegerMust(b)
	if i != 42 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Integer64(t *testing.T) {
	b, _ := Serialize.Raw(64)
	i, err := Deserialize.BytesTo.Integer64(b)
	if err != nil || i != 64 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Integer64Must(t *testing.T) {
	b, _ := Serialize.Raw(64)
	i := Deserialize.BytesTo.Integer64Must(b)
	if i != 64 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Integers(t *testing.T) {
	b, _ := Serialize.Raw([]int{1, 2})
	ints, err := Deserialize.BytesTo.Integers(b)
	if err != nil || len(ints) != 2 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_IntegersMust(t *testing.T) {
	b, _ := Serialize.Raw([]int{1})
	ints := Deserialize.BytesTo.IntegersMust(b)
	if len(ints) != 1 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Bool(t *testing.T) {
	b, _ := Serialize.Raw(true)
	v, err := Deserialize.BytesTo.Bool(b)
	if err != nil || !v {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_BoolMust(t *testing.T) {
	b, _ := Serialize.Raw(false)
	v := Deserialize.BytesTo.BoolMust(b)
	if v {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_MapAnyItem(t *testing.T) {
	b, _ := Serialize.Raw(map[string]any{"k": "v"})
	m, err := Deserialize.BytesTo.MapAnyItem(b)
	if err != nil || len(m) == 0 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_MapAnyItemMust(t *testing.T) {
	b, _ := Serialize.Raw(map[string]any{"k": "v"})
	m := Deserialize.BytesTo.MapAnyItemMust(b)
	if len(m) == 0 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_MapStringString(t *testing.T) {
	b, _ := Serialize.Raw(map[string]string{"k": "v"})
	m, err := Deserialize.BytesTo.MapStringString(b)
	if err != nil || m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_MapStringStringMust(t *testing.T) {
	b, _ := Serialize.Raw(map[string]string{"k": "v"})
	m := Deserialize.BytesTo.MapStringStringMust(b)
	if m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_BytesTo_Bytes(t *testing.T) {
	input := []byte(`"aGVsbG8="`)
	_, err := Deserialize.BytesTo.Bytes(input)
	_ = err
}

func TestDeserializer_ResultTo_String(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	s, err := Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_StringMust(t *testing.T) {
	r := NewResult.AnyPtr("x")
	s := Deserialize.ResultTo.StringMust(r)
	if s != "x" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_Bool(t *testing.T) {
	r := NewResult.AnyPtr(true)
	v, err := Deserialize.ResultTo.Bool(r)
	if err != nil || !v {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_BoolMust(t *testing.T) {
	r := NewResult.AnyPtr(true)
	v := Deserialize.ResultTo.BoolMust(r)
	if !v {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_Byte(t *testing.T) {
	r := NewResult.AnyPtr(byte(65))
	_, err := Deserialize.ResultTo.Byte(r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeserializer_ResultTo_MapAnyItem(t *testing.T) {
	r := NewResult.AnyPtr(map[string]any{"k": "v"})
	m, err := Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || len(m) == 0 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_MapAnyItemMust(t *testing.T) {
	r := NewResult.AnyPtr(map[string]any{"k": "v"})
	m := Deserialize.ResultTo.MapAnyItemMust(r)
	if len(m) == 0 {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_MapStringString(t *testing.T) {
	r := NewResult.AnyPtr(map[string]string{"k": "v"})
	m, err := Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_MapStringStringMust(t *testing.T) {
	r := NewResult.AnyPtr(map[string]string{"k": "v"})
	m := Deserialize.ResultTo.MapStringStringMust(r)
	if m["k"] != "v" {
		t.Fatal("unexpected")
	}
}

func TestDeserializer_ResultTo_StringsMust(t *testing.T) {
	r := NewResult.AnyPtr([]string{"a", "b"})
	lines := Deserialize.ResultTo.StringsMust(r)
	if len(lines) != 2 {
		t.Fatal("unexpected")
	}
}
