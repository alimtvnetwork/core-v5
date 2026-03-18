package coredynamic

import (
	"reflect"
	"testing"
)

// === Type ===

func TestType_Func(t *testing.T) {
	rt := Type("hello")
	if rt != reflect.TypeOf("") { t.Fatal("expected string type") }
}

// === TypeSameStatus ===

func TestTypeSameStatus_Same(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	if !ts.IsSame { t.Fatal("expected same") }
	if ts.IsNotSame() { t.Fatal("expected not-not-same") }
	if ts.IsNotEqualTypes() { t.Fatal("expected equal") }
}

func TestTypeSameStatus_Different(t *testing.T) {
	ts := TypeSameStatus("a", 1)
	if ts.IsSame { t.Fatal("expected not same") }
	if !ts.IsNotSame() { t.Fatal("expected not-same") }
}

func TestTypeSameStatus_NilBoth(t *testing.T) {
	ts := TypeSameStatus(nil, nil)
	if !ts.IsSame { t.Fatal("expected same") }
}

func TestTypeSameStatus_NilOne(t *testing.T) {
	ts := TypeSameStatus(nil, "a")
	if ts.IsSame { t.Fatal("expected not same") }
}

// === TypeStatus ===

func TestTypeStatus_IsValid(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	if !ts.IsValid() { t.Fatal("expected valid") }
	if ts.IsInvalid() { t.Fatal("expected not invalid") }
	// Call again for cached
	if !ts.IsValid() { t.Fatal("expected valid") }
}

func TestTypeStatus_NilReceiver(t *testing.T) {
	var ts *TypeStatus
	if ts.IsValid() { t.Fatal("expected false") }
	if !ts.IsInvalid() { t.Fatal("expected true") }
}

func TestTypeStatus_Pointer(t *testing.T) {
	s := "hello"
	ts := TypeSameStatus(&s, &s)
	if !ts.IsAnyPointer() { t.Fatal("expected true") }
	if !ts.IsBothPointer() { t.Fatal("expected true") }
	_ = ts.NonPointerLeft()
	_ = ts.NonPointerRight()
}

func TestTypeStatus_NonPointer(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	if ts.IsAnyPointer() { t.Fatal("expected false") }
	if ts.IsBothPointer() { t.Fatal("expected false") }
	_ = ts.NonPointerLeft()
	_ = ts.NonPointerRight()
}

func TestTypeStatus_SameRegardlessPointer(t *testing.T) {
	s := "hello"
	ts := TypeSameStatus(&s, "world")
	if !ts.IsSameRegardlessPointer() { t.Fatal("expected true") }
	ts2 := TypeSameStatus("a", "b")
	if !ts2.IsSameRegardlessPointer() { t.Fatal("expected true") }
}

func TestTypeStatus_Names(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	if ts.LeftName() == "" { t.Fatal("expected non-empty") }
	if ts.RightName() == "" { t.Fatal("expected non-empty") }
	if ts.LeftFullName() == "" { t.Fatal("expected non-empty") }
	if ts.RightFullName() == "" { t.Fatal("expected non-empty") }
	tsNil := TypeSameStatus(nil, nil)
	_ = tsNil.LeftName()
	_ = tsNil.RightName()
	_ = tsNil.LeftFullName()
	_ = tsNil.RightFullName()
}

func TestTypeStatus_NotMatchMessage(t *testing.T) {
	ts := TypeSameStatus("a", 1)
	msg := ts.NotMatchMessage("left", "right")
	if msg == "" { t.Fatal("expected non-empty") }
	ts2 := TypeSameStatus("a", "b")
	if ts2.NotMatchMessage("l", "r") != "" { t.Fatal("expected empty") }
}

func TestTypeStatus_NotMatchErr(t *testing.T) {
	ts := TypeSameStatus("a", 1)
	err := ts.NotMatchErr("left", "right")
	if err == nil { t.Fatal("expected error") }
	ts2 := TypeSameStatus("a", "b")
	if ts2.NotMatchErr("l", "r") != nil { t.Fatal("expected nil") }
}

func TestTypeStatus_ValidationError(t *testing.T) {
	ts := TypeSameStatus("a", 1)
	if ts.ValidationError() == nil { t.Fatal("expected error") }
	ts2 := TypeSameStatus("a", "b")
	if ts2.ValidationError() != nil { t.Fatal("expected nil") }
}

func TestTypeStatus_MustBeSame(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	ts.MustBeSame() // should not panic
}

func TestTypeStatus_SrcDestination(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	ts.SrcDestinationMustBeSame() // should not panic
	msg := ts.NotEqualSrcDestinationMessage()
	if msg != "" { t.Fatal("expected empty") }
	err := ts.NotEqualSrcDestinationErr()
	if err != nil { t.Fatal("expected nil") }
}

func TestTypeStatus_IsEqual(t *testing.T) {
	ts1 := TypeSameStatus("a", "b")
	ts2 := TypeSameStatus("a", "b")
	if !ts1.IsEqual(&ts2) { t.Fatal("expected equal") }
	var nilTS *TypeStatus
	if !nilTS.IsEqual(nil) { t.Fatal("expected equal") }
	if nilTS.IsEqual(&ts1) { t.Fatal("expected not equal") }
	if ts1.IsEqual(nil) { t.Fatal("expected not equal") }
	ts3 := TypeSameStatus("a", 1)
	if ts1.IsEqual(&ts3) { t.Fatal("expected not equal") }
}

// === LeftRight ===

func TestLeftRight_Basic(t *testing.T) {
	lr := &LeftRight{Left: "a", Right: "b"}
	if lr.IsEmpty() { t.Fatal("expected non-empty") }
	if !lr.HasAnyItem() { t.Fatal("expected true") }
	if !lr.HasLeft() { t.Fatal("expected true") }
	if !lr.HasRight() { t.Fatal("expected true") }
	if lr.IsLeftEmpty() { t.Fatal("expected false") }
	if lr.IsRightEmpty() { t.Fatal("expected false") }
}

func TestLeftRight_NilReceiver(t *testing.T) {
	var lr *LeftRight
	if !lr.IsEmpty() { t.Fatal("expected empty") }
	if lr.HasAnyItem() { t.Fatal("expected false") }
	if lr.HasLeft() { t.Fatal("expected false") }
	if lr.HasRight() { t.Fatal("expected false") }
	if !lr.IsLeftEmpty() { t.Fatal("expected true") }
	if !lr.IsRightEmpty() { t.Fatal("expected true") }
}

func TestLeftRight_ReflectSet(t *testing.T) {
	lr := &LeftRight{Left: "hello", Right: "world"}
	var s string
	err := lr.LeftReflectSet(&s)
	if err != nil { t.Fatal("unexpected") }
	err2 := lr.RightReflectSet(&s)
	if err2 != nil { t.Fatal("unexpected") }
	var nilLR *LeftRight
	_ = nilLR.LeftReflectSet(&s)
	_ = nilLR.RightReflectSet(&s)
}

func TestLeftRight_Deserialize(t *testing.T) {
	lr := &LeftRight{Left: "hello", Right: "world"}
	_ = lr.DeserializeLeft()
	_ = lr.DeserializeRight()
	var nilLR *LeftRight
	if nilLR.DeserializeLeft() != nil { t.Fatal("expected nil") }
	if nilLR.DeserializeRight() != nil { t.Fatal("expected nil") }
}

func TestLeftRight_ToDynamic(t *testing.T) {
	lr := &LeftRight{Left: "a", Right: "b"}
	_ = lr.LeftToDynamic()
	_ = lr.RightToDynamic()
	var nilLR *LeftRight
	if nilLR.LeftToDynamic() != nil { t.Fatal("expected nil") }
	if nilLR.RightToDynamic() != nil { t.Fatal("expected nil") }
}

func TestLeftRight_TypeStatus(t *testing.T) {
	lr := &LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	if !ts.IsSame { t.Fatal("expected same") }
	var nilLR *LeftRight
	ts2 := nilLR.TypeStatus()
	if !ts2.IsSame { t.Fatal("expected same for nil") }
}

// === TypedDynamic ===

func TestTypedDynamic_String(t *testing.T) {
	td := NewTypedDynamic[string]("hello", true)
	if td.Data() != "hello" { t.Fatal("expected hello") }
	if td.Value() != "hello" { t.Fatal("expected hello") }
	if !td.IsValid() { t.Fatal("expected valid") }
	if td.IsInvalid() { t.Fatal("expected not invalid") }
	if td.String() != "hello" { t.Fatal("expected hello") }
	if td.ValueString() != "hello" { t.Fatal("expected hello") }
}

func TestTypedDynamic_Int(t *testing.T) {
	td := NewTypedDynamicValid[int](42)
	if td.ValueInt() != 42 { t.Fatal("expected 42") }
	if td.ValueBool() { t.Fatal("expected false") }
	if td.ValueInt64() != -1 { t.Fatal("expected -1") }
}

func TestTypedDynamic_Invalid(t *testing.T) {
	td := InvalidTypedDynamic[string]()
	if !td.IsInvalid() { t.Fatal("expected invalid") }
	tdp := InvalidTypedDynamicPtr[int]()
	if tdp == nil { t.Fatal("expected non-nil") }
}

func TestTypedDynamic_Ptr(t *testing.T) {
	td := NewTypedDynamicPtr[string]("hello", true)
	if td == nil { t.Fatal("expected non-nil") }
}

func TestTypedDynamic_Clone(t *testing.T) {
	td := NewTypedDynamic[string]("hello", true)
	c := td.Clone()
	if c.Data() != "hello" { t.Fatal("expected hello") }
	_ = td.NonPtr()
	tdp := NewTypedDynamicPtr[string]("hello", true)
	cp := tdp.ClonePtr()
	if cp == nil { t.Fatal("expected non-nil") }
	_ = tdp.Ptr()
	var nilTD *TypedDynamic[string]
	if nilTD.ClonePtr() != nil { t.Fatal("expected nil") }
}

func TestTypedDynamic_JSON(t *testing.T) {
	td := NewTypedDynamic[string]("hello", true)
	b, err := td.JsonBytes()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
	_ = td.JsonResult()
	_ = td.Json()
	_ = td.JsonPtr()
	s, err2 := td.JsonString()
	if err2 != nil || s == "" { t.Fatal("unexpected") }
	_, _ = td.MarshalJSON()
	_, _ = td.ValueMarshal()
	_ = td.JsonModel()
	_ = td.JsonModelAny()
}

func TestTypedDynamic_UnmarshalJSON(t *testing.T) {
	td := &TypedDynamic[string]{}
	err := td.UnmarshalJSON([]byte(`"hello"`))
	if err != nil || td.Data() != "hello" { t.Fatal("unexpected") }
}

func TestTypedDynamic_Deserialize(t *testing.T) {
	td := &TypedDynamic[string]{}
	err := td.Deserialize([]byte(`"hello"`))
	if err != nil { t.Fatal("unexpected") }
	var nilTD *TypedDynamic[string]
	err2 := nilTD.Deserialize([]byte(`"x"`))
	if err2 == nil { t.Fatal("expected error") }
}

func TestTypedDynamic_Bytes(t *testing.T) {
	td := NewTypedDynamic[[]byte]([]byte("hello"), true)
	b, ok := td.Bytes()
	if !ok || len(b) == 0 { t.Fatal("unexpected") }
	td2 := NewTypedDynamic[string]("hello", true)
	b2, ok2 := td2.Bytes()
	if !ok2 || len(b2) == 0 { t.Fatal("unexpected") }
}

func TestTypedDynamic_GetAs(t *testing.T) {
	td := NewTypedDynamic[string]("hello", true)
	s, ok := td.GetAsString()
	if !ok || s != "hello" { t.Fatal("unexpected") }
	_, ok2 := td.GetAsInt()
	if ok2 { t.Fatal("expected false") }
	_, ok3 := td.GetAsInt64()
	if ok3 { t.Fatal("expected false") }
	_, ok4 := td.GetAsUint()
	if ok4 { t.Fatal("expected false") }
	_, ok5 := td.GetAsFloat64()
	if ok5 { t.Fatal("expected false") }
	_, ok6 := td.GetAsFloat32()
	if ok6 { t.Fatal("expected false") }
	_, ok7 := td.GetAsBool()
	if ok7 { t.Fatal("expected false") }
	_, ok8 := td.GetAsBytes()
	if ok8 { t.Fatal("expected false") }
	_, ok9 := td.GetAsStrings()
	if ok9 { t.Fatal("expected false") }
}

func TestTypedDynamic_ToDynamic(t *testing.T) {
	td := NewTypedDynamic[string]("hello", true)
	d := td.ToDynamic()
	if d.IsInvalid() { t.Fatal("expected valid") }
}

// === SimpleRequest ===

func TestSimpleRequest(t *testing.T) {
	sr := NewSimpleRequestValid("hello")
	if sr.Request() != "hello" { t.Fatal("expected hello") }
	if sr.Value() != "hello" { t.Fatal("expected hello") }
	if sr.Message() != "" { t.Fatal("expected empty") }
	if sr.InvalidError() != nil { t.Fatal("expected nil") }
	if !sr.IsValid() { t.Fatal("expected valid") }
}

func TestSimpleRequest_Invalid(t *testing.T) {
	sr := InvalidSimpleRequest("test error")
	if sr.IsValid() { t.Fatal("expected invalid") }
	if sr.Message() != "test error" { t.Fatal("expected test error") }
	err := sr.InvalidError()
	if err == nil { t.Fatal("expected error") }
	// Call again for caching
	_ = sr.InvalidError()
	sr2 := InvalidSimpleRequestNoMessage()
	if sr2.InvalidError() != nil { t.Fatal("expected nil") }
}

func TestSimpleRequest_New(t *testing.T) {
	sr := NewSimpleRequest("hello", true, "")
	if sr.Request() != "hello" { t.Fatal("expected hello") }
}

func TestSimpleRequest_TypeMismatch(t *testing.T) {
	sr := NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err != nil { t.Fatal("expected nil") }
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err2 == nil { t.Fatal("expected error") }
	err3 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	if err3 == nil { t.Fatal("expected error") }
}

func TestSimpleRequest_IsPointer(t *testing.T) {
	sr := NewSimpleRequestValid("hello")
	if sr.IsPointer() { t.Fatal("expected false") }
	// call again for caching
	if sr.IsPointer() { t.Fatal("expected false") }
}

func TestSimpleRequest_IsReflectKind(t *testing.T) {
	sr := NewSimpleRequestValid("hello")
	if !sr.IsReflectKind(reflect.String) { t.Fatal("expected true") }
}

// === SimpleResult ===

func TestSimpleResult(t *testing.T) {
	sr := NewSimpleResultValid("hello")
	if sr.Result != "hello" { t.Fatal("expected hello") }
	if sr.Message != "" { t.Fatal("expected empty") }
	if sr.InvalidError() != nil { t.Fatal("expected nil") }
}

func TestSimpleResult_Invalid(t *testing.T) {
	sr := InvalidSimpleResult("test error")
	if sr.IsValid() { t.Fatal("expected invalid") }
	err := sr.InvalidError()
	if err == nil { t.Fatal("expected error") }
	_ = sr.InvalidError() // cached
	sr2 := InvalidSimpleResultNoMessage()
	if sr2.InvalidError() != nil { t.Fatal("expected nil") }
}

func TestSimpleResult_New(t *testing.T) {
	sr := NewSimpleResult("hello", true, "")
	if sr.Result != "hello" { t.Fatal("expected hello") }
}

func TestSimpleResult_TypeMismatch(t *testing.T) {
	sr := NewSimpleResultValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err != nil { t.Fatal("expected nil") }
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err2 == nil { t.Fatal("expected error") }
	err3 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	if err3 == nil { t.Fatal("expected error") }
}

func TestSimpleResult_Clone(t *testing.T) {
	sr := NewSimpleResultValid("hello")
	c := sr.Clone()
	if c.Result != "hello" { t.Fatal("expected hello") }
	cp := sr.ClonePtr()
	if cp == nil { t.Fatal("expected non-nil") }
	var nilSR *SimpleResult
	if nilSR.ClonePtr() != nil { t.Fatal("expected nil") }
}

// === AnyToReflectVal ===

func TestAnyToReflectVal(t *testing.T) {
	rv := AnyToReflectVal("hello")
	if !rv.IsValid() { t.Fatal("expected valid") }
}
