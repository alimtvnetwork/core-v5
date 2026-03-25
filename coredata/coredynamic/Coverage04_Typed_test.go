package coredynamic

import (
	"testing"
)

// ── TypedDynamic ──

func TestTypedDynamic(t *testing.T) {
	d := NewTypedDynamic("hello", true)
	if d.Data() != "hello" || d.Value() != "hello" { t.Fatal("unexpected") }
	if !d.IsValid() || d.IsInvalid() { t.Fatal("unexpected") }
	if d.String() != "hello" { t.Fatal("unexpected") }
	_, _ = d.JsonBytes()
	_ = d.JsonResult()
	_ = d.Json()
	_ = d.JsonPtr()
	_, _ = d.JsonString()
	_, _ = d.MarshalJSON()
	_, _ = d.ValueMarshal()
	_, _ = d.Bytes()
	_ = d.JsonModel()
	_ = d.JsonModelAny()
	_ = d.Clone()
	_ = d.ClonePtr()
	_ = d.NonPtr()
	_ = d.Ptr()
	_ = d.ToDynamic()
	s, ok := d.GetAsString()
	if !ok || s != "hello" { t.Fatal("unexpected") }
	_, _ = d.GetAsInt()
	_, _ = d.GetAsInt64()
	_, _ = d.GetAsUint()
	_, _ = d.GetAsFloat64()
	_, _ = d.GetAsFloat32()
	_, _ = d.GetAsBool()
	_, _ = d.GetAsBytes()
	_, _ = d.GetAsStrings()
	_ = d.ValueString()
	_ = d.ValueInt()
	_ = d.ValueInt64()
	_ = d.ValueBool()

	inv := InvalidTypedDynamic[string]()
	if inv.IsValid() { t.Fatal("expected invalid") }
	invp := InvalidTypedDynamicPtr[string]()
	if invp.IsValid() { t.Fatal("expected invalid") }
	var nilTd *TypedDynamic[string]
	if nilTd.ClonePtr() != nil { t.Fatal("expected nil") }
}

func TestTypedDynamic_Valid(t *testing.T) {
	d := NewTypedDynamicValid("hello")
	if !d.IsValid() { t.Fatal("expected valid") }
	dp := NewTypedDynamicPtr("hello", true)
	if !dp.IsValid() { t.Fatal("expected valid") }
}

func TestTypedDynamic_Deserialize(t *testing.T) {
	d := NewTypedDynamic("", true)
	err := d.Deserialize([]byte(`"world"`))
	if err != nil || d.Value() != "world" { t.Fatal("unexpected") }
	var nilTd *TypedDynamic[string]
	if nilTd.Deserialize(nil) == nil { t.Fatal("expected error") }
}

func TestTypedDynamic_BytesFromNonBytes(t *testing.T) {
	d := NewTypedDynamic(42, true)
	b, ok := d.Bytes()
	if !ok || len(b) == 0 { t.Fatal("expected json bytes") }
}

// ── TypedSimpleRequest ──

func TestTypedSimpleRequest(t *testing.T) {
	r := NewTypedSimpleRequest("hello", true, "")
	if r.Data() != "hello" || r.Request() != "hello" || r.Value() != "hello" { t.Fatal("unexpected") }
	if !r.IsValid() || r.IsInvalid() { t.Fatal("unexpected") }
	_ = r.Message()
	_ = r.String()
	_ = r.InvalidError()
	_, _ = r.JsonBytes()
	_ = r.JsonResult()
	_ = r.Json()
	_ = r.JsonPtr()
	_, _ = r.MarshalJSON()
	_ = r.JsonModel()
	_ = r.JsonModelAny()
	_ = r.Clone()
	_ = r.ToSimpleRequest()
	_ = r.ToTypedDynamic()
	_ = r.ToDynamic()
	s, ok := r.GetAsString()
	if !ok || s != "hello" { t.Fatal("unexpected") }
	_, _ = r.GetAsInt()
	_, _ = r.GetAsInt64()
	_, _ = r.GetAsFloat64()
	_, _ = r.GetAsFloat32()
	_, _ = r.GetAsBool()
	_, _ = r.GetAsBytes()
	_, _ = r.GetAsStrings()

	v := NewTypedSimpleRequestValid("test")
	if !v.IsValid() { t.Fatal("expected valid") }
	inv := InvalidTypedSimpleRequest[string]("msg")
	if inv.IsValid() { t.Fatal("expected invalid") }
	_ = inv.InvalidError()
	inv2 := InvalidTypedSimpleRequestNoMessage[string]()
	_ = inv2

	var nilR *TypedSimpleRequest[string]
	if nilR.IsValid() { t.Fatal("expected invalid") }
	if !nilR.IsInvalid() { t.Fatal("expected invalid") }
	if nilR.Message() != "" { t.Fatal("expected empty") }
	if nilR.String() != "" { t.Fatal("expected empty") }
	if nilR.InvalidError() != nil { t.Fatal("expected nil") }
	if nilR.Clone() != nil { t.Fatal("expected nil") }
	_ = nilR.ToSimpleRequest()
	_ = nilR.ToTypedDynamic()
	_ = nilR.ToDynamic()
}

// ── TypedSimpleResult ──

func TestTypedSimpleResult(t *testing.T) {
	r := NewTypedSimpleResult("hello", true, "")
	if r.Data() != "hello" || r.Result() != "hello" { t.Fatal("unexpected") }
	if !r.IsValid() || r.IsInvalid() { t.Fatal("unexpected") }
	_ = r.Message()
	_ = r.String()
	_ = r.InvalidError()
	_, _ = r.JsonBytes()
	_ = r.JsonResult()
	_ = r.Json()
	_ = r.JsonPtr()
	_, _ = r.MarshalJSON()
	_ = r.JsonModel()
	_ = r.JsonModelAny()
	_ = r.Clone()
	_ = r.ClonePtr()
	_ = r.ToSimpleResult()
	_ = r.ToTypedDynamic()
	_ = r.ToDynamic()
	s, ok := r.GetAsString()
	if !ok || s != "hello" { t.Fatal("unexpected") }
	_, _ = r.GetAsInt()
	_, _ = r.GetAsInt64()
	_, _ = r.GetAsFloat64()
	_, _ = r.GetAsBool()
	_, _ = r.GetAsBytes()
	_, _ = r.GetAsStrings()

	v := NewTypedSimpleResultValid("test")
	if !v.IsValid() { t.Fatal("expected valid") }
	inv := InvalidTypedSimpleResult[string]("msg")
	if inv.IsValid() { t.Fatal("expected invalid") }
	_ = inv.InvalidError()
	inv2 := InvalidTypedSimpleResultNoMessage[string]()
	_ = inv2

	var nilR *TypedSimpleResult[string]
	if nilR.IsValid() { t.Fatal("expected invalid") }
	if nilR.Message() != "" { t.Fatal("expected empty") }
	if nilR.String() != "" { t.Fatal("expected empty") }
	if nilR.InvalidError() != nil { t.Fatal("expected nil") }
	if nilR.ClonePtr() != nil { t.Fatal("expected nil") }
	_ = nilR.ToSimpleResult()
	_ = nilR.ToTypedDynamic()
	_ = nilR.ToDynamic()
}

// ── BytesConverter ──

func TestBytesConverter(t *testing.T) {
	bc := NewBytesConverter([]byte(`"hello"`))
	var s string
	_ = bc.Deserialize(&s)
	_ = bc.SafeCastString()
	_, _ = bc.CastString()
	_, _ = bc.ToString()
	_ = bc.ToStringMust()

	bc2 := NewBytesConverter([]byte(`["a","b"]`))
	_, _ = bc2.ToStrings()
	_ = bc2.ToStringsMust()
	_, _ = bc2.ToCollection()
	_, _ = bc2.ToSimpleSlice()

	bc3 := NewBytesConverter([]byte(`true`))
	_, _ = bc3.ToBool()
	_ = bc3.ToBoolMust()

	bc4 := NewBytesConverter([]byte(`42`))
	_, _ = bc4.ToInt64()
	_ = bc4.ToInt64Must()

	bc5 := NewBytesConverter([]byte(`{"a":"b"}`))
	_, _ = bc5.ToHashmap()
	_, _ = bc5.ToMapAnyItems()

	bc6 := NewBytesConverter([]byte(`{"a":true}`))
	_, _ = bc6.ToHashset()

	// empty
	bc7 := NewBytesConverter([]byte{})
	_, _ = bc7.CastString()
	_ = bc7.SafeCastString()
}

// ── Misc functions ──

func TestAnyToReflectVal(t *testing.T) {
	rv := AnyToReflectVal(42)
	if rv.Int() != 42 { t.Fatal("expected 42") }
}

func TestIsAnyTypesOf(t *testing.T) {
	if !IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(""), reflect.TypeOf(0)) { t.Fatal("expected true") }
	if IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(0)) { t.Fatal("expected false") }
}

func TestLengthOfReflect(t *testing.T) {
	if LengthOfReflect(reflect.ValueOf([]int{1, 2})) != 2 { t.Fatal("expected 2") }
	if LengthOfReflect(reflect.ValueOf("hello")) != 0 { t.Fatal("expected 0") }
}

func TestSafeTypeName(t *testing.T) {
	_ = SafeTypeName(nil)
	_ = SafeTypeName("hello")
}
