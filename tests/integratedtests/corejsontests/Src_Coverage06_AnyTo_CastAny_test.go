package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"testing"
)

func TestAnyTo_SerializedJsonResult(t *testing.T) {
	// Result
	r := corejson.NewResult.Any("x")
	jr := corejson.AnyTo.SerializedJsonResult(r)
	if jr.HasError() {
		t.Fatal(jr.Error)
	}

	// *Result
	rp := corejson.NewResult.AnyPtr("x")
	jr2 := corejson.AnyTo.SerializedJsonResult(rp)
	if jr2.HasError() {
		t.Fatal(jr2.Error)
	}

	// []byte
	jr3 := corejson.AnyTo.SerializedJsonResult([]byte(`"x"`))
	if jr3.HasError() {
		t.Fatal(jr3.Error)
	}

	// string
	jr4 := corejson.AnyTo.SerializedJsonResult("hello")
	_ = jr4

	// nil
	jr5 := corejson.AnyTo.SerializedJsonResult(nil)
	if jr5.Error == nil {
		t.Fatal("expected error for nil")
	}

	// any
	jr6 := corejson.AnyTo.SerializedJsonResult(42)
	if jr6.HasError() {
		t.Fatal(jr6.Error)
	}
}

func TestAnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyTo_PrettyStringWithError(t *testing.T) {
	// string
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}

	// []byte
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" {
		t.Fatal("unexpected")
	}

	// Result
	r := corejson.NewResult.Any(42)
	s3, err3 := corejson.AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" {
		t.Fatal("unexpected")
	}

	// *Result
	rp := corejson.NewResult.AnyPtr(42)
	s4, err4 := corejson.AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" {
		t.Fatal("unexpected")
	}

	// any
	s5, err5 := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	if err5 != nil || s5 == "" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = corejson.AnyTo.SafeJsonPrettyString("hello")
	_ = corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func TestAnyTo_JsonString(t *testing.T) {
	_ = corejson.AnyTo.corejson.JsonString("hello")
	_ = corejson.AnyTo.corejson.JsonString([]byte("test"))
	_ = corejson.AnyTo.corejson.JsonString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.corejson.JsonString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.corejson.JsonString(42)
}

func TestAnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte("test"))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.Any(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.AnyPtr(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func TestAnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestAnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	_ = s
}

func TestAnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func TestCastAny_FromToDefault(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_WithReflection(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(true, "hello", &out)
	_ = err
}

func TestCastAny_FromToOption_Bytes(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_String(t *testing.T) {
	var out int
	err := corejson.CastAny.FromToOption(false, "42", &out)
	if err != nil || out != 42 {
		t.Fatal("unexpected")
	}
}

func TestCastAny_FromToOption_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCastAny_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) {
		return []byte(`"hello"`), nil
	}
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)
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
	err := corejson.CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.Name != "test" {
		t.Fatal("unexpected")
	}
}

func TestCastAny_OrDeserializeTo(t *testing.T) {
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hi"`), &out)
	if err != nil || out != "hi" {
		t.Fatal("unexpected")
	}
}
