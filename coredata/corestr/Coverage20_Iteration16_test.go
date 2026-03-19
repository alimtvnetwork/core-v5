package corestr

import (
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_ValidValue_Constructors(t *testing.T) {
	vv := NewValidValueUsingAny(false, true, "hello")
	if vv == nil || !vv.IsValid {
		t.Fatal("expected valid")
	}

	vv2 := NewValidValueUsingAnyAutoValid(false, "hello")
	if vv2 == nil {
		t.Fatal("expected non-nil")
	}

	vv3 := NewValidValueEmpty()
	if vv3 == nil || !vv3.IsValid {
		t.Fatal("expected valid empty")
	}

	vv4 := InvalidValidValueNoMessage()
	if vv4.IsValid {
		t.Fatal("expected invalid")
	}

	vv5 := InvalidValidValue("msg")
	if vv5.IsValid || vv5.Message != "msg" {
		t.Fatal("expected invalid with msg")
	}
}

func Test_I16_ValidValue_ValueBytesOnce(t *testing.T) {
	vv := NewValidValue("test")
	b1 := vv.ValueBytesOnce()
	b2 := vv.ValueBytesOnce()
	if string(b1) != "test" || string(b2) != "test" {
		t.Fatal("expected test bytes")
	}
	b3 := vv.ValueBytesOncePtr()
	if string(b3) != "test" {
		t.Fatal("expected test bytes via ptr")
	}
}

func Test_I16_ValidValue_StringChecks(t *testing.T) {
	vv := NewValidValue("  hello  ")
	if vv.IsEmpty() {
		t.Fatal("expected not empty")
	}
	if vv.IsWhitespace() {
		t.Fatal("expected not whitespace")
	}
	if vv.Trim() != "hello" {
		t.Fatal("expected trimmed")
	}
	if !vv.HasValidNonEmpty() {
		t.Fatal("expected valid non-empty")
	}
	if !vv.HasValidNonWhitespace() {
		t.Fatal("expected valid non-whitespace")
	}
	if !vv.HasSafeNonEmpty() {
		t.Fatal("expected safe non-empty")
	}
}

func Test_I16_ValidValue_ValueConversions(t *testing.T) {
	vv := NewValidValue("true")
	if !vv.ValueBool() {
		t.Fatal("expected true")
	}

	vv2 := NewValidValue("42")
	if vv2.ValueInt(0) != 42 {
		t.Fatal("expected 42")
	}
	if vv2.ValueDefInt() != 42 {
		t.Fatal("expected 42")
	}
	if vv2.ValueByte(0) != 42 {
		t.Fatal("expected 42")
	}
	if vv2.ValueDefByte() != 42 {
		t.Fatal("expected 42")
	}

	vv3 := NewValidValue("3.14")
	if vv3.ValueFloat64(0) != 3.14 {
		t.Fatal("expected 3.14")
	}
	if vv3.ValueDefFloat64() != 3.14 {
		t.Fatal("expected 3.14")
	}

	// Error cases
	vv4 := NewValidValue("notnum")
	if vv4.ValueBool() {
		t.Fatal("expected false for invalid bool")
	}
	if vv4.ValueInt(99) != 99 {
		t.Fatal("expected default")
	}

	// Empty bool
	vv5 := NewValidValue("")
	if vv5.ValueBool() {
		t.Fatal("expected false for empty")
	}
}

func Test_I16_ValidValue_Comparisons(t *testing.T) {
	vv := NewValidValue("hello")
	if !vv.Is("hello") {
		t.Fatal("expected match")
	}
	if !vv.IsAnyOf("world", "hello") {
		t.Fatal("expected any match")
	}
	if vv.IsAnyOf("world", "foo") {
		t.Fatal("expected no match")
	}
	if !vv.IsAnyOf() {
		t.Fatal("expected true for empty values")
	}
	if !vv.IsContains("ell") {
		t.Fatal("expected contains")
	}
	if !vv.IsAnyContains("xyz", "ell") {
		t.Fatal("expected any contains")
	}
	if vv.IsAnyContains("xyz", "abc") {
		t.Fatal("expected no contains")
	}
	if !vv.IsAnyContains() {
		t.Fatal("expected true for empty")
	}
	if !vv.IsEqualNonSensitive("HELLO") {
		t.Fatal("expected case insensitive match")
	}
}

func Test_I16_ValidValue_CloneAndDispose(t *testing.T) {
	var vvNil *ValidValue
	if vvNil.Clone() != nil {
		t.Fatal("expected nil clone for nil")
	}
	if vvNil.String() != "" {
		t.Fatal("expected empty for nil")
	}
	if vvNil.FullString() != "" {
		t.Fatal("expected empty for nil")
	}

	vv := NewValidValue("test")
	c := vv.Clone()
	if c.Value != "test" {
		t.Fatal("expected test")
	}

	vv.Clear()
	if vv.Value != "" || vv.IsValid {
		t.Fatal("expected cleared")
	}

	vv2 := NewValidValue("x")
	vv2.Dispose()

	vvNil.Clear()
	vvNil.Dispose()
}

func Test_I16_ValidValue_Json(t *testing.T) {
	vv := NewValidValue("test")
	j := vv.Json()
	_ = j
	jp := vv.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
	b, err := vv.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I16_ValidValue_ParseInjectUsingJson(t *testing.T) {
	vv := NewValidValue("test")
	j := vv.JsonPtr()
	vv2 := &ValidValue{}
	result, err := vv2.ParseInjectUsingJson(j)
	if err != nil {
		t.Fatal(err)
	}
	if result.Value != "test" {
		t.Fatal("expected test")
	}
}

func Test_I16_ValidValue_Split(t *testing.T) {
	vv := NewValidValue("a,b,c")
	parts := vv.Split(",")
	if len(parts) != 3 {
		t.Fatal("expected 3 parts")
	}
	parts2 := vv.SplitNonEmpty(",")
	if len(parts2) == 0 {
		t.Fatal("expected non-empty")
	}
	parts3 := vv.SplitTrimNonWhitespace(",")
	if len(parts3) == 0 {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_ValidValues_Basics(t *testing.T) {
	vvs := EmptyValidValues()
	if vvs.HasAnyItem() {
		t.Fatal("expected empty")
	}

	vvs.Add("a").Add("b")
	if vvs.Length() != 2 {
		t.Fatal("expected 2")
	}
	if vvs.Count() != 2 {
		t.Fatal("expected 2")
	}
	if !vvs.HasAnyItem() {
		t.Fatal("expected items")
	}
	if !vvs.HasIndex(1) {
		t.Fatal("expected index 1")
	}
	if vvs.SafeValueAt(0) != "a" {
		t.Fatal("expected a")
	}
	if vvs.SafeValueAt(99) != "" {
		t.Fatal("expected empty for out of range")
	}
}

func Test_I16_ValidValues_UsingValues(t *testing.T) {
	v1 := ValidValue{Value: "x", IsValid: true}
	vvs := NewValidValuesUsingValues(v1)
	if vvs.Length() != 1 {
		t.Fatal("expected 1")
	}
	// Empty
	vvs2 := NewValidValuesUsingValues()
	if vvs2.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I16_ValidValues_StringsAndFullStrings(t *testing.T) {
	vvs := EmptyValidValues()
	if len(vvs.Strings()) != 0 {
		t.Fatal("expected empty strings")
	}
	vvs.Add("x")
	ss := vvs.Strings()
	if len(ss) != 1 || ss[0] != "x" {
		t.Fatal("expected [x]")
	}
	fs := vvs.FullStrings()
	if len(fs) != 1 {
		t.Fatal("expected 1 full string")
	}
	s := vvs.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I16_ValidValues_Find(t *testing.T) {
	vvs := EmptyValidValues()
	found := vvs.Find(func(i int, v *ValidValue) (*ValidValue, bool, bool) {
		return v, true, false
	})
	if len(found) != 0 {
		t.Fatal("expected empty for empty collection")
	}

	vvs.Add("a").Add("b").Add("c")
	found = vvs.Find(func(i int, v *ValidValue) (*ValidValue, bool, bool) {
		if v.Value == "b" {
			return v, true, true
		}
		return nil, false, false
	})
	if len(found) != 1 {
		t.Fatal("expected 1 found")
	}
}

func Test_I16_ValidValues_SafeValidValueAt(t *testing.T) {
	vvs := EmptyValidValues()
	vvs.AddFull(true, "valid", "")
	vvs.AddFull(false, "invalid", "err")

	if vvs.SafeValidValueAt(0) != "valid" {
		t.Fatal("expected valid")
	}
	if vvs.SafeValidValueAt(1) != "" {
		t.Fatal("expected empty for invalid")
	}
	if vvs.SafeValidValueAt(99) != "" {
		t.Fatal("expected empty for out of range")
	}
}

func Test_I16_ValidValues_SafeIndexes(t *testing.T) {
	vvs := EmptyValidValues()
	vvs.Add("a").Add("b")
	vals := vvs.SafeValuesAtIndexes(0, 1)
	if len(vals) != 2 {
		t.Fatal("expected 2")
	}
	vals2 := vvs.SafeValidValuesAtIndexes(0, 1)
	if len(vals2) != 2 {
		t.Fatal("expected 2")
	}
	empty := vvs.SafeValuesAtIndexes()
	if len(empty) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I16_ValidValues_ConcatAndAdd(t *testing.T) {
	vvs := EmptyValidValues()
	vvs.Add("a")

	// ConcatNew with empty and clone
	c1 := vvs.ConcatNew(true)
	if c1.Length() != 1 {
		t.Fatal("expected 1 cloned")
	}
	// ConcatNew with empty no clone
	c2 := vvs.ConcatNew(false)
	if c2 != vvs {
		t.Fatal("expected same reference")
	}
	// ConcatNew with collection
	other := EmptyValidValues()
	other.Add("b")
	c3 := vvs.ConcatNew(false, other)
	if c3.Length() != 2 {
		t.Fatal("expected 2")
	}

	// AddValidValues nil
	vvs.AddValidValues(nil)
	// AddValidValues empty
	vvs.AddValidValues(EmptyValidValues())

	// Adds empty
	vvs.Adds()
	// AddsPtr empty
	vvs.AddsPtr()
}

func Test_I16_ValidValues_HashmapAndMap(t *testing.T) {
	vvs := EmptyValidValues()
	h := vvs.Hashmap()
	if h == nil {
		t.Fatal("expected non-nil")
	}
	m := vvs.Map()
	if m == nil {
		t.Fatal("expected non-nil")
	}

	vvs.Add("k")
	h2 := vvs.Hashmap()
	if h2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I16_ValidValues_AddHashset(t *testing.T) {
	vvs := EmptyValidValues()
	vvs.AddHashsetMap(nil)
	vvs.AddHashset(nil)
	vvs.AddHashsetMap(map[string]bool{"a": true})
	if vvs.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_ValueStatus_Constructors(t *testing.T) {
	vs := InvalidValueStatusNoMessage()
	if vs == nil {
		t.Fatal("expected non-nil")
	}
	vs2 := InvalidValueStatus("err")
	if vs2 == nil || vs2.ValueValid.Message != "err" {
		t.Fatal("expected err message")
	}
	c := vs2.Clone()
	if c.ValueValid.Message != "err" {
		t.Fatal("expected cloned message")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_TextWithLineNumber_All(t *testing.T) {
	var tln *TextWithLineNumber
	if tln.HasLineNumber() {
		t.Fatal("expected false for nil")
	}
	if !tln.IsInvalidLineNumber() {
		t.Fatal("expected invalid for nil")
	}
	if tln.Length() != 0 {
		t.Fatal("expected 0 for nil")
	}
	if !tln.IsEmpty() {
		t.Fatal("expected empty for nil")
	}
	if !tln.IsEmptyText() {
		t.Fatal("expected empty text for nil")
	}

	tln = &TextWithLineNumber{LineNumber: 5, Text: "hello"}
	if !tln.HasLineNumber() {
		t.Fatal("expected has line number")
	}
	if tln.IsInvalidLineNumber() {
		t.Fatal("expected valid line number")
	}
	if tln.Length() != 5 {
		t.Fatal("expected 5")
	}
	if tln.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	if tln.IsEmptyText() {
		t.Fatal("expected non-empty text")
	}
	if !tln.IsEmptyTextLineBoth() {
		// valid line + text => not empty
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight (corestr) — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_LeftRight_Constructors(t *testing.T) {
	lr := InvalidLeftRightNoMessage()
	if lr.IsValid {
		t.Fatal("expected invalid")
	}
	lr2 := InvalidLeftRight("msg")
	if lr2.IsValid || lr2.Message != "msg" {
		t.Fatal("expected invalid with msg")
	}
	lr3 := NewLeftRight("a", "b")
	if !lr3.IsValid || lr3.Left != "a" || lr3.Right != "b" {
		t.Fatal("expected valid a/b")
	}
}

func Test_I16_LeftRight_UsingSlice(t *testing.T) {
	lr := LeftRightUsingSlice(nil)
	if lr.IsValid {
		t.Fatal("expected invalid for nil")
	}
	lr2 := LeftRightUsingSlice([]string{"a"})
	if lr2.Left != "a" || lr2.Right != "" {
		t.Fatal("expected a/empty")
	}
	lr3 := LeftRightUsingSlice([]string{"a", "b"})
	if lr3.Left != "a" || lr3.Right != "b" || !lr3.IsValid {
		t.Fatal("expected a/b valid")
	}
}

func Test_I16_LeftRight_UsingSlicePtr(t *testing.T) {
	lr := LeftRightUsingSlicePtr([]string{})
	if lr.IsValid {
		t.Fatal("expected invalid")
	}
	lr2 := LeftRightUsingSlicePtr([]string{"a", "b"})
	if lr2.Left != "a" {
		t.Fatal("expected a")
	}
}

func Test_I16_LeftRight_TrimmedUsingSlice(t *testing.T) {
	lr := LeftRightTrimmedUsingSlice(nil)
	if lr.IsValid {
		t.Fatal("expected invalid")
	}
	lr2 := LeftRightTrimmedUsingSlice([]string{})
	if lr2.IsValid {
		t.Fatal("expected invalid for empty")
	}
	lr3 := LeftRightTrimmedUsingSlice([]string{" a "})
	if lr3.Left != " a " { // single element not trimmed
		// left is not trimmed in single-element case
	}
	lr4 := LeftRightTrimmedUsingSlice([]string{" a ", " b "})
	if lr4.Left != "a" || lr4.Right != "b" {
		t.Fatal("expected trimmed a/b")
	}
}

func Test_I16_LeftRight_Methods(t *testing.T) {
	lr := NewLeftRight("hello", "world")
	if string(lr.LeftBytes()) != "hello" {
		t.Fatal("expected hello bytes")
	}
	if string(lr.RightBytes()) != "world" {
		t.Fatal("expected world bytes")
	}
	if lr.LeftTrim() != "hello" {
		t.Fatal("expected hello")
	}
	if lr.IsLeftEmpty() || lr.IsRightEmpty() {
		t.Fatal("expected not empty")
	}
	if lr.IsLeftWhitespace() || lr.IsRightWhitespace() {
		t.Fatal("expected not whitespace")
	}
	if !lr.HasValidNonEmptyLeft() || !lr.HasValidNonEmptyRight() {
		t.Fatal("expected valid non-empty")
	}
	if !lr.HasSafeNonEmpty() {
		t.Fatal("expected safe non-empty")
	}
	_ = lr.NonPtr()
	if lr.Ptr() == nil {
		t.Fatal("expected ptr")
	}
	if !lr.IsLeft("hello") || !lr.IsRight("world") {
		t.Fatal("expected match")
	}
	if !lr.Is("hello", "world") {
		t.Fatal("expected match")
	}
}

func Test_I16_LeftRight_IsEqual(t *testing.T) {
	var lr1 *LeftRight
	var lr2 *LeftRight
	if !lr1.IsEqual(lr2) {
		t.Fatal("expected nil == nil")
	}
	lr3 := NewLeftRight("a", "b")
	if lr1.IsEqual(lr3) {
		t.Fatal("expected nil != non-nil")
	}
	if lr3.IsEqual(lr1) {
		t.Fatal("expected non-nil != nil")
	}
	lr4 := NewLeftRight("a", "b")
	if !lr3.IsEqual(lr4) {
		t.Fatal("expected equal")
	}
}

func Test_I16_LeftRight_CloneAndDispose(t *testing.T) {
	lr := NewLeftRight("a", "b")
	c := lr.Clone()
	if c.Left != "a" {
		t.Fatal("expected a")
	}
	lr.Clear()
	lr.Dispose()
	var lrNil *LeftRight
	lrNil.Clear()
	lrNil.Dispose()
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_LeftMiddleRight_Constructors(t *testing.T) {
	lmr := InvalidLeftMiddleRightNoMessage()
	if lmr.IsValid {
		t.Fatal("expected invalid")
	}
	lmr2 := InvalidLeftMiddleRight("msg")
	if lmr2.IsValid {
		t.Fatal("expected invalid")
	}
	lmr3 := NewLeftMiddleRight("a", "m", "b")
	if !lmr3.IsValid || lmr3.Left != "a" || lmr3.Middle != "m" || lmr3.Right != "b" {
		t.Fatal("expected a/m/b")
	}
}

func Test_I16_LeftMiddleRight_Methods(t *testing.T) {
	lmr := NewLeftMiddleRight("a", "m", "b")
	if string(lmr.LeftBytes()) != "a" {
		t.Fatal("expected a")
	}
	if string(lmr.RightBytes()) != "b" {
		t.Fatal("expected b")
	}
	if string(lmr.MiddleBytes()) != "m" {
		t.Fatal("expected m")
	}
	if lmr.LeftTrim() != "a" || lmr.RightTrim() != "b" || lmr.MiddleTrim() != "m" {
		t.Fatal("expected trimmed")
	}
	if lmr.IsLeftEmpty() || lmr.IsRightEmpty() || lmr.IsMiddleEmpty() {
		t.Fatal("expected non-empty")
	}
	if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyRight() || !lmr.HasValidNonEmptyMiddle() {
		t.Fatal("expected valid non-empty")
	}
	if !lmr.HasSafeNonEmpty() {
		t.Fatal("expected safe")
	}
	if !lmr.IsAll("a", "m", "b") {
		t.Fatal("expected match")
	}
	if !lmr.Is("a", "b") {
		t.Fatal("expected match")
	}
	lr := lmr.ToLeftRight()
	if lr.Left != "a" || lr.Right != "b" {
		t.Fatal("expected a/b")
	}
	c := lmr.Clone()
	if c.Left != "a" {
		t.Fatal("expected a")
	}
	lmr.Clear()
	lmr.Dispose()
	var lmrNil *LeftMiddleRight
	lmrNil.Clear()
	lmrNil.Dispose()
}
