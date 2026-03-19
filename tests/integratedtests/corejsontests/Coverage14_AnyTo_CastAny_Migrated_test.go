package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Migrated from Coverage06_AnyTo_CastAny_test.go and Coverage11 ──

func Test_C06_AnyTo_SerializedJsonResult(t *testing.T) {
	r := corejson.NewResult.Any("x")
	jr := corejson.AnyTo.SerializedJsonResult(r)
	if jr.HasError() { t.Fatal(jr.Error) }

	rp := corejson.NewResult.AnyPtr("x")
	jr2 := corejson.AnyTo.SerializedJsonResult(rp)
	if jr2.HasError() { t.Fatal(jr2.Error) }

	jr3 := corejson.AnyTo.SerializedJsonResult([]byte(`"x"`))
	if jr3.HasError() { t.Fatal(jr3.Error) }

	jr4 := corejson.AnyTo.SerializedJsonResult("hello")
	_ = jr4

	jr5 := corejson.AnyTo.SerializedJsonResult(nil)
	if jr5.Error == nil { t.Fatal("expected error for nil") }

	jr6 := corejson.AnyTo.SerializedJsonResult(42)
	if jr6.HasError() { t.Fatal(jr6.Error) }

	jr7 := corejson.AnyTo.SerializedJsonResult(errors.New("oops"))
	if jr7 == nil { t.Fatal("expected non-nil") }
}

func Test_C06_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func Test_C06_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
	_, err2 := corejson.AnyTo.SerializedString(nil)
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C06_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	if s == "" { t.Fatal("expected non-empty") }
	s2 := corejson.AnyTo.SerializedSafeString(nil)
	if s2 != "" { t.Fatal("expected empty") }
}

func Test_C06_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func Test_C06_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func Test_C06_AnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" { t.Fatal("unexpected") }
	r := corejson.NewResult.Any(42)
	s3, err3 := corejson.AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" { t.Fatal("unexpected") }
	rp := corejson.NewResult.AnyPtr(42)
	s4, err4 := corejson.AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" { t.Fatal("unexpected") }

	re := corejson.Result{Error: errors.New("e")}
	_, err5 := corejson.AnyTo.PrettyStringWithError(re)
	if err5 == nil { t.Fatal("expected error") }
	rep := &corejson.Result{Error: errors.New("e")}
	_, err6 := corejson.AnyTo.PrettyStringWithError(rep)
	if err6 == nil { t.Fatal("expected error") }

	s5, err5b := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	if err5b != nil || s5 == "" { t.Fatal("unexpected") }
}

func Test_C06_AnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = corejson.AnyTo.SafeJsonPrettyString("hello")
	_ = corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func Test_C06_AnyTo_JsonString(t *testing.T) {
	_ = corejson.AnyTo.JsonString("hello")
	_ = corejson.AnyTo.JsonString([]byte("test"))
	_ = corejson.AnyTo.JsonString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.JsonString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.JsonString(42)
}

func Test_C06_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" { t.Fatal("unexpected") }
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte("test"))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.Any(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.AnyPtr(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(42)

	_, err2 := corejson.AnyTo.JsonStringWithErr(corejson.Result{Error: errors.New("e")})
	if err2 == nil { t.Fatal("expected error") }
	_, err3 := corejson.AnyTo.JsonStringWithErr(&corejson.Result{Error: errors.New("e")})
	if err3 == nil { t.Fatal("expected error") }
}

func Test_C06_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	if s != "hello" { t.Fatal("unexpected") }
}

func Test_C06_AnyTo_PrettyStringMust(t *testing.T) {
	_ = corejson.AnyTo.PrettyStringMust("hello")
}

func Test_C06_AnyTo_SerializedFieldsMap(t *testing.T) {
	_, _ = corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
}

func Test_C06_CastAny_FromToDefault(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToOption_Bytes(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToOption_String(t *testing.T) {
	var out int
	err := corejson.CastAny.FromToOption(false, "42", &out)
	if err != nil || out != 42 { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToOption_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	if err != nil { t.Fatal(err) }
}

func Test_C06_CastAny_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	if err != nil { t.Fatal(err) }
}

func Test_C06_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToOption_AnyFallback(t *testing.T) {
	type simple struct{ Name string }
	src := simple{Name: "test"}
	var dst simple
	err := corejson.CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.Name != "test" { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToOption_WithReflection(t *testing.T) {
	var out string
	_ = corejson.CastAny.FromToOption(true, "hello", &out)
}

func Test_C06_CastAny_OrDeserializeTo(t *testing.T) {
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hi"`), &out)
	if err != nil || out != "hi" { t.Fatal("unexpected") }
}

func Test_C06_CastAny_FromToReflection(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)
	if err != nil || out != "hello" { t.Fatal("unexpected") }
}

type testSerializer14 struct{}
func (testSerializer14) Serialize() ([]byte, error) { return []byte(`"x"`), nil }

func Test_C06_AnyTo_UsingSerializer_Alt(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(testSerializer14{})
	if r == nil || r.HasError() { t.Fatal("unexpected") }
}

func Test_C06_AnyTo_PrettyStringMust_Map(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust(map[string]string{"a": "1"})
	if s == "" { t.Fatal("expected non-empty") }
}
