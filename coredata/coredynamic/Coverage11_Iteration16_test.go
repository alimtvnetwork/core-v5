package coredynamic

import (
	"reflect"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_AnyCollection_EmptyCollection(t *testing.T) {
	ac := EmptyAnyCollection()
	if ac == nil {
		t.Fatal("expected non-nil")
	}
	if !ac.IsEmpty() {
		t.Fatal("expected empty")
	}
	if ac.Items() == nil {
		t.Fatal("expected non-nil slice")
	}
	if ac.DynamicItems() == nil {
		t.Fatal("expected non-nil slice")
	}
	dc := ac.DynamicCollection()
	if dc == nil {
		t.Fatal("expected non-nil")
	}
	if ac.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if ac.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I16_AnyCollection_WithItems(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("hello")
	ac.Add(42)

	if ac.Length() != 2 {
		t.Fatal("expected 2")
	}
	if ac.At(0) != "hello" {
		t.Fatal("expected hello")
	}
	d := ac.AtAsDynamic(0)
	if !d.IsValid() {
		t.Fatal("expected valid")
	}
	if ac.First() != "hello" {
		t.Fatal("expected hello")
	}
	if ac.Last() != 42 {
		t.Fatal("expected 42")
	}
	if ac.FirstDynamic() != "hello" {
		t.Fatal("expected hello")
	}
	if ac.LastDynamic() != 42 {
		t.Fatal("expected 42")
	}
	if ac.FirstOrDefault() != "hello" {
		t.Fatal("expected hello")
	}
	if ac.LastOrDefault() != 42 {
		t.Fatal("expected 42")
	}
	dc := ac.DynamicCollection()
	if dc == nil || dc.Length() != 2 {
		t.Fatal("expected 2 dynamic items")
	}
}

func Test_I16_AnyCollection_ReflectSetAt(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("test")
	var s string
	err := ac.ReflectSetAt(0, &s)
	if err != nil {
		t.Fatal(err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// CastedResult — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_CastedResult_NilReceiver(t *testing.T) {
	var cr *CastedResult
	if !cr.IsInvalid() {
		t.Fatal("expected invalid for nil")
	}
	if cr.IsNotNull() {
		t.Fatal("expected false for nil")
	}
	if cr.IsNotPointer() {
		t.Fatal("expected false for nil")
	}
	if cr.IsNotMatchingAcceptedType() {
		t.Fatal("expected false for nil")
	}
	if cr.IsSourceKind(reflect.String) {
		t.Fatal("expected false for nil")
	}
	if cr.HasError() {
		t.Fatal("expected false for nil")
	}
}

func Test_I16_CastedResult_Valid(t *testing.T) {
	cr := &CastedResult{
		IsValid:              true,
		IsNull:               false,
		IsMatchingAcceptedType: true,
		IsPointer:            true,
		SourceKind:           reflect.String,
	}
	if cr.IsInvalid() {
		t.Fatal("expected valid")
	}
	if !cr.IsNotNull() {
		t.Fatal("expected not null")
	}
	if cr.IsNotPointer() {
		t.Fatal("expected pointer")
	}
	if cr.IsNotMatchingAcceptedType() {
		t.Fatal("expected matching")
	}
	if !cr.IsSourceKind(reflect.String) {
		t.Fatal("expected string kind")
	}
	if cr.HasAnyIssues() {
		t.Fatal("expected no issues")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicStatus — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DynamicStatus_Invalid(t *testing.T) {
	ds := InvalidDynamicStatusNoMessage()
	if ds == nil {
		t.Fatal("expected non-nil")
	}
	ds2 := InvalidDynamicStatus("test msg")
	if ds2.Message != "test msg" {
		t.Fatal("expected msg")
	}
}

func Test_I16_DynamicStatus_Clone(t *testing.T) {
	ds := InvalidDynamicStatus("msg")
	cloned := ds.Clone()
	if cloned.Message != "msg" {
		t.Fatal("expected same msg")
	}
}

func Test_I16_DynamicStatus_ClonePtr(t *testing.T) {
	var ds *DynamicStatus
	if ds.ClonePtr() != nil {
		t.Fatal("expected nil for nil receiver")
	}
	ds = InvalidDynamicStatus("msg")
	cp := ds.ClonePtr()
	if cp == nil || cp.Message != "msg" {
		t.Fatal("expected cloned")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_LeftRight_NilReceiver(t *testing.T) {
	var lr *LeftRight
	if !lr.IsEmpty() {
		t.Fatal("expected empty for nil")
	}
	if lr.HasAnyItem() {
		t.Fatal("expected no items for nil")
	}
	if lr.HasLeft() {
		t.Fatal("expected no left for nil")
	}
	if lr.HasRight() {
		t.Fatal("expected no right for nil")
	}
	if !lr.IsLeftEmpty() {
		t.Fatal("expected left empty for nil")
	}
	if !lr.IsRightEmpty() {
		t.Fatal("expected right empty for nil")
	}
	if err := lr.LeftReflectSet(nil); err != nil {
		t.Fatal("expected nil err for nil receiver")
	}
	if err := lr.RightReflectSet(nil); err != nil {
		t.Fatal("expected nil err for nil receiver")
	}
	if lr.DeserializeLeft() != nil {
		t.Fatal("expected nil for nil")
	}
	if lr.DeserializeRight() != nil {
		t.Fatal("expected nil for nil")
	}
	if lr.LeftToDynamic() != nil {
		t.Fatal("expected nil for nil")
	}
	if lr.RightToDynamic() != nil {
		t.Fatal("expected nil for nil")
	}
	ts := lr.TypeStatus()
	_ = ts
}

func Test_I16_LeftRight_WithValues(t *testing.T) {
	lr := &LeftRight{Left: "a", Right: "b"}
	if lr.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	if !lr.HasAnyItem() {
		t.Fatal("expected has items")
	}
	if !lr.HasLeft() {
		t.Fatal("expected has left")
	}
	if !lr.HasRight() {
		t.Fatal("expected has right")
	}
	ld := lr.LeftToDynamic()
	if ld == nil {
		t.Fatal("expected non-nil")
	}
	rd := lr.RightToDynamic()
	if rd == nil {
		t.Fatal("expected non-nil")
	}
	dl := lr.DeserializeLeft()
	if dl == nil {
		t.Fatal("expected non-nil")
	}
	dr := lr.DeserializeRight()
	if dr == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// IsAnyTypesOf — branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_IsAnyTypesOf_Found(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	if !IsAnyTypesOf(strType, intType, strType) {
		t.Fatal("expected found")
	}
}

func Test_I16_IsAnyTypesOf_NotFound(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	if IsAnyTypesOf(strType, intType) {
		t.Fatal("expected not found")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeTypeName — nil branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_SafeTypeName_Nil(t *testing.T) {
	if SafeTypeName(nil) != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_I16_SafeTypeName_String(t *testing.T) {
	if SafeTypeName("hello") != "string" {
		t.Fatal("expected 'string'")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PointerOrNonPointer
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_PointerOrNonPointer_NonPtr(t *testing.T) {
	s := "hello"
	out, rv := PointerOrNonPointer(false, &s)
	_ = rv
	if out == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I16_PointerOrNonPointer_Ptr(t *testing.T) {
	s := "hello"
	out, rv := PointerOrNonPointer(true, &s)
	_ = rv
	if out == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionTypes — factory functions
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_CollectionTypes_Factories(t *testing.T) {
	if NewStringCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if EmptyStringCollection() == nil {
		t.Fatal("expected non-nil")
	}
	if NewIntCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if EmptyIntCollection() == nil {
		t.Fatal("expected non-nil")
	}
	if NewInt64Collection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if NewByteCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if NewBoolCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if NewFloat64Collection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if NewAnyMapCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
	if NewStringMapCollection(5) == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Dynamic — constructors and clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_Dynamic_InvalidDynamic(t *testing.T) {
	d := InvalidDynamic()
	if d.IsValid() {
		t.Fatal("expected invalid")
	}
}

func Test_I16_Dynamic_NewDynamicValid(t *testing.T) {
	d := NewDynamicValid("test")
	if !d.IsValid() {
		t.Fatal("expected valid")
	}
}

func Test_I16_Dynamic_Clone(t *testing.T) {
	d := NewDynamic("x", true)
	c := d.Clone()
	if !c.IsValid() {
		t.Fatal("expected valid clone")
	}
}

func Test_I16_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *Dynamic
	if d.ClonePtr() != nil {
		t.Fatal("expected nil for nil")
	}
}

func Test_I16_Dynamic_ClonePtr_Valid(t *testing.T) {
	d := NewDynamicPtr("x", true)
	cp := d.ClonePtr()
	if cp == nil || !cp.IsValid() {
		t.Fatal("expected valid clone")
	}
}

func Test_I16_Dynamic_NonPtrPtr(t *testing.T) {
	d := NewDynamic("x", true)
	np := d.NonPtr()
	if !np.IsValid() {
		t.Fatal("expected valid")
	}
	p := d.Ptr()
	if p == nil || !p.IsValid() {
		t.Fatal("expected valid ptr")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — basic branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KeyVal_Basics(t *testing.T) {
	kv := KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamic()
	if !d.IsValid() {
		t.Fatal("expected valid")
	}
	vd := kv.ValueDynamic()
	if !vd.IsValid() {
		t.Fatal("expected valid")
	}
	kdp := kv.KeyDynamicPtr()
	if kdp == nil {
		t.Fatal("expected non-nil")
	}
	vdp := kv.ValueDynamicPtr()
	if vdp == nil {
		t.Fatal("expected non-nil")
	}
	if kv.IsKeyNull() {
		t.Fatal("expected not null")
	}
	if kv.IsKeyNullOrEmptyString() {
		t.Fatal("expected not null/empty")
	}
	if kv.IsValueNull() {
		t.Fatal("expected not null")
	}
	s := kv.String()
	if s == "" {
		t.Fatal("expected non-empty string")
	}
	rv := kv.ValueReflectValue()
	if !rv.IsValid() {
		t.Fatal("expected valid reflect value")
	}
}

func Test_I16_KeyVal_NullKey(t *testing.T) {
	kv := KeyVal{Key: nil, Value: nil}
	if !kv.IsKeyNull() {
		t.Fatal("expected null key")
	}
	if !kv.IsValueNull() {
		t.Fatal("expected null value")
	}
}
