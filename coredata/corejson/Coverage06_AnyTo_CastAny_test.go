package corejson

import (
	"testing"
)

func TestAnyTo_SerializedJsonResult(t *testing.T) {
	// Result
	r := NewResult.Any("x")
	jr := AnyTo.SerializedJsonResult(r)
	if jr.HasError() {
		t.Fatal(jr.Error)
	}

	// *Result
	rp := NewResult.AnyPtr("x")
	jr2 := AnyTo.SerializedJsonResult(rp)
	if jr2.HasError() {
		t.Fatal(jr2.Error)
	}

	// []byte
	jr3 := AnyTo.SerializedJsonResult([]byte(`"x"`))
	if jr3.HasError() {
		t.Fatal(jr3.Error)
	}

	// string
	jr4 := AnyTo.SerializedJsonResult("hello")
	_ = jr4

	// nil
	jr5 := AnyTo.SerializedJsonResult(nil)
	if jr5.Error == nil {
		t.Fatal("expected error for nil")
	}

	// any
	jr6 := AnyTo.SerializedJsonResult(42)
	if jr6.HasError() {
		t.Fatal(jr6.Error)
	}
}

func TestAnyTo_SerializedRaw(t *testing.T) {
	b, err := AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SerializedString(t *testing.T) {
	s, err := AnyTo.SerializedString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SerializedSafeString(t *testing.T) {
	s := AnyTo.SerializedSafeString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_SerializedStringMust(t *testing.T) {
	s := AnyTo.SerializedStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_SafeJsonString(t *testing.T) {
	s := AnyTo.SafeJsonString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_PrettyStringWithError(t *testing.T) {
	// string
	s, err := AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}

	// []byte
	s2, err2 := AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" {
		t.Fatal("unexpected")
	}

	// Result
	r := NewResult.Any(42)
	s3, err3 := AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" {
		t.Fatal("unexpected")
	}

	// *Result
	rp := NewResult.AnyPtr(42)
	s4, err4 := AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" {
		t.Fatal("unexpected")
	}

	// any
	s5, err5 := AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	if err5 != nil || s5 == "" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = AnyTo.SafeJsonPrettyString("hello")
	_ = AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = AnyTo.SafeJsonPrettyString(NewResult.Any(1))
	_ = AnyTo.SafeJsonPrettyString(NewResult.AnyPtr(1))
	_ = AnyTo.SafeJsonPrettyString(42)
}

func TestAnyTo_JsonString(t *testing.T) {
	_ = AnyTo.JsonString("hello")
	_ = AnyTo.JsonString([]byte("test"))
	_ = AnyTo.JsonString(NewResult.Any(1))
	_ = AnyTo.JsonString(NewResult.AnyPtr(1))
	_ = AnyTo.JsonString(42)
}

func TestAnyTo_JsonStringWithErr(t *testing.T) {
	s, err := AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
	_, _ = AnyTo.JsonStringWithErr([]byte("test"))
	_, _ = AnyTo.JsonStringWithErr(NewResult.Any(1))
	_, _ = AnyTo.JsonStringWithErr(NewResult.AnyPtr(1))
	_, _ = AnyTo.JsonStringWithErr(42)
}

func TestAnyTo_JsonStringMust(t *testing.T) {
	s := AnyTo.JsonStringMust("hello")
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_PrettyStringMust(t *testing.T) {
	s := AnyTo.PrettyStringMust("hello")
	_ = s
}

func TestAnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func TestCastAny_FromToDefault(t *testing.T) {
	var out string
	err := CastAny.FromToDefault([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_WithReflection(t *testing.T) {
	var out string
	err := CastAny.FromToOption(true, "hello", &out)
	_ = err
}

func TestCastAny_FromToOption_Bytes(t *testing.T) {
	var out string
	err := CastAny.FromToOption(false, []byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_String(t *testing.T) {
	var out int
	err := CastAny.FromToOption(false, "42", &out)
	if err != nil || out != 42 {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_Result(t *testing.T) {
	r := NewResult.Any("hello")
	var out string
	err := CastAny.FromToOption(false, r, &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCastAny_FromToOption_ResultPtr(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	var out string
	err := CastAny.FromToOption(false, r, &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) {
		return []byte(`"hello"`), nil
	}
	var out string
	err := CastAny.FromToOption(false, fn, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_AnyFallback(t *testing.T) {
	type simple struct {
		Name string
	}
	src := simple{Name: "test"}
	var dst simple
	err := CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.Name != "test" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_OrDeserializeTo(t *testing.T) {
	var out string
	err := CastAny.OrDeserializeTo([]byte(`"hi"`), &out)
	if err != nil || out != "hi" {
		t.Fatal("unexpected")
	}
}
