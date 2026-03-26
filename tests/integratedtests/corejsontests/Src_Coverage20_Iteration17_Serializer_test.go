package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"fmt"
	"testing"
)

// Covers: serializerLogic methods

func Test_I17_Serialize_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	if r.HasError() {
		t.Fatal("unexpected error")
	}
	if string(r.Bytes) != `["a","b"]` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte(`hello`))
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Serialize_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"x"})
	if r.HasError() || string(r.Bytes) != `["x"]` {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Serialize_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Serialize_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	if string(r.Bytes) != `"hello"` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	if string(r.Bytes) != `42` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(99)
	if string(r.Bytes) != `99` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	if string(r.Bytes) != `true` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2, 3})
	if string(r.Bytes) != `[1,2,3]` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_I17_Serialize_FromStringer(t *testing.T) {
	r := corejson.Serialize.FromStringer(testStringer{val: "test"})
	if string(r.Bytes) != `"test"` {
		t.Fatalf("unexpected: %s", string(r.Bytes))
	}
}

func Test_I17_Serialize_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr(map[string]int{"a": 1})
	if r.HasError() {
		t.Fatal("unexpected error")
	}

	// error path
	rFail := corejson.Serialize.UsingAnyPtr(func() {})
	if !rFail.HasError() {
		t.Fatal("expected error for func type")
	}
}

func Test_I17_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("test")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Serialize_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("test")
	if err != nil || string(b) != `"test"` {
		t.Fatal("unexpected raw result")
	}
}

func Test_I17_Serialize_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal(42)
	if err != nil || string(b) != `42` {
		t.Fatal("unexpected marshal result")
	}
}

func Test_I17_Serialize_ApplyMust_Success(t *testing.T) {
	r := corejson.Serialize.ApplyMust("ok")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_Serialize_ApplyMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for unmarshalable")
		}
	}()
	corejson.Serialize.ApplyMust(func() {})
}

func Test_I17_Serialize_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("hello")
	if string(b) != `"hello"` {
		t.Fatalf("unexpected: %s", string(b))
	}
}

func Test_I17_Serialize_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("hello")
	if len(b) == 0 {
		t.Fatal("expected non-empty safe bytes")
	}
}

func Test_I17_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr(func() {})
	if len(b) != 0 {
		t.Fatal("expected empty for error case")
	}

	b2 := corejson.Serialize.ToSafeBytesSwallowErr("ok")
	if len(b2) == 0 {
		t.Fatal("expected non-empty for valid case")
	}
}

func Test_I17_Serialize_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("ok")
	if string(b) != `"ok"` {
		t.Fatalf("unexpected: %s", string(b))
	}
}

func Test_I17_Serialize_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("ok")
	if err != nil || string(b) != `"ok"` {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	if s != `"hello"` {
		t.Fatalf("unexpected: %s", s)
	}
}

func Test_I17_Serialize_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("hello")
	if s != `"hello"` {
		t.Fatalf("unexpected: %s", s)
	}
}

func Test_I17_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	if err != nil || s != `"hello"` {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	if err != nil || s == "" {
		t.Fatal("unexpected result")
	}
}

func Test_I17_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty pretty string")
	}
}

func Test_I17_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty pretty string")
	}
}

func Test_I17_Serialize_Apply_Error(t *testing.T) {
	r := corejson.Serialize.Apply(func() {})
	if !r.HasError() {
		t.Fatal("expected error")
	}
	if r.TypeName == "" {
		t.Fatal("expected type name")
	}
}

// Covers fmt.Stringer interface usage
var _ fmt.Stringer = testStringer{}
