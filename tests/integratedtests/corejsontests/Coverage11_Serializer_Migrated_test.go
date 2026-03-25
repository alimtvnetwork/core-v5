package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Migrated from Coverage02_Serializer_test.go ──

func Test_C02_Serializer_Apply(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	if r.HasError() || r.JsonString() != `"hello"` { t.Fatal("unexpected") }
}

func Test_C02_Serializer_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	if r.HasError() { t.Fatal("unexpected error") }
}

func Test_C02_Serializer_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte(`"test"`))
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(64)
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	if r.HasError() || r.JsonString() != "true" { t.Fatal("unexpected") }
}

func Test_C02_Serializer_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2, 3})
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("x")
	if r.HasError() { t.Fatal(r.Error) }
	ch := make(chan int)
	r2 := corejson.Serialize.UsingAnyPtr(ch)
	if !r2.HasError() { t.Fatal("expected error") }
}

func Test_C02_Serializer_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("x")
	if r.HasError() { t.Fatal(r.Error) }
}

func Test_C02_Serializer_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("x")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func Test_C02_Serializer_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("x")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func Test_C02_Serializer_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("x")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func Test_C02_Serializer_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("x")
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func Test_C02_Serializer_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("x")
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func Test_C02_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("x")
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func Test_C02_Serializer_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("x")
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func Test_C02_Serializer_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	if s != `"hello"` { t.Fatal("unexpected") }
}

func Test_C02_Serializer_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("x")
	if s == "" { t.Fatal("expected string") }
}

func Test_C02_Serializer_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("x")
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func Test_C02_Serializer_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func Test_C02_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	if s == "" { t.Fatal("expected string") }
}

func Test_C02_Serializer_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	if s == "" { t.Fatal("expected string") }
}

func Test_C02_Serializer_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("x")
	if r.HasError() { t.Fatal("unexpected error") }
}
