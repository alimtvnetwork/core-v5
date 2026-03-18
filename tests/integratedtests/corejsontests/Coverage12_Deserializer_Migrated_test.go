package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Migrated from Coverage03_Deserializer_test.go ──

func Test_C03_Deserializer_UsingBytes(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytes([]byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
	err2 := corejson.Deserialize.UsingBytes([]byte(`invalid`), &out)
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C03_Deserializer_UsingString(t *testing.T) {
	var out int
	err := corejson.Deserialize.UsingString("42", &out)
	if err != nil || out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_FromString(t *testing.T) {
	var out int
	err := corejson.Deserialize.FromString("42", &out)
	if err != nil || out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_FromStringMust(t *testing.T) {
	var out int
	corejson.Deserialize.FromStringMust("42", &out)
	if out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_UsingStringPtr(t *testing.T) {
	s := `"hello"`
	var out string
	err := corejson.Deserialize.UsingStringPtr(&s, &out)
	if err != nil { t.Fatal(err) }
	err2 := corejson.Deserialize.UsingStringPtr(nil, &out)
	if err2 == nil { t.Fatal("expected error for nil") }
}

func Test_C03_Deserializer_UsingError(t *testing.T) {
	err := corejson.Deserialize.UsingError(nil, nil)
	if err != nil { t.Fatal("expected nil for nil error") }
	var out string
	err2 := corejson.Deserialize.UsingError(errors.New(`"hello"`), &out)
	if err2 != nil { t.Fatal(err2) }
}

func Test_C03_Deserializer_UsingStringOption(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingStringOption(true, "", &out)
	if err != nil { t.Fatal("expected nil for empty string skip") }
	err2 := corejson.Deserialize.UsingStringOption(false, `"x"`, &out)
	if err2 != nil { t.Fatal(err2) }
}

func Test_C03_Deserializer_UsingStringIgnoreEmpty(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &out)
	if err != nil { t.Fatal("expected nil") }
}

func Test_C03_Deserializer_UsingBytesPointer(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesPointer(nil, &out)
	if err == nil { t.Fatal("expected error for nil") }
	err2 := corejson.Deserialize.UsingBytesPointer([]byte(`"x"`), &out)
	if err2 != nil { t.Fatal(err2) }
}

func Test_C03_Deserializer_UsingBytesPointerMust(t *testing.T) {
	var out string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"x"`), &out)
	if out != "x" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_UsingBytesMust(t *testing.T) {
	var out int
	corejson.Deserialize.UsingBytesMust([]byte("42"), &out)
	if out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_UsingSafeBytesMust(t *testing.T) {
	var out int
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &out)
	corejson.Deserialize.UsingSafeBytesMust([]byte("42"), &out)
	if out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_UsingBytesIf(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &out)
	if err != nil { t.Fatal("expected nil when skip") }
	err2 := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &out)
	if err2 != nil { t.Fatal(err2) }
}

func Test_C03_Deserializer_UsingBytesPointerIf(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &out)
	if err != nil { t.Fatal("expected nil when skip") }
}

func Test_C03_Deserializer_Apply(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.Deserialize.Apply(r.Ptr(), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ApplyMust(t *testing.T) {
	r := corejson.NewResult.Any(42)
	var out int
	corejson.Deserialize.ApplyMust(r.Ptr(), &out)
	if out != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_UsingResult(t *testing.T) {
	r := corejson.NewResult.Any("hi")
	var out string
	err := corejson.Deserialize.UsingResult(r.Ptr(), &out)
	if err != nil { t.Fatal(err) }
}

func Test_C03_Deserializer_MapAnyToPointer(t *testing.T) {
	m := map[string]any{"key": "val"}
	var out map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, m, &out)
	if err != nil { t.Fatal(err) }
	err2 := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	if err2 != nil { t.Fatal("expected nil for empty skip") }
}

func Test_C03_Deserializer_AnyToFieldsMap(t *testing.T) {
	_, _ = corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
}

func Test_C03_Deserializer_BytesTo_Strings(t *testing.T) {
	b, _ := corejson.Serialize.Raw([]string{"a", "b"})
	lines, err := corejson.Deserialize.BytesTo.Strings(b)
	if err != nil || len(lines) != 2 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_StringsMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw([]string{"a"})
	lines := corejson.Deserialize.BytesTo.StringsMust(b)
	if len(lines) != 1 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_String(t *testing.T) {
	b, _ := corejson.Serialize.Raw("hello")
	s, err := corejson.Deserialize.BytesTo.String(b)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_StringMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw("x")
	s := corejson.Deserialize.BytesTo.StringMust(b)
	if s != "x" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Integer(t *testing.T) {
	b, _ := corejson.Serialize.Raw(42)
	i, err := corejson.Deserialize.BytesTo.Integer(b)
	if err != nil || i != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_IntegerMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw(42)
	i := corejson.Deserialize.BytesTo.IntegerMust(b)
	if i != 42 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Integer64(t *testing.T) {
	b, _ := corejson.Serialize.Raw(64)
	i, err := corejson.Deserialize.BytesTo.Integer64(b)
	if err != nil || i != 64 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Integer64Must(t *testing.T) {
	b, _ := corejson.Serialize.Raw(64)
	i := corejson.Deserialize.BytesTo.Integer64Must(b)
	if i != 64 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Integers(t *testing.T) {
	b, _ := corejson.Serialize.Raw([]int{1, 2})
	ints, err := corejson.Deserialize.BytesTo.Integers(b)
	if err != nil || len(ints) != 2 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_IntegersMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw([]int{1})
	ints := corejson.Deserialize.BytesTo.IntegersMust(b)
	if len(ints) != 1 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Bool(t *testing.T) {
	b, _ := corejson.Serialize.Raw(true)
	v, err := corejson.Deserialize.BytesTo.Bool(b)
	if err != nil || !v { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_BoolMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw(false)
	v := corejson.Deserialize.BytesTo.BoolMust(b)
	if v { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_MapAnyItem(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapAnyItem(b)
	if err != nil || len(m) == 0 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_MapAnyItemMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]any{"k": "v"})
	m := corejson.Deserialize.BytesTo.MapAnyItemMust(b)
	if len(m) == 0 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_MapStringString(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapStringString(b)
	if err != nil || m["k"] != "v" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_MapStringStringMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]string{"k": "v"})
	m := corejson.Deserialize.BytesTo.MapStringStringMust(b)
	if m["k"] != "v" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_BytesTo_Bytes(t *testing.T) {
	input := []byte(`"aGVsbG8="`)
	_, _ = corejson.Deserialize.BytesTo.Bytes(input)
}

func Test_C03_Deserializer_ResultTo_String(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_StringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("x")
	s := corejson.Deserialize.ResultTo.StringMust(r)
	if s != "x" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_Bool(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	v, err := corejson.Deserialize.ResultTo.Bool(r)
	if err != nil || !v { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_BoolMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	v := corejson.Deserialize.ResultTo.BoolMust(r)
	if !v { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_Byte(t *testing.T) {
	r := corejson.NewResult.AnyPtr(byte(65))
	_, err := corejson.Deserialize.ResultTo.Byte(r)
	if err != nil { t.Fatal(err) }
}

func Test_C03_Deserializer_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || len(m) == 0 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapAnyItemMust(r)
	if len(m) == 0 { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_MapStringString(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)
	if err != nil || m["k"] != "v" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapStringStringMust(r)
	if m["k"] != "v" { t.Fatal("unexpected") }
}

func Test_C03_Deserializer_ResultTo_StringsMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	lines := corejson.Deserialize.ResultTo.StringsMust(r)
	if len(lines) != 2 { t.Fatal("unexpected") }
}

// ── Additional Deserializer methods from Coverage11, 15 ──

func Test_C11_Deserialize_UsingErrorWhichJsonResult(t *testing.T) {
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func Test_C11_Deserialize_FromTo(t *testing.T) {
	var out string
	err := corejson.Deserialize.FromTo([]byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

func Test_C11_Deserialize_UsingDeserializerToOption(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingDeserializerToOption(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C11_Deserialize_UsingDeserializerDefined(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func Test_C11_Deserialize_UsingDeserializerFuncDefined(t *testing.T) {
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &struct{}{})
	if err == nil { t.Fatal("expected error") }
	err2 := corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error { return nil }, &struct{}{})
	if err2 != nil { t.Fatal("unexpected") }
}

func Test_C11_Deserialize_UsingJsonerToAny(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
	err2 := corejson.Deserialize.UsingJsonerToAny(false, nil, &struct{}{})
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C11_Deserialize_UsingJsonerToAnyMust(t *testing.T) {
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &struct{}{})
	if err != nil { t.Fatal("expected nil") }
}

func Test_C15_Deserialize_UsingSerializerFuncTo(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(fn, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func Test_C15_Deserialize_ResultTo_ByteMust(t *testing.T) {
	r := corejson.Serialize.Apply(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)
	if b != 65 { t.Fatal("unexpected") }
}

func Test_C15_Deserialize_BytesTo_BytesMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
}
