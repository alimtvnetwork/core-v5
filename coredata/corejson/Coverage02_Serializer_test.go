package corejson

import (
	"testing"
)

func TestSerializer_Apply(t *testing.T) {
	r := Serialize.Apply("hello")
	if r.HasError() || r.JsonString() != `"hello"` {
		t.Fatal("unexpected")
	}
}

func TestSerializer_StringsApply(t *testing.T) {
	r := Serialize.StringsApply([]string{"a", "b"})
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func TestSerializer_FromBytes(t *testing.T) {
	r := Serialize.FromBytes([]byte(`"test"`))
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromStrings(t *testing.T) {
	r := Serialize.FromStrings([]string{"a"})
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromStringsSpread(t *testing.T) {
	r := Serialize.FromStringsSpread("a", "b")
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromString(t *testing.T) {
	r := Serialize.FromString("hello")
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromInteger(t *testing.T) {
	r := Serialize.FromInteger(42)
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromInteger64(t *testing.T) {
	r := Serialize.FromInteger64(64)
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_FromBool(t *testing.T) {
	r := Serialize.FromBool(true)
	if r.HasError() || r.JsonString() != "true" {
		t.Fatal("unexpected")
	}
}

func TestSerializer_FromIntegers(t *testing.T) {
	r := Serialize.FromIntegers([]int{1, 2, 3})
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_UsingAnyPtr(t *testing.T) {
	r := Serialize.UsingAnyPtr("x")
	if r.HasError() {
		t.Fatal(r.Error)
	}

	ch := make(chan int)
	r2 := Serialize.UsingAnyPtr(ch)
	if !r2.HasError() {
		t.Fatal("expected error")
	}
}

func TestSerializer_UsingAny(t *testing.T) {
	r := Serialize.UsingAny("x")
	if r.HasError() {
		t.Fatal(r.Error)
	}
}

func TestSerializer_Raw(t *testing.T) {
	b, err := Serialize.Raw("x")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestSerializer_Marshal(t *testing.T) {
	b, err := Serialize.Marshal("x")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestSerializer_ToBytesErr(t *testing.T) {
	b, err := Serialize.ToBytesErr("x")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestSerializer_ToBytesMust(t *testing.T) {
	b := Serialize.ToBytesMust("x")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestSerializer_ToSafeBytesMust(t *testing.T) {
	b := Serialize.ToSafeBytesMust("x")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestSerializer_ToSafeBytesSwallowErr(t *testing.T) {
	b := Serialize.ToSafeBytesSwallowErr("x")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestSerializer_ToBytesSwallowErr(t *testing.T) {
	b := Serialize.ToBytesSwallowErr("x")
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestSerializer_ToString(t *testing.T) {
	s := Serialize.ToString("hello")
	if s != `"hello"` {
		t.Fatal("unexpected")
	}
}

func TestSerializer_ToStringMust(t *testing.T) {
	s := Serialize.ToStringMust("x")
	if s == "" {
		t.Fatal("expected string")
	}
}

func TestSerializer_ToStringErr(t *testing.T) {
	s, err := Serialize.ToStringErr("x")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestSerializer_ToPrettyStringErr(t *testing.T) {
	s, err := Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestSerializer_ToPrettyStringIncludingErr(t *testing.T) {
	s := Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected string")
	}
}

func TestSerializer_Pretty(t *testing.T) {
	s := Serialize.Pretty(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected string")
	}
}

func TestSerializer_ApplyMust(t *testing.T) {
	r := Serialize.ApplyMust("x")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}
